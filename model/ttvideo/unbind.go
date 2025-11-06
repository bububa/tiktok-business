package ttvideo

import "github.com/bububa/tiktok-business/util"

// UnbindRequest 解除 Spark Ads 帖子绑定 API Request
type UnbindRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// ItemID Spark Ads帖子ID
	ItemID string `json:"item_id,omitempty"`
}

// Encode implements PostRequest
func (r *UnbindRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
