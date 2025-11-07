package gmvmax

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
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
		IdentityList []Identity `json:"identity_list,omitempty"`
	} `json:"data"`
}

// Identity 认证身份
type Identity struct {
	// IdentityID 认证身份 ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。
	// 枚举值：
	// AUTH_CODE：授权用户。当您使用授权码访问 TikTok 账户或 TikTok 帖子时，系统将创建此类型的认证身份。
	// TT_USER：TikTok 用户。当您绑定 TikTok For Business 账户与 TikTok 企业号，或将 TikTok For Business 账户与普通 TikTok 账户绑定后，将 TikTok 账户升级为企业号时，系统将创建此类型的认证身份。
	// BC_AUTH_TT：添加到商务中心的 TikTok 用户。当您将 TikTok 账号添加到商务中心，并从 TikTok 账号主获得访问批准后，系统将创建此类型的认证身份。
	// TTS_TT：TikTok Shop 关联的 TikTok 用户。当您为 TikTok Shop 设置官方 TikTok 账号后，系统将创建此类型的认证身份
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 仅当 identity_type 为 BC_AUTH_TT 时返回本字段。
	// 与添加到商务中心的 TikTok 用户类型认证身份绑定的商务中心的 ID。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// IdentityAuthorizedShopID 仅为部分 BC_AUTH_TT 认证身份返回本字段。
	// 与添加到商务中心的 TikTok 用户类型认证身份绑定的 TikTok Shop 的 ID
	IdentityAuthorizedShopID string `json:"identity_authorized_shop_id,omitempty"`
	// ShopID 仅当 identity_type 为 TTS_TT 时返回本字段。
	// 与 TikTok Shop 关联的 TikTok 用户类型认证身份绑定的 TikTok Shop 的 ID
	ShopID string `json:"shop_id,omitempty"`
	// ProfileImage 与该认证身份绑定的 TikTok 账号的主页头像临时 URL。
	// 有效期：约 48 小时。过期时间包含在 x-expires 参数后的 URL 中，格式为以秒为单位的 Epoch/Unix 时间戳。
	// URL 过期后，您需要调用 /campaign/gmv_max/info/ 获取新的 URL。
	ProfileImage string `json:"profile_image,omitempty"`
	// DisplayName 与该认证身份绑定的 TikTok 账号的展示名称
	DisplayName string `json:"display_name,omitempty"`
	// UserName 与该认证身份绑定的 TikTok 账号的用户名
	UserName string `json:"user_name,omitempty"`
	// IsRunningCustomShopAds 该认证身份是否被处于启用状态的视频购物广告，商品购物广告或 直 播购物广告占用。
	// 支持的值：true，false
	IsRunningCustomShopAds bool `json:"is_running_custom_shop_ads,omitempty"`
	// ProductGMVMaxAvailable 该认证身份是否可用于商品 GMV Max 推广系列。
	// 支持的值：true，false
	ProductGMVMaxAvailable bool `json:"product_gmv_max_available,omitempty"`
	// LiveGMVMaxAvailable 该认证身份是否可用于直播 GMV Max 推广系列。
	// 支持的值：true，false
	LiveGMVMaxAvailable bool `json:"live_gmv_max_available,omitempty"`
	// UnavailableReason 仅当 live_gmv_max_available 为 false 时返回本字段。
	// 该认证身份不可用于直播 GMV Max 推广系列的原因。
	// 枚举值：
	// OCCUPIED：该认证身份已被一个处于启用状态的直播 GMV Max 推广系列占用。
	// UNAUTHORIZED：该认证身份无在直播 GMV Max 推广系列中使用的授权。
	UnavailableReason string `json:"unavailable_reason,omitempty"`
}
