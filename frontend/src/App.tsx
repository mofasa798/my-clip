import { useState, useEffect } from "react"
import HomePage from "./pages/HomePage"
import EditorPage from "./pages/EditorPage"
import HistoryPage from "./pages/HistoryPage"
import SettingsPage from "./pages/SettingsPage"
import type { DepResult, Config } from "./types"

type Page = "home" | "editor" | "history" | "settings"

export default function App() {
  const [currentPage, setCurrentPage] = useState<Page>("home")
  const [deps, setDeps] = useState<DepResult | null>(null)
  const [config, setConfig] = useState<Config | null>(null)
  const [theme, setTheme] = useState("dark")

  // Editor state
  const [editorFile, setEditorFile] = useState("")
  const [editorTitle, setEditorTitle] = useState("")

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
        if (cfg.theme) setTheme(cfg.theme)
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
        if (cfg.theme) setTheme(cfg.theme)
      }
    } catch (err) {
      console.error("Failed to save config:", err)
    }
  }

  const handleOpenFolder = (path: string) => {
    if (window.GoApp?.OpenFolder) {
      const dir = path.substring(0, path.lastIndexOf("\\"))
      window.GoApp.OpenFolder(dir || path).catch(console.error)
    }
  }

  const handleNavigateEditor = (filePath: string, title: string) => {
    setEditorFile(filePath)
    setEditorTitle(title)
    setCurrentPage("editor")
  }

  useEffect(() => {
    loadDeps()
    loadConfig()
  }, [])

  const isDark = theme === "dark" || (!theme && true)
  const bgClass = isDark ? "bg-gray-900 text-white" : "bg-white text-gray-900"
  const navBg = isDark ? "border-gray-800" : "border-gray-200"
  const footerBg = isDark ? "bg-gray-900 border-gray-800" : "bg-white border-gray-200"

  const navItems: { id: Page; label: string }[] = [
    { id: "home", label: "Home" },
    { id: "history", label: "History" },
    { id: "settings", label: "Settings" },
  ]

  return (
    <div className={`min-h-screen ${bgClass}`}>
      {/* Navigation */}
      <nav className={`border-b ${navBg} px-4 py-3`}>
        <div className="max-w-5xl mx-auto flex items-center justify-between">
          <div className="flex items-center gap-6">
            <span className={`font-bold text-lg ${isDark ? "text-white" : "text-gray-900"}`}>My Clip</span>
            <div className="flex gap-1">
              {navItems.map((item) => (
                <button
                  key={item.id}
                  onClick={() => setCurrentPage(item.id)}
                  className={`px-3 py-1.5 text-sm rounded-md transition-colors ${
                    currentPage === item.id
                      ? isDark ? "bg-gray-700 text-white" : "bg-gray-200 text-gray-900"
                      : isDark ? "text-gray-400 hover:text-white hover:bg-gray-800" : "text-gray-500 hover:text-gray-900 hover:bg-gray-100"
                  }`}
                >
                  {item.label}
                </button>
              ))}
            </div>
          </div>
          {currentPage === "editor" && (
            <button onClick={() => setCurrentPage("home")}
              className="text-sm text-indigo-400 hover:text-indigo-300">
              &larr; Home
            </button>
          )}
        </div>
      </nav>

      {/* Page Content */}
      <main className={currentPage !== "home" ? "pb-12" : ""}>
        {currentPage === "home" && (
          <HomePage deps={deps} onRefreshDeps={handleRefreshDeps} onNavigateEditor={handleNavigateEditor} />
        )}
        {currentPage === "editor" && (
          <EditorPage videoPath={editorFile} videoTitle={editorTitle} onBack={() => setCurrentPage("home")} />
        )}
        {currentPage === "history" && (
          <HistoryPage onOpenFolder={handleOpenFolder} />
        )}
        {currentPage === "settings" && (
          <SettingsPage config={config} deps={deps} onSave={handleSaveConfig} onRefreshDeps={handleRefreshDeps} />
        )}
      </main>

      {/* Status Bar */}
      <footer className={`fixed bottom-0 left-0 right-0 border-t ${footerBg} px-4 py-1.5`}>
        <div className="max-w-5xl mx-auto flex items-center justify-between text-xs text-gray-500">
          <span>Ready</span>
          <span>My Clip v0.3.0</span>
        </div>
      </footer>
    </div>
  )
}
