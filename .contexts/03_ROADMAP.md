# 03_ROADMAP.md

# Development Roadmap

This document defines the long-term development roadmap for the project.

The roadmap is organized around architectural milestones rather than individual features.

Each phase should produce a stable foundation for the next phase.

Implementation should always follow the roadmap unless explicitly overridden.

---

# Roadmap Principles

Development should prioritize:

* Strong foundations
* Incremental delivery
* Stable architecture
* Small iterations
* Continuous validation

Avoid implementing future features before their prerequisites exist.

---

# Phase Overview

```text id="w6u1ja"
Phase 1

Project Foundation

↓

Phase 2

Platform Layer

↓

Phase 3

Media Pipeline

↓

Phase 4

User Interface

↓

Phase 5

Application Features

↓

Phase 6

Stabilization & Release
```

Each phase depends on the previous one.

---

# Phase 1 — Project Foundation

Objective

Create a maintainable project foundation.

Deliverables

* Repository structure
* Wails v3 setup
* React + TypeScript setup
* Go backend setup
* Build configuration
* Application startup
* Dependency verification
* Logging
* Settings management
* Event infrastructure

Success Criteria

* Application launches successfully.
* Frontend communicates with backend.
* Dependencies are detected correctly.
* Development environment is stable.

---

# Phase 2 — Platform Layer

Objective

Create a platform-independent abstraction for supported video providers.

Deliverables

* Platform interface
* Platform resolver
* URL detection
* URL validation
* Video metadata model
* Stream information model
* Adapter registration
* YouTube adapter
* Kick adapter

Success Criteria

* Supported platforms are detected automatically.
* Metadata can be retrieved consistently.
* Platform implementations remain isolated.
* Core services do not depend on platform-specific logic.

---

# Phase 3 — Media Pipeline

Objective

Build the complete local media processing pipeline.

Deliverables

* Download workflow
* Local media storage
* FFprobe wrapper
* FFmpeg wrapper
* GPU detection
* Stream Copy export
* GPU encoding
* CPU fallback
* Progress reporting
* Export pipeline

Success Criteria

* Videos download successfully.
* Clips export correctly.
* GPU acceleration works when available.
* CPU fallback behaves reliably.

---

# Phase 4 — User Interface

Objective

Create a complete desktop experience.

Deliverables

* Home page
* Metadata display
* Video preview
* Timeline editor
* Export panel
* Progress display
* Settings page
* History page
* Theme support

Success Criteria

* Users can complete the primary workflow.
* Timeline interaction is responsive.
* Progress updates are reliable.
* UI remains responsive during long operations.

---

# Phase 5 — Application Features

Objective

Improve usability and workflow efficiency.

Deliverables

* Export presets
* Recent projects
* Download history
* Export history
* Keyboard shortcuts
* Output management
* Error recovery
* Improved validation

Success Criteria

* Daily workflows become faster.
* Configuration is easier to manage.
* Frequently used actions require fewer steps.

---

# Phase 6 — Stabilization & Release

Objective

Prepare the application for long-term use.

Deliverables

* Performance optimization
* Memory optimization
* Error handling improvements
* Test coverage improvements
* Documentation review
* Packaging
* Installer
* Release build

Success Criteria

* Stable long-running operation.
* Reliable export performance.
* Consistent user experience.
* Ready for production use.

---

# Feature Priority

Priority order:

```text id="4qnn7t"
Core Architecture

↓

Platform Support

↓

Media Processing

↓

User Experience

↓

Convenience Features
```

Never prioritize convenience features over core functionality.

---

# Out of Scope

The following are intentionally excluded from the current roadmap:

* Cloud synchronization
* User accounts
* Online editing
* Collaborative features
* Video uploading
* Plugin system
* Mobile application

These may be considered only after the core application is mature.

---

# Future Expansion

Possible future phases:

* Additional platform adapters
* AV1 encoding
* Hardware decoding
* Batch processing
* Advanced timeline editing
* Waveform visualization
* Subtitle editing

Future work should preserve the existing architecture.

---

# AI Guidelines

When planning or implementing work:

* Complete one phase before starting the next.
* Avoid skipping architectural milestones.
* Respect dependency order.
* Build reusable foundations.
* Do not implement speculative features.

If a requested feature depends on unfinished work, implement the prerequisite first.

---

# Roadmap Philosophy

The roadmap prioritizes architecture over features.

A stable foundation enables faster feature development and easier long-term maintenance.

Every completed phase should leave the project in a usable, testable, and extensible state.
