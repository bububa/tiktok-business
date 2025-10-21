package enum

// StoreProductStatus 商品的状态。
type StoreProductStatus string

const (
	// StoreProductStatus_AVAILABLE：商品可用于创建广告。
	StoreProductStatus_AVAILABLE StoreProductStatus = "AVAILABLE"
	// StoreProductStatus_NOT_AVAILABLE：商品不可用于创建广告。
	StoreProductStatus_NOT_AVAILABLE StoreProductStatus = "NOT_AVAILABLE"
)
