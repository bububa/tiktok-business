package identity

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// DeleteRequest 删除认证身份 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// IdentityID 认证身份ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。枚举值：CUSTOMIZED_USER, AUTH_CODE, TT_USER, BC_AUTH_TT。如果不传入，则返回所有结果。获取更多相关信息，请访问认证身份
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
}

// Encode implements PostRequest
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
