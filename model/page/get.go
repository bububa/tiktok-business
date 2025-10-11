package page

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取页面 ID
type GetRequest struct {
	// AdvertiserID advertiser_id和 library_id 必传其一。
	// 如果相应的页面在广告账号下（即未迁移到商务中心），必须传入广告账号 ID。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// LibraryID advertiser_id和 library_id 必传其一。
	// 如果相应的页面已经迁移到商务中心下，必须传入相应的表单库 ID。您可以将页面从广告账号下迁移到商务中心。
	LibraryID string `json:"library_id,omitempty"`
	// Page 当前页数。
	// 默认值: 1。取值范围: ≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值: 10。取值范围: 1-100。
	PageSize int `json:"page_size,omitempty"`
	// Status 页面状态。
	// 枚举值: EDITED（草稿）,PUBLISHED（可投放）。
	PageStatus enum.PageStatus `json:"page_status,omitempty"`
	// Title 页面标题。用于筛选标题中 包含 本字段值的页面
	Title string `json:"title,omitempty"`
	// UpdateTimeRange 用于筛选页面的页面更新时间范围。
	UpdateTimeRange *TimeRange `json:"update_time_range,omitempty"`
	// BusinessType 页面类型。
	// 枚举值: LEAD_GEN(即时表单（线索表单）), STORE_FRONT(商品橱窗页), APP_PROFILE_PAGE (应用介绍页), TIKTOK_INSTANT_PAGE ( 自定义页面 (自定义TikTok即时体验页面 ) ), SHOP_ADS_PLP (店铺广告产品列表页), SHOP_ADS_PDP (店铺广告产品详情页)，POP_UP_FORM(私信页面) 。
	// 默认值: LEAD_GEN。
	BusinessType enum.BusinessType `json:"business_type,omitempty"`
	// BusinessTypes 页面类型。
	// 枚举值：LEAD_GEN(即时表单（线索表单）), STORE_FRONT(商品橱窗页), APP_PROFILE_PAGE (应用介绍页), TIKTOK_INSTANT_PAGE ( 自定义页面 (自定义TikTok即时体验页面 ) ), SHOP_ADS_PLP (店铺广告产品列表页), SHOP_ADS_PDP (店铺广告产品详情页) ，POP_UP_FORM(私信页面)。
	// 注意: 如果该字段与 business_type同时传入, business_type的值将会被忽略。
	BusinessTypes []enum.BusinessType `json:"business_types,omitempty"`
}

type TimeRange struct {
	// Start 传入update_time_range时本字段必填。
	// 起始时间，YYYY-MM-DD HH:MM:SS（UTC 时间）。
	// 示例：2024-01-01 00:00:00。
	Start string `json:"start,omitempty"`
	// End 传入update_time_range时本字段必填。
	// 结束时间，YYYY-MM-DD HH:MM:SS（UTC 时间）。
	// 示例：2024-01-01 00:00:00。
	End string `json:"end,omitempty"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	if r.AdvertiserID != "" {
		values.Set("advertiser_id", r.AdvertiserID)
	}
	if r.LibraryID != "" {
		values.Set("library_id", r.LibraryID)
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.PageStatus != "" {
		values.Set("page_status", string(r.PageStatus))
	}
	if r.Title != "" {
		values.Set("title", r.Title)
	}
	if r.UpdateTimeRange != nil {
		values.Set("update_time_range", string(util.JSONMarshal(r.UpdateTimeRange)))
	}
	if r.BusinessType != "" {
		values.Set("business_type", string(r.BusinessType))
	}
	if len(r.BusinessTypes) > 0 {
		values.Set("business_types", string(util.JSONMarshal(r.BusinessTypes)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 获取页面 ID API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 页面列表
	List []Page `json:"list,omitempty"`
}
