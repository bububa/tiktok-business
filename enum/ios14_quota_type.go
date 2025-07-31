package enum

// Ios14QuotaType 是否占用iOS14专属推广系列配额。
type Ios14QuotaType string

const (
	// Ios14QuotaType_OCCUPIED（占用）
	Ios14QuotaType_OCCUPIED Ios14QuotaType = "OCCUPIED"
	// Ios14QuotaType_UNOCCUPIED（不占用）
	Ios14QuotaType_UNOCCUPIED Ios14QuotaType = "UNOCCUPIED"
)
