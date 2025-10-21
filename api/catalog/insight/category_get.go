package insight

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/insight"
)

// CategoryGet 获取热门商品类目列表
// 您可以使用本接口获取您的电商商品库中与基于 TikTok 用户互动数据的前 50 个热门商品类目相匹配的商品数量，以及这些商品的可售库存百分比。
func CategoryGet(ctx context.Context, clt *core.SDKClient, req *insight.CategoryGetRequest, accessToken string) ([]insight.Category, error) {
	var resp insight.CategoryGetResponse
	if err := clt.Get(ctx, "v1.3/catalog/insight/category/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.CategoryInsights, nil
}
