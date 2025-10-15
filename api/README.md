# SRE Quiz API

This is the backend API for the SRE Quiz application, built with Go and Gin.

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
