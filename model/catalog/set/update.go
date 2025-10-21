package set

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 更新商品系列 API Request
type UpdateRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// ProductSetID 商品系列 ID。
	ProductSetID string `json:"product_set_id,omitempty"`
	// ProductSetName 商品系列名称(小于28个字符)
	ProductSetName string `json:"product_set_name,omitempty"`
	// Conditions 筛选条件。
	// 注意：若您通过本对象传入多个筛选条件，建议您在所有筛选条件中传入的预期值总数不超过1024个。否则，很可能出现请求超时
	Conditions *Conditions `json:"conditions,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新商品系列 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *UpdateResult `json:"data,omitempty"`
}

type UpdateResult struct {
	// ProductSetID 商品系列 ID。
	ProductSetID string `json:"product_set_id,omitempty"`
	// ProductSetName 商品系列名称
	ProductSetName string `json:"product_set_name,omitempty"`
	// ProductCount 该商品系列下的商品数
	ProductCount int `json:"product_count,omitempty"`
}
