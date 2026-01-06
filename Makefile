.PHONY: help deps docker-up docker-down fiber gin echo chi httprouter hertz fasthttp bench test clean

help:
	@echo "Available targets:"
	@echo "  make deps        - Download and install dependencies"
	@echo "  make docker-up   - Start PostgreSQL container"
	@echo "  make docker-down - Stop PostgreSQL container"
	@echo "  make fiber       - Run Fiber server on :8080"
	@echo "  make gin         - Run Gin server on :8080"
	@echo "  make echo        - Run Echo server on :8080"
	@echo "  make chi         - Run Chi server on :8080"
	@echo "  make httprouter  - Run HttpRouter server on :8080"
	@echo "  make hertz       - Run Hertz server on :8080"
	@echo "  make fasthttp    - Run FastHTTP server on :8080"
	@echo "  make bench       - Run benchmark tests (all frameworks)"
	@echo "  make test        - Run all tests"
	@echo "  make clean       - Clean build artifacts"

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

docker-up:
	@echo "Starting PostgreSQL..."
	@docker compose up -d
	@echo "Waiting for PostgreSQL to be ready..."
	@until docker compose exec -T postgres pg_isready -U postgres > /dev/null 2>&1; do sleep 1; done
	@echo "PostgreSQL is ready!"

docker-down:
	@echo "Stopping PostgreSQL..."
	@docker compose down

fiber: deps docker-up
	@echo "Starting Fiber server on :8080..."
	go run ./cmd/fiber

gin: deps docker-up
	@echo "Starting Gin server on :8080..."
	go run ./cmd/gin

echo: deps docker-up
	@echo "Starting Echo server on :8080..."
	go run ./cmd/echo

chi: deps docker-up
	@echo "Starting Chi server on :8080..."
	go run ./cmd/chi

httprouter: deps docker-up
	@echo "Starting HttpRouter server on :8080..."
	go run ./cmd/httprouter

hertz: deps docker-up
	@echo "Starting Hertz server on :8080..."
	go run ./cmd/hertz

fasthttp: deps docker-up
	@echo "Starting FastHTTP server on :8080..."
	go run ./cmd/fasthttp

bench: deps docker-up
	@echo "Running benchmarks..."
	@echo "This will test each framework with 10000 requests"
	@echo ""
	go test -v ./benchmark -run TestAllFrameworks

test: deps
	@echo "Running all tests..."
	go test -v ./...

clean:
	@echo "Cleaning..."
	go clean
	rm -f go.sum
