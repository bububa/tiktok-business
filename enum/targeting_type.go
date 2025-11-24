package enum

// TargetingType 定向标签所适用的定向类型。
type TargetingType string

const (
	// TargetingType_GEO (地理位置定向，包括行政区域定向和邮政编码定向)
	TargetingType_GEO TargetingType = "GEO"
	// TargetingType_ISP ：互联网服务提供商定向。
	TargetingType_ISP TargetingType = "ISP"
)
