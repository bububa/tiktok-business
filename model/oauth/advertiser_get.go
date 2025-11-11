package oauth

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdvertiserGetRequest 获取广告账号列表 API Request
type AdvertiserGetRequest struct {
	// AppID 开发者申请的应用app_id，可通过应用管理界面查看
	AppID string `json:"app_id,omitempty"`
	// Secret 开发者应用的私钥，可通过应用管理界面查看
	Secret string `json:"secret,omitempty"`
}

// Encode implements GetRequest interface
func (r *AdvertiserGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("app_id", r.AppID)
	values.Set("secret", r.Secret)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AdvertiserGetResponse 获取广告账号列表 API Response
type AdvertiserGetResponse struct {
	model.BaseResponse
	Data struct {
		List []Advertiser `json:"list,omitempty"`
	} `json:"data"`
}

// Advertiser 广告主
type Advertiser struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserName 广告主名字
	AdvertiserName string `json:"advertiser_name,omitempty"`
}
