# My Clip

Multi-Platform Video Clipper — a fast, reliable desktop application for downloading, previewing, clipping, and exporting videos from multiple online sources.

Built with **Wails v3** (Go) + **React** (TypeScript, Vite, Tailwind CSS).

## Prerequisites

### Required
- **Go 1.25+** — [Download Go](https://go.dev/dl/)
- **Node.js 20+** — [Download Node.js](https://nodejs.org/)
- **Wails v3 CLI**
  ```bash
  go install github.com/wailsapp/wails/v3/cmd/wails3@latest
  ```

### External Dependencies (auto-detected at runtime)
- **FFmpeg + FFprobe** — media processing ([download](https://ffmpeg.org/download.html))
- **yt-dlp** — video downloading ([download](https://github.com/yt-dlp/yt-dlp))

> The application functions without these tools but features will be limited.
> Missing dependencies are reported in the UI.

## Installation

```bash
# Clone the repository
git clone https://github.com/mofasa798/my-clip.git
cd my-clip

# Install frontend dependencies
cd frontend && npm install
cd ..

# Generate Wails bindings
wails3 generate bindings

# Build for production
wails3 build
```

The compiled binary will be in `build/bin/my-clip.exe`.

## Development

```bash
# Run in development mode with hot reload
wails3 dev
```

The application will start with hot reload for both frontend and backend.

## Usage

1. **Paste a video URL** on the Home page (supports YouTube, Kick)
2. **Select quality** and click Download
3. **Click "Open in Editor"** after download completes
4. **Set clip range** using the timeline sliders (or press I/O keys)
5. **Choose a preset** (Fast/Balanced/Maximum) or select encoder manually
6. **Click Export** (or press Ctrl+Enter)

### Keyboard Shortcuts (Editor)

| Key | Action |
|---|---|
| `Space` | Play / Pause preview |
| `I` | Set clip In point |
| `O` | Set clip Out point |
| `Ctrl+Enter` | Start export |
| `Escape` | Back to Home |

## Project Structure

```
├── internal/         # Go backend packages
│   ├── app/          # Application service (Wails bindings)
│   ├── domain/       # Business models & value objects
│   ├── source/       # Video source adapters (YouTube, Kick)
│   │   ├── resolver/ # URL-to-adapter matching
│   │   ├── registry/ # Adapter registration
│   │   ├── ytdlp/    # Shared yt-dlp wrapper
│   │   └── download/ # Download orchestrator
│   ├── media/        # Media processing pipeline
│   │   ├── ffmpeg/   # FFmpeg wrapper & argument builder
│   │   ├── ffprobe/  # FFprobe wrapper
│   │   ├── gpu/      # GPU capability detection
│   │   └── export/   # Clip & export orchestration
│   ├── system/       # OS integration (logger, config, detector, history, presets)
│   └── shared/       # Shared utilities & error types
├── frontend/         # React + TypeScript + Vite
│   └── src/
│       ├── pages/      # Home, Editor, History, Settings
│       ├── components/ # Reusable UI components
│       ├── services/   # Backend communication layer
│       └── types/      # TypeScript type definitions
├── build/            # Wails build configurations
├── .contexts/        # Project documentation & ADRs
└── CHANGELOG.md      # Version history
```

## Supported Sources

- YouTube (`youtube.com`, `youtu.be`)
- Kick (`kick.com`)

## Encoding Strategy

```
Stream Copy → GPU (AMF / NVENC / QSV) → CPU (libx264)
```

Priority is automatic based on detected capabilities. Fall back to software encoding if no GPU is available.

## Testing

```bash
# Run all Go unit tests
go test ./internal/...

# Run specific package tests
go test ./internal/domain/...
go test ./internal/system/...
go test ./internal/media/ffmpeg/...
```

## License

MIT
