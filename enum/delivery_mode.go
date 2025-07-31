package enum

// DeliveryMode 覆盖和频次广告组中广告投放的排序与排期策略。
type DeliveryMode string

const (
	// DeliveryMode_STANDARD：标准投放。每个广告均匀投放，预计会获得大致相同的访问量。
	DeliveryMode_STANDARD DeliveryMode = "STANDARD"
	// DeliveryMode_SCHEDULE：计划投放。为每个广告设置投放的特定时间段。
	DeliveryMode_SCHEDULE DeliveryMode = "SCHEDULE"
	// DeliveryMode_SEQUENCE：顺序投放。为广告设置特定的投放顺序。
	DeliveryMode_SEQUENCE DeliveryMode = "SEQUENCE"
	// DeliveryMode_OPTIMIZE：优选投放。
	DeliveryMode_OPTIMIZE DeliveryMode = "OPTIMIZE"
)
