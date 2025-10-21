# Quiz Component

This is the main component of the quiz page. It matches the style of the homepage’s main panel and is responsible for the entire quiz-taking experience.

## Responsibilities

- Display the users name.
- Show quiz progress: current question number (e.g., 1/20).
- Show elapsed time since quiz start.
- Display the current question and answer options.
- Allow answer selection (only one at a time).
- Enable the submit button only after an answer is selected.
- On submit, call the `/submit-answer` API endpoint to record the answer for the user and question.
- If the API returns an error, display the error and allow the user to retry.
- On success, fetch the next question from the backend and display it.
- Repeat until all 20 questions are answered.
- After the last question, display a completion message, the users score, total time, and a button to view the leaderboard.

## Component Diagram

```
+------------------------------------------------------+
| [QUIZ PANEL]                                         |
|------------------------------------------------------|
| User: [NAME]         Question: X/20   Time: MM:SS    |
|                                                      |
| [QUESTION TEXT]                                      |
|                                                      |
| [ANSWER OPTION 1] (selectable button)                |
| [ANSWER OPTION 2]                                    |
| [ANSWER OPTION 3]                                    |
| [ANSWER OPTION 4]                                    |
|                                                      |
| [SUBMIT ANSWER BUTTON] (enabled when answer chosen)  |
|                                                      |
| (If complete)                                        |
|   [QUIZ COMPLETE!]                                   |
|   [SCORE: XX/20]   [TIME: MM:SS]                     |
|   [VIEW LEADERBOARD BUTTON]                          |
+------------------------------------------------------+
```

## Props

| Prop Name      | Type     | Required | Description                                                       |
| -------------- | -------- | -------- | ----------------------------------------------------------------- |
| userName       | string   | Yes      | The name of the user taking the quiz.                             |
| totalQuestions | number   | Yes      | Total number of questions in the quiz (default: 20).              |
| onQuizComplete | function | Yes      | Callback when the quiz is finished (score, time, etc).            |
| fetchQuestion  | function | Yes      | Function to fetch a question by index (returns question/answers). |
| submitAnswer   | function | Yes      | Function to submit an answer (questionId, answerId).              |
| initialTime    | number   | No       | Optionally set the starting time (for resume scenarios).          |

## State

- currentQuestionIndex: number
- currentQuestion: { id, text, answers: [{id, text}] }
- selectedAnswerId: string | null
- elapsedTime: number (seconds)
- score: number
- isComplete: boolean
- isLoading: boolean

## Behavior & Logic

- On mount, fetch the first question and start the timer.
- When an answer is selected, enable the submit button.
- On submit:
  - Call `submitAnswer` (which POSTs to `/submit-answer`) with the selected answer, username, and questionId.
  - If the API returns an error, display the error and allow the user to retry.
  - If not last question, fetch the next question from the backend and reset selection.
  - If last question, stop timer, calculate score, set `isComplete`.
- On completion, show score, time, and leaderboard button.

## Accessibility

- All buttons and options are keyboard accessible.
- High-contrast, large clickable areas.
- No color-only indicators for selection.

## Example Usage

```jsx
<QuizComponent
  userName="Alice"
  totalQuestions={20}
  fetchQuestion={fetchQuestion}
  submitAnswer={submitAnswer}
  onQuizComplete={handleQuizComplete}
/>
```
