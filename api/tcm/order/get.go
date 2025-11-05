package order

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/order"
)

// Get 获取 TTCM 订单
func Get(ctx context.Context, clt *core.SDKClient, req *order.GetRequest, accessToken string) (*order.GetResult, error) {
	var resp order.GetResponse
	if err := clt.Get(ctx, "v1.3/tcm/order/get/v2/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
