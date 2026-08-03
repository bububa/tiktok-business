package app

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest gets apps under an advertiser account.
type ListRequest struct {
	AdvertiserID   string   `json:"advertiser_id,omitempty"`
	AppPlatformIDs []string `json:"app_platform_ids,omitempty"`
}

// Encode implements model.GetRequest.
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.AppPlatformIDs) > 0 {
		values.Set("app_platform_ids", string(util.JSONMarshal(r.AppPlatformIDs)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ListResponse is the response to ListRequest.
type ListResponse struct {
	model.BaseResponse
	Data struct {
		Apps []App `json:"apps,omitempty"`
	} `json:"data"`
}
