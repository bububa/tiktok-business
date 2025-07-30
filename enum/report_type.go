package enum

// ReportType 报表类型
type ReportType string

const (
	// ReportType_BASIC：基础报表。
	ReportType_BASIC ReportType = "BASIC"
	// ReportType_AUDIENCE：受众分析报表。
	ReportType_AUDIENCE ReportType = "AUDIENCE"
	// ReportType_PLAYABLE_MATERIAL ：试玩素材报表。
	ReportType_PLAYABLE_MATERIAL ReportType = "PLAYABLE_MATERIAL"
	// ReportType_CATALOG：DSA 报表。
	ReportType_CATALOG ReportType = "CATALOG"
	// ReportType_BC ：商务中心报表。
	ReportType_BC ReportType = "BC"
	// ReportType_TT_SHOP：最大成交额广告报表。
	ReportType_TT_SHOP ReportType = "TT_SHOP"
)
