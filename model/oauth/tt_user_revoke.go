package oauth

import "github.com/bububa/tiktok-business/util"

// TTUserRevokeRequest 撤销短期访问令牌
type TTUserRevokeRequest struct {
	// ClientID 开发者应用 ID
	ClientID string `json:"client_id,omitempty"`
	// ClientSecret 开发者应用密钥
	ClientSecret string `json:"client_secret,omitempty"`
	// AccessToken 您想要撤销的 TikTok 账号访问令牌。
	AccessToken string `json:"access_token,omitempty"`
}

// Encode implement PostRequest interface
func (r *TTUserRevokeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
