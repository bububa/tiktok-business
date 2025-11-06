package ttvideo

import "github.com/bububa/tiktok-business/util"

// AuthorizeRequest 应用授权码 API Request
type AuthorizeRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AuthCode Spark Ads帖子授权码
	AuthCode string `json:"auth_code,omitempty"`
	// OriginalPostAuthCode 如果本帖子是一个拼接或者合拍帖或者提及了其他视频，你还需要获得源视频的授权，并使用本字段传入源视频的授权码。
	// 注意：目前，帖子只能提及一个视频。
	OriginalPostAuthCode string `json:"original_post_auth_code,omitempty"`
}

// Encode implements PostRequest
func (r *AuthorizeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
