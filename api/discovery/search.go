package discovery

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/discovery"
)

// Search Get trending search keywords
// Use this endpoint to retrieve the top 20 search keywords trending on TikTok or search keywords personalized for your Business Account.
func Search(ctx context.Context, clt *core.SDKClient, req *discovery.SearchRequest, accessToken string) ([]string, error) {
	var resp discovery.SearchResponse
	if err := clt.Get(ctx, "v1.3/discovery/search/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.SearchKeywords, nil
}
