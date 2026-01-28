package discovery

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/discovery"
)

// SearchKeyword Get recommended search keywords
// Use this endpoint to retrieve a list of 20 recommended search keywords for a search query or to get recommended search keywords personalized for your Business Account.
func SearchKeyword(ctx context.Context, clt *core.SDKClient, req *discovery.SearchKeywordRequest, accessToken string) ([]string, error) {
	var resp discovery.SearchKeywordResponse
	if err := clt.Get(ctx, "v1.3/discovery/search/keyword/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.SearchKeywords, nil
}
