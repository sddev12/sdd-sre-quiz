# Quiz Page Spec

## Purpose

The Quiz page presents the user with a series of questions, one at a time, and records their answers. It is the core interactive experience of the SRE Quiz.

## Layout & Structure

- **Header:**

  - Black background, white uppercase text: `QUIZ TERMINAL`.
  - Full-width, bold, and high-contrast (same as homepage).

- **Quiz Panel:**

  - Centered, rectangular white panel with a thick black border.
  - Contains:
    - **Quiz Progress:**
      - Top-left: `QUESTION X OF 20` (uppercase, technical font).
      - Top-right: Timer showing elapsed time (e.g., `TIME: 00:42`).
    - **Question:**
      - Large, bold, uppercase, technical font.
      - Centered in the panel.
    - **Answer Choices:**
      - List of possible answers, each as a large, outlined button (uppercase, black border, white background).
      - Only one answer can be selected at a time.
      - Selected answer is visually highlighted (e.g., filled background or bold border).
    - **Submit Button:**
      - Disabled until an answer is selected.
      - Uppercase, bold, black border, white background.
      - Label: `SUBMIT ANSWER`.
    - **Quiz Complete State:**
      - After the last question, show a message: `QUIZ COMPLETE!`
      - Display user's score and total time taken.
      - Show a button: `VIEW LEADERBOARD` (same style as homepage leaderboard button).

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
|   | QUESTION 3 OF 20         TIME: 01:12          |      |
|   |                                               |      |
|   | WHAT DOES SRE STAND FOR?                      |      |
|   |                                               |      |
|   | [SITE RELIABILITY ENGINEER   ]                |      |
|   | [SOFTWARE RELIABILITY ENGINEER]               |      |
|   | [SYSTEMS RELIABILITY ENGINEER ]               |      |
|   | [SECURE RELIABILITY ENGINEER  ]               |      |
|   |                                               |      |
|   | [SUBMIT ANSWER BUTTON]                        |      |
|   +-----------------------------------------------+      |
|                                                          |
+----------------------------------------------------------+
| VERSION 0.1                        © 2025 SRE QUIZ       |
+----------------------------------------------------------+
```

## Behavior & Logic

- User cannot skip questions; must answer each to proceed.
- Timer starts when the quiz begins and stops when the last question is submitted.
- After completion, user sees their score and time, and can view the leaderboard.
- Only one quiz attempt per user.
- When the user selects an answer and clicks "Submit Answer":
  - The frontend sends a POST request to `/submit-answer` with the username, questionId, and answerId.
  - If the API returns an error, the error is shown and the user can retry.
  - On success, the next question is fetched from the backend and displayed.
  - This repeats until all 20 questions are answered.
