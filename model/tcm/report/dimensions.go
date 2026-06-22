package report

import "github.com/bububa/tiktok-business/model"

// Dimensions 维度数据
type Dimensions struct {
	// TTOTCMAccountID TikTok One Creator Marketplace 账户 ID
	TTOTCMAccountID string `json:"tto_tcm_account_id,omitempty"`
	// CampaignID TikTok One Creator Marketplace 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CountryCode 与 TikTok One Creator Marketplace 推广系列关联的国家或地区代码
	CountryCode string `json:"country_code,omitempty"`
	// CampaignType TTO Creator Marketplace 推广系列类型
	CampaignType string `json:"campaign_type,omitempty"`
	// CampaignChannelType 推广系列创建渠道类型
	CampaignChannelType string `json:"campaign_channel_type,omitempty"`
	// TCMAccountID TTCM 账户 ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// OrderID 订单 ID
	OrderID string `json:"order_id,omitempty"`
	// OrderType 订单类型，代表订单的创建方式。
	// 枚举值：API_ORDER_V2 （订单通过 /tcm/order/create/v2/接口创建）
	OrderType string `json:"order_type,omitempty"`
	// VideoID 视频 ID
	VideoID string `json:"video_id,omitempty"`
	// HandleName TikTok创作者的用户名
	HandleName string `json:"handle_name,omitempty"`
	// VideoFirstReleaseTime 视频首次发布的时间，格式：YYYY-MM-DD（UTC+0 时间）。 示例： 2020-06-10
	VideoFirstReleaseTime string `json:"video_first_release_time,omitzero"`
}

// VideoInfo 视频信息
type VideoInfo struct {
	// CreateTime 视频发布时间，格式：YYYY-MM-DD HH:MM:SS（UTC+0 时间）
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// EmbedURL TikTok 视频嵌入链接
	EmbedURL string `json:"embed_url,omitempty"`
	// ThumbnailURL 视频缩略图临时链接
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// VideoID 视频 ID
	VideoID string `json:"video_id,omitempty"`
	// AuthorizationCode Spark Ads 授权码
	AuthorizationCode string `json:"authorization_code,omitempty"`
}

// CreatorInfo TikTok 创作者信息
type CreatorInfo struct {
	// DisplayName 创作者展示名称
	DisplayName string `json:"display_name,omitempty"`
	// HandleName 创作者用户名
	HandleName string `json:"handle_name,omitempty"`
}
