package enum

// BidBudgetStatus 出价和预算分析的结果。
type BidBudgetStatus string

const (
	// BidBudgetStatus_GOOD：良好。
	BidBudgetStatus_GOOD BidBudgetStatus = "GOOD"
	// BidBudgetStatus_LOW_BID_AND_BUDGET：出价和预算过低。
	BidBudgetStatus_LOW_BID_AND_BUDGET BidBudgetStatus = "LOW_BID_AND_BUDGET"
	// BidBudgetStatus_LOW_BUDGET：预算较低。
	BidBudgetStatus_LOW_BUDGET BidBudgetStatus = "LOW_BUDGET"
	// BidBudgetStatus_LOW_BID：出价较低。
	BidBudgetStatus_LOW_BID BidBudgetStatus = "LOW_BID"
	// BidBudgetStatus_NO_DATA：暂无数据。
	BidBudgetStatus_NO_DATA BidBudgetStatus = "NO_DATA"
)
