package product

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest  删除商品 API Request
type GetRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// Page 当前页数。
	// 注意: 仅可最多返回100,000条数据。您需确保page与page_size的乘积不超过100,000。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。默认值：10，最大值：500。
	PageSize int `json:"page_size,omitempty"`
	// ProductIDs 商品ID列表。最大数量：1000。ID传错的话，不返回对应商品信息。
	// 注意：
	// product_ids和sku_ids 不可同时指定。
	// 对于短剧商品库，仅支持使用 product_ids 筛选商品。
	ProductIDs []string `json:"product_ids,omitempty"`
	// SkuIDs 通过SKU ID列表获取对应的商品。最大数量：1000。
	// 长度限制：每个 SKU ID 100 字符，不能包含表情符号。
	// 注意：
	// 不支持重复的SKU ID。
	// product_ids和sku_ids 不可同时指定。
	SkuIDs []string `json:"sku_ids,omitempty"`
	// ProductSetIDs 通过商品系列ID获取商品。最大数量：100
	ProductSetIDs []string `json:"product_set_ids,omitempty"`
	// Order 排序选项。如果传入了SKU ID或者商品ID对结果进行筛选，则排序设置无效。
	Order *GetOrder `json:"order,omitempty"`
	// Conditions 筛选条件。当你使用本对象声明筛选条件时，除了case_sensitive外的所有字段都为必填。如果您传入了sku_ids或product_ids，该字段将被忽略
	Conditions *GetConditions `json:"conditions,omitempty"`
}

// GetOrder 排序选项。如果传入了SKU ID或者商品ID对结果进行筛选，则排序设置无效。
type GetOrder struct {
	// OrderCondition 排序条件。目前只支持根据商品可用状态 PRODUCT_AVAILABILITY排序。本字段不能与custom_order同时传入。
	OrderCondition string `json:"order_condition,omitempty"`
	// CustomOrder 按字段排序设置。本字段不能与order_condition同时传入。
	CustomOrder []CustomOrder `json:"custom_order,omitempty"`
}

// CustomOrder 按字段排序设置。本字段不能与order_condition同时传入。
type CustomOrder struct {
	// Field 排序所依据的字段。支持的字段请参考下面的支持字段列表
	Field string `json:"field,omitempty"`
	// Type 排序方式。枚举值: ASC (升序), DES (降序)
	Type enum.OrderType `json:"type,omitempty"`
}

// GetConditions 筛选条件。当你使用本对象声明筛选条件时，除了case_sensitive外的所有字段都为必填。如果您传入了sku_ids或product_ids，该字段将被忽略
type GetConditions struct {
	And []model.Filter `json:"and,omitempty"`
	Or  []model.Filter `json:"or,omitempty"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if len(r.ProductIDs) > 0 {
		values.Set("product_ids", string(util.JSONMarshal(r.ProductIDs)))
	}
	if len(r.SkuIDs) > 0 {
		values.Set("sku_ids", string(util.JSONMarshal(r.SkuIDs)))
	}
	if len(r.ProductSetIDs) > 0 {
		values.Set("product_set_ids", string(util.JSONMarshal(r.ProductSetIDs)))
	}
	if r.Order != nil {
		values.Set("order", string(util.JSONMarshal(r.Order)))
	}
	if r.Conditions != nil {
		values.Set("conditions", string(util.JSONMarshal(r.Conditions)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 删除商品 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 商品列表
	List []Product `json:"list,omitempty"`
}
