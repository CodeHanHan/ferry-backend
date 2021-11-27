# 启动项目
dev-up:
	@swag init
	@go build -o debug-main main.go && ./debug-main

# 构建数据库镜像
build-db:
	@cd db && make build

# 启动数据库容器
db-up:
	@cd db && make up

# 数据库正向迁移
migrate-up:
	@cd db && make migrate-up

# 数据库反向迁移
migrate-down:
	@cd db && make migrate-down

# 指定版本数据库迁移
migrate-to-%:
	@cd db && make migrate-to-%*

# 构建数据库、应用镜像
build-images:
	@cd db && make build
	@docker build -t ferry-main .

# 上线docker-compose
compose-up:
	@cd deploy/docker-compose && make compose-up

# 下线docker-compose
compose-down:
	@cd deploy/docker-compose && make compose-down

.PHONY: up, build-db, db-up, migrate-up, migrate-down