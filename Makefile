.PHONY : run-dev setup-dev help
.SILENT : run-dev

# TODO: needs to use these for the run-dev target and make it dynamic based on the flags
db-dsn := $(if $(db-dsn),$(db-dsn),$(""))
port := $(if $(port),$(port),"")
help:
	@echo "make run-dev <...FLAGS> - Run the application in development mode"
	@echo "		-db-dsn=<connection-string> - Set the database connection string (optional)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

setup-dev: ## Setup the development environment
	@echo "====REVERTING TO CLEAN STATE===="
	@echo "---Dropping database---"
	psql postgres -c "DROP DATABASE IF EXISTS my_chess_website"
	@echo "----------------------"

	@echo "---Dropping user---"
	psql postgres -c "DROP USER IF EXISTS chess_admin"
	@echo "----------------------"

	@echo "---Dropping extensions---"
	psql postgres -c "DROP EXTENSION IF EXISTS citext"
	psql postgres -c 'DROP EXTENSION IF EXISTS "uuid-ossp"'
	@echo "----------------------"

	@echo "---Downloading dependencies---"
	# go mod download
	# go mod vendor
	@echo "----------------------"

	@echo "====Setting up development environment===="
	@echo "---Creating database---"
	psql postgres -c "CREATE DATABASE my_chess_website"
	@echo "----------------------"

	@echo "---creating user and password---"
	psql postgres -c "CREATE USER chess_admin WITH PASSWORD 'pa55word'"
	@echo "----------------------"

	@echo "---changing ownership of the database to user just created---"
	psql postgres -c "ALTER DATABASE my_chess_website OWNER TO chess_admin"
	@echo "----------------------"

	@echo "creating extensions for the database with user just created"
	psql chess_admin -d my_chess_website -c "CREATE EXTENSION IF NOT EXISTS citext"
	psql chess_admin -d my_chess_website -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp"'
	@echo "----------------------"

	@echo "====Running migrations===="
	# migrate -path ./migrations -database ${CHESS_DB_DSN} up
	@echo "==========================="

	@echo "====Running seeds===="
	#ago run ./cmd/api/ -limiter-enabled=false;
	# go run ./cmd/seeds/...
	@echo "====================="

run-dev: ## Run the application in development mode
	# db-dsn = $(if $(db-dsn),$(db-dsn),$(""))
	# port = $(if $(port),$(port),"")
	@echo "Running in development mode"
	if [ -z "$(db-dsn)" ]; then \
		echo "DB DSN is not set, using default local connection string"; \
		go run ./cmd/api/ -limiter-enabled=false; \
		else \
		echo "DB DSN: $(db-dsn)"; \
		go run ./cmd/api/... -db-dsn=$(db-dsn); \
		fi
