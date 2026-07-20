import { useState, useEffect } from "react"
import HomePage from "./pages/HomePage"
import SettingsPage from "./pages/SettingsPage"
import type { DepResult, Config } from "./types"

type Page = "home" | "settings"

export default function App() {
  const [currentPage, setCurrentPage] = useState<Page>("home")
  const [deps, setDeps] = useState<DepResult | null>(null)
  const [config, setConfig] = useState<Config | null>(null)

  const loadDeps = async () => {
    try {
      if (window.GoApp) {
        const result = await window.GoApp.GetDependencies()
        setDeps(result)
      }
    } catch (err) {
      console.error("Failed to load dependencies:", err)
    }
  }

  const loadConfig = async () => {
    try {
      if (window.GoApp) {
        const cfg = await window.GoApp.GetConfig()
        setConfig(cfg as unknown as Config)
      }
    } catch (err) {
      console.error("Failed to load config:", err)
    }
  }

  const handleRefreshDeps = async () => {
    try {
      if (window.GoApp) {
        const result = await window.GoApp.RefreshDependencies()
        setDeps(result)
      }
    } catch (err) {
      console.error("Failed to refresh dependencies:", err)
    }
  }

  const handleSaveConfig = async (cfg: Config) => {
    try {
      if (window.GoApp) {
        await window.GoApp.SaveConfig(cfg as unknown as Record<string, string>)
        setConfig(cfg)
      }
    } catch (err) {
      console.error("Failed to save config:", err)
    }
  }

  useEffect(() => {
    loadDeps()
    loadConfig()
  }, [])

  const navItems: { id: Page; label: string }[] = [
    { id: "home", label: "Home" },
    { id: "settings", label: "Settings" },
  ]

  return (
    <div className="min-h-screen bg-gray-900 text-white">
      {/* Navigation */}
      <nav className="border-b border-gray-800 px-4 py-3">
        <div className="max-w-4xl mx-auto flex items-center justify-between">
          <div className="flex items-center gap-6">
            <span className="font-bold text-lg text-white">My Clip</span>
            <div className="flex gap-1">
              {navItems.map((item) => (
                <button
                  key={item.id}
                  onClick={() => setCurrentPage(item.id)}
                  className={`px-3 py-1.5 text-sm rounded-md transition-colors ${
                    currentPage === item.id
                      ? "bg-gray-700 text-white"
                      : "text-gray-400 hover:text-white hover:bg-gray-800"
                  }`}
                >
                  {item.label}
                </button>
              ))}
            </div>
          </div>
        </div>
      </nav>

      {/* Page Content */}
      <main>
        {currentPage === "home" && (
          <HomePage deps={deps} onRefreshDeps={handleRefreshDeps} />
        )}
        {currentPage === "settings" && (
          <SettingsPage config={config} onSave={handleSaveConfig} />
        )}
      </main>

      {/* Status Bar */}
      <footer className="fixed bottom-0 left-0 right-0 border-t border-gray-800 bg-gray-900 px-4 py-1.5">
        <div className="max-w-4xl mx-auto flex items-center justify-between text-xs text-gray-600">
          <span>Ready</span>
          <span>My Clip v0.1.0</span>
        </div>
      </footer>
    </div>
  )
}
