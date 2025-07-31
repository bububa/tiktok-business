package enum

// DiscountType 预算折扣类型。
type DiscountType string

const (
	// DiscountType_NO_DISCOUNT：无折扣。
	DiscountType_NO_DISCOUNT DiscountType = "NO_DISCOUNT"
	// DiscountType_BY_PERCENTAGE：百分比折扣。
	DiscountType_BY_PERCENTAGE DiscountType = "BY_PERCENTAGE"
	// DiscountType_BY_AMOUNT：固定金额折扣。
	DiscountType_BY_AMOUNT DiscountType = "BY_AMOUNT"
)
