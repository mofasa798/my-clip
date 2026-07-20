import { useState } from "react"
import DepStatusPanel from "../components/DepStatusPanel"
import type { DepResult, VideoMetadata, DownloadResult } from "../types"

// The Go App will be injected via Wails bindings.
// For now we use a global reference set after initialization.
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
    }
  }
}

interface Props {
  deps: DepResult | null
  onRefreshDeps: () => void
}

export default function HomePage({ deps, onRefreshDeps }: Props) {
  const [url, setUrl] = useState("")
  const [version, setVersion] = useState("")
  const [loading, setLoading] = useState(false)
  const [metadata, setMetadata] = useState<VideoMetadata | null>(null)
  const [error, setError] = useState("")
  const [supportedSources, setSupportedSources] = useState<string[]>([])

  // Load version and supported sources on mount
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

    try {
      if (window.GoApp) {
        // First resolve the source
        const source = await window.GoApp.ResolveSource(url)
        if (!source) {
          setError("URL not supported. Try a YouTube or Kick link.")
          setLoading(false)
          return
        }
        // Then load metadata
        const meta = await window.GoApp.GetMetadata(url)
        setMetadata(meta)
      }
    } catch (err: any) {
      setError(err?.message || "Failed to load video. Check the URL and try again.")
    } finally {
      setLoading(false)
    }
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
    <div className="flex flex-col items-center justify-center min-h-[80vh] px-4">
      <div className="w-full max-w-xl space-y-8">
        {/* Header */}
        <div className="text-center space-y-2">
          <h1 className="text-3xl font-bold text-white">My Clip</h1>
          <p className="text-gray-400">
            Multi-Platform Video Clipper
          </p>
          {version && (
            <p className="text-xs text-gray-600">v{version}</p>
          )}
        </div>

        {/* URL Input */}
        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="relative">
            <input
              type="url"
              value={url}
              onChange={(e) => {
                setUrl(e.target.value)
                setMetadata(null)
                setError("")
              }}
              placeholder="Paste video URL here..."
              className="w-full px-4 py-3 bg-gray-800 border border-gray-700 rounded-lg 
                         text-white placeholder-gray-500 focus:outline-none focus:border-indigo-500
                         transition-colors"
              autoFocus
            />
          </div>
          <button
            type="submit"
            disabled={!url.trim() || loading}
            className="w-full py-3 bg-indigo-600 hover:bg-indigo-500 disabled:bg-gray-700
                       disabled:text-gray-500 text-white font-medium rounded-lg
                       transition-colors"
          >
            {loading ? "Loading..." : "Load Video"}
          </button>
        </form>

        {/* Error Message */}
        {error && (
          <div className="bg-red-900/30 border border-red-700 rounded-lg p-3 text-sm text-red-400">
            {error}
          </div>
        )}

        {/* Metadata Preview */}
        {metadata && (
          <div className="bg-gray-800/50 rounded-lg p-4 border border-gray-700/50 space-y-3">
            {metadata.thumbnail && (
              <img
                src={metadata.thumbnail}
                alt={metadata.title}
                className="w-full rounded-lg aspect-video object-cover"
              />
            )}
            <div className="space-y-1">
              <h2 className="font-semibold text-white truncate">{metadata.title}</h2>
              <p className="text-sm text-gray-400">{metadata.author}</p>
              <div className="flex gap-3 text-xs text-gray-500">
                <span>{metadata.source}</span>
                <span>{formatDuration(metadata.duration)}</span>
                {metadata.streams.length > 0 && (
                  <span>{metadata.streams.length} formats</span>
                )}
              </div>
            </div>

            {/* Stream Selection */}
            {metadata.streams.length > 0 && (
              <div className="space-y-2">
                <h3 className="text-sm font-medium text-gray-300">Available Qualities</h3>
                <div className="grid grid-cols-2 gap-2 max-h-40 overflow-y-auto">
                  {metadata.streams
                    .filter((s) => s.has_video)
                    .slice(0, 10)
                    .map((stream) => (
                      <div
                        key={stream.id}
                        className="bg-gray-800 rounded p-2 border border-gray-700 text-xs"
                      >
                        <div className="text-gray-200 font-medium">{stream.quality || stream.resolution}</div>
                        <div className="text-gray-500">{stream.format} · {formatBytes(stream.size)}</div>
                      </div>
                    ))}
                </div>
              </div>
            )}
          </div>
        )}

        {/* Supported Sources */}
        <div className="text-center text-xs text-gray-600">
          {supportedSources.length > 0
            ? `Supported: ${supportedSources.join(", ")}`
            : "No sources configured"}
        </div>

        {/* Dependency Status */}
        <div className="bg-gray-800/50 rounded-lg p-4 border border-gray-700/50">
          <DepStatusPanel deps={deps} onRefresh={onRefreshDeps} />
        </div>
      </div>
    </div>
  )
}
