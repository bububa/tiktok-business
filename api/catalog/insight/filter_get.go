package insight

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/insight"
)

// FilterGet 获取商品库分析数据的筛选条件
// 您可以使用本接口获取商品库分析数据的可用筛选条件，这些条件可筛选电商商品库中的商品，生成更具针对性的分析数据。
// 返回的筛选值可与 /catalog/insight/product/get/ 结合使用，以获取商品库中热门商品的分析数据。
func FilterGet(ctx context.Context, clt *core.SDKClient, req *insight.FilterGetRequest, accessToken string) (*insight.FilterGetResult, error) {
	var resp insight.FilterGetResponse
	if err := clt.Get(ctx, "v1.3/catalog/insight/filter/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
