package enum

// LeadTaskStatus 线索下载任务状态
type LeadTaskStatus string

const (
	// LeadTaskStatus_CREATED
	LeadTaskStatus_CREATED LeadTaskStatus = "CREATED"
	// LeadTaskStatus_FAILED
	LeadTaskStatus_FAILED LeadTaskStatus = "FAILED"
	// LeadTaskStatus_RUNNING
	LeadTaskStatus_RUNNING LeadTaskStatus = "RUNNING"
	// LeadTaskStatus_SUCCEED
	LeadTaskStatus_SUCCEED LeadTaskStatus = "SUCCEED"
)
