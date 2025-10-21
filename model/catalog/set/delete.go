package set

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DeleteRequest 删除商品系列 API Request
type DeleteRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// ProductSetID 商品系列 ID, 长度范围[1,10]。
	ProductSetIDs []string `json:"product_set_ids,omitempty"`
}

// Encode implements PostRequest interface
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeleteResponse 更新商品系列 API Response
type DeleteResponse struct {
	model.BaseResponse
	Data struct {
		// ProductSetIDs 已删除的商品系列ID。
		ProductSetIDs []string `json:"product_set_ids,omitempty"`
	} `json:"data"`
}
