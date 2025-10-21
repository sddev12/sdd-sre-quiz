# Pages & User Journeys

## Pages

There will be 4 pages total for the application

### Homepage

The landing page where users can choose to take the quiz or view the leaderboard.

The user should be able to view the leaderboard at any time, not only when they have completed the quiz so the leaderboard button should always be active on the homepage.

The user should only be able to take the quiz once.

### Quiz page

This is a dynamic page which loads a question component.

The question component handles the logic for the quiz.

It calls for the first question, allows the user to select an answer and then submits that answer. Users cannot skip questions, they MUST provide an answer for each even if it is just a guess.

It will then load the next queston and the next until the quiz is complete

When the quiz is complete it will show the users score and allow them to view the leaderboard

The user is timed on the entire quiz so that when they complete the 20 questions, there is a log of how long it took them.

### Leaderboard page

This page will show a list of users arranged by their score so that users can see where they places vs everyone else in the quiz.
It may be that we need to log the time the user took to complete the quiz to make the leaderboard more granular.

The info for each user shown on the leaderboard is:

- Leaderboard position
- Name
- Score
- Time taken to complete quiz

The leaderboard is global meaning that it will show all users that ever took the quiz and their placement on the leaderboard.

## User Journeys

The application will have two user journeys:

- Take the quiz
- View the leaderboard

### Take the quiz

1. User arrives on the homepage
2. User enters their name and clicks the begin quiz button
3. User is moved to the Quiz page
4. A Question component is displayed on the page. The component makes a GET request to the backend API to retrieve the first question and the possible answer choices.
5. The question is displayed in the component and the user can select one answer.
6. Once an answer is selected a submit button on the component is enabled.
7. The user clicks submit, and the frontend POSTs to `/submit-answer` with the username, questionId, and answerId.
8. If the API returns an error, the error is shown and the user can retry submitting their answer.
9. On success, the next question is loaded for the user via another API call.
10. This process is repeated until the user has answered 20 questions.
11. Once the 20th and final question is answered the user sees a quiz completed message and is given their total score to see.

### View the Leaderboard

1. The user arrives on the homepage
2. There is a 'View Leaderboard' button on the homepage
3. The user clicks the 'View Leaderboard' button
4. The user is taken to the leaderboard page
5. The leaderboard page

More detail specs for each page will be provided in the specs/frontend/pages directory.
