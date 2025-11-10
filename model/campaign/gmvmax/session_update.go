package gmvmax

import (
	"github.com/bububa/tiktok-business/util"
)

// SessionUpdateRequest 更新最大投放量或创意作品加热时段 API Request
type SessionUpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 商品 GMV Max 推广系列的 ID。
	// 若要创建商品 GMV Max 推广系列并获取推广系列 ID，可使用/campaign/gmv_max/create/，并将 shopping_ads_type 设置为 PRODUCT。
	// 若要筛选广告账户中已有的商品 GMV Max 推广系列，可使用 /gmv_max/campaign/get/ 并将 filtering 设置为 {"gmv_max_promotion_types":["PRODUCT_GMV_MAX"]}。
	CampaignID string `json:"campaign_id,omitempty"`
	// StoreID TikTok Shop ID。
	// 若想获取和 GMV Max 推广系列绑定的 TikTok Shop 的 ID，可使用 /campaign/gmv_max/info/。
	StoreID string `json:"store_id,omitempty"`
	// SessionID 商品 GMV Max 推广系列中某一视频的创意作品加热时段 ID。
	// 若想获取商品 GMV Max 推广系列中生效状态的创意作品加热时段的 ID 列表，可使用 /campaign/gmv_max/session/list/ 并选择 bid_type 为 CREATIVE_NO_BID 的时段。
	SessionID string `json:"session_id,omitempty"`
	// Session 有关该推广系列最大投放量模式的详情
	Session *Session `json:"session,omitempty"`
}

// Encode implements PostRequest
func (r *SessionUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
