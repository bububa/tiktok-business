package savedaudience

import "github.com/bububa/tiktok-business/util"

// DeleteRequest 删除已保存受众 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// SavedAudienceIDs 已保存受众 ID 列表。
	// 最大数量：100。
	SavedAudienceIDs []string `json:"saved_audience_ids,omitempty"`
}

// Encode implements PostRequest
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
