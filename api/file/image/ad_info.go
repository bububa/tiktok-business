package image

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file/image"
)

// AdInfo 获取图片信息
func AdInfo(ctx context.Context, clt *core.SDKClient, req *image.AdInfoRequest, accessToken string) ([]image.Image, error) {
	var ret image.AdInfoResponse
	if err := clt.Get(ctx, "v1.3/file/image/ad/info", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.List, nil
}
