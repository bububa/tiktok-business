package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// BidRecommend 获取推荐 GMV Max ROI 目标和预算
// 您可以使用本接口获取使用了某一 TikTok Shop 的商品 GMV Max 推广系列或直播 GMV Max 推广系列的推荐 ROI （投资回报率）目标和预算
func BidRecommend(ctx context.Context, clt *core.SDKClient, req *gmvmax.BidRecommendRequest, accessToken string) (*gmvmax.BidRecommend, error) {
	var resp gmvmax.BidRecommendResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/bid/recommend/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
