package enum

// StoreProductGMVMaxAdsStatus 该商品在处于启用状态的商品 GMV Max 推广系列中的状态。
type StoreProductGMVMaxAdsStatus string

const (
	// StoreProductGMVMaxAdsStatus_OCCUPIED：该商品已被处于启用状态的商品 GMV Max 推广系列占用。
	StoreProductGMVMaxAdsStatus_OCCUPIED StoreProductGMVMaxAdsStatus = "OCCUPIED"
	// StoreProductGMVMaxAdsStatus_UNOCCUPIED: 该商品未被处于启用状态的商品 GMV Max 推广系列占用。
	StoreProductGMVMaxAdsStatus_UNOCCUPIED StoreProductGMVMaxAdsStatus = "UNOCCUPIED"
)
