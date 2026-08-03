package enum

// VerificationIndustryCode is an industry category used for business verification.
type VerificationIndustryCode string

const (
	// VerificationIndustryCode_NOT_APPLICABLE indicates that no regulated industry applies.
	VerificationIndustryCode_NOT_APPLICABLE VerificationIndustryCode = "NOT_APPLICABLE"
	// VerificationIndustryCode_ALCOHOL indicates the alcohol industry.
	VerificationIndustryCode_ALCOHOL VerificationIndustryCode = "ALCOHOL"
	// VerificationIndustryCode_OTC indicates the over-the-counter medicine industry.
	VerificationIndustryCode_OTC VerificationIndustryCode = "OTC"
	// VerificationIndustryCode_DATING_APP indicates the dating app industry.
	VerificationIndustryCode_DATING_APP VerificationIndustryCode = "DATING_APP"
	// VerificationIndustryCode_FINANCIAL_SERVICES indicates the financial services industry.
	VerificationIndustryCode_FINANCIAL_SERVICES VerificationIndustryCode = "FINANCIAL_SERVICES"
)
