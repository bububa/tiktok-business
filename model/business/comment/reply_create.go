package comment

import (
	"github.com/bububa/tiktok-business/util"
)

// ReplyCreateRequest 创建对自有视频评论的回复 API Request
type ReplyCreateRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// VideoID 自有 TikTok 视频的唯一标识符，用于列出对应评论（及评论回复）。可通过/business/video/list/ 接口的item_id字段获取
	VideoID string `json:"video_id,omitempty"`
	// CommentID 自有 TikTok 视频评论的唯一标识符，用于创建对应回复。可通过/business/comment/list/ 接口的comment_id字段获取评论 ID
	CommentID string `json:"comment_id,omitempty"`
	// Text 要创建的评论的文本内容。字符数上限为 150 个字符（UTF-8编码）
	Text string `json:"text,omitempty"`
	// ImageURI 必须传入 text 或 image_uri。
	// 用作回复的图片的 URI。
	// 若想获取图片 URI，需先调用 /business/comment/image/upload/ 接口。
	ImageURI string `json:"image_uri,omitempty"`
	// ImageWidth 当传入 image_uri 时，此为必填字段。
	// 图片的宽度。
	ImageWidth int64 `json:"image_width,omitempty"`
	// ImageHeight 当传入 image_uri 时，此为必填字段。
	// 图片的高度。
	ImageHeight int64 `json:"image_height,omitempty"`
}

// Encode implements PostRequest
func (r *ReplyCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
