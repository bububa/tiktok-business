package discovery

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CMLVideoListRequest Get trending videos related to tracks API Request
type CMLVideoListRequest struct {
	// BusinessID Application specific unique identifier for the TikTok account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// CommercialMusicID The ID of the track
	CommercialMusicID string `json:"commercial_music_id,omitempty"`
	// CountryCode The code of the location to filter the results by.
	// Default value: US.
	// See Location code for the supported values.
	CountryCode string `json:"country_code,omitempty"`
}

// Encode implements GetRequest
func (r *CMLVideoListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	values.Set("commercial_music_id", r.CommercialMusicID)
	if r.CountryCode != "" {
		values.Set("country_code", r.CountryCode)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// CMLVideoListResponse Get trending videos related to tracks API Response
type CMLVideoListResponse struct {
	model.BaseResponse
	Data *CommercialMusicLibrary `json:"data,omitempty"`
}
