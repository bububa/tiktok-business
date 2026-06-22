package report

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Dimensions is the dimension payload returned by basic reports.
type Dimensions struct {
	// AdvertiserID is returned when dimensions contains advertiser_id.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID is returned when dimensions contains campaign_id.
	CampaignID string `json:"campaign_id,omitempty"`
	// AdgroupID is returned when dimensions contains adgroup_id.
	AdgroupID string `json:"adgroup_id,omitempty"`

	// AdIDV2 groups by Ad ID.
	// The types of data you can get include:
	// the ad level data for Manual Ads,
	// the ad level data for Smart+ Campaigns,
	// and the ad (asset group) level data for Upgraded Smart+ Ads.
	AdIDV2 string `json:"ad_id_v2,omitempty"`
	// AdID groups by Ad ID.
	// The types of data you can get include:
	// the ad level data for Manual Ads,
	// the ad level data for Smart+ Campaigns,
	// and the creative level data for Upgraded Smart+ Ads.
	AdID string `json:"ad_id,omitempty"`

	// StatTimeDay groups by day.
	StatTimeDay model.DateTime `json:"stat_time_day,omitzero"`
	// StatTimeHour groups by hour.
	StatTimeHour model.DateTime `json:"stat_time_hour,omitzero"`

	// CountryCode groups by location code.
	CountryCode string `json:"country_code,omitempty"`
	// AdType groups by Ad Type.
	AdType enum.AdType `json:"ad_type,omitempty"`
	// CustomEventType groups by custom event type.
	CustomEventType enum.OptimizationEvent `json:"custom_event_type,omitempty"`
	// EventSourceID groups by event source ID.
	EventSourceID string `json:"event_source_id,omitempty"`
	// PageID groups by page ID.
	PageID string `json:"page_id,omitempty"`
	// ComponentName groups by component name on the TikTok Instant Form.
	ComponentName string `json:"component_name,omitempty"`
	// RoomID groups by live room ID.
	RoomID string `json:"room_id,omitempty"`
	// PostID groups by post ID.
	PostID string `json:"post_id,omitempty"`
	// ImageID groups by image ID.
	ImageID string `json:"image_id,omitempty"`
	// VideoMaterialID groups by video material ID.
	VideoMaterialID string `json:"video_material_id,omitempty"`
	// MinisID groups by minis ID.
	MinisID string `json:"minis_id,omitempty"`

	// SearchTerms groups by Search Term.
	SearchTerms string `json:"search_terms,omitempty"`
	// SearchKeyword groups by search keywords.
	SearchKeyword string `json:"search_keyword,omitempty"`
	// MatchType groups by match type.
	MatchType enum.KeywordMatchType `json:"match_type,omitempty"`

	// SkanPostbackSequence groups by SKAN Postback Sequence for iOS 14 Dedicated Campaigns.
	SkanPostbackSequence model.Int64 `json:"skan_postback_sequence,omitempty"`
}
