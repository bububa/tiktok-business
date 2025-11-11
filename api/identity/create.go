package identity

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/identity"
)

// Create 创建认证身份
// 您可以使用本接口创建自定义用户 (CUSTOMIZED_USER)类型的认证身份
func Create(ctx context.Context, clt *core.SDKClient, req *identity.CreateRequest, accessToken string) (string, error) {
	var resp identity.CreateResponse
	if err := clt.Post(ctx, "v1.3/identity/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.IdentityID, nil
}
