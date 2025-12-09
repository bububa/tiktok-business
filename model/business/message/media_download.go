package message

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// MediaDownloadRequest Download an image or a video from a message API Request
type MediaDownloadRequest struct {
	// BusinessID Application specific unique identifier for the TikTok Business Account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// ConversationID Conversation ID.
	// To get the conversation ID, use /business/message/conversation/list/.
	ConversationID string `json:"conversation_id,omitempty"`
	// MessageID Message ID.
	// To get a message ID, use /business/message/content/list/.
	MessageID string `json:"message_id,omitempty"`
	// MediaID Media ID.
	// To get a media ID, use /business/message/content/list/.
	MediaID string `json:"media_id,omitempty"`
	// MediaType Multimedia type.
	// Enum values:
	// IMAGE: image.
	// VIDEO: video.
	MediaType string `json:"media_type,omitempty"`
}

// Encode implements PostRequest
func (r *MediaDownloadRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// MediaDownloadResponse Download an image or a video from a message API Response
type MediaDownloadResponse struct {
	model.BaseResponse
	Data struct {
		// DownloadURL Download URL.
		// It is valid for 24 hours. If it expires, call this endpoint again to generate a new URL.
		// Note: When downloading the image through the URL, you must include a header. Set the header to x-user = {{Access-Token}}
		DownloadURL string `json:"download_url,omitempty"`
	} `json:"data"`
}
