package comment

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// DeleteRequest 回复评论 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdID 广告 ID
	AdID string `json:"ad_id,omitempty"`
	// TikTokItemID TikTok 视频 ID
	TikTokItemID string `json:"tiktok_item_id,omitempty"`
	// CommentID 您想要回复的评论 ID
	CommentID string `json:"comment_id,omitempty"`
	// IdentityType 认证身份类型。
	// 枚举值：CUSTOMIZED_USER (自定义用户）， TT_USER（TikTok 用户）。
	// 关于认证身份的更多信息，请访问认证身份
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityID 认证身份 ID
	IdentityID string `json:"identity_id,omitempty"`
}

// Encode implements PostRequest
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
