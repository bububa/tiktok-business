package app

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest updates an existing app.
type UpdateRequest struct {
	AdvertiserID      string                   `json:"advertiser_id,omitempty"`
	AppID             string                   `json:"app_id,omitempty"`
	Platform          enum.OperatingSystem     `json:"platform,omitempty"`
	DownloadURL       string                   `json:"download_url,omitempty"`
	EnableRetargeting enum.AppRetargetingState `json:"enable_retargeting,omitempty"`
	TrackingURL       *TrackingURL             `json:"tracking_url,omitempty"`
	Partner           string                   `json:"partner,omitempty"`
}

// Encode implements model.PostRequest.
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
