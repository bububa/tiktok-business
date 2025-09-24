package enum

// CustomConversionActiveStatus 自定义转化的活跃状态。
type CustomConversionActiveStatus string

const (
	// CustomConversionActiveStatus_NO_RECENT_ACTIVITY: 近期无活动。已收到回传活动，但并非是在过去 7 天内收到。
	CustomConversionActiveStatus_NO_RECENT_ACTIVITY CustomConversionActiveStatus = "NO_RECENT_ACTIVITY"
	// CustomConversionActiveStatus_ACTIVE: 活跃。在过去 7 天内收到回传活动。
	CustomConversionActiveStatus_ACTIVE CustomConversionActiveStatus = "ACTIVE"
	// CustomConversionActiveStatus_WAITING_FOR_ACTIVITY: 不活跃。过去 90 天内未收到回传活动。
	CustomConversionActiveStatus_WAITING_FOR_ACTIVITY CustomConversionActiveStatus = "WAITING_FOR_ACTIVITY"
)
