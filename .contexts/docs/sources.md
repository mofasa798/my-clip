# sources.md

# Source Layer Engineering Guide

This document defines the architecture, responsibilities, and engineering standards for the Source Layer.

The Source Layer is responsible for interacting with supported online video sources.

It isolates source-specific implementations from the rest of the application.

The core application should never depend on any individual video source.

---

# Purpose

The Source Layer exists to:

* Detect supported sources
* Validate video URLs
* Retrieve video metadata
* Discover available streams
* Download media
* Convert source-specific data into generic application models

Everything outside this layer should work with generic video objects.

---

# Design Principles

The Source Layer should be:

* Source-agnostic
* Extensible
* Predictable
* Stateless
* Easy to test

Every source should behave consistently from the perspective of the application.

---

# Supported Sources

Initial sources:

* YouTube
* Kick

Future sources may include:

* Twitch
* Vimeo
* Dailymotion

Adding a new source should require minimal changes outside this layer.

---

# Source Architecture

```text
Video URL

↓

Source Resolver

↓

Source Adapter

↓

Generic Video Metadata

↓

Download Service
```

Source-specific logic ends at the Source Adapter.

---

# Responsibilities

The Source Layer owns:

* Source detection
* URL validation
* Metadata retrieval
* Stream discovery
* Download orchestration
* Source capability reporting

It does not own:

* Video clipping
* Video preview
* Export
* GPU selection
* Timeline editing

Those belong to other layers.

---

# Source Resolver

The Source Resolver determines which adapter should handle a URL.

Responsibilities:

* Detect source
* Match supported URLs
* Select adapter

Example:

```text
https://youtube.com/...

↓

YouTube Adapter
```

```text
https://kick.com/...

↓

Kick Adapter
```

The resolver should not perform downloads.

---

# Source Adapter

Every source implements the same interface.

Responsibilities:

* Validate URL
* Retrieve metadata
* Retrieve stream information
* Download media

The adapter hides all implementation details.

---

# Source Interface

Recommended interface:

```go
type SourceAdapter interface {
    Name() string

    Match(url string) bool

    Metadata(ctx context.Context, url string) (*VideoMetadata, error)

    Download(ctx context.Context, request DownloadRequest) (*DownloadResult, error)
}
```

All adapters should implement this interface.

---

# Generic Models

After source resolution, all data should be represented using generic models.

Example:

```go
type VideoMetadata struct {
    Source      string
    Title       string
    Author      string
    Duration    time.Duration
    Thumbnail   string
    URL         string
    Streams     []StreamInfo
}
```

Business services should consume only these models.

---

# Source Capabilities

Every source may expose different capabilities.

Example:

```text
Supports Metadata

Supports Download

Supports Live Stream

Supports Multiple Streams

Supports Chapters
```

Capabilities should be reported through the adapter.

Business logic should not hardcode assumptions.

---

# URL Validation

Validation should occur before any network operation.

Validate:

* URL format
* Supported source
* Required identifiers

Return descriptive errors.

---

# Metadata Retrieval

Metadata should include:

* Title
* Author
* Duration
* Thumbnail
* Source
* Available streams
* Live status
* Original URL

Metadata retrieval should be lightweight.

---

# Stream Discovery

Every adapter should return available streams.

Example:

* Video streams
* Audio streams
* Combined streams

The application selects the appropriate stream later.

---

# Download Workflow

```text
Video URL

↓

Source Resolver

↓

Source Adapter

↓

Download

↓

Local Media File
```

The Download Service should never know which source is used.

---

# Error Handling

Errors should be meaningful.

Examples:

Good:

* Unsupported source
* Invalid video URL
* Video unavailable
* Download failed

Avoid exposing internal implementation details.

---

# External Dependencies

Source adapters may use external tools.

Examples:

* yt-dlp

These dependencies must remain internal to the adapter.

No other layer should invoke them directly.

---

# Source Registry

The application should maintain a registry of available adapters.

Example:

```text
Registry

↓

YouTube

Kick
```

The resolver selects an adapter from the registry.

Avoid hardcoding source selection throughout the application.

---

# Directory Structure

Recommended structure:

```text
source/

resolver/

registry/

adapters/

    youtube/

    kick/
```

Each source implementation should remain self-contained.

---

# Adding a New Source

Steps:

1. Implement the Source Adapter interface.
2. Register the adapter.
3. Verify URL detection.
4. Test metadata retrieval.
5. Test download workflow.

No changes should be required in:

* Clip Service
* Media Layer
* Export Pipeline
* Frontend

---

# Testing

Unit tests should verify:

* URL matching
* Metadata mapping
* Stream mapping
* Error handling

Integration tests should verify:

* Download
* Metadata retrieval
* External tool integration

---

# Performance

The Source Layer should:

* Minimize network requests
* Cache temporary metadata where appropriate
* Avoid duplicate downloads
* Return metadata quickly

Long-running operations should report progress.

---

# AI Guidelines

When generating code:

* Keep source implementations isolated.
* Follow the Source Adapter interface.
* Convert source-specific data into generic models.
* Never leak implementation details outside this layer.
* Avoid source-specific logic in business services.
* Prefer composition over inheritance.

---

# Source Layer Philosophy

The application does not process YouTube videos or Kick videos.

It processes videos obtained from supported sources.

Once a video has been resolved and downloaded, its origin becomes an implementation detail.

The rest of the application should operate exclusively on generic media models, ensuring a consistent architecture that remains easy to extend as new sources are added.
