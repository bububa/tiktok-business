package discovery

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/discovery"
)

// TrendingList Get popular hashtags
// Use this endpoint to get a list of the top 200 popular hashtags on TikTok.
func TrendingList(ctx context.Context, clt *core.SDKClient, req *discovery.TrendingListRequest, accessToken string) (*discovery.TrendingListResult, error) {
	var resp discovery.TrendingListResponse
	if err := clt.Get(ctx, "v1.3/discovery/trending_list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
