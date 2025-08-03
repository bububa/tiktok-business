package spc

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// QuotaGetRequest 获取投放中 Smart+ 推广系列的动态配额 API Request
type QuotaGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest interface
func (r *QuotaGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// QuotaGetResponse 获取投放中 Smart+ 推广系列的动态配额 API Response
type QuotaGetResponse struct {
	model.BaseResponse
	Data *Quota `json:"data,omitempty"`
}

type Quota struct {
	// TotalQuota 广告账户的投放中 Smart+ 推广系列总配额
	TotalQuota int64 `json:"total_quota,omitempty"`
	// UsedQuota 广告账户的投放中 Smart+ 推广系列已使用的配额
	UsedQuota int64 `json:"used_quota,omitempty"`
	// UsedCampaignIDs 仅当 used_quota 大于 0 时返回本字段。
	// 广告账户中已创建的投放中 Smart+ 推广系列的 ID 列表。
	// 若想暂停投放中的 Smart+ 推广系列，可使用 /campaign/status/update/ 并将 operation_status 设置为 DISABLE。
	UsedCampaignIDs []string `json:"used_campaign_ids,omitempty"`
}
