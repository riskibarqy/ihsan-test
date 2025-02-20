ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

MIGRATE=migrate -path migrations -database "$(DB_CONNECTION_STRING)"

.PHONY: migrate-up migrate-down migrate-new migrate-force migrate-version run

migrate-new:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -format unix $$name

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

migrate-force:
	@read -p "Enter version: " version; \
	$(MIGRATE) force $$version

migrate-version:
	$(MIGRATE) version

run:
	@go run cmd/main.go