# My Clip

Multi-Platform Video Clipper — a fast, reliable desktop application for downloading, previewing, clipping, and exporting videos from multiple online sources.

Built with **Wails v3** (Go) + **React** (TypeScript, Vite, Tailwind CSS).

## Getting Started

### Prerequisites

- Go 1.25+
- Node.js 20+
- npm

### Development

```bash
# Install frontend dependencies
cd frontend && npm install

# Generate Wails bindings
wails3 generate bindings

# Run in development mode (from project root)
wails3 dev
```

### Build

```bash
wails3 build
```

## Project Structure

```
├── internal/       # Go backend packages
│   ├── app/        # Application service (Wails bindings)
│   ├── domain/     # Business models
│   ├── source/     # Video source adapters
│   ├── media/      # Media processing (FFmpeg, GPU)
│   ├── system/     # OS integration (logger, config, detector)
│   └── shared/     # Shared utilities
├── frontend/       # React + TypeScript + Vite
│   └── src/
│       ├── pages/      # Home, Settings, Editor, History
│       ├── components/ # Reusable UI components
│       ├── services/   # Backend communication
│       └── types/      # TypeScript type definitions
├── build/          # Wails build configurations
└── .contexts/      # Project documentation
```

## Supported Sources

- YouTube
- Kick

## License

MIT
