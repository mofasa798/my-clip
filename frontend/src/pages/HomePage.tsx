import { useState } from "react"
import DepStatusPanel from "../components/DepStatusPanel"
import type { DepResult, VideoMetadata, DownloadResult } from "../types"

declare global {
  interface Window {
    GoApp?: {
      GetVersion(): Promise<string>
      GetDependencies(): Promise<DepResult>
      RefreshDependencies(): Promise<DepResult>
      GetConfig(): Promise<Record<string, string>>
      SaveConfig(cfg: Record<string, string>): Promise<void>
      SupportedSources(): Promise<string[]>
      ResolveSource(url: string): Promise<string>
      GetMetadata(url: string): Promise<VideoMetadata>
      StartDownload(url: string, streamId: string): Promise<DownloadResult>
      ProbeFile(path: string): Promise<any>
      GetHistory(): Promise<any[]>
    }
  }
}

interface Props {
  deps: DepResult | null
  onRefreshDeps: () => void
  onNavigateEditor: (filePath: string, title: string) => void
}

export default function HomePage({ deps, onRefreshDeps, onNavigateEditor }: Props) {
  const [url, setUrl] = useState("")
  const [version, setVersion] = useState("")
  const [loading, setLoading] = useState(false)
  const [downloading, setDownloading] = useState(false)
  const [downloadProgress, setDownloadProgress] = useState(0)
  const [metadata, setMetadata] = useState<VideoMetadata | null>(null)
  const [selectedStream, setSelectedStream] = useState("")
  const [downloadedFile, setDownloadedFile] = useState("")
  const [error, setError] = useState("")
  const [supportedSources, setSupportedSources] = useState<string[]>([])

  useState(() => {
    if (window.GoApp) {
      window.GoApp.GetVersion().then(setVersion).catch(console.error)
      window.GoApp.SupportedSources().then(setSupportedSources).catch(console.error)
    }
  })

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!url.trim()) return

    setLoading(true)
    setError("")
    setMetadata(null)
    setDownloadedFile("")
    setDownloadProgress(0)

    try {
      if (window.GoApp) {
        const source = await window.GoApp.ResolveSource(url)
        if (!source) {
          setError("URL not supported. Supported sources: YouTube (youtube.com, youtu.be), Kick (kick.com)")
          setLoading(false)
          return
        }
        const meta = await window.GoApp.GetMetadata(url)
        setMetadata(meta)
        if (meta.streams.length > 0) {
          const best = meta.streams.filter(s => s.has_video)[0]
          if (best) setSelectedStream(best.id)
        }
      }
    } catch (err: any) {
      const msg = err?.message || ""
      if (msg.includes("no video") || msg.includes("video unavailable")) {
        setError("Video unavailable. It may be private, age-restricted, or removed.")
      } else if (msg.includes("find in PATH") || msg.includes("yt-dlp")) {
        setError("yt-dlp is not installed or not found in PATH. Please install yt-dlp and try again.")
      } else if (msg.includes("network") || msg.includes("timeout") || msg.includes("connection")) {
        setError("Network error. Check your internet connection and try again.")
      } else {
        setError(`Failed to load video: ${msg}`)
      }
    } finally {
      setLoading(false)
    }
  }

  const handleDownload = async () => {
    if (!metadata || !selectedStream) return
    setDownloading(true)
    setDownloadProgress(0)
    setError("")

    // Retry logic: up to 3 attempts
    let lastError = ""
    for (let attempt = 1; attempt <= 3; attempt++) {
      try {
        if (window.GoApp) {
          const result = await window.GoApp.StartDownload(url, selectedStream)
          setDownloadProgress(100)
          setDownloadedFile(result.file_path)
          return
        }
      } catch (err: any) {
        lastError = err?.message || "Download failed."
        if (attempt < 3) {
          await new Promise(r => setTimeout(r, 2000 * attempt))
        }
      }
    }
    setError(`Download failed after 3 attempts: ${lastError}`)
  }

  const formatDuration = (seconds: number): string => {
    const h = Math.floor(seconds / 3600)
    const m = Math.floor((seconds % 3600) / 60)
    const s = Math.floor(seconds % 60)
    if (h > 0) return `${h}:${String(m).padStart(2, "0")}:${String(s).padStart(2, "0")}`
    return `${m}:${String(s).padStart(2, "0")}`
  }

  const formatBytes = (bytes: number): string => {
    if (bytes === 0) return "Unknown"
    const sizes = ["B", "KB", "MB", "GB"]
    const i = Math.floor(Math.log(bytes) / Math.log(1024))
    return `${(bytes / Math.pow(1024, i)).toFixed(1)} ${sizes[i]}`
  }

  return (
    <div className="flex flex-col items-center justify-start min-h-[80vh] px-4 pt-8">
      <div className="w-full max-w-xl space-y-6">
        <div className="text-center space-y-2">
          <h1 className="text-3xl font-bold text-white">My Clip</h1>
          <p className="text-gray-400">Multi-Platform Video Clipper</p>
          {version && <p className="text-xs text-gray-600">v{version}</p>}
        </div>

        {/* URL Input */}
        <form onSubmit={handleSubmit} className="space-y-3">
          <input
            type="url"
            value={url}
            onChange={(e) => { setUrl(e.target.value); setMetadata(null); setError(""); setDownloadedFile("") }}
            placeholder="Paste video URL here..."
            className="w-full px-4 py-3 bg-gray-800 border border-gray-700 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-indigo-500 transition-colors"
            autoFocus
          />
          <button
            type="submit"
            disabled={!url.trim() || loading}
            className="w-full py-3 bg-indigo-600 hover:bg-indigo-500 disabled:bg-gray-700 disabled:text-gray-500 text-white font-medium rounded-lg transition-colors"
          >
            {loading ? "Loading..." : "Load Video"}
          </button>
        </form>

        {error && (
          <div className="bg-red-900/30 border border-red-700 rounded-lg p-3 text-sm text-red-400">{error}</div>
        )}

        {/* Metadata + Download */}
        {metadata && (
          <div className="bg-gray-800/50 rounded-lg p-4 border border-gray-700/50 space-y-3">
            {metadata.thumbnail && (
              <img src={metadata.thumbnail} alt={metadata.title} className="w-full rounded-lg aspect-video object-cover" />
            )}
            <div>
              <h2 className="font-semibold text-white truncate">{metadata.title}</h2>
              <p className="text-sm text-gray-400">{metadata.author}</p>
              <div className="flex gap-3 text-xs text-gray-500 mt-1">
                <span>{metadata.source}</span><span>{formatDuration(metadata.duration)}</span>
              </div>
            </div>

            {/* Stream selector */}
            {metadata.streams.filter(s => s.has_video).length > 0 && (
              <div>
                <label className="text-xs text-gray-500 block mb-1">Quality</label>
                <select
                  value={selectedStream}
                  onChange={(e) => setSelectedStream(e.target.value)}
                  className="w-full px-3 py-2 bg-gray-800 border border-gray-700 rounded-md text-white text-sm focus:outline-none focus:border-indigo-500"
                >
                  {metadata.streams.filter(s => s.has_video).slice(0, 15).map(s => (
                    <option key={s.id} value={s.id}>
                      {s.quality || s.resolution || s.id} · {s.format} · {formatBytes(s.size)}
                    </option>
                  ))}
                </select>
              </div>
            )}

            {/* Download button */}
            {!downloadedFile && (
              <button
                onClick={handleDownload}
                disabled={downloading || !selectedStream}
                className="w-full py-2 bg-green-600 hover:bg-green-500 disabled:bg-gray-700 disabled:text-gray-500 text-white font-medium rounded-lg transition-colors"
              >
                {downloading ? "Downloading..." : "⬇ Download"}
              </button>
            )}

            {/* Download progress */}
            {downloading && (
              <div className="space-y-1">
                <div className="w-full bg-gray-700 rounded-full h-2">
                  <div className="bg-green-500 h-2 rounded-full transition-all" style={{ width: `${downloadProgress}%` }} />
                </div>
                <p className="text-xs text-gray-500">{downloadProgress.toFixed(0)}%</p>
              </div>
            )}

            {/* After download: navigate to editor */}
            {downloadedFile && (
              <button
                onClick={() => onNavigateEditor(downloadedFile, metadata.title)}
                className="w-full py-2 bg-indigo-600 hover:bg-indigo-500 text-white font-medium rounded-lg transition-colors"
              >
                ✂ Open in Editor
              </button>
            )}
          </div>
        )}

        <div className="text-center text-xs text-gray-600">
          {supportedSources.length > 0 ? `Supported: ${supportedSources.join(", ")}` : "No sources configured"}
        </div>

        <div className="bg-gray-800/50 rounded-lg p-4 border border-gray-700/50">
          <DepStatusPanel deps={deps} onRefresh={onRefreshDeps} />
        </div>
      </div>
    </div>
  )
}
