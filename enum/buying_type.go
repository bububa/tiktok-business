package enum

// BuyingTypes 购买类型。
type BuyingType string

const (
	// BuyingType_AUCTION：竞价广告。
	BuyingType_AUCTION BuyingType = "AUCTION"
	// BuyingType_RESERVATION_RF：合约覆盖和频次广告以及 TikTok Pulse 广告。
	BuyingType_RESERVATION_RF BuyingType = "RESERVATION_RF"
	// BuyingType_RESERVATION_TOP_VIEW：合约 TopView 广告。
	BuyingType_RESERVATION_TOP_VIEW BuyingType = "RESERVATION_TOP_VIEW"
)
