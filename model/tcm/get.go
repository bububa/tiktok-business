package tcm

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 通过访问令牌获取 TTCM 账号 API Request
type GetRequest struct {
	// AppID 开发者应用 ID
	AppID string `json:"app_id,omitempty"`
	// Secret 开发者应用密钥
	Secret string `json:"secret,omitempty"`
	// Page 当前页数。
	// 默认值: 1。
	// 取值范围: ≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值: 10。
	// 取值范围: 1-100。
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("app_id", r.AppID)
	values.Set("secret", r.Secret)
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

// GetResponse 通过访问令牌获取 TTCM 账号 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// TCMAccountIDs 该令牌可以访问的 TTCM 账号的 ID 列表
	TCMAccountIDs []string `json:"tcm_account_ids,omitempty"`
}
