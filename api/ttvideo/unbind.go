package ttvideo

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ttvideo"
)

// Unbind 解除 Spark Ads 帖子绑定
// 若 Spark Ads 帖子的授权已过期或已取消，您可以使用本接口对帖子进行解绑。解绑 Spark Ads 帖子意味着该帖子从广告账户的已授权 Spark Ads 帖子列表中移除。
func Unbind(ctx context.Context, clt *core.SDKClient, req *ttvideo.UnbindRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/tt_video/unbind/", req, nil, accessToken)
}
