package showcase

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Product 可用于定向所指定国家或地区的橱窗商品
type Product struct {
	// ItemGroupID 商品的 SPU ID
	ItemGroupID string `json:"item_group_id,omitempty"`
	// Title 商品的标题
	Title string `json:"title,omitempty"`
	// ProductImageURL 商品图片的 URL
	ProductImageURL string `json:"product_image_url,omitempty"`
	// MinPrice 商品的最低价格。
	// 注意：若您仅为商品设置了一个固定价格，则 min_price 和 max_price 将返回相同的值。
	MinPrice model.Float64 `json:"min_price,omitempty"`
	// MaxPrice 商品的最高价格。
	// 注意：若您仅为商品设置了一个固定价格，则 min_price 和 max_price 将返回相同的值。
	MaxPrice model.Float64 `json:"max_price,omitempty"`
	// Currency 商品价格（min_price 和 max_price）对应的币种代码。
	// 若想了解币种代码对应的币种，请查看不同币种的预算校验比例和取值范围。
	Currency string `json:"currency,omitempty"`
	// Category 商品的分类
	Category string `json:"category,omitempty"`
	// Status 商品的状态。
	// 枚举值：
	// AVAILABLE：商品可用于创建广告。
	// NOT_AVAILABLE：商品不可用于创建广告。
	Status enum.ProductAvailability `json:"status,omitempty"`
	// CatalogID 商品（item_group_id）所属的商品库 ID。
	// 注意：自 2024 年 6 月 30 日起，将不再返回本字段。
	CatalogID string `json:"catalog_id,omitempty"`
	// StoreID 商品（item_group_id）所属的店铺 ID。
	// 请注意本字段仅支持 TikTok Shop 店铺。
	StoreID string `json:"store_id,omitempty"`
}
