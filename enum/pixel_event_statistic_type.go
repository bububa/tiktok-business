package enum

// PixelEventStatisticType 统计类型
type PixelEventStatisticType string

const (
	// PixelEventStatisticType_EVERY_TIME（每一次）
	PixelEventStatisticType_EVERY_TIME PixelEventStatisticType = "EVERY_TIME"
	// PixelEventStatisticType_ONCE（仅一次）
	PixelEventStatisticType_ONCE PixelEventStatisticType = "ONCE"
)
