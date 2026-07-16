# 06_PROMPT.md

# AI Operating Instructions

This document defines how AI coding assistants should behave while contributing to this project.

It is intended to ensure consistency, maintainability, and predictable development throughout the project lifecycle.

---

# Project Documentation

Before making any code changes, read the project documentation in the following order:

1. INDEX.md
2. 00_PROJECT.md
3. 01_RULES.md
4. 02_ARCHITECTURE.md
5. 03_ROADMAP.md
6. 04_DECISIONS.md
7. 05_SPRINT.md
8. 06_PROMPT.md
9. 07_CODE_STYLE.md
10. 08_PROJECT_STRUCTURE.md

Reference additional documents inside `docs/` only when required.

Treat the documentation inside `.contexts/` as the single source of truth.

---

# Primary Objective

Your objective is to help develop this application while preserving:

* Simplicity
* Consistency
* Readability
* Maintainability
* Performance

Do not optimize for hypothetical future requirements.

---

# Before Writing Code

Always perform the following steps.

## Step 1

Understand the requested feature.

Do not assume missing requirements.

Ask questions if necessary.

---

## Step 2

Determine the current development phase.

Only implement work that belongs to the active phase.

---

## Step 3

Read architecture decisions.

Verify whether an ADR already exists.

Follow accepted decisions.

---

## Step 4

Read the current sprint.

Implement only the requested task.

Do not work ahead.

---

## Step 5

Review the existing implementation before writing new code.

Reuse existing components whenever possible.

Avoid duplicate functionality.

---

# While Writing Code

Always:

* Follow project architecture.
* Follow project coding rules.
* Keep implementations incremental.
* Prefer explicit code.
* Prefer readability over cleverness.

Every change should have a clear purpose.

---

# Code Generation Principles

Prefer:

Small changes

↓

Small commits

↓

Small pull requests

↓

Continuous progress

Avoid large rewrites.

Avoid introducing unnecessary complexity.

---

# Feature Development

Implement only what has been requested.

Do not automatically implement:

* Future roadmap items
* Additional improvements
* Nice-to-have features
* Optimizations without evidence

Stay focused on the current task.

---

# Refactoring

Refactor only when:

* It simplifies the implementation.
* It removes duplication.
* It improves readability.
* It fixes an architectural violation.

Do not refactor unrelated code.

---

# Dependencies

Before adding a dependency, ask:

* Does the Go standard library already solve this?
* Does the project already contain a suitable implementation?
* Is the dependency justified?

Avoid dependency bloat.

---

# Architecture

Respect module boundaries.

Never bypass application layers.

Do not introduce enterprise architecture.

Do not invent additional abstraction layers.

---

# Error Handling

Never ignore errors.

Provide meaningful error messages.

Return errors with sufficient context.

Handle external command failures gracefully.

---

# Logging

Log meaningful events only.

Avoid excessive logging.

Logs should help debugging without creating noise.

---

# Documentation

When implementation changes affect documentation:

Update the documentation.

Do not allow documentation to become outdated.

If documentation ownership belongs to another file, update that file instead of duplicating information.

---

# Testing

Before considering a task complete:

* Verify expected behavior.
* Verify error handling.
* Verify edge cases.
* Verify UI interaction when applicable.

Think through failure scenarios.

---

# When Requirements Are Unclear

Do not guess.

Instead:

* Explain the ambiguity.
* Ask a concise clarification question.
* Wait for confirmation.

---

# Things You Must Never Do

Never:

* Rewrite large portions of the project without approval.
* Change architecture decisions.
* Introduce enterprise patterns.
* Add unnecessary dependencies.
* Implement future roadmap phases.
* Ignore existing documentation.
* Duplicate business logic.
* Duplicate documentation.

---

# Completion Checklist

Before finishing any task, verify:

* The solution works.
* The code compiles.
* Error handling is complete.
* Logging is appropriate.
* Architecture is respected.
* Existing code style is preserved.
* No unnecessary complexity was introduced.

---

# Communication Style

When explaining technical decisions:

* Be concise.
* Be objective.
* Explain trade-offs.
* Avoid unnecessary verbosity.

If multiple solutions exist:

Recommend the simplest solution that satisfies the current requirements.

---

# Continuous Improvement

If you discover:

* Duplicate code
* Architectural inconsistencies
* Missing documentation
* Potential bugs

Do not immediately implement changes.

Instead:

1. Explain the observation.
2. Explain why it matters.
3. Ask whether the project owner wants to address it.

---

# Final Principle

This project values long-term maintainability over short-term convenience.

Every contribution should make the codebase easier to understand, easier to maintain, and easier to extend without introducing unnecessary complexity.

When in doubt, choose the simplest solution that fulfills the current requirement.
