package set

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取商品系列列表 API Request
type GetRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// ProductSetID 商品系列 ID。
	// 若未不传入本字段，系统将返回与商品库（catalog_id）绑定的所有商品系列。
	// 若传入了本字段，系统将返回指定的商品系列。
	ProductSetID string `json:"product_set_id,omitempty"`
	// ReturnProductCount 是否返回参数 product_count。
	// 支持的值：true，false。
	// 默认值：true。
	// 注意：若商品系列中的商品过多，返回 product_count 可能导致超时错误。若想避免此情况下的超时错误，可将本字段设置为 false。
	ReturnProductCount *bool `json:"return_product_count,omitempty"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	if r.ProductSetID != "" {
		values.Set("product_set_id", r.ProductSetID)
	}
	if r.ReturnProductCount != nil {
		if *r.ReturnProductCount {
			values.Set("return_product_count", "true")
		} else {
			values.Set("return_product_count", "false")
		}
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 获取商品系列列表 API Response
type GetResponse struct {
	model.BaseResponse
	Data struct {
		// List 商品系列列表
		List []ProductSet `json:"product_set,omitempty"`
	} `json:"data"`
}

type ProductSet struct {
	// ProductCount 仅当请求中return_product_count为true或未传入时返回本字段。
	// 该商品系列中的商品数量。
	ProductCount int `json:"product_count,omitempty"`
	// ProductSetID 商品系列ID
	ProductSetID string `json:"product_set_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// ProductSetName 商品系列名称
	ProductSetName string `json:"product_set_name,omitempty"`
	// Conditions 筛选规则
	Conditions *Conditions `json:"conditions,omitempty"`
}

// Conditions 筛选条件
type Conditions struct {
	And []model.Filter `json:"and,omitempty"`
	Or  []model.Filter `json:"or,omitempty"`
}
