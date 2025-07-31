package enum

// AttributionEventCount 广告组的事件统计方式（统计类型）。
// 统计用户在查看或点击了广告后采取操作的次数的方式。
type AttributionEventCount string

const (
	// AttributionEventCount_UNSET ：未设置。
	AttributionEventCount_UNSET AttributionEventCount = "UNSET"
	// AttributionEventCount_EVERY：每一次。某位用户进行的多个购买事件将分别计入转化量。
	AttributionEventCount_EVERY AttributionEventCount = "EVERY"
	// AttributionEventCount_ONCE：仅一次。某位用户进行的多个事件将只算作 1 次转化。
	AttributionEventCount_ONCE AttributionEventCount = "ONCE"
)
