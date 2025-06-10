include .env
export

.PHONY: migrate


# NAME=<migration_name> make migrate-create
migrate-create:
	migrate create -ext sql -dir migrations -seq $(NAME)

migrate-up:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)" -path ./migrations up

# STEP=<step> make migrate-up-step
migrate-down:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)" -path ./migrations down $(STEP)

migrate-status:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)" -path ./migrations version

# FORCE_VERSION=<version> make migrate-force-version
migrate-force-version:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)" -path ./migrations force $(FORCE_VERSION)
