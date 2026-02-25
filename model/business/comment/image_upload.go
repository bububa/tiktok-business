package comment

import (
	"io"

	"github.com/bububa/tiktok-business/model"
)

// ImageUploadRequest 上传评论图片 Request
type ImageUploadRequest struct {
	// BusinessID Application specific unique identifier for the TikTok Business Account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// ImageFile 用于生成新评论或回复现有评论的图片文件。
	// 大小上限：5 MB。
	// 最高分辨率：1080 x 1920 px 或 1920 x 1080 px。
	// 最低分辨率：360 x 360 px。
	// 支持文件类型：JPG、JPEG、WebP 或 PNG。
	ImageFile io.Reader `json:"-"`
}

// Encode implements UploadRequest
func (r *ImageUploadRequest) Encode() []model.UploadField {
	return []model.UploadField{
		{
			Key:   "business_id",
			Value: r.BusinessID,
		},
		{
			Key:    "image_file",
			Reader: r.ImageFile,
		},
	}
}

// ImageUploadResponse 上传评论图片 Response
type ImageUploadResponse struct {
	model.BaseResponse
	Data *ImageUploadResult `json:"data,omitempty"`
}

type ImageUploadResult struct {
	// ImageURI 已上传图片的 URI。
	// 当上传的图片为可接受类型时返回
	ImageURI string `json:"image_uri,omitempty"`
	// Width image_file 的宽度
	Width int64 `json:"width,omitempty"`
	// Height image_file 的高度。
	Height int64 `json:"height,omitempty"`
}
