# Copilot instructions

## Build, test, and lint
- Run the server: `go run ./cmd/server`
- Build the server: `go build ./cmd/server`
- Tests (none yet): `go test ./...`

## High-level architecture
- `cmd/server/main.go` is the backend entry point. It exposes `GET /api/hello` and serves the static front-end from `client/`.
- `client/` contains the front-end (HTML/CSS/JS) that calls `/api/hello` from the browser and renders the JSON response.
- The top-level folders (`api`, `build`, `configs`, `deployments`, `docs`, `internal`, `pkg`, `test`) are currently empty placeholders containing only `.gitkeep`.

## Key conventions
- Comments and identifier names are in Portuguese; keep new identifiers and comments consistent unless you are asked to localize.
- API responses use JSON with either `message` (success) or `error` (failure).
