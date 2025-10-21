package product

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UploadRequest 通过 JSON 结构上传商品 API Request
type UploadRequest struct {
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
func (r *UploadRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UploadResponse 通过 JSON 结构上传商品 API Response
type UploadResponse struct {
	model.BaseResponse
	Data struct {
		// FeedLogID 商品处理日志ID
		FeedLogID string `json:"feed_log_id,omitempty"`
	} `json:"data"`
}
