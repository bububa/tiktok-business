package ttvideo

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/ttvideo"
)

// Status 获取授权状态
// 您可以使用本接口获取现有订单的 Spark Ads 授权状态。
func Status(ctx context.Context, clt *core.SDKClient, req *ttvideo.StatusRequest, accessToken string) (*ttvideo.Video, error) {
	var resp ttvideo.StatusResponse
	if err := clt.Get(ctx, "v1.3/tcm/tt_video/status/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
