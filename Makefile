.PHONY : run-dev

help:
	@echo "make run-dev - Run the application in development mode"

run-dev:
	@echo "Running in development mode"
	go run ./cmd/chess/.
