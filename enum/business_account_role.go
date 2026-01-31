package enum

// BusinessAccountRole Business Center user's Business Account management permission to the TikTok account.
type BusinessAccountRole string

const (
	// BUSINESS_ACCOUNT_ADMIN: Business Account admin. Business Account admins have access to all Business Account features, including verification, member management, and linking to Business Center.
	BUSINESS_ACCOUNT_ADMIN BusinessAccountRole = "BUSINESS_ACCOUNT_ADMIN"
	// BUSINESS_ACCOUNT_OPERATOR: Business Account operator. Business Account operators can manage TikTok analytics, posts, messages, and profile information.
	BUSINESS_ACCOUNT_OPERATOR BusinessAccountRole = "BUSINESS_ACCOUNT_OPERATOR"
	// BUSINESS_ACCOUNT_ANALYST : Business Account analyst. Business Account analysts only have access to analytics and are not able to manage posts, messages, etc.
	BUSINESS_ACCOUNT_ANALYST BusinessAccountRole = "BUSINESS_ACCOUNT_ANALYST"
)
