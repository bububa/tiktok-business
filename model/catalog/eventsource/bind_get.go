package eventsource

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// BindGetRequest 获取事件源绑定信息 API Request
type BindGetRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
}

// Encode implements GetRequest interface
func (r *BindGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// BindGetResponse 获取事件源绑定信息 API Response
type BindGetResponse struct {
	model.BaseResponse
	Data *BindGetResult `json:"data,omitempty"`
}

type BindGetResult struct {
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// EventSources 事件源数据。每个事件源是一个对象。
	EventSources []EventSource `json:"event_sources,omitempty"`
}
