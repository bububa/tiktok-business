package enum

// TCMAuthCodeStatus 授权码状态。枚举值：
type TCMAuthCodeStatus string

const (
	// TCMAuthCodeStatus_NOT_USED：授权码已生成，但尚未使用。
	TCMAuthCodeStatus_NOT_USED TCMAuthCodeStatus = "NOT_USED"
	// TCMAuthCodeStatus_IN_USE：授权码已生成，且已经使用。
	TCMAuthCodeStatus_IN_USE TCMAuthCodeStatus = "IN_USE"
	// TCMAuthCodeStatus_EXPIRED：授权码已生成，但已经失效。
	TCMAuthCodeStatus_EXPIRED TCMAuthCodeStatus = "EXPIRED"
	// TCMAuthCodeStatus_DELETED：授权码已删除。
	TCMAuthCodeStatus_DELETED TCMAuthCodeStatus = "DELETED"
)
