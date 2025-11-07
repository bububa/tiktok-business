package gmvmax

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// StoreListRequest 获取 GMV Max 推广系列可用的 TikTok Shop API Request
type StoreListRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest
func (r *StoreListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// StoreListResponse 获取 GMV Max 推广系列可用的 TikTok Shop API Response
type StoreListResponse struct {
	model.BaseResponse
	Data struct {
		// StoreList 该广告账号有权限访问的 TikTok Shop 列表
		StoreList []Store `json:"store_list,omitempty"`
	} `json:"data"`
}

// Store TikTok Shop
type Store struct {
	// StoreID TikTok Shop ID
	StoreID string `json:"store_id,omitempty"`
	// IsGMVMaxAvailable 该 TikTok Shop 是否可用于 GMV Max 推广系列。
	// 支持的值：
	// true：可用。
	// false：不可用
	IsGMVMaxAvailable bool `json:"is_gmv_max_available,omitempty"`
	// StoreAuthorizedBcID 有权限访问该 TikTok Shop（store_id）的商务中心的 ID
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// IsOwnerBc 此商务中心（store_authorized_bc_id）是否拥有该 TikTok Shop （store_id ）。
	// 支持的值：
	// true：此商务中心拥有该 TikTok Shop。
	// false：此商务中心是有权限访问该 TikTok Shop 的合作伙伴商务中心，但不拥有该 TikTok Shop
	IsOwnerBc bool `json:"is_owner_bc,omitempty"`
	// StoreAuthorizedBcInfo 有权限访问该 TikTok Shop 的商务中心的有关信息
	StoreAuthorizedBcInfo *StoreAuthorizedBcInfo `json:"store_authorized_bc_info,omitempty"`
	// ThumbnailURL TikTok Shop 缩略图 URL
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// StoreName TikTok Shop 名称
	StoreName string `json:"store_name,omitempty"`
	// StoreCode TikTok Shop 的店铺代码
	StoreCode string `json:"store_code,omitempty"`
	// TargetingRegionCodes TikTok Shop 可定向的地域的代码。
	// 若想查看某一地域代码对应的地域，可查看附录-地区代码。
	// 例如，VN 代表越南
	TargetingRegionCodes []string `json:"targeting_region_codes,omitempty"`
	// StoreStatus TikTok Shop 的状态。
	// 枚举值：
	// ACTIVE：可用。
	// INACTIVE：不可用。
	// NEW_CREATE：新建
	StoreStatus enum.StoreStatus `json:"store_status,omitempty"`
	// ExclusiveAuthorzedAdvertiserInfo 有关该 TikTok Shop 的 GMV Max 专属授权信息
	ExclusiveAuthorizedAdvertiserInfo *ExclusiveAuthorization `json:"exclusive_authorized_advertiser_info,omitempty"`
}

// StoreAuthorizedBcInfo 有权限访问该 TikTok Shop 的商务中心的有关信息
type StoreAuthorizedBcInfo struct {
	// BcID 商务中心 ID
	BcID string `json:"bc_id,omitempty"`
	// BcProfileImage 商务中心头像图片 URL
	BcProfileImage string `json:"bc_profile_image,omitempty"`
	// BcName 商务中心名称
	BcName string `json:"bc_name,omitempty"`
	// UserRole 用户（成员）在商务中心中的角色。
	// 枚举值：
	// ADMIN：管理员。管理员可以访问商务中心内的所有功能。
	// STANDARD：标准。标准成员只能对分配给他们的账户进行操作
	UserRole enum.UserRole `json:"user_role,omitempty"`
}
