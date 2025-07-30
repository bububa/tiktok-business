package enum

// BidAlignType 专属推广系列的归因类型。此类型决定广告组层级设置的目标 CPA （conversion_bid_price） 或目标 ROAS（roas_bid） 基于的归因网络类型。
type BidAlignType string

const (
	// BidAlignType_SAN：自归因平台归因。
	BidAlignType_SAN BidAlignType = "SAN"
	// BidAlignType_SKAN：SKAN 归因。
	BidAlignType_SKAN BidAlignType = "SKAN"
)
