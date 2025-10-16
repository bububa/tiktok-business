package feed

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// SwitchRequest 修改更新源定时更新状态 API Request
type SwitchRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedID 更新源ID
	FeedID string `json:"feed_id,omitempty"`
	// Status 本字段用于修改更新源对应的定时更新状态。枚举值：ON，OFF。
	// 注意：
	// 您仅可通过本字段传入一个与当前定时更新状态不同的状态。例如，若更新源对应的定时更新状态为ON，您不可将本字段重复设置为ON。
	// 若您想将定时更新状态从OFF 改为 ON，您需要确保更新源已同时设置了 uri，interval_type，和timezone这三个参数，否则更新会失败。
	// 若您想为现有更新源设置uri，interval_type ，和timezone，可使用 /catalog/feed/update/接口。
	Status enum.OnOff `json:"status,omitempty"`
}

// Encode implements PostRequest interface
func (r *SwitchRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
