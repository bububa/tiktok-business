package customaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/customaudience"
)

// RuleCreate 通过规则创建受众
// 您可以使用本接口通过规则创建受众。
// 您需要传入一系列的包含规则及可选的排除规则来创建受众。
func RuleCreate(ctx context.Context, clt *core.SDKClient, req *customaudience.RuleCreateRequest, accessToken string) (string, error) {
	var resp customaudience.RuleCreateResponse
	if err := clt.Post(ctx, "v1.3/dmp/custom_audience/rule/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.CustomAudienceID, nil
}
