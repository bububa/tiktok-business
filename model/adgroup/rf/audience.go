package rf

import "github.com/bububa/tiktok-business/enum"

// Audience 受众定向规则
type Audience struct {
	// AudienceIDs 想要定向的受众ID列表。若要将受众在覆盖和频次广告中定向或排除，需要在创建或更新受众时，将audience_sub_type设置为REACH_FREQUENCY
	AudienceIDs []string `json:"audience_ids,omitempty"`
	// ExcludedAudienceIDs 要排除的受众ID列表。若要将受众在覆盖和频次广告中定向或排除，需要在创建或更新受众时，将audience_sub_type设置为REACH_FREQUENCY
	ExcludedAudienceIDs []string `json:"excluded_audience_ids,omitempty"`
	// AgeGroups 受众年龄区间。若定向地区包括以色列和巴西，则此字段不可设置为AGE_13_17。枚举值可见枚举值-受众年龄区间
	AgeGroups []enum.Age `json:"age_groups,omitempty"`
	// Gender 受众性别。 枚举值详见枚举值 - 性别
	Gender []enum.AudienceGender `json:"gender,omitempty"`
	// Languages 受众语言。枚举值详见枚举值-受众语言。
	Languages []string `json:"languages,omitempty"`
	// LocationIDs 定向地域ID。所定向地域可为国家或地区层级，省、指定市场区域或都会区层级，或城市层级地域。
	// 只支持定向到单个国家或地区。如果传入了多个ID，则这些ID必须属于同一国家或地区。
	// 支持投放覆盖和频次广告的国家和地区请参见覆盖和频次广告开头部分。地域ID列表可参见附录-地区定向。
	LocationIDs []string `json:"location_ids,omitempty"`
	// OperatingSystems 操作系统。枚举值详见枚举值-受众操作系统
	OperatingSystems []enum.OperatingSystem `json:"operating_systems,omitempty"`
	// NetworkTypes 网络类型。不传即不限制网络类型。枚举值详见枚举值-网络类型
	NetworkTypes []enum.NetworkType `json:"network_types,omitempty"`
	// DeviceModelIDs 想要定向的设备机型ID列表。目前只支持定向到品牌层级。传入本字段时，设备价格字段device_price_ranges不可同时传入。所有设备机型ID的枚举值可通过 设备机型接口获取。
	// 注意: 创建广告时需确保对应设备处于活跃状态（ /tool/device_model/ 接口返回中is_active 为 True）
	DeviceModelIDs []string `json:"device_model_ids,omitempty"`
	// DevicePriceRanges 设备价格区间（10000代表1000+）。该数字必须是50的倍数。
	// 重要提示: 最终的实际上限将在您设定的上限基础上增加50，您可以在TikTok广告管理平台的广告组设置中看到实际上限。例如，如果您设置的价格区间是[0, 250]，实际区间则为[0, 300]。
	DevicePriceRanges []int `json:"device_price_ranges,omitempty"`
	// CarrierIDs 运营商ID列表。您可使用 /tool/carrier/ 接口获取此字段的枚举值
	CarrierIDs []string `json:"carrier_ids,omitempty"`
	// InterestCategoryIDs 兴趣分类。不传即不限制兴趣分类，若想定向所有受众，请不要传入此字段。您可以使用/tool/target_recommend_tags/接口获取推荐兴趣分类，推荐结果将基于您的定向地区及所在行业。您还可以使用/tool/interest_category/ 接口获取兴趣分类枚举值
	InterestCategoryIDs []string `json:"interest_category_ids,omitempty"`
	// BrandSafetyType 覆盖和频次广告的品牌安全类型。默认值：NO_BRAND_SAFETY。枚举值：
	// NO_BRAND_SAFETY：不使用任何品牌安全解决方案，即使用全部库存。您的广告可能会在某些包含成人主题的内容周围展示。
	// EXPANDED_INVENTORY：使用TikTok的品牌安全解决方案，即使用扩展库存。您的广告将不会出现在不当内容或包含成人主题的内容后面。在下个API版本中EXPANDED_INVENTORY将替代NO_BRAND_SAFETY，并成为默认的品牌安全选项。
	// STANDARD_INVENTORY：使用TikTok的品牌安全解决方案中的标准库存。您的广告将展示在适合大多数品牌的内容周围。
	// LIMITED_INVENTORY：使用TikTok的品牌安全解决方案中的有限库存。您的广告将展示在不包含成人主题的内容周围。
	// THIRD_PARTY：使用第三方品牌安全解决方案。
	// 您可以使用/tool/region/接口查询品牌安全设置对应的广告可投放国家或地区。
	// 注意：
	// 出价前第三方品牌安全解决方案目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	BrandSafetyType enum.BrandSafetyType `json:"brand_safe_type,omitempty"`
	// ContextualTagIDs 内容相关定向标签ID列表。您可使用 /tool/contextual_tag/get/获取可用的内容相关定向标签。
	// 注意：
	// 本字段目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 仅支持推广系列的推广目标为RF_REACH 。
	// 不支持同时将brand_safety_type设置为THIRD_PARTY。
	// 不支持同时将feed_type设置为 TOP_FEED。
	// 若您使用/campaign/create/创建推广系列时，将rf_campaign_type 设置为PULSE，则需在本字段（contextual_tag_ids）传入 PREMIUM类型的内容相关定向标签，创建的广告组的CPM（cpm）将是固定的。
	ContextualTagIDs []string `json:"contextual_tag_ids,omitempty"`
}
