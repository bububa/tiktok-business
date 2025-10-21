package catalog

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 修改商品库名称 API Request
type UpdateRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// Name 商品库名称。
	// 长度限制：128 字符。
	Name string `json:"name,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 修改商品库名称 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data struct {
		// CatalogID 商品库ID
		CatalogID string `json:"catalog_id,omitempty"`
	} `json:"data"`
}
