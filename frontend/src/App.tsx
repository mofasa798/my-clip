import { useState, useEffect, lazy, Suspense, memo } from "react"
import { Backend } from "./services/backend"
import type { DepResult, Config } from "./types"

// Lazy-loaded pages for code splitting
const HomePage = lazy(() => import("./pages/HomePage"))
const EditorPage = lazy(() => import("./pages/EditorPage"))
const HistoryPage = lazy(() => import("./pages/HistoryPage"))
const SettingsPage = lazy(() => import("./pages/SettingsPage"))

const PageFallback = memo(function PageFallback() {
  return <div className="flex items-center justify-center min-h-[60vh] text-gray-500 text-sm">Loading...</div>
})

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
      const result = await Backend.GetDependencies()
      setDeps(result)
    } catch (err) {
      console.error("Failed to load dependencies:", err)
    }
  }

  const loadConfig = async () => {
    try {
      const cfg = await Backend.GetConfig()
      setConfig(cfg)
      if (cfg.theme) setTheme(cfg.theme)
    } catch (err) {
      console.error("Failed to load config:", err)
    }
  }

  const handleRefreshDeps = async () => {
    try {
      const result = await Backend.RefreshDependencies()
      setDeps(result)
    } catch (err) {
      console.error("Failed to refresh dependencies:", err)
    }
  }

  const handleSaveConfig = async (cfg: Config) => {
    try {
      await Backend.SaveConfig(cfg)
      setConfig(cfg)
      if (cfg.theme) setTheme(cfg.theme)
    } catch (err) {
      console.error("Failed to save config:", err)
    }
  }

  const handleOpenFolder = async (path: string) => {
    try {
      const dir = path.substring(0, path.lastIndexOf("\\"))
      await Backend.OpenFolder(dir || path)
    } catch (err) {
      console.error(err)
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
        <Suspense fallback={<PageFallback />}>
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
        </Suspense>
      </main>

      {/* Status Bar */}
      <footer className={`fixed bottom-0 left-0 right-0 border-t ${footerBg} px-4 py-1.5`}>
        <div className="max-w-5xl mx-auto flex items-center justify-between text-xs text-gray-500">
          <span>Ready</span>
          <span>My Clip v0.6.0</span>
        </div>
      </footer>
    </div>
  )
}
