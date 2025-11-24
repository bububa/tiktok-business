package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/comment"
)

// ReplyCreate 创建对自有视频评论的回复
// 您可以使用本接口为自有TikTok 账号所发布原生视频的现有评论创建回复。
func ReplyCreate(ctx context.Context, clt *core.SDKClient, req *comment.ReplyCreateRequest, accessToken string) (*comment.Comment, error) {
	var resp comment.CreateResponse
	if err := clt.Post(ctx, "v1.3/business/comment/reply/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
