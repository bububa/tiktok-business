package gmvmax

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ExclusiveAuthorizationGetRequest 获取广告账号的 TikTok Shop 专属授权状态 API Request
type ExclusiveAuthorizationGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID TikTok Shop ID。
	// 若想获取可用于 GMV Max 推广系列的 TikTok Shop，可使用 /gmv_max/store/list/ 并确认返回的 is_gmv_max_available 为 true
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID 有权限访问该 TikTok Shop（store_id）的商务中心的 ID
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
}

// Encode implements GetRequest
func (r *ExclusiveAuthorizationGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("store_id", r.StoreID)
	values.Set("store_authorized_bc_id", r.StoreAuthorizedBcID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ExclusiveAuthorizationGetResponse 获取广告账号的 TikTok Shop 专属授权状态 API Response
type ExclusiveAuthorizationGetResponse struct {
	model.BaseResponse
	Data *ExclusiveAuthorization `json:"data,omitempty"`
}

// ExclusiveAuthorization 有关该 TikTok Shop 的 GMV Max 专属授权信息
type ExclusiveAuthorization struct {
	// AdvertiserID 广告账号 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserName 已获得该 TikTok Shop 的专属授权，可使用此店铺创建 GMV Max 推广系列的唯 一广告账户的名称
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// AdvertiserStatus 已获得该 TikTok Shop 的专属授权，可使用此店铺创建 GMV Max 推广系列的 唯一广告账户的状态。
	// 枚举值：
	// STATUS_ENABLE：审核通过。
	// STATUS_CONFIRM_FAIL：未通过审核。
	// STATUS_PENDING_CONFIRM：正在审核。
	// STATUS_LIMIT：暂时停用。
	// STATUS_CONTRACT_PENDING：合同未生效。
	// STATUS_DISABLE：永久停用。
	// STATUS_PENDING_CONFIRM_MODIFY：修改待审核。
	// STATUS_PENDING_VERIFIED待验证。
	// STATUS_SELF_SERVICE_UNAUDITED：自助开户待验证资质。
	// STATUS_WAIT_FOR_BPM_AUDIT：CRM 系统待审核。
	// STATUS_CONFIRM_FAIL_END：CRM 系统未通过审核。
	// STATUS_CONFIRM_MODIFY_FAIL：修改未通过审核
	AdvertiserStatus enum.AdvertiserStatus `json:"advertiser_status,omitempty"`
	// AuthorizationStatus 该广告账户使用此 TikTok Shop 创建 GMV Max 推广系列的专属授权状态。
	// 枚举值：
	// EFFECTIVE：该广告账号目前拥有 GMV Max 专属授权。
	// INEFFECTIVE：该广告账号拥有过 GMV Max 专属授权，但专属授权已取消。
	// UNAUTHORIZED：该广告账号尚未拥有 GMV Max 专属授权
	AuthorizationStatus enum.ExclusiveAuthorizationStatus `json:"authorization_status,omitempty"`
	// StoreID TikTok Shop 的 ID
	StoreID string `json:"store_id,omitempty"`
	// IdentityID 仅当 TikTok Shop 设置了官方 TikTok 账号时返回本字段。
	// 与 TikTok Shop 的官方 TikTok 账号绑定的认证身份 ID
	IdentityID string `json:"identity_id,omitempty"`
}
