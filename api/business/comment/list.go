package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/comment"
)

// List 获取自有视频的评论
// 您可以使用本接口获取自有 TikTok 账号所发布的某一原生视频的所有或指定的相关公开和隐藏评论（以及相关信息）。
// 您可根据是否返回对应的parent_comment_id判断comment_id对应的是评论还是回复。仅回复会返回对应的父层级评论（parent_comment_id）
func List(ctx context.Context, clt *core.SDKClient, req *comment.ListRequest, accessToken string) (*comment.ListResult, error) {
	var resp comment.ListResponse
	if err := clt.Get(ctx, "v1.3/business/comment/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
