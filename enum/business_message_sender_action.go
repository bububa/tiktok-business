package enum

// BusinessMessageSenderAction The sender's action in the messaging session.
type BusinessMessageSenderAction string

const (
	// BusinessMessageSenderAction_TYPING: The sender is preparing a reply. This displays a “Typing” icon in the messaging window, which will automatically disappear after five seconds.
	BusinessMessageSenderAction_TYPING BusinessMessageSenderAction = "TYPING"
	// BusinessMessageSenderAction_MARK_READ: The sender marks all messages in the session as read, displaying a “Seen” icon for those messages.
	BusinessMessageSenderAction_MARK_READ BusinessMessageSenderAction = "MARK_READ"
)
