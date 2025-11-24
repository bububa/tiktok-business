package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TargetingInfoRequest 定向标签的信息 API Request
type TargetingInfoRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Scene 所指定的targeting_ids适用的定向类型。
	// 枚举值：
	// GEO ：地理位置定向，包括行政区域定向和邮政编码定向。
	// ISP ：互联网服务提供商定向。
	// 默认值：GEO。
	Scene enum.TargetingType `json:"scene,omitempty"`
	// TargetingIDs 定向标签ID列表。
	// 您可传入地域ID，邮政编码ID，地域ID和邮政编码ID的组合，或互联网服务提供商ID。
	// 最大数量： 1000。
	// 您可通过/tool/targeting/search/返回的geo_id获取地域ID或邮政编码ID，或通过/tool/region/返回的location_id获取地域ID。
	// 您可通过/tool/targeting/list/返回的isp_id获取互联网服务提供商ID。
	TargetingIDs []string `json:"targeting_ids,omitempty"`
	// ObjectiveType 推广目标。
	// 允许的枚举值：REACH，TRAFFIC，VIDEO_VIEWS，LEAD_GENERATION，ENGAGEMENT，APP_PROMOTION，WEB_CONVERSIONS，PRODUCT_SALES。
	// 枚举值的描述请查看推广目标。
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// PromotionType 推广对象类型（优化位置）。
	// 枚举值详见枚举值 - 推广对象类型
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// Placements 投放版位。枚举值请查看枚举值 - 版位。
	Placements []enum.Placement `json:"placements,omitempty"`
	// OperatingSystem 定向受众设备的操作系统。枚举值: ANDROID, IOS
	OperatingSystem enum.OperatingSystem `json:"operating_system,omitempty"`
	// BrandSafetyType 品牌安全类型。 仅当placements 设置为PLACEMENT_TIKTOK时生效。默认值：NO_BRAND_SAFETY。枚举值：
	// NO_BRAND_SAFETY：不使用任何品牌安全解决方案，即使用全部库存。您的广告可能会在某些包含成人主题的内容周围展示。
	// EXPANDED_INVENTORY：使用TikTok的品牌安全解决方案，即使用扩展库存。您的广告将不会出现在不当内容或包含成人主题的内容后面。在下个API版本中EXPANDED_INVENTORY将替代NO_BRAND_SAFETY，并成为默认的品牌安全选项。
	// STANDARD_INVENTORY：使用TikTok的品牌安全解决方案中的标准库存。您的广告将展示在适合大多数品牌的内容周围。
	// LIMITED_INVENTORY：使用TikTok的品牌安全解决方案中的有限库存。您的广告将展示在不包含成人主题的内容周围。
	// THIRD_PARTY：使用第三方品牌安全解决方案。若选择本选项，您需要向brand_safety_partner字段传入第三方合作伙伴名称。
	// 您可以使用/tool/region/接口查询品牌安全设置对应的广告可投放国家或地区。
	// 注意：
	// 出价前第三方品牌安全解决方案和在使用APP_PROMOTION，WEB_CONVERSIONS，TRAFFIC和LEAD_GENERATION推广目标的竞价广告中应用出价前第一方品牌安全解决方案目前均为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 请查看品牌安全，了解出价前品牌安全筛选的支持推广目标。
	BrandSafetyType enum.BrandSafetyType `json:"brand_safety_type,omitempty"`
	// BrandSafetyPartner 品牌安全合作伙伴。仅当brand_safety为THIRD_PARTY时必填， 此时placements 需设置为PLACEMENT_TIKTOK。枚举值: IAS, OPEN_SLATE（该合作伙伴在TikTok广告管理平台的对应名称为DoubleVerify，因为该合作伙伴已被DoubleVerify收购）。
	// 注意：出价前第三方品牌安全解决方案目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	BrandSafetyPartner string `json:"brand_saftey_partner,omitempty"`
}

// Encode implements GetRequest interface
func (r *TargetingInfoRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("targeting_ids", string(util.JSONMarshal(r.TargetingIDs)))
	if r.Scene != "" {
		values.Set("scene", string(r.Scene))
	}
	if r.ObjectiveType != "" {
		values.Set("objective_type", string(r.ObjectiveType))
	}
	if r.PromotionType != "" {
		values.Set("promotion_type", string(r.PromotionType))
	}
	if len(r.Placements) > 0 {
		values.Set("placements", string(util.JSONMarshal(r.Placements)))
	}
	if r.OperatingSystem != "" {
		values.Set("operating_system", string(r.OperatingSystem))
	}
	if r.BrandSafetyType != "" {
		values.Set("brand_safety_type", string(r.BrandSafetyType))
	}
	if r.BrandSafetyPartner != "" {
		values.Set("brand_saftey_partner", r.BrandSafetyPartner)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// TargetingInfoResponse 定向标签的信息 API Response
type TargetingInfoResponse struct {
	model.BaseResponse
	Data *TargetingInfoResult `json:"data,omitempty"`
}

// TargetingInfoResult 定向标签的信息
type TargetingInfoResult struct {
	// TargetingTagList 所指定的定向标签ID所匹配到的定向标签的信息
	TargetingTagList []TargetingTag `json:"targeting_tag_list,omitempty"`
	// ParentTags 所有父层级定向标签的相关信息
	ParentTags []TargetingTag `json:"targeting_tag,omitempty"`
}
