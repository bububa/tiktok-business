package enum

// PageStatus 落地页状态
type PageStatus string

const (
	// PageStatus_DRAFT（草稿）
	PageStatus_DRAFT PageStatus = "DRAFT"
	// PageStatus_READY（可用）
	PageStatus_READY PageStatus = "READY"
)
