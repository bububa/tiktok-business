package enum

// BusinessMessageButtonType The type of the button or text link.
type BusinessMessageButtonType string

const (
	// BusinessMessageButtonType_REPLY: Reply button or reply text link. When users click the button or text link, the system will send the text content of the button or text link as a reply message on behalf of the user in the conversation.
	BusinessMessageButtonType_REPLY BusinessMessageButtonType = "REPLY"
)
