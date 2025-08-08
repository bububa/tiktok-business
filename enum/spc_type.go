package enum

// SpcType 网站推广系列类型。
type SpcType string

const (
	// SpcType_WEB_ALL_IN_ONE：Smart+ 网站推广系列，支持的优化目标包括转化、价值、点击和落地页浏览量。
	SpcType_WEB_ALL_IN_ONE SpcType = "WEB_ALL_IN_ON"
	// SpcType_UNSET：智能效果网站推广系列，支持的优化目标仅包括转化、价值和点击。
	SpcType_UNSET SpcType = "UNSET"
)
