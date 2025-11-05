package showcase

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ProductGetRequest 获取橱窗中的可用商品 API Request
type ProductGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// IdentityID 与橱窗绑定且有橱窗权限的认证身份的ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。
	// 枚举值：TT_USER （TikTok用户）， BC_AUTH_TT（添加到商务中心的TikTok用户）。
	// 您可以查看认证身份了解不同认证身份类型。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 当 identity_type 为 BC_AUTH_TT 时必填。
	// 与添加到商务中心的 TikTok 用户这一认证身份绑定的商务中心的 ID。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// RegionCodes 您想要使用橱窗商品定向的国家或地区代码列表
	RegionCodes []string `json:"region_codes,omitempty"`
	// Filtering 筛选条件
	Filtering *ProductGetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值：1。取值范围：≥1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值：10。取值范围：[1, 1000]
	PageSize int `json:"page_size,omitempty"`
}

type ProductGetFilter struct {
	// ItemGroupIDs 商品的 SPU （标准化产品单元）ID 列表。
	// 最大数量：10。
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
}

// Encode implements GetRequest
func (r *ProductGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("identity_id", r.IdentityID)
	values.Set("identity_type", string(r.IdentityType))
	if r.IdentityAuthorizedBcID != "" {
		values.Set("identity_authorized_bc_id", r.IdentityAuthorizedBcID)
	}
	values.Set("region_codes", string(util.JSONMarshal(r.RegionCodes)))
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

// ProductGetResponse 获取橱窗中的可用商品 API Response
type ProductGetResponse struct {
	model.BaseResponse
	Data *ProductGetResult `json:"data,omitempty"`
}

type ProductGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ShowcaseProducts 可用于定向所指定国家或地区的橱窗商品列表
	ShowcaseProducts []Product `json:"showcase_products,omitempty"`
}
