package property

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AddRequest 将 URL 资源添加至广告账户 API Request
type AddRequest struct {
	// AppID 开发者应用ID。
	// 若想获取应用ID，您可以导航至应用管理 > App Detail > 基本信息。
	AppID string `json:"app_id,omitempty"`
	// Secret 开发者应用密钥。
	// 若想获取密钥，可以导航至应用管理 > App Detail > 基本信息。
	Secret string `json:"secret,omitempty"`
	// URLPropertyMeta 您想要添加和验证所有权的 URL 资源的信息
	URLPropertyMeta *URLProperty `json:"url_property_meta,omitempty"`
}

// Encode implements PostRequest
func (r *AddRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AddResponse 将 URL 资源添加至广告账户 API Response
type AddResponse struct {
	model.BaseResponse
	Data struct {
		URLPropertyInfo *URLProperty `json:"url_property_info,omitempty"`
	} `json:"data"`
}
