package gmvmax

import (
	"github.com/bububa/tiktok-business/util"
)

// SessionDeleteRequest 创建最大投放量或创意作品加热时段 API Request
type SessionDeleteRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// SessionID 商品 GMV Max 推广系列中某一视频的创意作品加热时段 ID。
	// 若想获取商品 GMV Max 推广系列中生效状态的最大投放量时段的 ID 列表，可使用 /campaign/gmv_max/session/list/ 并挑选 bid_type 为 NO_BID 的时段。
	// 若想获取商品 GMV Max 推广系列中生效状态的创意作品加热时段的 ID 列表，可使用 /campaign/gmv_max/session/list/ 并挑选 bid_type 为 CREATIVE_NO_BID 的时段。
	SessionID string `json:"session_id,omitempty"`
}

// Encode implements PostRequest
func (r *SessionDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
