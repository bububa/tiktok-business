package brand

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取 TTO Brand Profiles API Request
type GetRequest struct {
	// TTOTCMAccountID TTO Creator Marketplace 账户 ID
	TTOTCMAccountID string `json:"tto_tcm_account_id,omitempty"`
	// BrandProfileIDs Brand Profile ID 列表。
	// 最大数量：20。
	BrandProfileIDs []string `json:"brand_profile_ids,omitempty"`
	// Page 当前页数。
	// 默认值：1。
	// 取值范围：≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值：10。
	// 取值范围：1-100。
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("tto_tcm_account_id", r.TTOTCMAccountID)
	if len(r.BrandProfileIDs) > 0 {
		values.Set("brand_profile_ids", string(util.JSONMarshal(r.BrandProfileIDs)))
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

// GetResponse 获取 TTO Brand Profiles API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// Brands Brand Profile 列表
	Brands []Profile `json:"brands,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
