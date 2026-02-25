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
	// ReplyImageURI 用作回复的自定义图片的可公开访问 HTTP(s) URL。
	// 图片限制：
	// 最高分辨率：1080 x 1920 px 或 1920 x 1080 px。
	// 最低分辨率：360 x 360 px。
	// 文件大小：不超过 20 MB。
	// 文件类型：JPG、JPEG、WebP 或 PNG。
	// 示例：https://tiktok.com/tiktok-images/image.jpg
	// 注意：传入已手动或自动验证所有权的 URL。若想手动验证 URL 资源的所有权，请按照管理 URL 资源中所述的步骤操作。
	ReplyImageURI string `json:"reply_image_uri,omitempty"`
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
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 为自有视频创建新评论
type CreateResponse struct {
	model.BaseResponse
	Data *Comment `json:"data,omitempty"`
}
