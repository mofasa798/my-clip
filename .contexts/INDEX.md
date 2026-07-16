# INDEX.md

# AI Documentation Index

Welcome to the project documentation.

This directory contains the complete engineering documentation for the project and serves as the **single source of truth** for all AI coding assistants and contributors.

Before making any code changes, read this document and follow the recommended reading order.

---

# Purpose

The documentation is designed to:

* Keep development consistent.
* Prevent over-engineering.
* Maintain a clean architecture.
* Preserve coding style.
* Ensure AI follows the same engineering principles throughout the project.
* Provide a predictable development workflow.

Each document owns a single responsibility.

Avoid duplicating information across multiple documents.

---

# Documentation Structure

```text
.contexts/

INDEX.md

00_PROJECT.md
01_RULES.md
02_ARCHITECTURE.md
03_ROADMAP.md
04_DECISIONS.md
05_SPRINT.md
06_PROMPT.md
07_CODE_STYLE.md
08_PROJECT_STRUCTURE.md

docs/
```

---

# Reading Order

Always read the following documents in order.

## 1. Project Overview

**00_PROJECT.md**

Purpose

Understand:

* Vision
* Objectives
* Scope
* Technology Stack
* Hardware Target
* Project Philosophy

This document explains **what** is being built.

---

## 2. Engineering Rules

**01_RULES.md**

Purpose

Understand:

* Coding Rules
* Engineering Principles
* AI Constraints
* Error Handling
* Dependency Rules

This document explains **how code should be written**.

---

## 3. Architecture

**02_ARCHITECTURE.md**

Purpose

Understand:

* System Layers
* Module Responsibilities
* Dependency Flow
* Processing Pipeline

This document explains **how the application is designed**.

---

## 4. Roadmap

**03_ROADMAP.md**

Purpose

Understand:

* Development Phases
* Milestones
* Deliverables
* Exit Criteria

Only implement work belonging to the current phase.

---

## 5. Architecture Decisions

**04_DECISIONS.md**

Purpose

Understand:

* Accepted ADRs
* Technology Decisions
* Architecture Constraints

Accepted decisions should not be changed unless explicitly requested.

---

## 6. Current Sprint

**05_SPRINT.md**

Purpose

Understand:

* Current Sprint
* Active Tasks
* Completed Tasks
* Blocked Tasks
* Backlog

Only implement tasks listed in the current sprint.

---

## 7. AI Operating Instructions

**06_PROMPT.md**

Purpose

Understand:

* AI Workflow
* Development Process
* Communication Style
* Code Generation Process

Follow these instructions throughout the development session.

---

## 8. Code Style Guide

**07_CODE_STYLE.md**

Purpose

Understand:

* Naming Conventions
* Function Style
* File Organization
* Logging Style
* Error Style
* Formatting

Generated code should follow this style consistently.

---

## 9. Project Structure

**08_PROJECT_STRUCTURE.md**

Purpose

Understand:

* Folder Structure
* Package Responsibilities
* File Organization
* Dependency Rules
* Directory Layout
* Module Boundaries

This document explains **where code belongs**.

Read this document before creating new files, folders, or packages.

---

# Optional Technical Documentation

Read only when working in the corresponding domain.

## Backend

docs/backend.md

Topics

* Package Layout
* Services
* Workers
* Backend Design

---

## Frontend

docs/frontend.md

Topics

* React
* Components
* Hooks
* State Management
* Wails Bindings

---

## FFmpeg

docs/ffmpeg.md

Topics

* Stream Copy
* GPU Encoding
* CPU Fallback
* Command Generation
* Filters

---

## GPU

docs/gpu.md

Topics

* AMD AMF
* Encoder Detection
* Hardware Acceleration
* Fallback Strategy

---

## UI

docs/ui.md

Topics

* Layout
* Navigation
* User Experience
* Components
* Screens

---

# Standard Development Workflow

Every implementation should follow this workflow.

```text
Read INDEX

↓

Read Required Documents

↓

Understand Current Phase

↓

Read Current Sprint

↓

Review Existing Code

↓

Implement Small Increment

↓

Build

↓

Verify

↓

Update Documentation

↓

Update Sprint Progress
```

---

# Document Ownership

| Document                | Responsibility                                             |
| ----------------------- | ---------------------------------------------------------- |
| INDEX.md                | Documentation entry point                                  |
| 00_PROJECT.md           | Project vision, scope, and objectives                      |
| 01_RULES.md             | Engineering rules and constraints                          |
| 02_ARCHITECTURE.md      | System architecture                                        |
| 03_ROADMAP.md           | Development phases and milestones                          |
| 04_DECISIONS.md         | Architecture Decision Records (ADR)                        |
| 05_SPRINT.md            | Active sprint and current progress                         |
| 06_PROMPT.md            | AI operating instructions                                  |
| 07_CODE_STYLE.md        | Coding conventions and style guide                         |
| 08_PROJECT_STRUCTURE.md | Project layout, folders, packages, and module organization |

Every document owns exactly one responsibility.

Never duplicate information across documents.

---

# Working Principles

Always:

* Respect the project scope.
* Respect accepted architecture decisions.
* Respect the current roadmap phase.
* Respect the active sprint.
* Respect project architecture.
* Respect project structure.
* Respect the code style.

When uncertain:

Stop.

Explain the ambiguity.

Ask for clarification.

---

# Core Philosophy

This project prioritizes:

* Simplicity over abstraction.
* Readability over cleverness.
* Performance over unnecessary flexibility.
* Maintainability over premature optimization.
* Incremental development over large rewrites.

Every contribution should leave the project cleaner and easier to understand.

---

# AI Checklist

Before writing code, verify that you have:

* Read the required documentation.
* Understood the active roadmap phase.
* Reviewed the current sprint.
* Reviewed accepted ADRs.
* Followed the project structure.
* Followed the code style.
* Planned the smallest possible implementation.

If any item is missing, stop and review the documentation before continuing.

---

# Source of Truth

The `.contexts/` directory is the authoritative source of project documentation.

If code and documentation disagree:

* Do not guess.
* Do not silently modify either one.
* Explain the inconsistency.
* Ask the project owner for clarification.

Consistency between documentation and implementation is a fundamental engineering principle of this project.
