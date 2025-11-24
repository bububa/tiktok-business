package comment

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 为自有视频创建新评论
type CreateRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// VideoID 自有 TikTok 视频的唯一标识符，用于列出对应评论（及评论回复）。可通过/business/video/list/ 接口的item_id字段获取
	VideoID string `json:"video_id,omitempty"`
	// Text 要创建的评论的文本内容。字符数上限为 150 个字符（UTF-8编码）
	Text string `json:"text,omitempty"`
}

// Encode implements PostRequest
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 为自有视频创建新评论
type CreateResponse struct {
	model.BaseResponse
	Data *Comment `json:"data,omitempty"`
}
