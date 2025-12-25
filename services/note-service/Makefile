# Load .env if it exists
ifneq (,$(wildcard .env))
	include .env
	export
endif

.PHONY: build run test clean migrate-create migrate-up migrate-down psql

BINARY=note-service
SRC=./cmd/note-service/main.go
MIGRATIONS_DIR=./internal/db/migrations
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

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
	@goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) up

migrate-down:
	@goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) down

psql:
	@PGPASSWORD=$(DB_PASSWORD) psql \
		-h $(DB_HOST) \
		-p $(DB_PORT) \
		-U $(DB_USER) \
		-d $(DB_NAME)

clean:
	@rm -rf bin
