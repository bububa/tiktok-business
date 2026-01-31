# Agent Development Guidelines for tiktok-business SDK

## Build Commands

### Build the project:
```bash
go build -v ./...
```

### Build with race detection:
```bash
go build -race -v ./...
```

## Test Commands

### Run all tests:
```bash
go test -v ./...
```

### Run a single test file:
```bash
go test -v ./path/to/package/file_test.go
```

### Run a specific test function:
```bash
go test -run TestFunctionName -v ./path/to/package/
```

### Run tests with coverage:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Run tests with race detection:
```bash
go test -race -v ./...
```

## Lint Commands

### Standard Go linters:
```bash
# Install golangci-lint (if not already installed)
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.3

# Run linter
golangci-lint run
```

### Basic Go tools:
```bash
go fmt ./...
go vet ./...
```

## Code Style Guidelines

### Imports
- Group imports with standard library first, followed by third-party packages
- Use blank import for package side effects when necessary
- Example:
```go
import (
    "encoding/json"
    "fmt"
    "net/url"
    
    "github.com/google/go-cmp/cmp"
)
```

### Formatting
- Use `go fmt` for consistent formatting
- Limit line length to 120 characters when possible
- Use meaningful variable and function names
- Follow Go naming conventions (camelCase for exported items)

### Types
- Use descriptive type names
- Export types that are part of the public API
- Use struct tags appropriately for JSON/marshaling
- Prefer explicit types over type aliases unless there's a clear benefit

### Naming Conventions
- Use camelCase for exported functions/types
- Use PascalCase for exported constants
- Use descriptive names for functions and variables
- Use "Err" prefix for error variables (e.g., `var ErrNotFound`)
- Use "Default" prefix for default values (e.g., `const DefaultTimeout`)

### Error Handling
- Always check and handle errors appropriately
- Use errors wrapping when adding context: `fmt.Errorf("context: %w", err)`
- Return early when possible to reduce nesting
- Provide meaningful error messages

### Testing Patterns
- Use table-driven tests for multiple test cases
- Use the `testing.T` helper functions to properly report errors
- Name test functions descriptively (e.g., `TestFunctionName_CaseDescription`)
- Use subtests for organizing related test cases
- Example test structure:
```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name string
        input interface{}
        want interface{}
    }{
        // test cases here
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test logic
        })
    }
}
```

### Documentation
- Document exported functions, methods, and types with clear GoDoc comments
- Start comments with the function name for clarity
- Provide examples where beneficial
- Document expected behavior and edge cases

### Struct Tags
- Use `json` tags consistently for JSON marshaling
- Use `omitempty` option where appropriate to exclude empty values
- Use custom tag formats for specific encoding needs (like `layout` in query encoder)

### Performance Considerations
- Be mindful of allocations in frequently called functions
- Reuse buffers when appropriate
- Use appropriate data structures for the task
- Avoid unnecessary string concatenation in loops

## Project Structure
- `/api` - API definitions and implementations
- `/core` - Core business logic
- `/enum` - Enumerated types and constants
- `/model` - Data models and structures
- `/util` - Utility functions and helpers

## Dependencies
- Only add dependencies that are well-maintained and widely used
- Keep dependencies up-to-date with `go get -u`
- Use `go mod tidy` to clean up unused dependencies

## Git Workflow
- Write clear, descriptive commit messages
- Keep commits focused on a single logical change
- Use conventional commit format when possible
- Update tests when changing functionality