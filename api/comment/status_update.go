package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/comment"
)

// StatusUpdate 更新评论状态
// 您可以使用本接口将一个或一组评论设为隐藏或可见
func StatusUpdate(ctx context.Context, clt *core.SDKClient, req *comment.StatusUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/comment/status/update/", req, nil, accessToken)
}
