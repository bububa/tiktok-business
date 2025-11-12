package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/comment"
)

// Delete 删除自有视频的评论
func Delete(ctx context.Context, clt *core.SDKClient, req *comment.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/business/comment/delete/", req, nil, accessToken)
}
