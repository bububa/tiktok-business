package enum

// KeywordRelevanceStatus 关键词相关性分析的结果。
type KeywordRelevanceStatus string

const (
	// KeywordRelevanceStatus_TO_BE_CALCULATED：待计算。
	KeywordRelevanceStatus_TO_BE_CALCULATED KeywordRelevanceStatus = "TO_BE_CALCULATED"
	// KeywordRelevanceStatus_PARTIALLY_RELEVANT：部分相关。
	KeywordRelevanceStatus_PARTIALLY_RELEVANT KeywordRelevanceStatus = "PARTIALLY_RELEVANT"
	// KeywordRelevanceStatus_RELEVANT：相关。
	KeywordRelevanceStatus_RELEVANT KeywordRelevanceStatus = "RELEVANT"
	// KeywordRelevanceStatus_IRRELEVANT：不相关。
	KeywordRelevanceStatus_IRRELEVANT KeywordRelevanceStatus = "IRRELEVANT"
)
