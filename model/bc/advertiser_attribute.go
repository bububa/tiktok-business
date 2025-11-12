package bc

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdvertiserAttributeRequest 获取广告账户的币种和注册地区 API Request
type AdvertiserAttributeRequest struct {
	// BcID 商务中心 ID
	BcID string `json:"bc_id,omitempty"`
}

// Encode implements GetRequest
func (r *AdvertiserAttributeRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AdvertiserAttributeResponse 获取广告账户的币种和注册地区 API Response
type AdvertiserAttributeResponse struct {
	model.BaseResponse
	Data *AdvertiserAttribute `json:"data,omitempty"`
}

type AdvertiserAttribute struct {
	// Currencies 该商务中心内广告账户的注册地对应的地区代码。
	Currencies []string `json:"currencies,omitempty"`
	// RegionCodes 该商务中心内广告账户对应的币种代码
	RegionCodes []string `json:"region_codes,omitempty"`
}
