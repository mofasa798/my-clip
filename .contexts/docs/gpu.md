# gpu.md

# GPU Acceleration Guide

This document defines how GPU acceleration is detected, selected, and used by the application.

GPU acceleration is an optimization.

The application must continue to function correctly even when no supported GPU is available.

---

# Objectives

The GPU subsystem should provide:

* Reliable hardware detection
* Automatic encoder selection
* Fast export performance
* Graceful fallback
* Consistent behavior

---

# Target Hardware

Primary development hardware:

CPU

* AMD Ryzen 5 PRO 4650G

GPU

* AMD Radeon RX 6600 (8 GB)

Memory

* 16 GB DDR4

The implementation should remain compatible with other hardware whenever practical.

---

# Preferred Encoder

Preferred encoder priority:

```text id="3f8qyu"
h264_amf

↓

hevc_amf

↓

libx264
```

Use CPU encoding only when hardware encoding is unavailable or fails.

---

# Detection Strategy

GPU capability should be detected once during application startup.

Detection results should be cached.

Do not probe GPU capabilities before every export.

---

# Detection Flow

```text id="k2rjma"
Application Startup

↓

Detect FFmpeg

↓

Query Available Encoders

↓

Detect AMF Support

↓

Store Detection Result

↓

Ready
```

---

# Detection Responsibilities

The GPU module is responsible for:

* Detecting supported hardware encoders
* Reporting supported codecs
* Reporting encoder availability
* Selecting the preferred encoder

It should not execute export jobs.

---

# Encoder Selection

Selection priority:

1. Stream Copy (if no re-encoding is required)
2. h264_amf
3. hevc_amf
4. libx264

Selection should be automatic.

---

# Supported Operations

GPU encoding should be used for:

* Re-encoding
* Scaling
* Cropping
* Subtitle burn-in
* Codec conversion

GPU encoding is unnecessary for Stream Copy.

---

# Hardware Decoding

Hardware decoding is currently out of scope.

Use software decoding unless future requirements justify GPU decoding.

---

# Capability Model

Represent detected capabilities with a structured model.

Example fields:

* GPU Available
* Vendor
* Device Name
* Supported Encoders
* FFmpeg Version
* AMF Available

Avoid parsing FFmpeg output throughout the application.

---

# Startup Verification

During application startup:

Verify:

* FFmpeg exists
* FFprobe exists
* yt-dlp exists
* GPU encoder availability

Report results to the UI.

---

# Failure Handling

If GPU detection fails:

* Log the error
* Disable GPU encoding
* Continue application startup

Startup should not fail because GPU acceleration is unavailable.

---

# Encoder Validation

Before starting an export:

Verify:

* Requested encoder exists
* Requested encoder is supported
* FFmpeg can use the encoder

Reject invalid requests early.

---

# Fallback Strategy

If GPU encoding fails during export:

```text id="nrxlpn"
Attempt GPU Encoding

↓

Failure

↓

Log Failure

↓

Switch to CPU Encoding

↓

Continue Export
```

The fallback should happen once.

Do not retry GPU encoding repeatedly.

---

# User Configuration

Users may choose:

* Automatic
* GPU
* CPU

Automatic is the default.

If the selected mode is unavailable, explain the reason and use the best available option when appropriate.

---

# Performance Goals

Optimize for:

* Fast startup
* Low CPU usage
* Stable export speed
* Predictable behavior

Avoid unnecessary capability checks.

---

# Logging

Log meaningful GPU events.

Examples:

* GPU detected
* AMF encoder available
* GPU encoding selected
* GPU encoding failed
* CPU fallback activated

Avoid verbose hardware logging.

---

# Error Messages

Provide actionable messages.

Good

```text id="3cv39s"
AMF encoder is unavailable. Falling back to CPU encoding.
```

Avoid

```text id="8x4jja"
GPU error
```

---

# Testing

Verify:

* GPU detected
* GPU unavailable
* Unsupported encoder
* GPU fallback
* Automatic selection
* Manual GPU selection
* Manual CPU selection

Each scenario should behave predictably.

---

# Future Expansion

Potential future support:

* AV1 AMF
* Hardware decoding
* Multi-GPU selection
* Encoder benchmarking
* User-defined encoder presets

These features should be implemented only when scheduled in the roadmap.

---

# AI Guidelines

When generating GPU-related code:

* Keep detection isolated.
* Cache detection results.
* Do not duplicate encoder checks.
* Never execute FFmpeg directly from the GPU module.
* Prefer automatic selection.
* Preserve graceful fallback behavior.
* Keep the GPU module independent from UI code.

---

# GPU Philosophy

GPU acceleration improves performance but is not a requirement for correctness.

The application should always prioritize successful exports over forcing hardware acceleration.

Reliable fallback behavior is more valuable than aggressive GPU usage.
