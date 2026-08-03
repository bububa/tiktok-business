package app

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InfoRequest gets one app under an advertiser account.
type InfoRequest struct {
	AdvertiserID string `json:"advertiser_id,omitempty"`
	AppID        string `json:"app_id,omitempty"`
}

// Encode implements model.GetRequest.
func (r *InfoRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("app_id", r.AppID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// InfoResponse is the response to InfoRequest.
type InfoResponse struct {
	model.BaseResponse
	Data struct {
		App *App `json:"app,omitempty"`
	} `json:"data"`
}
