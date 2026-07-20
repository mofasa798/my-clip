# Changelog

## v0.6.0 (Milestone 6) — Stabilization & Release

### Performance
- Added React.memo to critical components to reduce re-renders
- Added lazy loading with React.lazy + Suspense for code splitting
- Added context cancellation for download and export operations
- Limited FFmpeg stderr buffer size to prevent memory growth
- Added `GOGC=off` recommendation for production builds

### Error Handling
- Wrapped all external command errors with context (command name, exit code)
- Added graceful degradation when FFmpeg, FFprobe, or yt-dlp are missing
- All errors returned to frontend are user-friendly (no raw stack traces)
- Added error recovery in download service (non-nil yt-dlp wrapper)

### Testing
- Unit tests for domain models (VideoMetadata, StreamInfo, validation)
- Unit tests for system/config (load, save, defaults)
- Unit tests for ffmpeg/args (timestamp formatting, argument building)
- All tests use standard library `testing` package only
- No external dependencies invoked in unit tests

### Build
- `wails3 build` confirmed working for Windows
- Binary output to `build/bin/my-clip.exe`
- Tested graceful startup without FFmpeg/yt-dlp

### Documentation
- Added CHANGELOG.md
- Updated README with full prerequisites, installation, and build instructions

## v0.5.0 (Milestone 5) — Workflow Enhancement

- Export presets (Fast Stream Copy, Balanced GPU, Maximum Quality)
- Keyboard shortcuts (Space, I, O, Ctrl+Enter, Escape)
- History search, filter, individual delete
- Download retry with backoff (max 3 attempts)
- System notifications, clipboard copy, open folder
- Contextual error messages

## v0.4.0 (Milestone 4) — Desktop Experience

- Four-page navigation (Home, Editor, History, Settings)
- Video preview with timeline sliders
- Export panel with encoder selection
- History page with persistent storage
- Settings page with GPU info and dependencies
- Light/dark theme support

## v0.3.0 (Milestone 3) — Media Pipeline

- FFprobe wrapper for local media metadata
- FFmpeg wrapper with argument builder
- GPU detection (AMF, NVENC, QSV) with auto-fallback
- Clip extraction and export pipeline
- Media storage management

## v0.2.0 (Milestone 2) — Source Layer

- SourceAdapter interface (Name, Match, Metadata, Download)
- YouTube adapter (youtube.com, youtu.be)
- Kick adapter (kick.com)
- yt-dlp wrapper for metadata and download
- Source resolver and registry

## v0.1.0 (Milestone 1) — Project Foundation

- Wails v3 project initialization
- Go backend with layered architecture
- React + TypeScript + Vite + Tailwind CSS
- Dependency detection (FFmpeg, FFprobe, yt-dlp, AMF)
- Logger, configuration, and basic UI
