ifneq (,$(wildcard .env))
    include .env
    export
endif

DB_CONNECTION_STRING=postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATE=migrate -path migrations -database "$(DB_CONNECTION_STRING)"

.PHONY: migrate-up migrate-down migrate-new migrate-force migrate-version run

migrate-new:
	@read -p "migration name ? : " migration_name \
	&& migrate  create  -ext sql -dir migrations -seq "$${migration_name}" 

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-force:
	@read -p "Enter version: " version; \
	$(MIGRATE) force $$version

migrate-version:
	$(MIGRATE) version

run:
	@go run cmd/main.go
