# ui.md

# User Interface & User Experience Guide

This document defines the UI and UX principles for the application.

The goal is to provide a clean, fast, and distraction-free desktop experience focused on clipping videos efficiently.

The interface should always prioritize usability over visual complexity.

---

# Design Principles

The interface should be:

* Simple
* Fast
* Responsive
* Predictable
* Consistent

Avoid unnecessary animations or decorative elements.

---

# Design Philosophy

The application is a productivity tool.

The UI should feel:

* Lightweight
* Professional
* Minimal
* Efficient

Users should complete common tasks with as few interactions as possible.

---

# User Workflow

The primary workflow is:

```text id="7bgmrd"
Paste URL

↓

Load Metadata

↓

Download (if needed)

↓

Select Clip

↓

Choose Export Options

↓

Export

↓

Open Output Folder
```

Every screen should support this workflow.

---

# Navigation

Keep navigation shallow.

Recommended pages:

```text id="qh7x8v"
Home

↓

Editor

↓

History

↓

Settings
```

Avoid multi-level navigation.

---

# Home Page

Responsibilities:

* Paste video URL
* Detect supported platform
* Display metadata
* Display thumbnail
* Select quality
* Start download

The home page is the application's entry point.

---

# Editor Page

Responsibilities:

* Video preview
* Timeline
* Clip range selection
* Export settings
* Export progress

The editor should focus entirely on clip creation.

---

# History Page

Responsibilities:

* Recent downloads
* Recent exports
* Quick access to output files

History should remain lightweight.

---

# Settings Page

Responsibilities:

* Output directory
* Export defaults
* Theme
* Encoder preference
* Dependency status

Advanced settings should remain hidden unless necessary.

---

# Layout

Recommended layout:

```text id="2jcp1m"
Header

↓

Main Content

↓

Status Bar
```

Avoid sidebars unless future requirements justify them.

---

# Header

Contains:

* Application title
* Current page
* Theme switch (optional)
* Settings shortcut

Keep the header compact.

---

# Status Bar

Display:

* Current task
* Export progress
* Application status

Avoid excessive information.

---

# Video Preview

The preview should provide:

* Playback controls
* Current timestamp
* Frame display

Do not overload the preview with controls.

---

# Timeline

The timeline should:

* Support dragging
* Display clip markers
* Allow precise seeking
* Show timestamps

Interaction should remain smooth.

---

# Export Panel

Display:

* Output filename
* Output directory
* Encoder
* Estimated duration
* Export button

Keep export options grouped together.

---

# Progress Display

Every long-running task should show:

* Progress percentage
* Current operation
* Cancel button

Never leave users wondering if work is still in progress.

---

# Error Messages

Errors should be:

* Clear
* Friendly
* Actionable

Example:

```text id="v92m8j"
FFmpeg could not be found.

Please configure the FFmpeg executable in Settings.
```

Avoid technical jargon unless helpful.

---

# Notifications

Notify users when:

* Download completes
* Export completes
* Errors occur
* Dependencies are missing

Notifications should be brief.

---

# Forms

Forms should:

* Validate input immediately
* Display inline errors
* Disable invalid actions

Do not overwhelm users with validation messages.

---

# Buttons

Primary actions:

* Paste URL
* Download
* Export

Secondary actions:

* Cancel
* Open Folder
* Retry

Use clear action labels.

Avoid vague labels like:

* OK
* Run
* Execute

---

# Icons

Use one icon library consistently.

Icons should support labels, not replace them.

---

# Colors

Use colors to communicate meaning.

Examples:

* Green — Success
* Yellow — Warning
* Red — Error
* Blue — Primary Action

Do not rely on color alone.

---

# Typography

Prefer:

* Clear hierarchy
* Consistent spacing
* Readable font sizes

Avoid excessive font variations.

---

# Spacing

Maintain consistent spacing throughout the application.

Prefer an 8px spacing system.

Example:

```text id="mxdtxw"
4

8

16

24

32
```

---

# Accessibility

Support:

* Keyboard navigation
* Focus indicators
* Screen reader labels where practical

Maintain sufficient contrast.

---

# Responsiveness

Although the application targets desktop use, layouts should adapt gracefully to different window sizes.

Avoid fixed-size layouts whenever possible.

---

# Empty States

Provide helpful empty states.

Examples:

* No history available
* No downloads yet
* No recent exports

Explain the next action.

---

# Loading States

Every asynchronous action should provide feedback.

Examples:

* Skeleton placeholders
* Loading indicators
* Disabled controls

Avoid blocking the interface unnecessarily.

---

# Confirmation Dialogs

Ask for confirmation only when actions are destructive.

Examples:

* Delete history
* Remove exported files

Do not interrupt normal workflows with unnecessary confirmations.

---

# Theme

Support:

* Light
* Dark

Themes should affect the entire application consistently.

---

# Performance

The UI should remain responsive during:

* Downloading
* Metadata retrieval
* Exporting
* GPU encoding

Long-running tasks must never freeze the interface.

---

# Future UI Features

Potential future enhancements:

* Keyboard shortcuts
* Customizable layout
* Export presets
* Recent project reopening

These features should only be implemented when scheduled in the roadmap.

---

# AI Guidelines

When generating UI code:

* Keep layouts simple.
* Reuse existing components.
* Avoid visual clutter.
* Follow consistent spacing.
* Use descriptive labels.
* Preserve accessibility.
* Keep user workflows intuitive.

Every screen should support the primary workflow with minimal friction.

---

# UI Philosophy

The interface exists to help users complete tasks quickly and confidently.

A clean, predictable interface is more valuable than a feature-rich but cluttered design.

Every UI element should have a clear purpose.

If an element does not improve the user experience, it should not exist.
