package identity

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/identity"
)

// Get 获取认证身份列表
// 您可以使用本接口获取广告账户下的认证身份列表。您可以根据身份类型或显示名称筛选结果
func Get(ctx context.Context, clt *core.SDKClient, req *identity.GetRequest, accessToken string) (*identity.GetResult, error) {
	var resp identity.GetResponse
	if err := clt.Get(ctx, "v1.3/identity/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
