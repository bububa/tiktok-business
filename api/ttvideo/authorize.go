package ttvideo

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ttvideo"
)

// Authorize 应用授权码
// 您可以使用本接口应用授权码。从视频主获得授权码后，授权并不会立即生效。您需要使用授权码将广告账户与Spark Ads帖子相连接。
// 对于拼接或者合拍帖子或者源视频提及了其他视频，广告主在推广帖子前还需要获得源视频的授权。在本接口中，你需要使用auth_code来传入当前视频的授权码，并用original_post_auth_code来传入源视频的授权码。
func Authorize(ctx context.Context, clt *core.SDKClient, req *ttvideo.AuthorizeRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/tt_video/authorize/", req, nil, accessToken)
}
