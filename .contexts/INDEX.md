# INDEX.md

# AI Context Index

This directory contains the engineering documentation that defines the project's architecture, standards, and development workflow.

These documents are the primary source of truth for both developers and AI agents.

When generating code, always follow these documents before introducing new patterns or abstractions.

---

# Reading Order

Read the documents in the following order:

```text
00_PROJECT.md
        │
        ▼
01_RULES.md
        │
        ▼
02_ARCHITECTURE.md
        │
        ▼
04_DECISIONS.md
        │
        ▼
08_PROJECT_STRUCTURE.md
        │
        ▼
Remaining documents as needed
```

Project vision and architecture always take precedence over implementation details.

---

# Core Documents

These documents define the project's foundations.

| Document                  | Purpose                                        |
| ------------------------- | ---------------------------------------------- |
| `00_PROJECT.md`           | Project vision, scope, goals, and philosophy   |
| `01_RULES.md`             | Development rules and engineering standards    |
| `02_ARCHITECTURE.md`      | System architecture and layer responsibilities |
| `03_ROADMAP.md`           | Long-term architectural roadmap                |
| `04_DECISIONS.md`         | Architecture Decision Records (ADR)            |
| `05_SPRINT.md`            | Current development priorities                 |
| `06_PROMPT.md`            | AI prompting guidelines                        |
| `07_CODE_STYLE.md`        | Coding standards and naming conventions        |
| `08_PROJECT_STRUCTURE.md` | Directory structure and package organization   |
| `09_TESTING.md`           | Testing strategy and standards                 |

---

# Engineering Guides

These documents explain subsystem implementations.

| Document           | Purpose                           |
| ------------------ | --------------------------------- |
| `backend.md`       | Backend engineering guide         |
| `frontend.md`      | Frontend engineering guide        |
| `sources.md`       | Source Layer engineering guide    |
| `ffmpeg.md`        | FFmpeg integration                |
| `gpu.md`           | GPU acceleration                  |
| `ui.md`            | User interface guidelines         |
| `design_system.md` | Visual language and design system |

Additional engineering guides may be added as the project evolves.

---

# Document Priority

When multiple documents discuss the same topic, follow this priority:

```text
00_PROJECT

↓

02_ARCHITECTURE

↓

04_DECISIONS

↓

08_PROJECT_STRUCTURE

↓

Engineering Guides

↓

Sprint Planning
```

Higher-priority documents override lower-priority documents.

---

# Document Responsibilities

## Vision

Defines:

* Why the project exists
* Scope
* Goals
* Philosophy

---

## Architecture

Defines:

* Layer boundaries
* Responsibilities
* Dependencies

---

## Decisions

Defines:

* Permanent architectural decisions
* Accepted design patterns
* Long-term constraints

---

## Structure

Defines:

* Directory layout
* Package organization
* File locations

---

## Engineering Guides

Define:

* Implementation standards
* Best practices
* Technology-specific guidance

---

## Sprint

Defines:

* Current implementation priorities
* Active tasks
* Short-term goals

Sprint planning should evolve frequently.

---

# AI Instructions

Before generating code:

1. Read the relevant engineering guides.
2. Follow the documented architecture.
3. Reuse existing models and services.
4. Respect architectural boundaries.
5. Avoid introducing alternative patterns without approval.

When uncertain, prefer existing project conventions.

---

# Documentation Philosophy

The documentation exists to keep developers and AI aligned.

Architectural consistency is more valuable than introducing new abstractions.

When the documentation and implementation differ, update the documentation first if the architecture is intentionally changing.

The goal is a predictable, maintainable, and extensible codebase that remains easy to understand over time.
