package creator

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// PublicRequest 获取公开账户洞察 API Request
type PublicRequest struct {
	// TCMAccountID TikTok Creator Marketplace 账户 ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// HandleName 您想要搜索的 TikTok 创作者的用户名
	HandleName string `json:"handle_name,omitempty"`
}

// Encode implements GetRequest
func (r *PublicRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("tcm_account_id", r.TCMAccountID)
	values.Set("handle_name", r.HandleName)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// PublicResponse 获取公开账户洞察 API Response
type PublicResponse struct {
	model.BaseResponse
	Data *Creator `json:"data,omitempty"`
}
