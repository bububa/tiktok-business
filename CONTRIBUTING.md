# Contributing to TikTok Business API Golang SDK

Thank you for your interest in contributing to the TikTok Business API Golang SDK! We appreciate your help in making this project better for everyone.

## Table of Contents

- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Code Style](#code-style)
- [Adding New Features](#adding-new-features)
- [Submitting Changes](#submitting-changes)
- [Testing](#testing)
- [Documentation](#documentation)

## Getting Started

Before you begin contributing, please:

1. Fork the repository on GitHub
2. Clone your fork locally
3. Create a new branch for your changes
4. Make sure you have Go 1.24 or later installed

## Development Setup

1. Clone your fork:

```bash
git clone https://github.com/YOUR_USERNAME/tiktok-business.git
cd tiktok-business
```

2. Install dependencies:

```bash
go mod download
```

3. Run tests to ensure everything is working:

```bash
go test ./...
```

## Code Style

We follow standard Go conventions and have specific guidelines:

### Imports
- Group imports with standard library first, followed by third-party packages
- Use blank import for package side effects when necessary

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

### Documentation
- Document exported functions, methods, and types with clear GoDoc comments
- Start comments with the function name for clarity
- Provide examples where beneficial
- Document expected behavior and edge cases

### Struct Tags
- Use `json` tags consistently for JSON marshaling
- Use `omitempty` option where appropriate to exclude empty values
- Use custom tag formats for specific encoding needs

## Adding New Features

When adding new features to the SDK:

1. **API Endpoints**: Add new API endpoints in the appropriate subdirectory under `/api`
2. **Models**: Define request/response models in the `/model` directory
3. **Enums**: Add new enumeration values in the `/enum` directory
4. **Tests**: Include comprehensive tests for your new functionality

### API Endpoint Structure

Each API endpoint should follow this pattern:

```
api/{category}/{endpoint}.go
```

Example structure for a new endpoint:

```go
package {category}

import (
    "context"
    "github.com/bububa/tiktok-business/core"
    "github.com/bububa/tiktok-business/model/{category}"
)

// GetRequest represents the request for getting {resource}s
type GetRequest struct {
    // Request fields
}

// Get implements the get {resource} API endpoint
func Get(ctx context.Context, clt *core.SDKClient, req *GetRequest, accessToken string) (*GetResponse, error) {
    var resp GetResponse
    err := clt.Get(ctx, req, &resp, accessToken)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

// GetResponse represents the response for getting {resource}s
type GetResponse struct {
    // Response fields
}
```

## Submitting Changes

1. **Create an Issue**: If your change addresses a known issue, reference it in your pull request
2. **Branch Naming**: Use descriptive branch names like `feature/new-api-endpoint` or `fix/authentication-bug`
3. **Commits**: Write clear, descriptive commit messages
4. **Pull Request**: Submit a pull request with a clear description of your changes

### Pull Request Requirements

- All tests must pass
- Code must be properly formatted with `go fmt`
- New features must include tests
- Documentation must be updated if applicable
- Changes should follow the project's architecture patterns

### Commit Message Guidelines

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line

## Testing

### Running Tests

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

Run tests with coverage:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Test Structure

Tests should follow the table-driven test pattern:

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

## Documentation

### API Documentation

When adding new API endpoints, ensure you:

1. Add comprehensive godoc comments
2. Include examples in the README if the feature is commonly used
3. Update any relevant guides or tutorials

### README Updates

Keep the README updated with:
- New features and capabilities
- Updated usage examples
- Correct installation instructions
- Links to relevant resources

## Questions?

If you have any questions about contributing, please open an issue on GitHub and we'll help you out.

Thank you for contributing to the TikTok Business API Golang SDK!