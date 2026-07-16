# 08_PROJECT_STRUCTURE.md

# Project Structure

This document defines the physical organization of the project.

Its purpose is to provide a predictable structure for both developers and AI assistants.

Every directory should have a single, well-defined responsibility.

---

# Design Principles

The project structure should be:

* Domain-oriented
* Platform-agnostic
* Easy to navigate
* Easy to extend
* Consistent

Avoid unnecessary nesting.

Avoid generic folders.

---

# Top-Level Structure

```text
project/

.contexts/
docs/

backend/
frontend/

bin/
storage/
scripts/

go.mod
package.json
README.md
```

---

# Directory Overview

## .contexts

Engineering documentation for AI and developers.

Contains:

* Project rules
* Architecture
* Roadmap
* ADR
* Sprint
* Code style
* Testing

No source code.

---

## docs

Technical documentation.

Examples:

* backend.md
* frontend.md
* ffmpeg.md
* gpu.md
* platform.md
* ui.md
* design_system.md

---

## backend

Contains the complete Go application.

All business logic lives here.

---

## frontend

Contains the React application.

Presentation only.

---

## bin

External executables.

Examples:

```text
ffmpeg
ffprobe
yt-dlp
```

Platform-specific binaries may be organized into subdirectories.

---

## storage

Application data.

Examples:

```text
settings/
history/
cache/
exports/
temp/
logs/
```

Generated files should never be committed.

---

## scripts

Development utilities.

Examples:

* build
* release
* clean

---

# Backend Structure

```text
backend/

main.go

internal/

    app/

    domain/

        clip/
        download/
        history/
        settings/

    platform/

        resolver/

        adapters/

            youtube/
            kick/

    media/

        ffmpeg/
        ffprobe/
        gpu/

    system/

        filesystem/
        logger/
        worker/

    shared/
```

The backend is organized by architectural boundaries rather than implementation details.

---

# Backend Responsibilities

## app

Application lifecycle.

* Startup
* Shutdown
* Dependency initialization
* Service registration

---

## domain

Core business logic.

Contains:

* Clip creation
* Download workflow
* Settings
* History

The domain layer must never depend on platform-specific implementations.

---

## platform

Responsible for supported video platforms.

Contains:

* Platform detection
* URL validation
* Metadata retrieval
* Stream discovery
* Media download

Each supported platform implements the same interface.

Examples:

```text
youtube/
kick/
```

Future platforms should be added here.

---

## media

Responsible for local media operations.

Contains:

* FFmpeg wrapper
* FFprobe wrapper
* GPU detection
* Media inspection
* Export pipeline

The media layer never communicates directly with online platforms.

---

## system

Responsible for operating system interaction.

Contains:

* File system
* Logging
* Background workers

System services should remain reusable.

---

## shared

Small reusable utilities.

Allowed:

* Constants
* Shared types
* Small helper functions

Forbidden:

* Business logic
* Platform logic
* FFmpeg execution

---

# Frontend Structure

```text
frontend/

src/

assets/
components/
constants/
hooks/
layouts/
pages/
services/
stores/
types/
utils/
```

The frontend should remain shallow.

Avoid deeply nested folders.

---

# Frontend Responsibilities

## assets

Static resources.

* Images
* Icons
* Fonts

---

## components

Reusable UI components.

Examples:

* Timeline
* ProgressBar
* VideoPreview
* MetadataCard

---

## constants

Application constants.

Examples:

* Event names
* Default values
* Theme tokens

---

## hooks

Reusable React hooks.

Examples:

* useDownload
* useTimeline
* useExport

---

## layouts

Shared page layouts.

---

## pages

Top-level application pages.

Examples:

* Home
* Editor
* History
* Settings

---

## services

Frontend communication layer.

Responsible for:

* Wails bindings
* Backend communication

No business logic.

---

## stores

Global UI state.

Examples:

* Theme
* Sidebar state
* Current page

Business state belongs in the backend.

---

## types

Shared TypeScript types.

---

## utils

Small utility functions.

No business logic.

---

# Platform Adapter Structure

Every supported platform should follow the same structure.

Example:

```text
platform/

adapters/

    youtube/

        adapter.go
        metadata.go
        download.go

    kick/

        adapter.go
        metadata.go
        download.go
```

This consistency simplifies future platform additions.

---

# Data Storage Structure

```text
storage/

cache/
exports/
history/
logs/
settings/
temp/
```

Each directory has a single responsibility.

---

# Folder Creation Rules

Before creating a new folder:

Ask:

* Does an existing folder already fit?
* Will the new folder contain multiple related files?
* Does it represent a new domain?

If the answer is "No", do not create it.

---

# File Creation Rules

Create new files only when:

* Responsibility changes.
* Reuse justifies separation.
* File size becomes difficult to maintain.

Avoid speculative files.

---

# Naming Conventions

Folders

* lowercase

Files

* lowercase_with_underscores

Packages

* singular

React Components

* PascalCase

Hooks

* useSomething

Interfaces

* Descriptive names

Avoid generic names such as:

* Helper
* Manager
* Common
* UtilsService

---

# Dependency Rules

Allowed:

```text
Frontend

↓

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

Forbidden:

```text
Platform

↓

React
```

Forbidden:

```text
Media

↓

Platform Adapter
```

Forbidden:

```text
System

↓

Domain
```

Lower layers must never depend on higher layers.

---

# Future Expansion

Adding a new platform should require only:

1. Create a new adapter.
2. Register it in the Platform Resolver.

No changes should be required to:

* Clip
* Export
* FFmpeg
* GPU
* UI

---

# AI Guidelines

When generating code:

* Follow the defined directory structure.
* Reuse existing packages.
* Keep responsibilities focused.
* Avoid creating generic folders.
* Respect architectural boundaries.
* Keep platform implementations isolated.

Never reorganize the project without explicit approval.

---

# Project Structure Philosophy

The directory structure reflects the architecture.

Every folder should have an obvious responsibility.

If a file has no obvious location, improve the structure instead of creating a miscellaneous folder.

A predictable structure is more valuable than a clever one.
