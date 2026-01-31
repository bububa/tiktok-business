package discovery

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// VideoListRequest Get trending videos related to hashtags
type VideoListRequest struct {
	// AdvertiserID Advertiser ID.
	// Pass one of the advertiser IDs returned via advertiser_ids from /oauth2/access_token/.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// DiscoveryType The type of discovery item.
	// Currently, we only support HASHTAG.
	DiscoveryType enum.DiscoveryType `json:"discovery_type,omitempty"`
	// HashtagIDs A list of hashtag IDs.
	// Max size: 10.
	// Pass one or more popular hashtag IDs obtained via hashtag_id from /discovery/trending_list/.
	HashtagIDs []string `json:"hashtag_ids,omitempty"`
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
func (r *VideoListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("discovery_type", string(r.DiscoveryType))
	values.Set("hashtag_ids", string(util.JSONMarshal(r.HashtagIDs)))
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

// VideoListResponse Get trending videos related to hashtags
type VideoListResponse struct {
	model.BaseResponse
	Data struct {
		// List The list of trending videos using popular hashtags.
		// This object array returns the top 20 videos related to each popular hashtag in the selected region and timeframe.
		List []Hashtag `json:"list,omitempty"`
	} `json:"data"`
}
