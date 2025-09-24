package customconversion

import "github.com/bububa/tiktok-business/util"

// DeleteRequest 删除自定义转化 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomConversionID 自定义转化的 ID
	CustomConversionID string `json:"custom_conversion_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
