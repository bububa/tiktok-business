package store

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取广告账户下可用的商店 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID 您想筛选的第一方商店（TikTok Shop）的商店ID。 若未传入本字段，则将返回指定的广告账户下所有的可用store_id信息。
	// 注意：欲获取 TikTok Shop 的商店 ID，您可以调用/bc/asset/get/：
	// 当返回中的asset_type 为 TIKTOK_SHOP时，对应的asset_id为TikTok Shop 的商店 ID。
	StoreID string `json:"store_id,omitempty"`
	// StoreType 您想筛选的商店类型。
	// 枚举值：TIKTOK_SHOP （TikTok Shop）。
	// 若未传入本字段，则将会返回广告账户下所有商店的结果。
	StoreType enum.StoreType `json:"store_type,omitempty"`
}

// Encode implements GetRequest interface
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if r.StoreID != "" {
		values.Set("store_id", r.StoreID)
	}
	if r.StoreType != "" {
		values.Set("store_type", string(r.StoreType))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ListResponse 获取广告账户下可用的商店 API Response
type ListResponse struct {
	model.BaseResponse
	Data struct {
		// Stores 广告账户下可用商店的信息
		Stores []Store `json:"stores,omitempty"`
	} `json:"data"`
}

// Store 广告账户下商店的信息
type Store struct {
	// StoreAuthorizedBcID 已经授权可管理该商店（store_id）的商务中心的ID。 经授权后多个商务中心可管理同一商店。
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// StoreID TikTok Shop 的商店 ID。请查看创建商品来源为TikTok Shopping的视频购物广告-创建广告组，了解如何将store_id用于创建购物广告。
	StoreID string `json:"store_id,omitempty"`
	// StoreType 商店类型。枚举值：TIKTOK_SHOP （TikTok Shop）
	StoreType enum.StoreType `json:"store_type,omitempty"`
	// StoreName 商店名称
	StoreName string `json:"store_name,omitempty"`
	// StoreCode 仅当商店的store_type为TIKTOK_SHOP时返回。
	// 第一方商店（TikTok Shop）的商店代码。
	StoreCode string `json:"store_code,omitempty"`
	// CatalogID 与该商店绑定的商品库的ID。一个商品库仅可与一个商店绑定
	CatalogID string `json:"catalog_id,omitempty"`
	// TargetingRegionCodes 商店可用于定向的地域代码。请查看附录-地区代码了解代码代表的具体地域。
	TargetingRegionCodes []string `json:"targeting_region_codes,omitempty"`
}
