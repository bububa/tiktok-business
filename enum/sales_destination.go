package enum

// SalesDestination 销售目标页面，即想要推动销售的渠道。
type SalesDestination string

const (
	// SalesDestination_TIKTOK_SHOP：TikTok Shop。推动 TikTok Shop 上的销售。
	SalesDestination_TIKTOK_SHOP SalesDestination = "TIKTOK_SHOP"
	// SalesDestination_WEBSITE：网站。推动网站上的销售。
	SalesDestination_WEBSITE SalesDestination = "WEBSITE"
	// SalesDestination_APP：应用。推动应用上的销售（需要商品库）。
	SalesDestination_APP SalesDestination = "APP"
	// SalesDestination_WEB_AND_APP：网站和应用。提升网站和应用上的销量
	SalesDestination_WEB_AND_APP SalesDestination = "WEB_AND_APP"
)
