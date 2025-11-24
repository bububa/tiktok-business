package enum

// TargetingSearchType 搜索类型。
type TargetingSearchType string

const (
	// TargetingSearchType_FUZZY_SEARCH：对单个地域 ID 或邮政编码 ID 进行模糊搜索。
	TargetingSearchType_FUZZY_SEARCH TargetingSearchType = "FUZZY_SEARCH"
	// 对于此搜索类型，您仅可向 keywords 字段传入一个关键词，最多返回100条搜索结果。
	// TargetingSearchType_BATCH_REGION_SEARCH：对多个地域ID进行批量模糊搜索。
	TargetingSearchType_BATCH_REGION_SEARCH TargetingSearchType = "BATCH_REGION_SEARCH"
	// 对于此搜索类型，您最多可向 keywords 字段传入1,000个关键词，最多为每个关键词返回5条按相关度排序的搜索结果。
	// TargetingSearchType_BATCH_ZIPCODE_SEARCH : 对多个邮政编码ID进行批量精确搜索。
	// 对于此搜索类型，您最多可向 keywords 字段传入1,000个关键词，最多为每个关键词返回 1 条搜索结果，或无匹配结果时返回 0 条搜索结果。
	TargetingSearchType_BATCH_ZIPCODE_SEARCH TargetingSearchType = "BATCH_ZIPCODE_SEARCH"
)
