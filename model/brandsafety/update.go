package brandsafety

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 设置或更新广告账户的品牌安全中心设置 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CoverAllAdObjectives 是否将品牌安全设置（brand_safety_type、category_exclusion_ids 和 vertical_sensitivity_id）应用于所有推广目标。
	// 支持的值:
	// true：将品牌安全设置应用于所有目标，包括覆盖人数、视频播放量、社区互动、访问量、应用推广、线索收集、网站转化量和商品销量。
	// false：仅将品牌安全设置应用于覆盖人数、视频播放量和社区互动目标。
	// 注意: 本字段设置为 true 后，不允许更新为 false。
	CoverAllAdObjectives bool `json:"cover_all_ad_objectives"`
	// BrandSafetyType 库存筛选等级。
	// 所选等级将应用于“推荐”、“关注”、“个人资料页”和“搜索”推荐内容（广告会显示在要查看的搜索结果旁）的广告位。
	// 枚举值:
	// EXPANDED_INVENTORY：扩展库存。您的广告不会展示在明显不适当的内容前后，但可能会展示在成人主题内容前后。
	// STANDARD_INVENTORY：标准库存。您的广告将展示在适合大多数品牌的内容前后。
	// LIMITED_INVENTORY：有限库存。您的广告不会展示在包含成人主题的内容前后。
	// 注意: 若想更新品牌安全设置，您需同时传入 brand_safety_type、category_exclusion_ids 和 vertical_sensitivity_id。
	BrandSafetyType enum.BrandSafetyType `json:"brand_safety_type,omitempty"`
	// CategoryExclusionIDs 仅当 brand_safety_type 为 STANDARD_INVENTORY 或 LIMITED_INVENTORY 时生效。
	// 不希望在您的广告前后显示的内容排除类别 ID 列表。
	// 若想获取内容排除类别 ID 列表，可使用 /tool/content_exclusion/get/ 并查看返回的 excluded_category_list。
	// 注意: 目前，商品销量暂不支持类别排除。
	CategoryExclusionIDs []string `json:"category_exclusion_ids,omitempty"`
	// VerticalSensitivityID 仅当 brand_safety_type 为 STANDARD_INVENTORY 或 LIMITED_INVENTORY 时生效。
	// 不希望在您的广告前后显示的行业敏感内容类别的 ID。
	// 若想获取行业敏感内容类别 ID 列表，可使用 /tool/content_exclusion/get/ 并查看返回的 vertical_sensitivity_list。
	// 注意: 目前，以下推广目标暂不支持行业敏感内容：访问量、应用推广、线索收集、网站转化量和商品销量。
	VerticalSensitivityID string `json:"vertical_sensitivity_id,omitempty"`
}

// Encode implements PostRequest
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 设置或更新广告账户的品牌安全中心设置 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *BrandSafetySetting `json:"data,omitempty"`
}
