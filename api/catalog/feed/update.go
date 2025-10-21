package feed

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/feed"
)

// Update 修改更新源
func Update(ctx context.Context, clt *core.SDKClient, req *feed.UpdateRequest, accessToken string) (*feed.Feed, error) {
	var ret feed.UpdateResponse
	if err := clt.Post(ctx, "v1.3/catalog/feed/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
