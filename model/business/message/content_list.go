package message

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ContentListRequest Get a list of messages API Request
type ContentListRequest struct {
	// BusinessID Application specific unique identifier for the TikTok Business Account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// ConversationID Conversation ID.
	// To get the conversation ID, use /business/message/conversation/list/.
	// Note: If you receive a "Param conversation_id is invalid." error, it may be due to encoding issues. Try replacing any "+" characters with "%2B" to resolve potential encoding issues.
	ConversationID string `json:"conversation_id,omitempty"`
}

// Encode implements GetRequest
func (r *ContentListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	values.Set("conversation_id", r.ConversationID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ContentListResponse Get a list of messages API Response
type ContentListResponse struct {
	model.BaseResponse
	Data *ContentListResult `json:"data,omitempty"`
}

type ContentListResult struct {
	// Messages Messages in the conversation
	Messages []Message `json:"messages,omitempty"`
	// Participants Information of participants in the conversation
	Participants []Participant `json:"participants,omitempty"`
}
