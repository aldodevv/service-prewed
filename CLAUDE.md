@AGENTS.md

# Quick Reference

- **Read `plan.md`** before any changes — it has the full refactoring roadmap
- Go backend, Clean Architecture: domain → repository → usecase → delivery
- Build: `go build ./...` | Run: `cd cmd/api && go run main.go`
- Database: Neon PostgreSQL, migrations in `migrations/`
- Seeder: `cd cmd/seeder && go run main.go`
- API port: 8080, all routes under `/v1/`
