package enum

// CatalogType 商品库类型。
type CatalogType string

const (
	// CatalogType_ECOM：电商商品库。
	CatalogType_ECOM CatalogType = "ECOM"
	// CatalogType_HOTEL：酒店商品库。
	CatalogType_HOTEL CatalogType = "HOTEL"
	// CatalogType_FLIGHT：航班商品库。
	CatalogType_FLIGHT CatalogType = "FLIGHT"
	// CatalogType_DESTINATION：目的地商品库。
	CatalogType_DESTINATION CatalogType = "DESTINATION"
	// CatalogType_ENTERTAINMENT：娱乐商品库。
	CatalogType_ENTERTAINMENT CatalogType = "ENTERTAINMENT"
	// CatalogType_AUTO_VEHICLE: Auto-Inventory catalog.
	CatalogType_AUTO_VEHICLE CatalogType = "AUTO_VEHICLE"
	// CatalogType_AUTO_MODEL: Auto-Model catalog.
	CatalogType_AUTO_MODEL CatalogType = "AUTO_MODEL"
	// CatalogType_MINI_SERIES：短剧商品库。
	CatalogType_MINI_SERIES CatalogType = "MINI_SERIES"
	// CatalogType_GENERIC: generic catalog.
	CatalogType_GENERIC CatalogType = "GENERIC"
	// CatalogType_ONLINE_TO_OFFLINE: online-to-offline catalog.
	CatalogType_ONLINE_TO_OFFLINE CatalogType = "ONLINE_TO_OFFLINE"
	// CatalogType_ECOMMERCE: e-commerce
	CatalogType_ECOMMERCE CatalogType = "ECOMMERCE"
	// CatalogType_TRAVEL_ENTERTAINMENT: travel and entertainment.
	CatalogType_TRAVEL_ENTERTAINMENT CatalogType = "TRAVEL_ENTERTAINMENT"
)
