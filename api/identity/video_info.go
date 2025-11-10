package identity

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/identity"
)

// VideoInfo 获取 TikTok 帖子信息
// 您可以使用本接口获取通过AUTH_CODE，TT_USER或BC_AUTH_TT类型的认证身份下的 TikTok 帖子的信息
func VideoInfo(ctx context.Context, clt *core.SDKClient, req *identity.VideoInfoRequest, accessToken string) (*identity.VideoInfoResult, error) {
	var resp identity.VideoInfoResponse
	if err := clt.Get(ctx, "v1.3/identity/video/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
