package post

import (
	"github.com/bububa/tiktok-business/util"
)

// AuthorizeDeleteRequest 删除 TikTok 帖子的授权码 API Request
type AuthorizeDeleteRequest struct {
	// BusinessID TikTok 账户的应用唯一识别 ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// ItemID TikTok 帖子 ID。
	// 若想获取 TikTok 账号的公开帖子，可使用/business/video/list/。
	// 示例：7123456789123456789。
	// 注意：若 TikTok 帖子无授权码，无论是未生成过授权码或生成的授权码已删除，将出现报错。
	ItemID string `json:"item_id,omitempty"`
}

// Encode implements PostRequest
func (r *AuthorizeDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
