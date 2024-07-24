include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #
.PHONY: run-dev setup-dev help audit vendor confirm build/api
.SILENT : run-dev help

## help: Show this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


confirm:
	@read -p "Are you sure? [y/N] " response; \
	if [ "$$response" != "y" ]; then \
		echo "Exiting..."; \
		exit 1; \
	fi
# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
## setup-dev: prepare the development environment
setup-dev: confirm
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

## run-dev: run the application in development mode
run-dev:
	@echo "Running in development mode"
	if [ -z "$(db-dsn)" ]; then \
		echo "DB DSN is not set, using default local connection string"; \
		go run ./cmd/api/ -limiter-enabled=false; \
		else \
		echo "DB DSN: $(db-dsn)"; \
		go run ./cmd/api/... -db-dsn=$(db-dsn); \
		fi
# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy dependencies and vendor them
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/api: build the cmd/api application
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s -w' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o=./bin/linux_amd64/api ./cmd/api
