package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// MinisGet 获取广告账号内的 TikTok Minis。
func MinisGet(ctx context.Context, clt *core.SDKClient, req *tool.MinisGetRequest, accessToken string) (*tool.MinisGetResult, error) {
	var resp tool.MinisGetResponse
	if err := clt.Get(ctx, "v1.3/minis/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
