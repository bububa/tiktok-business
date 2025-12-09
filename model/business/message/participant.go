package message

import "github.com/bububa/tiktok-business/enum"

// Participant Information of participant in the conversation
type Participant struct {
	// Role The TikTok account type of the participant.
	// Enum values:
	// BUSINESS_ACCOUNT: Business Account.
	// PERSONAL_ACCOUNT: Personal Account.
	Role enum.BusinessMessageParticipantRole `json:"role,omitempty"`
	// ID The identifier of the participant.
	// When role is PERSONAL_ACCOUNT, this field is the same as the unique_identifier, the globally unique identifier assigned to the TikTok Personal Account user in the conversation. This identifier remains consistent across different APIs, such as the Accounts comment listing endpoint, facilitating cross-referencing and integration.
	// When role is BUSINESS_ACCOUNT, this field is the same as the business_id, the application specific unique identifier for the TikTok Business Account.
	ID string `json:"id,omitempty"`
	// DisplayName The TikTok display name (nickname) of the participant
	DisplayName string `json:"display_name,omitempty"`
	// ProfileImage Temporary URL for the TikTok profile image of the participant and needs to be re-acquired after expiration.
	// The expiration time is included in the URL after the x-expires parameter, in the format of an Epoch/Unix timestamp in seconds.
	// Example: https://p16-sign-sg.tiktokcdn.com/aweme/100x100/tos-alisg-avt-1234/5678a9b1c2d345ef6789123b4ghijkl5.jpeg?lk3s=a5d48078&nonce=56434&refresh_token=3b1a321962699a37d22f145bf20cf766&x-expires=1719579600&x-signature=RTkvJazVm0fjkS%2FGfNHOuqIpla4%3D&shp=a5d48078&shcp=8aecc5ac
	ProfileImage string `json:"profile_image,omitempty"`
	// IsFollower Returned only when role is PERSONAL_ACCOUNT.
	// Whether the Personal Account follows the Business Account in the conversation.
	// Supported values: true, false.
	IsFollower bool `json:"is_follower,omitempty"`
}
