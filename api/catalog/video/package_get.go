package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/video"
)

// PackageGet 获取视频包
func PackageGet(ctx context.Context, clt *core.SDKClient, req *video.PackageGetRequest, accessToken string) (*video.Package, error) {
	var resp video.PackageGetResponse
	if err := clt.Get(ctx, "v1.3/catalog/video_package/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
