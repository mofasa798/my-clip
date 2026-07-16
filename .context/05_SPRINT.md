# 05_SPRINT.md

# Current Sprint

This document tracks the current development progress.

Unlike the roadmap, this document changes frequently.

Only tasks listed under **Current Sprint** should be actively implemented.

Completed tasks should be moved to **Completed**.

Future work belongs in **Backlog**.

---

# Current Phase

**Phase 1 — Project Foundation**

Status

🟢 In Progress

---

# Current Sprint

## Project Initialization

* [ ] Initialize Wails v3 project
* [ ] Configure React
* [ ] Configure TypeScript
* [ ] Configure Vite
* [ ] Configure Tailwind CSS
* [ ] Configure Go modules

---

## Project Structure

* [ ] Create backend folder structure
* [ ] Create frontend folder structure
* [ ] Create storage directories
* [ ] Create bin directory
* [ ] Create settings.json

---

## Dependency Detection

* [ ] Detect FFmpeg
* [ ] Detect FFprobe
* [ ] Detect yt-dlp
* [ ] Detect AMD AMF encoder

---

## Application

* [ ] Application startup
* [ ] Application shutdown
* [ ] Dependency verification
* [ ] Initialize logger
* [ ] Initialize settings

---

## User Interface

* [ ] Home page
* [ ] Settings page
* [ ] Dependency status panel

---

# Blocked

No blocked tasks.

If a task becomes blocked, document:

* Reason
* Required action
* Expected resolution

---

# Completed

Nothing completed yet.

Completed tasks should remain here until the current phase is finished.

Example:

* [x] Configure Go modules

---

# Backlog

The following tasks belong to future phases.

Do not implement them until the roadmap reaches the appropriate phase.

## Phase 2

* Video metadata
* Thumbnail
* Video download
* Resolution selector

## Phase 3

* Video preview
* Timeline
* Clip editor

## Phase 4

* Export engine
* Stream Copy
* GPU encoding

## Phase 5

* History
* Notifications
* Theme

## Phase 6

* Playlist download
* Batch clipping
* Livestream clipping
* Watermark
* Subtitle burn-in

---

# Notes

Use this section for temporary development notes.

Remove notes once they are no longer relevant.

Example:

* Investigate FFmpeg AMF detection.
* Verify Wails v3 API changes.

---

# Sprint Rules

Always follow these rules.

* Complete one task at a time.
* Keep pull requests small.
* Finish current sprint before starting new work.
* Avoid unrelated refactoring.
* Keep implementations incremental.

---

# Completion Checklist

Before marking a task as completed, verify:

* [ ] Feature works correctly.
* [ ] Code builds successfully.
* [ ] Error handling is complete.
* [ ] Logging is implemented.
* [ ] Documentation is updated if necessary.
* [ ] No unnecessary dependencies were introduced.

---

# AI Instructions

Before writing code:

1. Read the current sprint.
2. Select the next unfinished task.
3. Implement only that task.
4. Do not implement backlog items.
5. Update this document after completing work.
6. Move completed tasks to the **Completed** section.
7. Keep this document synchronized with the actual project state.

This file is the single source of truth for current development progress.
