package enum

// TargetingOptimizationMode Audience targeting optimization mode.
type TargetingOptimizationMode string

const (
	// TargetingOptimizationMode_MANUAL: Custom targeting. You can use custom targeting settings to precisely control who sees your ads. This may limit delivery and impact campaign performance.
	TargetingOptimizationMode_MANUAL TargetingOptimizationMode = "MANUAL"
	// TargetingOptimizationMode_AUTOMATIC: Automatic targeting. You can use automatic targeting to leverage real-time data and machine learning to target audiences most likely to engage with your ads.
	TargetingOptimizationMode_AUTOMATIC TargetingOptimizationMode = "AUTOMATIC"
)
