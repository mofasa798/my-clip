# frontend.md

# Frontend Engineering Guide

This document defines the frontend implementation guidelines for the project.

The frontend is responsible only for presentation and user interaction.

Business logic belongs to the Go backend.

---

# Responsibilities

The frontend owns:

* User Interface
* User Interaction
* Navigation
* Forms
* Video Preview
* Timeline
* Progress Display
* Notifications
* Theme

The frontend should remain lightweight.

---

# Technology Stack

Framework

* React

Language

* TypeScript

Bundler

* Vite

Styling

* Tailwind CSS

Desktop Bridge

* Wails v3

---

# Frontend Directory

```text id="pjx3kn"
frontend/src/

components/
hooks/
layouts/
pages/
services/
stores/
types/
utils/
assets/
```

Each directory has one responsibility.

---

# Directory Responsibilities

## components

Reusable UI components.

Examples

* Button
* Input
* Timeline
* ProgressBar
* Modal
* VideoPreview

Components should remain small.

---

## pages

Application screens.

Examples

* Home
* Download
* Editor
* Settings
* History

Pages compose components.

Pages should not contain business logic.

---

## layouts

Shared page layouts.

Examples

* MainLayout
* SettingsLayout

Layouts organize pages.

---

## hooks

Reusable React hooks.

Examples

* useSettings
* useDownload
* useExport
* useProgress

Hooks encapsulate UI behavior.

---

## services

Frontend communication layer.

Responsibilities

* Call Wails bindings
* Transform responses
* Hide implementation details

Business logic does not belong here.

---

## stores

Global UI state only.

Examples

* Theme
* Sidebar state
* Current page

Avoid storing business state.

---

## types

Shared TypeScript types.

Examples

* VideoMetadata
* ExportOptions
* Settings

Keep types centralized.

---

## utils

Small helper functions.

No business logic.

---

## assets

Static assets.

Examples

* Icons
* Images
* Fonts

---

# Component Design

Each component should have one responsibility.

Good

```text id="od7ovx"
Timeline

VideoPreview

ProgressBar

SettingsForm

DownloadCard
```

Avoid large components with multiple unrelated responsibilities.

---

# Component Hierarchy

Preferred hierarchy

```text id="gpnexr"
Page

↓

Layout

↓

Components

↓

Small UI Elements
```

Avoid deeply nested component trees.

---

# Page Responsibilities

Pages should:

* Arrange components
* Manage local UI state
* Handle navigation

Pages should NOT:

* Execute FFmpeg
* Execute yt-dlp
* Perform business logic

---

# State Management

Keep state as local as possible.

Preferred order

```text id="98rymg"
Component State

↓

Custom Hooks

↓

Global Store
```

Avoid global state unless multiple pages require it.

---

# Wails Communication

All backend communication should be centralized.

Preferred flow

```text id="g3vfbz"
React Component

↓

Frontend Service

↓

Wails Binding

↓

Go Backend
```

Never call backend bindings directly from multiple components.

---

# Event Handling

Progress updates should use Wails Events.

Example

```text id="bbyjiv"
Backend

↓

Emit Event

↓

Frontend Service

↓

React State

↓

UI Update
```

---

# Forms

Forms should:

* Validate user input
* Display errors
* Disable actions while processing

Complex validation belongs to the backend.

---

# Video Preview

The preview component should:

* Display current frame
* Show playback position
* Synchronize with the timeline

The preview should not perform video processing.

---

# Timeline

Responsibilities

* Select clip range
* Seek video
* Display timestamps

Timeline interactions should remain smooth.

Business logic belongs to the backend.

---

# Progress Display

Long-running operations should provide:

* Percentage
* Current operation
* Estimated progress

The UI should remain responsive.

---

# Error Handling

Display user-friendly error messages.

Technical details should remain available in logs.

Example

Good

```text id="5nncc7"
Failed to export the selected clip.
```

Avoid

```text id="zw0gt4"
Error 127
```

---

# Loading States

Every asynchronous action should provide visual feedback.

Examples

* Spinner
* Progress Bar
* Disabled Button

Avoid leaving the user without feedback.

---

# Navigation

Keep navigation simple.

Expected pages

```text id="jwlqnu"
Home

Download

Editor

History

Settings
```

Avoid deeply nested navigation.

---

# Styling

Use Tailwind CSS.

Prefer utility classes.

Maintain consistent spacing.

Avoid inline styles unless necessary.

---

# Theme

Support:

* Light
* Dark

Theme should be controlled globally.

---

# Icons

Use one icon library consistently.

Avoid mixing icon sets.

---

# Notifications

Notify users about:

* Download completed
* Export completed
* Errors
* Dependency issues

Avoid excessive notifications.

---

# Accessibility

Provide:

* Keyboard navigation
* Focus states
* Accessible labels

Maintain reasonable color contrast.

---

# Performance

Optimize rendering.

Avoid unnecessary re-renders.

Memoization should only be used when beneficial.

---

# TypeScript

Avoid:

* any
* implicit types

Prefer explicit interfaces.

Keep types close to their domain.

---

# Services

Frontend services should:

* Call backend APIs
* Transform responses
* Normalize data

Services should NOT:

* Execute business logic
* Maintain application state

---

# Custom Hooks

Hooks should encapsulate reusable UI behavior.

Examples

```text id="ryjlwm"
useDownload()

useExport()

useTimeline()

useSettings()

useProgress()
```

Avoid hooks with multiple unrelated responsibilities.

---

# File Naming

Use lowercase with underscores.

Examples

```text id="l7sjbr"
video_preview.tsx
settings_page.tsx
download_card.tsx
progress_bar.tsx
```

---

# Component Naming

Use PascalCase.

Examples

```text id="z2edfd"
VideoPreview

ProgressBar

DownloadCard

SettingsForm
```

Component names should describe their purpose.

---

# AI Guidelines

When generating frontend code:

* Keep components focused.
* Reuse existing components.
* Keep pages lightweight.
* Avoid business logic.
* Use hooks for reusable behavior.
* Follow the project structure.
* Keep UI responsive.
* Preserve consistent styling.

---

# Frontend Philosophy

The frontend should present information, not own it.

Business rules belong to the backend.

The frontend exists to provide a clean, responsive, and intuitive desktop experience.

A simple interface with predictable behavior is preferred over unnecessary complexity.
