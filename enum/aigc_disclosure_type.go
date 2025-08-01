package enum

// AigcDisclosureType 是否启用 AI 生成内容自主声明开关，以表明广告中包含 AI 生成内容。启用该开关后，当用户查看完整广告时，您的广告将带有“广告主标记为 AI 生成”的标签。
type AigcDisclosureType string

const (
	// AigcDisclosureType_SELF_DISCLOSURE：启用开关，声明广告包含 AI 生成内容。
	AigcDisclosureType_SELF_DISCLOSURE AigcDisclosureType = "SELF_DISCLOSURE"
	// AigcDisclosureType_NOT_DECLARED：不声明广告包含 AI 生成内容。
	AigcDisclosureType_NOT_DECLARED AigcDisclosureType = "NOT_DECLARED"
)
