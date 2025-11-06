package post

import (
	"github.com/bububa/tiktok-business/util"
)

// AuthorizeRequest 延长 TikTok 帖子的授权有效期 API Request
type AuthorizeRequest struct {
	// BusinessID TikTok 账户的应用唯一识别 ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// ItemID TikTok 帖子 ID。
	// 若想获取 TikTok 账号的公开帖子，可使用/business/video/list/。
	// 示例：7123456789123456789。
	// 注意：若 TikTok 帖子无授权码，无论是未生成过授权码或生成的授权码已删除，将出现报错。
	ItemID string `json:"item_id,omitempty"`
	// AuthorizationDays 授权的有效期，单位为天。
	// 支持的值：7，30，60，365。
	// 默认值：30。
	AuthorizationDays int `json:"authorization_days,omitempty"`
}

// Encode implements PostRequest
func (r *AuthorizeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
