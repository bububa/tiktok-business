package enum

// ReportTaskStatus 任务状态。
type ReportTaskStatus string

const (
	// ReportTaskStatus_QUEUING（队列中）
	ReportTaskStatus_QUEUING ReportTaskStatus = "QUEUING"
	// ReportTaskStatus_PROCESSING（ 处理中）
	ReportTaskStatus_PROCESSING ReportTaskStatus = "PROCESSING"
	// ReportTaskStatus_SUCCESS（ 任务成功）
	ReportTaskStatus_SUCCESS ReportTaskStatus = "SUCCESS"
	// ReportTaskStatus_FAILED（任务失败）
	ReportTaskStatus_FAILED ReportTaskStatus = "FAILED"
	// ReportTaskStatus_CANCELED（任务已取消）。
	ReportTaskStatus_CANCELED ReportTaskStatus = "CANCELED"
)
