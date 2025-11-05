package ttuser

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TokenInfoGetRequest 获取已授权的创作者权限 API Request
type TokenInfoGetRequest struct {
	// AppID 开发者应用 ID
	AppID string `json:"app_id,omitempty"`
	// AccessToken 用于权限验证的访问令牌。
	// 欲获取在TTCM创作者授予TikTok创作者账户权限后生成的访问令牌，请查看获取或刷新访问令牌。
	AccessToken string `json:"access_token,omitempty"`
}

// Encode implement PostRequest interface
func (r *TokenInfoGetRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TokenInfoGetResponse 获取已授权的创作者权限 API Response
type TokenInfoGetResponse struct {
	model.BaseResponse
	Data *TokenInfo `json:"data,omitempty"`
}

type TokenInfo struct {
	// AppID 开发者应用ID
	AppID string `json:"app_id,omitempty"`
	// Scope 所指定访问令牌的权限范围。
	// 例如，"user.info.basic,video.list"
	Scope string `json:"scope,omitempty"`
	// CreatorID TikTok创作者账户的应用唯一识别ID。
	// 注意：您可以将本字段的值传给TikTok Creator Marketplace接口 /tcm/creator/authorized/，/tcm/order/update/creator/，和/tcm/creator/authorized/video/list/中的请求参数creator_id。
	CreatorID string `json:"creator_id,omitempty"`
}
