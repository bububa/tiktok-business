package enum

// TikTokItemType 想要获取的 TikTok 帖子类型。
type TikTokItemType string

const (
	// TikTokItemType_VIDEO：视频帖子。
	TikTokItemType_VIDEO TikTokItemType = "VIDEO"
	// TikTokItemType_CAROUSEL：照片帖子。
	TikTokItemType_CAROUSEL TikTokItemType = "CAROUSEL"
)
