package enum

// KeywordStatus 其他兴趣分类状态，即分类是否有效。
type KeywordStatus string

const (
	// KeywordStatus_EFFECTIVE：有效。有效其他兴趣分类指一个分类能覆盖的受众人数足够大。
	KeywordStatus_EFFECTIVE KeywordStatus = "EFFECTIVE"
	// KeywordStatus_INEFFECTIVE：无效。无效其他兴趣分类指一个分类能覆盖的受众人数太少。
	KeywordStatus_INEFFECTIVE KeywordStatus = "INEFFECTIVE"
)
