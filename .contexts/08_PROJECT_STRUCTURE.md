# 08_PROJECT_STRUCTURE.md

# Project Structure

This document defines the directory structure and organization of the project.

Every file should have a clear and predictable location.

The project structure reflects the architectural boundaries described in `02_ARCHITECTURE.md`.

---

# Project Layout

```text
project/

├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── internal/
│       ├── app/
│       ├── domain/
│       ├── source/
│       │   ├── resolver/
│       │   ├── registry/
│       │   ├── youtube/
│       │   └── kick/
│       ├── media/
│       │   ├── ffmpeg/
│       │   ├── ffprobe/
│       │   └── gpu/
│       ├── system/
│       └── shared/
│
├── frontend/
│   ├── src/
│   │   ├── pages/
│   │   ├── components/
│   │   ├── hooks/
│   │   ├── services/
│   │   ├── stores/
│   │   ├── types/
│   │   ├── constants/
│   │   └── assets/
│   └── package.json
│
├── docs/
│   ├── backend.md
│   ├── frontend.md
│   ├── sources.md
│   ├── ffmpeg.md
│   ├── gpu.md
│   ├── ui.md
│   └── design_system.md
│
├── .contexts/
│   ├── INDEX.md
│   ├── 00_PROJECT.md
│   ├── 01_RULES.md
│   ├── 02_ARCHITECTURE.md
│   ├── 03_ROADMAP.md
│   ├── 04_DECISIONS.md
│   ├── 05_SPRINT.md
│   ├── 06_PROMPT.md
│   ├── 07_CODE_STYLE.md
│   ├── 08_PROJECT_STRUCTURE.md
│   └── 09_TESTING.md
│
└── README.md
```

---

# Backend Structure

The backend is organized by architectural responsibility.

## app/

Application bootstrap.

Responsibilities:

* Startup
* Shutdown
* Dependency wiring
* Wails integration

No business logic belongs here.

---

## domain/

Business models and workflows.

Examples:

* Clip
* Export
* History
* Settings

The Domain Layer contains the application's core business rules.

---

## source/

Responsible for online video sources.

Contains:

* Source Resolver
* Source Registry
* Source implementations

Example:

```text
source/

resolver/

registry/

youtube/

kick/
```

Every source implements the same interface.

---

## media/

Responsible for local media processing.

Contains:

* FFmpeg wrapper
* FFprobe wrapper
* GPU support
* Media utilities

This layer never communicates with online video sources.

---

## system/

Operating system integration.

Examples:

* File system
* Configuration
* Logging
* External processes
* Temporary files

Infrastructure concerns belong here.

---

## shared/

Reusable components shared across multiple layers.

Examples:

* Utilities
* Generic helpers
* Common value objects
* Shared constants

Do not place business logic here.

---

# Frontend Structure

The frontend follows feature-oriented organization.

## pages/

Top-level application screens.

Examples:

* Home
* Editor
* History
* Settings

---

## components/

Reusable UI components.

Examples:

* Button
* Card
* Dialog
* Timeline
* ProgressBar

Components should remain small and composable.

---

## hooks/

Reusable React hooks.

Examples:

* useTimeline
* useExport
* useVideoPlayer

Hooks manage UI behavior only.

---

## services/

Frontend communication with the backend.

Responsibilities:

* Wails bindings
* Event subscriptions
* Request mapping

Business logic belongs in the backend.

---

## stores/

Global application state.

Examples:

* Theme
* UI state
* Export progress

Avoid storing business logic.

---

## types/

Shared TypeScript models.

These should mirror backend models whenever practical.

---

## constants/

Application constants.

Examples:

* Routes
* Event names
* UI configuration

---

## assets/

Static frontend assets.

Examples:

* Icons
* Images
* Fonts

---

# Documentation Structure

## .contexts/

Defines project standards.

These documents guide both developers and AI.

They represent the project's engineering rules and architectural decisions.

---

## docs/

Contains implementation-specific documentation.

Examples:

* Backend
* Frontend
* Source Layer
* FFmpeg
* GPU
* UI

These documents explain how individual subsystems work.

---

# Naming Conventions

Directories use:

* lowercase
* singular nouns
* descriptive names

Examples:

```text
source/

media/

domain/

system/
```

Avoid:

```text
sources/

platform/

modules/

misc/
```

---

# Package Organization

Prefer domain-oriented packages.

Good:

```text
clip/

history/

settings/

source/
```

Avoid:

```text
manager/

controller/

handler/

helper/
```

---

# File Organization

Every file should have a single responsibility.

Avoid extremely large files.

Recommended limits:

* Go files: approximately 300–500 lines
* React components: approximately 200–300 lines

Split files when responsibilities grow.

---

# Future Expansion

Adding a new source should require only:

```text
source/

vimeo/
```

or

```text
source/

twitch/
```

No structural changes should be required elsewhere.

---

# AI Guidelines

When generating code:

* Follow the documented directory structure.
* Respect architectural boundaries.
* Do not create alternative directory layouts.
* Reuse existing packages whenever possible.
* Keep packages focused on a single responsibility.
* Prefer extending existing domains over creating new top-level folders.

---

# Project Structure Philosophy

A predictable structure reduces cognitive load.

Every directory should communicate its purpose immediately.

Developers and AI should be able to determine where new code belongs without introducing new architectural patterns or unnecessary complexity.
