# 04_DECISIONS.md

# Architecture Decision Records (ADR)

This document records permanent architectural decisions for the project.

These decisions should remain stable throughout development.

New decisions may be added over time.

Existing decisions should only change when absolutely necessary.

---

# ADR-001

## Platform-Agnostic Architecture

Status

Accepted

Decision

The application shall be platform-agnostic.

No core service may depend on a specific video platform.

Platform-specific logic must be isolated within the Platform Layer.

Reason

This allows new platforms to be added without changing business logic.

---

# ADR-002

## Platform Adapter Pattern

Status

Accepted

Decision

Every supported platform must implement the same adapter interface.

Examples:

* YouTube
* Kick
* Future platforms

Reason

A consistent interface simplifies maintenance and future expansion.

---

# ADR-003

## Local-First Processing

Status

Accepted

Decision

All media processing shall occur locally.

The application does not rely on cloud services.

Reason

Local processing provides better privacy, lower latency, and predictable performance.

---

# ADR-004

## Stream Copy First

Status

Accepted

Decision

The preferred export strategy is:

```text id="mk6g0x"
Stream Copy

↓

GPU Encoding

↓

CPU Encoding
```

Reason

Stream Copy provides the fastest export while preserving quality.

---

# ADR-005

## GPU as an Optimization

Status

Accepted

Decision

GPU acceleration is optional.

The application must continue functioning when GPU acceleration is unavailable.

Reason

Correctness is more important than performance.

---

# ADR-006

## Layered Architecture

Status

Accepted

Decision

The application follows a layered architecture.

```text id="jokx4j"
User Interface

↓

Application Layer

↓

Platform Layer

↓

Media Layer

↓

System Layer
```

Dependencies always point downward.

Reason

This keeps responsibilities clear and prevents architectural drift.

---

# ADR-007

## Generic Video Model

Status

Accepted

Decision

After platform resolution, every video shall be represented by the same generic data model.

Example:

* VideoMetadata
* StreamInfo
* ClipRequest

Core services must never distinguish between YouTube, Kick, or future platforms.

Reason

The application processes media, not platforms.

---

# ADR-008

## FFmpeg Isolation

Status

Accepted

Decision

Only the FFmpeg wrapper may execute FFmpeg.

No service may build or execute FFmpeg commands directly.

Reason

Centralized command generation improves maintainability and testing.

---

# ADR-009

## FFprobe Isolation

Status

Accepted

Decision

Media inspection shall be performed exclusively through the FFprobe wrapper.

Reason

Keeps metadata retrieval centralized and reusable.

---

# ADR-010

## yt-dlp Isolation

Status

Accepted

Decision

The application must never call yt-dlp directly from business services.

Only Platform Adapters may interact with yt-dlp.

Reason

yt-dlp is an implementation detail of supported platforms, not part of the application's business logic.

---

# ADR-011

## Thin Frontend

Status

Accepted

Decision

The frontend is responsible only for presentation and user interaction.

Business logic belongs exclusively to the Go backend.

Reason

Maintains a clear separation of concerns.

---

# ADR-012

## Explicit Dependencies

Status

Accepted

Decision

Use constructor injection.

Avoid dependency injection frameworks.

Reason

Explicit dependencies are easier to understand and debug.

---

# ADR-013

## Domain-Oriented Packages

Status

Accepted

Decision

Packages are organized by business domain rather than technical type.

Examples:

* clip
* download
* platform
* settings

Avoid generic packages such as:

* helpers
* managers
* common

Reason

Domain-oriented organization improves discoverability and long-term maintenance.

---

# ADR-014

## Local Media Pipeline

Status

Accepted

Decision

After download completes, all subsequent workflows operate exclusively on local media files.

Reason

Separating platform access from media processing keeps the architecture modular.

---

# ADR-015

## Stable Public Interfaces

Status

Accepted

Decision

Public interfaces should evolve slowly.

Internal implementations may change without affecting higher layers.

Reason

Stable contracts reduce refactoring and improve long-term maintainability.

---

# ADR-016

## Simplicity Over Flexibility

Status

Accepted

Decision

Prefer straightforward implementations over highly configurable abstractions.

Avoid speculative architecture.

Reason

The project is maintained by a solo developer with AI assistance.

Simplicity improves development speed and readability.

---

# ADR-017

## Minimal Dependencies

Status

Accepted

Decision

Introduce third-party libraries only when they provide substantial value.

Avoid dependencies for simple functionality.

Reason

A smaller dependency graph is easier to maintain and update.

---

# ADR-018

## Testing Strategy

Status

Accepted

Decision

Business logic should be covered primarily by unit tests.

Integration tests should focus on wrappers and external tools.

Reason

Fast, deterministic tests support frequent refactoring.

---

# ADR-019

## Documentation as Source of Truth

Status

Accepted

Decision

The documentation inside `.contexts/` defines the project's engineering standards.

Implementation should follow documentation.

If implementation and documentation diverge, update the documentation first before making architectural changes.

Reason

Shared documentation keeps AI and developer aligned over time.

---

# ADR Maintenance

When introducing a significant architectural decision:

1. Create a new ADR.
2. Assign the next sequential number.
3. Mark its status.
4. Explain the decision.
5. Explain the reasoning.

Do not rewrite existing ADRs unless the architecture intentionally changes.

---

# AI Guidelines

When generating code:

* Follow accepted ADRs.
* Do not introduce alternative architectures.
* Do not replace existing patterns without approval.
* Respect architectural boundaries.
* Treat ADRs as permanent project rules.

---

# Decisions Philosophy

Architecture should evolve intentionally, not accidentally.

Every major decision should be documented once and followed consistently.

Well-documented decisions reduce uncertainty, prevent unnecessary refactoring, and keep the project coherent as it grows.
