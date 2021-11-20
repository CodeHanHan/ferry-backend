dev-up:
	@go build -o debug-main main.go && ./debug-main

build-db:
	@cd db && make build

db-up:
	@cd db && make up

migrate-up:
	@cd db && make migrate-up

migrate-down:
	@cd db && make migrate-down

.PHONY: up, build-db, db-up, migrate-up, migrate-down