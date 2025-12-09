package enum

// BusinessAutoMessageAuditStatus The review status of the automatic message.
type BusinessAutoMessageAuditStatus string

const (
	// BusinessAutoMessageAuditStatus_REVIEWING: The automatic message is pending review.
	BusinessAutoMessageAuditStatus_REVIEWING BusinessAutoMessageAuditStatus = "REVIEWING"
	// BusinessAutoMessageAuditStatus_APPROVED: The automatic message has been approved.
	BusinessAutoMessageAuditStatus_APPROVED BusinessAutoMessageAuditStatus = "APPROVED"
	// BusinessAutoMessageAuditStatus_REJECTED: The automatic message has been rejected.
	BusinessAutoMessageAuditStatus_REJECTED BusinessAutoMessageAuditStatus = "REJECTED"
)
