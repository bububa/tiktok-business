package rf

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup/rf"
)

// InventoryEstimate 获取覆盖和频次广告库存预估
// 您可以使用本接口获取覆盖和频次广告的库存预估。参考这里了解更多库存预估的细节，及该接口返回结果与TikTok广告管理平台预测结果的匹配关系。
func InventoryEstimate(ctx context.Context, clt *core.SDKClient, req *rf.InventoryEstimateRequest, accessToken string) (*rf.InventoryEstimateResult, error) {
	var resp rf.InventoryEstimateResponse
	if err := clt.Get(ctx, "v1.3/rf/inventory/estimate/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
