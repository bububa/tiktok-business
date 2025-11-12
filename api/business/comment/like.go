package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/comment"
)

// Like 点赞/取消点赞自有视频的评论
// 您可以使用本接口点赞或取消点赞自有TikTok 账号所发布原生视频的某条现有评论
func Like(ctx context.Context, clt *core.SDKClient, req *comment.LikeRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/business/comment/like/", req, nil, accessToken)
}
