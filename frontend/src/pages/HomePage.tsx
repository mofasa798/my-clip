import { useState } from "react"
import DepStatusPanel from "../components/DepStatusPanel"
import type { DepResult } from "../types"

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

  // Load version on mount
  useState(() => {
    if (window.GoApp) {
      window.GoApp.GetVersion().then(setVersion).catch(console.error)
    }
  })

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (!url.trim()) return
    // Will navigate to editor in a future milestone
    console.log("URL submitted:", url)
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
              onChange={(e) => setUrl(e.target.value)}
              placeholder="Paste video URL here..."
              className="w-full px-4 py-3 bg-gray-800 border border-gray-700 rounded-lg 
                         text-white placeholder-gray-500 focus:outline-none focus:border-indigo-500
                         transition-colors"
              autoFocus
            />
          </div>
          <button
            type="submit"
            disabled={!url.trim()}
            className="w-full py-3 bg-indigo-600 hover:bg-indigo-500 disabled:bg-gray-700
                       disabled:text-gray-500 text-white font-medium rounded-lg
                       transition-colors"
          >
            Load Video
          </button>
        </form>

        {/* Supported Sources */}
        <div className="text-center text-xs text-gray-600">
          Supported: YouTube, Kick
        </div>

        {/* Dependency Status */}
        <div className="bg-gray-800/50 rounded-lg p-4 border border-gray-700/50">
          <DepStatusPanel deps={deps} onRefresh={onRefreshDeps} />
        </div>
      </div>
    </div>
  )
}
