package tool

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TargetingCategoryRecommendRequest 获取推荐兴趣及行为分类 API Request
type TargetingCategoryRecommendRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// RegionCodes 要投放的国家或地区代码列表。 例如: ["US", "JP", "KR"]
	RegionCodes []string `json:"region_codes,omitempty"`
	// AppID 您想要推广的移动应用ID。可选字段，仅当您要推广的是一款移动应用时可传入
	AppID string `json:"app_id,omitempty"`
}

// Encode implements GetRequest interface
func (r *TargetingCategoryRecommendRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("region_codes", string(util.JSONMarshal(r.RegionCodes)))
	if r.AppID != "" {
		values.Set("app_id", string(r.AppID))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// TargetingCategoryRecommendResponse 获取推荐兴趣及行为分类 API Response
type TargetingCategoryRecommendResponse struct {
	model.BaseResponse
	Data *TargetingCategoryRecommendResult `json:"data,omitempty"`
}

type TargetingCategoryRecommendResult struct {
	// InterestCategories 推荐的兴趣分类
	InterestCategories []InterestCategory `json:"interest_categories,omitempty"`
	// ActionCategories 推荐的行为分类
	ActionCategories []ActionCategory `json:"action_categories,omitempty"`
}
