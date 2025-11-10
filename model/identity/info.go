package identity

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InfoRequest 获取认证身份详情 API Request
type InfoRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// IdentityID 认证身份ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。枚举值：CUSTOMIZED_USER, AUTH_CODE, TT_USER, BC_AUTH_TT。如果不传入，则返回所有结果。获取更多相关信息，请访问认证身份
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。当identity_type为BC_AUTH_TT时必填
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
}

// Encode implements GetRequest
func (r *InfoRequest) Encode() string {
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

// InfoResponse 获取认证身份详情 API Response
type InfoResponse struct {
	model.BaseResponse
	Data struct {
		// IdentityInfo 认证身份信息
		IdentityInfo *Identity `json:"identity_info,omitempty"`
	} `json:"data"`
}
