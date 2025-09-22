package enum

// DiagnosisIssueCategory 问题类别。
type DiagnosisIssueCategory string

const (
	// DiagnosisIssueCategory_CREATIVE：创意相关问题。
	DiagnosisIssueCategory_CREATIVE DiagnosisIssueCategory = "CREATIVE"
	// DiagnosisIssueCategory_BID_AND_BUDGET：出价和预算相关问题。
	DiagnosisIssueCategory_BID_AND_BUDGET DiagnosisIssueCategory = "BID_AND_BUDGET"
	// DiagnosisIssueCategory_EVENT_TRACK：事件追踪相关问题。
	DiagnosisIssueCategory_EVENT_TRACK DiagnosisIssueCategory = "EVENT_TRACK"
)
