package enum

// AuthCodeStatus 授权码的状态。
type AuthCodeStatus string

const (
	// AuthCodeStatus_NOT_USED：授权码已生成，但尚未使用。
	AuthCodeStatus_NOT_USED AuthCodeStatus = "NOT_USED"
	// AuthCodeStatus_IN_USE：授权码已生成，且已经使用。
	AuthCodeStatus_IN_USE AuthCodeStatus = "IN_USE"
	// AuthCodeStatus_EXPIRED：授权码已生成，但已经失效。
	AuthCodeStatus_EXPIRED AuthCodeStatus = "EXPIRED"
	// AuthCodeStatus_DELETED：授权码已删除
	AuthCodeStatus_DELETED AuthCodeStatus = "DELETED"
)
