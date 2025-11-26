package customaudience

import "github.com/bububa/tiktok-business/util"

// LookalikeUpdateRequest 刷新相似受众 API Request
type LookalikeUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceIDs 自定义受众ID列表
	CustomAudienceIDs []string `json:"custom_audience_ids,omitempty"`
}

// Encode implements PostRequest
func (r *LookalikeUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
