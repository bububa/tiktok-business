package enum

// DeeplinkFormatType 深度链接的格式类型。
type DeeplinkFormatType string

const (
	// UNIVERSAL_OR_APP_LINK：通用/应用链接，即以http://或 https:// 开头的 iOS 通用链接或 Android 应用链接。
	UNIVERSAL_OR_APP_LINK DeeplinkFormatType = "UNIVERSAL_OR_APP_LINK"
	// SCHEME_LINK：URL 架构，即格式为scheme://resource 的自定义网址架构。例如，WhatsApp 的自定义网址架构格式应为whatsapp://resource。
	SCHEME_LINK DeeplinkFormatType = "SCHEME_LINK"
	// DeeplinkFormatType_NONE：非深度链接。
	DeeplinkFormatType_NONE DeeplinkFormatType = "NONE"
)
