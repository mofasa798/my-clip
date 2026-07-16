# 09_TESTING.md

# Testing Strategy

This document defines the testing philosophy and strategy for the project.

Testing exists to verify correctness, prevent regressions, and maintain confidence when refactoring.

Tests should focus on business behavior rather than implementation details.

---

# Testing Philosophy

Prioritize:

* Reliability
* Simplicity
* Fast execution
* Deterministic results
* Maintainability

Tests should be easy to understand.

---

# Testing Pyramid

Follow this priority:

```text
Unit Tests

↓

Integration Tests

↓

End-to-End Tests (Future)
```

The majority of tests should be unit tests.

---

# What Should Be Tested

Backend

* Business logic
* Validation
* Command generation
* Error handling
* State transitions
* Configuration loading

Frontend

* Components
* Custom hooks
* User interaction
* Rendering
* Service communication

Do not test third-party libraries.

---

# Backend Testing

Language

* Go

Framework

* testing (standard library)

Prefer the Go standard library whenever possible.

Avoid unnecessary testing frameworks.

---

# Backend Test Structure

Example:

```text
clip/

clip_service.go
clip_service_test.go
```

Keep test files close to the implementation.

---

# Unit Test Scope

Unit tests should verify:

* Input validation
* Business rules
* Expected outputs
* Error handling

Do not invoke FFmpeg or yt-dlp in unit tests.

---

# Integration Tests

Integration tests verify interaction between modules.

Examples:

* FFmpeg wrapper
* FFprobe wrapper
* GPU detection
* yt-dlp wrapper

Integration tests may require external dependencies.

---

# External Dependencies

External binaries should be tested separately.

Examples:

* FFmpeg
* FFprobe
* yt-dlp

Mock them during unit testing.

---

# Mocking

Mock only external dependencies.

Examples:

* File system
* FFmpeg
* yt-dlp
* Network

Do not mock simple business logic.

---

# Test Naming

Use descriptive names.

Examples:

```go
func TestCreateClip_WithValidRange(t *testing.T)

func TestCreateClip_InvalidTimestamp(t *testing.T)

func TestDetectGPU_AMFAvailable(t *testing.T)
```

Names should describe behavior.

---

# Frontend Testing

Recommended:

* Vitest
* React Testing Library

Focus on user-visible behavior.

Avoid testing implementation details.

---

# Component Tests

Verify:

* Rendering
* User interaction
* Disabled states
* Loading states
* Error messages

Avoid snapshot-heavy testing.

---

# Hook Tests

Test:

* State transitions
* Async behavior
* Error handling

Hooks should be tested independently.

---

# Service Tests

Frontend services should verify:

* Backend request generation
* Response transformation
* Error propagation

Business rules remain in the backend.

---

# End-to-End Testing

Not required initially.

Future recommendation:

* Playwright

Use E2E tests only for critical workflows.

Examples:

* Download video
* Create clip
* Export clip

---

# Manual Testing Checklist

Before merging changes, verify:

* Application starts
* Dependencies detected
* Download works
* Preview works
* Timeline works
* Export works
* GPU detection works
* CPU fallback works
* Settings persist
* History updates

---

# Regression Testing

Whenever fixing a bug:

1. Reproduce the bug.
2. Write a test if practical.
3. Fix the bug.
4. Verify the test passes.

Every bug fix should reduce future regressions.

---

# Performance Testing

Occasionally verify:

* Startup time
* Export speed
* Memory usage
* CPU usage
* GPU utilization

Performance tests should not run with every unit test.

---

# Code Coverage

Prioritize meaningful coverage.

Suggested targets:

* Business Logic: 90%+
* Wrappers: 70%+
* UI Components: 60%+

Coverage is a guide, not a goal.

Quality is more important than percentage.

---

# CI Readiness

Although this project is local-first, tests should remain suitable for future CI integration.

Tests should:

* Be deterministic
* Avoid machine-specific paths
* Avoid hidden dependencies

---

# AI Guidelines

When generating code:

* Generate tests for business logic.
* Do not generate unnecessary tests.
* Prefer unit tests over integration tests.
* Keep tests readable.
* Avoid brittle assertions.
* Test behavior, not implementation.

---

# Testing Philosophy

Tests are part of the codebase.

A feature is considered complete only when its critical behavior can be verified consistently.

Simple, reliable tests are more valuable than a large number of fragile tests.
