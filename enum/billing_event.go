package enum

// BillingEvent 计费事件
type BillingEvent string

const (
	// CPC 点击计费
	CPC BillingEvent = "CPC"
	// CPM 每千人计费
	CPM BillingEvent = "CPM"
	// CPV 有效播放计费
	CPV BillingEvent = "CPV"
	// OCPC 已废弃
	// oCPC（CPA 即 oCPC）
	OCPC BillingEvent = "OCPC"
	// GD 保量
	GD BillingEvent = "GD"
	// OCPM oCPM
	OCPM BillingEvent = "OCPM"
)
