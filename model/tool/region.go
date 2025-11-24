package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// RegionRequest 基于具体设置获取可投放地域 API Request
type RegionRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Placements 投放版位。枚举值请查看枚举值 - 版位。
	Placements []enum.Placement `json:"placements,omitempty"`
	// ObjectiveType 推广目标。枚举值及描述请查看推广目标
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// LevelRange 想要返回的地域层级。 枚举值:
	// ALL: 返回所有层级的地域。
	// TO_COUNTRY：只返回国家或地区级地域。
	// TO_PROVINCE：返回国家和省级地域。指定市场区域和都会区属于省级地域。
	// TO_CITY: 返回国家、省、市级地域。
	// TO_DISTRICT: 返回国家、省、市、区级地域。
	LevelRange enum.RegionLevelRange `json:"level_range,omitempty"`
	// Language 返回的地域名称（name）的语言。
	// 默认值： en（英语）。
	// 若想获取可选值，请查看 language 可选值。
	// 注意：
	// 若不支持指定的语言，则将忽略指定的语言，并使用默认值 en 。
	// 若指定了支持的非英语的语言，则以该语言为本土语言的国家或地区之外的地域名称可能未进行翻译。例如，若您将 language 设置为ja（日语），则日本之外的地域的名称可能不会被翻译。
	Language string `json:"language,omitempty"`
	// ShoppingAdsType 仅当 objective_type 为 PRODUCT_SALES 时生效。
	// 购物广告类型。
	// 枚举值：
	// VIDEO：视频购物广告。
	// LIVE：直播购物广告。
	// PRODUCT_SHOPPING_ADS：商品购物广告。
	// 若想了解不同类型的购物广告以及创建视频购物广告、直播购物广告和商品购物广告的步骤，请查看创建购物广告。
	ShoppingAdsType enum.ShoppingAdsType `json:"shopping_ads_type,omitempty"`
	// PromotionType 推广对象类型（优化位置）。
	// 枚举值详见枚举值 - 推广对象类型
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// PromotionTargetType 仅当objective_type 设置为 LEAD_GENERATION时有效。推广目标线索收集的专属推广对象类型。
	// 枚举值：
	// INSTANT_PAGE：即时表单（线索表单）。创建可快速加载的应用内 TikTok 即时表单来收集更多线索。
	// EXTERNAL_WEBSITE：第三方网站表单。使用包含第三方网站表单的落地页，或使用会重定向至包含第三方网站表单网站的TikTok即时体验页面来收集更多线索。
	PromotionTargetType enum.PromotionTargetType `json:"promotion_target_type,omitempty"`
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
	// RfCampaignType 仅当推广目标（objective_type）设置为RF_REACH时，可用此字段将推广系列设置为TikTok Pulse推广系列，从而获取PREMIUM标签。
	// 枚举值： STANDARD （普通覆盖和频次推广系列）， PULSE（TikTok Pulse推广系列）。
	// 若指定为PULSE，则仅返回一个国家或地区ID（location_id）。
	// 注意：
	// 推广目标非RF_REACH时，本字段无需传入。
	// 使用/adgroup/rf/create/ 创建TikTok Pulse广告组时，请求中需将本接口返回的location_id值传给location_ids。
	RfCampaignType enum.RfCampaignType `json:"rf_campaign_type,omitempty"`
}

// Encode implements GetRequest interface
func (r *RegionRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("placements", string(util.JSONMarshal(r.Placements)))
	values.Set("objective_type", string(r.ObjectiveType))
	if r.LevelRange != "" {
		values.Set("level_range", string(r.LevelRange))
	}
	if r.Language != "" {
		values.Set("language", r.Language)
	}
	if r.ShoppingAdsType != "" {
		values.Set("shopping_ads_type", string(r.ShoppingAdsType))
	}
	if r.PromotionType != "" {
		values.Set("promotion_type", string(r.PromotionType))
	}
	if r.PromotionTargetType != "" {
		values.Set("promotion_target_type", string(r.PromotionTargetType))
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
	if r.RfCampaignType != "" {
		values.Set("rf_campaign_type", string(r.RfCampaignType))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// RegionResponse 基于具体设置获取可投放地域 API Response
type RegionResponse struct {
	model.BaseResponse
	Data *RegionResult `json:"data,omitempty"`
}

type RegionResult struct {
	// RegionList 国家或地区代码列表
	RegionList []string `json:"region_list,omitempty"`
	// RegionInfo 可投放地域信息
	RegionInfo []Region `json:"region_info,omitempty"`
}

// Region 可投放地域信息
type Region struct {
	// LocationID 地域ID
	LocationID string `json:"location_id,omitempty"`
	// Name 地域名称
	Name string `json:"name,omitempty"`
	// ParentID 上级地域ID。如果没有上级地域，返回0
	ParentID string `json:"parent_id,omitempty"`
	// RegionCode 国家或地区代码。例如： "DE"
	RegionCode string `json:"region_code,omitempty"`
	// CountryCode 地域所属的国家代码
	CountryCode string `json:"country_code,omitempty"`
	// NextLevelIDs 下级地域ID列表
	NextLevelIDs []string `json:"next_level_ids,omitempty"`
	// AreaType 地域类型。枚举值: ADMIN, METROPOLITAN_OR_DMA
	AreaType enum.AreaType `json:"area_type,omitempty"`
	// Level 地域层级。枚举值: COUNTRY（国家或地区级）, PROVINCE, CITY（市或郡/县级）, DISTRICT（区或市级）。 指定市场区域以及都会区属于省级地域
	Level enum.RegionLevel `json:"level,omitempty"`
	// SupportBelow18 该地域是否支持向18岁以下人群投放广告
	SupportBelow18 bool `json:"support_below_18,omitempty"`
}
