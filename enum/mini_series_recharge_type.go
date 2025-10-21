package enum

// MiniSeriesType 短剧收费类型。
type MiniSeriesType string

const (
	// MiniSeriesType_BY_TIERS: 按级别收费。
	MiniSeriesType_BY_TIERS MiniSeriesType = "BY_TIERS"
	// MiniSeriesType_SUBSCRIPTION: 按周期订阅收费。
	MiniSeriesType_SUBSCRIPTION MiniSeriesType = "SUBSCRIPTION"
	// MiniSeriesType_BY_EPISODES: 按集数收费。
	MiniSeriesType_BY_EPISODES MiniSeriesType = "BY_EPISODES"
)
