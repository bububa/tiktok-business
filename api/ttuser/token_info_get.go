package ttuser

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ttuser"
)

// TokenInfoGet 获取已授权的创作者权限
// 您可以使用本接口获取已加入TTCM (TikTok Creator Marketplace) 的创作者授予的 TikTok 创作者账户（此时为 TikTok 个人账号）权限范围信息，或获取企业号用户授予的企业号的权限范围信息。
// 若您想获取 TikTok 用户授予的 TikTok 账号的权限范围信息，请查看这里。
func TokenInfoGet(ctx context.Context, clt *core.SDKClient, accessToken string) (*ttuser.TokenInfo, error) {
	req := ttuser.TokenInfoGetRequest{
		AppID:       clt.AppID(),
		AccessToken: accessToken,
	}
	var resp ttuser.TokenInfoGetResponse
	if err := clt.Post(ctx, "v1.3/tt_user/token_info/get/", &req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
