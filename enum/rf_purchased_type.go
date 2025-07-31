package enum

// RfPurchasedType 覆盖和频次广告购买方式
type RfPurchasedType string

const (
	// RfPurchasedType_FIXED_SHOW 按固定展示量购买
	RfPurchasedType_FIXED_SHOW RfPurchasedType = "FIXED_SHOW"
	// RfPurchasedType_FIXED_REACH 按固定触达人数购买
	RfPurchasedType_FIXED_REACH RfPurchasedType = "FIXED_REACH"
)
