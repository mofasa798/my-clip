# 00_PROJECT.md

# Project Overview

## Project Name

Multi-Platform Video Clipper

---

# Vision

Build a fast, reliable, and easy-to-use desktop application that allows users to download, preview, clip, and export videos from supported online platforms.

The application should provide a consistent user experience regardless of the original video platform.

All platform-specific implementations must remain isolated from the core application.

---

# Mission

Develop a local-first desktop application capable of:

* Detecting supported video platforms
* Retrieving video metadata
* Downloading videos
* Creating high-quality clips
* Exporting clips efficiently using GPU acceleration whenever possible

The application should prioritize performance, reliability, and simplicity.

---

# Core Principles

The project follows these principles:

* Multi-platform by design
* Platform-agnostic architecture
* Local-first processing
* Performance-oriented implementation
* Simple and maintainable codebase
* Minimal dependencies
* Consistent user experience

---

# Supported Platforms

Current targets:

* YouTube
* Kick

Future expansion may include additional platforms through the Platform Adapter layer without requiring changes to the core application.

---

# Primary Workflow

```text id="q5r3ha"
Paste Video URL

↓

Detect Platform

↓

Load Metadata

↓

Download Video

↓

Preview Video

↓

Select Clip Range

↓

Export Clip

↓

Open Output Folder
```

Every feature should support this workflow.

---

# Platform Independence

The application should never assume a specific video platform.

After a URL has been resolved by the Platform Layer, every video should be treated as a generic video source.

The clipping, exporting, previewing, and processing pipelines must remain completely independent of the originating platform.

---

# Target Environment

Operating System

* Windows (Primary)

Future support:

* Linux
* macOS

Desktop Framework

* Wails v3

Execution Model

* Local Desktop Application

Cloud deployment is out of scope.

---

# Processing Philosophy

Video processing should follow this priority:

```text id="v1amhs"
Stream Copy

↓

GPU Encoding

↓

CPU Encoding
```

The application should always choose the fastest reliable processing strategy.

---

# Performance Goals

The application should:

* Start quickly
* Consume minimal memory
* Keep the UI responsive
* Utilize hardware acceleration when available
* Avoid unnecessary processing
* Handle large video files efficiently

---

# Architecture Philosophy

The application is divided into independent layers.

```text id="yjlwmk"
User Interface

↓

Application Layer

↓

Platform Layer

↓

Media Processing Layer

↓

System Layer
```

Each layer has a single responsibility.

Platform-specific logic must never leak into higher layers.

---

# User Experience Goals

The application should provide:

* Fast metadata retrieval
* Responsive timeline editing
* Smooth video preview
* Reliable downloads
* Predictable exports
* Clear progress reporting
* Friendly error messages

Users should be able to create clips with minimal configuration.

---

# Scope

Included:

* Multi-platform video support
* Metadata retrieval
* Video downloading
* Video preview
* Clip creation
* GPU-accelerated export
* Export history
* Configurable settings

Excluded:

* Cloud services
* User accounts
* Video uploading
* Collaborative editing
* Online processing

---

# Success Criteria

The project is considered successful when it can:

* Detect supported platforms automatically
* Download videos reliably
* Export clips efficiently
* Use GPU acceleration when available
* Fall back gracefully to CPU encoding
* Remain stable during long-running operations
* Maintain a clean and understandable architecture

---

# Long-Term Vision

The application should remain easy to extend.

Adding support for a new platform should require implementing a new Platform Adapter without modifying the clipping, processing, or export pipelines.

The core application should remain independent from any individual video platform.

---

# Project Philosophy

This project is not a YouTube clipper.

It is a multi-platform desktop application for clipping online videos.

Platform support is a replaceable implementation detail.

The core application should always remain platform-agnostic, modular, and maintainable.
