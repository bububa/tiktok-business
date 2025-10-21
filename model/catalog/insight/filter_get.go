package insight

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// FilterGetRequest 获取商品库分析数据的筛选条件 API Request
type FilterGetRequest struct {
	// BcID 商务中心 ID。
	// 若想获取您有权访问的商务中心列表，可使用/bc/get/。
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库 ID。
	// 需指定一个包含至少 20 个商品的电商商品库。
	// 若想获取某个商务中心中的电商商品库列表，可使用 /catalog/get/，筛选对应的 catalog_type 为 ECOM 的 catalog_id 值。
	// 若想验证商品库中包含的商品至少为 20 个，可使用 /catalog/overview/，查看返回的字段 approved，rejected 和 processing 的值之和是否大于等于 20。
	CatalogID string `json:"catalog_id,omitempty"`
	// FilterType 筛选条件的类型。
	// 枚举值：
	// CATEGORY_ID：商品类目 ID。
	// BRAND：商品品牌名称。
	// AVAILABILITY：商品库存状态。
	FilterType enum.CatalogInsightFilterType `json:"filter_type,omitempty"`
	// Page 当前页数。
	// 默认值: 1。取值范围: ≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值: 10。取值范围：1-200。
	// 注意：您可根据自身需要决定是否使用 page 和 page_size 字段分批次获取分页结果。不过，需注意返回中不包含分页信息。
	// 若您不想获取分页结果，可将 page_size 设置为 50，从而在单个返回中获取完整结果
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r *FilterGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	values.Set("filter_type", string(r.FilterType))
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

// FilterGetResponse 获取商品库分析数据的筛选条件 API Response
type FilterGetResponse struct {
	model.BaseResponse
	Data *FilterGetResult `json:"data,omitempty"`
}

type FilterGetResult struct {
	// Brands 仅当请求中的 filter_type 设置为 BRAND 时返回本字段。
	// 可用于筛选该商品库中商品的品牌名称列表。
	Brands []string `json:"brands,omitempty"`
	// Availabilities 仅当请求中的 filter_type 设置为  AVAILABILITY 时返回本字段。
	//  可用于筛选该商品库中商品的库存状态列表。
	Availabilities []enum.ProductAvailability `json:"availabilities,omitempty"`
	// Categories 仅当请求中的 filter_type 设置为 CATEGORY_ID 时返回本字段。
	// 可用于筛选该商品库中商品的商品类目有关信息。
	Categories []Category `json:"categories,omitempty"`
}
