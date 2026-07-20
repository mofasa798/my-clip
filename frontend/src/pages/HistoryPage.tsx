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

  useEffect(() => {
    loadHistory()
  }, [])

  return (
    <div className="max-w-3xl mx-auto py-8 px-4">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold text-white">History</h1>
        <div className="flex gap-2">
          <button
            onClick={loadHistory}
            className="px-3 py-1.5 text-sm text-indigo-400 hover:text-indigo-300 transition-colors"
          >
            Refresh
          </button>
          {entries.length > 0 && (
            <button
              onClick={handleClear}
              className="px-3 py-1.5 text-sm text-red-400 hover:text-red-300 transition-colors"
            >
              Clear All
            </button>
          )}
        </div>
      </div>

      {loading ? (
        <div className="text-gray-400 text-sm">Loading history...</div>
      ) : entries.length === 0 ? (
        <div className="text-gray-500 text-sm text-center py-12">
          No history yet. Download or export a video to see it here.
        </div>
      ) : (
        <div className="space-y-2">
          {entries.map((entry, i) => (
            <div
              key={entry.id || i}
              className="bg-gray-800/50 rounded-lg p-3 border border-gray-700/50 flex items-center justify-between"
            >
              <div className="min-w-0 flex-1">
                <div className="flex items-center gap-2">
                  <span className={`text-xs px-1.5 py-0.5 rounded ${
                    entry.type === "download" ? "bg-blue-900/50 text-blue-400" : "bg-green-900/50 text-green-400"
                  }`}>
                    {entry.type}
                  </span>
                  <span className="text-sm text-white truncate">{entry.title}</span>
                </div>
                <div className="flex gap-3 text-xs text-gray-500 mt-1">
                  <span>{entry.source}</span>
                  <span>{formatBytes(entry.file_size)}</span>
                  <span>{formatDate(entry.timestamp)}</span>
                  <span className={entry.status === "completed" ? "text-green-500" : "text-red-500"}>
                    {entry.status}
                  </span>
                </div>
              </div>
              {entry.file_path && (
                <button
                  onClick={() => onOpenFolder(entry.file_path)}
                  className="ml-3 px-3 py-1 text-xs bg-gray-700 hover:bg-gray-600 text-gray-300 rounded-md transition-colors shrink-0"
                >
                  Open
                </button>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
