package creator

import "github.com/bububa/tiktok-business/model"

// Post 帖子
type Post struct {
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// Caption 视频标题
	Caption string `json:"caption,omitempty"`
	// DisplayName 创作者的显示名称
	DisplayName string `json:"display_name,omitempty"`
	// Likes 视频的点赞数
	Likes int64 `json:"likes,omitempty"`
	// Comments 视频获得的评论数
	Comments int64 `json:"comments,omitempty"`
	// Shares 视频的分享次数
	Shares int64 `json:"shares,omitempty"`
	// VideoViews 视频播放量。指用户首次播放视频的次数。重播不计算在内
	VideoViews int64 `json:"video_views,omitempty"`
	// EngagementCount 点赞、评论和分享次数合计
	EngagementCount int64 `json:"engagement_count,omitempty"`
	// EngagementRate 计算公式为：engagement_count / video_views
	EngagementRate float64 `json:"engagement_rate,omitempty"`
	// EmbedURL 视频 URL
	EmbedURL string `json:"embed_url,omitempty"`
	// ThumbnailURL 视频缩略图的 URL
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// CreateTime 视频创建时间
	CreateTime model.DateTime `json:"create_time,omitzero"`
}
