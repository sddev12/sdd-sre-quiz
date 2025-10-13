# Homepage Specification

## Overview

This document describes the specification for the SRE Quiz Application homepage, as implemented in the initial version.

## Purpose

The homepage serves as the entry point for users to begin the SRE quiz. It sets the tone for the application's minimal, high-contrast, and technical aesthetic, inspired by control-panel UIs.

## Layout & Structure

- **Header:**
  - Black background, white uppercase text: `QUIZ TERMINAL`.
  - Full-width, bold, and high-contrast.
- **Main Panel:**
  - Centered, rectangular white panel with a thick black border.
  - Contains:
    - Title: `SRE QUIZ` (large, uppercase, technical font).
    - Welcome message: Brief, uppercase, centered.
    - Name input:
      - Centered, monospaced, uppercase.
      - Only a thick black bottom border (no side or top borders).
      - Placeholder text: `ENTER YOUR NAME` (disappears on focus).
      - Custom flashing block cursor for a terminal-like effect.
      - Input is required to enable the "Begin Quiz" button.
    - "Begin Quiz" button:
      - Uppercase, bold, black border, white background.
      - Disabled until a name is entered.
- **Footer:**
  - Full-width, white background, black top border.
  - Left: `VERSION 0.1`.
  - Right: `© [current year] SRE QUIZ`.

## Color Palette

- Background: White (`#FFFFFF`)
- Foreground/Text/Borders: Black (`#000000`)
- Accent/Hover: Light Gray (`#EAEAEA`)

## Typography

- Monospaced or technical font (e.g., IBM Plex Mono, Roboto Mono).
- All text is uppercase.
- Title: 24–28px, bold.
- Body/Input: 12–16px.

## Interactivity

- Input placeholder disappears on focus.
- Input shows a custom flashing block cursor.
- "Begin Quiz" button is disabled until a name is entered.
- Button and input have subtle hover/focus effects (no animation longer than 150ms).

## Accessibility

- High contrast for all text and interactive elements.
- Keyboard accessible (tab focus, enter key for button).
- No color-only indicators.

## Visual Example

```
+----------------------------------------------------------+
| [BLACK HEADER] QUIZ TERMINAL                             |
+----------------------------------------------------------+
|                                                          |
|   +-----------------------------------------------+      |
|   | SRE QUIZ                                      |      |
|   |                                               |      |
|   | WELCOME TO THE SITE RELIABILITY...            |      |
|   |                                               |      |
|   | [ENTER YOUR NAME________█]                    |      |
|   |                                               |      |
|   | [BEGIN QUIZ BUTTON]                           |      |
|   +-----------------------------------------------+      |
|                                                          |
+----------------------------------------------------------+
| VERSION 0.1                        © 2025 SRE QUIZ       |
+----------------------------------------------------------+
```

## Version

- Spec version: 1.0
- Date: 2025-10-13
- Author: Sam Dickinson
