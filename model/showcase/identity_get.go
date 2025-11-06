package showcase

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// IdentityGetRequest 获取广告账户下有橱窗权限的认证身份 API Request
type IdentityGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest
func (r *IdentityGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// IdentityGetResponse 获取广告账户下有橱窗权限的认证身份 API Response
type IdentityGetResponse struct {
	model.BaseResponse
	Data struct {
		// IdentityList 有橱窗权限的认证身份列表
		IdentityList []Identity `json:"identity_list,omitempty"`
	} `json:"data"`
}
