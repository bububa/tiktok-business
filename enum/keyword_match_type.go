package enum

// KeywordMatchType 否定关键词的匹配类型。
type KeywordMatchType string

const (
	// PRECISE_WORD ：精确匹配。
	PRECISE_WORD KeywordMatchType = "PRECISE_WORD"
	// PHRASE_WORD：词组匹配。
	PHRASE_WORD KeywordMatchType = "PHRASE_WORD"
	// BROAD_WORD：广泛匹配。
	BROAD_WORD KeywordMatchType = "BROAD_WORD"
)
