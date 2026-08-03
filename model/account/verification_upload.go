package account

import (
	"io"

	"github.com/bububa/tiktok-business/model"
)

const (
	defaultVerificationImageFile1Name = "image_file1"
	defaultVerificationImageFile2Name = "image_file2"
)

// VerificationUploadRequest is the multipart request for uploading account verification documents.
type VerificationUploadRequest struct {
	// ImageFile1 is the primary verification document or the front image of a personal ID.
	ImageFile1 io.Reader `json:"-"`
	// ImageFile1Name is the filename sent for ImageFile1.
	ImageFile1Name string `json:"-"`
	// ImageFile2 is an optional additional document or the back image of a personal ID.
	ImageFile2 io.Reader `json:"-"`
	// ImageFile2Name is the filename sent for ImageFile2.
	ImageFile2Name string `json:"-"`
}

// Encode implements model.UploadRequest.
func (r *VerificationUploadRequest) Encode() []model.UploadField {
	fields := make([]model.UploadField, 0, 2)
	if r.ImageFile1 != nil {
		fields = append(fields, model.UploadField{
			Reader: r.ImageFile1,
			Key:    "image_file1",
			Value:  verificationImageFileName(r.ImageFile1Name, defaultVerificationImageFile1Name),
		})
	}
	if r.ImageFile2 != nil {
		fields = append(fields, model.UploadField{
			Reader: r.ImageFile2,
			Key:    "image_file2",
			Value:  verificationImageFileName(r.ImageFile2Name, defaultVerificationImageFile2Name),
		})
	}
	return fields
}

func verificationImageFileName(name string, fallback string) string {
	if name != "" {
		return name
	}
	return fallback
}

// VerificationUploadResponse is the response for uploading account verification documents.
type VerificationUploadResponse struct {
	model.BaseResponse
	Data *VerificationUploadResult `json:"data,omitempty"`
}

// VerificationUploadResult contains the IDs of uploaded verification documents.
type VerificationUploadResult struct {
	QualificationImageIDs []string `json:"qualification_image_ids,omitempty"`
}
