package post

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/post"
)

// AuthorizeStatus 获取 TikTok 帖子的授权状态
func AuthorizeStatus(ctx context.Context, clt *core.SDKClient, req *post.AuthorizeStatusRequest, accessToken string) (*post.AuthorizeStatus, error) {
	var resp post.AuthorizeStatusResponse
	if err := clt.Get(ctx, "v1.3/business/post/authorize/status", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
