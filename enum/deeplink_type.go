package enum

// DeeplinkType 深度链接类型。
type DeeplinkType string

const (
	// DeeplinkType_NORMAL：非延迟深度链接。
	DeeplinkType_NORMAL DeeplinkType = "NORMAL"
	// DeeplinkType_DEFERRED_DEEPLINK：延迟深度链接
	DeeplinkType_DEFERRED_DEEPLINK DeeplinkType = "DEFERRED_DEEPLINK"
)
