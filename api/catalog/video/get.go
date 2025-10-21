package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/video"
)

// Get 获取商品库中上传的商品库视频
func Get(ctx context.Context, clt *core.SDKClient, req *video.GetRequest, accessToken string) (*video.GetResult, error) {
	var resp video.GetResponse
	if err := clt.Get(ctx, "v1.3/catalog/video/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
