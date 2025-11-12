package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// MemberUpdate 修改成员信息
// 商务中心的管理员用户可以使用本接口修改成员的用户名以及他们在商务中心的角色
func MemberUpdate(ctx context.Context, clt *core.SDKClient, req *bc.MemberUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/member/update/", req, nil, accessToken)
}
