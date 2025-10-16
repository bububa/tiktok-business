package store

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/store"
)

// List 获取广告账户下可用的商店
// 您可以使用本接口获取广告账户下可用的第一方商店（TikTok Shop）列表。
func List(ctx context.Context, clt *core.SDKClient, req *store.ListRequest, accessToken string) ([]store.Store, error) {
	var ret store.ListResponse
	if err := clt.Get(ctx, "v1.3/store/list/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.Stores, nil
}
