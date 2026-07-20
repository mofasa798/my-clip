# 00_PROJECT.md

# Project Overview

## Project Name

Multi-Platform Video Clipper

---

# Vision

Build a fast, reliable, and intuitive desktop application for downloading, previewing, clipping, and exporting videos from multiple online video sources.

The application should provide a consistent user experience regardless of where the video originates.

Source-specific implementations must remain isolated from the rest of the application.

---

# Mission

Develop a local-first desktop application capable of:

* Detecting supported video sources
* Retrieving video metadata
* Downloading media
* Previewing videos
* Creating clips
* Exporting clips efficiently
* Utilizing hardware acceleration whenever available

The application should prioritize performance, maintainability, and simplicity.

---

# Core Principles

The project follows these principles:

* Multi-source by design
* Source-agnostic architecture
* Local-first processing
* Capability-driven hardware support
* Performance-oriented implementation
* Simple and maintainable codebase
* Minimal dependencies
* Consistent user experience

---

# Supported Sources

Current supported sources:

* YouTube
* Kick

Future sources should be added by implementing a new Source Adapter without requiring changes to the core application.

---

# Primary Workflow

```text
Paste Video URL
        │
        ▼
Detect Source
        │
        ▼
Load Metadata
        │
        ▼
Download Media
        │
        ▼
Preview Video
        │
        ▼
Select Clip Range
        │
        ▼
Export Clip
        │
        ▼
Open Output Folder
```

Every feature should support this workflow.

---

# Source Independence

The application should never assume a specific video source.

After a URL has been resolved by the Source Layer, every video should be treated as a generic media source.

Clipping, previewing, processing, and exporting must remain completely independent of the originating source.

---

# Target Environment

## Execution Model

* Local desktop application
* Offline-first media processing
* No cloud services required

## Operating Systems

Primary target:

* Windows

Future targets:

* Linux
* macOS

## Desktop Framework

* Wails v3

---

# Hardware Philosophy

The application should adapt to the user's hardware capabilities instead of requiring specific hardware models.

Whenever possible, processing should automatically select the most efficient available method.

Supported hardware acceleration includes:

* NVIDIA NVENC
* AMD AMF
* Intel Quick Sync Video (QSV)

If no hardware acceleration is available, the application must automatically fall back to software encoding.

---

# Processing Strategy

Preferred processing order:

```text
Stream Copy
        │
        ▼
GPU Encoding
        │
        ▼
CPU Encoding
```

The application should always choose the fastest reliable strategy.

---

# Performance Goals

The application should:

* Launch quickly
* Minimize memory usage
* Keep the interface responsive
* Detect hardware capabilities automatically
* Avoid unnecessary processing
* Handle large media files efficiently
* Maintain stable long-running operations

---

# Architecture Philosophy

The application is divided into independent layers.

```text
User Interface
        │
        ▼
Application Layer
        │
        ▼
Source Layer
        │
        ▼
Media Layer
        │
        ▼
System Layer
```

Each layer has one clear responsibility.

Dependencies always point downward.

---

# User Experience Goals

The application should provide:

* Fast metadata retrieval
* Responsive timeline editing
* Smooth local video preview
* Reliable downloads
* Predictable exports
* Clear progress reporting
* Friendly error messages

Users should be able to create clips with minimal configuration.

---

# Scope

## Included

* Multiple video sources
* Metadata retrieval
* Media downloading
* Local video preview
* Clip creation
* GPU-accelerated export
* Export history
* User settings

## Excluded

* Cloud synchronization
* User accounts
* Online editing
* Video uploading
* Collaborative workflows
* Remote media processing

---

# Success Criteria

The project is considered successful when it can:

* Detect supported sources automatically
* Retrieve metadata consistently
* Download media reliably
* Export clips efficiently
* Utilize hardware acceleration when available
* Fall back gracefully to software encoding
* Maintain a clean and extensible architecture

---

# Long-Term Vision

The application should remain easy to extend.

Supporting a new video source should require only implementing a new Source Adapter and registering it with the Source Resolver.

No modifications should be required to:

* Media processing
* Clip creation
* Export pipeline
* User interface

---

# Project Philosophy

This project is not a YouTube clipper.

It is a multi-platform desktop application for clipping videos from supported online sources.

Video sources are replaceable implementations.

The core application should always remain source-agnostic, modular, maintainable, and easy to extend.
