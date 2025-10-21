package enum

// TikTokURLType TikTok公开网址类型
type TikTokURLType string

const (
	// TikTokURLType_USER_PROFILE：用户档案页面。
	TikTokURLType_USER_PROFILE TikTokURLType = "USER_PROFILE"
	// TikTokURLType_VIDEO：视频详情页面。
	TikTokURLType_VIDEO TikTokURLType = "VIDEO"
	// TikTokURLType_HASHTAG_CHALLENGE：挑战赛页面。
	TikTokURLType_HASHTAG_CHALLENGE TikTokURLType = "HASHTAG_CHALLENGE"
	// TikTokURLType_MUSIC：音乐页面。
	TikTokURLType_MUSIC TikTokURLType = "MUSIC"
	// TikTokURLType_STICKER：贴纸（特效）页面。
	TikTokURLType_STICKER TikTokURLType = "STICKER"
	// TikTokURLType_STICKER_SHOOTER：贴纸开拍页面。
	TikTokURLType_STICKER_SHOOTER TikTokURLType = "STICKER_SHOOTER"
)
