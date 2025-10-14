# Leaderboard Page Spec

## Purpose

The Leaderboard page displays a ranked list of all users who have completed the quiz, showing their scores and completion times.

## Layout & Structure

- **Header:**

  - Black background, white uppercase text: `QUIZ TERMINAL`.
  - Full-width, bold, and high-contrast (same as homepage).

- **Leaderboard Panel:**

  - Centered, rectangular white panel with a thick black border.
  - Contains:
    - **Title:**
      - Large, uppercase, technical font: `LEADERBOARD`.
      - Centered at the top of the panel.
    - **Leaderboard Table:**
      - Columns:
        - `#` (Leaderboard position)
        - `NAME`
        - `SCORE`
        - `TIME`
      - Each row represents a user who has completed the quiz.
      - Table is sorted by score (descending), then by time taken (ascending for tie-breaks).
      - Table is scrollable if there are many users.
    - **Navigation:**
      - Button: `BACK TO HOME` (outlined, uppercase, black border, white background, left-aligned under the panel).

- **Footer:**
  - Full-width, white background, black top border.
  - Left: `VERSION 0.1`.
  - Right: `© 2025 SRE QUIZ`.

## Visual Example

```
+----------------------------------------------------------+
| [BLACK HEADER] QUIZ TERMINAL                             |
+----------------------------------------------------------+
|                                                          |
|   +-----------------------------------------------+      |
|   | LEADERBOARD                                  |      |
|   |                                               |      |
|   |  #   NAME         SCORE   TIME                |      |
|   |  1   ALICE        19      01:12               |      |
|   |  2   BOB          18      01:30               |      |
|   |  3   CAROL        18      01:45               |      |
|   |  ...                                      ... |      |
|   +-----------------------------------------------+      |
|   | [BACK TO HOME BUTTON]                             |  |
|   +-----------------------------------------------+      |
|                                                          |
+----------------------------------------------------------+
| VERSION 0.1                        © 2025 SRE QUIZ       |
+----------------------------------------------------------+
```

## Behavior & Logic

- Leaderboard is global: shows all users who have completed the quiz.
- Sorted by score (desc), then by time (asc).
- User can return to the homepage via the back button.
