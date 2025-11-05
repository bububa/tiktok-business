package showcase

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/showcase"
)

// IdentityGet 获取广告账户下有橱窗权限的认证身份
func IdentityGet(ctx context.Context, clt *core.SDKClient, req *showcase.IdentityGetRequest, accessToken string) ([]showcase.Identity, error) {
	var ret showcase.IdentityGetResponse
	if err := clt.Get(ctx, "v1.3/showcase/identity/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.IdentityList, nil
}
