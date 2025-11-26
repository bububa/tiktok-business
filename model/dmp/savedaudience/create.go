package savedaudience

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建已保存受众
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	Audience
}

// Encode implements PostRequest
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建已保存受众
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// SavedAudienceID 已保存受众ID
		SavedAudienceID string `json:"saved_audience_id,omitempty"`
	} `json:"data"`
}
