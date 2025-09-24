package pixel

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取Pixel列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Code Pixel 代码。如传入则按照 Pixel 代码进行筛选。
	Code string `json:"code,omitempty"`
	// PixelID Pixel ID。如传入则按照 Pixel ID 进行筛选。
	PixelID string `json:"pixel_id,omitempty"`
	// Name Pixel 名称。模糊匹配，如传入则筛选名称包含传入数据的 Pixel。
	Name string `json:"name,omitempty"`
	// OrderBy 排序方式。枚举值: EARLIEST_CREATE（最早创建优先）,LATEST_CREATE（最晚创建优先）。 默认值: EARLIEST_CREATE
	OrderBy string `json:"order_by,omitempty"`
	// Filtering 筛选条件
	Filtering *ListFilter `json:"filtering,omitempty"`
	// Page 当前页数。默认值: 1。取值范围: ≥ 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。默认值: 10。取值范围: 1-20
	PageSize int `json:"page_size,omitempty"`
}

// ListFilter 筛选条件
type ListFilter struct {
	// AvailableForCatalogOnly Pixel 是否可作为事件源与商品库绑定。默认值：false。若想将Pixel与商品库绑定，您可以使用/catalog/eventsource/bind/。
	AvailableForCatalogOnly bool `json:"available_for_catalog_only"`
}

// Encode implements GetRequest interface
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if r.Code != "" {
		values.Set("code", r.Code)
	}
	if r.PixelID != "" {
		values.Set("pixel_id", r.PixelID)
	}
	if r.Name != "" {
		values.Set("name", r.Name)
	}
	if r.OrderBy != "" {
		values.Set("order_by", r.OrderBy)
	}
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

// ListResponse 获取Pixel列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// Pixels Pixel 列表
	Pixels []Pixel `json:"pixels,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
