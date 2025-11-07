package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// VideoGet 获取商品 GMV Max 推广系列可用的帖子
// 您可以使用本接口获取使用了某一 TikTok Shop 的商品 GMV Max 推广系列中可用的 TikTok 帖子列表。
func VideoGet(ctx context.Context, clt *core.SDKClient, req *gmvmax.VideoGetRequest, accessToken string) (*gmvmax.VideoGetResult, error) {
	var resp gmvmax.VideoGetResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/video/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
