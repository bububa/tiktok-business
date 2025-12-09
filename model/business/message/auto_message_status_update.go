package message

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// AutoMessageStatusUpdateRequest Turn on or turn off an automatic message for a Business Account API Request
type AutoMessageStatusUpdateRequest struct {
	// BusinessID Application specific unique identifier for the TikTok Business Account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// AutoMessageType The type of automatic message.
	// Enum values:
	// WELCOME_MESSAGE: welcome message. A welcome message is a message that is automatically sent to a user when they start a direct message conversation with you. To learn about how to manage the welcome message, see Manage the welcome message for a Business Account.
	// SUGGESTED_QUESTION: suggested question. A suggested question is a frequently asked question displayed at the beginning of a chat. When the user selects a suggested question, the user will automatically receive the corresponding preset answer. To learn about how to manage suggested questions, see Manage the suggested questions for a Business Account.
	// CHAT_PROMPT: chat prompt. A chat prompt is an interactive button displayed above the input box. When users tap a chat prompt, the system will automatically generate a corresponding question for users to start the conversation. To learn about how to manage chat prompts, see Manage the chat prompts for a Business Account.
	AutoMessageType enum.BusinessAutoMessageType `json:"auto_message_type,omitempty"`
	// OperationStatus The operation to perform on the type of automatic message.
	// Enum values:
	// ENABLE: To turn on the type of automatic message.
	// DISABLE: To turn off the type of automatic message.
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
}

// Encode implements PostRequest
func (r *AutoMessageStatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
