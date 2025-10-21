# User Journey: Take Quiz

## Overview

This specification describes the SRE Quiz user journey, including starting the quiz and answering questions. It covers the technical flow between the frontend and backend, error handling, and user experience requirements.

## Flow Description

### Step 1: Start Quiz

1. **Homepage Load**

- The user navigates to the homepage of the SRE Quiz application.
- The homepage displays a prompt for the user to enter their name and a "Begin Quiz" button (initially disabled).

2. **User Input**

- The user enters their name into the input field.
- The "Begin Quiz" button becomes enabled when the input is non-empty.

3. **Start Quiz Action**

- The user clicks the "Begin Quiz" button.
- The frontend sends a POST request to the backend API endpoint `/start-quiz` with the following JSON body:
  ```json
  {
    "username": "<entered_username>"
  }
  ```

4. **API Response Handling**

- **Success (201 Created):**
  - The backend creates a new user record in the database if the username does not already exist.
  - The API responds with status 201 and a success message.
  - The frontend navigates the user to the quiz page, passing the username as a parameter.
- **Error (400 or other):**
  - If the username already exists or there is a validation/database error, the API responds with an error status (e.g., 400) and a message.
  - The frontend displays an error message to the user (e.g., "User already exists" or a generic error message).
  - The user remains on the homepage and can try again.

### Step 2: Fetch and Display First Question

1. **Quiz Page Load**

- After a successful start, the user is navigated to the quiz page (e.g., `/quiz?name=<username>`).
- The quiz page/component initializes and reads the username from the query parameter.

2. **Fetch First Question**

- On mount, the quiz component makes a GET request to the backend API endpoint `/question?questionId=01` to fetch the first question.
- The request is sent directly to the backend API (e.g., `http://localhost:8080/question?questionId=01`).

3. **API Response Handling**

- **Success (200):**
  - The API responds with the question data, including the question text and possible answers.
  - The frontend populates the quiz component with the received question and answer options.
- **Error (500 or other):**
  - If there is an error, the API responds with an error status and message.
  - The frontend displays an error message to the user and may offer a retry option.

### Step 3: Submit Answer and Progress

1. **Submit Answer**

   - After selecting an answer, the user clicks the "Submit Answer" button.
   - The frontend sends a POST request to the backend API endpoint `/submit-answer` with the following JSON body:
     ```json
     {
       "username": "<entered_username>",
       "questionId": "<current_question_id>",
       "answerId": "<selected_answer_id>"
     }
     ```

2. **API Response Handling**

   - **Success (200):**
     - The API responds with a success message (e.g., `{ "message": "answer recorded" }`).
     - The frontend increments the question number and makes a GET request for the next question.
     - The next question is loaded and displayed as before.
   - **Error (400, 404, 500):**
     - If there is a validation error, user not found, or server error, the API responds with an error status and message.
     - The frontend displays the error message to the user and prompts them to try submitting their answer again.
     - The user remains on the current question until the answer is successfully submitted.

3. **Repeat Process**
   - Steps 1 and 2 repeat for each question until the user has answered all 20 questions.
   - After the final question is answered and submitted, the frontend can proceed to quiz evaluation or completion steps (not covered here).

## Technical Details

- **Frontend:**
  - The homepage form manages the username input and button state.
  - On submit, it makes a POST request to `/start-quiz` using `fetch` or a similar HTTP client.
  - Handles API responses to either proceed or show an error.
  - On the quiz page, the quiz component reads the username from the query parameter.
  - On mount, the quiz component makes a GET request to `/question?questionId=01` to fetch the first question.
  - Populates the UI with the question and answer options from the API response.
  - Handles API errors and displays user-friendly messages.
- **Backend:**
  - The `/start-quiz` endpoint validates the input and checks for existing users.
  - On success, creates a new user and returns 201.
  - On error, returns an appropriate error status and message.
  - The `/question` endpoint returns the question and answers for the requested questionId.
  - On error, returns an appropriate error status and message.
- **API Contract:**
  - POST `/start-quiz`: `{ "username": "<entered_username>" }`
    - Success: `201 { "message": "user created successfully" }`
    - Error: `400 { "message": "user already exists" }` or other error messages as appropriate.
  - GET `/question?questionId=01`
    - Success: `200 { "questionId": "01", "question": "...", "answers": [ ... ] }`
    - Error: `500 { "message": "Friendly Error Message" }` or other error messages as appropriate.

## Acceptance Criteria

- The user cannot proceed without entering a name.
- The frontend only navigates to the quiz page after a successful API response.
- All error messages from the API are shown to the user in a user-friendly way.
- No duplicate users are created in the database.
- When the user lands on the quiz page, the first question is fetched from the backend and displayed.
- The frontend and backend follow the API contract for all steps described above.
