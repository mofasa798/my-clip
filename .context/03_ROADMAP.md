# Current Phase

Phase 1

Only implement tasks from the active phase.

Future phases must not be implemented unless explicitly requested.

# DEVELOPMENT PHASES

---

# Phase 1 — Foundation

Goal

Build the project foundation.

Tasks

- Initialize Wails project
- Setup React + TypeScript
- Setup TailwindCSS
- Configure Go project
- Implement application settings
- Detect FFmpeg
- Detect FFprobe
- Detect yt-dlp
- Detect GPU support
- Create folder structure
- Create Settings page

Deliverable

Application starts successfully and verifies all required dependencies.

---

# Phase 2 — Video Discovery

Goal

Allow users to load videos.

Tasks

- URL input
- Validate URL
- Fetch metadata
- Fetch thumbnail
- Fetch duration
- Fetch available resolutions
- Display video information
- Download selected quality

Deliverable

Users can successfully download videos locally.

---

# Phase 3 — Clip Editor

Goal

Create a simple clipping interface.

Tasks

- Video preview
- Timeline
- Start marker
- End marker
- Frame preview
- Time input
- Output filename
- Export options

Deliverable

Users can select clip ranges visually.

---

# Phase 4 — Video Processing

Goal

Generate clips efficiently.

Tasks

- Smart Stream Copy
- GPU Encoding
- CPU fallback
- Progress reporting
- Cancel processing
- Error handling
- Output validation

Priority

Always prefer Stream Copy whenever possible.

Deliverable

High-speed clip generation.

---

# Phase 5 — User Experience

Goal

Improve usability.

Tasks

- History
- Recent videos
- Remember settings
- Progress notifications
- Output folder shortcut
- Theme support
- Better error messages
- Logging viewer

Deliverable

Complete desktop user experience.

---

# Phase 6 — Advanced Features

Goal

Optional power-user functionality.

Tasks

- Playlist download
- Livestream clipping
- Batch clipping
- Thumbnail generator
- Watermark
- Subtitle burn-in
- Presets
- AV1 encoding
- Multi-export
- Automatic updates

These features should only be implemented after all previous phases are stable.

---

# Non Goals

This project will NEVER include:

- Cloud deployment
- Authentication
- Multi-user support
- Remote API
- Distributed processing
- Docker
- Kubernetes
- Microservices
- PostgreSQL
- Redis
- RabbitMQ
- Kafka

The application is intentionally designed to remain a **fast, lightweight, single-user desktop application**.

---
