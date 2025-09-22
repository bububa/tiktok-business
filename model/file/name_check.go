package file

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// NameCheckRequest 查验文件名是否重名 API Request
type NameCheckRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	NameCheckFile
	// Files 若想查验单个文件名是否重名，file_name 和 file_type必填。
	// 若想查验多个文件名是否重名，files必填。
	// 想要查验是否重名的文件名信息。
	// 最大数量：20。
	Files []NameCheckFile `json:"files,omitempty"`
}

type NameCheckFile struct {
	// FileName 文件名称
	FileName string `json:"file_name,omitempty"`
	// FileType 文件类型。
	// 枚举值：VIDEO（视频），IMAGE（图片）。
	// 默认值：VIDEO
	FileType string `json:"file_type,omitempty"`
}

// Encode implements GetRequest interface
func (r *NameCheckRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.Files) > 0 {
		values.Set("files", string(util.JSONMarshal(r.Files)))
	} else {
		values.Set("file_name", r.FileName)
		values.Set("file_type", r.FileType)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// NameCheckResponse 查验文件名是否重名 API Response
type NameCheckResponse struct {
	model.BaseResponse
	Data *NameCheckResult
}

type NameCheckResult struct {
	// Duplicate 仅当请求中未传入 files 时返回本字段。
	//
	// 单个文件名是否重名。
	//
	// 允许值：true，false。
	Duplicate bool `json:"duplicate,omitempty"`
	// DuplicateTemplateID 仅当 duplicate 为 true 时返回。
	// 已使用此文件名的图片或视频的素材 ID。
	DuplicateTemplateID string `json:"duplicate_template_id,omitempty"`
	// FileName 文件名称
	FileName string `json:"file_name,omitempty"`
	// BatchResults 仅当请求中传入 files 时返回本字段。
	// 文件名列表的查验结果。
	BatchResults []NameCheckResult `json:"batch_results,omitempty"`
}
