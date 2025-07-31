package enum

// ScheduleType 广告投放时间类型。
type ScheduleType string

const (
	// SCHEDULE_FROM_NOW 只需要明确投放开始时间
	SCHEDULE_FROM_NOW ScheduleType = "SCHEDULE_FROM_NOW"
	// SCHEDULE_START_END 需要明确投放开始和结束时间
	SCHEDULE_START_END ScheduleType = "SCHEDULE_START_END"
)
