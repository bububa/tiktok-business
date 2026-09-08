package enum

// KeywordAuditStatus The review status of the search keyword.
type KeywordAuditStatus string

const (
	// KeywordAuditStatus_AUDITING: The keyword is under review.
	KeywordAuditStatus_AUDITING KeywordAuditStatus = "AUDITING"
	// KeywordAuditStatus_PASS: The keyword has passed review and can be delivered.
	KeywordAuditStatus_PASS KeywordAuditStatus = "PASS"
	// KeywordAuditStatus_REJECTED: The keyword failed to pass the review and cannot be delivered.
	KeywordAuditStatus_REJECTED KeywordAuditStatus = "REJECTED"
)
