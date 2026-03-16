package adspreview

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// SmartPlusAdPreviewRequest Preview Upgraded Smart+ Ads API Request
type SmartPlusAdPreviewRequest struct {
	CreateRequest
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// SmartPlusAdID ID of the Upgraded Smart+ Ad that you want to preview
	SmartPlusAdID string `json:"smart_plus_ad_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *SmartPlusAdPreviewRequest) Encode() []byte {
	r.PreviewType = enum.AdPreviewType_AD
	return util.JSONMarshal(r)
}

func (r *SmartPlusAdPreviewRequest) Gateway() string {
	return "smart_plus/ad/preview/"
}
