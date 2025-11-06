package creator

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/creator"
)

// Invite 邀请创作者加入 TTCM
// 您可以使用本接口按用户名将 TikTok Creator Marketplace 入驻邀请发送给列表中的 TikTok 创作者。
func Invite(ctx context.Context, clt *core.SDKClient, req *creator.InviteRequest, accessToken string) (*creator.InviteResult, error) {
	var resp creator.InviteResponse
	if err := clt.Post(ctx, "v1.3/tcm/creator/invite/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
