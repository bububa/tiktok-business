package enum

// IdentityAvailableStatus 认证身份的可用状态
type IdentityAvailableStatus string

const (
	// IdentityAvailableStatus_AVAILABLE
	IdentityAvailableStatus_AVAILABLE IdentityAvailableStatus = "AVAILABLE"
	// IdentityAvailableStatus_NO_VALID_BIND_ACCOUNT
	IdentityAvailableStatus_NO_VALID_BIND_ACCOUNT IdentityAvailableStatus = "NO_VALID_BIND_ACCOUNT"
	// IdentityAvailableStatus_SCOPE_UNAVAILABLE
	IdentityAvailableStatus_SCOPE_UNAVAILABLE IdentityAvailableStatus = "SCOPE_UNAVAILABLE"
	// IdentityAvailableStatus_IS_PRIVATE_ACCOUNT
	IdentityAvailableStatus_IS_PRIVATE_ACCOUNT IdentityAvailableStatus = "IS_PRIVATE_ACCOUNT"
	// IdentityAvailableStatus_NOT_BUSINESS_ACCOUNT
	IdentityAvailableStatus_NOT_BUSINESS_ACCOUNT IdentityAvailableStatus = "NOT_BUSINESS_ACCOUNT"
)
