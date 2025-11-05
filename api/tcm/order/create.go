package order

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/order"
)

// Create 创建 TTCM 订单
// 您可以使用本接口为创作者创建TTCM (TikTok Creator Marketplace)订单，无论创作者是否在TTCM平台注册均适用。
// 您只需传入基本的订单和创作者信息（tcm_account_id，brand_name，description，handle_names），随后本接口会自动生成TTCM推广系列代码，邀请链接及多个订单ID，能够使创作者为同一推广系列发布多个视频。
func Create(ctx context.Context, clt *core.SDKClient, req *order.CreateRequest, accessToken string) (*order.CreateResult, error) {
	var resp order.CreateResponse
	if err := clt.Post(ctx, "v1.3/tcm/order/create/v2/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
