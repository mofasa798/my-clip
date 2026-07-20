# gpu.md

# GPU Acceleration Guide

This document defines the GPU acceleration strategy for the application.

GPU acceleration is an optimization layer within the Media Layer.

The application should automatically detect available hardware capabilities and select the most efficient encoding strategy.

GPU acceleration must never be treated as a hard requirement.

---

# Purpose

GPU acceleration exists to:

* Reduce export time
* Lower CPU utilization
* Improve responsiveness during long-running exports
* Provide efficient hardware video encoding

If hardware acceleration is unavailable, the application must continue operating using software encoding.

---

# Design Principles

GPU support should be:

* Capability-driven
* Vendor-neutral
* Runtime-detected
* Automatically selected
* Transparent to users

Business logic must never depend on a specific GPU vendor.

---

# Supported Hardware Accelerators

The application should support the following FFmpeg hardware encoders:

| Vendor | Encoder                |
| ------ | ---------------------- |
| NVIDIA | NVENC                  |
| AMD    | AMF                    |
| Intel  | Quick Sync Video (QSV) |

Support is determined by runtime capability detection, not by hardware model.

---

# Hardware Detection

The application should detect hardware capabilities when it starts.

Detection should verify:

* FFmpeg availability
* Supported hardware encoders
* Supported hardware decoders
* Available codecs
* Driver compatibility

Hardware capabilities should be cached for the current session.

---

# Capability Detection Flow

```text
Application Starts
        │
        ▼
Detect FFmpeg
        │
        ▼
Query Available Hardware Encoders
        │
        ▼
Build Capability Profile
        │
        ▼
Expose Capabilities to Backend
```

The application should never assume a particular GPU is installed.

---

# Encoder Selection Strategy

Preferred processing order:

```text
Stream Copy Available
        │
        ▼
Yes
        │
        ▼
Stream Copy
        │
        ▼
No
        │
        ▼
Hardware Encoder Available
        │
        ▼
Yes
        │
        ▼
GPU Encoding
        │
        ▼
No
        │
        ▼
CPU Encoding
```

This strategy should be followed consistently across all exports.

---

# Encoder Priority

When multiple hardware encoders are available, select one based on detected capabilities and user preferences.

The backend should expose supported encoders to the frontend, allowing users to override the automatic selection if desired.

---

# Decoder Strategy

Hardware decoding may be used when supported.

Preferred order:

1. Hardware decoding
2. Software decoding

Decoder selection should remain transparent to the user.

---

# Stream Copy

Whenever possible, prefer Stream Copy.

Advantages:

* Fastest processing
* No quality loss
* Minimal CPU usage
* Minimal GPU usage

If Stream Copy satisfies the export request, re-encoding should be avoided.

---

# CPU Fallback

GPU acceleration is optional.

If hardware encoding fails:

1. Log the failure.
2. Attempt another supported hardware encoder if available.
3. Fall back to software encoding.

Export should continue whenever possible.

---

# Runtime Capability Model

Example:

```text
Capabilities

• Stream Copy
• NVENC
• AMF
• QSV
• Hardware Decode
• Software Encode
```

The backend should expose capabilities through generic models rather than vendor-specific flags.

---

# User Preferences

Users may configure:

* Preferred encoder
* Automatic encoder selection
* Enable or disable hardware acceleration

The backend determines whether the requested configuration is supported.

---

# Error Handling

Examples:

Good:

* Hardware encoder unavailable
* Selected encoder not supported
* Driver initialization failed
* Falling back to software encoding

Avoid exposing raw FFmpeg output directly to users.

---

# Performance Guidelines

The application should:

* Detect capabilities only once per session
* Reuse capability information
* Avoid unnecessary encoder probing
* Prefer efficient encoding paths
* Keep the UI responsive during export

---

# Testing

Unit tests should verify:

* Capability model generation
* Encoder selection logic
* Fallback decisions

Integration tests should verify:

* NVIDIA NVENC
* AMD AMF
* Intel QSV
* Software encoding fallback

Tests should validate behavior rather than specific hardware models.

---

# AI Guidelines

When generating GPU-related code:

* Never assume a specific GPU vendor.
* Detect capabilities dynamically.
* Keep vendor-specific logic isolated.
* Prefer generic capability models.
* Respect the processing priority:

  * Stream Copy
  * Hardware Encoding
  * Software Encoding
* Always implement graceful fallback behavior.

---

# GPU Philosophy

GPU acceleration is a performance optimization, not a dependency.

The application should adapt automatically to the capabilities of the host system, providing the fastest reliable export strategy while remaining fully functional on any modern desktop computer.
