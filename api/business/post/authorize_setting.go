package post

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/post"
)

// AuthorizeSetting 启用或关闭 TikTok 帖子的广告授权设置
// 您可以使用本接口启用或关闭 TikTok 帖子的“广告授权”设置。
// 若想了解管理 TikTok 帖子“广告授权”设置的详细步骤，可查看管理 TikTok 帖子的广告授权。
func AuthorizeSetting(ctx context.Context, clt *core.SDKClient, req *post.AuthorizeSettingRequest, accessToken string) (*post.AuthorizeStatus, error) {
	var resp post.AuthorizeStatusResponse
	if err := clt.Post(ctx, "v1.3/business/post/authorize/setting", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
