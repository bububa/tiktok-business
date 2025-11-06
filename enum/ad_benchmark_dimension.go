package enum

// AdBenchmarkDimension 对比维度
type AdBenchmarkDimension string

const (
	// AdBenchmarkDimension_LOCATION（地域）
	AdBenchmarkDimension_LOCATION AdBenchmarkDimension = "LOCATION"
	// AdBenchmarkDimension_AD_CATEGORY（兴趣分类）
	AdBenchmarkDimension_AD_CATEGORY AdBenchmarkDimension = "AD_CATEGORY"
	// AdBenchmarkDimension_EXTERNAL_ACTION（转化事件）
	AdBenchmarkDimension_EXTERNAL_ACTION AdBenchmarkDimension = "EXTERNAL_ACTION"
	// AdBenchmarkDimension_PLACEMENT（版位）。
	AdBenchmarkDimension_PLACEMENT AdBenchmarkDimension = "PLACEMENT"
)
