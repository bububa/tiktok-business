package enum

// BusinessMessageConversationType Conversation type.
type BusinessMessageConversationType string

const (
	// BusinessMessageConversationType_STRANGER: Conversations with users who have sent a message to the Business Account of the business, but the business has not yet accepted the message request and responded to the message.
	BusinessMessageConversationType_STRANGER BusinessMessageConversationType = "STRANGER"
	// BusinessMessageConversationType_SINGLE: One-on-one conversations with users who have sent a message to the Business Account of the business, and the business has previously responded to the user.
	BusinessMessageConversationType_SINGLE BusinessMessageConversationType = "SINGLE"
)
