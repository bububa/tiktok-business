package feed

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建更新源 API Request
type CreateRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedName 更新源名称。
	FeedName string `json:"feed_name,omitempty"`
	// UpdateMode 更新模式。
	// 枚举值:
	// OVERWRITE：覆盖数据源。若历史商品在新的商品库中不存在，则历史商品将被删除。若存在新商品，将在商品库中添加新商品；若商品已存在，则会用上传的商品信息更新原有的商品信息。
	// INCREMENTAL：增量式更新。若历史商品在新的商品库中不存在，历史商品不会被删除。若存在新商品，将在商品库中添加新商品；若商品已存在，则会用上传的商品信息更新原有的商品信息。
	UpdateMode enum.CatalogFeedUpdateMode `json:"update_mode,omitempty"`
	// ScheduleParam 定时更新的设置
	ScheduleParam *ScheduleParam `json:"schedule_param,omitempty"`
}

// Encode implements PostRequest interface
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建更新源 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Feed `json:"data,omitempty"`
}

type (
	UpdateRequest  = CreateRequest
	UpdateResponse = CreateResponse
)
