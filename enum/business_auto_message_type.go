package enum

// BusinessAutoMessageType The type of automatic message.
type BusinessAutoMessageType string

const (
	// BusinessAutoMessageType_WELCOME_MESSAGE: welcome message. A welcome message is a message that is automatically sent to a user when they start a direct message conversation with you.
	BusinessAutoMessageType_WELCOME_MESSAGE BusinessAutoMessageType = "WELCOME_MESSAGE"
	// BusinessAutoMessageType_SUGGESTED_QUESTION: suggested question. A suggested question is a frequently asked question displayed at the beginning of a chat. When the user selects a suggested question, the user will automatically receive the corresponding preset answer.
	BusinessAutoMessageType_SUGGESTED_QUESTION BusinessAutoMessageType = "SUGGESTED_QUESTION"
	// BusinessAutoMessageType_AUTO_REPLY: keyword reply. A keyword reply is an automated response triggered by specific keywords in a message.
	BusinessAutoMessageType_AUTO_REPLY BusinessAutoMessageType = "AUTO_REPLY"
	// BusinessAutoMessageType_Note: When the message is not an automatic message, this field will not be returned.
	BusinessAutoMessageType_Note BusinessAutoMessageType = "Note"
	// BusinessAutoMessageType_CHAT_PROMPT: chat prompt. A chat prompt is an interactive button displayed above the input box. When users tap a chat prompt, the system will automatically generate a corresponding question for users to start the conversation. To learn about how to manage chat prompts, see Manage the chat prompts for a Business Account.
	BusinessAutoMessageType_CHAT_PROMPT BusinessAutoMessageType = "CHAT_PROMPT"
)
