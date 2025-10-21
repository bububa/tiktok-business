package diagnostic

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/diagnostic"
)

// ProductTaskGet 异步下载商品库商品诊断信息
// 您可以使用本接口下载包含商品库商品诊断信息的 CSV 文件。
// 需注意， CSV 文件仅返回每个检测到的问题的 Product ID（商品ID，对应 product_id ），SKU ID（ 自定义的商品唯一ID，对应sku_id / hotel_id / flight_id / destination_id ) ，Item title（商品标题，对应 title ），Issue title（问题标题，对应 issue_title ），Severity（问题级别，对应 issue_level ）和 Source（更新源名称）。若想获取更为详细的诊断信息，使用/diagnostic/catalog/接口。
func ProductTaskGet(ctx context.Context, clt *core.SDKClient, req *diagnostic.ProductTaskGetRequest, accessToken string) (*diagnostic.ProductTaskGetResult, error) {
	var resp diagnostic.ProductTaskGetResponse
	if err := clt.Get(ctx, "v1.3/diagnostic/catalog/product/task/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
