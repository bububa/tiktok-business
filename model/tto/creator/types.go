package creator

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

type Label struct {
	LabelID   model.Int64 `json:"label_id,omitempty"`
	LabelName string      `json:"label_name,omitempty"`
}

type Distribution struct {
	Country    string              `json:"country,omitempty"`
	Gender     enum.Gender         `json:"gender,omitempty"`
	Age        enum.TTOAudienceAge `json:"age,omitempty"`
	Device     string              `json:"device,omitempty"`
	Usage      string              `json:"usage,omitempty"`
	Percentage model.Float64       `json:"percentage,omitempty"`
}

type CreatorRate struct {
	Rate     model.Float64 `json:"rate,omitempty"`
	Currency string        `json:"currency,omitempty"`
}

type Creator struct {
	CreatorID         string         `json:"creator_id,omitempty"`
	HandleName        string         `json:"handle_name,omitempty"`
	Handle            string         `json:"handle,omitempty"`
	DisplayName       string         `json:"display_name,omitempty"`
	ProfileImage      string         `json:"profile_image,omitempty"`
	Bio               string         `json:"bio,omitempty"`
	FollowersCount    model.Int64    `json:"followers_count,omitempty"`
	FollowingCount    model.Int64    `json:"following_count,omitempty"`
	LikesCount        model.Int64    `json:"likes_count,omitempty"`
	VideosCount       model.Int64    `json:"videos_count,omitempty"`
	MedianViews       model.Int64    `json:"median_views,omitempty"`
	EngagementRate    model.Float64  `json:"engagement_rate,omitempty"`
	CreatorPrice      model.Float64  `json:"creator_price,omitempty"`
	Currency          string         `json:"currency,omitempty"`
	CreatorRate       *CreatorRate   `json:"creator_rate,omitempty"`
	CountryCode       string         `json:"country_code,omitempty"`
	IndustryLabels    []Label        `json:"industry_labels,omitempty"`
	ContentLabels     []Label        `json:"content_labels,omitempty"`
	AudienceCountries []Distribution `json:"audience_countries,omitempty"`
	AudienceGenders   []Distribution `json:"audience_genders,omitempty"`
	AudienceAges      []Distribution `json:"audience_ages,omitempty"`
	AudienceDevices   []Distribution `json:"audience_devices,omitempty"`
	AudienceUsages    []Distribution `json:"audience_usages,omitempty"`
}

type Post struct {
	VideoID             string         `json:"video_id,omitempty"`
	DisplayName         string         `json:"display_name,omitempty"`
	ThumbnailURL        string         `json:"thumbnail_url,omitempty"`
	EmbedURL            string         `json:"embed_url,omitempty"`
	Caption             string         `json:"caption,omitempty"`
	Likes               model.Int64    `json:"likes,omitempty"`
	Comments            model.Int64    `json:"comments,omitempty"`
	Shares              model.Int64    `json:"shares,omitempty"`
	Favorites           model.Int64    `json:"favorites,omitempty"`
	VideoViews          model.Int64    `json:"video_views,omitempty"`
	CreateTime          model.DateTime `json:"create_time,omitzero"`
	CaptionHashtags     []string       `json:"caption_hashtags,omitempty"`
	MentionedAccounts   []string       `json:"mentioned_accounts,omitempty"`
	CollaborationStatus string         `json:"collaboration_status,omitempty"`
}
