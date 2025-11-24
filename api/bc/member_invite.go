package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// MemberInvite 邀请成员
// 商务中心的管理员用户可以使用本接口邀请成员加入商务中心，并给成员分配权限和资产。
// 每个商务中心最多支持4000个成员，最多支持20个管理员用户同时在线。
func MemberInvite(ctx context.Context, clt *core.SDKClient, req *bc.MemberInviteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/member/invite/", req, nil, accessToken)
}
