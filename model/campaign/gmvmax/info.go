package gmvmax

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InfoRequest 获取 GMV Max 推广系列的详情 API Request
type InfoRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID GMV Max 推广系列 ID。
	// 若想创建 GMV Max 推广系列并获取对应 ID，可使用/campaign/gmv_max/create/。
	// 若想筛选广告账户中的现有 GMV Max 推广系列，可使用/gmv_max/campaign/get/。
	CampaignID string `json:"campaign_id,omitempty"`
}

// Encode implements GetRequest interface
func (r *InfoRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("campaign_id", r.CampaignID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// InfoResponse 获取 GMV Max 推广系列的详情 API Response
type InfoResponse struct {
	model.BaseResponse
	Data *Campaign `json:"data,omitempty"`
}
