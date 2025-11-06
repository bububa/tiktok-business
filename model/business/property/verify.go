package property

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// VerifyRequest 检查 URL 资源验证结果 API Request
type VerifyRequest struct {
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
func (r *VerifyRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// VerifyResponse 检查 URL 资源验证结果 API Response
type VerifyResponse struct {
	model.BaseResponse
	Data struct {
		URLPropertyInfo *URLProperty `json:"url_property_info,omitempty"`
	} `json:"data"`
}
