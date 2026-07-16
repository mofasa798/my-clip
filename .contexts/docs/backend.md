# backend.md

# Backend Engineering Guide

This document defines the backend implementation guidelines for the project.

The backend is responsible for all business logic, system orchestration, and communication with external tools.

The frontend must never contain business logic.

---

# Responsibilities

The backend owns:

* Application lifecycle
* Business logic
* Video processing
* Video downloading
* Metadata retrieval
* GPU detection
* Settings
* Logging
* File operations
* Progress reporting
* External process execution

The backend should expose a clean API to the frontend through Wails.

---

# Backend Directory

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
logger/
settings/
worker/
event/
utils/
```

Each package should have one responsibility.

---

# Package Responsibilities

## app

Responsible for:

* Startup
* Shutdown
* Dependency verification
* Initialization
* Service registration

---

## clip

Responsible for:

* Clip creation
* Export workflow
* Processing strategy
* Output validation

---

## download

Responsible for:

* Metadata retrieval
* Video downloading
* Audio downloading
* Progress reporting

---

## ffmpeg

Wrapper around FFmpeg.

Responsibilities:

* Build command arguments
* Execute FFmpeg
* Parse progress
* Return structured errors

No other package should directly execute FFmpeg.

---

## ffprobe

Wrapper around FFprobe.

Responsibilities:

* Metadata
* Codec information
* Duration
* Streams

---

## gpu

Responsible for:

* Detect GPU
* Detect supported encoders
* Select preferred encoder

---

## settings

Responsible for:

* Load settings
* Save settings
* Default values
* Validation

---

## logger

Responsible for:

* File logging
* Log formatting
* Log levels

Logging should be centralized.

---

## history

Responsible for:

* Recent downloads
* Recent exports

Should remain lightweight.

---

## worker

Responsible for:

* Long-running jobs
* Progress updates
* Cancellation

---

## event

Responsible for:

* Wails Events
* Progress events
* UI notifications

---

## utils

Utility functions only.

No business logic.

No external command execution.

---

# Service Design

Every service should represent a real business capability.

Examples

Good

```text
DownloadService
ClipService
SettingsService
GPUService
HistoryService
```

Avoid

```text
Manager
Helper
Processor
Utility
Common
```

---

# Service Responsibilities

Each service should own:

* One domain
* One responsibility
* One public API

Avoid giant services.

---

# Service Communication

Preferred flow

```text
UI

↓

Application Service

↓

Infrastructure Service

↓

External Tool
```

Never allow the UI to bypass services.

---

# External Commands

Only wrappers may execute:

* FFmpeg
* FFprobe
* yt-dlp

Every wrapper should:

* Validate arguments
* Capture stdout
* Capture stderr
* Support Context
* Return structured errors

---

# Worker Design

Workers are used only for long-running operations.

Examples

* Download
* Export
* Thumbnail extraction

Workers should:

* Report progress
* Support cancellation
* Return results

No scheduling system is required.

---

# Progress Reporting

Every long-running task should emit progress events.

Examples

```text
Download Started

↓

Progress

↓

Completed
```

or

```text
Export Started

↓

Progress

↓

Completed
```

Progress should be emitted through Wails Events.

---

# Error Handling

Every service should return meaningful errors.

Example

Good

```text
failed to detect FFmpeg executable
```

Good

```text
failed to retrieve video metadata
```

Avoid

```text
unknown error
```

---

# Context Usage

Every long-running service should accept:

context.Context

This allows:

* Cancellation
* Timeout
* Resource cleanup

---

# Configuration

Services should never hardcode:

* Directories
* Encoders
* Output paths

Read configuration through the Settings Service.

---

# File Operations

All file operations belong in the backend.

Examples

* Create directories
* Move files
* Delete temporary files
* Verify permissions

Never expose filesystem logic to React.

---

# Logging

Every important backend event should be logged.

Examples

* Startup
* Shutdown
* Dependency detection
* Download
* Export
* Error

Avoid excessive logging.

---

# Package Dependencies

Allowed dependency flow

```text
app

↓

services

↓

wrappers

↓

external executables
```

Forbidden

```text
ffmpeg

↓

React
```

or

```text
download

↓

UI
```

---

# Dependency Injection

Use constructor injection only.

Example

```go
NewDownloadService(...)
```

Do not introduce a dependency injection framework.

---

# State

Keep services stateless whenever possible.

Persistent state belongs in:

* Settings
* History

Avoid global variables.

---

# Models

Models should represent business entities.

Examples

```text
VideoMetadata
ClipRequest
ExportOptions
DownloadTask
Settings
HistoryEntry
```

Avoid generic names.

---

# Event Naming

Prefer descriptive event names.

Examples

```text
download.started
download.progress
download.completed

export.started
export.progress
export.completed
```

Avoid:

```text
update
change
done
```

---

# Validation

Validate input as early as possible.

Examples

* URL
* File path
* Output directory
* Clip duration
* Encoder

Return validation errors immediately.

---

# Temporary Files

Temporary files should:

* Be created in a temporary directory
* Be cleaned automatically
* Never overwrite user files

---

# Resource Management

Always:

* Close files
* Release resources
* Wait for processes
* Cancel contexts
* Remove temporary files

Resource leaks are unacceptable.

---

# Performance

Optimize for:

* Fast startup
* Low memory usage
* Efficient I/O
* Minimal process spawning

Avoid unnecessary allocations.

---

# Testing Strategy

Backend logic should be testable.

Separate:

* Business logic
* External command execution

Wrappers should isolate system dependencies.

---

# AI Guidelines

When creating backend code:

* Reuse existing services.
* Keep packages cohesive.
* Keep services focused.
* Keep APIs explicit.
* Prefer composition over abstraction.
* Do not create packages without a clear responsibility.
* Do not execute external binaries outside dedicated wrappers.
* Follow the dependency flow defined in the architecture.

---

# Backend Philosophy

The backend is the heart of the application.

Its responsibilities are to:

* Coordinate workflows.
* Execute external tools.
* Manage application state.
* Provide reliable business logic.
* Remain simple, modular, and maintainable.

A backend that is easy to understand is more valuable than one that is overly flexible.
