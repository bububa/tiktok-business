package customaudience

import "github.com/bububa/tiktok-business/util"

// LookalikeDeleteRequest 删除受众 API Request
type LookalikeDeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceIDs 自定义受众ID列表
	CustomAudienceIDs []string `json:"custom_audience_ids,omitempty"`
}

// Encode implements PostRequest
func (r *LookalikeDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
