# API Contract for SRE Quiz App

This specification details the API contract for the backend REST API of the SRE Quiz App.

## Endpoints

### POST /start-quiz

**Method**: `POST`

**Path**: `/start-quiz`

**Description**

The `POST /start-quiz` endpoint recevies the users name in the body of the request and creates a record for the user in the databse so that the users answers can be tracked against that record as they proceed through the quiz.

The front end will call this endpoint when the user enters their name on the homepage and clicks the start quiz button.

**Endpoint Logic**

Receive POST request

Parse body to get username

Add a user record to the database under the users collection. The record will conform to the users model under the database/collections spec.

If the user record already exists in the database, respond with 400 as defined below

If the record is created successfull, respond with 201 as defined below

**Request Body**

```
{
    "username": "JohnSmith"
}
```

**Response**

| Code | Description               | Example                                     |
| ---- | ------------------------- | ------------------------------------------- |
| 201  | User created successfully | `{ "message": "user created successfully"}` |
| 400  | Validation Error          | `{ "message": "user already exists" }`      |

<br>

### GET /question

**Method**: `GET`

**Path**: `/question`

**Query Params**

`questionId` : `Number (01 - 20)`

Example: `?questionId=05`

**Description**

Gets a question based on the ID (number) of the question requested.
This endpoint is called by the question component of the Quiz page in the frontend application.

**Example Request**

```
GET /question?questionId=01
```

**Response**

| Code | Description               | Example                                                                                                                                                                                                                                            |
| ---- | ------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 200  | User created successfully | `{ "questionId": "01", "question": "What is toil in SRE", "answers": [{ "answerId": "01", "answer": "A programming language" }, { "answerId": "02", "answer": "A monitoring system" }, { "answerId", "03", "answer": "Repetetive manual work" }]}` |
| 500  | Internal Server Error     | `{ "message": "Friendly Error Message" }`                                                                                                                                                                                                          |

<br>

### POST /submit-answer

**Method**: `POST`

**Path**: `/submit-answer`

**Description**

Submits the user's answer for a specific question. Updates the user's record with their answer.

**Request Body**

```
{
    "username": "JohnSmith",
    "questionId": "05",
    "answerId": "02"
}
```

**Response**

| Code | Description           | Example                                   |
| ---- | --------------------- | ----------------------------------------- |
| 200  | Answer submitted      | `{ "message": "answer recorded" }`        |
| 400  | Validation Error      | `{ "message": "invalid input" }`          |
| 404  | User not found        | `{ "message": "user not found" }`         |
| 500  | Internal Server Error | `{ "message": "Friendly Error Message" }` |

<br>

### GET /evaluate

**Method**: `GET`

**Path**: `/evaluate`

**Query Params**

`username` : `String`

Example: `?username=JohnSmith`

**Description**

Evaluates the user's answers after all questions are completed, calculates the score, and returns the result.

**Example Request**

```
GET /evaluate?username=JohnSmith
```

**Response**

| Code | Description           | Example                                                         |
| ---- | --------------------- | --------------------------------------------------------------- |
| 200  | Evaluation complete   | `{ "score": 18, "totalQuestions": 20, "totalTimeSeconds": 95 }` |
| 404  | User not found        | `{ "message": "user not found" }`                               |
| 400  | Quiz not complete     | `{ "message": "quiz not complete" }`                            |
| 500  | Internal Server Error | `{ "message": "Friendly Error Message" }`                       |

<br>

### GET /leaderboard

**Method**: `GET`

**Path**: `/leaderboard`

**Description**

Returns the current leaderboard, sorted by score (descending) and time (ascending for tie-breaks).

**Example Request**

```
GET /leaderboard
```

**Response**

| Code | Description           | Example                                                                                                                        |
| ---- | --------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| 200  | Leaderboard returned  | `[ { "username": "Alice", "score": 19, "totalTimeSeconds": 72 }, { "username": "Bob", "score": 18, "totalTimeSeconds": 90 } ]` |
| 500  | Internal Server Error | `{ "message": "Friendly Error Message" }`                                                                                      |
