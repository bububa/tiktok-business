package library

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取表单库 API Request
type GetRequest struct {
	// Page 返回中的页数
	Page int `json:"page,omitempty"`
	// PageSize 每页的记录数目
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 获取表单库 API Response
type GetResponse struct {
	model.BaseResponse
	Data struct {
		// List 表单库列表
		List []Library `json:"list,omitempty"`
	} `json:"dat"`
}

// Library 表单库
type Library struct {
	// LibraryID 表单库ID
	LibraryID string `json:"library_id,omitempty"`
	// LibraryName 表单库名称
	LibraryName string `json:"library_name,omitempty"`
	// AdvertiserID 广告账号ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CreateTime 表单库创建的时间，格式为YYYY-MM-DD HH:MM:SS。例如：2017-01-01 00:00:00
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// UpdateTime 表单库更新的时间，格式为 YYYY-MM-DD HH:MM:SS。例如： 2017-01-01 00:00:00
	UpdateTime model.DateTime `json:"update_time,omitzero"`
}
