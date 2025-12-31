package term

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CheckRequest 检查协议状态 API Request
type CheckRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// TermType 协议的类型。
	// 枚举值: LeadAds
	TermType enum.TermType `json:"term_type,omitempty"`
}

// Encode implements GetRequest
func (r *CheckRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("term_type", string(r.TermType))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// CheckResponse 检查协议状态 API Response
type CheckResponse struct {
	model.BaseResponse
	Data struct {
		// Confirmed 协议签订状态。TRUE表示签订成功，FALSE表示签订未完成。
		Confirmed bool `json:"confirmed,omitempty"`
	} `json:"data"`
}
