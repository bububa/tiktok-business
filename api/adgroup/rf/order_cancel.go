package rf

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup/rf"
)

// OrderCancel 撤销覆盖和频次广告订单
// 您可以使用本接口撤单，从而停止您的覆盖和频次广告投放流程。
// 撤单成功后，对应广告组的状态会变为已关闭。已进入投放中状态的广告组不可撤单。
func OrderCancel(ctx context.Context, clt *core.SDKClient, req *rf.OrderCancelRequest, accessToken string) ([]string, error) {
	var resp rf.OrderCancelResponse
	if err := clt.Post(ctx, "v1.3/rf/order/cancel/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.FailAdgroupIDs, nil
}
