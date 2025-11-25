package dmp

import "github.com/bububa/tiktok-business/util"

// CustomAudienceShareRequest 分享受众 API Request
type CustomAudienceShareRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceIDs 您想要分享的受众ID。数量大小：1-10
	CustomAudienceIDs []string `json:"custom_audience_ids,omitempty"`
	// SharedAdvertiserIDs 您想要分享给的广告主ID，这些广告主必须与您在同一个商务中心下。数量大小：1-10
	SharedAdvertiserIDs []string `json:"shared_advertiser_ids,omitempty"`
}

// Encode implements PostRequest
func (r *CustomAudienceShareRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
