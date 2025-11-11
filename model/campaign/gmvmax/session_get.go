package gmvmax

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SessionGetRequest 获取最大投放量或创意作品加热时段的详情 API Request
type SessionGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// SessionIDs 商品 GMV Max 推广系列中的商品的最大投放量时段或创意作品加热时段 ID 列表。
	// 最大数量: 20。
	// 若想获取商品 GMV Max 推广系列中生效状态的最大投放量时段的 ID 列表，可使用 /campaign/gmv_max/session/list/ 并挑选 bid_type 为 NO_BID 的时段。
	// 若想获取商品 GMV Max 推广系列中生效状态的创意作品加热时段的 ID 列表，可使用 /campaign/gmv_max/session/list/ 并挑选 bid_type 为 CREATIVE_NO_BID 的时段。
	SessionIDs []string `json:"session_ids,omitempty"`
}

// Encode implements GetRequest
func (r *SessionGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("session_ids", string(util.JSONMarshal(r.SessionIDs)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// SessionGetResponse 获取最大投放量或创意作品加热时段的详情 API Response
type SessionGetResponse struct {
	model.BaseResponse
	Data struct {
		// SessionList 最大投放量时段或创意作品加热时段列表
		SessionList []Session `json:"session,omitempty"`
	} `json:"data"`
}
