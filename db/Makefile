MYSQL_ROOT_PASSWORD := codehanhan

up:
	@docker run --name ferry-db --restart=always -p 13306:3306 -v ~/.ferry-data/mysql:/var/lib/mysql -d ferry-mysql

build:
	@docker build -t ferry-mysql .

migrate-up:
	@migrate -source=file://schema -database "mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(localhost:13306)/ferry" -verbose up

migrate-down:
	@migrate -source=file://schema -database "mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(localhost:13306)/ferry" -verbose down

migrate-up-%:
	@migrate -source=file://schema -database "mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(localhost:13306)/ferry" -verbose up $*

migrate-down-%:
	@migrate -source=file://schema -database "mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(localhost:13306)/ferry" -verbose down $*

migrate-to-%:
	@migrate -source=file://schema -database "mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(localhost:13306)/ferry" -verbose goto $*

.PHONY: up, build

# migrate create -ext sql -dir schema/hello -seq table_name