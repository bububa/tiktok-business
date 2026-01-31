package discovery

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TrendingListRequest Get popular hashtags API Request
type TrendingListRequest struct {
	// AdvertiserID Advertiser ID.
	// Pass one of the advertiser IDs returned via advertiser_ids from /oauth2/access_token/.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// DiscoveryType The type of discovery item.
	// Currently, we only support HASHTAG.
	DiscoveryType enum.DiscoveryType `json:"discovery_type,omitempty"`
	// CountryCode Location code.
	// Default value: US.
	// See Location code for the supported values.
	CountryCode string `json:"country_code,omitempty"`
	// CategoryName The name of the hashtag industry category.
	// To learn about the enum values, see Hashtag industry category names. You can pass either a level-1 category name or a level-2 category name.
	// Default value: ALL.
	// Example: FOOD_AND_BEVERAGE.
	// If you pass category_name and industry_id at the same time, industry_id will be ignored.
	CategoryName string `json:"category_name,omitempty"`
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
func (r *TrendingListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("discovery_type", string(r.DiscoveryType))
	if r.CountryCode != "" {
		values.Set("country_code", r.CountryCode)
	}
	if r.CategoryName != "" {
		values.Set("category_name", r.CategoryName)
	}
	if r.DateRange != "" {
		values.Set("date_range", r.DateRange)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// TrendingListResponse Get popular hashtags API Response
type TrendingListResponse struct {
	model.BaseResponse
	Data *TrendingListResult `json:"data,omitempty"`
}

type TrendingListResult struct {
	// FilterInfo Filtering information that is specified in your request.
	// If you don't include the following parameters in the request, we will return the default values for them, which are date_range = 7DAY, country_code = US, industry_id = ALL, and category_name = ALL.
	FilterInfo *TrendingListRequest `json:"filter_info,omitempty"`
	// List The list of popular hashtags
	List []Hashtag `json:"list,omitempty"`
}

type Hashtag struct {
	// HashtagID Hashtag ID.
	HashtagID string `json:"hashtag_id,omitempty"`
	// HashtagName Hashtag name
	HashtagName string `json:"hashtag_name,omitempty"`
	// RankPosition Current rank position of the hashtag among hashtags within the same category (category_name) or industry (industry_id) in the selected region (country_code).
	RankPosition model.Int `json:"rank_position,omitempty"`
	// RankChange The change in rank position compared to the rank in the last time period.
	// A positive rank_change indicates an increase in rank. For example, if the rank position shifts from 5 to 1, rank_change will be 4.
	// A negative rank_change indicates a decrease in rank. For example, if the rank position shifts from 1 to 5, rank_change will be -4.
	// A rank_change of 0 indicates no changes in rank.
	// A rank_change marked as NEW indicates that the hashtag was not present in the top 200 ranking list in the last time period.
	RankChange model.Int `json:"rank_change,omitempty"`
	// Views The number of views for related videos that feature this hashtag in the selected region and timeframe.
	// Note: If you fail to specify country_code or date_range in your request, we will return data based on the default values for these two parameters, which are date_range = 7DAY or country_code = US.
	Views int64 `json:"views,omitempty"`
	// ViewsGlobalLifetime The number of views for related videos that feature this hashtag, globally and across their whole lifetime
	ViewsGlobalLifetime int64 `json:"views_global_lifetime,omitempty"`
	// Posts The number of posts for related videos that feature this hashtag in the selected region and timeframe.
	// Note: If you fail to specify country_code or date_range in your request, we will return data based on the default values for these two parameters, which are date_range = 7DAY or country_code = US.
	Posts int64 `json:"posts,omitempty"`
	// PostsGlobalLifetime The number of posts for related videos that feature this hashtag, globally and across their whole lifetime
	PostsGlobalLifetime int64 `json:"posts_global_lifetime,omitempty"`
	// TopCountryList The location codes of the top five locations in the world where this hashtag is most popular.
	// To learn about the location that a location code represents, see Location code.
	TopCountryList []string `json:"top_country_list,omitempty"`
	// HashtagStatus The status of the hashtag.
	// Enum values:
	// ONLINE: available for ad creation.
	// OFFLINE: unavailable for ad creation.
	HashtagStatus enum.OnlineOffline `json:"hashtag_status,omitempty"`
	// TrendingHistory The historical ranking data for the hashtag, spanning up to the past 30 days
	TrendingHistory []TrendingHistory `json:"trending_history,omitempty"`
	// AudienceInsights Information about viewers that watch videos associated with this hashtag.
	// Note: If hashtag_status is OFFLINE or the expected audience size for the hashtag is less than 1,000, audience_insights will be returned as empty.
	AudienceInsights *AudienceInsights `json:"audience_insights,omitempty"`
	// TopVideoList The list of trending videos for the hashtag.
	TopVideoList []Video `json:"top_video_list,omitempty"`
}

// AudienceInsights Information about viewers that watch videos associated with this hashtag.
type AudienceInsights struct {
	// AudienceAges Information about viewers' age
	AudienceAges []AudienceAge `json:"audience_ages,omitempty"`
}

// AudienceAge Information about viewers' age
type AudienceAge struct {
	// Age Age group
	Age string `json:"age,omitempty"`
	// Percentage The approximate percentage of viewers associated with different age groups
	Percentage float64 `json:"percentage,omitempty"`
}

type Video struct {
	// VideoID Unique identifier for the TikTok video
	VideoID string `json:"video_id,omitempty"`
	// EmbedURL An embeddable link for the TikTok video.
	// This URL can be used to embed the TikTok video in external websites or applications.
	// If the privacy setting of the post is "Friends" or "Only you", the video will not be viewable through this link.
	EmbedURL string `json:"embed_url,omitempty"`
	// ShareURL A shareable URL for the TikTok video
	ShareURL string `json:"share_url,omitempty"`
}
