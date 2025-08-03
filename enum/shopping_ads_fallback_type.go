package enum

// ShoppingAdsFallbackType 在购物广告再营销场景下，深度链接唤起失败后的返回落地页类型。
type ShoppingAdsFallbackType string

const (
	// ShoppingAdsFallbackType_DEFAULT：未设置。
	ShoppingAdsFallbackType_DEFAULT ShoppingAdsFallbackType = "DEFAULT"
	// ShoppingAdsFallbackType_CUSTOM: 自定义链接。您可以在广告组中提供一个网页地址, 并回退到这个地址。
	ShoppingAdsFallbackType_CUSTOM ShoppingAdsFallbackType = "CUSTOM"
	// ShoppingAdsFallbackType_SHOPPING_ADS：商品库链接。 回退到目录中为各商品添加的网页地址。
	ShoppingAdsFallbackType_SHOPPING_ADS ShoppingAdsFallbackType = "SHOPPING_ADS"
)
