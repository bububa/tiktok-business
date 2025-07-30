package enum

// BidType 出价策略
type BidType string

const (
	// BID_TYPE_CUSTOM 手动出价。在TikTok广告管理平台中，该出价策略被称为“成本上限”。
	BID_TYPE_CUSTOM BidType = "BID_TYPE_CUSTOM"
	// BID_TYPE_NO_BID 关闭出价控制，获取最佳的广告投放效果。在TikTok广告管理平台中，该出价策略被称为“最大投放量”。
	BID_TYPE_NO_BID BidType = "BID_TYPE_NO_BID"
)
