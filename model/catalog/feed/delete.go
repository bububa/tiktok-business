package feed

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DeleteRequest 删除更新源 API Request
type DeleteRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedID 更新源ID
	FeedID string `json:"feed_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeleteResponse 删除更新源 API Response
type DeleteResponse struct {
	model.BaseResponse
	Data struct {
		// FeedID 更新源 ID
		FeedID model.Uint64 `json:"feed_id,omitempty"`
	} `json:"data"`
}
