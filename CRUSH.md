# CRUSH.md - TikTok Business API Go SDK

## Build Commands
```bash
go build ./...
```

## Test Commands
```bash
# Run all tests
go test ./...

# Run a specific test file
go test -v ./path/to/package -run TestName

# Run tests with coverage
go test -cover ./...

# Run tests in verbose mode
go test -v ./...
```

## Lint Commands
```bash
# Install golangci-lint first: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run

# Run with specific linters
golangci-lint run --disable-all --enable=goimports --enable=govet --enable=errcheck
```

## Code Style Guidelines

### Imports
- Use standard library imports first, followed by external packages, then internal packages
- Group imports with blank lines between groups
- Use descriptive import aliases when package names conflict or are unclear

### Formatting
- Use `gofmt` for all code formatting
- Use `goimports` to manage imports automatically
- Line length should be reasonable (soft limit 100 characters)
- Use tabs for indentation (default Go style)

### Naming Conventions
- Use camelCase for variables and functions
- Use PascalCase for exported identifiers
- Use descriptive names; avoid abbreviations except for well-known ones (e.g., URL, HTTP)
- Receiver names should be short (one or two letters) and consistent within a package

### Types
- Use structs for data models
- Use interfaces for defining behavior contracts
- Prefer defining types close to where they are used
- Use type aliases sparingly and only when adding semantic meaning

### Error Handling
- Always handle errors explicitly
- Use multi-value returns with error as the last return value
- Wrap errors with context using `fmt.Errorf("context: %w", err)` when appropriate
- Don't ignore errors with `_` unless explicitly intended and documented

### Documentation
- All exported identifiers should have comments
- Use Godoc style comments
- Comments should explain why, not what
- Include examples for complex functionality

## Project Structure
- `api/` - API endpoint implementations
- `model/` - Data models and structures
- `enum/` - Enumerated types and constants
- `core/` - Core client and shared functionality
- `util/` - Utility functions and helpers

## Additional Notes
- This is a Go SDK for TikTok Business API
- Follow REST API conventions where applicable
- Use context.Context for request cancellation and timeouts
- All API calls should be thread-safe
