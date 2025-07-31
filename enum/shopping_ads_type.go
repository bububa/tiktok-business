package enum

// ShoppingAdsType  购物广告类型。
type ShoppingAdsType string

const (
	// ShoppingAdsType_VIDEO：视频购物广告。
	ShoppingAdsType_VIDEO ShoppingAdsType = "VIDEO"
	// ShoppingAdsType_LIVE：直播购物广告。
	ShoppingAdsType_LIVE ShoppingAdsType = "LIVE"
	// ShoppingAdsType_PRODUCT_SHOPPING_ADS：商品购物广告。
	ShoppingAdsType_PRODUCT_SHOPPING_ADS ShoppingAdsType = "PRODUCT_SHOPPING_ADS"
	// ShoppingAdsType_CATALOG_LISTING_ADS：产品目录广告。
	ShoppingAdsType_CATALOG_LISTING_ADS ShoppingAdsType = "CATALOG_LISTING_ADS"
	// ShoppingAdsType_UNSET：未设置。
	ShoppingAdsType_UNSET ShoppingAdsType = "UNSET"
)
