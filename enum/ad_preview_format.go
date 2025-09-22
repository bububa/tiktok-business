package enum

// AdPreviewFormat 不同placement的预览格式有所不同：
type AdPreviewFormat string

const (
	// AdPreviewFormat_IN_FEED：信息流。广告将投放在“推荐”动态中，可能投放在“个人资料页”和“关注”推荐内容中。
	AdPreviewFormat_IN_FEED AdPreviewFormat = "IN_FEED"
	// AdPreviewFormat_SEARCH_RESULTS：搜索结果页面。
	AdPreviewFormat_SEARCH_RESULTS AdPreviewFormat = "SEARCH_RESULTS"
	// AdPreviewFormat_SEARCH_FEED：搜索推荐内容。
	AdPreviewFormat_SEARCH_FEED AdPreviewFormat = "SEARCH_FEED"
	// AdPreviewFormat_TIKTOK_LITE：TikTok Lite，占用内容更小、视频加载速度更快的精简版 TikTok 应用。TikTok Lite 子版位仅支持定向日本或韩国时使用。
	AdPreviewFormat_TIKTOK_LITE AdPreviewFormat = "TIKTOK_LITE"
	// AdPreviewFormat_PRODUCT_SEARCH: 搜索结果页面（Shopping Center）。在“商城”选项卡的搜索结果页面以及 TikTok 应用的常规搜索结果页面中展示您的商品。了解关于可投放店铺广告的“商城”选项卡广告位的详情。
	AdPreviewFormat_PRODUCT_SEARCH AdPreviewFormat = "PRODUCT_SEARCH"
	// AdPreviewFormat_PRODUCT_SHOP_CENTER: “为你推荐”板块（Shopping Center）。在“为你推荐”板块，将您的商品展示给可能会购买的用户。
	AdPreviewFormat_PRODUCT_SHOP_CENTER AdPreviewFormat = "PRODUCT_SHOP_CENTER"
	// AdPreviewFormat_INTERSTITIAL
	AdPreviewFormat_INTERSTITIAL AdPreviewFormat = "INTERSTITIAL"
	// AdPreviewFormat_REWARDED
	AdPreviewFormat_REWARDED AdPreviewFormat = "REWARDED"
	// AdPreviewFormat_APP_OPEN
	AdPreviewFormat_APP_OPEN AdPreviewFormat = "APP_OPEN"
	// AdPreviewFormat_NORMAL_BANNER
	AdPreviewFormat_NORMAL_BANNER AdPreviewFormat = "NORMAL_BANNER"
	// AdPreviewFormat_VIDEO_THUMBNAIL_BANNER
	AdPreviewFormat_VIDEO_THUMBNAIL_BANNER AdPreviewFormat = "VIDEO_THUMBNAIL_BANNER"
	// AdPreviewFormat_SMALL_VIDEO_BANNER
	AdPreviewFormat_SMALL_VIDEO_BANNER AdPreviewFormat = "SMALL_VIDEO_BANNER"
	// AdPreviewFormat_ICON_ONLY_BANNER
	AdPreviewFormat_ICON_ONLY_BANNER AdPreviewFormat = "ICON_ONLY_BANNER"
	// AdPreviewFormat_NORMAL_NATIVE
	AdPreviewFormat_NORMAL_NATIVE AdPreviewFormat = "NORMAL_NATIVE"
	// AdPreviewFormat_VIDEO_THUMBNAIL_NATIVE
	AdPreviewFormat_VIDEO_THUMBNAIL_NATIVE AdPreviewFormat = "VIDEO_THUMBNAIL_NATIVE"
)
