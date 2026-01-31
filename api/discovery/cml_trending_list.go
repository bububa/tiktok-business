package discovery

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/discovery"
)

// CMLTrendingList Get popular tracks from the Commercial Music Library
// Use this endpoint to retrieve a list of up to 100 popular tracks that are pre-cleared for use in organic content and ad creation from the Commercial Music Library.
func CMLTrendingList(ctx context.Context, clt *core.SDKClient, req *discovery.CMLTrendingListRequest, accessToken string) ([]discovery.CommercialMusicLibrary, error) {
	var resp discovery.CMLTrendingListResponse
	if err := clt.Get(ctx, "v1.3/discovery/cml/trending_list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
