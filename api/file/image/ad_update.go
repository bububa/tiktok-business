package image

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file/image"
)

// AdUpdate 更新图片名称
func AdUpdate(ctx context.Context, clt *core.SDKClient, req *image.AdUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/file/image/ad/update/", req, nil, accessToken)
}
