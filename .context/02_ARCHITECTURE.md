# Architecture

## Overview

React UI

↓

Wails

↓

Go Application Layer

↓

Services

↓

FFmpeg / yt-dlp

↓

Output

---

## Layers

Presentation

- React
- Components
- Pages

Application

- Business Logic
- Services

Infrastructure

- FFmpeg
- yt-dlp
- File System

---

## Dependency Rule

Presentation

↓

Application

↓

Infrastructure

Dependencies must never point upward.

---

## External Processes

Only Go services may execute:

- FFmpeg
- FFprobe
- yt-dlp

React must never invoke external binaries directly.