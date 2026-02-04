.PHONY: help test test-coverage clean install dev build

help:
	@echo "Available commands:"
	@echo "  make test           - Run all tests"
	@echo "  make test-coverage  - Run tests with coverage report"
	@echo "  make install        - Install dependencies"
	@echo "  make dev            - Run application in development mode"
	@echo "  make build          - Build application"
	@echo "  make clean          - Clean build artifacts"

install:
	@echo "📦 Installing dependencies..."
	go mod download
	go mod tidy

test:
	@echo "🧪 Running tests..."
	go test -v -race -timeout 30s ./...

test-coverage:
	@echo "🧪 Running tests with coverage..."
	go test -v -race \
		-coverprofile=coverage.out \
		-covermode=atomic \
		-coverpkg=./services/...,./controllers/...,./repositories/...,./utils/...,./middlewares/...,./models/...,./config/...,./dto/... \
		-timeout 30s \
		./tests/... 2>&1 || true
	@echo ""
	@if [ -f coverage.out ]; then \
		echo "📊 Generating coverage report..."; \
		go tool cover -html=coverage.out -o coverage.html; \
		echo "✅ Coverage report generated: coverage.html"; \
		if command -v open > /dev/null; then \
			open coverage.html; \
		elif command -v xdg-open > /dev/null; then \
			xdg-open coverage.html; \
		else \
			echo "📝 Open coverage.html in your browser"; \
		fi; \
	else \
		echo "❌ Coverage file not generated"; \
	fi

clean:
	@echo "🧹 Cleaning up..."
	rm -f coverage.out coverage.html
	go clean

dev:
	@echo "🚀 Starting development server..."
	go run main.go

build:
	@echo "🔨 Building application..."
	go build -o bin/app main.go
	@echo "✅ Build complete: bin/app"
