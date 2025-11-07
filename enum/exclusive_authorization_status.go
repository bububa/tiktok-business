package enum

// ExclusiveAuthorizationStatus 该广告账户使用此 TikTok Shop 创建 GMV Max 推广系列的专属授权状态。
type ExclusiveAuthorizationStatus string

const (
	// ExclusiveAuthorizationStatus_EFFECTIVE：该广告账号目前拥有 GMV Max 专属授权。
	ExclusiveAuthorizationStatus_EFFECTIVE ExclusiveAuthorizationStatus = "EFFECTIVE"
	// ExclusiveAuthorizationStatus_INEFFECTIVE：该广告账号拥有过 GMV Max 专属授权，但专属授权已取消。
	ExclusiveAuthorizationStatus_INEFFECTIVE ExclusiveAuthorizationStatus = "INEFFECTIVE"
	// ExclusiveAuthorizationStatus_UNAUTHORIZED：该广告账号尚未拥有 GMV Max 专属授权
	ExclusiveAuthorizationStatus_UNAUTHORIZED ExclusiveAuthorizationStatus = "UNAUTHORIZED"
)
