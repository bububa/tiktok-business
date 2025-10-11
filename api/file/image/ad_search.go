package image

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file/image"
)

// AdSearch 搜索图片
func AdSearch(ctx context.Context, clt *core.SDKClient, req *image.AdSearchRequest, accessToken string) (*image.AdSearchResult, error) {
	var ret image.AdSearchResponse
	if err := clt.Get(ctx, "v1.3/file/image/ad/search/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
