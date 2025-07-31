package enum

// TiktokSubplacement TikTok 子版位，用于选择广告出现的具体位置。
type TiktokSubplacement string

const (
	// TiktokSubplacement_IN_FEED：信息流。广告将投放在“推荐”动态中，可能投放在“个人资料页”和“关注”推荐内容中。
	TiktokSubplacement_IN_FEED TiktokSubplacement = "IN_FEED"
	// TiktokSubplacement_SEARCH_FEED：搜索推荐内容。
	TiktokSubplacement_SEARCH_FEED TiktokSubplacement = "SEARCH_FEED"
	// TiktokSubplacement_TIKTOK_LITE：TikTok Lite，占用内容更小、视频加载速度更快的精简版 TikTok 应用。TikTok Lite 子版位仅支持定向日本或韩国时使用。
	TiktokSubplacement_TIKTOK_LITE TiktokSubplacement = "TIKTOK_LITE"
)
