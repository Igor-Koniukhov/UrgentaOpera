#MIGRATION=users make migrate-create
include .env

DB_CONNECTION = "${DB_USER}:${DB_PASS}@(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true"

migrate-create:
	mkdir -p ./migrations
	goose -dir ./migrations -table migrations mysql $(DB_CONNECTION) create $(MIGRATION) sql
migrate-up:
	goose -dir ./migrations -table migrations mysql $(DB_CONNECTION) up
migrate-redo:
	goose -dir ./migrations -table migrations mysql $(DB_CONNECTION) redo
migrate-down:
	goose -dir ./migrations -table migrations mysql $(DB_CONNECTION) down
migrate-reset:
	goose -dir ./migrations -table migrations mysql $(DB_CONNECTION) reset
migrate-status:
	goose -dir ./migrations -table migrations mysql $(DB_CONNECTION) status

