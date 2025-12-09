package message

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SendRequest Send a message to a conversation API Request
type SendRequest struct {
	// BusinessID Application specific unique identifier for the TikTok Business Account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// RecipientType Required when direct_reply is not specified.
	// Not supported when direct_reply is specified.
	// Currently, the only supported value is CONVERSATION
	RecipientType enum.BusinessMessageRecipientType `json:"recipient_type,omitempty"`
	// Recipient Required when direct_reply is not specified.
	// Not supported when direct_reply is specified.
	// Recipient.
	// When recipient_type is CONVERSATION, you need to specify conversation_id in this field.
	// To get the conversation ID, use /business/message/conversation/list/.
	Recipient string `json:"recipient,omitempty"`
	// MessageType The type of the message.
	// Enum values:
	// TEXT: text message.
	// IMAGE: image message.
	// SHARE_POST: TikTok post message.
	// TEMPLATE: message template.
	// SENDER_ACTION: Sender action. Note that for this type, the returned message ID is an empty string ("") because no actual message is sent.
	MessageType enum.BusinessMessageType `json:"message_type,omitempty"`
	// Text Required when message_type is TEXT.
	// Text information.
	// A message cannot include both a text and an image.
	Text *Text `json:"text,omitempty"`
	// Image Required when message_type is IMAGE.
	// Image information.
	// A message cannot include both a text and an image.
	// Sending and receiving image attachments via API are only available when both the sender and receiver are in countries that support image attachments in direct messages. Image sharing is not available in all markets.
	Image *Image `json:"image,omitempty"`
	// SharePost Required when message_type is SHARE_POST.
	// Information about the post to be shared.
	SharePost *SharePost `json:"share_post,omitempty"`
	// Template Required when message_type is TEMPLATE.
	// Details about the message template to send.
	Template *Template `json:"template,omitempty"`
	// SenderAction Required when message_type is SENDER_ACTION.
	// The sender's action in the messaging session.
	// TYPING: The sender is preparing a reply. This displays a “Typing” icon in the messaging window, which will automatically disappear after five seconds.
	// MARK_READ: The sender marks all messages in the session as read, displaying a “Seen” icon for those messages.
	SenderAction enum.BusinessMessageSenderAction `json:"sender_action,omitempty"`
	// ReferencedMessageInfo Required when you want to send a reply to a message.
	// Information about the original message to reply to.
	ReferencedMessageInfo *ReferencedMessageInfo `json:"referenced_message_info,omitempty"`
	// DirectReply Required when you are sending a direct reply.
	// Details of the direct reply.
	// Note:
	// To send a direct reply, ensure that Comment-to-Message is enabled for the Business Account. To obtain the Comment-to-Message setting for a Business Account, use /business/message/direct_reply/get/.
	// The Comment-to-Message feature is only available for Business Accounts registered in Vietnam, Indonesia, and Thailand.
	// Once you enable Comment-to-Message for a Business Account, the Business Account can only reply to comments published by TikTok accounts that are registered in the APAC, LATAM, and METAP regions.
	DirectReply *DirectReply `json:"direct_reply,omitempty"`
}

// Encode implements PostRequest
func (r *SendRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// SendResponse Send a message to a conversation API Response
type SendResponse struct {
	model.BaseResponse
	Data struct {
		// Message Message details
		Message struct {
			// MessageID Message ID.
			// Note: If message_type is set to SENDER_ACTION in the request, the value of this field will be an empty string ("") because no actual message is sent.
			MessaegID string `json:"message_id,omitempty"`
		}
	}
}
