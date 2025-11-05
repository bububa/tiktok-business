package creator

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// StatusGetRequest 查询创作者 TTCM 状态 API Request
type StatusGetRequest struct {
	// TCMAccountID TikTok Creator Marketplace 账户 ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// HandleNames TikTok 创作者的用户名。最多可指定 20 个用户名。
	HandleNames []string `json:"handle_names,omitempty"`
	// AdvertiserID 广告账户 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest
func (r *StatusGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("tcm_account_id", r.TCMAccountID)
	values.Set("handle_names", string(util.JSONMarshal(r.HandleNames)))
	if r.AdvertiserID != "" {
		values.Set("advertiser_id", r.AdvertiserID)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// StatusGetResponse 查询创作者 TTCM 状态 API Response
type StatusGetResponse struct {
	model.BaseResponse
	Data struct {
		// OnboardingStatus 入驻状态数据
		OnboardingStatus []OnboardingStatus `json:"onboarding_status,omitempty"`
	} `json:"data"`
}

// OnboardingStatus 入驻状态数据
type OnboardingStatus struct {
	// HandlerName 创作者用户名
	HandlerName string `json:"handler_name,omitempty"`
	// Status 创作者入驻状态。枚举值：IN，NOT_IN，INVALID，INVITED
	Status enum.TCMCreatorStatus `json:"status,omitempty"`
}
