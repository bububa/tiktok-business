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
	// ReportType_VIDEO_INSIGHT 为一个或多个广告账号生成的视频洞察报表
	ReportType_VIDEO_INSIGHT ReportType = "VIDEO_INSIGHT"
	// ReportType_AD AD即将废弃：提供广告秒级表现数据的报表。
	ReportType_AD ReportType = "AD"
	// ReportType_VIDEO VIDEO：视频素材的视频洞察报表。
	ReportType_VIDEO ReportType = "VIDEO"
)
