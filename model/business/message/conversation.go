package message

// Conversation conversation details
type Conversation struct {
	// ConversationID Conversation ID
	ConversationID string `json:"conversation_id,omitempty"`
	// UpTime The time of the last message in the conversation, in the format of an Epoch/Unix timestamp in milliseconds
	UpTime int64 `json:"up_time,omitempty"`
	// Referral Referral information of the conversation
	Referral *ConversationReferral `json:"referral,omitempty"`
}

// ConversationReferral Referral information of the conversation
type ConversationReferral struct {
	// Ad Related referral ad information of the conversation, including all historical referral ads of the conversation
	Ad []ConversationReferralAd `json:"ad,omitempty"`
	// ShortLink Information about the tiktok.me link.
	// A tiktok.me link is a shareable TikTok shortlink that allows users to initiate a one-on-one conversation with a business in the TikTok App.
	ShortLink []ConversationReferralShortLink `json:"short_link,omitempty"`
}

// ConversationReferralAd Related referral ad information of the conversation, including all historical referral ads of the conversation
type ConversationReferralAd struct {
	// AdvertiserID The ID of the advertiser account that is associated with the ad
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdID Ad ID of the related ad
	AdID string `json:"ad_id,omitempty"`
	// Timestamp The time of the referral event, in the format of an Epoch/Unix timestamp in seconds
	Timestamp int64 `json:"timestamp,omitempty"`
	// AdName The name of the related ad
	AdName string `json:"ad_name,omitempty"`
	// EmbedURL An embeddable link for the TikTok post used in the related ad.
	// This URL can be used to embed the TikTok post in external websites or applications.
	// If the privacy setting of the post is "Friends" or "Only you", the post will not be viewable through this link
	EmbedURL string `json:"embed_url,omitempty"`
}

// ConversationReferralShortLink Information about the tiktok.me link.
// A tiktok.me link is a shareable TikTok shortlink that allows users to initiate a one-on-one conversation with a business in the TikTok App.
type ConversationReferralShortLink struct {
	// Ref The referral parameter (ref) that is configured in the tiktok.me link, offering context about the conversation. You can incorporate the referral parameter in your tiktok.me link to monitor various links posted on different channels.
	// This parameter is a URL-encoded string with a maximum length of 2,083 characters.
	// Note: The ref parameter only supports alphanumeric characters and the special characters -, _, and =.
	Ref string `json:"ref,omitempty"`
	// PrefilledMessage The prefilled message that is configured in the tiktok.me link.
	PrefilledMessage string `json:"prefilled_message,omitempty"`
	// PrefilledMessageAuditStatus The audit status of the prefilled message.
	// Enum values:
	// REJECT: The message was rejected during the review. If prefilled_message_audit_status is REJECT, the prefilled_message will not be prefilled into the message input box.
	// PASS: The message passed the review.
	PrefilledMessageAuditStatus string `json:"prefilled_message_audit_status,omitempty"`
}
