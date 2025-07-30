package oauth

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AccessTokenRequest 获取长期访问令牌 API Request
type AccessTokenRequest struct {
	// AppID 开发者应用 ID。
	// 若想获取应用 ID，可以导航至应用管理 > App Detail > 基本信息。
	AppID string `json:"app_id,omitempty"`
	// Secret 开发者应用密钥。
	// 若想获取密钥，可以导航至应用管理 > App Detail > 基本信息。
	Secret string `json:"secret,omitempty"`
	// AuthCode 您的广告主授权 URL 生成的授权码（auth_code）。
	// 授权码有效期为 1 小时且只能使用一次。
	// 若想了解您使用的 API 类别对应的授权码获取方式，可查看不同 API 的授权流程对比中的"OAuth 授权 URL"行。
	AuthCode string `json:"auth_code,omitempty"`
}

// Encode implements PostRequest interface
func (r *AccessTokenRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

type AccessTokenResponse struct {
	model.BaseResponse
	Data *AccessToken `json:"data,omitempty"`
}

type AccessToken struct {
	// AccessToken 用于权限验证的长期访问令牌。
	AccessToken string `json:"access_token,omitempty"`
	// AdvertiserID 该访问令牌可以访问的广告账号列表。
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// Scope 该访问令牌所拥有的权限。
	// 您可查看权限范围了解权限范围 ID（scope的值）对应的权限。
	Scope []int64 `json:"scope,omitempty"`
	// TokenType 令牌类型。枚举值：Bearer （不记名令牌，拥有该令牌后任何人均可访问相关资源）。
	TokenType string `json:"token_type,omitempty"`
	// ExpiresIn 访问令牌的剩余有效时长，以秒为单位。
	ExpiresIn int64 `json:"expires_in,omitempty"`
	// RefreshToken 用于刷新访问令牌的刷新令牌。刷新令牌有效期为 1 年。
	RefreshToken string `json:"refresh_token,omitempty"`
	// RefreshTokenExpiresIn 刷新令牌的剩余有效时长，以秒为单位。
	RefreshTokenExpiresIn int64 `json:"refresh_token_expires_in,omitempty"`
	// OpenID TikTok 账户的应用唯一识别 ID
	OpenID string `json:"open_id,omitempty"`
}
