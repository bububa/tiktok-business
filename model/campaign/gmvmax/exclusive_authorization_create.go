package gmvmax

import "github.com/bububa/tiktok-business/util"

// ExclusiveAuthorizationCreateRequest 授予广告账号 TikTok Shop 的专属授权 API Request
type ExclusiveAuthorizationCreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID TikTok Shop ID。
	// 若想获取可用于 GMV Max 推广系列的 TikTok Shop，可使用 /gmv_max/store/list/ 并确认返回的 is_gmv_max_available 为 true
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID 有权限访问该 TikTok Shop（store_id）的商务中心的 ID
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
}

// Encode implements PostRequest
func (r *ExclusiveAuthorizationCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
