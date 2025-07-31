package enum

// PacingMode 广告投放速度类型。
type PacingMode string

const (
	// PACING_MODE_SMOOTH 在预定的时间内平均分配预算
	PACING_MODE_SMOOTH PacingMode = "PACING_MODE_SMOOTH"
	// PACING_MODE_FAST 尽快消耗预算并产出结果
	PACING_MODE_FAST PacingMode = "PACING_MODE_FAST"
)
