package customaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/customaudience"
)

// LookalikeCreate 创建相似受众
// 您可以使用本接口创建相似受众。您需要使用广告主账号下已有的受众人群创建相似受众。
func LookalikeCreate(ctx context.Context, clt *core.SDKClient, req *customaudience.LookalikeCreateRequest, accessToken string) (string, error) {
	var resp customaudience.LookalikeCreateResponse
	if err := clt.Post(ctx, "v1.3/dmp/custom_audience/lookalike/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.CustomAudienceID, nil
}
