package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/comment"
)

// ReplyList 获取评论的所有回复
// 若自有TikTok 账号某原生视频的某条评论已有回复，您可以使用本接口查看所有公开和隐藏的回复（以及相关信息）
func ReplyList(ctx context.Context, clt *core.SDKClient, req *comment.ReplyListRequest, accessToken string) (*comment.ListResult, error) {
	var resp comment.ListResponse
	if err := clt.Get(ctx, "v1.3/business/comment/reply/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
