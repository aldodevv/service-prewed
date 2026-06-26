.PHONY: build run db-up db-down tidy

build:
	go build -o tmp/main cmd/api/main.go

run:
	go run cmd/api/main.go

db-up:
	docker compose up -d

db-down:
	docker compose down

tidy:
	go mod tidy
