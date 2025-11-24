package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// MemberDelete 删除成员或撤销成员邀请
// 商务中心的管理员用户可以使用本接口在商务中心删除成员或撤销成员邀请。
func MemberDelete(ctx context.Context, clt *core.SDKClient, req *bc.MemberDeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/member/delete/", req, nil, accessToken)
}
