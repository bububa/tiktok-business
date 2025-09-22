package file

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// FinishUploadRequest 完成分片上传 API Request
type FinishUploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// UploadID 上传任务ID
	UploadID string `json:"upload_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *FinishUploadRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// FinishUploadResponse 完成分片上传 API Response
type FinishUploadResponse struct {
	model.BaseResponse
	Data *File `json:"data,omitempty"`
}
