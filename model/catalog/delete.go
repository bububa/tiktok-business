package catalog

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DeleteRequest 删除商品库 API Request
type DeleteRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeleteResponse 删除商品库名称 API Response
type DeleteResponse struct {
	model.BaseResponse
	Data struct {
		// CatalogID 商品库ID
		CatalogID string `json:"catalog_id,omitempty"`
	} `json:"data"`
}
