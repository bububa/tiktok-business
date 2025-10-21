package diagnostic

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/diagnostic"
)

// ProductTaskCreate 创建商品库商品诊断信息的异步下载任务
// 您可以使用本接口创建商品库商品诊断信息的异步下载任务。
// 异步下载任务创建完成后，您可以使用/diagnostic/catalog/product/task/get/ 下载生成的 CSV 文件。需注意与异步下载任务提供的诊断信息相比，/diagnostic/catalog/接口同步获取的诊断信息更为详细。异步下载任务仅返回每个检测到的问题的 Product ID（商品ID，对应 product_id ），SKU ID（ 自定义的商品唯一ID，对应sku_id / hotel_id / flight_id / destination_id ) ，Item title（商品标题，对应 title ），Issue title（问题标题，对应 issue_title ），Severity（问题级别，对应 issue_level ）和 Source（更新源名称）。
func ProductTaskCreate(ctx context.Context, clt *core.SDKClient, req *diagnostic.ProductTaskCreateRequest, accessToken string) (string, error) {
	var resp diagnostic.ProductTaskCreateResponse
	if err := clt.Post(ctx, "v1.3/diagnostic/catalog/product/task/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.TaskID, nil
}
