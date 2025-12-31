package term

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// ConfirmRequest 签订协议 API Request
type ConfirmRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// TermType 协议的类型。
	// 枚举值: LeadAds
	TermType enum.TermType `json:"term_type,omitempty"`
}

// Encode implements PostRequest
func (r *ConfirmRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
