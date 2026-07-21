import { useState, useRef, useEffect, useCallback, memo } from "react"
import { Backend } from "../services/backend"
import type { ExportPreset, EncoderOption } from "../types"

interface Props {
  videoPath: string
  videoTitle: string
  onBack: () => void
  onExportDone?: () => void
}

// Convert a local file path to a proper file:// URL for the video element.
// Windows: C:\path\to\file.mp4 → file:///C:/path/to/file.mp4
// Unix:    /path/to/file.mp4  → file:///path/to/file.mp4
function toFileUrl(path: string): string {
  if (!path) return ""
  // Already a URL
  if (path.startsWith("file://") || path.startsWith("http://") || path.startsWith("https://")) return path
  // Windows path: replace backslashes and prepend file:///
  const normalized = path.replace(/\\/g, "/")
  return normalized.startsWith("/") ? `file://${normalized}` : `file:///${normalized}`
}

const EditorPage = memo(function EditorPage({ videoPath, videoTitle, onBack, onExportDone }: Props) {
  const videoRef = useRef<HTMLVideoElement>(null)
  const [startTime, setStartTime] = useState(0)
  const [endTime, setEndTime] = useState(100)
  const [duration, setDuration] = useState(100)
  const [encoder, setEncoder] = useState("auto")
  const [availableEncoders, setAvailableEncoders] = useState<EncoderOption[]>([])
  const [exporting, setExporting] = useState(false)
  const [exportProgress, setExportProgress] = useState(0)
  const [exportDone, setExportDone] = useState(false)
  const [exportError, setExportError] = useState("")
  const [presets, setPresets] = useState<ExportPreset[]>([])
  const [activePreset, setActivePreset] = useState("")
  const [notification, setNotification] = useState("")

  // --- Functions used by keyboard shortcuts must be defined first ---

  const showNotif = (msg: string) => {
    setNotification(msg)
    setTimeout(() => setNotification(""), 4000)
  }

  const handleExport = async () => {
    setExporting(true)
    setExportDone(false)
    setExportError("")
    setExportProgress(0)

    const progressTimer = setInterval(() => {
      setExportProgress(prev => Math.min(prev + 5, 90))
    }, 800)

    try {
      await Backend.ExportFile(videoPath, encoder, "mp4")
      setExportProgress(100)
      setExportDone(true)
      showNotif("✓ Export completed!")
      Backend.ShowNotification("My Clip", "Export completed successfully!").catch(() => {})
      onExportDone?.()
      const dir = await Backend.GetClipOutputDir()
      Backend.OpenFolder(dir).catch(() => {})
    } catch (err: any) {
      const msg = err?.message || "Export failed. Check that FFmpeg is available and the file is valid."
      setExportError(msg)
      showNotif("✗ " + msg)
    } finally {
      clearInterval(progressTimer)
      setExporting(false)
    }
  }

  // Keyboard shortcuts
  const handleKeyDown = useCallback((e: KeyboardEvent) => {
    if (!videoRef.current) return
    switch (e.code) {
      case "Space":
        e.preventDefault()
        videoRef.current.paused ? videoRef.current.play() : videoRef.current.pause()
        break
      case "KeyI":
        e.preventDefault()
        setStartTime(videoRef.current.currentTime)
        break
      case "KeyO":
        e.preventDefault()
        setEndTime(videoRef.current.currentTime)
        break
      case "Enter":
        if (e.ctrlKey || e.metaKey) { e.preventDefault(); handleExport() }
        break
      case "Escape":
        if (exporting) return; onBack()
        break
    }
  }, [startTime, endTime, exporting, handleExport, onBack])

  useEffect(() => {
    window.addEventListener("keydown", handleKeyDown)
    return () => window.removeEventListener("keydown", handleKeyDown)
  }, [handleKeyDown])

  // Load presets and available encoders on mount
  useEffect(() => {
    Backend.GetPresets().then(setPresets).catch(console.error)
    Backend.GetAvailableEncoders().then(setAvailableEncoders).catch(console.error)
  }, [])

  const handleLoadedMetadata = () => {
    if (videoRef.current) {
      setDuration(videoRef.current.duration)
      setEndTime(videoRef.current.duration)
    }
  }

  const handlePreview = () => {
    if (!videoRef.current) return
    videoRef.current.currentTime = startTime
    videoRef.current.play()
  }

  const handlePresetChange = (name: string) => {
    setActivePreset(name)
    const preset = presets.find(p => p.name === name)
    if (preset) setEncoder(preset.encoder === "auto" ? "auto" : preset.encoder)
  }

  const handleCopyPath = async () => {
    await Backend.CopyPathToClipboard(videoPath)
    showNotif("📋 Path copied to clipboard!")
  }

  const handleOpenFolder = async () => {
    const dir = await Backend.GetOutputDir()
    Backend.OpenFolder(dir).catch(() => {})
  }

  const formatTime = (s: number) => {
    const h = Math.floor(s / 3600)
    const m = Math.floor((s % 3600) / 60)
    const sec = Math.floor(s % 60)
    if (h > 0) return `${h}:${String(m).padStart(2, "0")}:${String(sec).padStart(2, "0")}`
    return `${m}:${String(sec).padStart(2, "0")}`
  }

  return (
    <div className="max-w-4xl mx-auto py-6 px-4 space-y-6">
      {/* Toast notification */}
      {notification && (
        <div className="fixed top-4 right-4 z-50 bg-gray-800 border border-gray-700 rounded-lg px-4 py-2 text-sm text-white shadow-lg animate-pulse">
          {notification}
        </div>
      )}

      {/* Back button */}
      <button onClick={onBack} className="text-sm text-indigo-400 hover:text-indigo-300 transition-colors">
        &larr; Back to Home
      </button>

      <div className="flex items-center justify-between">
        <h1 className="text-xl font-bold text-white truncate">{videoTitle}</h1>
        <div className="flex gap-2">
          <button onClick={handleCopyPath} className="text-xs text-gray-400 hover:text-white transition-colors">📋 Copy Path</button>
          <button onClick={handleOpenFolder} className="text-xs text-gray-400 hover:text-white transition-colors">📂 Open Folder</button>
        </div>
      </div>

      {/* Video Preview */}
      <div className="bg-black rounded-lg overflow-hidden">
        <video
          ref={videoRef}
          src={toFileUrl(videoPath)}
          onLoadedMetadata={handleLoadedMetadata}
          className="w-full max-h-[400px]"
          controls
        />
      </div>

      {/* Timeline / Clip Range */}
      <div className="bg-gray-800/50 rounded-lg p-4 border border-gray-700/50 space-y-3">
        <div className="flex items-center justify-between">
          <h2 className="text-sm font-medium text-gray-300">Clip Range</h2>
          <span className="text-xs text-gray-600">I=In point &middot; O=Out point &middot; Space=Play</span>
        </div>

        <div className="flex gap-4 text-sm flex-wrap">
          <div className="flex items-center gap-2">
            <label className="text-gray-500">In (I):</label>
            <span className="text-white font-mono w-20">{formatTime(startTime)}</span>
            <input type="range" min={0} max={duration} step={0.5} value={startTime}
              onChange={(e) => { const v = parseFloat(e.target.value); if (v < endTime) setStartTime(v) }}
              className="w-24 accent-indigo-500" />
          </div>
          <div className="flex items-center gap-2">
            <label className="text-gray-500">Out (O):</label>
            <span className="text-white font-mono w-20">{formatTime(endTime)}</span>
            <input type="range" min={0} max={duration} step={0.5} value={endTime}
              onChange={(e) => { const v = parseFloat(e.target.value); if (v > startTime) setEndTime(v) }}
              className="w-24 accent-indigo-500" />
          </div>
          <button onClick={handlePreview}
            className="px-4 py-1.5 bg-gray-700 hover:bg-gray-600 text-white text-sm rounded-md transition-colors">
            ▶ Preview
          </button>
        </div>
      </div>

      {/* Export Panel */}
      <div className="bg-gray-800/50 rounded-lg p-4 border border-gray-700/50 space-y-3">
        <div className="flex items-center justify-between">
          <h2 className="text-sm font-medium text-gray-300">Export</h2>
          <span className="text-xs text-gray-600">Ctrl+Enter to export</span>
        </div>

        {/* Presets */}
        {presets.length > 0 && (
          <div className="flex gap-2 flex-wrap">
            {presets.map(p => (
              <button key={p.name} onClick={() => handlePresetChange(p.name)}
                className={`px-3 py-1 text-xs rounded-md transition-colors ${
                  activePreset === p.name
                    ? "bg-indigo-600 text-white"
                    : "bg-gray-700 text-gray-300 hover:bg-gray-600"
                }`}>
                {p.name}
              </button>
            ))}
          </div>
        )}

        <div className="flex gap-4 items-end flex-wrap">
          <div className="space-y-1">
            <label className="text-xs text-gray-500">Encoder</label>
            <select value={encoder} onChange={(e) => { setEncoder(e.target.value); setActivePreset("") }}
              className="px-3 py-1.5 bg-gray-800 border border-gray-700 rounded-md text-white text-sm focus:outline-none focus:border-indigo-500">
              <option value="auto">Auto (recommended)</option>
              {availableEncoders.filter(e => e.available).map(e => (
                <option key={e.value} value={e.value}>{e.name}</option>
              ))}
            </select>
          </div>

          <button onClick={handleExport} disabled={exporting}
            className="px-6 py-1.5 bg-indigo-600 hover:bg-indigo-500 disabled:bg-gray-700 text-white text-sm font-medium rounded-md transition-colors">
            {exporting ? "Exporting..." : "Export Clip"}
          </button>
        </div>

        {/* Progress bar */}
        {exporting && (
          <div className="space-y-1">
            <div className="w-full bg-gray-700 rounded-full h-2">
              <div className="bg-indigo-500 h-2 rounded-full transition-all" style={{ width: `${Math.min(exportProgress, 100)}%` }} />
            </div>
            <p className="text-xs text-gray-500">{exportProgress.toFixed(1)}%</p>
          </div>
        )}

        {exportDone && <div className="text-sm text-green-400">✓ Export completed!</div>}
        {exportError && <div className="text-sm text-red-400">{exportError}</div>}
      </div>
    </div>
  )
})

export default EditorPage
