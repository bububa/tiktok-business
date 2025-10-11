package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file/video"
)

// AdSearch 搜索视频
func AdSearch(ctx context.Context, clt *core.SDKClient, req *video.AdSearchRequest, accessToken string) (*video.AdSearchResult, error) {
	var ret video.AdSearchResponse
	if err := clt.Get(ctx, "v1.3/file/video/ad/search/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
