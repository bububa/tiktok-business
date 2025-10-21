package insight

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/insight"
)

// ProductGet 获取热门商品列表
// 您可以使用本接口获取电商商品库中基于 TikTok 用户互动数据的热门商品，至多可获取前 50 个热门商品。
func ProductGet(ctx context.Context, clt *core.SDKClient, req *insight.ProductGetRequest, accessToken string) ([]insight.Product, error) {
	var resp insight.ProductGetResponse
	if err := clt.Get(ctx, "v1.3/catalog/insight/product/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ProductInsights, nil
}
