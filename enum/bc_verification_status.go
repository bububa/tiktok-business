package enum

	// BcVerificationStatus Business verification status of the Business Center.
type BcVerificationStatus string

const (
// BcVerificationStatus_NOT_SUBMITTED: The business verification has not been submitted.
	BcVerificationStatus_NOT_SUBMITTED BcVerificationStatus = "NOT_SUBMITTED"
	// BcVerificationStatus_REVIEWING: The business verification is currently under review.
	BcVerificationStatus_REVIEWING BcVerificationStatus = "REVIEWING"
	// BcVerificationStatus_VERIFIED: The business has been verified successfully.
	BcVerificationStatus_VERIFIED BcVerificationStatus = "VERIFIED"
	// BcVerificationStatus_FAILED: The business verification has failed.
	BcVerificationStatus_FAILED BcVerificationStatus = "FAILED"
	// BcVerificationStatus_EXPIRED: The business verification has expired.
	BcVerificationStatus_EXPIRED BcVerificationStatus = "EXPIRED"
)
