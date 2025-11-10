package identity

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取认证身份列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// IdentityType 认证身份类型。枚举值：CUSTOMIZED_USER, AUTH_CODE, TT_USER, BC_AUTH_TT。如果不传入，则返回所有结果。获取更多相关信息，请访问认证身份
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。当identity_type为BC_AUTH_TT时必填
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// Filtering 筛选条件。
	// 注意：本字段仅当identity_type 为 CUSTOMIZED_USER 或不传入identity_type时生效。
	Filtering *GetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 取值范围：1-100。
	// 默认值：20
	PageSize int `json:"page_size,omitempty"`
}

// GetFilter 筛选条件。
type GetFilter struct {
	// Keyword 筛选关键词。
	// 当 identity_type 为 CUSTOMIZED_USER时，您可以通过本字段传入显示名称（display_name）进行模糊匹配
	Keyword string `json:"keyword,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("identity_type", string(r.IdentityType))
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

// GetResponse 获取认证身份列表 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// IdentityList 认证身份列表
	IdentityList []Identity `json:"identity_list,omitempty"`
}
