package enum

// ReviewStatus 审核状态
type ReviewStatus string

const (
	// ReviewStatus_INIT
	ReviewStatus_INIT ReviewStatus = "INIT"
	// ReviewStatus_REVIEWING
	ReviewStatus_REVIEWING ReviewStatus = "REVIEWING"
	// ReviewStatus_SUCCEED
	ReviewStatus_SUCCEED ReviewStatus = "SUCCEED"
	// ReviewStatus_FAIL
	ReviewStatus_FAIL ReviewStatus = "FAIL"
)
