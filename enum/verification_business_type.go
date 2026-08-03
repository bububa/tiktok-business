package enum

// VerificationBusinessType is the type of account verification to perform.
type VerificationBusinessType string

const (
	// VerificationBusinessType_BUSINESS performs business verification.
	VerificationBusinessType_BUSINESS VerificationBusinessType = "BUSINESS"
	// VerificationBusinessType_INDIVIDUAL performs individual identity verification.
	VerificationBusinessType_INDIVIDUAL VerificationBusinessType = "INDIVIDUAL"
)
