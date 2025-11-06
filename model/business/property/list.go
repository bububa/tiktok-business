package property

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取添加至广告账户的 URL 资源列表 API Request
type ListRequest struct {
	// AppID 开发者应用ID。
	// 若想获取应用ID，您可以导航至应用管理 > App Detail > 基本信息。
	AppID string `json:"app_id,omitempty"`
	// Secret 开发者应用密钥。
	// 若想获取密钥，可以导航至应用管理 > App Detail > 基本信息。
	Secret string `json:"secret,omitempty"`
}

// Encode implements GetRequest
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("app_id", r.AppID)
	values.Set("secret", r.Secret)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ListResponse 获取添加至广告账户的 URL 资源列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data struct {
		// URLPropertyInfoList 已添加至广告账户的 URL 资源的信息。
		URLPropertyInfoList []URLProperty `json:"url_property_info_list,omitempty"`
	} `json:"data"`
}
