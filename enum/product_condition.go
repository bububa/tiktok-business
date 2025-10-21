package enum

	// ProductCondition 商店中商品的当前状况。
type ProductCondition string

const (
	// ProductCondition_NEW：全新。
	ProductCondition_NEW ProductCondition = "NEW"
	// ProductCondition_REFURBISHED：翻新。
	ProductCondition_REFURBISHED ProductCondition = "REFURBISHED"
	// ProductCondition_USED：二手。
ProductCondition_USED ProductCondition = "USED"
)
