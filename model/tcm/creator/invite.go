package creator

import (
	"encoding/json"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InviteRequest 邀请创作者加入 TTCM API Request
type InviteRequest struct {
	// TCMAccountID TikTok Creator Marketplace 账户 ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// HandleNames TikTok 创作者的用户名。最多可指定 20 个用户名。
	HandleNames []string `json:"handle_names,omitempty"`
	// AdvertiserID 广告账户 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest
func (r *InviteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// InviteResponse 邀请创作者加入 TTCM API Response
type InviteResponse struct {
	model.BaseResponse
	Data *InviteResult `json:"data,omitempty"`
}

type InviteResult struct {
	// InTCM 已加入 TTCM的创作者
	InTCM []string `json:"in_tcm,omitempty"`
	// Invited 满足所有资格要求且已获得邀请的创作者
	Invited []bool `json:"invited,omitempty"`
	// Rejected 未通过资格检查的创作者
	Rejected []bool `json:"rejected,omitempty"`
	// RejectReasons 针对每个创作者用户名的拒绝原因，帮助 API 合作伙伴了解创作者为何被拒绝加入 TTCM
	RejectReasons json.RawMessage `json:"reject_reasons,omitempty"`
}
