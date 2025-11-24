package comment

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// HideRequest 隐藏/取消隐藏自有视频的评论 API Request
type HideRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// VideoID 自有 TikTok 视频的唯一标识符，用于列出对应评论（及评论回复）。可通过/business/video/list/ 接口的item_id字段获取
	VideoID string `json:"video_id,omitempty"`
	// CommentID 自有 TikTok 视频评论的唯一标识符，用于创建对应回复。可通过/business/comment/list/ 接口的comment_id字段获取评论 ID
	CommentID string `json:"comment_id,omitempty"`
	// Action 要对评论进行的特定操作 - 隐藏或取消隐藏。
	// 枚举值：
	// HIDE
	// UNHIDE
	Action enum.CommentAction `json:"action,omitempty"`
}

// Encode implements PostRequest
func (r *HideRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
