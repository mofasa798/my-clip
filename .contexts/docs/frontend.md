# frontend.md

# Frontend Engineering Guide

This document defines the frontend architecture, responsibilities, and engineering standards.

The frontend is responsible for presenting the user interface and communicating with the backend.

It should remain lightweight, predictable, and free of business logic.

---

# Frontend Responsibilities

The frontend owns:

* User interaction
* Navigation
* State presentation
* Timeline interaction
* Video preview
* Form validation
* Progress visualization
* Theme management

The frontend must never implement business logic.

---

# Frontend Architecture

```text
React UI

↓

Components

↓

Hooks

↓

Services

↓

Wails Bindings

↓

Go Backend
```

Communication always flows downward.

---

# Core Principles

The frontend should be:

* Thin
* Reactive
* Component-based
* Accessible
* Easy to maintain

Business rules belong in the backend.

---

# Application Workflow

The UI follows a simple workflow.

```text
Paste Video URL

↓

Load Video Metadata

↓

Preview Video

↓

Select Clip Range

↓

Configure Export

↓

Export Clip

↓

View History
```

The workflow is identical for every supported video source.

---

# Page Structure

Recommended pages:

```text
Home

Editor

History

Settings
```

Each page owns one responsibility.

---

# Home Page

Responsibilities:

* Accept video URL
* Display supported providers
* Validate input
* Request metadata
* Navigate to the editor

The Home page should not detect providers.

Provider detection belongs to the backend.

---

# Editor Page

Responsibilities:

* Display video metadata
* Preview video
* Timeline editing
* Export configuration
* Progress monitoring

The Editor operates only on generic video metadata.

---

# History Page

Responsibilities:

* Display previous exports
* Display download history
* Search
* Filter
* Open export directory

History should remain read-only.

---

# Settings Page

Responsibilities:

* Output directory
* Preferred encoder
* Theme
* GPU preferences
* Application settings

Settings should be loaded from the backend.

---

# Component Organization

Example:

```text
components/

Button/
Card/
Dialog/
Input/
ProgressBar/

MetadataCard/

Timeline/

VideoPlayer/

ExportPanel/
```

Components should remain small and reusable.

---

# Component Rules

Each component should:

* Have one responsibility
* Receive data through props
* Avoid hidden side effects
* Remain reusable

Avoid overly large components.

---

# Hooks

Hooks manage UI behavior.

Examples:

```text
useExport()

useTimeline()

useVideoPlayer()

useSettings()
```

Hooks should not contain business logic.

---

# Services

Services communicate with the backend.

Responsibilities:

* Wails bindings
* Request mapping
* Response mapping
* Error propagation

Services should never implement business rules.

---

# State Management

Global state should be minimal.

Examples:

* Theme
* Active page
* Current export progress
* Dialog state

Business state belongs to the backend.

---

# Data Models

The frontend should consume generic models.

Examples:

```text
VideoMetadata

StreamInfo

ClipRequest

ExportOptions

ExportProgress
```

The frontend should not distinguish between YouTube, Kick, or future providers.

---

# Event Flow

Backend events:

```text
download.started

download.progress

download.completed

export.started

export.progress

export.completed
```

The frontend subscribes to events.

Avoid polling.

---

# Video Preview

The preview is based on local media.

The frontend should never stream directly from an online provider.

Workflow:

```text
Video URL

↓

Backend Download

↓

Local Media

↓

Video Preview
```

---

# Timeline

The timeline is a presentation component.

Responsibilities:

* Display duration
* Select start/end points
* Display playhead
* Zoom
* Seek

Timeline calculations remain simple.

Complex validation belongs to the backend.

---

# Form Validation

Frontend validation should cover:

* Required fields
* Empty input
* Basic formatting

Business validation belongs to the backend.

---

# Error Handling

Display user-friendly messages.

Avoid exposing internal implementation details.

Example:

Good:

```text
Unable to retrieve video metadata.
```

Avoid:

```text
yt-dlp exited with code 1
```

Implementation details should remain hidden.

---

# Performance

Prioritize:

* Fast rendering
* Minimal re-renders
* Lazy loading where appropriate
* Lightweight components

Avoid unnecessary state.

---

# Accessibility

The UI should support:

* Keyboard navigation
* Focus management
* Screen reader compatibility
* Clear visual feedback

Accessibility should be considered from the beginning.

---

# Styling

Use Tailwind CSS.

Design should follow the Design System documentation.

Avoid inline styles.

---

# Testing

Frontend tests should verify:

* Rendering
* User interaction
* Component behavior
* Hook behavior

Do not test implementation details.

---

# AI Guidelines

When generating frontend code:

* Keep components focused.
* Keep hooks lightweight.
* Place business logic in the backend.
* Reuse existing components.
* Consume generic backend models.
* Never assume a specific video provider.

---

# Frontend Philosophy

The frontend presents workflows, not Source implementations.

Users interact with videos through a consistent interface, regardless of where the media originated.

A well-designed frontend is unaware of provider-specific behavior and focuses entirely on delivering a smooth editing experience.
