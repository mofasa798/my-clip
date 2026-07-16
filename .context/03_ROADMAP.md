# 03_ROADMAP.md

# Project Roadmap

This document defines the development roadmap for the project.

The roadmap is the single source of truth for development progress.

Only features within the **current active phase** should be implemented.

Future phases must not be implemented unless explicitly requested.

---

# Current Development Phase

## Active Phase

**Phase 1 — Project Foundation**

Status

🟢 Active

Current Objective

Build a stable project foundation before implementing any user-facing functionality.

---

# Development Workflow

Development follows a sequential approach.

```text
Phase 1

↓

Phase 2

↓

Phase 3

↓

Phase 4

↓

Phase 5

↓

Phase 6
```

A phase should be completed before moving to the next one.

Avoid skipping phases.

---

# Phase 1 — Project Foundation

## Goal

Establish the project foundation.

Create the application structure, configure development tools, verify external dependencies, and prepare the project architecture.

---

## Tasks

Project

* Initialize Wails v3
* Configure React
* Configure TypeScript
* Configure Tailwind CSS
* Configure Go project
* Setup Git repository

Application

* Application bootstrap
* Startup sequence
* Shutdown sequence

Dependencies

* Detect FFmpeg
* Detect FFprobe
* Detect yt-dlp
* Detect AMD AMF support

Configuration

* Create settings system
* Load configuration
* Save configuration
* Default configuration

Storage

* Create storage directories
* Verify write permissions

UI

* Home page
* Settings page

Logging

* Logger initialization
* Startup logs

---

## Deliverables

The application should:

* Launch successfully.
* Verify required dependencies.
* Detect GPU support.
* Save settings.
* Load settings.
* Display a basic UI.

---

## Exit Criteria

Phase 1 is complete when:

* Project builds successfully.
* Application launches without errors.
* Required binaries are detected.
* Settings persist correctly.
* Logging works.
* Folder structure is complete.

---

# Phase 2 — Video Discovery

## Goal

Allow users to discover and download videos.

---

## Tasks

Metadata

* Parse URLs
* Validate supported platforms
* Retrieve metadata
* Retrieve thumbnails
* Retrieve duration
* Retrieve available formats

Download

* Select quality
* Download video
* Download audio
* Merge streams when necessary

UI

* Metadata preview
* Thumbnail preview
* Resolution selector
* Download progress

---

## Deliverables

Users can:

* Load supported URLs.
* View metadata.
* Download videos locally.

---

## Exit Criteria

* Metadata retrieval is reliable.
* Downloads complete successfully.
* Progress updates function correctly.
* Cancellation works.

---

# Phase 3 — Clip Editor

## Goal

Provide an intuitive interface for selecting clip ranges.

---

## Tasks

Preview

* Video player
* Frame preview

Timeline

* Interactive timeline
* Start marker
* End marker
* Manual timestamp editing

Output

* Filename
* Output directory
* Export options

UI

* Responsive layout
* Timeline controls

---

## Deliverables

Users can visually select clip ranges before exporting.

---

## Exit Criteria

* Timeline works correctly.
* Timestamp selection is accurate.
* Preview playback is stable.
* Export options are functional.

---

# Phase 4 — Video Processing

## Goal

Generate clips efficiently.

---

## Tasks

Processing

* Stream Copy
* GPU Encoding
* CPU Fallback

Progress

* Progress updates
* Cancellation
* Completion events

Validation

* Output validation
* Error reporting

Optimization

* Minimize processing time
* Preserve video quality

---

## Deliverables

Users can export clips successfully.

The application automatically selects the optimal processing strategy.

---

## Exit Criteria

* Stream Copy works.
* GPU encoding works.
* CPU fallback works.
* Progress reporting works.
* Output validation succeeds.

---

# Phase 5 — User Experience

## Goal

Improve usability and workflow.

---

## Tasks

History

* Recent downloads
* Recent exports

Settings

* Theme
* Preferences
* Remember user choices

Convenience

* Open output folder
* Notifications
* Better error messages

Logging

* Log viewer

UI

* UX improvements
* Accessibility improvements

---

## Deliverables

The application provides a polished desktop experience.

---

## Exit Criteria

* Settings are remembered.
* History functions correctly.
* Notifications work.
* Error messages are user-friendly.

---

# Phase 6 — Advanced Features

## Goal

Implement optional power-user features.

These features are intentionally postponed until the core application is stable.

---

## Tasks

Video

* Playlist download
* Livestream clipping
* Batch clipping

Export

* Presets
* Watermark
* Subtitle burn-in
* Thumbnail generator

Encoding

* AV1 support
* Multi-export

Application

* Automatic update checker

---

## Deliverables

Advanced functionality is available without compromising application simplicity.

---

## Exit Criteria

All advanced features operate independently and do not negatively affect the core workflow.

---

# Deferred Features

The following features are intentionally postponed.

* Cloud synchronization
* Remote storage
* Team collaboration
* AI clip detection
* Automatic highlight generation
* OCR
* Speech recognition

These features should not be considered during normal development.

---

# Out of Scope

The following are permanently outside the scope of this project.

* Cloud deployment
* Authentication
* User accounts
* REST API
* GraphQL
* Database server
* PostgreSQL
* Redis
* RabbitMQ
* Kafka
* Docker
* Kubernetes
* Microservices
* Distributed processing

---

# Definition of Done

A task is considered complete when:

* The implementation works.
* Error handling is complete.
* Logging is implemented.
* The code follows project architecture.
* The code follows project rules.
* The feature has been manually tested.
* No unnecessary dependencies were introduced.

---

# Roadmap Rules

Always follow these rules:

* Only work on the active phase.
* Complete the current phase before moving forward.
* Do not implement future features.
* Do not skip milestones.
* Keep implementations incremental.
* Prefer small, reviewable changes.

---

# Progress Tracking

Progress should be tracked in:

* 05_SPRINT.md

Architecture decisions belong in:

* 04_DECISIONS.md

Implementation rules belong in:

* 01_RULES.md

This document should only describe **what** will be built and **when** it should be built.
