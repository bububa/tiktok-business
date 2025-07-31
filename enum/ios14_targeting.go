package enum

// Ios14Targeting 您想定向的 iOS 设备版本。
type Ios14Targeting string

const (
	// Ios14Targeting_UNSET：运行 iOS 14.4 及更早版本的设备。
	Ios14Targeting_UNSET Ios14Targeting = "UNSET"
	// Ios14Targeting_IOS14_MINUS：不受 iOS 14 新的 ATT 框架影响的 iOS 设备（运行 iOS 14.0 及更早版本的设备）。
	Ios14Targeting_IOS14_MINUS Ios14Targeting = "IOS14_MINUS"
	// Ios14Targeting_IOS14_PLUS：运行iOS 14.5 及之后版本、已执行 ATT 框架的 iOS 设备。此值仅支持在专属推广系列中使用。
	Ios14Targeting_IOS14_PLUS Ios14Targeting = "IOS14_PLUS"
	// Ios14Targeting_ALL：运行 iOS 14.5 及之后版本、已执行 ATT 框架的 iOS 设备。此值仅支持在 iOS 应用再营销广告或 iOS 重定向类型的视频购物广告（商品来源为商品库且优化位置为应用）中使用。
	Ios14Targeting_ALL Ios14Targeting = "ALL"
)
