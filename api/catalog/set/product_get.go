package set

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/set"
)

// ProductGet 获取商品信息
// 您可以使用本接口获取指定商品系列下的商品信息
func ProductGet(ctx context.Context, clt *core.SDKClient, req *set.ProductGetRequest, accessToken string) (*set.ProductGetResult, error) {
	var resp set.ProductGetResponse
	if err := clt.Get(ctx, "v1.3/catalog/set/product/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
