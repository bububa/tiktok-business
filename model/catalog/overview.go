package catalog

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// OverviewRequest 获取商品库概览 API Request
type OverviewRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
}

// Encode implements GetRequest interface
func (r *OverviewRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// OverviewResponse 获取商品库概览 API Response
type OverviewResponse struct {
	model.BaseResponse
	Data *Overview `json:"data,omitempty"`
}

type Overview struct {
	// Approved 商品库中已通过审核的商品数量。
	Approved int `json:"approved,omitempty"`
	// Rejected 商品库中审核被拒的商品数量
	Rejected int `json:"rejected,omitempty"`
	// Processing 商品库中审核处理中的商品数量
	Processing int `json:"processing,omitempty"`
	// OrganicApproved 与TikTok Shopping绑定的已通过审核的商品数量
	OrganicApproved int `json:"organic_approved,omitempty"`
	// OrganicRejected 与TikTok Shopping绑定的审核被拒的商品数量
	OrganicRejected int `json:"organic_rejected,omitempty"`
	// OrganicProcessing 与TikTok Shopping绑定的审核处理中的商品数量
	OrganicProcessing int `json:"organic_processing,omitempty"`
}
