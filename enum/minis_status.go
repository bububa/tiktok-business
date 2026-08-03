package enum

// MinisStatus TikTok Minis 状态。
type MinisStatus string

const (
	// MinisStatus_ACTIVE 表示 TikTok Minis 已启用。
	MinisStatus_ACTIVE MinisStatus = "ACTIVE"
	// MinisStatus_INACTIVE 表示 TikTok Minis 未启用。
	MinisStatus_INACTIVE MinisStatus = "INACTIVE"
)
