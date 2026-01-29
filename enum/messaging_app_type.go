package enum

// MessagingAppType The type of instant messaging app or customized URL to use in the Instant Messaging Ad Group.
type MessagingAppType string

const (
	// MessagingAppType_MESSENGER: Messenger.
	MessagingAppType_MESSENGER MessagingAppType = "MESSENGER"
	// MessagingAppType_WHATSAPP: WhatsApp.
	MessagingAppType_WHATSAPP MessagingAppType = "WHATSAPP"
	// MessagingAppType_ZALO: Zalo.
	MessagingAppType_ZALO MessagingAppType = "ZALO"
	// MessagingAppType_LINE: Line.
	MessagingAppType_LINE MessagingAppType = "LINE"
	// MessagingAppType_IM_URL: Instant Messaging URL.
	MessagingAppType_IM_URL MessagingAppType = "IM_URL"
)
