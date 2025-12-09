package message

import "github.com/bububa/tiktok-business/enum"

// AutoMessage The details of automatic message
type AutoMessage struct {
	// AutoMessageID The ID of the automatic message
	AutoMessageID string `json:"auto_message_id,omitempty"`
	// AutoMessageType The type of automatic message.
	// Enum values:
	// WELCOME_MESSAGE: welcome message. A welcome message is a message that is automatically sent to a user when they start a direct message conversation with you.
	// SUGGESTED_QUESTION: suggested question. A suggested question is a frequently asked question displayed at the beginning of a chat. When the user selects a suggested question, the user will automatically receive the corresponding preset answer.
	// CHAT_PROMPT: chat prompt. A chat prompt is an interactive button displayed above the input box. When users tap a chat prompt, the system will automatically generate a corresponding question for users to start the conversation.
	AutoMessageType enum.BusinessAutoMessageType `json:"auto_message_type,omitempty"`
	// AuditStatus The review status of the automatic message.
	// Enum values:
	// REVIEWING: The automatic message is pending review.
	// APPROVED: The automatic message has been approved.
	// REJECTED: The automatic message has been rejected.
	// Note:
	// If the audit_status of the automatic message is REJECTED, you cannot update the automatic message by using /business/message/auto_message/status/update/.
	// If audit_status of the message is REVIEWING, you need to wait for some time before rechecking the review status. The exact waiting time may vary depending on the content and other factors, but normally ranges from a few seconds to 2-3 hours.
	AuditStatus enum.BusinessAutoMessageAuditStatus `json:"audit_status,omitempty"`
	// WelcomeMessage Returned only when auto_message_type is WELCOME_MESSAGE.
	// Information about the welcome message.
	WelcomeMessage *WelcomeMessage `json:"welcome_message,omitempty"`
	// SuggestedQuestion Returned only when auto_message_type is SUGGESTED_QUESTION.
	// Information about the suggested question.
	SuggestedQuestion *SuggestedQuestion `json:"suggested_question,omitempty"`
	// ChatPrompt Returned only when auto_message_type is CHAT_PROMPT.
	// Information about the chat prompt.
	ChatPrompt *ChatPrompt `json:"chat_prompt,omitempty"`
}

// WelcomeMessage Required when auto_message_type is WELCOME_MESSAGE.
// Information about the welcome message.
type WelcomeMessage struct {
	// Content Required when welcome_message is passed.
	// The text of the welcome message.
	// Length limit: 250 characters.
	// Example: Hi there, Please leave us a message and we will respond to you shortly. Thank you!.
	// Note: You can use "\n" as a line break to display the welcome message in multiple lines within the mobile TikTok App. To ensure proper display, check the welcome message in the TikTok mobile app after deployment.
	Content string `json:"content,omitempty"`
}

// SuggestedQuestion Required when auto_message_type is SUGGESTED_QUESTION.
// Information about the suggested question.
type SuggestedQuestion struct {
	// Question Required when suggested_question is passed.
	// A frequently asked question to help users start a conversation.
	// Length limit: 80 characters.
	Question string `json:"question,omitempty"`
	// Answer Required when suggested_question is passed.
	// The preset answer that will be automatically sent when a user selects the question.
	// Length limit: 200 characters.
	Answer string `json:"answer,omitempty"`
}

// ChatPrompt Required when auto_message_type is CHAT_PROMPT.
// Information about the chat prompt.
type ChatPrompt struct {
	// Title Required when chat_prompt is passed.
	// The name of the prompt button that is displayed above the input box.
	// Length limit: 18 characters.
	// You can use text, system emojis, and both English and non-English languages.
	// To ensure optimal display, keep the name concise and clear.
	Title string `json:"title,omitempty"`
	// Content Required when chat_prompt is passed.
	// The corresponding question for the prompt that the system automatically generates for users to start the conversation when users tap the chat prompt.
	// Length limit: 150 characters.
	// You can use text, system emojis, and both English and non-English languages.
	Content string `json:"content,omitempty"`
}
