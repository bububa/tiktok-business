package identity

import "github.com/bububa/tiktok-business/enum"

// Identity 认证身份详情
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
	// ProfileImage 与该认证身份绑定的 TikTok 账号的主页头像临时 URL。
	// 有效期：约 48 小时。过期时间包含在 x-expires 参数后的 URL 中，格式为以秒为单位的 Epoch/Unix 时间戳。
	// URL 过期后，您需要调用 /campaign/gmv_max/info/ 获取新的 URL。
	ProfileImage string `json:"profile_image,omitempty"`
	// DisplayName 与该认证身份绑定的 TikTok 账号的展示名称
	DisplayName string `json:"display_name,omitempty"`
	// UserName 与该认证身份绑定的 TikTok 账号的用户名
	UserName string `json:"user_name,omitempty"`
	// AvailableStatus 认证身份的可用状态。只对TT_USER和BC_AUTH_TT类型的身份认证有效。
	// 枚举值： AVAILABLE, NO_VALID_BIND_ACCOUNT, SCOPE_UNAVAILABLE, IS_PRIVATE_ACCOUNT, NOT_BUSINESS_ACCOUNT
	AvailableStatus enum.IdentityAvailableStatus `json:"available_status,omitempty"`
	// CanPushVideo 该 BC_AUTH_TT 或 TT_USER 认证身份是否可以创建或编辑视频。
	// 支持的值：true， false。
	CanPushVideo bool `json:"can_push_video,omitempty"`
	// CanPullVideo 该 BC_AUTH_TT 或 TT_USER 认证身份是否可以获取 TikTok 账户下的所有视频。
	// 支持的值：true， false
	CanPullVideo bool `json:"can_pull_video,omitempty"`
	// CanUseLiveAds 该 BC_AUTH_TT 或 TT_USER 认证身份是否可以访问 TikTok 账户的直播间。
	// 支持的值：true， false
	CanUseLiveAds bool `json:"can_use_live_ads,omitempty"`
	// CanManageMessage 该 BC_AUTH_TT 或 TT_USER 认证身份是否可以管理私信。
	// 支持的值：true， false
	CanManageMessage bool `json:"can_manage_message,omitempty"`
	// ShopID 仅当 identity_type 为 TTS_TT 时返回本字段。
	// 与 TikTok Shop 关联的 TikTok 用户类型认证身份绑定的 TikTok Shop 的 ID
	ShopID string `json:"shop_id,omitempty"`
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
