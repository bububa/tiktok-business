package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// CustomAudienceShareLog 获取受众分享记录
// 您可以使用本接口获取某一受众的分享记录。分享记录包含曾经分享过受众的广告主ID和名称。
func CustomAudienceShareLog(ctx context.Context, clt *core.SDKClient, req *dmp.CustomAudienceShareLogRequest, accessToken string) ([]dmp.CustomAudienceShareLog, error) {
	var resp dmp.CustomAudienceShareLogResponse
	if err := clt.Get(ctx, "v1.3/dmp/custom_audience/share/log/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
