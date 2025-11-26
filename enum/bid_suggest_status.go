package enum

// BidSuggestStatus 针对出价的诊断结果。
type BidSuggestStatus string

const (
	// BidSuggestStatus_GOOD：良好。
	BidSuggestStatus_GOOD BidSuggestStatus = "GOOD"
	// BidSuggestStatus_LOW：出价较低。
	BidSuggestStatus_LOW BidSuggestStatus = "LOW"
	// BidSuggestStatus_NO_DATA：暂无数据。
	BidSuggestStatus_NO_DATA BidSuggestStatus = "NO_DATA"
)
