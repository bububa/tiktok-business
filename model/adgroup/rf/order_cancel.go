package rf

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// OrderCancelRequest 撤销覆盖和频次广告订单 API Request
type OrderCancelRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdgroupID 要撤单的广告组ID。需对应覆盖和频次广告组
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
}

// Encode implements PostRequest
func (r *OrderCancelRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// OrderCancelResponse 撤销覆盖和频次广告订单 API Response
type OrderCancelResponse struct {
	model.BaseResponse
	Data struct {
		// FailAdgroupIDs 撤单失败的广告组ID
		FailAdgroupIDs []string `json:"fail_adgroup_ids,omitempty"`
	} `json:"data"`
}
