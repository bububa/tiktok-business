package gmvmax

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// StoreShopAdUsageCheckRequest 检查 TikTok Shop 是否可用于商品 GMV Max 推广系列 API Request
type StoreShopAdUsageCheckRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID TikTok Shop ID。
	// 若想获取可用于 GMV Max 推广系列的 TikTok Shop，可使用 /gmv_max/store/list/ 并确认返回的 is_gmv_max_available 为 true
	StoreID string `json:"store_id,omitempty"`
}

// Encode implements GetRequest
func (r *StoreShopAdUsageCheckRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("store_id", r.StoreID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// StoreShopAdUsageCheckResponse 检查 TikTok Shop 是否可用于商品 GMV Max 推广系列 API Response
type StoreShopAdUsageCheckResponse struct {
	model.BaseResponse
	Data *StoreShopAdUsageCheckResult `json:"data,omitempty"`
}

type StoreShopAdUsageCheckResult struct {
	// PromoteAllProductsAllowed 是否支持使用此 TikTok Shop 创建推广此 TikTok Shop 中的所有商品的商品 GMV Max 推广系列。
	// 支持的值：true，false
	PromoteAllProductsAllowed bool `json:"promote_all_products_allowed,omitempty"`
	// IsRunningCustomShopAds 此 TikTok Shop 是否已在处于启用状态的视频购物广告或商品购物广告中占用。
	// 支持的值：true，false
	IsRunningCustomShopAds bool `json:"is_running_custom_shop_ads,omitempty"`
}
