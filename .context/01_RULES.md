# AI Coding Rules

## General

- Simplicity first.
- Readability over cleverness.
- Avoid over-engineering.
- Never implement future features unless explicitly requested.
- Make the smallest possible change.

---

## Backend

- Business logic belongs in Go.
- React must never call FFmpeg directly.
- Wrap every external command.
- Use Context for cancellation.
- Prefer standard library.

---

## Frontend

- Functional Components only.
- Keep components small.
- Avoid unnecessary state.
- Prefer composition.

---

## Architecture

Do NOT introduce:

- Repository Pattern
- Factory Pattern
- CQRS
- Event Bus
- Dependency Injection Framework
- Generic abstractions

unless explicitly requested.

---

## Performance

Always prefer:

Stream Copy

↓

GPU Encoding

↓

CPU Encoding

Only encode when necessary.

---

## Dependencies

Every dependency must have a clear justification.

Prefer built-in Go packages whenever possible.

---

## Documentation Rules

Every piece of information must have a single source of truth.

Do not duplicate documentation across multiple files.

If information belongs to another document, reference it instead of copying it.

---