import { useState, useEffect } from "react"
import type { Config, DepResult, GPUInfo, EncoderOption } from "../types"

interface Props {
  config: Config | null
  deps: DepResult | null
  onSave: (cfg: Config) => void
  onRefreshDeps: () => void
}

export default function SettingsPage({ config, deps, onSave, onRefreshDeps }: Props) {
  const [outputDir, setOutputDir] = useState("")
  const [theme, setTheme] = useState("dark")
  const [encoder, setEncoder] = useState("auto")
  const [gpuInfo, setGpuInfo] = useState<GPUInfo | null>(null)
  const [encoders, setEncoders] = useState<EncoderOption[]>([])

  // Load config values when config changes
  useEffect(() => {
    if (!config) return
    setOutputDir(config.output_dir)
    setTheme(config.theme)
    setEncoder(config.preferred_encoder)
  }, [config])

  // Fetch GPU info and encoders once on mount
  useEffect(() => {
    window.GoApp?.GetGPUInfo?.().then(setGpuInfo).catch(console.error)
    window.GoApp?.GetAvailableEncoders?.().then(setEncoders).catch(console.error)
  }, [])

  const handleSave = () => {
    onSave({ output_dir: outputDir, theme, preferred_encoder: encoder })
  }

  const depItems = deps ? [
    { label: "FFmpeg", found: deps.ffmpeg.found, version: deps.ffmpeg.version },
    { label: "FFprobe", found: deps.ffprobe.found, version: deps.ffprobe.version },
    { label: "yt-dlp", found: deps.yt_dlp.found, version: deps.yt_dlp.version },
  ] : []

  return (
    <div className="max-w-2xl mx-auto py-8 px-4 space-y-8">
      <h1 className="text-2xl font-bold text-white">Settings</h1>

      {/* General */}
      <section className="space-y-4">
        <h2 className="text-sm font-semibold text-gray-400 uppercase tracking-wider">General</h2>
        <div className="space-y-2">
          <label className="block text-sm font-medium text-gray-300">Output Directory</label>
          <input type="text" value={outputDir} onChange={(e) => setOutputDir(e.target.value)}
            className="w-full px-3 py-2 bg-gray-800 border border-gray-700 rounded-lg text-white focus:outline-none focus:border-indigo-500" />
        </div>
        <div className="space-y-2">
          <label className="block text-sm font-medium text-gray-300">Theme</label>
          <select value={theme} onChange={(e) => setTheme(e.target.value)}
            className="w-full px-3 py-2 bg-gray-800 border border-gray-700 rounded-lg text-white focus:outline-none focus:border-indigo-500">
            <option value="dark">Dark</option>
            <option value="light">Light</option>
          </select>
        </div>
        <div className="space-y-2">
          <label className="block text-sm font-medium text-gray-300">Preferred Encoder</label>
          <select value={encoder} onChange={(e) => setEncoder(e.target.value)}
            className="w-full px-3 py-2 bg-gray-800 border border-gray-700 rounded-lg text-white focus:outline-none focus:border-indigo-500">
            <option value="auto">Auto</option>
            {encoders.filter(e => e.available).map(e => (
              <option key={e.value} value={e.value}>{e.name}</option>
            ))}
            <option value="libx264">CPU (libx264)</option>
          </select>
        </div>
        <button onClick={handleSave}
          className="px-6 py-2 bg-indigo-600 hover:bg-indigo-500 text-white font-medium rounded-lg transition-colors">
          Save Settings
        </button>
      </section>

      {/* GPU Info */}
      {gpuInfo && (
        <section className="space-y-3">
          <h2 className="text-sm font-semibold text-gray-400 uppercase tracking-wider">GPU Acceleration</h2>
          <div className="bg-gray-800/50 rounded-lg p-4 border border-gray-700/50 space-y-2">
            <div className="flex items-center justify-between text-sm">
              <span className="text-gray-300">Status</span>
              <span className={gpuInfo.gpu_available ? "text-green-400" : "text-gray-500"}>
                {gpuInfo.gpu_available ? `Available (${gpuInfo.gpu_vendor})` : "Not available"}
              </span>
            </div>
            <div className="flex items-center justify-between text-sm">
              <span className="text-gray-300">Preferred</span>
              <span className="text-gray-400">{gpuInfo.preferred}</span>
            </div>
            <div className="text-xs text-gray-500 space-y-1 pt-1">
              {(gpuInfo.encoders as Array<{ name: string; available: boolean; vendor: string }>).map((enc) => (
                <div key={enc.name} className="flex items-center gap-2">
                  <span className={`w-1.5 h-1.5 rounded-full ${enc.available ? "bg-green-500" : "bg-gray-600"}`} />
                  <span>{enc.name}</span>
                  {enc.vendor && <span className="text-gray-600">{enc.vendor}</span>}
                </div>
              ))}
            </div>
          </div>
        </section>
      )}

      {/* Dependencies */}
      <section className="space-y-3">
        <div className="flex items-center justify-between">
          <h2 className="text-sm font-semibold text-gray-400 uppercase tracking-wider">Dependencies</h2>
          <button onClick={onRefreshDeps} className="text-xs text-indigo-400 hover:text-indigo-300">Refresh</button>
        </div>
        <div className="bg-gray-800/50 rounded-lg p-4 border border-gray-700/50 space-y-1">
          {depItems.length === 0 ? (
            <div className="text-gray-500 text-sm">Checking...</div>
          ) : depItems.map((item) => (
            <div key={item.label} className="flex items-center justify-between text-sm">
              <div className="flex items-center gap-2">
                <span className={`w-2 h-2 rounded-full ${item.found ? "bg-green-500" : "bg-red-500"}`} />
                <span className="text-gray-300">{item.label}</span>
              </div>
              <span className="text-gray-500 text-xs truncate max-w-[200px]">
                {item.found ? item.version : "Not found"}
              </span>
            </div>
          ))}
        </div>
      </section>
    </div>
  )
}
