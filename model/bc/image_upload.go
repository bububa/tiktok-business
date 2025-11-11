package bc

import (
	"io"

	"github.com/bububa/tiktok-business/model"
)

// ImageUploadRequest 上传资质证书 API Request
type ImageUploadRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// ImageFile 资质证书文件，仅支持图片格式(JPG/JPEG/PNG)。文件大小最大为10 MB
	ImageFile io.Reader `json:"image_file,omitempty"`
}

// Encode implements UploadRequest
func (r *ImageUploadRequest) Encode() []model.UploadField {
	return []model.UploadField{
		{
			Key:   "bc_id",
			Value: r.BcID,
		},
		{
			Key:    "image_file",
			Reader: r.ImageFile,
			Value:  "image_file",
		},
	}
}

// ImageUploadResponse 上传资质证书 API Response
type ImageUploadResponse struct {
	model.BaseResponse
	Data *ImageUploadResult `json:"data,omitempty"`
}

type ImageUploadResult struct {
	// ImageID 资质证书图片ID。您可以将该字段的值传入/bc/advertiser/create/接口，在商务中心下创建广告账户
	ImageID string `json:"image_id,omitempty"`
	// ImageURL 资质证书图片预览链接，查看一次后失效
	ImageURL string `json:"image_url,omitempty"`
}
