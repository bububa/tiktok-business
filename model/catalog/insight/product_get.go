package insight

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ProductGetRequest 获取热门商品列表 API Request
type ProductGetRequest struct {
	// BcID 商务中心 ID。
	// 若想获取您有权访问的商务中心列表，可使用/bc/get/。
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库 ID。
	// 需指定一个包含至少 20 个商品的电商商品库。
	// 若想获取某个商务中心中的电商商品库列表，可使用 /catalog/get/，筛选对应的 catalog_type 为 ECOM 的 catalog_id 值。
	// 若想验证商品库中包含的商品至少为 20 个，可使用 /catalog/overview/，查看返回的字段 approved，rejected 和 processing 的值之和是否大于等于 20。
	CatalogID string `json:"catalog_id,omitempty"`
	// Filtering 筛选条件
	Filtering *ProductGetFilter `json:"filtering,omitempty"`
}

// ProductGetFilter 筛选条件
type ProductGetFilter struct {
	// CategoryIDs 若传入了 filtering，需至少指定 category_ids，brands和 availabilities 其中之一。
	// TikTok 商品类目 ID 列表，可用于筛选分析数据。
	// 最大数量：50。
	// 需传入 /catalog/insight/filter/get/ 返回的 category_id 字段值中的一个或多个商品类目 ID。
	Categories []string `json:"categories,omitempty"`
	// Brands 若传入了 filtering，需至少指定 category_ids，brands和 availabilities 其中之一。
	// 商品品牌名称列表，可用于筛选分析数据。
	// 最大数量：50。
	// 需传入 /catalog/insight/filter/get/ 返回的 brands 字段值中的一个或多个品牌名称。
	Brands []string `json:"brands,omitempty"`
	// Availabilities 若传入了 filtering，需至少指定 category_ids，brands和 availabilities 其中之一。
	// 商品的库存状态列表，可用于筛选分析数据。
	// 枚举值：
	// IN_STOCK：有库存。
	// AVAILABLE_FOR_ORDER：可订购。
	// PREORDER：可预订。
	// OUT_OF_STOCK：缺货。
	// DISCONTINUED：停售。
	Availabilities []enum.ProductAvailability `json:"availabilities,omitempty"`
}

// Encode implements GetRequest interface
func (r *ProductGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ProductGetResponse 获取热门商品列表 API Response
type ProductGetResponse struct {
	model.BaseResponse
	Data struct {
		// ProductInsights 电商商品库中的热门商品，至多可获取前 50 个热门商品，列表按热门程度倒序排列。
		// 例如，列表中第一个商品排名第 1，即最热门的商品，第二个商品排名第 2，即第二热门的商品。
		ProductInsights []Product `json:"product_insights,omitempty"`
	} `json:"data"`
}

// Product 电商商品库中的热门商品
type Product struct {
	// ProductID 系统生成的商品 ID
	ProductID string `json:"product_id,omitempty"`
	// ImageURL 商品主图的 URL
	ImageURL string `json:"image_url,omitempty"`
	// Title 商品标题
	Title string `json:"title,omitempty"`
	// Description 关于商品的简短描述
	Description string `json:"description,omitempty"`
	// SkuID 广告主自定义的商品 SKU ID
	SkuID string `json:"sku_id,omitempty"`
	// CategoryInfo 商品的商品类目的有关信息
	CategoryInfo *Category `json:"category_info,omitempty"`
	// Brand 商品品牌名称
	Brand string `json:"brand,omitempty"`
	// Price 商品价格信息
	Price *PriceInfo `json:"price,omitempty"`
	// Availability 商品库存情况。
	// 枚举值：
	// IN_STOCK：有库存。
	// AVAILABLE_FOR_ORDER：可订购。
	// PREORDER：可预订。
	// OUT_OF_STOCK：缺货。
	// DISCONTINUED：停售。
	Availability enum.ProductAvailability `json:"availability,omitempty"`
}

// PriceInfo 价格信息
type PriceInfo struct {
	// Price 商品的基础价格。
	// 示例：30。
	Price float64 `json:"price,omitempty" csv:"price"`
	// Currency 货币单位。
	// 若指定本字段，则需填入商品库所定向地域对应的币种。
	// 若未指定本字段，则默认设置为商品库所定向地域对应的币种。
	// 示例：USD。
	Currency string `json:"currency,omitempty" csv:"currency"`
	// SalePrice 商品的促销价格（如果正在折价销售）。
	// 示例：20。
	SalePrice float64 `json:"sale_price,omitempty" csv:"sale_price"`
	// SalePriceEffectiveDate 促销开始和结束的日期及时间，格式为：YYYY-MM-DDT23:59+00:00/YYYY-MM-DDT23:59+00:00。
	// 开始日期应在结束日期之前。
	// 如果商品的促销价没有限定期限，请不要传入本字段。
	// 示例：2023-12-11T23:59+00:00/2023-12-15T23:59+00:00。
	SalePriceEffectiveDate string `json:"sale_price_effective_date,omitempty" csv:"sale_price_effective_date"`
	// TotalPrice 仅对酒店商品库中的酒店返回本字段。
	// 总价格。
	TotalPrice float64 `json:"total_price,omitempty" csv:"total_price"`
}
