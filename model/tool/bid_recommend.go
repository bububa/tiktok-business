package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// BidRecommendRequest 获取建议出价 API Request
type BidRecommendRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列ID
	CampaignID string `json:"campaign_id,omitempty"`
	// ObjectiveType 推广系列的推广目标。该接口目前仅支持访问量和网站转化量目标。枚举值及描述请查看推广目标。
	// 注意: 该接口目前正在进行优化，以扩大支持的功能范围、提高结果准确性。请密切关注更新信息。
	ObjectiveType []enum.ObjectiveType `json:"objective_type,omitempty"`
	// LocationIDs 你想要定向的地域ID列表
	LocationIDs []string `json:"location_ids,omitempty"`
	// ExternalAction 广告组的转化事件
	ExternalAction enum.OptimizationEvent `json:"external_action,omitempty"`
}

// Encode implements PostRequest interface
func (r *BidRecommendRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BidRecommendResponse 获取建议出价 API Response
type BidRecommendResponse struct {
	model.BaseResponse
	Data struct {
		// Bid 建议出价的值
		Bid float64 `json:"bid,omitempty"`
	} `json:"data"`
}
