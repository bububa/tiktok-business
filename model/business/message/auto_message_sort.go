package message

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// AutoMessageSortRequest Sort the automatic message for a Business Account API Request
type AutoMessageSortRequest struct {
	// BusinessID Application specific unique identifier for the TikTok Business Account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// AutoMessageType The type of automatic message.
	// Enum values:
	// SUGGESTED_QUESTION: suggested question. A suggested question is a frequently asked question displayed at the beginning of a chat. When the user selects a suggested question, the user will automatically receive the corresponding preset answer. To learn about how to manage suggested questions, see Manage the suggested questions for a Business Account.
	// CHAT_PROMPT: chat prompt. A chat prompt is an interactive button displayed above the input box. When users tap a chat prompt, the system will automatically generate a corresponding question for users to start the conversation. To learn about how to manage chat prompts, see Manage the chat prompts for a Business Account.
	AutoMessageType enum.BusinessAutoMessageType `json:"auto_message_type,omitempty"`
	// AutoMessageID The ID of the automatic message.
	// To retrieve the automatic messages for a Business Account, use /business/message/auto_message/get/.
	AutoMessageID string `json:"auto_message_id,omitempty"`
}

// Encode implements PostRequest
func (r *AutoMessageSortRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
