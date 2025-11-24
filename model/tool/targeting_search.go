package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TargetingSearchRequest 搜索地域定向标签 API Request
type TargetingSearchRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// ObjectiveType 推广目标。
	// 允许的枚举值：REACH，TRAFFIC，VIDEO_VIEWS，LEAD_GENERATION，ENGAGEMENT，APP_PROMOTION，WEB_CONVERSIONS，PRODUCT_SALES。
	// 枚举值的描述请查看推广目标。
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// PromotionType 推广对象类型（优化位置）。
	// 枚举值详见枚举值 - 推广对象类型
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// Placements 投放版位。枚举值请查看枚举值 - 版位。
	Placements []enum.Placement `json:"placements,omitempty"`
	// SearchType 搜索类型。
	// 枚举值：
	// FUZZY_SEARCH：对单个地域 ID 或邮政编码 ID 进行模糊搜索。
	// 对于此搜索类型，您仅可向 keywords 字段传入一个关键词，最多返回100条搜索结果。
	// BATCH_REGION_SEARCH：对多个地域ID进行批量模糊搜索。
	// 对于此搜索类型，您最多可向 keywords 字段传入1,000个关键词，最多为每个关键词返回5条按相关度排序的搜索结果。
	// BATCH_ZIPCODE_SEARCH : 对多个邮政编码ID进行批量精确搜索。
	// 对于此搜索类型，您最多可向 keywords 字段传入1,000个关键词，最多为每个关键词返回 1 条搜索结果，或无匹配结果时返回 0 条搜索结果。
	// 注意：
	// 巴西、印度尼西亚、泰国、越南的邮政编码定向目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 枚举值BATCH_ZIPCODE_SEARCH和BATCH_REGION_SEARCH仅可用于搜索邮政编码 ID 或地域 ID，而FUZZY_SEARCH可用于搜索单个邮政编码 ID 或地域 ID。
	SearchType enum.TargetingSearchType `json:"search_type,omitempty"`
	// Keywords 用于搜索定向标签的关键词列表。
	// 长度限制：
	// search_type 为 FUZZY_SEARCH时为1。
	// search_type为BATCH_REGION_SEARCH或 BATCH_ZIPCODE_SEARCH时为1,000。
	// 若您将search_type设置为BATCH_ZIPCODE_SEARCH，您需传入邮政编码作为关键词，因为此搜索类型仅支持精确搜索。
	// 对于美国邮政编码定向，请传入 5 位美国邮政编码作为关键词。示例：["10001","10002"]。
	// 对于加拿大邮政编码定向，请传入加拿大邮政编码的前 3 个字符（前分拣区）作为关键词。示例：["A0A","A0B"]。
	// 对于巴西邮政编码定向，请传入巴西邮政编码的前 5 位数字作为关键词。示例：["81480","18072"]。
	// 对于印度尼西亚邮政编码定向，请传入 5 位的印度尼西亚邮政编码作为关键词。示例：["20371","12120"]。
	// 对于泰国邮政编码定向，请传入 5 位的泰国邮政编码作为关键词。示例：["30000","40000"]。
	// 对于越南邮政编码定向，请传入越南邮政编码的前 3 位数字作为关键词。示例：["718","719"]。
	Keywords []string `json:"keywords,omitempty"`
	// GeoTypes 用于筛选结果的地域类型。
	// 枚举值：COUNTRY（国家或地区），PROVINCE（省），CITY（城市/郡/县），DISTRICT（区/城市），DMA（指定市场区域），ZIP_CODE（邮政编码对应地域）。
	// 当 search_type设置为FUZZY_SEARCH时，本字段允许值为ZIP_CODE，COUNTRY，PROVINCE，CITY，DISTRICT和DMA。
	// 当 search_type设置为BATCH_ZIPCODE_SEARCH时，本字段允许值为ZIP_CODE。
	// 当search_type设置为BATCH_REGION_SEARCH时，本字段允许值为COUNTRY，PROVINCE，CITY，DISTRICT和DMA。
	GeoTypes []enum.GeoType `json:"geo_types,omitempty"`
	// RegionCodes 用于筛选结果的定向国家或地区代码。
	// 当search_type设置为 BATCH_ZIPCODE_SEARCH 或BATCH_REGION_SEARCH时本字段必填，且只允许设置为US（美国）、 CA（加拿大）、BR（巴西）、ID（印度尼西亚）、TH（泰国）或 VN （越南）。
	// 当search_type设置为 FUZZY_SEARCH时，本字段可传，允许的国家或地区代码枚举值参见附录-地区代码 。
	RegionCodes []string `json:"region_codes,omitempty"`
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
func (r *TargetingSearchRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("objective_type", string(r.ObjectiveType))
	if r.PromotionType != "" {
		values.Set("promotion_type", string(r.PromotionType))
	}
	values.Set("placements", string(util.JSONMarshal(r.Placements)))
	values.Set("search_type", string(r.SearchType))
	values.Set("keywords", string(util.JSONMarshal(r.Keywords)))
	if len(r.GeoTypes) > 0 {
		values.Set("geo_types", string(util.JSONMarshal(r.GeoTypes)))
	}
	if len(r.RegionCodes) > 0 {
		values.Set("region_codes", string(util.JSONMarshal(r.RegionCodes)))
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

// TargetingSearchResponse 搜索地域定向标签 API Response
type TargetingSearchResponse struct {
	model.BaseResponse
	Data *TargetingSearchResult `json:"data,omitempty"`
}

// TargetingSearchResult 定向标签的信息
type TargetingSearchResult struct {
	// TargetingTagList 所指定的关键词所匹配到的定向标签的信息。
	// 注意：对于同一个关键词（keyword）可能返回多个匹配到的定向标签，对应不同名称（name）。
	TargetingTagList []TargetingTag `json:"targeting_tag_list,omitempty"`
	// ParentTags 所有父层级定向标签的相关信息
	ParentTags []TargetingTag `json:"targeting_tag,omitempty"`
}
