# 04_DECISIONS.md

# Architecture Decision Records (ADR)

This document records important architectural and technical decisions made for this project.

Each accepted decision is considered **final** unless explicitly changed by the project owner.

AI assistants should follow these decisions without proposing alternative architectures unless requested.

---

# ADR-001

## Title

Desktop Framework

## Decision

Use **Wails v3**.

## Status

Accepted

## Rationale

* Native Go integration.
* Small executable size.
* Low memory usage.
* Excellent desktop performance.
* No embedded Chromium runtime.

---

# ADR-002

## Title

Backend Language

## Decision

Use **Go**.

## Status

Accepted

## Rationale

* Fast compilation.
* Excellent concurrency.
* Simple deployment.
* Native integration with Wails.
* Strong support for external process management.

---

# ADR-003

## Title

Frontend Framework

## Decision

Use **React + TypeScript + Vite**.

## Status

Accepted

## Rationale

* Mature ecosystem.
* Excellent Wails support.
* Fast development.
* Strong TypeScript integration.

---

# ADR-004

## Title

Styling

## Decision

Use **Tailwind CSS**.

## Status

Accepted

## Rationale

* Rapid UI development.
* Consistent design system.
* Minimal custom CSS.

---

# ADR-005

## Title

Video Processing Engine

## Decision

Use **FFmpeg**.

## Status

Accepted

## Rationale

* Industry standard.
* High performance.
* Broad codec support.
* Reliable and well-tested.

---

# ADR-006

## Title

Video Downloader

## Decision

Use **yt-dlp**.

## Status

Accepted

## Rationale

* Reliable metadata extraction.
* Supports YouTube and Kick.
* Mature project.
* Actively maintained.

---

# ADR-007

## Title

GPU Acceleration

## Decision

Prefer **AMD AMF**.

## Status

Accepted

## Rationale

Target hardware includes an AMD Radeon RX6600.

Hardware encoding significantly reduces processing time when re-encoding is required.

---

# ADR-008

## Title

Video Processing Priority

## Decision

Always use:

1. Stream Copy
2. GPU Encoding
3. CPU Encoding

## Status

Accepted

## Rationale

This order provides the best balance between performance and output quality.

---

# ADR-009

## Title

Application Type

## Decision

Desktop only.

## Status

Accepted

## Rationale

This application is intended exclusively for local personal use.

Cloud deployment is intentionally excluded.

---

# ADR-010

## Title

Architecture Style

## Decision

Use a simple layered architecture.

## Status

Accepted

## Rationale

The project does not require enterprise architecture.

Simple layering improves maintainability.

---

# ADR-011

## Title

Database

## Decision

No database.

## Status

Accepted

## Rationale

The application is single-user.

JSON configuration and lightweight storage are sufficient.

---

# ADR-012

## Title

Queue System

## Decision

Use Go Worker Pool.

## Status

Accepted

## Rationale

No Redis or external queue is required.

Go provides efficient concurrency primitives.

---

# ADR-013

## Title

Configuration

## Decision

Store configuration in:

settings.json

## Status

Accepted

## Rationale

Human-readable.

Easy backup.

Easy debugging.

---

# ADR-014

## Title

History Storage

## Decision

Store history locally.

## Status

Accepted

## Rationale

No persistent database is necessary.

History remains lightweight.

---

# ADR-015

## Title

External Commands

## Decision

Wrap every external executable inside dedicated Go services.

## Status

Accepted

## Rationale

Centralized error handling.

Better testing.

Cleaner architecture.

---

# ADR-016

## Title

Progress Communication

## Decision

Use Wails Events.

## Status

Accepted

## Rationale

Native desktop communication.

No WebSocket required.

---

# ADR-017

## Title

Logging

## Decision

Store logs locally.

## Status

Accepted

## Rationale

Simple debugging.

No external logging service.

---

# ADR-018

## Title

Dependency Philosophy

## Decision

Prefer the Go standard library whenever practical.

## Status

Accepted

## Rationale

Reduce dependencies.

Improve maintainability.

Reduce upgrade risks.

---

# ADR-019

## Title

Architecture Complexity

## Decision

Avoid enterprise patterns.

## Status

Accepted

## Rationale

The project is intentionally simple.

Avoid introducing:

* Repository Pattern
* CQRS
* Event Bus
* Service Locator
* Generic Dependency Injection

unless explicitly requested.

---

# ADR-020

## Title

Processing Location

## Decision

All video processing occurs locally.

## Status

Accepted

## Rationale

No cloud processing.

No remote workers.

No external services.

---

# ADR-021

## Title

Project Philosophy

## Decision

The application is optimized for:

* Simplicity
* Performance
* Maintainability
* Local-first workflow

## Status

Accepted

## Rationale

The application should remain understandable and maintainable by a single developer.

---

# ADR-022

## Title

Project Scope

## Decision

This project intentionally excludes:

* Authentication
* Multi-user support
* REST API
* Cloud deployment
* Docker
* Kubernetes
* PostgreSQL
* Redis
* RabbitMQ
* Kafka
* Microservices

## Status

Accepted

## Rationale

These technologies provide no meaningful benefit for a single-user desktop application.

---

# ADR Lifecycle

Each decision has one of the following statuses:

* Proposed
* Accepted
* Deprecated
* Replaced

Only **Accepted** decisions should guide implementation.

---

# Modifying Decisions

An accepted decision should only be changed when:

* The project owner explicitly requests it.
* A new ADR supersedes the previous one.
* The original decision is no longer technically valid.

Do not replace accepted decisions automatically.

---

# AI Instructions

Before suggesting architectural changes:

1. Read this document.
2. Verify whether an ADR already exists.
3. If an accepted ADR exists, follow it.
4. Do not propose alternative architectures unless explicitly requested.
5. If a new architectural decision is necessary, propose a new ADR instead of silently changing existing behavior.

---

# Guiding Principle

Architecture decisions exist to preserve consistency over time.

A simpler architecture that remains consistent is preferred over a more flexible architecture that introduces unnecessary complexity.
