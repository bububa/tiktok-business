package app

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest creates an app under an advertiser account.
type CreateRequest struct {
	AdvertiserID      string                   `json:"advertiser_id,omitempty"`
	DownloadURL       string                   `json:"download_url,omitempty"`
	TrackingURL       *TrackingURL             `json:"tracking_url,omitempty"`
	Partner           string                   `json:"partner,omitempty"`
	EnableRetargeting enum.AppRetargetingState `json:"enable_retargeting,omitempty"`
}

// Encode implements model.PostRequest.
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse is the response to CreateRequest.
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		AppID string `json:"app_id,omitempty"`
	} `json:"data"`
}
