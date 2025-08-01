package enum

// TiktokPageCategory 您希望推广的 TikTok 页面的类型。
type TiktokPageCategory string

const (
	// TiktokPageCategory_PROFILE_PAGE：TikTok 账号主页。
	TiktokPageCategory_PROFILE_PAGE TiktokPageCategory = "PROFILE_PAGE"
	// TiktokPageCategory_OTHER_TIKTOK_PAGE：其他 TikTok 页面，包括播放列表页面、话题标签页面、音乐页面和商业化特效页面。
	TiktokPageCategory_OTHER_TIKTOK_PAGE TiktokPageCategory = "OTHER_TIKTOK_PAGE"
	// TiktokPageCategory_TIKTOK_INSTANT_PAGE：自定义 TikTok 即时体验页面。
	TiktokPageCategory_TIKTOK_INSTANT_PAGE TiktokPageCategory = "TIKTOK_INSTANT_PAGE"
)
