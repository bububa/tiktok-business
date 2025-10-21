package enum

// IntervalType 定时更新间隔单位。
type IntervalType string

const (
	// IntervalType_HOURLY（小时）
	IntervalType_HOURLY IntervalType = "HOURLY"
	// IntervalType_DAILY（日）
	IntervalType_DAILY IntervalType = "DAILY"
	// IntervalType_MONTHLY（月）。
	IntervalType_MONTHLY IntervalType = "MONTHLY"
)
