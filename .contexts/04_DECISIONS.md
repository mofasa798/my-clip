# 04_DECISIONS.md

# Architecture Decision Records (ADR)

This document records the permanent architectural decisions for the project.

Each ADR captures a significant design choice, its rationale, and the constraints it introduces.

Accepted ADRs should remain stable throughout the lifetime of the project.

---

# ADR-001

## Source-Agnostic Architecture

**Status**

Accepted

**Decision**

The application shall be source-agnostic.

The core application must never depend on a specific online video source.

**Rationale**

Supporting a new video source should require only implementing a new Source Adapter.

The application's business logic must remain unchanged.

---

# ADR-002

## Source Adapter Pattern

**Status**

Accepted

**Decision**

Every supported source must implement the same Source interface.

Examples:

* YouTube
* Kick
* Future sources

**Rationale**

A consistent interface keeps source implementations isolated and simplifies future expansion.

---

# ADR-003

## Local-First Processing

**Status**

Accepted

**Decision**

All media processing shall occur locally.

No cloud services are required.

**Rationale**

Local processing provides predictable performance, improved privacy, and complete offline media processing after download.

---

# ADR-004

## Capability-Driven Hardware Support

**Status**

Accepted

**Decision**

The application shall detect hardware capabilities at runtime.

Processing strategy must depend on available capabilities rather than specific hardware models.

**Rationale**

The application should remain compatible with future hardware upgrades without requiring architectural changes.

---

# ADR-005

## Processing Strategy

**Status**

Accepted

**Decision**

Preferred processing order:

```text
Stream Copy

↓

GPU Encoding

↓

CPU Encoding
```

**Rationale**

Always prefer the fastest reliable processing strategy while preserving compatibility.

---

# ADR-006

## Layered Architecture

**Status**

Accepted

**Decision**

The application shall follow this architecture:

```text
User Interface

↓

Application

↓

Domain

↓

Source
Media

↓

System
```

Dependencies always point downward.

**Rationale**

Layered responsibilities improve maintainability and reduce coupling.

---

# ADR-007

## Generic Media Model

**Status**

Accepted

**Decision**

Once a video has been resolved, it shall be represented using generic media models.

Examples:

* VideoMetadata
* StreamInfo
* ClipRequest
* ExportRequest

The Domain Layer must never distinguish between YouTube, Kick, or future sources.

**Rationale**

The application processes media rather than individual providers.

---

# ADR-008

## FFmpeg Isolation

**Status**

Accepted

**Decision**

Only the Media Layer may execute FFmpeg.

No business service may construct or execute FFmpeg commands directly.

**Rationale**

Centralized command generation improves maintainability and testing.

---

# ADR-009

## FFprobe Isolation

**Status**

Accepted

**Decision**

Media inspection shall occur exclusively through the FFprobe wrapper.

**Rationale**

Metadata retrieval remains centralized and reusable.

---

# ADR-010

## External Tool Isolation

**Status**

Accepted

**Decision**

External tools such as yt-dlp shall remain internal implementation details of the Source Layer.

The Domain Layer must never communicate with external tools directly.

**Rationale**

Business logic should remain independent from infrastructure.

---

# ADR-011

## Thin Frontend

**Status**

Accepted

**Decision**

The frontend shall remain responsible only for presentation and user interaction.

Business logic belongs exclusively in the backend.

**Rationale**

A thin frontend is easier to maintain and keeps architectural boundaries clear.

---

# ADR-012

## Explicit Dependencies

**Status**

Accepted

**Decision**

Dependencies shall be injected explicitly through constructors.

Dependency injection frameworks are not permitted.

**Rationale**

Explicit dependencies improve readability and simplify debugging.

---

# ADR-013

## Domain-Oriented Packages

**Status**

Accepted

**Decision**

Packages shall be organized by business domain instead of technical categories.

Examples:

* clip
* history
* settings
* source
* media

Avoid generic packages such as:

* helpers
* managers
* common

**Rationale**

Domain-oriented organization scales better as the project grows.

---

# ADR-014

## Stateless Services

**Status**

Accepted

**Decision**

Services should remain stateless whenever possible.

Persistent state belongs only to dedicated storage components.

**Rationale**

Stateless services simplify testing, concurrency, and maintenance.

---

# ADR-015

## Stable Public Interfaces

**Status**

Accepted

**Decision**

Public interfaces should evolve slowly.

Internal implementations may change freely without affecting higher layers.

**Rationale**

Stable contracts reduce unnecessary refactoring.

---

# ADR-016

## Simplicity Over Abstraction

**Status**

Accepted

**Decision**

Prefer simple, explicit implementations over highly configurable abstractions.

Avoid speculative architecture.

**Rationale**

The project is maintained by a solo developer with AI assistance.

Simplicity improves development speed and readability.

---

# ADR-017

## Minimal Dependencies

**Status**

Accepted

**Decision**

Introduce third-party libraries only when they provide significant value.

Avoid dependencies for functionality that can be implemented simply within the project.

**Rationale**

A smaller dependency graph is easier to maintain.

---

# ADR-018

## Testing Strategy

**Status**

Accepted

**Decision**

Business logic should primarily be covered by unit tests.

Integration tests should focus on external wrappers and infrastructure.

**Rationale**

Fast, deterministic tests encourage frequent refactoring.

---

# ADR-019

## Documentation as the Source of Truth

**Status**

Accepted

**Decision**

Documentation inside `.contexts/` defines the engineering standards of the project.

Implementation should follow documentation.

Architectural documentation should be updated before significant architectural changes are implemented.

**Rationale**

Shared documentation keeps developers and AI aligned over time.

---

# ADR-020

## Desktop-First Distribution

**Status**

Accepted

**Decision**

The application shall be distributed as a standalone desktop application.

It must not require:

* Servers
* Databases
* Containers
* Cloud infrastructure

**Rationale**

The project targets local desktop usage and should remain simple to install and operate.

---

# ADR Maintenance

When introducing a new architectural decision:

1. Create a new ADR.
2. Assign the next sequential identifier.
3. Record its status.
4. Describe the decision.
5. Explain the rationale.

Existing ADRs should only change when the architecture intentionally evolves.

---

# AI Guidelines

When generating code:

* Follow accepted ADRs.
* Respect architectural boundaries.
* Preserve dependency direction.
* Avoid introducing alternative architectures.
* Do not replace established patterns without explicit approval.

When in doubt, prioritize existing ADRs over implementation convenience.

---

# Decision Philosophy

Architecture should evolve intentionally, not accidentally.

Every major decision should be documented once, understood by both developers and AI, and followed consistently throughout the project.

Well-defined decisions reduce uncertainty, prevent architectural drift, and make future development predictable.
