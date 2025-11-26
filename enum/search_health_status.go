package enum

// SearchHealthStatus 整体搜索健康状况
type SearchHealthStatus string

const (
	// SearchHealthStatus_GOOD：良好。您的设置有助于提升广告成效。
	SearchHealthStatus_GOOD SearchHealthStatus = "GOOD"
	// SearchHealthStatus_NEED_IMPROVEMENT：需要改进。
	SearchHealthStatus_NEED_IMPROVEMENT SearchHealthStatus = "NEED_IMPROVEMENT"
	// SearchHealthStatus_NO_DATA：暂无数据。
	SearchHealthStatus_NO_DATA SearchHealthStatus = "NO_DATA"
)
