package gmvmax

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/identity"
	"github.com/bububa/tiktok-business/util"
)

// IdentityGetRequest 获取 GMV Max 推广系列可用的认证身份 API Request
type IdentityGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID TikTok Shop ID。
	// 若想获取可用于 GMV Max 推广系列的 TikTok Shop，可使用 /gmv_max/store/list/ 并确认返回的 is_gmv_max_available 为 true
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID 有权限访问该 TikTok Shop（store_id）的商务中心的 ID
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
}

// Encode implements GetRequest
func (r *IdentityGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("store_id", r.StoreID)
	values.Set("store_authorized_bc_id", r.StoreAuthorizedBcID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// IdentityGetResponse 获取 GMV Max 推广系列可用的认证身份 API Response
type IdentityGetResponse struct {
	model.BaseResponse
	Data struct {
		// IdentityList 与该 TikTok Shop 绑定的认证身份列表
		IdentityList []identity.Identity `json:"identity_list,omitempty"`
	} `json:"data"`
}
