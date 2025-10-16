package store

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/store"
)

// ProductGet 获取 TikTok Shop 中的商品
func ProductGet(ctx context.Context, clt *core.SDKClient, req *store.ProductGetRequest, accessToken string) (*store.ProductGetResult, error) {
	var ret store.ProductGetResponse
	if err := clt.Get(ctx, "v1.3/store/product/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
