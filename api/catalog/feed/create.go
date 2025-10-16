package feed

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/feed"
)

// Create 创建更新源
// 您可以使用本接口创建更新源。
func Create(ctx context.Context, clt *core.SDKClient, req *feed.CreateRequest, accessToken string) (*feed.Feed, error) {
	var ret feed.CreateResponse
	if err := clt.Post(ctx, "v1.3/catalog/feed/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
