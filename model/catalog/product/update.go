package product

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 更新商品 API Request
type UpdateRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedId 更新源ID
	FeedID string `json:"feed_id,omitempty"`
	// Products 商品列表
	Products []Product `json:"products,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新商品 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data struct {
		// FeedLogID 商品处理日志ID
		FeedLogID string `json:"feed_log_id,omitempty"`
	} `json:"data"`
}
