package oauth

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TokenRequest 获得访问令牌 API Request
type TokenRequest struct {
	// ClientID 开发者应用 ID
	ClientID string `json:"client_id,omitempty"`
	// ClientSecret 开发者应用密钥
	ClientSecret string `json:"client_secret,omitempty"`
	// Code 您的广告主授权 URL 生成的授权码（auth_code）。
	// 授权码有效期为 1 小时且只能使用一次。
	// 若想了解您使用的 API 类别对应的授权码获取方式，可查看不同 API 的授权流程对比中的"OAuth 授权 URL"行。
	Code string `json:"code,omitempty"`
	// GrantType 生成访问令牌的方式。
	// 允许的枚举值：
	// authorization_code：通过授权码生成访问令牌。
	GrantType string `json:"grant_type,omitempty"`
}

// Encode implements PostRequest interface
func (r *TokenRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

type TokenResponse struct {
	model.BaseResponse
	AccessToken
}
