# MongoDB Collections for SRE Quiz Backend

## 1. users

Stores user information and quiz attempt status.

**Fields:**

- \_id (ObjectId)
- name (string, required, unique per quiz session)
- createdAt (ISODate)
- updatedAt (ISODate)

## 2. questions

Stores quiz questions and their answer options.

**Fields:**

- \_id (ObjectId)
- questionId (string, e.g. "01")
- text (string)
- answers (array of objects)
  - text (string)
  - isCorrect (boolean)

## 3. attempts

Tracks each user's quiz attempt, answers, score, and timing.

**Fields:**

- \_id (ObjectId)
- userId (ObjectId, ref: users)
- userName (string, denormalized for leaderboard)
- answers (array of objects)
  - questionId (string)
  - answerText (string)
  - isCorrect (boolean)
- score (number)
- totalTimeSeconds (number)
- startedAt (ISODate)
- completedAt (ISODate)

## 4. leaderboard

Precomputed leaderboard entries for fast access (optional, can be derived from attempts).

**Fields:**

- \_id (ObjectId)
- userId (ObjectId, ref: users)
- userName (string)
- score (number)
- totalTimeSeconds (number)
- completedAt (ISODate)

---

**Notes:**

- The leaderboard can be generated dynamically from the `attempts` collection, but a separate collection can be used for caching or performance.
- The `questions` collection is static and can be seeded from the quiz-questions.json file.
- User identity is minimal (name only) unless authentication is added.
- Each user can have only one attempt (enforced by backend logic).
