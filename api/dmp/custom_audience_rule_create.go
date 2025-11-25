package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// CustomAudienceRuleCreate 通过规则创建受众
// 您可以使用本接口通过规则创建受众。
// 您需要传入一系列的包含规则及可选的排除规则来创建受众。
func CustomAudienceRuleCreate(ctx context.Context, clt *core.SDKClient, req *dmp.CustomAudienceRuleCreateRequest, accessToken string) (string, error) {
	var resp dmp.CustomAudienceRuleCreateResponse
	if err := clt.Post(ctx, "v1.3/dmp/custom_audience/rule/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.CustomAudienceID, nil
}
