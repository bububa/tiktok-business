package discovery

import "github.com/bububa/tiktok-business/model"

// TrendingHistory The historical ranking data for the track, spanning up to the past 30 days.
type TrendingHistory struct {
	// Date Date, in the format of YYYY-MM-DD
	Date string `json:"date,omitempty"`
	// RankPositionDaily The rank position of the track in the selected region (country_code) on the date.
	// Example: 73.
	// Note: If the track does not perform well compared to the last time period, the track will not be considered a popular track and the value of this field will be NULL.
	RankPositionDaily model.Int `json:"rank_position_daily,omitempty"`
	// ViewsDaily The number of views for related videos that feature this hashtag in the selected region (country_code) on the date.
	// Example: 21215985.
	// Note: If the daily views are less than 1,000, the value of this field will be null.
	ViewsDaily int64 `json:"views_daily,omitempty"`
}
