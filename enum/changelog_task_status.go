package enum

// ChangelogTaskStatus Task status
type ChangelogTaskStatus string

const (
	ChangelogTaskStatus_PROCESSING ChangelogTaskStatus = "PROCESSING"
	ChangelogTaskStatus_SUCCESS    ChangelogTaskStatus = "SUCCESS"
	ChangelogTaskStatus_FAILED     ChangelogTaskStatus = "FAILED"
)
