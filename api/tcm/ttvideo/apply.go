package ttvideo

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/ttvideo"
)

// Apply 申请 Spark Ads 授权
// 您可以使用本接口发送申请，请求 TikTok 创作者授权合作伙伴为现有订单推广其视频。订单可为推广系列代码订单或非推广系列代码订单。
// 请注意，若视频要在非推广系列代码订单中使用，您可以继续使用/tcm/spark_ad/apply/申请Spark Ads授权码，您需传入order_id 而非video_id。
func Apply(ctx context.Context, clt *core.SDKClient, req *ttvideo.ApplyRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/tcm/tt_video/apply/", req, nil, accessToken)
}
