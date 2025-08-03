package enum

// IdentityType 认证身份类型
type IdentityType string

const (
	// IdentityType_CUSTOMIZED_USER (自定义用户）
	IdentityType_CUSTOMIZED_USER IdentityType = "CUSTOMIZED_USER"
	//  IdentityType_AUTH_CODE (授权码认证)
	IdentityType_AUTH_CODE IdentityType = "AUTH_CODE"
	// IdentityType_TT_USER (TikTok商务用户)
	IdentityType_TT_USER IdentityType = "TT_USER"
	//  IdentityType_BC_AUTH_TT（已添加到商务中心的TikTok用户）
	IdentityType_BC_AUTH_TT IdentityType = "BC_AUTH_TT"
	// IdentityType_CUSTOMIZED_USER
	IdentityType_CUSTOMIZED_USER IdentityType = "CUSTOMIZED_USER"
)
