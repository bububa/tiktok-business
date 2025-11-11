package enum

// GMVMaxBidType 时段类型
type GMVMaxBidType string

const (
	// GMVMaxBidType_NO_BID: 最大投放量时段。
	GMVMaxBidType_NO_BID GMVMaxBidType = "NO_BID"
	// GMVMaxBidType_CREATIVE_NO_BID: 创意作品加热时段
	GMVMaxBidType_CREATIVE_NO_BID GMVMaxBidType = "CREATIVE_NO_BID"
)
