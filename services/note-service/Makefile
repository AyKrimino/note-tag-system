.PHONY: build run test clean migrate-create migrate-up migrate-down

BINARY=note-service
SRC=main.go
MIGRATIONS_DIR="./internal/db/migrations"

build:
	@go build -o bin/$(BINARY) $(SRC)

run: build
	@./bin/$(BINARY)

test:
	@go test ./...

# Usage: make migrate-create name=migration_name
migrate-create:
	@goose -dir $(MIGRATIONS_DIR) create $(name) sql

migrate-up:
	@goose -dir $(MIGRATIONS_DIR) up

migrate-down:
	@goose -dir $(MIGRATIONS_DIR) down

clean:
	@rm -rf bin
