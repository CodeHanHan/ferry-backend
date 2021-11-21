compose-up:
	@docker-compose up -d ferry-db
	@docker-compose up -d ferry-main

compose-down:
	docker-compose down
	@echo "compose-down ok"

.PHONY: compose-up, compose-down