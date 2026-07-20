import { useState, useEffect } from "react"
import type { Config } from "../types"

interface Props {
  config: Config | null
  onSave: (cfg: Config) => void
}

export default function SettingsPage({ config, onSave }: Props) {
  const [outputDir, setOutputDir] = useState("")
  const [theme, setTheme] = useState("dark")
  const [encoder, setEncoder] = useState("auto")

  useEffect(() => {
    if (config) {
      setOutputDir(config.output_dir)
      setTheme(config.theme)
      setEncoder(config.preferred_encoder)
    }
  }, [config])

  const handleSave = () => {
    onSave({
      output_dir: outputDir,
      theme,
      preferred_encoder: encoder,
    })
  }

  return (
    <div className="max-w-2xl mx-auto py-8 px-4">
      <h1 className="text-2xl font-bold text-white mb-8">Settings</h1>

      <div className="space-y-6">
        {/* Output Directory */}
        <div className="space-y-2">
          <label className="block text-sm font-medium text-gray-300">
            Output Directory
          </label>
          <input
            type="text"
            value={outputDir}
            onChange={(e) => setOutputDir(e.target.value)}
            className="w-full px-3 py-2 bg-gray-800 border border-gray-700 rounded-lg
                       text-white placeholder-gray-500 focus:outline-none focus:border-indigo-500"
          />
        </div>

        {/* Theme */}
        <div className="space-y-2">
          <label className="block text-sm font-medium text-gray-300">Theme</label>
          <select
            value={theme}
            onChange={(e) => setTheme(e.target.value)}
            className="w-full px-3 py-2 bg-gray-800 border border-gray-700 rounded-lg
                       text-white focus:outline-none focus:border-indigo-500"
          >
            <option value="dark">Dark</option>
            <option value="light">Light</option>
          </select>
        </div>

        {/* Preferred Encoder */}
        <div className="space-y-2">
          <label className="block text-sm font-medium text-gray-300">
            Preferred Encoder
          </label>
          <select
            value={encoder}
            onChange={(e) => setEncoder(e.target.value)}
            className="w-full px-3 py-2 bg-gray-800 border border-gray-700 rounded-lg
                       text-white focus:outline-none focus:border-indigo-500"
          >
            <option value="auto">Auto</option>
            <option value="gpu">GPU</option>
            <option value="cpu">CPU</option>
          </select>
        </div>

        <button
          onClick={handleSave}
          className="px-6 py-2 bg-indigo-600 hover:bg-indigo-500 text-white font-medium
                     rounded-lg transition-colors"
        >
          Save Settings
        </button>
      </div>
    </div>
  )
}
