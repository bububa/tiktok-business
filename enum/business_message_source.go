package enum

// BusinessMessageSource The source of the message, indicating the platform from which the message was sent.
type BusinessMessageSource string

const (
	// BusinessMessageSource_APP: TikTok mobile app.
	BusinessMessageSource_APP BusinessMessageSource = "APP"
	// BusinessMessageSource_WEB: TikTok Web or TikTok Business Center.
	BusinessMessageSource_WEB BusinessMessageSource = "WEB"
	// BusinessMessageSource_API: TikTok Business Messaging API.
	BusinessMessageSource_API BusinessMessageSource = "API"
	// BusinessMessageSource_OTHERS: Other TikTok internal servers or tools, including automatic bots.
	BusinessMessageSource_OTHERS BusinessMessageSource = "OTHERS"
	// BusinessMessageSource_UNKNOWN_SOURCE: Unknown source.
	BusinessMessageSource_UNKNOWN_SOURCE BusinessMessageSource = "UNKNOWN_SOURCE"
)
