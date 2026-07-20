# 03_ROADMAP.md

# Development Roadmap

This document defines the long-term development roadmap for the project.

The roadmap is organized around architectural milestones rather than individual features.

Each milestone establishes a stable foundation for the next stage of development.

Detailed implementation tasks belong in `05_SPRINT.md`.

---

# Roadmap Principles

Development should prioritize:

* Strong architectural foundations
* Incremental delivery
* Stable abstractions
* Small iterative improvements
* Continuous validation

Avoid implementing future features before their prerequisites exist.

---

# Roadmap Overview

```text
Milestone 1
Project Foundation
        │
        ▼
Milestone 2
Source Layer
        │
        ▼
Milestone 3
Media Pipeline
        │
        ▼
Milestone 4
Desktop Experience
        │
        ▼
Milestone 5
Workflow Enhancement
        │
        ▼
Milestone 6
Stabilization & Release
```

Each milestone depends on the successful completion of the previous one.

---

# Milestone 1 — Project Foundation

## Objective

Establish a stable and maintainable project foundation.

## Deliverables

* Repository structure
* Wails v3 setup
* React + TypeScript setup
* Go backend setup
* Development tooling
* Build configuration
* Logging
* Configuration management
* Event infrastructure

## Success Criteria

* The application launches successfully.
* Frontend communicates with the backend.
* Development environment is fully operational.
* Project structure follows the documented architecture.

---

# Milestone 2 — Source Layer

## Objective

Implement a source-agnostic architecture for supported video sources.

## Deliverables

* Source interface
* Source resolver
* Source registry
* Generic metadata models
* Generic stream models
* YouTube adapter
* Kick adapter

## Success Criteria

* Supported sources are detected automatically.
* Metadata retrieval is consistent.
* Download workflow is source-independent.
* Adding a new source requires only a new adapter.

---

# Milestone 3 — Media Pipeline

## Objective

Build a reliable local media processing pipeline.

## Deliverables

* Download pipeline
* Local media storage
* FFprobe integration
* FFmpeg integration
* Hardware capability detection
* Stream Copy support
* GPU encoding
* CPU fallback
* Export progress reporting

## Success Criteria

* Media downloads reliably.
* Clips export correctly.
* Hardware acceleration is used when available.
* Software encoding remains a reliable fallback.

---

# Milestone 4 — Desktop Experience

## Objective

Create a polished desktop user experience.

## Deliverables

* Home page
* Editor page
* Timeline
* Video preview
* Settings
* Export panel
* Progress reporting
* History page
* Theme support

## Success Criteria

* Users can complete the full clipping workflow.
* The interface remains responsive during long-running operations.
* Timeline interaction is smooth.

---

# Milestone 5 — Workflow Enhancement

## Objective

Improve productivity and usability.

## Deliverables

* Export presets
* Download history
* Export history
* Keyboard shortcuts
* Improved validation
* Better error recovery
* Output management

## Success Criteria

* Common workflows require fewer manual steps.
* Repeated tasks become faster.
* Configuration becomes easier to manage.

---

# Milestone 6 — Stabilization & Release

## Objective

Prepare the application for long-term daily use.

## Deliverables

* Performance optimization
* Memory optimization
* Improved error handling
* Expanded test coverage
* Documentation review
* Packaging
* Installer
* Release build

## Success Criteria

* Stable long-running operation.
* Reliable export performance.
* Consistent user experience.
* Production-ready desktop application.

---

# Feature Priority

Development priority should always follow:

```text
Architecture

↓

Source Layer

↓

Media Pipeline

↓

Desktop Experience

↓

Workflow Improvements

↓

Quality & Release
```

Convenience features should never take priority over architectural stability.

---

# Out of Scope

The following are intentionally excluded from the current roadmap:

* Cloud synchronization
* User accounts
* Collaborative editing
* Remote media processing
* Video uploading
* Plugin system
* Mobile applications

These may be considered in future major versions.

---

# Future Expansion

Potential future milestones include:

* Additional Source Adapters
* Batch processing
* Hardware decoding
* AV1 encoding
* Subtitle support
* Waveform visualization
* Advanced timeline editing

Future work should preserve the established architecture.

---

# AI Guidelines

When planning or implementing work:

* Complete milestones sequentially.
* Build stable foundations before adding features.
* Avoid speculative implementations.
* Respect documented architectural boundaries.
* Keep the project source-agnostic.

If a requested feature depends on an unfinished milestone, complete the prerequisite first.

---

# Relationship to Sprint Planning

This roadmap defines **what** the project aims to achieve.

`05_SPRINT.md` defines **what is currently being built**.

The roadmap should change infrequently.

Sprint planning should evolve continuously throughout development.

---

# Roadmap Philosophy

The roadmap exists to guide long-term development, not daily implementation.

A stable architecture enables faster iteration, simpler maintenance, and easier expansion as new video sources and capabilities are added over time.
