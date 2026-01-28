package discovery

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DetailRequest Get details of a popular hashtag API Request
type DetailRequest struct {
	// AdvertiserID Advertiser ID.
	// Pass one of the advertiser IDs returned via advertiser_ids from /oauth2/access_token/.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// DiscoveryType The type of discovery item.
	// Currently, we only support HASHTAG.
	DiscoveryType enum.DiscoveryType `json:"discovery_type,omitempty"`
	// HashtagID Hashtag ID.
	// Pass in a hashtag ID that is returned via the /discovery/trending_list/ endpoint.
	// If the hashtag you provide isn't included in the list returned by the /discovery/trending_list/ endpoint, you will receive an empty list in response.
	HashtagID string `json:"hashtag_id,omitempty"`
	// CountryCode Location code.
	// Default value: US.
	// See Location code for the supported values.
	CountryCode string `json:"country_code,omitempty"`
	// DateRange The timeframe in which the items have been popular.
	// Enum values:
	// 1DAY: last 1 day.
	// 7DAY: last 7 days.
	// 30DAY: last 30 days.
	// 120DAY: last 120 days.
	// Default value: 7DAY.
	DateRange string `json:"date_range,omitempty"`
}

// Encode implements GetRequest
func (r *DetailRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("discovery_type", string(r.DiscoveryType))
	values.Set("hashtag_id", r.HashtagID)
	if r.CountryCode != "" {
		values.Set("country_code", r.CountryCode)
	}
	if r.DateRange != "" {
		values.Set("date_range", r.DateRange)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// DetailResponse Get details of a popular hashtag API Response
type DetailResponse struct {
	model.BaseResponse
	Data *Hashtag `json:"data,omitempty"`
}
