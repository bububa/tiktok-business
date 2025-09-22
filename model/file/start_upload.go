package file

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// StartUploadRequest 请求分片上传 API Request
type StartUploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Size 文件大小(以字节为单位，不小于20MB)
	Size int64 `json:"size,omitempty"`
	// ContentType 文件类型。
	// 枚举值：image，music，video，playable
	ContentType enum.UploadContentType `json:"content_type,omitempty"`
	// Name 文件名称，不超过100个字符
	Name string `json:"name,omitempty"`
}

// Encode implements PostRequest interface
func (r *StartUploadRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StartUploadResponse 请求分片上传 API Response
type StartUploadResponse struct {
	model.BaseResponse
	Data *StartUploadResult `json:"data,omitempty"`
}

type StartUploadResult struct {
	// UploadID 上传任务ID，后续请求都需要填入这个ID
	UploadID string `json:"upload_id,omitempty"`
	// FileName 文件名称，不超过100个字符
	FileName string `json:"file_name,omitempty"`
	// StartOffset 第一个分片的开始偏移量
	StartOffset int64 `json:"start_offset,omitempty"`
	// EndOffset 第一个分片的结束偏移量
	EndOffset int64 `json:"end_offset,omitempty"`
}
