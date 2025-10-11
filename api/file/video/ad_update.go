package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file/video"
)

// AdUpdate 更新视频名称
func AdUpdate(ctx context.Context, clt *core.SDKClient, req *video.AdUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/file/video/ad/update/", req, nil, accessToken)
}
