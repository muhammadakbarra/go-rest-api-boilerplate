include .env
export

.PHONY: run swagger migrate build test help

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run: Start the development server
run:
	go run cmd/api/main.go

## swagger: Regenerate Swagger documentation
swagger:
	go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/api/main.go --parseDependency --parseInternal

## migrate: Run database migrations
migrate:
	psql "$(DATABASE_URL)" < migrations/000001_create_posts_table.up.sql

## build: Build the application binary
build:
	go build -o bin/api cmd/api/main.go

## test: Run all tests
test:
	go test ./...
