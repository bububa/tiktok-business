package spc

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取 Smart+ 推广系列 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignIDs 推广系列 ID 列表。
	// 最大数量：100。
	CampaignIDs []string `json:"campaign_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("campaign_ids", string(util.JSONMarshal(r.CampaignIDs)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

type GetResponse struct {
	model.BaseResponse
	Data struct {
		// List Smart+ 推广系列或智能效果网站推广系列列表
		List []Campaign `json:"list,omitempty"`
	} `json:"data"`
}
