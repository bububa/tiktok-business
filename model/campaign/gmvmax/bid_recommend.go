package gmvmax

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// BidRecommendRequest 获取推荐 GMV Max ROI 目标和预算 API Request
type BidRecommendRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID TikTok Shop ID。
	// 若想获取可用于 GMV Max 推广系列的 TikTok Shop，可使用 /gmv_max/store/list/ 并确认返回的 is_gmv_max_available 为 true
	StoreID string `json:"store_id,omitempty"`
	// ShoppingAdsType GMV Max 推广系列类型。
	// 枚举值：
	// PRODUCT：商品 GMV Max 推广系列。
	// LIVE：直播 GMV Max 推广系列。
	ShoppingAdsType enum.ShoppingAdsType `json:"shopping_ads_type,omitempty"`
	// OptimizationGoal 优化目标。
	// 枚举值：
	// VALUE：总收入。
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// ItemGroupIDs 仅当shopping_ads_type 为 PRODUCT 时生效。
	// 商品 SPU（标准化产品单元）ID 列表。
	// 若您想在商品 GMV Max 推广系列中推广此 TikTok Shop 的所有商品，无需传入本字段。
	// 若您想在商品 GMV Max 推广系列中推广此 TikTok Shop 的特定商品，需通过本字段指定这些商品。
	// 最大数量：50。
	// 若想获取某个 TikTok Shop 中商品的 SPU ID 列表，可使用/store/product/get/。将 ad_creation_eligible 设置为 GMV_MAX 并从返回中挑选 status 为 AVAILABLE 且 gmv_max_ads_status 为 UNOCCUPIED 的 item_group_id 值。
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
	// IdentityID 仅当 shopping_ads_type 为 LIVE 时生效且必填。
	// 直播 GMV Max 推广系列在要使用的直播来源认证身份。
	// 若想获取直播 GMV Max 推广系列可用的认证身份，可调用/gmv_max/identity/get/并挑选 live_gmv_max_available 为 true 的认证身份
	IdentityID string `json:"identity_id,omitempty"`
}

// Encode implements GetRequest
func (r *BidRecommendRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("store_id", r.StoreID)
	values.Set("shopping_ads_type", string(r.ShoppingAdsType))
	values.Set("optimization_goal", string(r.OptimizationGoal))
	if len(r.ItemGroupIDs) > 0 {
		values.Set("item_group_ids", string(util.JSONMarshal(r.ItemGroupIDs)))
	}
	if r.IdentityID != "" {
		values.Set("identity_id", r.IdentityID)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// BidRecommendResponse 获取推荐 GMV Max ROI 目标和预算 API Response
type BidRecommendResponse struct {
	model.BaseResponse
	Data *BidRecommend `json:"data,omitempty"`
}

type BidRecommend struct {
	// RoasBid 为优化目标总收入推荐的 ROI 目标。
	// 建议的 ROI 目标是根据您的商品类目的平均 ROI 计算得出的
	RoasBid float64 `json:"roas_bid,omitempty"`
	// Budget 为优化目标总收入推荐的日预算，对应币种为广告账户的币种。
	// 建议的预算是根据过去 14 天您所选商品的平均广告支出或你的商品类别的平均广告支出计算得出的。建议增加预算以确保全天持续投放，因为预算越低，费用消耗得越快，还可能会限制推广系列的效果。
	// 若想获取某个广告账户的币种，可使用/advertiser/info/并查看返回的currency
	Budget float64 `json:"budget,omitempty"`
}
