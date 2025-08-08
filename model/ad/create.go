package ad

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建广告 API Request
type CreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdgroupID 广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// Creatives 广告创意。
	// 最大数量：20。
	// 单次请求中可在一个广告组中创建最多 20 个广告。若想在同一广告组中创建额外的广告，需多次调用/ad/create/。
	Creatives []Creative `json:"creatives,omitempty"`
}

// Encode implements PostRequest interface
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建广告 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Ad `json:"data,omitempty"`
}
