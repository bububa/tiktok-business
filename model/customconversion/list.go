package customconversion

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取与事件源关联的自定义转化 API Request
type ListRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// EventSourceType 事件源类型。
	// 枚举值:
	// PIXEL: Pixel。
	// APP: 应用。
	EventSourceType enum.EventSourceType `json:"event_source_type,omitempty"`
	// EventSourceID 事件源 ID。
	// 当 event_source_type 为 PIXEL 时，指定一个 Pixel ID。
	// 要获取广告账户中的 Pixel ID，可使用/pixel/list/并检查返回的pixel_id。
	// 当 event_source_type 为 APP 时，指定自归因网络应用的应用平台 ID。
	// 要获取广告账户中自归因应用的应用平台 ID，可使用/app/list/。确认应用的 self_attribution_enabled 返回值为 true 并检查返回的 app_platform_id。
	EventSourceID string `json:"event_source_id,omitempty"`
	// SearchKeyword 用于搜索自定义转化的关键字。
	// 您可以通过自定义转化名称或自定义转化 ID 进行搜索。
	//
	// 按自定义转化名称搜索时，支持模糊匹配。
	// 按自定义转化 ID 搜索时，支持精确匹配。
	SearchKeyword string `json:"search_keyword,omitempty"`
	// SortField 排序字段。
	// 枚举值:
	// CREATE_TIME: 自定义转化的创建时间。
	// TOTAL_COUNT: 自定义转化的总转化数。
	// 默认值: CREATE_TIME。
	SortField string `json:"sort_field,omitempty"`
	// SortType 排序类型。
	// 枚举值:
	// DESC: 降序。
	// ASC: 升序。
	// 默认值: DESC。
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Page 当前页数。
	// 取值范围: ≥ 1。
	// 默认值: 1。
	Page int `json:"page,omitempty"`
	// PageSize 页面大小。
	// 取值范围: 1-1,000。
	// 默认值: 10。
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("event_source_type", string(r.EventSourceType))
	values.Set("event_source_id", r.EventSourceID)
	if r.SearchKeyword != "" {
		values.Set("search_keyword", r.SearchKeyword)
	}
	if r.SortField != "" {
		values.Set("sort_field", r.SortField)
	}
	if r.SortType != "" {
		values.Set("sort_type", string(r.SortType))
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

// ListResponse 获取与事件源关联的自定义转化 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// List 与事件源关联的自定义转化列表
	List []CustomConversion `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
