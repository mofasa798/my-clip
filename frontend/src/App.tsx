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

  // Editor state — persists across tab switches
  const [editorFile, setEditorFile] = useState("")
  const [editorTitle, setEditorTitle] = useState("")

  // History refresh counter — incremented after download/export
  const [historyRevision, setHistoryRevision] = useState(0)

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

  const handleExportDone = () => {
    setHistoryRevision(v => v + 1)
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
    { id: "editor", label: "Editor" },
    { id: "history", label: "History" },
    { id: "settings", label: "Settings" },
  ]

  // Hide / show pages instead of unmounting, so state persists
  const pageClass = (page: Page) =>
    `w-full ${currentPage === page ? "block" : "hidden"}`

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
          {currentPage === "editor" && editorFile && (
            <button onClick={() => setCurrentPage("home")}
              className="text-sm text-indigo-400 hover:text-indigo-300">
              &larr; Home
            </button>
          )}
        </div>
      </nav>

      {/* Page Content — kept mounted to preserve state */}
      <main>
        <Suspense fallback={<PageFallback />}>
          <div className={pageClass("home")}>
            <HomePage deps={deps} onRefreshDeps={handleRefreshDeps} onNavigateEditor={handleNavigateEditor} />
          </div>
          <div className={pageClass("editor")}>
            {editorFile ? (
              <EditorPage
                videoPath={editorFile}
                videoTitle={editorTitle}
                onBack={() => setCurrentPage("home")}
                onExportDone={handleExportDone}
              />
            ) : (
              <div className="flex flex-col items-center justify-center min-h-[60vh] text-gray-500">
                <p className="text-lg mb-2">No video loaded</p>
                <p className="text-sm text-gray-600">Download a video from Home first, then open it in the Editor.</p>
              </div>
            )}
          </div>
          <div className={pageClass("history")}>
            <HistoryPage onOpenFolder={handleOpenFolder} revision={historyRevision} />
          </div>
          <div className={pageClass("settings")}>
            <SettingsPage config={config} deps={deps} onSave={handleSaveConfig} onRefreshDeps={handleRefreshDeps} />
          </div>
        </Suspense>
      </main>

      {/* Status Bar */}
      <footer className={`fixed bottom-0 left-0 right-0 border-t ${footerBg} px-4 py-1.5`}>
        <div className="max-w-5xl mx-auto flex items-center justify-between text-xs text-gray-500">
          <span>{currentPage === "editor" && editorFile ? `Editing: ${editorTitle}` : "Ready"}</span>
          <span>My Clip v0.6.0</span>
        </div>
      </footer>
    </div>
  )
}
