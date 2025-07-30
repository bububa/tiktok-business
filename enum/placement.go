package enum

// Placement 版位
type Placement string

const (
	// PLACEMENT_TIKTOK TikTok
	PLACEMENT_TIKTOK Placement = "PLACEMENT_TIKTOK"
	// PLACEMENT_PANGLE Pangle（原TikTok_Audience_Network）
	PLACEMENT_PANGLE Placement = "PLACEMENT_PANGLE"
	// PLACEMENT_GLOBAL_APP_BUNDLE 全球化应用组合。目前包括Capcut和Fizzo。
	// 注意：
	// 全球化应用组合版位目前不支持优化目标落地页浏览量（optimization_goal=TRAFFIC_LANDING_PAGE_VIEW）。
	PLACEMENT_GLOBAL_APP_BUNDLE Placement = "PLACEMENT_GLOBAL_APP_BUNDLE"
	// PLACEMENT_TOPBUZZ 已废弃 BuzzVideo/Babe
	PLACEMENT_TOPBUZZ Placement = "PLACEMENT_TOPBUZZ"
	// PLACEMENT_HELO 已废弃 Helo
	PLACEMENT_HELO Placement = "PLACEMENT_HELO"
)
