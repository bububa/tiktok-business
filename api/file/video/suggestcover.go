package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file/video"
)

// SuggestCover 获取视频推荐封面
func SuggestCover(ctx context.Context, clt *core.SDKClient, req *video.SuggestCoverRequest, accessToken string) ([]video.SuggestCover, error) {
	var ret video.SuggestCoverResponse
	if err := clt.Get(ctx, "v1.3/file/video/suggestcover", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.List, nil
}
