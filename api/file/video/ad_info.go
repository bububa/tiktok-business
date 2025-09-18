package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file/video"
)

// AdInfo 获取视频信息
func AdInfo(ctx context.Context, clt *core.SDKClient, req *video.AdInfoRequest, accessToken string) ([]video.Video, error) {
	var ret video.AdInfoResponse
	if err := clt.Get(ctx, "v1.3/file/video/ad/info", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.List, nil
}
