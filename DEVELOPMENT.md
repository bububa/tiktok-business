# Development Guide

This guide provides detailed information for developers who want to contribute to the TikTok Business API Golang SDK or extend its functionality.

## Table of Contents

- [Environment Setup](#environment-setup)
- [Building the SDK](#building-the-sdk)
- [Running Tests](#running-tests)
- [Code Generation](#code-generation)
- [Adding New API Endpoints](#adding-new-api-endpoints)
- [Testing Strategies](#testing-strategies)
- [Debugging](#debugging)
- [Performance](#performance)
- [Best Practices](#best-practices)

## Environment Setup

### Prerequisites

- Go 1.24 or later
- Git
- A Unix-like operating system (Linux, macOS) or Windows with WSL

### Setting Up Your Environment

1. Clone the repository:

```bash
git clone https://github.com/bububa/tiktok-business.git
cd tiktok-business
```

2. Verify Go installation:

```bash
go version
```

3. Download dependencies:

```bash
go mod download
```

4. Run a basic test to confirm setup:

```bash
go test ./core
```

## Building the SDK

### Standard Build

To build the SDK:

```bash
go build ./...
```

### With Race Detection

For detecting data races during development:

```bash
go build -race ./...
```

### Cross-compilation

To build for different platforms:

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build ./...

# Build for Windows
GOOS=windows GOARCH=amd64 go build ./...

# Build for macOS
GOOS=darwin GOARCH=amd64 go build ./...
```

## Running Tests

### All Tests

Run all tests in the project:

```bash
go test ./...
```

### Verbose Output

Run tests with verbose output:

```bash
go test -v ./...
```

### Specific Package Tests

Run tests for a specific package:

```bash
go test ./api/ad
go test ./model/ad
go test ./core
```

### Test Coverage

Generate coverage report:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Race Condition Testing

Test with race detection:

```bash
go test -race ./...
```

## Code Generation

The SDK follows consistent patterns for API endpoints. While there's no automatic code generation, the following patterns should be maintained:

### API Endpoint Template

```go
package {category}

import (
    "context"
    "github.com/bububa/tiktok-business/core"
    "github.com/bububa/tiktok-business/model/{category}"
)

// {Operation}Request represents the request for {operation}
type {Operation}Request struct {
    // Fields for the request
}

// Encode implements the request interface
func (req *{Operation}Request) Encode() string {
    // Encoding logic
}

// {Operation}Response represents the response for {operation}
type {Operation}Response struct {
    // Fields for the response
}

// {Operation} executes the {operation} API endpoint
func {Operation}(ctx context.Context, clt *core.SDKClient, req *{Operation}Request, accessToken string) (*{Operation}Response, error) {
    var resp {Operation}Response
    err := clt.Post(ctx, req, &resp, accessToken)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
```

## Adding New API Endpoints

### Step 1: Define Models

First, create the request and response models in the appropriate directory under `/model`:

```go
// model/{category}/types.go
package {category}

// {Resource} represents a {resource}
type {Resource} struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    // Other fields
}

// {Operation}Request represents the request for {operation}
type {Operation}Request struct {
    AdvertiserID string `json:"advertiser_id"`
    // Other request fields
}

// {Operation}Response represents the response for {operation}
type {Operation}Response struct {
    Code    int        `json:"code"`
    Message string     `json:"message"`
    Data    *{Resource} `json:"data"`
}
```

### Step 2: Define API Endpoint

Create the API endpoint in the appropriate directory under `/api`:

```go
// api/{category}/{operation}.go
package {category}

import (
    "context"
    "github.com/bububa/tiktok-business/core"
    "github.com/bububa/tiktok-business/model/{category}"
)

// {Operation} implements the {operation} API endpoint
func {Operation}(ctx context.Context, clt *core.SDKClient, req *{Operation}Request, accessToken string) (*{Operation}Response, error) {
    var resp {Operation}Response
    err := clt.Post(ctx, req, &resp, accessToken)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
```

### Step 3: Add Tests

Create comprehensive tests for your new endpoint:

```go
// api/{category}/{operation}_test.go
package {category}

import (
    "testing"
    // other imports
)

func Test{Operation}(t *testing.T) {
    // Test implementation
}
```

### Step 4: Update Documentation

Update the README and other documentation to reflect the new functionality.

## Testing Strategies

### Unit Tests

Focus on testing individual functions and methods:

```go
func TestFunctionality(t *testing.T) {
    tests := []struct {
        name     string
        input    interface{}
        want     interface{}
        wantErr  bool
    }{
        {
            name: "valid_input",
            input:    validInput,
            want:     expectedResult,
            wantErr:  false,
        },
        {
            name: "invalid_input",
            input:    invalidInput,
            want:     nil,
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test logic
        })
    }
}
```

### Integration Tests

Test the interaction between components:

```go
func TestAPICall(t *testing.T) {
    // Mock server setup
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"code": 0, "message": "success", "data": {...}}`))
    }))
    defer server.Close()

    // Test with mock server
    // ...
}
```

### Error Testing

Ensure errors are properly handled:

```go
func TestErrorResponse(t *testing.T) {
    // Test error responses from the API
    // Verify error wrapping and propagation
}
```

## Debugging

### Enable Debug Mode

Enable debugging to see detailed request/response information:

```go
client := core.NewSDKClient("app_id", "secret")
client.EnableDebug()
```

### Logging

The SDK uses Go's standard logging for debug output. You can customize the logger if needed.

### Network Debugging

Use tools like Wireshark or proxy servers to inspect the actual HTTP traffic.

## Performance

### Connection Management

The SDK reuses HTTP connections by default. For high-throughput applications:

```go
// Customize HTTP client settings
httpClient := &http.Client{
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
    Timeout: 30 * time.Second,
}

client := core.NewSDKClientWithHTTPClient(appID, secret, httpClient)
```

### Memory Optimization

- Reuse request objects when possible
- Use buffer pools for frequently allocated objects
- Minimize string concatenation in loops

### Concurrency

The SDK is designed to be thread-safe and can handle concurrent requests:

```go
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(i int) {
        defer wg.Done()
        // Make API calls
    }(i)
}
wg.Wait()
```

## Best Practices

### Error Handling

Always check and handle errors appropriately:

```go
resp, err := SomeAPI(ctx, client, req, token)
if err != nil {
    return fmt.Errorf("failed to call API: %w", err)
}
```

### Context Usage

Always use context for cancellation and timeouts:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

resp, err := API(ctx, client, req, token)
```

### Resource Management

Properly close resources:

```go
// Files, connections, etc.
```

### Input Validation

Validate inputs before making API calls:

```go
if req.AdvertiserID == "" {
    return nil, errors.New("advertiser_id is required")
}
```

### Consistent Naming

Follow the project's naming conventions:

- camelCase for exported identifiers
- Descriptive names for functions and variables
- Consistent terminology across the codebase

### Documentation

Maintain good documentation:

```go
// FunctionName performs the {operation}.
// It returns the {resource} or an error if the operation fails.
func FunctionName(ctx context.Context, clt *core.SDKClient, req *Request, token string) (*Response, error) {
    // Implementation
}
```

## Code Quality

### Linting

Run linters to maintain code quality:

```bash
# Install golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.3

# Run linter
golangci-lint run
```

### Formatting

Use go fmt for consistent formatting:

```bash
go fmt ./...
```

### Vet

Use go vet to catch common mistakes:

```bash
go vet ./...
```

This development guide should help you effectively contribute to the TikTok Business API Golang SDK. If you have any questions, please open an issue in the repository.