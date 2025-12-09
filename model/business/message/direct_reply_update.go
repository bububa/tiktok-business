package message

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// DirectReplyUpdateRequest Enable or disable Comment-to-Message for a Business Account API Request
type DirectReplyUpdateRequest struct {
	// BusinessID Application specific unique identifier for the TikTok Business Account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// DirectReplyType The type of direct reply.
	// Enum value:
	// COMMENT_TO_MESSAGE: Comment-to-Message. Once you enable this feature, the TikTok system will identify high-intent comments, enabling you to send customized replies to assist these users more effectively.
	DirectReplyType enum.BusinessMessageDirectReplyType `json:"direct_reply_type,omitempty"`
	// OperationStatus The operation to perform.
	// Enum values:
	// ENABLE: To enable the Comment-to-Message feature.
	// DISABLE: To disable the Comment-to-Message feature.
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
}

// Encode implements PostRequest
func (r *DirectReplyUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
