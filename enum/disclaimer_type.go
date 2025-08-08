package enum

// DisclaimerType 广告中的免责声明类型。
type DisclaimerType string

const (
	// DisclaimerType_TEXT_LINK（文字链免责声明）
	DisclaimerType_TEXT_LINK DisclaimerType = "TEXT_LINK"
	// DisclaimerType_TEXT_ONLY（纯文本免责声明）。
	DisclaimerType_TEXT_ONLY DisclaimerType = "TEXT_ONLY"
)
