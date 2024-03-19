.PHONY : run-dev
.SILENT : run-dev

# TODO: needs to use these for the run-dev target and make it dynamic based on the flags
db-dsn := $(if $(db-dsn),$(db-dsn),$(""))
port := $(if $(port),$(port),"")
help:
	@echo "make run-dev <...FLAGS> - Run the application in development mode"
	@echo "		-db-dsn=<connection-string> - Set the database connection string (optional)"

run-dev:
	# db-dsn = $(if $(db-dsn),$(db-dsn),$(""))
	@echo "Running in development mode"
	if [ -z "$(db-dsn)" ]; then \
		echo "DB DSN is not set, using default local connection string"; \
		go run ./cmd/chess/...; \
		else \
		echo "DB DSN: $(db-dsn)"; \
		go run ./cmd/chess/... -db-dsn=$(db-dsn); \
		fi
