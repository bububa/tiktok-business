package oauth

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// RevokeTokenRequest 撤销长期访问令牌 API Request
type RevokeTokenRequest struct {
	// AppID 开发者应用 ID
	AppID string `json:"app_id,omitempty"`
	// Secret 开发者应用密钥
	Secret string `json:"secret,omitempty"`
	// AccessToken 您想要撤销的长期访问令牌
	AccessToken string `json:"access_token,omitempty"`
}

// Encode implements PostRequest interface
func (r *RevokeTokenRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// RevokeTokenResponse 撤销长期访问令牌 API Response
type RevokeTokenResponse struct {
	model.BaseResponse
	Data *RevokeTokenResult `json:"data,omitempty"`
}

type RevokeTokenResult struct {
	// AppID 开发者应用 ID
	AppID string `json:"app_id,omitempty"`
	// AdvertiserIDs 该访问令牌不再有权限访问的广告账户 ID 列表
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
}
