# 07_CODE_STYLE.md

# Code Style Guide

This document defines the preferred coding style for the project.

Its purpose is to ensure all generated code follows a consistent style regardless of when or how it is produced.

When existing project code differs from this guide, follow the existing codebase unless instructed otherwise.

---

# General Principles

Prefer code that is:

* Simple
* Explicit
* Readable
* Predictable
* Easy to maintain

Code is read far more often than it is written.

Optimize for readability.

---

# Simplicity

Always choose the simplest implementation.

Avoid solving problems that do not yet exist.

Avoid speculative abstractions.

Good:

* Small functions
* Clear logic
* Straightforward control flow

Bad:

* Clever tricks
* Deep inheritance
* Generic frameworks
* Overly flexible APIs

---

# File Size

Recommended limits.

Go files

* Target: 100–300 lines
* Soft limit: 500 lines

React Components

* Target: 50–200 lines

Functions

* Target: 10–40 lines

These are guidelines, not strict limits.

---

# File Naming

Use lowercase.

Separate words with underscores.

Examples

```text
clip_service.go
gpu_detector.go
download_manager.go
settings_store.go
```

React

```text
home_page.tsx
settings_page.tsx
video_preview.tsx
timeline.tsx
```

Avoid:

```text
Helper.go
Manager.go
Util.go
Misc.go
```

---

# Package Naming

Packages should be singular.

Examples

```text
clip
download
gpu
worker
settings
history
ffmpeg
ffprobe
```

Avoid

```text
helpers
utilities
common
misc
```

---

# Function Naming

Functions should describe actions.

Good

```go
LoadSettings()

SaveSettings()

DetectGPU()

ExportClip()

DownloadVideo()
```

Bad

```go
Run()

Handle()

Execute()

Do()

Process()
```

---

# Variable Naming

Use descriptive names.

Good

```go
outputDirectory

downloadProgress

videoMetadata

clipDuration
```

Avoid

```go
d

tmp

obj

val

x
```

---

# Constants

Group related constants together.

Prefer typed constants where appropriate.

Example

```go
const (
    EncoderCPU = "libx264"
    EncoderAMF = "h264_amf"
)
```

---

# Error Handling

Return errors immediately.

Example

```go
metadata, err := loader.Load(url)
if err != nil {
    return err
}
```

Avoid

```go
if err == nil {
    ...
}
```

Keep the happy path obvious.

---

# Error Messages

Errors should explain:

* What failed
* Why it failed (if known)

Good

```text
failed to detect FFmpeg executable
```

Good

```text
failed to export clip using h264_amf
```

Avoid

```text
error
```

Avoid

```text
failed
```

---

# Logging Style

Log meaningful events.

Good

```text
Application started

FFmpeg detected

Export completed
```

Avoid

```text
Starting...

Done

Running...
```

Logs should help debugging.

---

# Comments

Write comments only when they add context.

Good

```go
// Detect available GPU encoders during startup.
```

Avoid

```go
// Increment i.
i++
```

Code should explain itself whenever possible.

---

# Function Design

Each function should perform one task.

Good

```text
Load Settings

↓

Validate

↓

Return
```

Avoid

```text
Load

↓

Validate

↓

Save

↓

Log

↓

Notify UI

↓

Cleanup
```

Split responsibilities.

---

# Control Flow

Prefer early returns.

Good

```go
if err != nil {
    return err
}

return nil
```

Avoid deep nesting.

---

# Switch Statements

Prefer switch over long if chains.

Example

```go
switch encoder {
case EncoderAMF:
    ...
case EncoderCPU:
    ...
}
```

---

# Struct Design

Structs should model real concepts.

Good

```go
Clip

DownloadTask

Settings

VideoMetadata
```

Avoid

```go
Data

Manager

Helper

Object
```

---

# Interfaces

Create interfaces only when multiple implementations exist.

Prefer concrete types.

Bad

```go
type Downloader interface{}
```

Good

```go
type Downloader interface {
    Download(context.Context, string) error
}
```

Only if multiple implementations are expected.

---

# Concurrency

Keep concurrency explicit.

Good

```text
Worker Pool

↓

Job

↓

Progress

↓

Done
```

Avoid hidden goroutines.

Every goroutine should have a clear lifecycle.

---

# External Commands

Never execute external commands directly from business logic.

Always use wrappers.

Good

```text
DownloadService

↓

YT-DLP Wrapper

↓

exec.CommandContext()
```

Never

```text
React

↓

exec.Command()
```

---

# React Components

Each component should have one responsibility.

Good

```text
Timeline

VideoPreview

ProgressBar

SettingsForm
```

Avoid giant pages containing all logic.

---

# React Hooks

Extract reusable logic into custom hooks.

Examples

```text
useSettings()

useDownload()

useExport()

useProgress()
```

Avoid large hooks with unrelated responsibilities.

---

# Tailwind CSS

Prefer utility classes.

Reuse existing design patterns.

Avoid excessive inline styling.

Keep spacing consistent.

---

# Imports

Keep imports organized.

Go

1. Standard library
2. Third-party
3. Internal packages

---

# Folder Structure

Each folder should represent one domain.

Avoid catch-all folders.

Bad

```text
helpers/
misc/
shared/
```

Good

```text
clip/
download/
gpu/
worker/
```

---

# Code Duplication

Before writing new code:

Check whether similar functionality already exists.

Prefer reuse.

Avoid copy-paste programming.

---

# Performance

Optimize only after correctness.

Prefer:

Correct

↓

Readable

↓

Fast

Premature optimization is discouraged.

---

# Configuration

Avoid hardcoded values.

Use configuration whenever appropriate.

Provide sensible defaults.

---

# Formatting

Use:

* gofmt
* prettier

Never manually fight the formatter.

Formatting should remain automatic.

---

# AI Guidelines

When generating code:

* Match the surrounding style.
* Reuse existing modules.
* Prefer incremental changes.
* Keep implementations small.
* Avoid unnecessary abstraction.
* Preserve consistency.

When multiple valid implementations exist:

Choose the one that is:

* Easiest to understand
* Easiest to debug
* Easiest to maintain

---

# Example Philosophy

Good code should feel like it was written by the same developer over time.

Consistency is more valuable than individual cleverness.

If unsure about style, prefer the simpler and more explicit implementation.
