package pixel

import "github.com/bububa/tiktok-business/util"

// EventCreateRequest 创建Pixel事件 API Request
type EventCreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// PixelID Pixel ID
	PixelID string `json:"pixel_id,omitempty"`
	// PixelEvents Pixel 事件列表
	PixelEvents []Event `json:"pixel_events,omitempty"`
}

// Encode implements PostRequest interface
func (r *EventCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
