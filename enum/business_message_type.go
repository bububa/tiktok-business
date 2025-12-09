package enum

// BusinessMessageType The type of the message.
type BusinessMessageType string

const (
	// BusinessMessageType_TEXT: text message.
	BusinessMessageType_TEXT BusinessMessageType = "TEXT"
	// BusinessMessageType_IMAGE: image message.
	BusinessMessageType_IMAGE BusinessMessageType = "IMAGE"
	// BusinessMessageType_SHARE_POST: TikTok post message.
	BusinessMessageType_SHARE_POST BusinessMessageType = "SHARE_POST"
	// BusinessMessageType_VIDEO: video.
	BusinessMessageType_VIDEO BusinessMessageType = "VIDEO"
	// BusinessMessageType_EMOJI: emoji.
	BusinessMessageType_EMOJI BusinessMessageType = "EMOJI"
	// BusinessMessageType_STICKER: sticker.
	BusinessMessageType_STICKER BusinessMessageType = "STICKER"
	// BusinessMessageType_TEMPLATE: message template.
	BusinessMessageType_TEMPLATE BusinessMessageType = "TEMPLATE"
	// BusinessMessageType_SENDER_ACTION: Sender action. Note that for this type, the returned message ID is an empty string ("") because no actual message is sent.
	BusinessMessageType_SENDER_ACTION BusinessMessageType = "SENDER_ACTION"
	// BusinessMessageType_OTHER: other message types.
	BusinessMessageType_OTHER BusinessMessageType = "OTHER"
)
