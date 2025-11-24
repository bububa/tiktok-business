package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ContextualTagGetRequest 获取可用内容定向标签 API Request
type ContextualTagGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// ObjectiveType 推广目标。仅支持REACH， VIDEO_VIEWS，RF_REACH
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// RegionCodes 定向的国家或地区代码
	RegionCodes []string `json:"region_codes,omitempty"`
	// BrandSafetyType 品牌安全类型。默认值：STANDARD_INVENTORY。枚举值：
	// EXPANDED_INVENTORY：使用TikTok的品牌安全解决方案，即使用扩展库存。您的广告将不会出现在不当内容或包含成人主题的内容后面。在下个API版本中EXPANDED_INVENTORY将替代NO_BRAND_SAFETY，并成为默认的品牌安全选项。
	// STANDARD_INVENTORY：使用TikTok的品牌安全解决方案中的标准库存。您的广告将展示在适合大多数品牌的内容周围。
	// LIMITED_INVENTORY：使用TikTok的品牌安全解决方案中的有限库存。您的广告将展示在不包含成人主题的内容周围。
	// 您可以使用/tool/region/接口查询品牌安全设置对应的广告可投放国家或地区。
	// 注意：
	// 出价前第三方品牌安全解决方案和在使用APP_PROMOTION，WEB_CONVERSIONS，TRAFFIC和LEAD_GENERATION推广目标的竞价广告中应用出价前第一方品牌安全解决方案目前均为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 请查看品牌安全，了解出价前品牌安全筛选的支持推广目标。
	BrandSafetyType enum.BrandSafetyType `json:"brand_safety_type,omitempty"`
	// RfCampaignType 注意：
	// 仅当推广目标（objective_type）设置为RF_REACH时，可用此字段将推广系列设置为TikTok Pulse推广系列，从而获取PREMIUM标签。
	// 推广目标非RF_REACH时，本字段无需传入。
	// 枚举值： STANDARD （普通覆盖和频次推广系列）， PULSE（TikTok Pulse推广系列）。
	RfCampaignType enum.RfCampaignType `json:"rf_campaign_type,omitempty"`
}

// Encode implements GetRequest interface
func (r *ContextualTagGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("objective_type", string(r.ObjectiveType))
	if len(r.RegionCodes) > 0 {
		values.Set("region_codes", string(util.JSONMarshal(r.RegionCodes)))
	}
	if r.BrandSafetyType != "" {
		values.Set("brand_safety_type", string(r.BrandSafetyType))
	}
	if r.RfCampaignType != "" {
		values.Set("rf_campaign_type", string(r.RfCampaignType))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ContextualTagGetResponse 获取可用内容定向标签 API Response
type ContextualTagGetResponse struct {
	model.BaseResponse
	Data struct {
		// ContextualTagList 内容相关定向标签列表。
		// 注意：内容相关定向目前为白名单功能。若您未申请进入白名单，返回值将为空列表（[]）。
		ContextualTagList []ContextualTag `json:"contextual_tag_list,omitempty"`
	} `json:"data"`
}

// ContextualTag 内容相关定向标签
type ContextualTag struct {
	// ContextualTagID 内容相关定向标签ID。
	// 注意：目前内容相关定向标签仅支持与定向以下国家或地区的受众定位同时使用：
	// 一般标签: US （美国）， CA （美国），GB （英国）， DE （德国）， FR （法国）， IT （意大利）， ES （西班牙）， JP （日本）， KR（韩国），AU （澳大利亚），BR （巴西）， MX（墨西哥）。
	// 优选标签: US （美国）， CA （美国），GB （英国）， DE （德国）， FR （法国）， IT （意大利）， ES （西班牙）， JP （日本）， KR（韩国），AU （澳大利亚），BR （巴西）， MX（墨西哥）。
	ContextualTagID string `json:"contextual_tag_id,omitempty"`
	// Name 内容相关定向标签名称
	Name string `json:"name,omitempty"`
	// Type 内容分组类型。枚举值：GENERAL (一般标签), PREMIUM(优选标签)
	Type enum.ContextualTagType `json:"type,omitempty"`
	// ContentLineupType 内容相关定向标签类型。
	// 枚举值：
	// MAX_PULSE（Max Pulse：您的广告将在 TikTok 上任意主题的热门内容前后展示。）
	// CUSTOM（类别分组：您的广告将在所选类别或季节性活动的热门内容前后展示。）
	// 注意：目前，您可以在 AE（阿联酋）、AU（澳大利亚）、BR（巴西）、CA（加拿大）、DE（德国）、ES（西班牙）、FR（法国）、GB（英国）、IT（意大利）、MX（墨西哥）、SA（沙特阿拉伯）、TR（土耳其） 和 US（美国）使用 Max Pulse，以及在US(美国）、CA（加拿大）、BR（巴西）、AU（澳大利亚）、GB（英国）、FR（法国）、IT（意大利）、ES（西班牙）、DE（德国）使用类别分组。Max Pulse 和类别分组为白名单功能。如果您想在上述市场使用该两种功能，请联系您的 TikTok 销售代表。
	ContentLineupType enum.ContentLineupType `json:"content_lineup_type,omitempty"`
	// Status 内容相关定向标签状态。枚举值: ONLINE, OFFLINE。仅上线状态（ status=ONLINE ）的标签可用于创建广告
	Status enum.OnlineOffline `json:"status,omitempty"`
	// Description 内容相关定向标签描述
	Description string `json:"description,omitempty"`
	// ObjectiveTypes 内容相关定向标签支持的推广目标
	ObjectiveTypes []enum.ObjectiveType `json:"objective_types,omitempty"`
	// RegionCodes 内容相关定向标签支持的国家或地区代码
	RegionCodes []string `json:"region_codes,omitempty"`
}
