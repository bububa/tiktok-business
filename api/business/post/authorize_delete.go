package post

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/post"
)

// AuthorizeDelete 删除 TikTok 帖子的授权码
func AuthorizeDelete(ctx context.Context, clt *core.SDKClient, req *post.AuthorizeDeleteRequest, accessToken string) (*post.AuthorizeStatus, error) {
	var resp post.AuthorizeStatusResponse
	if err := clt.Post(ctx, "v1.3/business/post/authorize/delete", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
