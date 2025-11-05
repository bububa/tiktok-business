package enum

// TTVideoReviewStatus 审核状态
type TTVideoReviewStatus string

const (
	// TTVideoReviewStatus_NO_REVIEW
	TTVideoReviewStatus_NO_REVIEW TTVideoReviewStatus = "NO_REVIEW"
	// TTVideoReviewStatus_CONTENT_REVIEWING
	TTVideoReviewStatus_CONTENT_REVIEWING TTVideoReviewStatus = "CONTENT_REVIEWING"
	// TTVideoReviewStatus_AD_REVIEWING
	TTVideoReviewStatus_AD_REVIEWING TTVideoReviewStatus = "AD_REVIEWING"
	// TTVideoReviewStatus_CONTENT_REVIEW_FAIL
	TTVideoReviewStatus_CONTENT_REVIEW_FAIL TTVideoReviewStatus = "CONTENT_REVIEW_FAIL"
	// TTVideoReviewStatus_AD_REVIEW_FAIL
	TTVideoReviewStatus_AD_REVIEW_FAIL TTVideoReviewStatus = "AD_REVIEW_FAIL"
	// TTVideoReviewStatus_SUCCEED
	TTVideoReviewStatus_SUCCEED TTVideoReviewStatus = "SUCCEED"
	// TTVideoReviewStatus_ALL_FAIL
	TTVideoReviewStatus_ALL_FAIL TTVideoReviewStatus = "ALL_FAIL"
)
