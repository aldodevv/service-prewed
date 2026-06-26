# service-wedding — Go Backend

## Project Overview
Wedding Invitation Platform backend. Go + PostgreSQL + Gin REST API + Clean Architecture.

**ALWAYS read `plan.md` first** before making any changes. It contains the complete refactoring roadmap.

## Architecture
```
cmd/api/main.go              — API server entrypoint
cmd/seeder/main.go            — Database seeder (themes + demo data)
internal/domain/              — Entity structs + Repository interfaces
internal/repository/postgres/  — PostgreSQL implementations
internal/usecase/             — Business logic layer
internal/delivery/http/       — HTTP handlers + router + middleware
migrations/                   — SQL migration files
```

## Critical Rules

### Code Standards
- Clean Architecture: `domain → repository → usecase → delivery/http`
- All SQL queries MUST use parameterized statements (no string concatenation)
- Error wrapping with `fmt.Errorf("context: %w", err)`
- Context (`context.Context`) as first parameter in all functions
- No hardcoded HTML/CSS in Go source files — use template files
- Max 200 lines per function, max 500 lines per file

### Database
- PostgreSQL with JSONB for flexible schema
- Connection string from `DATABASE_URL` env var
- SSL mode required for Neon cloud DB
- Run migrations before seeder: `psql $DATABASE_URL -f migrations/*.sql`

### Build & Run
```bash
# Build API
go build -o cmd/api/api cmd/api/main.go

# Build seeder
go build -o cmd/seeder/seeder cmd/seeder/main.go

# Run API server (port 8080)
cd cmd/api && ./api

# Run seeder
cd cmd/seeder && ./seeder

# Check compilation
go build ./...
go vet ./...
```

### Environment Variables
- `DATABASE_URL` — PostgreSQL connection string (Neon)
- `JWT_SECRET` — JWT signing secret
- `CLOUDINARY_NAME` — Cloudinary cloud name
- `CLOUDINARY_KEY` — Cloudinary API key
- `CLOUDINARY_SECRET` — Cloudinary API secret

### API Base URL
- Development: `http://localhost:8080`
- All routes under `/v1/` prefix
- Public routes: `/v1/public/wedding/:theme/:guest`, `/v1/contacts`
- Admin routes (JWT protected): `/v1/themes`, `/v1/contexts`, `/v1/guests`, `/v1/assets`

### Seeder
- The seeder file (`cmd/seeder/main.go`) seeds 3 premium themes + demo contexts + guests
- Theme HTML templates should be extracted to `cmd/seeder/templates/` directory
- Theme JSON data should be extracted to `cmd/seeder/data/` directory
- Target: seeder main.go max 200 lines (read from template files)

## Refactoring Notes
- See `plan.md` for the complete list of tasks
- Priority: RSVP feature, seeder extraction, public handler optimization
- New tables needed: `rsvps`, `context_assets`
- Public handler currently fetches ALL guests then loops — must use direct slug query
