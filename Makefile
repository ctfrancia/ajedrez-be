.PHONY : run-dev
.SILENT : run-dev

help:
	@echo "make run-dev <...FLAGS> - Run the application in development mode"
	@echo "		-db-dsn=<connection-string> - Set the database connection string (optional)"

run-dev:
	@echo "Running in development mode"
	if [ -z "$(db-dsn)" ]; then \
		echo "DB DSN is not set, using default local connection string"; \
		go run ./cmd/chess/...; \
	else \
		echo "DB DSN: $(db-dsn)"; \
		go run ./cmd/chess/... -db-dsn=$(db-dsn); \
	fi
