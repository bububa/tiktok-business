package dmp

import "github.com/bububa/tiktok-business/util"

// CustomAudienceLookalikeDeleteRequest 删除受众 API Request
type CustomAudienceLookalikeDeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceIDs 自定义受众ID列表
	CustomAudienceIDs []string `json:"custom_audience_ids,omitempty"`
}

// Encode implements PostRequest
func (r *CustomAudienceLookalikeDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
