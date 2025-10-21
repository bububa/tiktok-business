package product

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/product"
)

// Update 更新商品
// 您可以使用本接口批量更新商品库商品信息（单次最多可更新5000个商品），系统将会异步录入请求中的商品。您将会从返回中得到一个任务处理IDfeed_log_id。通过该ID您可以调用/catalog/product/log/查看任务的处理状态和最终结果。
// 和/catalog/product/file/接口不同，该接口主要用于少量商品上传及部分字段更新。
func Update(ctx context.Context, clt *core.SDKClient, req *product.UpdateRequest, accessToken string) (string, error) {
	var resp product.UpdateResponse
	if err := clt.Post(ctx, "v1.3/catalog/product/update/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.FeedLogID, nil
}
