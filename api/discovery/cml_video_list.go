package discovery

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/discovery"
)

// CMLVideoList Get trending videos related to tracks
// Use this endpoint to retrieve the top trending videos that use a popular track in a selected region.
func CMLVideoList(ctx context.Context, clt *core.SDKClient, req *discovery.CMLVideoListRequest, accessToken string) (*discovery.CommercialMusicLibrary, error) {
	var resp discovery.CMLVideoListResponse
	if err := clt.Get(ctx, "v1.3/discovery/cml/video_list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
