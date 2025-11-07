package gmvmax

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/ad"
	"github.com/bububa/tiktok-business/util"
)

// OccupiedCustomShopAdsListRequest 检查认证身份和商品是否在购物广告中占用 API Request
type OccupiedCustomShopAdsListRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID TikTok Shop ID。
	// 若想获取可用于 GMV Max 推广系列的 TikTok Shop，可使用 /gmv_max/store/list/ 并确认返回的 is_gmv_max_available 为 true
	StoreID string `json:"store_id,omitempty"`
	// OccupiedAssetType 想要检查是否被处于启用状态的视频购物广告，商品购物广告或直播购物广告占用的资产类型。
	// 枚举值：
	// IDENTITY_TT_USER: TikTok 用户（TT_USER）认证身份。
	// IDENTITY_BC_AUTH_TT: 添加到商务中心的 TikTok 用户（BC_AUTH_TT）认证身份。
	// IDENTITY_TTS_TT: TikTok Shop 关联的 TikTok 用户（TTS_TT）认证身份。
	// SPU：SPU
	OccupiedAssetType enum.OccupiedAssetType `json:"occupied_asset_type,omitempty"`
	// AssetIDs 资产 ID 列表。
	// 最大数量：1。
	// occupied_asset_type 为 IDENTITY_TT_USER 时，需通过本字段指定一个 TT_USER 认证身份的 ID。
	// occupied_asset_type 为 IDENTITY_BC_AUTH_TT 时，需通过本字段指定一个 BC_AUTH_TT 认证身份的 ID。
	// occupied_asset_type 为 IDENTITY_TTS_TT 时，需通过本字段指定一个 TTS_TT 认证身份的 ID。
	// occupied_asset_type 为SPU时，需通过本字段指定此 TikTok Shop 中一个商品的 SPU ID。
	// 若想获取认证身份列表，可使用/gmv_max/identity/get/。
	// 若想获取 TikTok Shop 中商品的 SPU ID 列表，可使用/store/product/get/。需将 ad_creation_eligible 设置为 GMV_MAX 并查看返回的item_group_id。
	AssetIDs []string `json:"asset_ids,omitempty"`
}

// Encode implements GetRequest
func (r *OccupiedCustomShopAdsListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("store_id", r.StoreID)
	values.Set("occupied_asset_type", string(r.OccupiedAssetType))
	values.Set("asset_ids", string(util.JSONMarshal(r.AssetIDs)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// OccupiedCustomShopAdsListResponse 检查认证身份和商品是否在购物广告中占用 API Response
type OccupiedCustomShopAdsListResponse struct {
	model.BaseResponse
	Data struct {
		// OccupiedShopAds 目前使用这些素材且处于启用状态的购物广告列表
		OccupiedShopAds []ad.Ad `json:"occupied_shop_ads,omitempty"`
	} `json:"data"`
}
