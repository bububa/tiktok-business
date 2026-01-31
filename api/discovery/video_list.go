package discovery

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/discovery"
)

// VideoList Get trending videos related to hashtags
// Use this endpoint to retrieve the top 20 trending videos that use a popular hashtag in the selected top country of the hashtag and timeframe.
func VideoList(ctx context.Context, clt *core.SDKClient, req *discovery.VideoListRequest, accessToken string) ([]discovery.Hashtag, error) {
	var resp discovery.VideoListResponse
	if err := clt.Get(ctx, "v1.3/discovery/video_list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
