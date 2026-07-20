import { useState, useRef, useEffect } from "react"

interface Props {
  videoPath: string
  videoTitle: string
  onBack: () => void
}

export default function EditorPage({ videoPath, videoTitle, onBack }: Props) {
  const videoRef = useRef<HTMLVideoElement>(null)
  const [startTime, setStartTime] = useState(0)
  const [endTime, setEndTime] = useState(100)
  const [duration, setDuration] = useState(100)
  const [encoder, setEncoder] = useState("auto")
  const [exporting, setExporting] = useState(false)
  const [exportProgress, setExportProgress] = useState(0)
  const [exportDone, setExportDone] = useState(false)
  const [encoders, setEncoders] = useState<{ name: string; value: string }[]>([])

  useEffect(() => {
    if (window.GoApp?.GetAvailableEncoders) {
      window.GoApp.GetAvailableEncoders().then((list: any[]) => {
        const available = list.filter((e: any) => e.available)
        setEncoders(available.map((e: any) => ({ name: e.name, value: e.value })))
      }).catch(console.error)
    }
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

  const handleExport = async () => {
    setExporting(true)
    setExportDone(false)
    try {
      if (window.GoApp?.ExportFile) {
        await window.GoApp.ExportFile(videoPath, encoder, "mp4")
        setExportDone(true)
      }
    } catch (err) {
      console.error("Export failed:", err)
    } finally {
      setExporting(false)
    }
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
      {/* Back button */}
      <button onClick={onBack} className="text-sm text-indigo-400 hover:text-indigo-300 transition-colors">
        &larr; Back to Home
      </button>

      <h1 className="text-xl font-bold text-white truncate">{videoTitle}</h1>

      {/* Video Preview */}
      <div className="bg-black rounded-lg overflow-hidden">
        <video
          ref={videoRef}
          src={`file://${videoPath}`}
          onLoadedMetadata={handleLoadedMetadata}
          controls
          className="w-full max-h-[400px]"
        />
      </div>

      {/* Timeline / Clip Range */}
      <div className="bg-gray-800/50 rounded-lg p-4 border border-gray-700/50 space-y-3">
        <h2 className="text-sm font-medium text-gray-300">Clip Range</h2>

        <div className="flex gap-4 text-sm">
          <div className="flex items-center gap-2">
            <label className="text-gray-500">Start:</label>
            <span className="text-white font-mono">{formatTime(startTime)}</span>
            <input
              type="range"
              min={0}
              max={duration}
              step={0.5}
              value={startTime}
              onChange={(e) => {
                const val = parseFloat(e.target.value)
                if (val < endTime) setStartTime(val)
              }}
              className="w-24 accent-indigo-500"
            />
          </div>
          <div className="flex items-center gap-2">
            <label className="text-gray-500">End:</label>
            <span className="text-white font-mono">{formatTime(endTime)}</span>
            <input
              type="range"
              min={0}
              max={duration}
              step={0.5}
              value={endTime}
              onChange={(e) => {
                const val = parseFloat(e.target.value)
                if (val > startTime) setEndTime(val)
              }}
              className="w-24 accent-indigo-500"
            />
          </div>
        </div>

        {/* Preview button */}
        <button
          onClick={handlePreview}
          className="px-4 py-1.5 bg-gray-700 hover:bg-gray-600 text-white text-sm rounded-md transition-colors"
        >
          ▶ Preview Clip
        </button>
      </div>

      {/* Export Panel */}
      <div className="bg-gray-800/50 rounded-lg p-4 border border-gray-700/50 space-y-3">
        <h2 className="text-sm font-medium text-gray-300">Export</h2>

        <div className="flex gap-4 items-end">
          <div className="space-y-1">
            <label className="text-xs text-gray-500">Encoder</label>
            <select
              value={encoder}
              onChange={(e) => setEncoder(e.target.value)}
              className="px-3 py-1.5 bg-gray-800 border border-gray-700 rounded-md text-white text-sm focus:outline-none focus:border-indigo-500"
            >
              <option value="auto">Auto</option>
              {encoders.map((e) => (
                <option key={e.value} value={e.value}>{e.name}</option>
              ))}
              <option value="libx264">CPU (libx264)</option>
            </select>
          </div>

          <div className="space-y-1">
            <label className="text-xs text-gray-500">Format</label>
            <select
              defaultValue="mp4"
              className="px-3 py-1.5 bg-gray-800 border border-gray-700 rounded-md text-white text-sm focus:outline-none focus:border-indigo-500"
            >
              <option value="mp4">MP4</option>
            </select>
          </div>

          <button
            onClick={handleExport}
            disabled={exporting}
            className="px-6 py-1.5 bg-indigo-600 hover:bg-indigo-500 disabled:bg-gray-700 text-white text-sm font-medium rounded-md transition-colors"
          >
            {exporting ? "Exporting..." : "Export Clip"}
          </button>
        </div>

        {/* Progress bar */}
        {exporting && (
          <div className="space-y-1">
            <div className="w-full bg-gray-700 rounded-full h-2">
              <div
                className="bg-indigo-500 h-2 rounded-full transition-all"
                style={{ width: `${Math.min(exportProgress, 100)}%` }}
              />
            </div>
            <p className="text-xs text-gray-500">{exportProgress.toFixed(1)}%</p>
          </div>
        )}

        {exportDone && (
          <div className="text-sm text-green-400">
            ✓ Export completed! Check your output folder.
          </div>
        )}
      </div>
    </div>
  )
}
