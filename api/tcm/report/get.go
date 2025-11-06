package report

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/report"
)

// Get 请求 TTCM 订单表现报表
// 您可以使用本接口请求通过/tcm/order/create/v2/创建的TTCM (TikTok Creator Marketplace)订单的报表数据。通过/tcm/order/create/v2/创建的订单可供已在TTCM平台注册或未注册的创作者使用。
func Get(ctx context.Context, clt *core.SDKClient, req *report.GetRequest, accessToken string) (*report.GetResult, error) {
	var resp report.GetResponse
	if err := clt.Get(ctx, "v1.3/tcm/report/get/v2/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
