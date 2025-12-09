package message

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AutoMessageCreateRequest Create an automatic message for a Business Account API Request
type AutoMessageCreateRequest struct {
	// BusinessID Application specific unique identifier for the TikTok Business Account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// AutoMessageType The type of automatic message.
	// Enum values:
	// WELCOME_MESSAGE: welcome message. A welcome message is a message that is automatically sent to a user when they start a direct message conversation with you. To learn about how to manage the welcome message, see Manage the welcome message for a Business Account.
	// SUGGESTED_QUESTION: suggested question. A suggested question is a frequently asked question displayed at the beginning of a chat. When the user selects a suggested question, the user will automatically receive the corresponding preset answer. To learn about how to manage suggested questions, see Manage the suggested questions for a Business Account.
	// CHAT_PROMPT: chat prompt. A chat prompt is an interactive button displayed above the input box. When users tap a chat prompt, the system will automatically generate a corresponding question for users to start the conversation. To learn about how to manage chat prompts, see Manage the chat prompts for a Business Account.
	AutoMessageType enum.BusinessAutoMessageType `json:"auto_message_type,omitempty"`
	// WelcomMessage Required when auto_message_type is WELCOME_MESSAGE.
	// Information about the welcome message.
	WelcomeMessage *WelcomeMessage `json:"welcome_message,omitempty"`
	// SuggestedQuestion Required when auto_message_type is SUGGESTED_QUESTION.
	// Information about the suggested question.
	// You can only create one suggested question per call. For each Business Account, you can customize up to three suggested questions. If you try to create more than three suggested questions, an error will occur.
	// To create multiple suggested questions, call /business/message/auto_message/create/ multiple times.
	// Note: Duplicate suggested questions are not supported.
	SuggestedQuestion *SuggestedQuestion `json:"suggested_question,omitempty"`
	// ChatPrompt Required when auto_message_type is CHAT_PROMPT.
	// Information about the chat prompt.
	// You can only create one chat prompt per call. For each Business Account, you can customize up to six chat prompts. If you try to create more than six chat prompts, an error will occur.
	// To create multiple chat prompts, call /business/message/auto_message/create/ multiple times.
	// If you set up multiple chat prompts for your Business Account, the chat prompt that was created earliest will appear as the first prompt above the input box. To sort the chat prompts, use /business/message/auto_message/sort/.
	// Note: Duplicate chat prompts are not supported.
	ChatPrompt *ChatPrompt `json:"chat_prompt,omitempty"`
}

// Encode implements PostRequest
func (r *AutoMessageCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AutoMessageCreateResponse Create an automatic message for a Business Account API Response
type AutoMessageCreateResponse struct {
	model.BaseResponse
	Data struct {
		// AutoMessage Information about the automatic message.
		AutoMessage struct {
			// AutoMessageID The ID of the automatic message
			AutoMessageID string `json:"auto_message_id,omitempty"`
		} `json:"auto_message"`
	} `json:"data"`
}
