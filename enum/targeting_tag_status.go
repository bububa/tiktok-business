package enum

// TargetingTagStatus 定向标签的状态。
type TargetingTagStatus string

const (
	// TargetingTagStatus_ENABLED（可用于定向）
	TargetingTagStatus_ENABLED TargetingTagStatus = "ENABLED"
	//  TargetingTagStatus_DISABLED（不可用于定向）
	TargetingTagStatus_DISABLED TargetingTagStatus = "DISABLED"
)
