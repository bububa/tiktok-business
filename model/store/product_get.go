package store

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ProductGetRequest 获取 TikTok Shop 中的商品 API Request
type ProductGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// BcID 商务中心 ID。
	// 若想获取您有权限访问的商务中心列表，请使用/bc/get/。
	BcID string `json:"bc_id,omitempty"`
	// StoreID TikTok Shop ID。
	// 若想获取与广告账号绑定的 TikTok Shop ID 列表，可使用/store/list/。
	StoreID string `json:"store_id,omitempty"`
	// Filtering 筛选条件
	Filtering *ProductGetFilter `json:"filtering,omitempty"`
	// SortField 排序字段。
	// 枚举值：
	// min_price：返回字段 min_price。
	// historical_sales：返回字段historical_sales 。
	SortField string `json:"sort_field,omitempty"`
	// SortType 排序类型。
	// 枚举值：
	// ASC：升序。
	// DESC：降序。
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Page 当前页数。
	// 默认值: 1。
	// 取值范围: ≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	PageSize int `json:"page_size,omitempty"`
}

// ProductGetFilter 筛选条件
type ProductGetFilter struct {
	// ItemGroupIDs 商品的 SPU（标准化产品单元）ID 列表。
	// 最大数量：10。
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
	// AdCreationEligible 商品可用于的广告或推广系列类型。
	// 枚举值：
	// CUSTOM_SHOP_ADS：购物广告，包括视频购物广告和商品购物广告。
	// GMV_MAX：商品 GMV Max 推广系列。
	// 若不传入本字段，返回中将不包含商品是否对于商品 GMV Max 推广系列可用的信息。
	AdCreationEligible string `json:"ad_creation_eligible,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"product_name,omitempty"`
}

// Encode implements GetRequest interface
func (r *ProductGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("store_id", r.StoreID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.AdvertiserID != "" {
		values.Set("advertiser_id", r.AdvertiserID)
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

// ProductGetResponse 获取 TikTok Shop 中的商品 API Response
type ProductGetResponse struct {
	model.BaseResponse
	Data *ProductGetResult `json:"data,omitempty"`
}

type ProductGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// StoreProducts TikTok Shop 中的商品列表。
	StoreProducts []Product `json:"store_products,omitempty"`
}

// Product TikTok Shop 中的商品
type Product struct {
	// StoreID TikTok Shop ID
	StoreID string `json:"store_id,omitempty"`
	// ItemGroupID 商品的 SPU ID
	ItemGroupID string `json:"item_group_id,omitempty"`
	// CatalogID 商品（item_group_id）所属的商品库 ID。
	// 注意：自 2024 年 6 月 30 日起，本字段返回 null 值。
	CatalogID string `json:"catalog_id,omitempty"`
	// Title 商品的标题
	Title string `json:"title,omitempty"`
	// ProductImageURL 商品图片的 URL
	ProductImageURL string `json:"product_image_url,omitempty"`
	// MinPrice 商品的最低价格。
	// 注意：若您仅为商品设置了一个固定价格，则 min_price 和 max_price 将返回相同的值
	MinPrice model.Float64 `json:"min_price,omitempty"`
	// MaxPrice 商品的最高价格。
	// 注意：若您仅为商品设置了一个固定价格，则 min_price 和 max_price 将返回相同的值。
	MaxPrice model.Float64 `json:"max_price,omitempty"`
	// Currencty 商品价格（min_price 和 max_price）对应的币种代码。
	// 若想了解币种代码对应的币种，请查看不同币种的预算校验比例和取值范围。
	Currency string `json:"currency,omitempty"`
	// HistoricalSales 商品的历史销售额，即该商品已销售的数量
	HistoricalSales int64 `json:"historical_sales,omitempty"`
	// Category 商品的分类
	Category string `json:"category,omitempty"`
	// Status 商品的状态。
	// 枚举值：
	// AVAILABLE：商品可用于创建广告。
	// NOT_AVAILABLE：商品不可用于创建广告。
	Status enum.StoreProductStatus `json:"status,omitempty"`
	// GMVMaxAdsStatus 仅当请求中传入筛选字段 ad_creation_eligible 时返回本字段。
	// 该商品在处于启用状态的商品 GMV Max 推广系列中的状态。
	// 枚举值：
	// OCCUPIED：该商品已被处于启用状态的商品 GMV Max 推广系列占用。
	// UNOCCUPIED: 该商品未被处于启用状态的商品 GMV Max 推广系列占用。
	// 注意：若 status 为 AVAILABLE 且 gmv_max_ads_status 为 UNOCCUPIED，则该商品可用于商品 GMV Max 推广系列。
	GMVMaxAdsStatus enum.StoreProductGMVMaxAdsStatus `json:"gmv_max_ads_status,omitempty"`
	// IsRunningCustomShopAds 仅当请求中筛选字段 ad_creation_eligible 设置为 GMV_MAX 时返回本字段。
	// 该商品是否被处于启用状态的视频购物广告或商品购物广告占用。
	// 支持的值：true，false。
	// 若 is_running_custom_shop_ads 为 true，您可调用 /gmv_max/occupied_custom_shop_ads/list/ 并将 occupied_asset_type 设置为 SPU，从而找出正在使用此商品的处于启用状态的视频购物广告或商品购物广告。若想暂停这些广告，可使用/ad/status/update/ ，并将 operation_status 设置为 DISABLE。
	IsRunningCustomShopAds bool `json:"is_running_custom_shop_ads,omitempty"`
}
