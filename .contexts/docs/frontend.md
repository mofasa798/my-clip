# frontend.md

# Frontend Engineering Guide

This document defines the engineering standards, responsibilities, and implementation guidelines for the React frontend.

The frontend is responsible for presenting the user interface and communicating with the backend.

Business logic must remain in the backend.

---

# Responsibilities

The frontend owns:

* User interaction
* Navigation
* Video preview
* Timeline editing
* Progress visualization
* Form validation
* Theme management
* Window state

The frontend should never implement business logic.

---

# Frontend Philosophy

The frontend presents workflows, not implementations.

Users interact with videos through a consistent interface regardless of where the media originates.

The frontend should remain unaware of source-specific behavior.

---

# Frontend Architecture

```text id="uv0z3r"
React Components

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
* Predictable
* Easy to maintain

Keep business rules in the backend.

---

# Backend-Driven UI

The backend is the source of truth.

The frontend should:

* Render backend models
* Display backend state
* Trigger backend actions
* Subscribe to backend events

The frontend should not decide:

* Which source is used
* Which encoder is selected
* Which hardware is available
* Which processing strategy is executed

Those decisions belong to the backend.

---

# Application Workflow

```text id="7h2v9f"
Paste Video URL

↓

Load Metadata

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

This workflow is identical for every supported source.

---

# Pages

Recommended pages:

```text id="2dcdf6"
Home

Editor

History

Settings
```

Each page should own one responsibility.

---

# Home Page

Responsibilities:

* Accept video URL
* Validate input
* Load metadata
* Display supported sources
* Navigate to the editor

The Home page should never detect the source.

Source detection belongs to the backend.

---

# Editor Page

Responsibilities:

* Display metadata
* Video preview
* Timeline editing
* Clip selection
* Export configuration
* Export progress

The Editor works exclusively with generic video models.

---

# History Page

Responsibilities:

* Display previous exports
* Display download history
* Search
* Filter
* Open output directory

History should remain read-only.

---

# Settings Page

Responsibilities:

* Output directory
* Preferred encoder
* Theme
* Export defaults
* Application preferences

Settings should be loaded from the backend.

---

# Component Organization

Example:

```text id="o7b54t"
components/

Button/

Card/

Dialog/

Input/

ProgressBar/

Timeline/

VideoPlayer/

MetadataCard/

ExportPanel/
```

Components should remain small and reusable.

---

# Hooks

Hooks manage UI behavior only.

Examples:

```text id="i8h8n7"
useTimeline()

useVideoPlayer()

useExport()

useHistory()

useSettings()
```

Avoid business logic inside hooks.

---

# Services

Services communicate with the backend.

Responsibilities:

* Wails bindings
* Event subscriptions
* Request mapping
* Response mapping

Services should never contain business rules.

---

# State Management

Global state should remain minimal.

Examples:

* Theme
* Active page
* Dialog state
* Export progress
* Playback state

Business state belongs to the backend.

---

# Data Models

The frontend should consume generic models.

Examples:

* VideoMetadata
* StreamInfo
* ClipRequest
* ExportOptions
* ExportProgress
* HistoryEntry

The frontend should never distinguish between YouTube, Kick, or future sources.

---

# Event Flow

The frontend reacts to backend events.

Examples:

```text id="3o4u3d"
metadata.loaded

download.started

download.progress

download.completed

export.started

export.progress

export.completed
```

Prefer event-driven communication over polling.

---

# Video Preview

The preview operates exclusively on local media.

Workflow:

```text id="7mhj2i"
Video URL

↓

Backend Download

↓

Local Media

↓

Video Preview
```

The frontend should never stream directly from an online source.

---

# Timeline

The timeline is a presentation component.

Responsibilities:

* Display duration
* Select clip range
* Move playhead
* Zoom
* Seek

Complex validation belongs to the backend.

---

# Form Validation

Frontend validation should be lightweight.

Validate:

* Required fields
* Empty input
* Basic formatting

Business validation belongs to the backend.

---

# Error Handling

Display user-friendly messages.

Examples:

Good:

```text id="v2n5g1"
Unable to load video metadata.
```

Avoid:

```text id="j0g8h2"
yt-dlp exited with status code 1.
```

Internal implementation details should never be exposed.

---

# Styling

Use:

* Tailwind CSS

Follow the project's Design System.

Avoid inline styles.

---

# Accessibility

The interface should support:

* Keyboard navigation
* Focus management
* Screen readers
* Visible focus indicators
* Clear interaction feedback

Accessibility should be considered from the beginning.

---

# Performance

The frontend should:

* Minimize unnecessary re-renders
* Lazy-load heavy components
* Keep component state local when possible
* Avoid excessive global state

Rendering should remain responsive during long-running backend operations.

---

# Testing

Frontend tests should focus on:

* Component rendering
* User interaction
* Hook behavior
* Event handling

Do not test backend business logic from the frontend.

---

# AI Guidelines

When generating frontend code:

* Keep components focused.
* Keep hooks lightweight.
* Reuse existing components.
* Consume backend models directly.
* Avoid duplicate state.
* Respect Backend-Driven UI.
* Never implement source-specific logic.

---

# Frontend Philosophy

The frontend exists to present a consistent editing experience.

It should not understand how media is downloaded, how clips are exported, or how hardware acceleration works.

Its responsibility is to display state, collect user input, and provide a responsive interface while the backend performs the application's core workflows.
