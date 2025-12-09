package message

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CapabilitiesGetRequest Check the capability of a Business Account for a conversation API Request
type CapabilitiesGetRequest struct {
	// BusinessID Application specific unique identifier for the TikTok Business Account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// CapabilityTypes A list of capability types that you want to query for the Business Account.
	// Enum value:
	// IMAGE_SEND: The Business Account can send images and/or receive images in the specified conversation.
	CapabilityTypes []enum.BusinessMessageCapabilityType `json:"capability_types,omitempty"`
	// ConversationID Required when the value of capability_types includes IMAGE_SEND.
	// Conversation ID.
	// To get a list of conversation IDs, use /business/message/conversation/list/.
	// Note: If you receive a "Param conversation_id is invalid." error, it may be due to encoding issues. Try replacing any "+" characters with "%2B" to resolve potential encoding issues.
	ConversationID string `json:"conversation_id,omitempty"`
	// ConversationType Required when the value of capability_types includes IMAGE_SEND.
	// Conversation type.
	// Currently, the only allowed values are:
	// STRANGER: Conversations with users who have sent a message to the Business Account of the business, but the business has not yet accepted the message request and responded to the message.
	// SINGLE: One-on-one conversations with users who have sent a message to the Business Account of the business, and the business has previously responded to the user.
	ConversationType enum.BusinessMessageConversationType `json:"conversation_type,omitempty"`
}

// Encode implements GetRequest
func (r *CapabilitiesGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	values.Set("capability_types", string(util.JSONMarshal(r.CapabilityTypes)))
	if r.ConversationID != "" {
		values.Set("conversation_id", r.ConversationID)
	}
	if r.ConversationType != "" {
		values.Set("conversation_type", string(r.ConversationType))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// CapabilitiesGetResponse Check the capability of a Business Account for a conversation API Response
type CapabilitiesGetResponse struct {
	model.BaseResponse
	Data struct {
		// CapabilityInfos The results regarding the queried capabilities for the Business Account.
		CapabilityInfos []CapabilityInfo `json:"capability_infos,omitempty"`
	} `json:"data"`
}

// CapabilityInfo The results regarding the queried capabilities for the Business Account.
type CapabilityInfo struct {
	// CapabilityType The capability type.
	// Enum value:
	// IMAGE_SEND: The Business Account can send images and/or receive images in the specified conversation.
	CapabilityType enum.BusinessMessageCapabilityType `json:"capability_type,omitempty"`
	// CapabilityResult Whether the Business Account has this capability.
	// Supported values: true, false.
	CapabilityResult bool `json:"capability_result,omitempty"`
}
