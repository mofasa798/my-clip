# 01_RULES.md

# AI Coding Rules

This document defines the mandatory coding standards and engineering principles for this project.

These rules apply to every code change unless explicitly overridden by the project owner.

When multiple implementation approaches are possible, always choose the simplest solution that satisfies the current requirements.

---

# Core Principles

Always prioritize:

1. Simplicity
2. Readability
3. Performance
4. Maintainability
5. Consistency

Do not optimize for hypothetical future requirements.

Avoid speculative programming.

---

# Documentation Rules

The documentation inside `.contexts/` is the **single source of truth**.

Never duplicate information across multiple documents.

If a document already owns a topic, reference it instead of copying it.

When documentation conflicts with the codebase, ask for clarification instead of making assumptions.

---

# Scope Rules

Only implement features that belong to the current development phase.

Never implement future roadmap items unless explicitly requested.

Never introduce functionality "just in case."

---

# Architecture Rules

Follow the architecture defined in:

* 02_ARCHITECTURE.md

Never bypass architectural boundaries.

Dependencies must always flow downward.

UI → Application → Infrastructure

Never reverse the dependency direction.

---

# Keep It Simple

Always choose:

* Small files
* Small functions
* Clear responsibilities
* Explicit behavior

Avoid unnecessary abstractions.

Avoid unnecessary indirection.

Prefer obvious code over clever code.

---

# Package Design

Packages should be cohesive.

Each package should have one clear responsibility.

Avoid packages that contain unrelated functionality.

Prefer many small focused packages over one large package.

---

# Interfaces

Do not create interfaces unless they provide real value.

Good reasons:

* Multiple implementations exist.
* Testing requires mocking.
* Clear abstraction already exists.

Bad reasons:

* Future-proofing.
* "Maybe we'll need it later."

Prefer concrete types whenever possible.

---

# Structs

Keep structs focused.

Avoid large "God Objects."

Each struct should own a single responsibility.

---

# Functions

Functions should:

* Do one thing.
* Have descriptive names.
* Return explicit errors.
* Remain easy to read.

Avoid deeply nested logic.

Extract helper functions when appropriate.

---

# Error Handling

Errors must never be ignored.

Always return meaningful errors.

Wrap external command failures with context.

Example:

* Which command failed.
* Exit code.
* stderr output.

Avoid panic except during unrecoverable startup failures.

---

# Logging

Log meaningful events only.

Examples:

* Startup
* Shutdown
* Download started
* Download completed
* Export started
* Export completed
* External process failed

Do not log excessive debug information unless debugging is enabled.

---

# Dependencies

Prefer the Go standard library.

Every third-party dependency must have a clear justification.

Never introduce a dependency merely for convenience.

Small utilities should be implemented directly when practical.

---

# External Commands

External binaries must never be called directly from the UI.

All command execution must be wrapped inside dedicated Go services.

Supported external tools:

* FFmpeg
* FFprobe
* yt-dlp

Every wrapper should:

* Validate arguments.
* Capture stdout.
* Capture stderr.
* Return structured errors.
* Support Context cancellation.

---

# Context Usage

Long-running operations must accept a Context.

Examples:

* Download
* Metadata retrieval
* Encoding
* Thumbnail extraction

Cancellation should terminate external processes whenever possible.

---

# Concurrency

Use Goroutines only when concurrency provides real benefits.

Avoid unnecessary parallelism.

Worker pools should remain lightweight.

Never introduce complex job schedulers for local processing.

---

# Backend Rules

Business logic belongs inside Go.

The backend owns:

* Video processing
* Downloads
* Settings
* File operations
* GPU detection
* FFmpeg wrappers
* yt-dlp wrappers

Never move business logic into React.

---

# Frontend Rules

The frontend is responsible for presentation only.

Responsibilities include:

* User interaction
* Layout
* Forms
* Timeline
* Progress display
* Notifications

Avoid placing business logic inside React components.

---

# React Rules

Use:

* Functional Components
* Hooks
* TypeScript

Avoid:

* Class Components
* Large monolithic components
* Deep prop drilling

Split components when responsibilities become unclear.

---

# State Management

Keep state as local as possible.

Prefer:

Component State

↓

Custom Hooks

↓

Global State (only when necessary)

Do not introduce state management libraries unless justified.

---

# Styling

Use Tailwind CSS.

Avoid inline styles unless necessary.

Keep styling consistent.

Reuse utility classes when appropriate.

---

# Performance Rules

Always prefer the fastest valid implementation.

Video processing priority:

1. Stream Copy
2. GPU Encoding
3. CPU Encoding

Never re-encode unless required.

Avoid unnecessary file copies.

Avoid unnecessary memory allocations.

---

# GPU Rules

Automatically detect GPU capabilities.

Never assume hardware acceleration is available.

Gracefully fall back to CPU encoding.

GPU usage must remain transparent to the user.

---

# File System Rules

Never hardcode user directories.

Always use configurable paths.

Ensure directories exist before writing files.

Avoid deleting user files automatically.

---

# Configuration Rules

Configuration must be stored locally.

Configuration should remain human-readable.

Missing configuration should automatically use sensible defaults.

---

# Naming Rules

Use descriptive names.

Avoid abbreviations unless universally understood.

Examples:

Good

* DownloadService
* ClipExporter
* GPUDetector

Bad

* DS
* Helper
* Manager
* Utils

Names should describe purpose rather than implementation.

---

# Comments

Prefer self-explanatory code.

Write comments only when they provide useful context.

Do not explain obvious code.

Keep comments up to date.

---

# Refactoring

Refactor only when it improves:

* Readability
* Maintainability
* Simplicity

Avoid unrelated refactoring during feature implementation.

Keep pull requests focused.

---

# Code Generation Rules

When generating new code:

* Match the existing project style.
* Reuse existing utilities.
* Avoid duplicate logic.
* Prefer composition.
* Keep implementations incremental.

Never rewrite large sections of code without explicit approval.

---

# Forbidden Patterns

Do not introduce the following unless explicitly requested:

* Repository Pattern
* Factory Pattern
* Service Locator
* CQRS
* Event Bus
* Mediator Pattern
* Generic Dependency Injection Framework
* Plugin Framework
* Enterprise Architecture patterns
* Premature optimization

The project intentionally remains simple.

---

# Testing Philosophy

Write code that is naturally testable.

Avoid designs created solely to satisfy testing frameworks.

Prefer deterministic behavior.

Business logic should be isolated from external commands whenever practical.

---

# AI Behavior

Before making changes:

1. Read the required project documentation.
2. Understand the current development phase.
3. Verify existing architecture.
4. Make the smallest necessary change.
5. Preserve coding consistency.

When requirements are ambiguous:

Stop.

Ask for clarification.

Do not guess.

---

# Success Criteria

Every implementation should be:

* Simple
* Readable
* Predictable
* Performant
* Maintainable
* Consistent with the rest of the project

If a proposed solution increases complexity without providing proportional value, it should be rejected.
