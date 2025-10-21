package product

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/product"
)

// Get 获取商品信息
func Get(ctx context.Context, clt *core.SDKClient, req *product.GetRequest, accessToken string) (*product.GetResult, error) {
	var resp product.GetResponse
	if err := clt.Get(ctx, "v1.3/catalog/product/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
