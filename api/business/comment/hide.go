package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/comment"
)

// Hide 隐藏/取消隐藏自有视频的评论
// 您可以使用本接口隐藏或取消隐藏自有TikTok 账号所发布原生视频的现有评论。
func Hide(ctx context.Context, clt *core.SDKClient, req *comment.HideRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/business/comment/hide/", req, nil, accessToken)
}
