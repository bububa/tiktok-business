package enum

// ProductAvailability 商店中该商品的当前库存情况。
type ProductAvailability string

const (
	// ProductAvailability_IN_STOCK：有库存。
	ProductAvailability_IN_STOCK ProductAvailability = "IN_STOCK"
	// ProductAvailability_AVAILABLE_FOR_ORDER：可订购。
	ProductAvailability_AVAILABLE_FOR_ORDER ProductAvailability = "AVAILABLE_FOR_ORDER"
	// ProductAvailability_PREORDER：可预订。
	ProductAvailability_PREORDER ProductAvailability = "PREORDER"
	// ProductAvailability_OUT_OF_STOCK：缺货。
	ProductAvailability_OUT_OF_STOCK ProductAvailability = "OUT_OF_STOCK"
	// ProductAvailability_DISCONTINUED：停售。
	ProductAvailability_DISCONTINUED ProductAvailability = "DISCONTINUED"
	// ProductAvailability_AVAILABLE
	ProductAvailability_AVAILABLE ProductAvailability = "AVAILABLE"
	// ProductAvailability_NOT_AVAILABLE
	ProductAvailability_NOT_AVAILABLE ProductAvailability = "NOT_AVAILABLE"
	// ProductAvailability_ON_DEMAND
	ProductAvailability_ON_DEMAND ProductAvailability = "ON_DEMAND"
	// ProductAvailability_SUBSCRIBERS_ONLY
	ProductAvailability_SUBSCRIBERS_ONLY ProductAvailability = "SUBSCRIBERS_ONLY"
)
