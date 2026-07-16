# 00_PROJECT.md

# Project Overview

## Vision

Build a fast, lightweight, and reliable desktop application for creating high-quality video clips from YouTube and Kick videos or livestreams.

The application is designed exclusively for **local desktop usage** and focuses on providing the fastest possible clipping workflow while maintaining excellent video quality.

The project prioritizes simplicity, maintainability, and performance over unnecessary flexibility or enterprise architecture.

---

# Objectives

The application should allow users to:

* Download videos or livestreams from YouTube and Kick.
* Preview video metadata before downloading.
* Select precise clip ranges using a visual timeline.
* Export clips quickly with minimal quality loss.
* Automatically choose the most efficient processing strategy.
* Utilize GPU hardware acceleration whenever frame re-encoding is required.
* Provide a clean, responsive, and intuitive desktop experience.

---

# Project Scope

This application is intended for:

* Personal use only.
* Windows desktop.
* Local processing.
* Single user.

The application does not require internet services beyond downloading videos from supported platforms.

No cloud infrastructure is required.

---

# Target Environment

## Operating System

* Windows 11

## Hardware

CPU

* AMD Ryzen 5 Pro 4650G
* 6 Cores / 12 Threads

GPU

* AMD Radeon RX 6600 8GB

Memory

* 16GB DDR4

Storage

* SSD recommended

The application should be optimized specifically for this hardware configuration.

---

# Technology Stack

## Desktop Framework

* Wails v3

## Frontend

* React
* TypeScript
* Vite
* Tailwind CSS

## Backend

* Go 1.24+

## Video Processing

* FFmpeg
* FFprobe

## Video Download

* yt-dlp

## Version Control

* Git

---

# Core Features

The application should support the following capabilities.

## Video Discovery

* Load YouTube videos.
* Load Kick videos.
* Load livestreams (future enhancement).
* Retrieve metadata.
* Retrieve thumbnails.
* Retrieve available video formats.

## Download

* Download selected video quality.
* Download audio when required.
* Display download progress.
* Support cancellation.

## Clip Editor

* Video preview.
* Interactive timeline.
* Start marker.
* End marker.
* Frame preview.
* Manual timestamp editing.

## Export

* Stream Copy clipping.
* GPU accelerated encoding.
* CPU fallback.
* Output filename customization.
* Output directory selection.

## Settings

* Output directory.
* Download directory.
* Theme.
* Preferred codec.
* GPU preference.

---

# Processing Strategy

The application should always select the fastest processing method available.

Processing priority:

1. Stream Copy
2. GPU Encoding
3. CPU Encoding

## Stream Copy

When the clip does not require visual modifications:

* No re-encoding.
* Original quality preserved.
* Extremely fast.
* Minimal CPU usage.
* GPU is not required.

This should always be the preferred strategy.

## GPU Encoding

GPU encoding is only used when frame processing is required.

Examples include:

* Resize
* Crop
* Watermark
* Subtitle burn-in
* Overlay
* Rotation
* Video filters
* Codec conversion
* FPS conversion

Preferred encoder:

* h264_amf

Alternative encoder:

* hevc_amf

## CPU Encoding

Used only when GPU acceleration is unavailable or unsupported.

Preferred encoder:

* libx264

---

# GPU Support

The application should automatically detect available GPU encoders during startup.

Preferred order:

1. AMD AMF
2. CPU

No manual configuration should be required for normal usage.

---

# Application Philosophy

This project follows several guiding principles.

## Local First

Everything runs locally.

No cloud services.

No remote processing.

No online authentication.

## Performance First

Always prioritize:

* Fast startup.
* Fast clipping.
* Low memory usage.
* Efficient CPU utilization.
* Efficient GPU utilization.

## Simplicity First

The project intentionally avoids unnecessary complexity.

The implementation should remain understandable by a single developer.

## Maintainability

Readable code is preferred over clever code.

Simple architecture is preferred over flexible architecture.

Avoid unnecessary abstractions.

---

# Directory Overview

The project is expected to follow this high-level structure.

```text
backend/
frontend/
bin/
storage/
.contexts/
```

Each major area has a dedicated responsibility.

Detailed structure is documented separately in:

* 02_ARCHITECTURE.md

---

# External Dependencies

Required binaries:

* FFmpeg
* FFprobe
* yt-dlp

These binaries should be bundled inside the project.

```text
bin/

ffmpeg.exe
ffprobe.exe
yt-dlp.exe
```

The backend is responsible for invoking and managing all external processes.

---

# Configuration

Application configuration should be stored locally.

Example:

```text
settings.json
```

Configuration may include:

* Output directory
* Download directory
* Preferred codec
* GPU preference
* Theme

The configuration format should remain human-readable.

---

# Logging

The application should generate logs for debugging purposes.

Suggested location:

```text
storage/logs/
```

Logs should include:

* Application startup
* External process execution
* Download events
* Export events
* Error messages

---

# Documentation

Project documentation is organized inside the `.contexts` directory.

Each document has a single responsibility.

Project documentation should never duplicate information across multiple files.

When information belongs to another document, reference that document instead.

---

# Related Documents

* INDEX.md
* 01_RULES.md
* 02_ARCHITECTURE.md
* 03_ROADMAP.md
* 04_DECISIONS.md
* 05_SPRINT.md
* 06_PROMPT.md
* 07_CODE_STYLE.md
* 08_PROJECT_STRUCTURE.md

---

# Non Goals

The following are intentionally excluded from this project.

* Cloud deployment
* Multi-user support
* Authentication
* REST API server
* Database server
* PostgreSQL
* MySQL
* Redis
* RabbitMQ
* Kafka
* Docker
* Kubernetes
* Microservices
* Distributed processing

The application is intentionally designed to remain a lightweight, local-first desktop application optimized for a single user.

---

# Success Criteria

The project is considered successful when it achieves the following goals:

* Fast application startup.
* Reliable video downloading.
* Responsive user interface.
* Accurate clip generation.
* Minimal quality loss.
* Efficient GPU utilization.
* Stable desktop experience.
* Easy long-term maintenance.
* Clear and consistent project architecture.
