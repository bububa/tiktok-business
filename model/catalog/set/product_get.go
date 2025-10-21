package set

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ProductGetRequest 获取商品信息 API Request
type ProductGetRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// ProductSetID 商品系列 ID。
	ProductSetID string `json:"product_set_id,omitempty"`
	// Page 当前页数。默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。默认值：20
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r *ProductGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	values.Set("product_set_id", r.ProductSetID)
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

// ProductGetResponse 获取商品信息 API Response
type ProductGetResponse struct {
	model.BaseResponse
	Data *ProductGetResult `json:"data,omitempty"`
}

type ProductGetResult struct {
	// ProductSetID 商品系列 ID。
	ProductSetID string `json:"product_set_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// ProductCount 该商品系列下的商品数
	ProductCount int `json:"product_count,omitempty"`
	// Products 商品列表
	Products []Product `json:"products,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// Product 商品
type Product struct {
	// ProductID 商品 ID
	ProductID string `json:"product_id,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"product_name,omitempty"`
	// SkuID 仅为电商商品库中的商品返回本字段。
	// 电商商品的唯一 ID。
	SkuID string `json:"sku_id,omitempty"`
	// HotelID 仅为酒店商品库中的酒店返回本字段。
	// 酒店的唯一 ID
	HotelID string `json:"hotel_id,omitempty"`
	// FlightID 仅为航班商品库中的航班返回本字段。
	// 航班的唯一 ID。
	FlightID string `json:"flight_id,omitempty"`
	// DestinationID 仅为目的地商品库中的目的地返回本字段。
	// 目的地的唯一 ID。
	DestinationID string `json:"destination_id,omitempty"`
}
