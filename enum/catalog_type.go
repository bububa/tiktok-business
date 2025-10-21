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
	// CatalogType_MINI_SERIES：短剧商品库。
	CatalogType_MINI_SERIES CatalogType = "MINI_SERIES"
)
