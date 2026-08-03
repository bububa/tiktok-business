package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// AvailableAttributionSource 获取应用可用的归因来源和数据源。
func AvailableAttributionSource(ctx context.Context, clt *core.SDKClient, req *tool.AvailableAttributionSourceRequest, accessToken string) (*tool.AvailableAttributionSourceResult, error) {
	var resp tool.AvailableAttributionSourceResponse
	if err := clt.Get(ctx, "v1.3/tool/available/attribution_source/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
