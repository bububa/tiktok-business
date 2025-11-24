package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// PartnerGetRequest 获取合作伙伴 API Request
type PartnerGetRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// Page 当前页数。
	// 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值：10
	PageSize int `json:"page_size,omitempty"`
	// Filtering 筛选条件
	Filtering *PartnerGetFilter `json:"filtering,omitempty"`
}

// PartnerGetFilter 筛选条件
type PartnerGetFilter struct {
	// Name 合作伙伴姓名
	Name string `json:"name,omitempty"`
}

// Encode implements GetRequest
func (r *PartnerGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
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

// PartnerGetResponse 获取合作伙伴 API Response
type PartnerGetResponse struct {
	model.BaseResponse
	Data *PartnerGetResult `json:"data,omitempty"`
}

type PartnerGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 合作伙伴列表
	List []Partner `json:"list,omitempty"`
}

// Partner 合作伙伴
type Partner struct {
	// BcID 合作伙伴商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// BcName 合作伙伴名
	BcName string `json:"bc_name,omitempty"`
}
