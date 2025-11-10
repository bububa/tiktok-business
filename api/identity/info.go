package identity

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/identity"
)

// Info 获取认证身份详情
func Info(ctx context.Context, clt *core.SDKClient, req *identity.InfoRequest, accessToken string) (*identity.Identity, error) {
	var resp identity.InfoResponse
	if err := clt.Get(ctx, "v1.3/identity/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.IdentityInfo, nil
}
