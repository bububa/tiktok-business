package enum

// BidDisplayMode 每次浏览成本的计算和测算模式。
type BidDisplayMode string

const (
	// BidDisplayMode_CPMV 每千次浏览成本
	BidDisplayMode_CPMV BidDisplayMode = "CPMV"
	// BidDisplayMode_CPV 每次浏览成本
	BidDisplayMode_CPV BidDisplayMode = "CPV"
)
