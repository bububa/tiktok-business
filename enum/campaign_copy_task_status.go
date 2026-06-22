package enum

// CampaignCopyTaskStatus The status of the asynchronous campaign copy task.
type CampaignCopyTaskStatus string

const (
	// CampaignCopyTaskStatus_RUNNING: The task is being processed.
	CampaignCopyTaskStatus_RUNNING CampaignCopyTaskStatus = "RUNNING"
	// CampaignCopyTaskStatus_SUCCESS: The task has been processed. Check the task_result to see if the task has succeeded.
	CampaignCopyTaskStatus_SUCCESS CampaignCopyTaskStatus = "SUCCESS"
	// CampaignCopyTaskStatus_FAILURE: The task fails to be processed.
	CampaignCopyTaskStatus_FAILURE CampaignCopyTaskStatus = "FAILURE"
)
