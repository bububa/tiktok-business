package comment

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// PostRequest 回复评论 API Request
type PostRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdID 广告 ID
	AdID string `json:"ad_id,omitempty"`
	// TikTokItemID TikTok 视频 ID
	TikTokItemID string `json:"tiktok_item_id,omitempty"`
	// CommentID 您想要回复的评论 ID
	CommentID string `json:"comment_id,omitempty"`
	// CommentType 评论类型。目前仅支持 REPLY
	CommentType enum.CommentType `json:"comment_type,omitempty"`
	// Text 回复内容
	Text string `json:"text,omitempty"`
	// IdentityType 认证身份类型。
	// 枚举值：CUSTOMIZED_USER (自定义用户）， TT_USER（TikTok 用户）。
	// 关于认证身份的更多信息，请访问认证身份
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityID 认证身份 ID
	IdentityID string `json:"identity_id,omitempty"`
}

// Encode implements PostRequest
func (r *PostRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// PostResponse 回复评论 API Response
type PostResponse struct {
	model.BaseResponse
	Data *Comment `json:"data,omitempty"`
}
