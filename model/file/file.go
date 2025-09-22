package file

import "github.com/bububa/tiktok-business/model"

// File 文件
type File struct {
	// FileID 文件的ID，您可以通过file_id获取文件库中的文件，上传到TikTok投放平台素材库。有效时间24h。
	FileID string `json:"file_id,omitempty"`
	// FileName 文件名称
	FileName string `json:"file_name,omitempty"`
	// Signature 文件的MD5
	Signature string `json:"signature,omitempty"`
	// FileSize 文件的大小，单位为B
	FileSize int64 `json:"file_size,omitempty"`
	// Size 文件大小，单位为字节
	Size int64 `json:"size,omitempty"`
	// CreateTime 文件上传时间。格式为：YYYY-MM-DD HH:MM:SS，UTC时区。示例：2020-06-10 07:39:14
	CreateTime model.DateTime `json:"create_time,omitzero"`
}
