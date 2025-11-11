package enum

// BcStatus 商务中心状态
type BcStatus string

const (
	// BcStatus_REVIEWING(审核中)
	BcStatus_REVIEWING BcStatus = "REVIEWING"
	// BcStatus_DENY(已拒绝)
	BcStatus_DENY BcStatus = "DENY"
	// BcStatus_ENABLE(已通过)
	BcStatus_ENABLE BcStatus = "ENABLE"
	// BcStatus_PUNISH(已禁用)
	BcStatus_PUNISH BcStatus = "PUNISH"
)
