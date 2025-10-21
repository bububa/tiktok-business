package feed

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/feed"
)

// Get 获取更新源
// 您可以使用本接口获取商品库的单个更新源或所有更新源的信息。
func Get(ctx context.Context, clt *core.SDKClient, req *feed.GetRequest, accessToken string) ([]feed.Feed, error) {
	var ret feed.GetResponse
	if err := clt.Get(ctx, "v1.3/catalog/feed/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.FeedList, nil
}
