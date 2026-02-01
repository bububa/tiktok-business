# User Guide

This guide provides practical information for developers integrating the TikTok Business API Golang SDK into their applications.

## Table of Contents

- [Getting Started](#getting-started)
- [Authentication](#authentication)
- [Making API Calls](#making-api-calls)
- [Handling Responses](#handling-responses)
- [Error Handling](#error-handling)
- [Rate Limiting](#rate-limiting)
- [Common Use Cases](#common-use-cases)
- [Troubleshooting](#troubleshooting)
- [Resources](#resources)

## Getting Started

### Installation

Add the SDK to your project:

```bash
go get github.com/bububa/tiktok-business
```

### Import the SDK

```go
import (
    "github.com/bububa/tiktok-business/core"
    "github.com/bububa/tiktok-business/api/ad"
    "github.com/bububa/tiktok-business/model/ad"
)
```

### Initialize the Client

```go
// Initialize with your app credentials
client := core.NewSDKClient("your_app_id", "your_app_secret")
```

## Authentication

### Access Token

Most API calls require an access token. You can obtain access tokens through TikTok's OAuth flow or use existing tokens.

```go
// Using an existing access token
accessToken := "your_access_token"

// Make API calls with the token
resp, err := ad.Get(context.Background(), client, req, accessToken)
```

### OAuth Integration

For OAuth token management:

```go
import "github.com/bububa/tiktok-business/api/ttuser"

// Generate OAuth URL
oauthURL := ttuser.OAuthURL("your_app_id", "your_redirect_uri", "state", []string{"ads_read", "ads_manage"})

// Exchange authorization code for access token
tokenResp, err := ttuser.Token(context.Background(), client, &ttuser.TokenRequest{
    AppID:          "your_app_id",
    Secret:         "your_app_secret",  // Note: For token exchange, secret might be passed differently depending on API
    AuthCode:       "authorization_code_from_callback",
    GrantType:      "authorization_code",
    RedirectURI:    "your_redirect_uri",
}, "")
```

## Making API Calls

### Basic API Call Pattern

All API calls follow this pattern:

```go
// 1. Create a request object
req := &ad.GetRequest{
    AdvertiserID: "your_advertiser_id",
    // Set other parameters
}

// 2. Call the API function
resp, err := ad.Get(context.Background(), client, req, "access_token")

// 3. Handle response/error
if err != nil {
    // Handle error
    log.Printf("API call failed: %v", err)
    return
}

// Process response
for _, adItem := range resp.Data.List {
    fmt.Printf("Ad: %s - %s\n", adItem.AdID, adItem.AdName)
}
```

### Common Parameters

Most API requests include common parameters:

```go
// Advertiser ID (required for most operations)
AdvertiserID: "your_advertiser_id"

// Filtering options
Filtering: &model.Filtering{
    Status: []string{"ACTIVE", "PAUSED"}, // Filter by status
    // Other filter options
}

// Pagination
Page:     1,      // Page number (1-indexed)
PageSize: 10,     // Number of items per page
```

## Handling Responses

### Response Structure

API responses typically follow this structure:

```go
type Response struct {
    Code    int         `json:"code"`    // API response code (0 = success)
    Message string      `json:"message"` // Human-readable message
    Data    interface{} `json:"data"`    // Actual response data
    RequestID string   `json:"request_id"` // Unique request identifier for debugging
}
```

### Processing Successful Responses

```go
if resp.Code == 0 { // Success code
    // Process the data
    for _, item := range resp.Data.Items {
        // Handle each item
    }
} else {
    // Handle API error
    log.Printf("API returned error: %s (code: %d)", resp.Message, resp.Code)
}
```

### Pagination

Handle paginated responses:

```go
page := 1
pageSize := 100

for {
    req := &ad.GetRequest{
        AdvertiserID: "your_advertiser_id",
        Page:         page,
        PageSize:     pageSize,
    }
    
    resp, err := ad.Get(context.Background(), client, req, "access_token")
    if err != nil {
        log.Printf("Error fetching page %d: %v", page, err)
        break
    }
    
    // Process current page items
    for _, item := range resp.Data.List {
        // Process item
    }
    
    // Check if there are more pages
    if len(resp.Data.List) < pageSize {
        break // Last page
    }
    
    page++
}
```

## Error Handling

### SDK Errors

The SDK returns Go error values for various scenarios:

```go
resp, err := ad.Get(context.Background(), client, req, "access_token")
if err != nil {
    // Check for different types of errors
    if errors.Is(err, context.DeadlineExceeded) {
        // Request timed out
        log.Println("Request timed out")
    } else if errors.Is(err, context.Canceled) {
        // Request was canceled
        log.Println("Request was canceled")
    } else {
        // Other error
        log.Printf("API call failed: %v", err)
    }
    return
}
```

### API-Specific Errors

Some errors come from the TikTok API itself:

```go
if resp.Code != 0 {
    switch resp.Code {
    case 40001:
        // Invalid access token
        log.Println("Invalid or expired access token")
        // Refresh token logic here
    case 40002:
        // Insufficient permissions
        log.Println("Insufficient API permissions")
    case 40003:
        // Rate limit exceeded
        log.Println("Rate limit exceeded")
        time.Sleep(60 * time.Second) // Wait before retrying
    default:
        log.Printf("Unknown API error: %d - %s", resp.Code, resp.Message)
    }
}
```

### Retry Logic

Implement retry logic for transient failures:

```go
func callWithRetry(ctx context.Context, maxRetries int, fn func() error) error {
    var lastErr error
    
    for i := 0; i <= maxRetries; i++ {
        err := fn()
        if err == nil {
            return nil // Success
        }
        
        lastErr = err
        
        // Don't sleep after the last attempt
        if i < maxRetries {
            // Exponential backoff
            waitTime := time.Duration(math.Pow(2, float64(i))) * time.Second
            select {
            case <-time.After(waitTime):
            case <-ctx.Done():
                return ctx.Err()
            }
        }
    }
    
    return lastErr
}
```

## Rate Limiting

### Understanding Limits

TikTok Business API has rate limits that vary by endpoint. Common limits include:

- **General API calls**: Usually limited to a certain number of requests per minute/hour
- **High-volume endpoints**: May have stricter limits
- **Read vs Write**: Different limits for read and write operations

### Handling Rate Limits

```go
func handleRateLimiting() {
    // When you receive a rate limit error (typically code 40003)
    // Wait before making the next request
    
    // Implement exponential backoff
    baseDelay := time.Second
    maxDelay := 60 * time.Second
    
    for attempt := 0; attempt < maxRetries; attempt++ {
        resp, err := ad.Get(context.Background(), client, req, token)
        if err != nil {
            // Handle non-rate limit errors
            return err
        }
        
        if resp.Code == 40003 { // Rate limit exceeded
            delay := baseDelay * time.Duration(math.Pow(2, float64(attempt)))
            if delay > maxDelay {
                delay = maxDelay
            }
            
            time.Sleep(delay)
            continue
        }
        
        // Success or other error
        return processResponse(resp)
    }
    
    return errors.New("max retries exceeded for rate limiting")
}
```

## Common Use Cases

### Managing Ad Campaigns

```go
func manageCampaigns(client *core.SDKClient, token, advertiserID string) error {
    // Get existing campaigns
    getCampaignReq := &ad.GetRequest{
        AdvertiserID: advertiserID,
        Filtering: &model.AdFiltering{
            Status: []string{"ALL"},
        },
    }
    
    campaignResp, err := ad.Get(context.Background(), client, getCampaignReq, token)
    if err != nil {
        return fmt.Errorf("failed to get campaigns: %w", err)
    }
    
    // Print campaign information
    for _, campaign := range campaignResp.Data.List {
        fmt.Printf("Campaign ID: %s, Name: %s, Status: %s\n", 
                   campaign.AdID, campaign.AdName, campaign.Status)
    }
    
    return nil
}
```

### Creating New Ads

```go
func createAd(client *core.SDKClient, token, advertiserID string) error {
    req := &ad.CreateRequest{
        AdvertiserID: advertiserID,
        Data: []*model.AdCreateData{
            {
                CampaignID: "existing_campaign_id",
                AdGroupName: "New Ad Group",
                AdName: "New Ad Creative",
                OptimizationGoal: "LINK_CLICKS",
                BillingEvent: "CLICK_THROUGH",
                BidAmount: 10000, // Amount in smallest currency unit
                Targeting: &model.Targeting{
                    AgeGroups: []string{"AGE_18_24", "AGE_25_34"},
                    Genders: []string{"GENDER_ALL"},
                    Locations: []string{"US", "CA"}, // Country codes
                },
            },
        },
    }
    
    resp, err := ad.Create(context.Background(), client, req, token)
    if err != nil {
        return fmt.Errorf("failed to create ad: %w", err)
    }
    
    if resp.Code != 0 {
        return fmt.Errorf("API error creating ad: %s", resp.Message)
    }
    
    fmt.Printf("Successfully created ad with ID: %s\n", resp.Data.AdIDs[0])
    return nil
}
```

### Retrieving Performance Reports

```go
func getPerformanceReports(client *core.SDKClient, token, advertiserID string) error {
    req := &report.GetIntegratedGetRequest{
        AdvertiserID: advertiserID,
        StartDate: "2023-01-01",
        EndDate: "2023-01-31",
        Dimensions: []string{"campaign_id", "date"},
        Metrics: []string{"spend", "impressions", "clicks", "reach"},
    }
    
    resp, err := report.GetIntegratedGet(context.Background(), client, req, token)
    if err != nil {
        return fmt.Errorf("failed to get reports: %w", err)
    }
    
    // Process report data
    for _, row := range resp.Data.List {
        fmt.Printf("Date: %s, Spend: %.2f, Impressions: %d, Clicks: %d\n",
                   row.Date, row.Spend, row.Impressions, row.Clicks)
    }
    
    return nil
}
```

## Troubleshooting

### Common Issues

#### 1. Invalid Access Token

**Symptoms**: API returns error code 40001 or similar authentication error
**Solution**: 
- Verify the access token hasn't expired
- Refresh the token if necessary
- Check that the token has the required scopes

#### 2. Insufficient Permissions

**Symptoms**: API returns permission-related error
**Solution**:
- Verify your app has the required permissions
- Check that the user granted all necessary scopes during OAuth
- Ensure the advertiser ID belongs to the authorized account

#### 3. Rate Limiting

**Symptoms**: API returns rate limit error (often code 40003)
**Solution**:
- Implement exponential backoff in your code
- Reduce the frequency of API calls
- Spread calls across time windows

#### 4. Parameter Validation Errors

**Symptoms**: API returns validation errors
**Solution**:
- Verify all required parameters are provided
- Check that parameter values meet API requirements
- Refer to official TikTok Business API documentation for valid values

### Debugging Tips

Enable debug mode to see detailed request/response information:

```go
client.EnableDebug()
```

Log request IDs for troubleshooting:

```go
if resp.RequestID != "" {
    log.Printf("Request ID: %s", resp.RequestID)
}
```

### Testing in Sandbox

Use sandbox mode for development:

```go
client.UseSandbox() // Use sandbox environment
```

## Resources

### Official Documentation

- [TikTok Business API Documentation](https://business-api.tiktok.com/)
- [TikTok for Business Developer Portal](https://developers.tiktok.com/)

### Support Channels

- GitHub Issues: [bububa/tiktok-business/issues](https://github.com/bububa/tiktok-business/issues)
- Stack Overflow: Tag questions with `tiktok-business-api`

### Community

- Join developer communities discussing TikTok advertising
- Participate in forums and Q&A sites
- Connect with other developers using the SDK

This user guide should help you effectively use the TikTok Business API Golang SDK in your applications. For additional support, please refer to the official TikTok Business API documentation or open an issue in the SDK repository.