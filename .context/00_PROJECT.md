# PROJECT.md

# Project Overview

This project is a **local-first Windows desktop application** for creating high-performance video clips from **YouTube** and **Kick** videos or livestreams.

The application is intended **only for personal use** and will **never be deployed to the cloud**.

The primary objective is to provide a fast, lightweight, and responsive clipping experience by leveraging **FFmpeg**, **yt-dlp**, and **AMD GPU Hardware Acceleration (AMF)** whenever re-encoding is required.

---

# Development Principles

This project should always prioritize:

- Simplicity over complexity.
- Performance over unnecessary abstraction.
- Native desktop experience.
- Clean and maintainable code.
- Local processing only.
- Minimal dependencies.
- No microservices.
- No cloud services.
- No Docker.
- No Kubernetes.
- No distributed architecture.

Avoid over-engineering.

---

# Target Environment

Operating System

- Windows 11

Hardware

CPU
- AMD Ryzen 5 Pro 4650G (6C / 12T)

GPU
- AMD Radeon RX6600 8GB

Memory
- 16GB DDR4

Development should be optimized specifically for this hardware configuration.

---

# Technology Stack

## Desktop

- Wails v3
- React
- TypeScript
- Vite
- TailwindCSS

## Backend

- Go 1.24+
- Goroutines
- Context
- Worker Pool

## Video Processing

- FFmpeg
- FFprobe

## Video Download

- yt-dlp

External binaries are stored inside:

```
bin/

ffmpeg.exe
ffprobe.exe
yt-dlp.exe
```

---

# Project Structure

```
backend/

cmd/

internal/

app/
clip/
download/
ffmpeg/
ffprobe/
gpu/
history/
settings/
worker/
utils/

frontend/

components/
pages/
hooks/
types/

storage/

downloads/
output/
cache/
thumbnail/
logs/

settings.json
```

---

# Architecture

```
React UI

↓

Wails

↓

Go Application Layer

↓

────────────────────────

yt-dlp

↓

FFmpeg

↓

Output Video

────────────────────────
```

Everything runs locally.

No external server.

No REST API.

No database server.

---

# Video Processing Strategy

There are only two processing modes.

## 1. Stream Copy (Preferred)

If the user only trims a video without applying any modification:

- No re-encoding
- Preserve original quality
- Extremely fast
- GPU is NOT used

Always prefer:

```
-c copy
```

This should be the default clipping strategy.

---

## 2. GPU Encoding

GPU encoding is only required when video processing modifies frames, such as:

- Resize
- Crop
- Watermark
- Subtitle burn-in
- Overlay
- Rotation
- Filters
- FPS conversion
- Codec conversion

Preferred encoder:

```
h264_amf
```

Alternative:

```
hevc_amf
```

Fallback:

```
libx264
```

---

# GPU Detection

At application startup:

Run

```
ffmpeg -encoders
```

If available:

```
h264_amf
```

Enable GPU mode.

Otherwise automatically fallback to CPU encoding.

The user should not need to configure this manually.

---

# Storage

No database is required.

Store application data using:

```
storage/

downloads/
output/
thumbnail/
logs/
```

Configuration:

```
settings.json
```

History may be stored as JSON if needed.

---

# Concurrency

Use Goroutines for:

- Metadata loading
- Downloading
- Thumbnail generation
- Video processing
- Progress updates

All long-running operations must support cancellation using Context.

---

# Progress Updates

Backend emits events.

```
Backend

↓

Wails Events

↓

Frontend

↓

Progress UI
```

Do not use WebSocket.

---

# Logging

Every external process should capture:

- stdout
- stderr
- exit code
- execution time

Store logs under:

```
storage/logs/
```

---

# Coding Guidelines

Always follow these rules.

- Keep business logic independent from UI.
- Never call FFmpeg directly from React.
- Wrap FFmpeg inside Go services.
- Wrap yt-dlp inside Go services.
- Prefer composition over inheritance.
- Avoid global mutable state.
- Use interfaces only where they provide clear value.
- Keep packages cohesive.
- Avoid unnecessary abstractions.
- Favor readable code over clever code.

---
# AI Instructions

When generating code for this project:

- Follow the current development phase only.
- Do not implement future phases unless explicitly requested.
- Keep implementations modular and maintainable.
- Avoid unnecessary libraries.
- Prefer native Go capabilities whenever possible.
- Keep FFmpeg and yt-dlp interactions isolated behind service layers.
- Prioritize performance and readability over architectural complexity.