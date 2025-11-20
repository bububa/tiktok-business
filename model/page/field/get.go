package field

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取即时表单字段值 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// PageID 即时表单ID
	PageID string `json:"page_id,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("page_id", r.PageID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 获取即时表单字段值 API Response
type GetResponse struct {
	model.BaseResponse
	Data *Field `json:"data,omitempty"`
}

// Field 即使表单字段
type Field struct {
	// Fields 即时表单字段信息
	Fields []string `json:"fields,omitempty"`
	// MetaData 即时表单元信息
	MetaData *Meta `json:"meta,omitempty"`
}

// Meta 即时表单元信息
type Meta struct {
	// PageID 即时表单ID
	PageID string `json:"page_id,omitempty"`
	// CreateTime 即时表单的创建时间
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// PageName 即时表单名称
	PageName string `json:"page_name,omitempty"`
	// PageURL 即时表单URL
	PageURL string `json:"page_url,omitempty"`
}
