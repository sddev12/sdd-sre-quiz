# Backend Overview

This space defines the context for the backend systems of the SRE Quiz app.
The space that provides context for the frontend application is sre-quiz-frontend and will be added to this space.

The backend services will consist of:

- A Go REST API using the Gin web framework
- MongoDB

Full specs for the API service will be added to the space and be present in the sdd-sre-quiz repository.

The API and Database will facilitate the front end application making calls to store and retrieve data on the questions and answers for the quiz, and also to track users scores.
There will be a leaderboard that users can view to see where they placed in the quiz.

The Go REST API will be a containerized application that will run in AWS ECS.
