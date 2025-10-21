package product

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/product"
)

// Log 获取商品操作结果
// 您可以使用本接口了解商品上传或删除是否成功，以及上传或删除失败的后续操作建议
func Log(ctx context.Context, clt *core.SDKClient, req *product.LogRequest, accessToken string) (*product.Log, error) {
	var resp product.LogResponse
	if err := clt.Get(ctx, "v1.3/catalog/product/log/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ProductFeedLog, nil
}
