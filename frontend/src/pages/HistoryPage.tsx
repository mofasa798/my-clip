import { useState, useEffect } from "react"
import type { HistoryEntry } from "../types"

interface Props {
  onOpenFolder: (path: string) => void
}

function formatBytes(bytes: number): string {
  if (bytes === 0) return "Unknown"
  const sizes = ["B", "KB", "MB", "GB"]
  const i = Math.floor(Math.log(bytes) / Math.log(1024))
  return `${(bytes / Math.pow(1024, i)).toFixed(1)} ${sizes[i]}`
}

function formatDate(ts: string): string {
  try {
    const d = new Date(ts)
    return d.toLocaleDateString() + " " + d.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" })
  } catch {
    return ts
  }
}

export default function HistoryPage({ onOpenFolder }: Props) {
  const [entries, setEntries] = useState<HistoryEntry[]>([])
  const [loading, setLoading] = useState(true)
  const [search, setSearch] = useState("")
  const [typeFilter, setTypeFilter] = useState<"all" | "download" | "export">("all")

  const loadHistory = async () => {
    setLoading(true)
    try {
      if (window.GoApp?.GetHistory) {
        const result = await window.GoApp.GetHistory()
        setEntries(result || [])
      }
    } catch (err) {
      console.error("Failed to load history:", err)
    } finally {
      setLoading(false)
    }
  }

  const handleClear = async () => {
    try {
      if (window.GoApp?.ClearHistory) {
        await window.GoApp.ClearHistory()
        setEntries([])
      }
    } catch (err) {
      console.error("Failed to clear history:", err)
    }
  }

  const handleDelete = async (index: number) => {
    try {
      if (window.GoApp?.DeleteHistoryEntry) {
        await window.GoApp.DeleteHistoryEntry(index)
        setEntries(prev => prev.filter((_, i) => i !== index))
      }
    } catch (err) {
      console.error("Failed to delete entry:", err)
    }
  }

  useEffect(() => {
    loadHistory()
  }, [])

  // Filter entries
  const filtered = entries.filter((e, i) => {
    if (typeFilter !== "all" && e.type !== typeFilter) return false
    if (search.trim()) {
      const q = search.toLowerCase()
      return e.title?.toLowerCase().includes(q) ||
             e.source?.toLowerCase().includes(q)
    }
    return true
  })

  // Reverse to show newest first
  const displayed = [...filtered].reverse()

  return (
    <div className="max-w-3xl mx-auto py-8 px-4">
      <div className="flex items-center justify-between mb-4">
        <h1 className="text-2xl font-bold text-white">History</h1>
        <div className="flex gap-2">
          <button onClick={loadHistory} className="px-3 py-1.5 text-sm text-indigo-400 hover:text-indigo-300 transition-colors">Refresh</button>
          {entries.length > 0 && (
            <button onClick={handleClear} className="px-3 py-1.5 text-sm text-red-400 hover:text-red-300 transition-colors">Clear All</button>
          )}
        </div>
      </div>

      {/* Search & Filter */}
      <div className="flex gap-2 mb-4">
        <input type="text" value={search} onChange={(e) => setSearch(e.target.value)}
          placeholder="Search by title or source..."
          className="flex-1 px-3 py-1.5 bg-gray-800 border border-gray-700 rounded-md text-white text-sm placeholder-gray-500 focus:outline-none focus:border-indigo-500" />
        <select value={typeFilter} onChange={(e) => setTypeFilter(e.target.value as any)}
          className="px-3 py-1.5 bg-gray-800 border border-gray-700 rounded-md text-white text-sm focus:outline-none focus:border-indigo-500">
          <option value="all">All</option>
          <option value="download">Downloads</option>
          <option value="export">Exports</option>
        </select>
      </div>

      {loading ? (
        <div className="text-gray-400 text-sm">Loading history...</div>
      ) : displayed.length === 0 ? (
        <div className="text-gray-500 text-sm text-center py-12">
          {entries.length === 0 ? "No history yet." : "No matching entries."}
        </div>
      ) : (
        <div className="space-y-2">
          {displayed.map((entry, displayIdx) => {
            // Find actual index in original entries array (reversed)
            const origIdx = entries.length - 1 - displayed.indexOf(entry)
            return (
              <div key={entry.id || origIdx}
                className="bg-gray-800/50 rounded-lg p-3 border border-gray-700/50 flex items-center justify-between group">
                <div className="min-w-0 flex-1">
                  <div className="flex items-center gap-2">
                    <span className={`text-xs px-1.5 py-0.5 rounded ${
                      entry.type === "download" ? "bg-blue-900/50 text-blue-400" : "bg-green-900/50 text-green-400"
                    }`}>{entry.type}</span>
                    <span className="text-sm text-white truncate">{entry.title}</span>
                  </div>
                  <div className="flex gap-3 text-xs text-gray-500 mt-1">
                    <span>{entry.source}</span>
                    <span>{formatBytes(entry.file_size)}</span>
                    <span>{formatDate(entry.timestamp)}</span>
                    <span className={entry.status === "completed" ? "text-green-500" : "text-red-500"}>{entry.status}</span>
                  </div>
                </div>
                <div className="flex gap-1 shrink-0 ml-2">
                  {entry.file_path && (
                    <button onClick={() => onOpenFolder(entry.file_path)}
                      className="px-2 py-1 text-xs bg-gray-700 hover:bg-gray-600 text-gray-300 rounded-md transition-colors">Open</button>
                  )}
                  <button onClick={() => handleDelete(origIdx)}
                    className="px-2 py-1 text-xs bg-gray-700 hover:bg-red-700 text-gray-300 hover:text-white rounded-md transition-colors opacity-0 group-hover:opacity-100">✕</button>
                </div>
              </div>
            )
          })}
        </div>
      )}
    </div>
  )
}
