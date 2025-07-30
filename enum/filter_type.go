package enum

// FilterType 筛选类型。
type FilterType string

const (
	// FilterType_IN：包含，当筛选类型为此项时，筛选的值必须是合法的JSON数组字符串。
	FilterType_IN FilterType = "IN"
	// FilterType_MATCH：模糊匹配，相当于 MySQL 中的LIKE 操作。
	FilterType_MATCH FilterType = "MATCH"
	// FilterType_GREATER_EQUAL：大于等于（简称GE）。
	FilterType_GREATER_EQUAL FilterType = "GREATER_EQUAL"
	// FilterType_GREATER_THAN：大于（简称GT）。
	FilterType_GREATER_THAN FilterType = "GREATER_THAN"
	// FilterType_LOWER_EQUAL：小于等于（简称LE）。
	FilterType_LOWER_EQUAL FilterType = "LOWER_EQUAL"
	// FilterType_LOWER_THAN：小于（简称LT）。
	FilterType_LOWER_THAN FilterType = "LOWER_THAN"
	// FilterType_BETWEEN： 在……之间。当筛选类型为此项时，筛选的值需要是合法的、包含两个元素的JSON数组字符串。
	FilterType_BETWEEN FilterType = "BETWEEN"
)
