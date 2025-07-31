package enum

// MessageAppType 即时通讯广告组中使用的即时通讯应用或自定义 URL。
type MessageAppType string

const (
	// MessageAppType_MESSENGER：Messenger。
	MessageAppType_MESSENGER MessageAppType = "MESSENGER"
	// MessageAppType_WHATSAPP：WhatsApp。
	MessageAppType_WHATSAPP MessageAppType = "WHATSAPP"
	// MessageAppType_ZALO: Zalo。
	MessageAppType_ZALO MessageAppType = "ZALO"
	// MessageAppType_LINE: Line。
	MessageAppType_LINE MessageAppType = "LINE"
	// MessageAppType_IM_URL：即时通讯 URL。
	MessageAppType_IM_URL MessageAppType = "IM_URL"
)
