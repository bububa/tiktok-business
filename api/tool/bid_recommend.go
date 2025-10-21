package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// BidRecommend 获取建议出价
// 您可以使用本接口获取系统建议出价，建议出价基于基本的推广系列及广告组的设置（如推广目标、转化事件以及地域定向等）。获取的建议出价仅能用于成本上限出价策略（即bid_type=BID_TYPE_CUSTOM)。
func BidRecommend(ctx context.Context, clt *core.SDKClient, req *tool.BidRecommendRequest, accessToken string) (float64, error) {
	var ret tool.BidRecommendResponse
	if err := clt.Post(ctx, "v1.3/tool/bid/recommend/", req, &ret, accessToken); err != nil {
		return 0, err
	}
	return ret.Data.Bid, nil
}
