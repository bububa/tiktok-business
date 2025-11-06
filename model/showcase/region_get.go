package showcase

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// RegionGetRequest 通过认证身份获取橱窗可定向的地区 API Request
type RegionGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// IdentityID 与橱窗绑定且有橱窗权限的认证身份的ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。
	// 枚举值：TT_USER （TikTok用户）， BC_AUTH_TT（添加到商务中心的TikTok用户）。
	// 您可以查看认证身份了解不同认证身份类型。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 当 identity_type 为 BC_AUTH_TT 时必填。
	// 与添加到商务中心的 TikTok 用户这一认证身份绑定的商务中心的 ID。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
}

// Encode implements GetRequest
func (r *RegionGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("identity_id", r.IdentityID)
	values.Set("identity_type", string(r.IdentityType))
	if r.IdentityAuthorizedBcID != "" {
		values.Set("identity_authorized_bc_id", r.IdentityAuthorizedBcID)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// RegionGetResponse 通过认证身份获取橱窗可定向的地区 API Response
type RegionGetResponse struct {
	model.BaseResponse
	Data struct {
		// RegionCodes 该橱窗（即该认证身份）可以定向的国家或地区代码列表。
		// 例如，["ID"]。
		RegionCodes []string `json:"region_codes,omitempty"`
	} `json:"data"`
}
