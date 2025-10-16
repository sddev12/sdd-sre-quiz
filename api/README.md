# SRE Quiz API

## Running Tests Locally

To run the API tests locally, you need a running MongoDB instance. The recommended way is to use Docker Compose:

1. **Start MongoDB using Docker Compose:**

   ```bash
   cd ../localdb
   docker-compose up -d
   ```

   This will start MongoDB and initialize the `sre_quiz` database.

2. **Run the Go tests:**
   ```bash
   cd ../api
   go test ./...
   ```

> **Note:**
>
> - The tests will connect to `mongodb://localhost:27017/sre_quiz` by default (see `.env.example`).
> - Make sure the database is running before running the tests.
> - Test data is cleaned up automatically by the test suite.
>   This is the backend API for the SRE Quiz application, built with Go and Gin.

## Features

- Gin web framework
- Environment variable support via `.env`
- Example root endpoint (`/`)

## Getting Started

1. **Install dependencies:**
   ```bash
   cd api
   go mod tidy
   ```
2. **Copy and configure environment variables:**
   ```bash
   cp .env.example .env
   # Edit .env as needed
   ```
3. **Run the server:**
   ```bash
   go run ./cmd/main.go
   ```

## Example Endpoint

- `GET /` → `{ "message": "SRE Quiz API is running" }`

## Next Steps

- Implement API endpoints as per the [specification](../specs/backend/api/api-contract.md).
