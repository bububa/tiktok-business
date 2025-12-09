package message

import (
	"io"

	"github.com/bububa/tiktok-business/model"
)

// MediaUploadRequest Upload an image API Request
type MediaUploadRequest struct {
	// BusinessID Application specific unique identifier for the TikTok Business Account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// File Multimedia File.
	// File size limit: 3MB.
	// Supported image file format: JPG and PNG.
	File io.Reader `json:"-"`
	// MediaType Multimedia type.
	// Currently, we only support IMAGE as the multimedia type.
	MediaType string `json:"media_type,omitempty"`
}

// Encode implements UploadRequest
func (r *MediaUploadRequest) Encode() []model.UploadField {
	return []model.UploadField{
		{
			Key:   "business_id",
			Value: r.BusinessID,
		},
		{
			Key:    "file",
			Reader: r.File,
		},
		{
			Key:   "media_type",
			Value: r.MediaType,
		},
	}
}

// MediaUploadResponse Upload an image API Response
type MediaUploadResponse struct {
	model.BaseResponse
	Data struct {
		// MediaID Media ID of the uploaded image.
		// Validity period: 30 days. Once the Media ID expires, you need to re-upload the image to obtain a new Media ID.
		MediaID string `json:"media_id,omitempty"`
	} `json:"data"`
}
