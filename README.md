# TikTok Business API Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/tiktok-business.svg)](https://pkg.go.dev/github.com/bububa/tiktok-business)
[![Go](https://github.com/bububa/tiktok-business/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/tiktok-business/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/tiktok-business/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/tiktok-business/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/tiktok-business.svg)](https://github.com/bububa/tiktok-business)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/tiktok-business)](https://goreportcard.com/report/github.com/bububa/tiktok-business)
[![GitHub license](https://img.shields.io/github/license/bububa/tiktok-business.svg)](https://github.com/bububa/tiktok-business/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/tiktok-business.svg)](https://GitHub.com/bububa/tiktok-business/releases/)

Golang SDK for TikTok Business API. This SDK provides a complete implementation of the TikTok Business API, allowing you to interact with TikTok's advertising and business platform programmatically.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Usage Examples](#usage-examples)
- [API Endpoints](#api-endpoints)
- [Authentication](#authentication)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Installation

Install the SDK using Go modules:

```bash
go get github.com/bububa/tiktok-business
```

## Quick Start

Here's a quick example to get you started with the SDK:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/bububa/tiktok-business/core"
    "github.com/bububa/tiktok-business/api/ad"
    "github.com/bububa/tiktok-business/model/ad"
)

func main() {
    // Initialize the client with your App ID and Secret
    client := core.NewSDKClient("your_app_id", "your_app_secret")

    // Create a request to get ads
    req := &ad.GetRequest{
        AdvertiserID: "your_advertiser_id",
        Filtering: &model.AdFiltering{
            Status: []string{"ALL"},
        },
        Page:     1,
        PageSize: 10,
    }

    // Execute the request
    resp, err := ad.Get(context.Background(), client, req, "your_access_token")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Retrieved %d ads\n", len(resp.Data.List))
    for _, adInfo := range resp.Data.List {
        fmt.Printf("Ad ID: %s, Name: %s, Status: %s\n", 
                   adInfo.AdID, adInfo.AdName, adInfo.Status)
    }
}
```

## Configuration

### Client Initialization

Create a new SDK client with your application credentials:

```go
client := core.NewSDKClient("your_app_id", "your_app_secret")
```

### Sandbox Mode

For testing purposes, you can enable sandbox mode:

```go
client := core.NewSDKClient("your_app_id", "your_app_secret")
client.UseSandbox()
```

### Debug Mode

Enable debug mode to see detailed request/response logs:

```go
client := core.NewSDKClient("your_app_id", "your_app_secret")
client.EnableDebug()
```

## Usage Examples

### Getting Ads

```go
req := &ad.GetRequest{
    AdvertiserID: "your_advertiser_id",
    Filtering: &model.AdFiltering{
        Status: []string{"ALL"},
    },
    Page:     1,
    PageSize: 10,
}

resp, err := ad.Get(context.Background(), client, req, "your_access_token")
if err != nil {
    log.Fatal(err)
}

for _, adInfo := range resp.Data.List {
    fmt.Printf("Ad: %s - %s\n", adInfo.AdID, adInfo.AdName)
}
```

### Creating an Ad

```go
req := &ad.CreateRequest{
    AdvertiserID: "your_advertiser_id",
    Data: []*model.AdCreateData{
        {
            CampaignID:    "campaign_id",
            AdGroupName:   "New Ad Group",
            AdName:        "New Ad",
            OptimizationGoal: "LINK_CLICKS",
            BillingEvent:  "CLICK_THROUGH",
            BidAmount:     10000, // Amount in cents
            Targeting: &model.Targeting{
                AgeGroups: []string{"AGE_18_24", "AGE_25_34"},
                Genders:   []string{"GENDER_MALE", "GENDER_FEMALE"},
            },
        },
    },
}

resp, err := ad.Create(context.Background(), client, req, "your_access_token")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Created ad with ID: %s\n", resp.Data.AdID)
```

### Updating an Ad

```go
req := &ad.UpdateRequest{
    AdvertiserID: "your_advertiser_id",
    Data: []*model.AdUpdateData{
        {
            AdID:   "ad_id_to_update",
            Status: "PAUSED",
        },
    },
}

resp, err := ad.Update(context.Background(), client, req, "your_access_token")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Updated ad: %s\n", resp.Message)
```

## API Endpoints

The SDK covers various TikTok Business API endpoints organized in different categories:

### Advertising
- `ad` - Manage ad campaigns, ad groups, and ads
- `ad/aco` - Advanced campaign optimization features
- `report` - Access advertising performance reports
- `pangle` - Pangle advertising platform APIs

### Catalog Management
- `catalog` - Manage product catalogs
- `catalog/product` - Product management
- `catalog/feed` - Catalog feeds
- `catalog/set` - Product sets
- `catalog/video` - Video management for catalogs

### Events & Tracking
- `event` - Track conversion events
- `pixel` - Pixel tracking management

### User & Showcase
- `user` - User information
- `showcase` - Showcase management
- `ttuser` - TikTok user OAuth

### Search Ads
- `searchad/negativekeyword` - Negative keyword management

### Additional Features
- `report` - Various reporting endpoints

## Authentication

The SDK requires an access token for most API calls. You can obtain access tokens through TikTok's OAuth flow.

For OAuth token management:

```go
// Generate OAuth URL for user authorization
oauthURL := ttuser.OAuthURL("your_app_id", "redirect_uri", "state", []string{"ads_read", "ads_manage"})

// Exchange authorization code for access token
tokenResp, err := ttuser.Token(context.Background(), client, &ttuser.TokenRequest{
    AppID:          "your_app_id",
    Secret:         "your_app_secret",
    AuthCode:       "authorization_code",
    GrantType:      "authorization_code",
    RedirectURI:    "redirect_uri",
}, "")

if err != nil {
    log.Fatal(err)
}

accessToken := tokenResp.AccessToken
```

## Testing

Run the complete test suite:

```bash
go test -v ./...
```

Run tests with coverage:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Run tests with race detection:

```bash
go test -race -v ./...
```

## Contributing

We welcome contributions! Please see our [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute to this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
