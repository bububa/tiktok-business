package post

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/post"
)

// Authorize 延长 TikTok 帖子的授权有效期
func Authorize(ctx context.Context, clt *core.SDKClient, req *post.AuthorizeRequest, accessToken string) (*post.AuthorizeStatus, error) {
	var resp post.AuthorizeStatusResponse
	if err := clt.Post(ctx, "v1.3/business/post/authorize/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
