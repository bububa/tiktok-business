package enum

// ProductAuditStatus 审核状态
type ProductAuditStatus string

const (
	// ProductAuditStatusApproved approved（通过）
	ProductAuditStatusApproved ProductAuditStatus = "approved"
	// ProductAuditStatusRejected rejected（未通过）
	ProductAuditStatusRejected ProductAuditStatus = "rejected"
	// ProductAuditStatusProcessing processing（处理中）。
	ProductAuditStatusProcessing ProductAuditStatus = "processing"
)
