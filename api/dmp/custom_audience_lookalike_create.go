package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// CustomAudienceLookalikeCreate 创建相似受众
// 您可以使用本接口创建相似受众。您需要使用广告主账号下已有的受众人群创建相似受众。
func CustomAudienceLookalikeCreate(ctx context.Context, clt *core.SDKClient, req *dmp.CustomAudienceLookalikeCreateRequest, accessToken string) (string, error) {
	var resp dmp.CustomAudienceLookalikeCreateResponse
	if err := clt.Post(ctx, "v1.3/dmp/custom_audience/lookalike/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.CustomAudienceID, nil
}
