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

build-images:
	@cd db && make build
	@docker build -t ferry-main .

compose-up:
	@cd deploy/docker-compose && make compose-up

compose-down:
	@cd deploy/docker-compose && make compose-down

.PHONY: up, build-db, db-up, migrate-up, migrate-down