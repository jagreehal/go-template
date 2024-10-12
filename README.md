# Go Project Template

My template for Go projects.

## Features

- Testing with `testify`
- Linting with `golangci-lint`
- Code formatting with `goimports`
- Mocking with `mockery`
- Dependency management with Go Modules
- Logging with `logrus`
- Configuration management with `viper`
- Error handling with `pkg/errors`
- Security scanning with `gosec`
- Continuous Integration support
- Docker support
- Code coverage reporting

## Getting Started

### Clone the repository

```bash
git clone https://github.com/jagreehal/go-template.git
```

Install dependencies:

```bash
go mod download
```

## Common Commands

Run Tests

```bash
make test
```

Run Linter

```bash
make lint
```

Format Code

```bash
make format
```

Generate Mocks

```bash
make mock
```

Update Dependencies

```bash
make update-deps
```

Tidy

```bash
make tidy
```

Run Security Checks

```bash
make security
```

Build Docker Image

```bash
make docker-build
```

Generate Coverage Report

```bash
make coverage
```
