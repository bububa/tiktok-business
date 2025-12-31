package term

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 查看协议 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Lang 想查看的协议的语言。
	// 枚举值: EN,JA,ZH。
	// 默认值: EN
	Lang string `json:"lang,omitempty"`
	// TermType 协议的类型。
	// 枚举值: LeadAds
	TermType enum.TermType `json:"term_type,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("term_type", string(r.TermType))
	if r.Lang != "" {
		values.Set("lang", r.Lang)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 查看协议 API Response
type GetResponse struct {
	model.BaseResponse
	Data struct {
		// Terms 协议的具体内容
		Terms []string `json:"terms,omitempty"`
	} `json:"data"`
}
