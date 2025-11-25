package dmp

import "github.com/bububa/tiktok-business/util"

// CustomAudienceShareCancelRequest 取消受众分享 API Request
type CustomAudienceShareCancelRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceID 您想要停止分享的受众ID
	CustomAudienceID string `json:"custom_audience_id,omitempty"`
	// SharedAdvertiserID 您要停止与其进行共享的广告主ID。只有当您想要取消与广告主分享受众时，您才需要传入该字段
	SharedAdvertiserID string `json:"shared_advertiser_id,omitempty"`
}

// Encode implements PostRequest
func (r *CustomAudienceShareCancelRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
