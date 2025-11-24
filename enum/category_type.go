package enum

// CategoryType 类别种类。
type CategoryType string

const (
	// CATEGORY_TYPE_EXCLUSION（内容排除类别）
	CATEGORY_TYPE_EXCLUSION CategoryType = "CATEGORY_TYPE_EXCLUSION"
	// CATEGORY_TYPE_VERTICAL（行业敏感内容控制类别）
	CATEGORY_TYPE_VERTICAL CategoryType = "CATEGORY_TYPE_VERTICAL"
)
