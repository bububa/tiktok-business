package gmvmax

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SessionListRequest 获取推广系列中的最大投放量或创意作品加热时段 API Request
type SessionListRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 商品 GMV Max 推广系列的 ID。
	// 若要创建商品 GMV Max 推广系列并获取推广系列 ID，可使用/campaign/gmv_max/create/，并将 shopping_ads_type 设置为 PRODUCT。
	// 若要筛选广告账户中已有的商品 GMV Max 推广系列，可使用 /gmv_max/campaign/get/ 并将 filtering 设置为 {"gmv_max_promotion_types":["PRODUCT_GMV_MAX"]}。
	CampaignID string `json:"campaign_id,omitempty"`
}

// Encode implements GetRequest
func (r *SessionListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("campaign_id", r.CampaignID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// SessionListResponse 获取推广系列中的最大投放量或创意作品加热时段 API Response
type SessionListResponse struct {
	model.BaseResponse
	Data struct {
		// SessionList 最大投放量时段或创意作品加热时段列表
		SessionList []Session `json:"session,omitempty"`
	} `json:"data"`
}
