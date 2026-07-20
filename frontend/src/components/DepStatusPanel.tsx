import type { DepResult } from "../types"

interface Props {
  deps: DepResult | null
  onRefresh: () => void
}

function StatusBadge({ found }: { found: boolean }) {
  return (
    <span
      className={`inline-block w-2 h-2 rounded-full ${
        found ? "bg-green-500" : "bg-red-500"
      }`}
    />
  )
}

export default function DepStatusPanel({ deps, onRefresh }: Props) {
  if (!deps) {
    return <div className="text-gray-400 text-sm">Checking dependencies...</div>
  }

  const items = [
    { label: "FFmpeg", status: deps.ffmpeg },
    { label: "FFprobe", status: deps.ffprobe },
    { label: "yt-dlp", status: deps.yt_dlp },
    { label: "AMF Encoder", status: deps.amf },
  ]

  return (
    <div className="space-y-2">
      <div className="flex items-center justify-between">
        <h3 className="text-sm font-medium text-gray-300">Dependencies</h3>
        <button
          onClick={onRefresh}
          className="text-xs text-indigo-400 hover:text-indigo-300 transition-colors"
        >
          Refresh
        </button>
      </div>
      <div className="space-y-1">
        {items.map((item) => (
          <div
            key={item.label}
            className="flex items-center justify-between text-sm"
          >
            <div className="flex items-center gap-2">
              <StatusBadge found={item.status.found} />
              <span className="text-gray-300">{item.label}</span>
            </div>
            <span className="text-gray-500 text-xs truncate max-w-[200px]">
              {item.status.found ? item.status.version : "Not found"}
            </span>
          </div>
        ))}
      </div>
    </div>
  )
}
