package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/video"
)

// Delete 删除上传的商品库视频
func Delete(ctx context.Context, clt *core.SDKClient, req *video.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/catalog/video/delete/", req, nil, accessToken)
}
