# 02_ARCHITECTURE.md

# System Architecture

This document defines the architectural foundations of the application.

The architecture prioritizes simplicity, maintainability, extensibility, and clear separation of responsibilities.

The system is designed around **generic media workflows**, not individual video sources.

---

# Architecture Principles

The architecture follows these principles:

* Source-agnostic
* Layered architecture
* Domain-driven workflows
* Local-first processing
* Explicit dependencies
* Composition over inheritance
* Simplicity over abstraction

Every layer owns a single responsibility.

---

# High-Level Architecture

```text
                  User Interface
                         │
                         ▼
                Application Layer
                         │
                         ▼
                   Domain Layer
                  ┌─────────────┐
                  ▼             ▼
             Source Layer   Media Layer
                  └──────┬──────┘
                         ▼
                    System Layer
```

Business logic belongs to the Domain Layer.

Infrastructure details belong to the lower layers.

---

# Layer Responsibilities

## User Interface

Responsible for:

* User interaction
* Navigation
* Timeline editing
* Video preview
* Progress visualization
* Settings pages

The UI never owns business logic.

---

## Application Layer

Responsible for coordinating application use cases.

Examples:

* Start download
* Create clip
* Export media
* Load settings
* Restore history

The Application Layer orchestrates workflows but should contain minimal business rules.

---

## Domain Layer

The Domain Layer is the heart of the application.

Responsibilities:

* Business rules
* Workflow orchestration
* Validation
* Generic media models
* Export requests
* Clip requests

The Domain Layer must never depend on any specific video source.

---

## Source Layer

Responsible for interacting with supported online video sources.

Responsibilities:

* Source detection
* URL validation
* Metadata retrieval
* Stream discovery
* Media download

Examples:

* YouTube
* Kick

After download completes, source-specific behavior ends.

---

## Media Layer

Responsible for processing local media.

Responsibilities:

* FFmpeg wrapper
* FFprobe wrapper
* Hardware detection
* Clip generation
* Export pipeline
* Media inspection

The Media Layer only works with local media files.

---

## System Layer

Responsible for operating system interaction.

Examples:

* File system
* External processes
* Logging
* Configuration
* Background workers

No business logic belongs here.

---

# Dependency Rules

Allowed:

```text
UI

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

Dependencies always point downward.

Lower layers must never depend on higher layers.

---

# Source Workflow

```text
Video URL

↓

Source Resolver

↓

Source Adapter

↓

Video Metadata

↓

Download

↓

Local Media
```

After this point, the source becomes an implementation detail.

---

# Media Workflow

```text
Local Media

↓

FFprobe

↓

Clip Request

↓

FFmpeg

↓

Export
```

Media processing is completely independent of the originating source.

---

# Export Strategy

Preferred processing order:

```text
Stream Copy

↓

GPU Encoding

↓

CPU Encoding
```

The application should always select the fastest available strategy.

---

# Hardware Strategy

The application should never assume a specific hardware configuration.

At runtime it should automatically detect available capabilities.

Preferred order:

```text
Stream Copy Available

↓

Yes

↓

Stream Copy

↓

No

↓

Hardware Encoder Available

↓

Yes

↓

GPU Encoding

↓

No

↓

CPU Encoding
```

Hardware capabilities determine processing strategy.

---

# Backend Architecture

Recommended structure:

```text
backend/

main.go

internal/

    app/

    domain/

    source/

        resolver/

        registry/

        adapters/

    media/

        ffmpeg/

        ffprobe/

        gpu/

    system/

    shared/
```

Each directory represents an architectural boundary.

---

# Frontend Architecture

Recommended structure:

```text
frontend/

src/

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

```text
User

↓

React UI

↓

Application

↓

Domain

↓

Source

↓

Local Media

↓

Media

↓

Export

↓

History
```

Every workflow should follow this pattern.

---

# Event Flow

```text
Backend

↓

Wails Events

↓

Frontend Services

↓

React State

↓

UI
```

The frontend should react to backend events rather than polling.

---

# Error Flow

Errors propagate upward.

```text
System

↓

Media

↓

Source

↓

Domain

↓

Application

↓

User Interface
```

Each layer should add meaningful context without exposing implementation details.

---

# Extension Strategy

Adding support for a new source should require only:

1. Implement a Source Adapter.
2. Register the adapter.
3. Verify metadata retrieval.
4. Verify download.

No changes should be required in:

* Domain
* Media Layer
* Export pipeline
* UI

---

# Architectural Constraints

The application must never:

* Mix source logic with media processing.
* Execute external binaries outside dedicated wrappers.
* Place business logic in the frontend.
* Couple the Domain Layer to any specific source.
* Bypass architectural boundaries.

---

# AI Guidelines

When generating architecture-related code:

* Keep business logic inside the Domain Layer.
* Keep source implementations isolated.
* Treat media processing as source-independent.
* Prefer explicit interfaces.
* Avoid unnecessary abstractions.
* Respect dependency direction.
* Preserve architectural boundaries.

---

# Architecture Philosophy

The application is designed around **generic media workflows**.

Video sources are interchangeable implementations.

Once media has been downloaded, its origin is no longer relevant.

The architecture should make adding a new source inexpensive while keeping the core application stable, predictable, and easy to maintain.
