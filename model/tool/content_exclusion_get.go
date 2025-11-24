package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ContentExclusionGetRequest 获取可用的内容排除类别 API Request
type ContentExclusionGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// ObjectiveType 推广目标。目前仅支持 REACH, VIDEO_VIEWS 和 ENGAGEMENT.
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// BrandSafetyType 品牌安全类型。默认值： NO_BRAND_SAFETY. 枚举值：
	// NO_BRAND_SAFETY: 不使用任何品牌安全解决方案，即使用全部库存。您的广告可能会在某些包含成人主题的内容周围展示。
	// EXPANDED_INVENTORY：使用TikTok的品牌安全解决方案，即使用扩展库存。您的广告将不会出现在不当内容或包含成人主题的内容后面。在下个API版本中EXPANDED_INVENTORY将替代NO_BRAND_SAFETY，并成为默认的品牌安全选项。
	// 注意 EXPANDED_INVENTORY目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// STANDARD_INVENTORY: 使用TikTok的品牌安全解决方案中的标准库存。您的广告将展示在适合大多数品牌的内容周围。
	// LIMITED_INVENTORY: 使用TikTok的品牌安全解决方案中的有限库存。您的广告将展示在不包含成人主题的内容周围。
	// 注意:
	//
	// 若您想获取行业敏感内容控制类型ID，您仅可以使用STANDARD_INVENTORY或LIMITED_INVENTORY。
	// 若您不在行业敏感内容控制的白名单，返回中的vertical_sensitivity_list将为null。
	BrandSafetyType enum.BrandSafetyType `json:"brand_safety_type,omitempty"`
}

// Encode implements GetRequest interface
func (r *ContentExclusionGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("objective_type", string(r.ObjectiveType))
	if r.BrandSafetyType != "" {
		values.Set("brand_safety_type", string(r.BrandSafetyType))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ContentExclusionGetResponse 获取可用的内容排除类别 API Response
type ContentExclusionGetResponse struct {
	model.BaseResponse
	Data *ContentExclusionGetResult `json:"data,omitempty"`
}

type ContentExclusionGetResult struct {
	// ExcludedCategoryList 内容排除类别列表
	ExcludedCategoryList []Category `json:"excluded_category_list,omitempty"`
	// VerticalSensitivityList 包含敏感内容的行业敏感内容控制类型
	VerticalSensitivityList []Category `json:"vertical_sensitivity_list,omitempty"`
}

// Category 内容类别
type Category struct {
	// CategoryID 类别ID
	CategoryID string `json:"category_id,omitempty"`
	// CategoryName 类别名称
	CategoryName string `json:"category_name,omitempty"`
	// SupportedRegions 类别支持的国家或地区代码
	SupportedRegions []string `json:"supported_regions,omitempty"`
	// Description 类别描述
	Description string `json:"description,omitempty"`
	// CategoryType 类别种类。对于excluded_category_list，将返回CATEGORY_TYPE_EXCLUSION（内容排除类别）
	CategoryType enum.CategoryType `json:"category_type,omitempty"`
}
