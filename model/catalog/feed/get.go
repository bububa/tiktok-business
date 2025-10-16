package feed

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取更新源 API Request
type GetRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedID 您想获取的商品库更新源 ID。如果不传入，则返回商品库的所有更新源。
	FeedID string `json:"feed_id,omitempty"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	if r.FeedID != "" {
		values.Set("feed_id", r.FeedID)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 获取更新源 API Response
type GetResponse struct {
	model.BaseResponse
	Data struct {
		// FeedList 更新源列表
		FeedList []Feed `json:"feed_list,omitempty"`
	} `json:"data"`
}
