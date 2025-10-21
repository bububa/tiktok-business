package diagnostic

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/diagnostic"
)

// Catalog 同步获取商品库商品诊断信息
// 您可以使用本接口同步获取您的商品库商品的诊断信息。
// - 若商品库中发现问题，则当前能够显示的最新诊断信息为前一天的诊断信息，因为商品库诊断数据按UTC+0 时间存在一天的延迟。
// - 若商品库中未发现问题，则返回的商品库诊断信息为空。
//
// 若想创建商品库商品诊断信息的异步下载任务，需使用/diagnostic/catalog/product/task/create/。需注意与异步下载任务提供的诊断信息相比，本接口同步获取的诊断信息更为详细。异步下载任务仅返回每个检测到的问题的 Product ID（商品ID，对应 product_id ），SKU ID（ 自定义的商品唯一ID，对应sku_id / hotel_id / flight_id / destination_id ) ，Item title（商品标题，对应 title ），Issue title（问题标题，对应 issue_title ），Severity（问题级别，对应 issue_level ）和 Source（更新源名称）。
func Catalog(ctx context.Context, clt *core.SDKClient, req *diagnostic.CatalogRequest, accessToken string) (*diagnostic.CatalogResult, error) {
	var resp diagnostic.CatalogResponse
	if err := clt.Get(ctx, "v1.3/diagnostic/catalog/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
