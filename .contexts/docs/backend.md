# backend.md

# Backend Engineering Guide

This document defines the backend architecture, responsibilities, and engineering standards.

The backend is the heart of the application.

It coordinates workflows, owns all business logic, and manages communication between the Platform Layer, Media Layer, and System Layer.

The frontend must never contain business logic.

---

# Backend Responsibilities

The backend owns:

* Application lifecycle
* Business logic
* Platform orchestration
* Download orchestration
* Media processing
* GPU selection
* Settings management
* History management
* Logging
* File management
* Progress reporting

---

# Backend Architecture

```text id="q1xp4r"
Application

↓

Domain

↓

Platform

↓

Media

↓

System
```

Dependencies always point downward.

---

# Layer Responsibilities

## Application

Coordinates workflows.

Examples:

* Download workflow
* Export workflow
* Settings workflow

The Application layer coordinates services but should contain minimal business rules.

---

## Domain

Contains business logic.

Examples:

* Clip creation
* Export requests
* Download requests
* Settings validation
* History management

The Domain layer must never know which video platform is being used.

---

## Platform

Responsible for supported video providers.

Responsibilities:

* Detect platform
* Validate URL
* Retrieve metadata
* Retrieve streams
* Download media

Every provider implements the same interface.

Examples:

* YouTube
* Kick

The rest of the backend communicates only through the Platform interface.

---

## Media

Responsible for local media operations.

Examples:

* FFmpeg
* FFprobe
* GPU detection
* Export
* Media inspection

The Media layer operates only on local files.

---

## System

Responsible for operating system interaction.

Examples:

* File system
* Logging
* Background workers
* Process execution

This layer should contain no business logic.

---

# Suggested Directory Structure

```text id="h2wbmz"
backend/

main.go

internal/

    app/

    domain/

    platform/

    media/

    system/

    shared/
```

Each package should own one domain.

---

# Platform Module

The Platform module is responsible for online video providers.

Suggested structure:

```text id="ewzy5u"
platform/

resolver/

adapters/

    youtube/

    kick/
```

Each adapter should expose identical capabilities.

---

# Domain Services

Examples:

```text id="8zjlwm"
ClipService

DownloadService

HistoryService

SettingsService
```

Services should describe business capabilities.

Avoid names such as:

```text id="x9c94n"
Manager

Helper

Processor

Utility
```

---

# Platform Flow

```text id="gdjlwm"
Video URL

↓

Platform Resolver

↓

Platform Adapter

↓

Video Metadata

↓

Download
```

Business services should never communicate directly with platform implementations.

---

# Media Flow

```text id="ck8m8r"
Local Media

↓

FFprobe

↓

Clip Service

↓

FFmpeg

↓

Export
```

Media processing remains platform-independent.

---

# System Flow

```text id="b9jlwm"
Media Layer

↓

Filesystem

↓

Process Execution

↓

Operating System
```

The System layer should expose reusable capabilities.

---

# External Processes

External binaries include:

* FFmpeg
* FFprobe
* yt-dlp

Rules:

* Execute only through dedicated wrappers.
* Support context cancellation.
* Capture stdout and stderr.
* Return structured errors.

Never execute external binaries from business services.

---

# Dependency Rules

Allowed:

```text id="tkjlwm"
Application

↓

Domain

↓

Platform

↓

Media

↓

System
```

Forbidden:

```text id="h9rjlwm"
Platform

↓

React
```

Forbidden:

```text id="mjlwm3"
Media

↓

Platform
```

Forbidden:

```text id="vjlwm4"
System

↓

Domain
```

Dependencies must remain acyclic.

---

# Service Design

Every service should:

* Own one responsibility
* Expose a small public API
* Hide implementation details
* Return structured errors

Avoid large services.

---

# State Management

Keep services stateless whenever practical.

Persistent state belongs only in:

* Settings
* History

Avoid global variables.

---

# Configuration

Configuration should be centralized.

Examples:

* Output directory
* Preferred encoder
* Theme
* Download options

Never hardcode configuration values.

---

# Logging

Centralize logging.

Log meaningful events only.

Examples:

* Startup
* Shutdown
* Platform detection
* Download started
* Export completed
* Errors

Avoid verbose logging.

---

# Event System

The backend emits events through Wails.

Examples:

```text id="kjlwm5"
download.started

download.progress

download.completed

export.started

export.progress

export.completed
```

The frontend subscribes to these events.

---

# Validation

Validate as early as possible.

Examples:

* URL
* Output directory
* Timestamp range
* Export options
* Encoder availability

Reject invalid requests before starting long-running operations.

---

# Error Handling

Errors should include:

* Context
* Cause
* Recovery information (when appropriate)

Avoid generic errors such as:

```text id="sjlwm6"
unknown error
```

---

# Background Workers

Workers should be used only for long-running operations.

Examples:

* Download
* Export
* Metadata retrieval

Workers should:

* Support cancellation
* Report progress
* Return structured results

---

# Performance

Optimize for:

* Fast startup
* Low memory usage
* Efficient disk I/O
* Minimal process spawning

Avoid unnecessary abstractions.

---

# Testing

Business logic should be tested independently of external tools.

Mock:

* FFmpeg
* FFprobe
* yt-dlp
* Filesystem

Do not mock business rules.

---

# AI Guidelines

When generating backend code:

* Keep layers isolated.
* Preserve dependency direction.
* Reuse existing services.
* Keep business logic inside the Domain layer.
* Treat platforms as interchangeable providers.
* Keep media processing platform-independent.
* Never execute external tools outside dedicated wrappers.

---

# Backend Philosophy

The backend exists to orchestrate workflows, not to expose implementation details.

Platform-specific logic ends at the Platform Layer.

From that point forward, the application works exclusively with generic media and business models.

A clean backend is one where replacing a video provider does not affect the rest of the application.
