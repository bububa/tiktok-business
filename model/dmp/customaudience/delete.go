package customaudience

import "github.com/bububa/tiktok-business/util"

// DeleteRequest 删除受众 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceIDs 受众群体ID列表。长度范围[1, 100]
	CustomAudienceIDs []string `json:"custom_audience_ids,omitempty"`
}

// Encode implements PostRequest
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
