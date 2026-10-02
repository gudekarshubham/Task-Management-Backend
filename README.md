# Task-Management-Backend

Task management REST API in Go (Gin, GORM, PostgreSQL) with JWT-protected routes.

## Setup

1. Start a PostgreSQL database.
2. Copy `internal/infrastructure/.env.example` to `internal/infrastructure/.env` and fill in your values.
   The `.env` file is git-ignored. You can also skip the file and set `PORT`, `DBURL` and `JWTSECRET`
   as environment variables instead.
3. Run the server from the repo root:

   ```sh
   go run ./server
   ```

Tables are created automatically on startup.
