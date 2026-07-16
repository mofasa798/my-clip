# 08_PROJECT_STRUCTURE.md

# Project Structure

This document defines the physical structure of the project.

Its purpose is to ensure all source code is organized consistently throughout the project's lifetime.

This document describes **where code belongs**, not how it works.

---

# Design Principles

The project structure should be:

* Simple
* Predictable
* Domain-oriented
* Easy to navigate
* Easy to maintain

Avoid unnecessary nesting.

---

# Top-Level Structure

```text id="9jzxtv"
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

# Directory Responsibilities

## .contexts

Contains AI documentation.

Examples:

* Project rules
* Architecture
* Sprint
* Code style

No source code.

---

## docs

Contains technical documentation.

Examples:

* FFmpeg
* GPU
* Backend
* Frontend
* UI

No source code.

---

## backend

Contains all Go application code.

Business logic belongs here.

---

## frontend

Contains the React application.

Presentation only.

---

## storage

Application data.

Examples:

* Settings
* History
* Temporary files
* Export output

Do not commit generated files.

---

## bin

External executables.

Examples:

```text id="lw7hpg"
ffmpeg

ffprobe

yt-dlp
```

Platform-specific binaries may be stored in subdirectories if needed.

---

## scripts

Development scripts.

Examples:

* Build
* Clean
* Release

---

# Backend Structure

```text id="8ebn0q"
backend/

main.go

internal/

    app/
    clip/
    download/
    event/
    ffmpeg/
    ffprobe/
    gpu/
    history/
    logger/
    settings/
    worker/
    shared/
```

Everything inside `internal/` is private to the application.

---

# Package Responsibilities

## app

Application lifecycle.

Startup.

Shutdown.

Initialization.

---

## clip

Clip creation.

Export orchestration.

Processing pipeline.

---

## download

Metadata.

Download.

Video retrieval.

---

## ffmpeg

FFmpeg wrapper.

Argument builder.

Progress parser.

---

## ffprobe

Media inspection.

Metadata extraction.

---

## gpu

GPU detection.

Encoder selection.

Capability reporting.

---

## history

Export history.

Download history.

---

## logger

Logging.

Log formatting.

---

## settings

Settings management.

Configuration.

Validation.

---

## worker

Background jobs.

Progress.

Cancellation.

---

## event

Wails Events.

Frontend notifications.

---

## shared

Reusable backend utilities.

Allowed:

* Constants
* Small helper functions
* Shared models

Forbidden:

* Business logic
* External command execution

---

# Frontend Structure

```text id="dfe68v"
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

Keep the frontend shallow.

Avoid deeply nested folders.

---

# Frontend Responsibilities

## assets

Images.

Fonts.

Icons.

---

## components

Reusable UI components.

---

## constants

Application constants.

Examples:

* Event names
* Default values
* Route names

---

## hooks

Reusable React hooks.

---

## layouts

Shared layouts.

---

## pages

Application screens.

---

## services

Backend communication.

Wails bindings.

---

## stores

Global UI state.

No business logic.

---

## types

Shared TypeScript types.

---

## utils

Small helper functions.

---

# Folder Creation Rules

Before creating a new folder:

Ask:

1. Does an existing folder already fit?
2. Is the new folder responsible for one domain?
3. Will it contain multiple files?

If the answer is "No", do not create it.

---

# File Creation Rules

Create a new file only when:

* Responsibility changes.
* File size becomes difficult to maintain.
* Reuse justifies separation.

Avoid creating files prematurely.

---

# Package Rules

Each package should own one domain.

Good

```text id="hbm2ax"
clip

download

settings
```

Avoid

```text id="lfzrj5"
helpers

common

misc

manager
```

---

# Dependency Rules

Allowed flow

```text id="jlwm5r"
Frontend

↓

Wails

↓

App

↓

Domain Services

↓

Infrastructure Wrappers

↓

External Tools
```

Forbidden

```text id="j1cyrq"
React

↓

FFmpeg
```

Forbidden

```text id="kaxqgi"
Download

↓

UI
```

Dependencies should always point downward.

---

# Business Logic

Business logic belongs only in:

```text id="2gnb1x"
backend/internal/
```

Never inside:

* React components
* React hooks
* Utility functions

---

# External Executables

Only the wrappers may execute:

* FFmpeg
* FFprobe
* yt-dlp

Never execute external binaries elsewhere.

---

# Models

Business models belong to their owning package.

Avoid creating a global models directory.

Example

```text id="o9bnx8"
clip/

ClipRequest

ClipResult
```

instead of

```text id="0m5h8u"
models/

everything.go
```

---

# Configuration Files

Store configuration under:

```text id="6x3dwp"
storage/

settings.json
history.json
```

Future configuration files should remain here.

---

# Temporary Files

Use:

```text id="aq9j5u"
storage/temp/
```

Temporary files must be cleaned automatically.

---

# Export Directory

Default location:

```text id="9ehpt4"
storage/exports/
```

Users may configure another output directory.

---

# Naming Conventions

Folders

* lowercase

Files

* lowercase_with_underscores

Packages

* singular

Components

* PascalCase

Hooks

* useSomething

---

# AI Guidelines

Before creating a file:

* Check whether an existing location is appropriate.
* Avoid duplicate responsibilities.
* Follow the defined structure.
* Keep packages cohesive.
* Keep folders domain-oriented.

Never reorganize the project structure without explicit approval.

---

# Future Expansion

If new features require additional folders:

* Introduce the smallest possible change.
* Document the new structure.
* Update this document.
* Keep the hierarchy simple.

Avoid speculative folder creation.

---

# Project Philosophy

A good project structure minimizes the number of decisions developers must make.

Every file should have an obvious home.

If the correct location is unclear, the structure should be improved rather than working around it.

Consistency is more valuable than clever organization.
