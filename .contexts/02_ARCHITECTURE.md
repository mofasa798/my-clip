# 02_ARCHITECTURE.md

# Project Architecture

This document defines the high-level architecture of the application.

Its purpose is to ensure every module has a clear responsibility, maintain a clean dependency flow, and keep the project simple as it grows.

Detailed implementation belongs to each module and should not violate the architecture defined here.

---

# Architecture Philosophy

The architecture follows these principles:

* Local-first
* Modular
* Layered
* Maintainable
* Performance-oriented
* Minimal abstraction

The application is intentionally designed as a **single executable desktop application**.

There are no remote services.

There are no microservices.

There is no distributed processing.

---

# High-Level Architecture

```text
┌─────────────────────────────┐
│         React UI            │
└──────────────┬──────────────┘
               │
               ▼
┌─────────────────────────────┐
│           Wails             │
└──────────────┬──────────────┘
               │
               ▼
┌─────────────────────────────┐
│     Go Application Layer    │
└──────────────┬──────────────┘
               │
 ┌─────────────┴─────────────┐
 ▼                           ▼
Business Services      Infrastructure
                             │
          ┌──────────────────┴──────────────────┐
          ▼                                     ▼
      FFmpeg / FFprobe                     yt-dlp
```

---

# Layer Responsibilities

## Presentation Layer

Technology:

* React
* TypeScript
* TailwindCSS

Responsibilities:

* User Interface
* User Interaction
* Timeline
* Preview
* Progress
* Forms
* Notifications

Must NOT:

* Execute FFmpeg
* Execute yt-dlp
* Access files directly
* Implement business logic

---

## Application Layer

Technology:

* Go

Responsibilities:

* Business Logic
* Workflow orchestration
* Validation
* Settings
* History
* Event handling
* Progress updates

This layer coordinates all modules.

It should remain independent from UI implementation details.

---

## Infrastructure Layer

Responsibilities:

* FFmpeg
* FFprobe
* yt-dlp
* File System
* GPU Detection
* Process Execution

Every external dependency belongs here.

---

# Dependency Direction

Dependencies must always flow downward.

```text
Presentation

↓

Application

↓

Infrastructure
```

Never reverse this direction.

Forbidden:

Infrastructure → UI

Infrastructure → React

Infrastructure → Wails

Business Logic → React

---

# Module Responsibilities

## app

Application bootstrap.

Responsibilities:

* Startup
* Shutdown
* Initialization
* Dependency verification

---

## clip

Handles clip generation.

Responsibilities:

* Trim
* Export
* Processing strategy
* Output validation

---

## download

Handles downloading.

Responsibilities:

* Metadata
* Video download
* Thumbnail download
* Progress reporting

---

## ffmpeg

Wrapper around FFmpeg.

Responsibilities:

* Command generation
* Command execution
* Progress parsing
* Error handling

Never expose raw FFmpeg commands to other modules.

---

## ffprobe

Wrapper around FFprobe.

Responsibilities:

* Metadata
* Streams
* Duration
* Codec information

---

## gpu

GPU capability detection.

Responsibilities:

* Detect AMF
* Detect supported encoders
* Determine encoding strategy

---

## settings

Application configuration.

Responsibilities:

* Load configuration
* Save configuration
* Validate configuration
* Default values

---

## history

Stores application history.

Responsibilities:

* Recent downloads
* Recent exports

Should remain lightweight.

---

## worker

Runs background jobs.

Responsibilities:

* Download
* Export
* Thumbnail extraction

Worker Pool should remain simple.

No queue server.

No scheduler.

---

## utils

Reusable helper functions.

Must NOT contain business logic.

Avoid turning utils into a dumping ground.

---

# Directory Structure

```text
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
hooks/
pages/
types/

storage/

downloads/
output/
thumbnail/
logs/

bin/

ffmpeg.exe
ffprobe.exe
yt-dlp.exe
```

---

# Event Flow

Communication between Backend and Frontend uses Wails Events.

```text
Backend

↓

Emit Event

↓

Frontend

↓

Update UI
```

Examples:

* Download Progress
* Export Progress
* Completion
* Errors

No WebSocket is required.

---

# Processing Flow

Video download:

```text
User

↓

Paste URL

↓

yt-dlp

↓

Metadata

↓

Download

↓

Local File
```

Clip generation:

```text
User

↓

Select Range

↓

Determine Strategy

↓

FFmpeg

↓

Output File
```

---

# Video Processing Strategy

Always determine the most efficient strategy.

Priority:

1. Stream Copy
2. GPU Encoding
3. CPU Encoding

Never encode unless necessary.

---

# External Processes

Only dedicated wrappers may execute:

* FFmpeg
* FFprobe
* yt-dlp

The rest of the application must interact only through service methods.

Never construct command-line arguments outside these wrappers.

---

# File Storage

Application data:

```text
storage/

downloads/
output/
thumbnail/
logs/
```

Configuration:

```text
settings.json
```

History may be stored as JSON.

No database server is required.

---

# Error Flow

Every external command should return:

* Success
* Exit Code
* stdout
* stderr
* Duration
* Structured Error

Errors should propagate upward.

Presentation layer is responsible only for displaying them.

---

# Concurrency

Concurrency should remain simple.

Preferred model:

```text
User Request

↓

Worker Pool

↓

External Process

↓

Progress Event

↓

UI
```

Do not introduce complex scheduling systems.

Go Goroutines are sufficient.

---

# GPU Architecture

GPU detection occurs once during application startup.

Preferred encoder priority:

1. h264_amf
2. hevc_amf
3. libx264

GPU availability should be transparent to users.

---

# Configuration Flow

```text
Startup

↓

Load settings.json

↓

Validate

↓

Apply Defaults

↓

Ready
```

Configuration should never prevent application startup.

Invalid values should fall back to defaults whenever possible.

---

# Logging Flow

Application

↓

Logger

↓

storage/logs/

Log categories:

* Startup
* Download
* Export
* Settings
* External Process
* Errors

Logging should aid debugging without becoming excessively verbose.

---

# Extensibility

Future modules should integrate without changing existing architecture.

Examples:

* Playlist Download
* Livestream Support
* Batch Processing
* Subtitle Burn-in

New modules should follow the same dependency rules.

---

# Architecture Constraints

The following are intentionally excluded:

* REST API
* GraphQL
* Database Server
* Redis
* RabbitMQ
* Kafka
* Docker
* Kubernetes
* Plugin Framework
* Distributed Workers
* Microservices

The architecture is intentionally optimized for a single-user desktop application.

---

# Success Criteria

A successful architecture should provide:

* Clear module boundaries.
* Simple dependency flow.
* Minimal coupling.
* High cohesion.
* Easy navigation.
* Easy maintenance.
* Straightforward debugging.
* Predictable behavior.

If a new design increases complexity without providing significant value, it should not be adopted.
