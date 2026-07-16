# design_system.md

# Design System

This document defines the visual design language for the application.

Its purpose is to ensure a consistent, modern, and maintainable user interface.

The design system should be followed by all UI components.

---

# Design Philosophy

The application should feel:

* Modern
* Minimal
* Professional
* Fast
* Clean

Avoid decorative elements.

Every visual element should have a purpose.

---

# Design Goals

Prioritize:

* Consistency
* Clarity
* Accessibility
* Simplicity
* Readability

Visual consistency is more important than uniqueness.

---

# Layout Grid

Use an 8px spacing system.

Preferred spacing scale:

```text
4px
8px
12px
16px
24px
32px
40px
48px
64px
```

Avoid arbitrary spacing values.

---

# Border Radius

Use consistent corner radius.

Recommended scale:

```text
Small   : 6px
Medium  : 8px
Large   : 12px
XL      : 16px
```

Avoid excessive rounding.

---

# Shadows

Use subtle elevation.

Levels:

```text
None

↓

Small

↓

Medium

↓

Large
```

Do not stack multiple shadows.

---

# Typography

Use a single font family.

Recommended:

```text
Inter
```

Fallback:

```text
system-ui
```

---

# Typography Scale

Recommended sizes:

```text
Title

32px

↓

Heading

24px

↓

Subheading

20px

↓

Body

16px

↓

Small

14px

↓

Caption

12px
```

Maintain consistent hierarchy.

---

# Font Weight

Recommended:

```text
Regular

Medium

Semibold

Bold
```

Avoid excessive font weight variation.

---

# Colors

Semantic colors only.

Primary

* Main actions

Success

* Completed operations

Warning

* Recoverable problems

Error

* Failures

Neutral

* Background
* Borders
* Text

Avoid assigning meaning to arbitrary colors.

---

# Light Theme

Characteristics:

* Bright background
* High readability
* Subtle borders
* Soft shadows

---

# Dark Theme

Characteristics:

* Low eye strain
* Strong contrast
* Consistent surface colors

Avoid pure black backgrounds.

---

# Surface Hierarchy

Recommended hierarchy:

```text
Window

↓

Page

↓

Section

↓

Card

↓

Control
```

Every level should be visually distinguishable.

---

# Buttons

Primary

* Main action

Secondary

* Alternative action

Danger

* Destructive action

Ghost

* Minimal action

Buttons should clearly communicate intent.

---

# Button Sizes

Recommended:

```text
Small

Medium

Large
```

Use consistent padding.

---

# Inputs

All inputs should have:

* Label
* Placeholder (when helpful)
* Validation state
* Disabled state

Avoid placeholder-only forms.

---

# Cards

Cards should group related content.

Examples:

* Video Metadata
* Export Settings
* Download Information

Avoid nested cards.

---

# Icons

Use one icon library consistently.

Recommended size:

```text
16px

20px

24px
```

Icons should support labels.

Do not replace text with icons.

---

# Dividers

Use dividers sparingly.

Prefer whitespace before separators.

---

# Spacing

Maintain consistent spacing.

Examples:

```text
Component

↓

16px

↓

Component
```

Avoid crowded layouts.

---

# Forms

Keep forms compact.

Group related controls.

Use vertical alignment.

---

# Dialogs

Dialogs should:

* Focus on one task
* Explain consequences
* Provide clear actions

Avoid large modal workflows.

---

# Tables

Use tables only when structured data benefits from comparison.

Examples:

* Download History
* Export History

Avoid tables for simple lists.

---

# Lists

Prefer cards or lists over tables for small datasets.

Keep list items consistent.

---

# Progress Indicators

Every long-running operation should display:

* Progress Bar
* Percentage
* Current Operation

Avoid indefinite loading indicators when progress is available.

---

# Status Indicators

Recommended statuses:

* Ready
* Running
* Completed
* Warning
* Failed

Use both color and text.

---

# Animations

Animations should be subtle.

Recommended duration:

```text
100ms

150ms

200ms
```

Avoid slow animations.

---

# Transitions

Use transitions only for:

* Hover
* Focus
* Dialog open/close
* Expand/collapse

Avoid unnecessary movement.

---

# Responsive Behavior

The application targets desktop.

Support:

* Window resizing
* Large displays
* Small laptop screens

Avoid fixed-width layouts.

---

# Empty States

Every empty view should include:

* Title
* Description
* Suggested action

Example:

```text
No recent exports

Create your first clip to see it here.
```

---

# Error States

Errors should clearly explain:

* What happened
* Why
* How to recover

Avoid technical messages.

---

# Loading States

Prefer:

* Skeleton loaders
* Progress bars
* Disabled actions

Avoid freezing the interface.

---

# Accessibility

Provide:

* Keyboard navigation
* Visible focus indicators
* Accessible labels
* Adequate contrast

Accessibility is a default requirement.

---

# Tailwind Guidelines

Prefer utility classes.

Avoid excessive custom CSS.

Extract repeated patterns into reusable components.

---

# Component Consistency

Components with similar behavior should:

* Share spacing
* Share sizing
* Share interaction patterns

Avoid creating visually inconsistent variants.

---

# Naming

UI components should have descriptive names.

Examples:

```text
PrimaryButton

ProgressBar

VideoCard

SettingsPanel
```

Avoid generic names.

---

# AI Guidelines

When generating UI:

* Follow the spacing scale.
* Preserve typography hierarchy.
* Use semantic colors.
* Keep layouts uncluttered.
* Reuse existing components.
* Follow the design system before introducing new styles.

If a design choice is unclear, prefer consistency over creativity.

---

# Design Philosophy

A design system exists to reduce decisions.

The best interface is one that feels familiar, predictable, and effortless to use.

Consistency is the highest priority.
