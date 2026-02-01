# Architecture Overview

This document provides a comprehensive overview of the TikTok Business API Golang SDK architecture, explaining how different components work together.

## Table of Contents

- [Project Structure](#project-structure)
- [Core Components](#core-components)
- [API Layer](#api-layer)
- [Model Layer](#model-layer)
- [Enum Definitions](#enum-definitions)
- [Utility Functions](#utility-functions)
- [Client Architecture](#client-architecture)
- [Error Handling](#error-handling)
- [Testing Strategy](#testing-strategy)

## Project Structure

```
tiktok-business/
├── api/                    # API endpoint implementations
│   ├── ad/                 # Ad management APIs
│   ├── ad/aco/             # Advanced campaign optimization APIs
│   ├── catalog/            # Catalog management APIs
│   ├── catalog/feed/       # Feed management APIs
│   ├── catalog/product/    # Product management APIs
│   ├── catalog/set/        # Product set APIs
│   ├── catalog/video/      # Video management APIs
│   ├── event/              # Event tracking APIs
│   ├── pixel/              # Pixel tracking APIs
│   ├── pangle/             # Pangle advertising APIs
│   ├── report/             # Reporting APIs
│   ├── searchad/           # Search advertising APIs
│   ├── showcase/           # Showcase APIs
│   ├── ttuser/             # TikTok user APIs
│   ├── user/               # User APIs
│   └── report/smartplus/   # SmartPlus reporting APIs
├── core/                   # Core SDK functionality
│   ├── client.go           # Main SDK client implementation
│   ├── const.go            # API endpoints and constants
│   └── doc.go              # Package documentation
├── enum/                   # Enumerated types and constants
├── model/                  # Request/response models
│   ├── ad/                 # Ad-related models
│   ├── catalog/            # Catalog-related models
│   ├── event/              # Event-related models
│   ├── pixel/              # Pixel-related models
│   ├── pangle/             # Pangle-related models
│   ├── report/             # Report-related models
│   ├── searchad/           # Search ad-related models
│   ├── showcase/           # Showcase-related models
│   ├── ttuser/             # TikTok user models
│   └── user/               # User models
├── util/                   # Utility functions
├── go.mod                  # Go module definition
├── go.sum                  # Go checksums
├── README.md               # Main documentation
├── CONTRIBUTING.md         # Contribution guidelines
├── ARCHITECTURE.md         # This file
└── .goreleaser.yml         # Release configuration
```

## Core Components

### Core Package (`core/`)

The core package contains the fundamental components that power the SDK:

- **SDKClient**: The main client that handles HTTP requests, authentication, and API communication
- **Constants**: Base URLs and API endpoints for different environments
- **Error Handling**: Common error structures and handling patterns

### Main Client (`core/client.go`)

The `SDKClient` is the central component that:

1. Manages authentication and API credentials
2. Handles HTTP requests to the TikTok Business API
3. Provides methods for GET, POST, and file upload operations
4. Supports both production and sandbox environments
5. Includes debugging capabilities for development

## API Layer

The API layer is organized by functional areas, with each directory containing:

- **Request types**: Structured representations of API requests
- **Response types**: Structured representations of API responses
- **Service functions**: Functions that perform the actual API calls

### API Pattern

Each API follows a consistent pattern:

```go
// Request type definition
type GetRequest struct {
    // Request fields
}

// Response type definition
type GetResponse struct {
    // Response fields
}

// Service function
func Get(ctx context.Context, clt *core.SDKClient, req *GetRequest, accessToken string) (*GetResponse, error) {
    // Implementation that uses the core client
}
```

## Model Layer

The model layer defines the data structures used in API requests and responses:

- **Request Models**: Define the structure of data sent to the API
- **Response Models**: Define the structure of data received from the API
- **Shared Models**: Common data structures used across multiple APIs

Models include proper JSON tags, validation, and helper methods where appropriate.

## Enum Definitions

The enum package contains all enumerated types used throughout the SDK:

- **Ad Formats**: Different ad format options
- **Status Values**: Various status enums (active, paused, deleted, etc.)
- **Targeting Options**: Available targeting criteria
- **Action Scenarios**: Different action scenarios for tracking
- **API Service Types**: Different services within the TikTok ecosystem

## Utility Functions

The utility package contains helper functions that support the SDK:

- **Encoding utilities**: Functions for encoding parameters
- **Validation utilities**: Functions for validating input
- **Formatting utilities**: Functions for data formatting

## Client Architecture

### SDKClient Structure

The SDKClient manages all API interactions:

```go
type SDKClient struct {
    appID        string
    secret       string
    baseURL      string
    httpClient   *http.Client
    debug        bool
    // Other fields
}
```

### Request Flow

1. **Initialization**: Client is initialized with app credentials
2. **Request Preparation**: Request objects are prepared with necessary parameters
3. **Authentication**: Access token is attached to the request
4. **HTTP Call**: The core client makes the actual HTTP request
5. **Response Processing**: Response is decoded and returned
6. **Error Handling**: Errors are properly wrapped and returned

### Environment Support

The client supports multiple environments:

- **Production**: Default environment using the main TikTok Business API
- **Sandbox**: Testing environment for development and testing
- **Custom**: Ability to override base URLs for special requirements

## Error Handling

The SDK follows Go's standard error handling patterns:

- **Specific Error Types**: Custom error types for different scenarios
- **Wrapped Errors**: Errors include context from multiple layers
- **API Error Responses**: Proper parsing of API-specific error responses
- **Network Errors**: Distinction between network and API errors

## Testing Strategy

The SDK includes comprehensive testing at multiple levels:

- **Unit Tests**: Individual functions and methods are tested
- **Integration Tests**: API interactions are tested with mock responses
- **Model Validation**: Request/response models are validated
- **Error Cases**: Error handling scenarios are tested

Tests follow the table-driven pattern for consistency and comprehensive coverage.

## Security Considerations

- **Credential Management**: Secure handling of API credentials
- **HTTPS Enforcement**: All API communications use HTTPS
- **Token Handling**: Proper management of access tokens
- **Input Validation**: Requests are validated before sending

## Performance Considerations

- **Connection Reuse**: HTTP connections are reused for efficiency
- **Minimal Allocations**: Efforts made to minimize memory allocations
- **Efficient Marshaling**: JSON marshaling is optimized
- **Context Support**: All operations support context for cancellation/timeout