package comment

import (
	"github.com/bububa/tiktok-business/util"
)

// DeleteRequest 删除自有视频的评论 API Request
type DeleteRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// CommentID 自有 TikTok 视频评论的唯一标识符，用于创建对应回复。可通过/business/comment/list/ 接口的comment_id字段获取评论 ID
	CommentID string `json:"comment_id,omitempty"`
}

// Encode implements PostRequest
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
