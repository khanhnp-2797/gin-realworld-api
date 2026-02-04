.PHONY: test test-verbose test-coverage test-unit test-integration clean install-deps

# Install dependencies
install-deps:
	go mod download
	go mod tidy

# Run all tests (using manual mocks)
test: install-deps
	go test ./tests/... -v -race -timeout 30s

# Run tests with verbose output
test-verbose: install-deps
	go test ./tests/... -v -race -count=1 -timeout 30s

# Run tests with coverage
test-coverage: install-deps
	go test ./tests/... -v -race -cover -coverprofile=coverage.out -timeout 30s
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run only unit tests
test-unit: install-deps
	go test ./tests/repositories/... ./tests/services/... -v -race -timeout 30s

# Run only integration tests
test-integration: install-deps
	go test ./tests/controllers/... ./tests/middlewares/... -v -race -timeout 30s

# Clean
clean:
	rm -f coverage.out coverage.html

# Run app
run:
	go run main.go

# Build app
build:
	go build -o bin/api .
