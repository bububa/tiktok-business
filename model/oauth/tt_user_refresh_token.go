package oauth

import "github.com/bububa/tiktok-business/util"

// TTUserRefreshTokenRequest 刷新短期访问令牌
type TTUserRefreshTokenRequest struct {
	// ClientID 开发者应用 ID
	ClientID string `json:"client_id,omitempty"`
	// ClientSecret 开发者应用密钥
	ClientSecret string `json:"client_secret,omitempty"`
	// GrantType 生成访问令牌的方式。
	// 允许的枚举值：
	// authorization_code：通过授权码生成访问令牌。
	GrantType string `json:"grant_type,omitempty"`
	// RefreshToken 刷新令牌，用于刷新现有的访问令牌。
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Encode implement PostRequest interface
func (r *TTUserRefreshTokenRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
