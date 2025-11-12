package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/comment"
)

// Create 为自有视频创建新评论
// 您可以使用本接口为自有TikTok 账号所发布原生视频创建新的评论
func Create(ctx context.Context, clt *core.SDKClient, req *comment.CreateRequest, accessToken string) (*comment.Comment, error) {
	var resp comment.CreateResponse
	if err := clt.Post(ctx, "v1.3/business/comment/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
