package enum

// ShoppingAdsDeeplinkType 在购物广告中要使用的深度链接的来源。
type ShoppingAdsDeeplinkType string

const (
	// ShoppingAdsDeeplinkType_NONE：不设置深度链接。
	ShoppingAdsDeeplinkType_NONE ShoppingAdsDeeplinkType = "NONE"
	// ShoppingAdsDeeplinkType_CUSTOM：自定义深度链接。在广告中使用你提供的深度链接。
	ShoppingAdsDeeplinkType_CUSTOM ShoppingAdsDeeplinkType = "CUSTOM"
	// ShoppingAdsDeeplinkType_SHOPPING_ADS：商品库深度链接。使用商品库中为各商品添加的深度链接。
	ShoppingAdsDeeplinkType_SHOPPING_ADS ShoppingAdsDeeplinkType = "SHOPPING_ADS"
)
