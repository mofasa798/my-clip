# backend.md

# Backend Engineering Guide

This document defines the engineering standards, responsibilities, and implementation guidelines for the Go backend.

The backend owns all business logic, media processing, source integration, and communication with the operating system.

The frontend should remain a presentation layer.

---

# Responsibilities

The backend is responsible for:

* Business logic
* Source resolution
* Media processing
* Export pipeline
* Hardware capability detection
* Settings management
* History management
* Communication with external tools

The backend should remain the single source of truth.

---

# Backend Architecture

```text
Frontend (React)

↓

Wails Bindings

↓

Application

↓

Domain

↓

Source
Media

↓

System
```

Business logic must never exist above the Domain layer.

---

# Core Principles

The backend should be:

* Source-agnostic
* Local-first
* Capability-driven
* Predictable
* Explicit
* Easy to test
* Easy to extend

Favor simplicity over abstraction.

---

# Layer Responsibilities

## Application

Coordinates use cases.

Examples:

* Load metadata
* Download media
* Export clip
* Save settings

Application services orchestrate workflows.

---

## Domain

Owns business rules.

Examples:

* Clip validation
* Export validation
* Workflow coordination
* Generic media models

The Domain layer must not depend on infrastructure.

---

## Source

Responsible for online video sources.

Responsibilities:

* Source detection
* URL validation
* Metadata retrieval
* Download

The Source layer ends once local media is available.

---

## Media

Responsible for local media processing.

Responsibilities:

* FFmpeg wrapper
* FFprobe wrapper
* Hardware detection
* Export pipeline
* Stream Copy

The Media layer operates exclusively on local files.

---

## System

Responsible for infrastructure.

Examples:

* File system
* Logging
* Configuration
* External process execution

Business logic never belongs here.

---

# Data Flow

```text
Video URL

↓

Source

↓

Video Metadata

↓

Local Media

↓

Clip Request

↓

Export

↓

History
```

The workflow should remain consistent regardless of the original source.

---

# Service Design

Each service should have one responsibility.

Good examples:

```text
MetadataService

DownloadService

ExportService

HistoryService

SettingsService
```

Avoid services that combine unrelated responsibilities.

---

# Dependency Rules

Allowed:

```text
Application

↓

Domain

↓

Source
Media

↓

System
```

Never introduce upward dependencies.

---

# Source Integration

Every supported source implements the same interface.

Example:

```go
type Source interface {
    Name() string
    Match(url string) bool
    Metadata(ctx context.Context, url string) (*VideoMetadata, error)
    Download(ctx context.Context, req DownloadRequest) (*DownloadResult, error)
}
```

Business services should only depend on this interface.

---

# Generic Models

Backend services should exchange generic models.

Examples:

* VideoMetadata
* StreamInfo
* MediaFile
* ClipRequest
* ExportOptions
* ExportResult

Avoid source-specific models outside the Source layer.

---

# External Tools

External tools are infrastructure.

Examples:

* yt-dlp
* FFmpeg
* FFprobe

Only dedicated wrappers should invoke them.

No business service should execute external commands directly.

---

# Hardware Strategy

The backend should detect available hardware capabilities dynamically.

Supported acceleration may include:

* NVIDIA NVENC
* AMD AMF
* Intel Quick Sync Video

Hardware vendors should never be hardcoded into business logic.

---

# Processing Strategy

Preferred order:

```text
Stream Copy

↓

GPU Encoding

↓

CPU Encoding
```

The backend should automatically select the most efficient available method.

---

# Event Communication

The backend communicates progress through Wails events.

Examples:

```text
metadata.loaded

download.started

download.progress

download.completed

export.started

export.progress

export.completed
```

Avoid polling whenever possible.

---

# Error Handling

Errors should be meaningful.

Good examples:

* Invalid source URL
* Unsupported source
* Metadata unavailable
* Download failed
* Export failed

Internal implementation details should remain hidden.

---

# Concurrency

Long-running operations should execute asynchronously.

Examples:

* Downloads
* Metadata retrieval
* Export
* Hardware detection

The UI should remain responsive throughout these operations.

---

# Configuration

Application settings belong in dedicated configuration models.

Examples:

* Output directory
* Preferred encoder
* Temporary directory
* Theme
* Export defaults

Avoid scattered configuration values.

---

# Logging

Log entries should provide operational insight.

Log:

* Application startup
* Source resolution
* Download progress
* Export lifecycle
* Hardware detection
* Unexpected failures

Avoid excessive logging.

---

# Testing

Prioritize testing for:

* Business rules
* Source adapters
* Media wrappers
* Export workflow
* Error handling

External tools should be covered by integration tests.

---

# Performance

The backend should:

* Minimize unnecessary allocations
* Avoid duplicate work
* Stream data when practical
* Clean temporary resources
* Prefer efficient processing paths

Performance optimizations should never compromise readability.

---

# AI Guidelines

When generating backend code:

* Keep services focused.
* Preserve layer boundaries.
* Use generic models.
* Keep source-specific logic inside the Source layer.
* Keep FFmpeg inside the Media layer.
* Avoid global state.
* Prefer explicit dependencies.
* Reuse existing services before introducing new ones.

---

# Backend Philosophy

The backend owns the application's intelligence.

It is responsible for understanding video sources, processing media, coordinating workflows, and exposing a clean interface to the frontend.

The frontend should request actions.

The backend should decide how those actions are performed.
