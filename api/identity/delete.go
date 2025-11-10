package identity

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/identity"
)

// Delete 删除认证身份
// 您可以使用本接口删除自定义用户认证身份（identity_type = CUSTOMIZED_USER）。如果您想要编辑认证身份，需要先使用本接口删除该身份，再使用 /v1.3/identity/create/ 创建一个新的认证身份
func Delete(ctx context.Context, clt *core.SDKClient, req *identity.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/identity/delete/", req, nil, accessToken)
}
