# 02_ARCHITECTURE.md

# System Architecture

This document defines the high-level architecture of the application.

Its purpose is to establish clear responsibilities, dependency boundaries, and communication flow between all major components.

The architecture is designed to be platform-agnostic, modular, and easy to extend.

---

# Architecture Principles

The application follows these principles:

* Platform-independent
* Layered architecture
* Single responsibility
* Explicit dependencies
* Local-first execution
* Composition over inheritance
* Simplicity over abstraction

Every layer should have one clear purpose.

---

# High-Level Architecture

```text id="0c6xra"
                 User Interface
                       │
                       ▼
              Application Layer
                       │
                       ▼
               Platform Layer
                       │
                       ▼
            Media Processing Layer
                       │
                       ▼
                System Layer
```

Dependencies always point downward.

No layer may bypass another.

---

# Layer Responsibilities

## User Interface

Responsible for:

* User interaction
* Timeline editing
* Video preview
* Progress display
* Settings
* Navigation

The UI never owns business logic.

---

## Application Layer

Coordinates application workflows.

Responsibilities:

* Clip creation
* Export workflow
* Download workflow
* Settings management
* History management
* Progress orchestration

The Application Layer coordinates services but does not implement platform-specific behavior.

---

## Platform Layer

The Platform Layer isolates every supported video platform.

Responsibilities:

* Detect platform
* Validate URL
* Retrieve metadata
* Retrieve available streams
* Download media

The rest of the application must never communicate directly with platform implementations.

---

## Media Processing Layer

Responsible for local media processing.

Responsibilities:

* Clip extraction
* Video encoding
* Stream copy
* GPU acceleration
* Metadata inspection

This layer operates only on local media.

It does not know where the media originated.

---

## System Layer

Responsible for operating system interaction.

Examples:

* File system
* External processes
* Logging
* Configuration
* Background workers

---

# Dependency Flow

```text id="lrbhtf"
React UI

↓

Application Services

↓

Platform Services

↓

Media Services

↓

System Services

↓

Operating System
```

Communication must always follow this direction.

---

# Platform Architecture

The application should support multiple platforms through adapters.

```text id="wwb55u"
Video URL

↓

Platform Resolver

↓

Platform Adapter

↓

Generic Video Source
```

Every supported platform should expose the same capabilities.

---

# Platform Adapter

Each adapter should implement the same interface.

Responsibilities:

* URL validation
* Metadata retrieval
* Stream discovery
* Download support

Platform adapters should not contain business logic.

---

# Generic Video Source

Once a video has been resolved, it becomes a generic media source.

Example model:

```text id="u4m3ax"
VideoMetadata

Platform

Title

Author

Duration

Thumbnail

Streams

Live

URL
```

Higher layers should only consume this model.

---

# Download Pipeline

```text id="rmfjnv"
Video URL

↓

Platform Resolver

↓

Platform Adapter

↓

Download Service

↓

Local Media File
```

The Download Service never knows which platform provided the media.

---

# Clip Pipeline

```text id="lya9v2"
Local Media File

↓

Timeline

↓

Clip Service

↓

FFmpeg

↓

Exported Clip
```

The Clip Pipeline operates exclusively on local files.

---

# Export Pipeline

Preferred processing strategy:

```text id="xgcrki"
Stream Copy

↓

GPU Encoding

↓

CPU Encoding
```

The Export Pipeline is completely platform-independent.

---

# Platform Independence

Platform-specific logic ends after download.

Everything beyond this point operates on generic local media.

This separation is mandatory.

---

# Backend Structure

```text id="31odlu"
backend/internal/

app/

domain/
    clip/
    download/
    history/
    settings/

platform/
    resolver/
    adapters/

media/
    ffmpeg/
    ffprobe/
    gpu/

system/
    logger/
    worker/
    filesystem/
```

Each directory represents a distinct architectural boundary.

---

# Frontend Structure

```text id="tezzjp"
frontend/src/

pages/
components/
hooks/
services/
stores/
types/
constants/
```

The frontend remains presentation-focused.

---

# Data Flow

```text id="k44ks4"
User Action

↓

UI

↓

Application Service

↓

Platform Layer

↓

Download

↓

Local File

↓

Media Processing

↓

Export

↓

History
```

Every workflow should follow this pattern.

---

# Event Flow

```text id="pm0g7g"
Backend

↓

Wails Events

↓

Frontend Services

↓

React State

↓

UI Update
```

The frontend should never poll backend progress.

---

# Error Flow

Errors should propagate upward.

```text id="mddjlwm"
System

↓

Media

↓

Platform

↓

Application

↓

User Interface
```

Every layer should enrich errors with useful context.

---

# Extension Strategy

Adding a new platform should require only:

1. Creating a new Platform Adapter.
2. Registering it in the Platform Resolver.

No changes should be required in:

* Clip Service
* Export Service
* FFmpeg Wrapper
* GPU Module
* UI

---

# Dependency Rules

Allowed:

```text id="7kqfjb"
Application

↓

Platform

↓

Media

↓

System
```

Forbidden:

```text id="svf3qv"
Platform

↓

React
```

Forbidden:

```text id="kqzsz8"
FFmpeg

↓

Platform Adapter
```

Lower layers must never depend on higher layers.

---

# Future Expansion

The architecture should support future platforms such as:

* Twitch
* Vimeo
* Other online video providers

without requiring changes to the application's core workflows.

---

# AI Guidelines

When generating architecture-related code:

* Keep platform logic isolated.
* Keep media processing platform-independent.
* Prefer explicit interfaces.
* Avoid unnecessary abstractions.
* Follow the defined dependency flow.
* Never bypass architectural boundaries.

---

# Architecture Philosophy

The architecture is centered around **generic video sources**, not individual platforms.

Platforms are interchangeable implementations.

The core application should only understand media, workflows, and user actions.

A new platform should integrate by implementing the Platform Adapter, while the rest of the application remains unchanged.
