package enum

// AudienceSubType 受众子类型，表明可以使用的广告类型。
type AudienceSubType string

const (
	// AudienceSubType_NORMAL: 常规受众，可用于非覆盖和频次广告
	AudienceSubType_NORMAL AudienceSubType = "NORMAL"
	// AudienceSubType_REACH_FREQUENCY：覆盖和频次广告受众，只可用于覆盖和频次广告
	AudienceSubType_REACH_FREQUENCY AudienceSubType = "REACH_FREQUENCY"
)
