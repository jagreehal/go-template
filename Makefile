.PHONY: quality test lint format mock update-deps tidy security coverage run docker-build

# Quality check (runs test, lint, format, security checks)
quality: format lint security test

# Testing
test:
	go test -v ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Linting and formatting
lint:
	golangci-lint run

format:
	goimports -w .

# Mock generation
mock:
	mockery --all --keeptree

# Dependency management
update-deps:
	go get -u ./...

tidy:
	go mod tidy

# Security scanning
security:
	gosec ./...

# Environment and running
run:
	gotenv -f .env go run main.go

# Docker build
docker-build:
	docker build -t your-app-name .