package enum

// TTVideoAuthStatus 授权状态。枚举值：
type TTVideoAuthStatus string

const (
	// TTVideoAuthStatus_WAITING：创作者尚未批准或拒绝您的授权申请。
	TTVideoAuthStatus_WAITING TTVideoAuthStatus = "WAITING"
	// TTVideoAuthStatus_REJECTED：创作者已拒绝您的授权申请。
	TTVideoAuthStatus_REJECTED TTVideoAuthStatus = "REJECTED"
	// TTVideoAuthStatus_ACCEPTED：创作者已批准您的授权申请。
	TTVideoAuthStatus_ACCEPTED TTVideoAuthStatus = "ACCEPTED"
)
