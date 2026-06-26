# service-wedding — Go Backend

**MANDATORY: Read `plan.md` before any code changes.**

## Stack
Go + Gin + PostgreSQL (Neon) + Clean Architecture

## Architecture
domain → repository → usecase → delivery/http

## Build
```bash
go build ./...
go vet ./...
```

## Run
```bash
# API server (port 8080)
go run cmd/api/main.go

# Seeder
go run cmd/seeder/main.go
```

## Key Rules
- Parameterized SQL only (`$1, $2`) — no string concatenation
- Error wrapping: `fmt.Errorf("context: %w", err)`
- Context as first param in all functions
- No hardcoded HTML in Go files — use template files
- Max 200 lines per function
