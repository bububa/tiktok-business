package message

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ConversationListRequest Get a list of conversations API Request
type ConversationListRequest struct {
	// BusinessID Application specific unique identifier for the TikTok Business Account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// ConversationType Conversation type.
	// Currently, the only allowed values are:
	// STRANGER: Conversations with users who have sent a message to the Business Account of the business, but the business has not yet accepted the message request and responded to the message.
	// SINGLE: One-on-one conversations with users who have sent a message to the Business Account of the business, and the business has previously responded to the user.
	ConversationType enum.BusinessMessageConversationType `json:"conversation_type,omitempty"`
	// Limit The maximum number of conversations to return for this request.
	// Value range: 1-100.
	// Default value: 100.
	Limit int `json:"limit,omitempty"`
	// Cursor Cursor for pagination.
	// If the response parameter has_more is true, pass the returned cursor value in this parameter in the subsequent request to receive the next page of conversations.
	// Default value: 0.
	Cursor int64 `json:"cursor,omitempty"`
}

// Encode implements GetRequest
func (r *ConversationListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	values.Set("conversation_type", string(r.ConversationType))
	if r.Limit > 0 {
		values.Set("limit", strconv.Itoa(r.Limit))
	}
	if r.Cursor > 0 {
		values.Set("cursor", strconv.FormatInt(r.Cursor, 10))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ConversationListResponse Get a list of conversations API Response
type ConversationListResponse struct {
	model.BaseResponse
	Data *ConversationListResult `json:"data,omitempty"`
}

type ConversationListResult struct {
	// Conversations Conversation List.
	Conversations []Conversation `json:"conversations,omitempty"`
	// HasMore Whether an additional page of conversations is available.
	HasMore bool `json:"has_more,omitempty"`
	// Cursor Cursor for the next page of conversations, to be passed in the cursor parameter of the subsequent API request.
	Cursor int64 `json:"cursor,omitempty"`
}
