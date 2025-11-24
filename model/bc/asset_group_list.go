package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AssetGroupListRequest 获取资产组列表 API Request
type AssetGroupListRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// Filtering 筛选条件
	Filtering *AssetGroupListFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 取值范围：1-50。
	// 默认值：10
	PageSize int `json:"page_size,omitempty"`
}

type AssetGroupListFilter struct {
	// Keyword 筛选关键词。
	Keyword string `json:"keyword,omitempty"`
}

// Encode implements GetRequest
func (r *AssetGroupListRequest) Encode() string {
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

// AssetGroupListResponse 获取资产组列表 API Response
type AssetGroupListResponse struct {
	model.BaseResponse
	Data *AssetGroupListResult `json:"data,omitempty"`
}

type AssetGroupListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AssetGroups 资产组列表
	AssetGroups []AssetGroup `json:"asset_groups,omitempty"`
}
