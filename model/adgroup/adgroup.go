package adgroup

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Adgroup 广告组
type Adgroup struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName 该广告组所属的推广系列的名称。
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignSystemOrigin 仅为内容加热推广系列返回本字段。
	// 非内容加热推广系列不返回本字段。
	//
	// 推广系列来源。
	//
	// 枚举值：
	// PROMOTE：该推广系列为通过 TikTok 移动应用创建的内容加热推广系列。
	//
	// 仅为内容加热推广系列（campaign_system_origin 为 PROMOTE）返回以下信息：
	// advertiser_id
	// campaign_id
	// create_time
	// campaign_name
	// operation_status
	// secondary_status
	CampaignSystemOrigin enum.CampaignSystemOrigin `json:"campaign_system_origin,omitempty"`
	// IsSmartPerformanceCampaign 是否为自动化类型的推广系列。
	// 支持的值：
	// true：推广系列为 Smart+ 推广系列或智能效果网站推广系列。
	// false：推广系列为普通类型的推广系列。
	// 注意：若is_smart_performance_campaign 为 true 且 objective_type 为 WEB_CONVERSIONS，您可根据/campaign/spc/get/返回的spc_type查看该网站推广系列为 Smart+ 网站推广系列还是智能效果网站推广系列。
	IsSmartPerformanceCampaign bool `json:"is_smart_performance_campaign,omitempty"`
	// AdgroupID 广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName 广告组名称
	AdgroupName string `json:"adgroup_name,omitempty"`
	// CreateTime 广告组创建时间。格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2023-01-01 00:00:01".
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ModifyTime 广告组修改时间。格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2023-01-01 00:00:01".
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
	// PromotionType 推广对象类型（优化位置）。若对应的推广系列推广目标不是REACH，VIDEO_VIEWS或ENGAGEMENT，则此字段必填。
	// 枚举值详见 枚举值-推广对象类型。
	// 注意: 如果在推广系列层级的objective_type是 ENGAGEMENT，则该字段只能使用 EXTERNAL_OR_DISPLAY。
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// ShoppingAdsType 当对应的推广系列 objective_type 为 PRODUCT_SALES 时返回。
	// 购物广告类型。
	//
	// 枚举值：
	// VIDEO：视频购物广告。
	// LIVE：直播购物广告。
	// PRODUCT_SHOPPING_ADS：商品购物广告。
	// CATALOG_LISTING_ADS：产品目录广告。
	// UNSET：未设置。
	ShoppingAdsType enum.ShoppingAdsType `json:"shopping_ads_type,omitempty"`
	// IdentityID 当 shopping_ads_type 为 VIDEO 且 product_source 为 SHOWCASE 或 shopping_ads_type 为 LIVE 时返回。
	// 认证身份 ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 当 shopping_ads_type 为 VIDEO 且 product_source 为 SHOWCASE 或 shopping_ads_type 为 LIVE 时返回。
	// 认证身份类型。枚举值: AUTH_CODE (授权码认证)， TT_USER (TikTok商务用户)， BC_AUTH_TT（已添加到商务中心的TikTok用户）。参阅认证身份 获取详细信息。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 当identity_type 为 BC_AUTH_TT时返回。
	// 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// ProductSource 仅视频购物广告和商品购物广告返回本字段。
	// 商品来源，即您要推广的商品的来源。
	// 枚举值：UNSET ,CATALOG(商品库)，STORE (TikTok Shop 店铺) ，SHOWCASE（TikTok 橱窗）
	ProductSource enum.ProductSource `json:"product_source,omitempty"`
	// CatalogID 以下任意场景下返回：
	// product_source 设置为 CATALOG 或 STORE。
	// shopping_ads_type 设置为 LIVE 且请求中传入了本字段。
	//
	// 商品库 ID。
	//
	// 注意：自 2024 年 6 月 30 日起，将不再为新创建的 product_source 为 STORE 的购物广告返回 catalog_id，传入的 catalog_id 字段将被忽略。存量的 product_source 为 STORE 的购物广告若未被更新，则不受影响。若您更新存量的这类广告组，且广告组已绑定 catalog_id ，则绑定的 catalog_id 将清空且该 catalog_id 字段将不再返回。
	CatalogID string `json:"catalog_id,omitempty"`
	// CatalogAuthorizedBcID 以下任意场景下返回：
	// shopping_ads_type 设置为 VIDEO 且 product_source 设置为CATALOG 或 STORE。
	// shopping_ads_type 设置为 LIVE 且请求中传入了本字段。
	// 如果要操作的是已经资产化的商品库，本字段会返回相应的商务中心 ID。
	//
	// 注意：自 2024 年 6 月 30 日起，将不再为新创建的 product_source 为 STORE 的购物广告返回 catalog_authorized_bc_id，传入的 catalog_authorized_bc_id 字段将被忽略。存量的 product_source 为 STORE 的购物广告若未被更新，则不受影响。若您更新存量的这类广告组，且广告组已绑定 catalog_authorized_bc_id ，则绑定的 catalog_authorized_bc_id 将清空且该 catalog_authorized_bc_id 字段将不再返回。
	CatalogAuthorizedBcID string `json:"catalog_authorized_bc_id,omitempty"`
	// StoreID 以下任意场景下返回：
	// shopping_ads_type 设置为 VIDEO 且 product_source 设置为 STORE。
	// shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS 且 product_source 设置为 STORE。
	// shopping_ads_type 设置为LIVE 且请求中传入了本字段。
	//
	// TikTok Shop 店铺 ID。
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID 请求中传入 store_id 时返回。
	// 有权管理该商店 (store_id）的商务中心的 ID。
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// PromotionTargetType 推广目标线索收集的专属推广对象类型。
	// 枚举值：
	// INSTANT_PAGE：即时表单（线索表单）。创建可快速加载的应用内 TikTok 即时表单来收集更多线索。
	// EXTERNAL_WEBSITE：第三方网站表单。使用包含第三方网站表单的落地页，或使用会重定向至包含第三方网站表单网站的TikTok即时体验页面来收集更多线索。
	// UNSET。
	PromotionTargetType enum.PromotionTargetType `json:"promotion_target_type,omitempty"`
	// MessageAppType 即时通讯广告组中使用的即时通讯应用或自定义 URL。
	// 枚举值：
	// MESSENGER：Messenger。
	// WHATSAPP：WhatsApp。
	// ZALO: Zalo。
	// LINE: Line。
	// IM_URL：即时通讯 URL。
	MesssageAppType enum.MessageAppType `json:"message_app_type,omitempty"`
	// MessageAppAccountID 即时通讯应用账号 ID。
	// messaging_app_type 为 MESSENGER 时，本字段代表 Facebook 公共主页编号。
	// messaging_app_type 为 LINE 时，本字段代表 LINE Business ID。
	// messaging_app_type 为 WHATSAPP 时，本字段代表根据指定的phone_region_code、phone_region_calling_code 和 phone_number 自动填充的 WhatsApp 电话号码。
	MessageAppAccountID string `json:"message_app_account_id,omitempty"`
	// PhoneRegionCode WhatsApp 或 Zalo 电话号码的地区代码。
	// 示例：US。
	PhoneRegionCode string `json:"phone_region_code,omitempty"`
	// PhoneRegionCallingCode WhatsApp 或 Zalo 电话号码的区号。
	// 示例：+1。
	PhoneRegionCallingCode string `json:"phone_region_calling_code,omitempty"`
	// PhoneNumber WhatsApp 或 Zalo 电话号码。
	PhoneNumber string `json:"phone_number,omitempty"`
	// PromotionWebsiteType 广告组中的 TikTok 即时体验页面类型。
	// 枚举值：
	// UNSET：不使用TikTok 即时体验页面。
	// TIKTOK_NATIVE_PAGE：使用 TikTok 即时体验页面。
	PromotionWebsiteType enum.PromotionWebsiteType `json:"promotion_website_type,omitempty"`
	// AppID 推广的App的ID。您可通过/app/list/获取app_id
	AppID string `json:"app_id,omitempty"`
	// AppType 推广的App的的类型。枚举值：APP_ANDROID（Android），APP_IOS（iOS）
	AppType enum.AppType `json:"app_type,omitempty"`
	// AppDownloadURL App下载链接
	AppDownloadURL string `json:"app_download_url,omitempty"`
	// PixelID Pixel ID，仅对落地页生效
	PixelID string `json:"pixel_id,omitempty"`
	// OptimizationEvent 广告组转化事件。参阅转化事件 ，获取详细信息
	OptimizationEvent enum.OptimizationEvent `json:"optimization_event,omitempty"`
	// DeepFunnelOptimizationStatus 深层漏斗优化的状态。
	// 深层漏斗优化同时优化浅层和深层漏斗事件。您可以在主要优化事件之外再选择一个次要事件，这有助于提升推广 系列的效果。
	// 枚举值：
	// ON ：启用深层漏斗优化。
	// OFF：不启用深层漏斗优化。
	DeepFunnelOptimizationStatus enum.OnOff `json:"deep_funnel_optimization_status,omitempty"`
	// DeepFunnelEventSource deep_funnel_optimization_status 为 ON 时返回本字段。
	// 事件源类型。
	// 枚举值：
	// PIXEL：Pixel。
	// OFFLINE：线下事件组。
	// CRM：CRM 事件组。
	DeepFunnelEventSource enum.DeepFunnelEventSource `json:"deep_funnel_event_source,omitempty"`
	// DeepFunnelEventSourceID deep_funnel_optimization_status 为 ON 时返回本字段。
	//  事件源 ID。
	// deep_funnel_event_source 为 PIXEL 时，本字段代表 Pixel ID。
	// deep_funnel_event_source 为 OFFLINE 时，本字段代表线下事件组 ID。
	// deep_funnel_event_source为CRM时，本字段代表 CRM 事件组 ID。
	DeepFunnelEventSourceID string `json:"deep_funnel_event_source_id,omitempty"`
	// DeepFunnelOptimizationEvent deep_funnel_optimization_status 为 ON 时返回本字段。
	// 深层漏斗优化事件。
	// 示例： SHOPPING。
	DeepFunnelOptimizationEvent enum.OptimizationEvent `json:"deep_funnel_optimization_event,omitempty"`
	// PlacementType 版位类型。
	// 枚举值：PLACEMENT_TYPE_AUTOMATIC（自动版位），PLACEMENT_TYPE_NORMAL （编辑版位）。
	PlacementType enum.PlacementType `json:"placement_type,omitempty"`
	// Placements 投放版位。获取枚举值，详见枚举值-版位。
	// 注意：
	// 当placement_type为PLACEMENT_TYPE_AUTOMATIC（自动版位）时，本字段返回指定的推广目标下实际支持的版位。例如，若实际仅支持 TikTok版位，本字段返回的值将为PLACEMENT_TIKTOK 。
	// 若指定的推广目标支持 PLACEMENT_GLOBAL_APP_BUNDLE 版位，但您未申请全球化应用组合版位的白名单，则本字段的值将不会包含PLACEMENT_GLOBAL_APP_BUNDLE。
	// 当 placement_type 为 PLACEMENT_TYPE_NORMAL（手动版位）时，本字段通常将返回您在请求中通过placements字段指定的版位。
	// 唯一的例外是您在请求中通过 placements 字段指定了 PLACEMENT_GLOBAL_APP_BUNDLE 版位 ，但是您未申请 PLACEMENT_GLOBAL_APP_BUNDLE 的白名单，则本字段的返回值中将筛除PLACEMENT_GLOBAL_APP_BUNDLE 。例如，若您在请求中将 placements 设置为["PLACEMENT_TIKTOK", "PLACEMENT_GLOBAL_APP_BUNDLE"]，但是未申请全球化应用组合版位的白名单，本字段的返回值将为["PLACEMENT_TIKTOK"]。
	Placements []enum.Placement `json:"placements,omitempty"`
	// TiktokSubplacements TikTok 子版位，用于选择广告出现的具体位置。
	// 枚举值：
	// IN_FEED：信息流。广告将投放在“推荐”动态中，可能投放在“个人资料页”和“关注”推荐内容中。
	// SEARCH_FEED：搜索推荐内容。
	// TIKTOK_LITE：TikTok Lite，占用内容更小、视频加载速度更快的精简版 TikTok 应用。TikTok Lite 子版位仅支持定向日本或韩国时使用。
	// 注意：若创建广告组时未指定 tiktok_subplacements，本字段将为空列表（[]）。
	TiktokSubplacements []enum.TiktokSubplacement `json:"tiktok_subplacements,omitempty"`
	// SearchResultEnabled 是否在搜索广告中展示您的广告，即当用户在 TikTok上搜索您的业务相关信息时，系统是否向他们展示您的广告。
	SearchResultEnabled bool `json:"search_result_enabled,omitempty"`
	// CommentDisabled 是否允许用户在TikTok上评论您的广告
	CommentDisabled bool `json:"comment_disabled,omitempty"`
	// VideoDownloadDisabled 用户是否可以在 TikTok 上下载您的广告视频
	VideoDownloadDisabled bool `json:"video_download_disabled,omitempty"`
	// ShareDisabled 本广告组中的广告是否禁止分享到第三方平台
	ShareDisabled bool `json:"share_disabled,omitempty"`
	// BlockedPangleAppIDs Pangle 媒体屏蔽列表 ID
	BlockedPangleAppIDs []string `json:"blocked_pangle_app_ids,omitempty"`
	// AudienceType 应用重定向受众类型。枚举值可见枚举值 - 应用重定向受众类型
	AudienceType enum.AudienceType `json:"audience_type,omitempty"`
	// AudienceRule 受众规则
	AudienceRule *AudienceRule `json:"audience_rule,omitempty"`
	// AutoTargetingEnabled 是否开启自动定向。
	// 注意：自 2024 年 6 月起，您将无法为广告组开启自动定向或定向扩量。为保证流畅的 API 体验，我们建议您使用智能定向。
	AutoTargetingEnabled bool `json:"auto_targeting_enabled,omitempty"`
	// ShoppingAdsRetargetingType shopping_ads_type 设置为 VIDEO 且 product_source 设置为 CATALOG 时返回。
	// 购物广告受众再营销类型。
	//
	// 枚举值：
	// LAB1 : 对浏览过商品或者将商品加过购物车，但未购买过商品的受众进行再营销。
	// LAB2：对将商品加过购物车,但未购买过商品的受众进行再营销。
	// LAB3 : 使用自定义条件组合对受众进行再营销。
	// OFF : 不进行受众再营销。
	ShoppingAdsRetargetingType enum.ShoppingAdsRetargetingType `json:"shopping_ads_retargeting_type,omitempty"`
	// ShoppingAdsRetargetingActionsDays 所指定的受众行为的有效时间范围。在该时间范围内完成所指定行为的受众将成为购物广告受众再营销的对象。
	// 取值范围：1，2，3，7, 14，30，60，90，180。
	ShoppingAdsRetargetingActionsDays int `json:"shopping_ads_retargeting_actions_days,omitempty"`
	// IncludeCustomActions 作为包含条件筛选购物广告受众再营销对象的自定义行为
	IncludeCustomActions []CustomAction `json:"include_custom_actions,omitempty"`
	// ExcludeCustomActions 作为排除条件筛选购物广告受众再营销对象的自定义行为。
	ExcludeCustomActions []CustomAction `json:"exclude_custom_actions,omitempty"`
	// ShoppingAdsRetargetingCustomAudienceRelation 由 shopping_ads_retargeting_type 字段指定的视频购物广告重定向受众与由 audience_ids 指定的自定义受众间的逻辑关系。
	// 枚举值：
	// OR：取视频购物广告重定向受众和自定义受众的并集构建定向受众。定向受众中包含属于视频购物广告重定向受众或自定义受众的人群。
	// AND：取视频购物广告重定向受众和自定义受众的交集构建定向受众。定向受众中仅包含同时属于视频购物广告重定向受众和自定义受众的人群。
	ShoppingAdsRetargetingCustomAudienceRelation enum.FilterSetOperator `json:"shopping_ads_retargeting_custom_audience_relation,omitempty"`
	// LocationIDs 定向地域 ID。使用tool/region 获取定向地域ID
	LocationIDs []string `json:"location_ids,omitempty"`
	// ZipcodeIDs 定向地域的邮政编码 ID。
	// 注意：
	// 邮政编码定向目前仅支持定向美国、加拿大、巴西、印度尼西亚、泰国、越南的地域。
	// 欲获取邮政编码 ID 的信息，您仅可使用 /tool/targeting/info/。
	ZipcodeIDs []string `json:"zipcode_ids,omitempty"`
	// Langauges 受众语言。
	// 枚举值：详见枚举值-受众语言
	Languages []string `json:"languages,omitempty"`
	// Gender 定向受众性别
	// 枚举值: GENDER_FEMALE,GENDER_MALE,GENDER_UNLIMITED。
	Gender enum.GENDER `json:"gender,omitempty"`
	// AgeGroups 定向受众年龄
	// 枚举值：详见枚举值-受众年龄区间
	AgeGroups []enum.Age `json:"age_groups,omitempty"`
	// ExcludeAgeUnderEighteen 是否排除十八岁以下受众
	ExcludeAgeUnderEighteen bool `json:"exclude_age_under_eighteen,omitempty"`
	// SpendingPower 消费能力。枚举值：ALL, HIGH。如果设为HIGH，您就可以定向和普通用户相比 TikTok 广告投入支出通常更多的高消费能力用户。
	SpendingPower enum.SpendingPower `json:"spending_power,omitempty"`
	// HouseholdIncome 您想定向的收入群体。枚举值：TOP5(家庭收入群体的前5%)， TOP10(家庭收入群体的前10%)，TOP10_25(家庭收入群体的前10-25%)， TOP25_50(家庭收入群体的前25-50%)。
	HouseholdIncome enum.HouseholdIncome `json:"household_income,omitempty"`
	// AudienceIDs 受众 ID 列表。
	AudienceIDs []string `json:"audience_ids,omitempty"`
	// SmartAudienceEnabled 是否已启用智能受众。
	// 枚举值： true， false。
	// 若想了解智能受众详情及如何启用智能受众，请查看智能定向。
	SmartAudienceEnabled bool `json:"smart_audience_enabled,omitempty"`
	// ExcludeAudienceIDs 排除受众 ID 列表
	ExcludeAudienceIDs []string `json:"exclude_audience_ids,omitempty"`
	// InterestCategoryIDs 用于定向受众的一般兴趣分类 ID 列表
	InterestCategoryIDs []string `json:"interest_category_ids,omitempty"`
	// InterestKeywordIDs 用于定向受众的其他兴趣分类 ID 列表
	InterestKeywordIDs []string `json:"interest_keyword_ids,omitempty"`
	// PurchaseIntentionKeywordIDs 购买意向分类 ID 列表，用于定向近期行为表明其可能会购买某内容类别相关产品的受众。
	// 若想搜索或列举购买意向分类 ID，可使用/targeting/search/。
	// 注意：
	// 不支持同时传入 purchase_intention_keyword_ids 和interest_keyword_ids ，否则会出现关键词冲突。
	// purchase_intention_keyword_ids 只支持竞价广告，且对应的 objective_type（推广目标）应为 APP_PROMOTION，WEB_CONVERSIONS ，LEAD_GENERATION，TRAFFIC 或 PRODUCT_SALES（对应的 product_source 需设置为CATALOG），同时版位需包含 TikTok（PLACEMENT_TIKTOK）或 Pangle （PLACEMENT_PANGLE）。
	PurchaseIntentionKeywordIDs []string `json:"purchase_intention_keyword_ids,omitempty"`
	// Actions 用于定向的行为分类对象数组
	Actions []TargetingAction `json:"actions,omitempty"`
	// SmartInterestBehaviorEnabled 是否已启用智能型兴趣和行为。
	// 枚举值： true， false。
	// 若想了解智能型兴趣和行为详情及如何启用智能型兴趣和行为，请查看智能定向。
	SmartInterestBehaviorEnabled bool `json:"smart_interest_behavior_enabled,omitempty"`
	// IncludedPangleAudiencePackageIDs 包含 Pangle 流量包 ID。您可以使用接口/pangle_audience_package/get/获取受众ID (package_id) 。bind_type设置为INCLUDE。请不要同时传入该字段和excluded_pangle_audience_package_ids
	IncludedPangleAudiencePackageIDs []string `json:"include_pangle_audience_package_ids,omitempty"`
	// ExcludePangleAudiencePackageIDs 排除 Pangle 流量包 ID。您可以使用接口/pangle_audience_package/get/获取受众ID (package_id) 。bind_type设置为EXCLUDE。请不要同时传入该字段和included_pangle_audience_package_ids
	ExcludePangleAudiencePackageIDs []string `json:"exclude_pangle_audience_package_ids,omitempty"`
	// OperatingSystems 受众操作系统，枚举值：ANDROID, IOS
	OperatingSystems []enum.OperatingSystem `json:"operating_systems,omitempty"`
	// MinAndroidVersion 受众最低Android版本。枚举值：详见枚举值-最低Android版本。
	MinAndroidVersion string `json:"min_android_version,omitempty"`
	// Ios14Targeting 您想定向的 iOS 设备版本。
	// 枚举值：
	// UNSET：运行 iOS 14.4 及更早版本的设备。
	// IOS14_MINUS：不受 iOS 14 新的 ATT 框架影响的 iOS 设备（运行 iOS 14.0 及更早版本的设备）。
	// IOS14_PLUS：运行iOS 14.5 及之后版本、已执行 ATT 框架的 iOS 设备。此值仅支持在专属推广系列中使用。
	// ALL：运行 iOS 14.5 及之后版本、已执行 ATT 框架的 iOS 设备。此值仅支持在 iOS 应用再营销广告或 iOS 重定向类型的视频购物广告（商品来源为商品库且优化位置为应用）中使用。
	Ios14Targeting enum.Ios14Targeting `json:"ios14_targeting,omitempty"`
	// MinIosVersion 受众最低 iOS 版本。枚举值：详见枚举值-最低 iOS 版本
	MinIosVersion string `json:"min_ios_version,omitempty"`
	// Ios14QuotaType 本推广系列是否占用iOS14专属推广系列配额。枚举值：OCCUPIED（占用）， UNOCCUPIED（不占用）。对于非覆盖和频次推广系列，如果ios14_targeting设为 IOS14_PLUS， 本字段会自动设为OCCUPIED
	Ios14QuotaType enum.Ios14QuotaType `json:"ios14_quota_type,omitempty"`
	// DeviceModelIDs 设备机型ID列表。关于设备机型更多信息，可参考获取设备型号列表
	DeviceModelIDs []string `json:"device_model_ids,omitempty"`
	// NetworkTypes 网络类型。
	// 详见枚举值-网络类型。
	NetworkTypes []enum.NetworkType `json:"network_types,omitempty"`
	// CarrierIDs 运营商ID。 只有当运营商in_use字段为true时生效。获取详细信息，参阅获取运营商列表
	CarrierIDs []string `json:"carrier_ids,omitempty"`
	// IspIDs 定向的互联网服务提供商的ID
	IspIDs []string `json:"isp_ids,omitempty"`
	// DevicePriceRanges 设备价格区间（10000代表1000+）。该数字必须是50的倍数。
	// 重要提示: 最终的实际上限将在您设定的上限基础上增加50，您可以在TikTok广告管理平台的广告组设置中看到实际上限。例如，如果您设置的价格区间是[0, 250]，实际区间则为[0, 300]。
	DevicePriceRanges []int `json:"device_price_ranges,omitempty"`
	// SavedAudienceID 当您在请求中指定了saved_audience_id时，将返回此信息。
	// 已保存受众ID。
	SavedAudienceID string `json:"saved_audience_id,omitempty"`
	// ContextualTagIDs 内容相关定向标签ID列表。请查看内容相关定向了解如何使用内容相关定向。
	ContextualTagIDs []string `json:"contextual_tag_ids,omitempty"`
	// BrandSafetyType 品牌安全类型。枚举值：
	// NO_BRAND_SAFETY：不使用任何品牌安全解决方案，即使用全部库存。您的广告可能会在某些包含成人主题的内容周围展示。
	// EXPANDED_INVENTORY：使用TikTok的品牌安全解决方案，即使用扩展库存。您的广告将不会出现在不当内容或包含成人主题的内容后面。在下个API版本中EXPANDED_INVENTORY将替代NO_BRAND_SAFETY，并成为默认的品牌安全选项。
	// STANDARD_INVENTORY：使用TikTok的品牌安全解决方案中的标准库存。您的广告将展示在适合大多数品牌的内容周围。
	// LIMITED_INVENTORY：使用TikTok的品牌安全解决方案中的有限库存。您的广告将展示在不包含成人主题的内容周围。
	// THIRD_PARTY：使用第三方品牌安全解决方案。
	// 您可以使用/tool/region/接口查询品牌安全设置对应的广告可投放国家或地区。
	// 注意：
	//
	// 出价前第一方品牌安全筛选对于使用应用推广（APP_PROMOTION），网站转化量（WEB_CONVERSIONS），访问量（TRAFFIC），或线索收集（LEAD_GENERATION） 或商品销量（PRODUCT_SALES）推广目标的竞价广告目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 出价前第三方品牌安全解决方案目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	BrandSafetyType enum.BrandSafetyType `json:"brand_safety_type,omitempty"`
	// BrandSafetyPartner 品牌安全合作伙伴。只有在品牌安全为第三方（brand_safety_type = THIRD_PARTY）时有效。枚举值：IAS， OPEN_SLATE（该合作伙伴在TikTok广告管理平台的对应名称为DoubleVerify，因为该合作伙伴已被DoubleVerify收购）。您可以使用/tool/region/接口，并根据自己的品牌安全设置获取可投放的国家和地区。您需要传入品牌安全类型和品牌安全合作伙伴。
	// 注意：
	// 出价前第三方品牌安全解决方案目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	BrandSafetyPartner string `json:"brand_safety_partner,omitempty"`
	// InventoryFilterEnabled 库存筛选（筛选安全视频，如果不安全则不展示），仅对于 PLACEMENT_TIKTOK 版位有效，true表示筛选，false表示不筛选。
	InventoryFilterEnabled bool `json:"inventory_filter_enabled,omitempty"`
	// CategoryExcludeIDs 内容排除类别ID
	CategoryExcludeIDs []string `json:"category_exclude_ids,omitempty"`
	// VerticalSensitiveID 行业敏感内容控制类型ID
	VerticalSensitiveID string `json:"vertical_sensitive_id,omitempty"`
	// BudgetMode 广告预算类型。如果开启了推广系列预算优化(budget_optimize_on)， 将返回BUDGET_MODE_INFINITE。枚举值及设置详见预算
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 广告组预算。
	// 当开启了推广系列预算优化(budget_optimize_on)时，返回0.0。
	// 对于 TopView 广告，本字段代表预估折后预算。
	Budget model.Float64 `json:"budget,omitempty"`
	// ScheduledBudget 表示隔日预算值。若是非0则表示设置了隔日预算并表示了隔日预算的数值；若是为0则表示未设置隔日预算。
	ScheduledBudget model.Float64 `json:"scheduled_budget,omitempty"`
	// ScheduleType 广告投放时间类型。
	// 枚举值: SCHEDULE_FROM_NOW，SCHEDULE_START_END。如果您选择了SCHEDULE_START_END，您需要明确投放开始和结束时间。 如果您选择了SCHEDULE_FROM_NOW，您只需要明确投放开始时间。
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime 广告投放起始时间(UTC+0)。
	// 格式：YYYY-MM-DD HH:MM:SS。
	ScheduleStartTime model.DateTime `json:"schedule_start_time,omitzero"`
	// ScheduleEndTime 广告投放结束时间(UTC+0)。
	// 格式： YYYY-MM-DD HH:MM:SS
	ScheduleEndTime model.DateTime `json:"schedule_end_time,omitzero"`
	// PredictImpression 仅为 TopView 广告返回本字段。
	// 预估展示量，单位为千。
	// 例如，字段值 1291 代表预估展示量为 1,291,000
	PredictImpression int64 `json:"predict_impression,omitempty"`
	// TopviewReachRange 仅为 TopView 广告返回本字段。
	// 预估覆盖人数区间，单位为千。
	// 例如，字段值 [342, 456] 代表预估覆盖人数为 342,000 到 456,000 。
	// 若想获取预估频次区间，可使用以下公式：
	// 频次 = 展示量/覆盖人数。
	// 例如，若 predict_impression 为 1291（即 1,291,000）且topview_reach_range 为 [342, 456]（即 342,000 到 456,000），则预估频次区间为 2.8 到 3.8。
	TopviewReachRange []int64 `json:"topview_reach_range,omitempty"`
	// PreDiscountCPM 仅为 TopView 广告返回本字段。
	// 预估折前 CPM。
	// 示例：11.33。
	PreDiscountCPM model.Float64 `json:"pre_discount_cpm,omitempty"`
	// CPM 仅为 TopView 广告返回本字段。
	// 预估折后 CPM。
	// 示例：0。
	CPM model.Float64 `json:"cpm,omitempty"`
	// DiscountType 仅为 TopView 广告返回本字段。
	// 预算折扣类型。
	// 枚举值：
	// NO_DISCOUNT：无折扣。
	// BY_PERCENTAGE：百分比折扣。
	// BY_AMOUNT：固定金额折扣。
	DiscountType enum.DiscountType `json:"discount_type,omitempty"`
	// DiscountAmount 仅为discount_type为 BY_AMOUNT 的 TopView 广告返回本字段。
	// 预算折扣的固定金额。
	// 示例：14615.8。
	DiscountAmount model.Float64 `json:"discount_amount,omitempty"`
	// DiscountPercentage 仅为discount_type为 BY_PERCENTAGE 的 TopView 广告返回本字段。
	// 预算折扣的百分比。
	DiscountPercentage model.Float64 `json:"discount_percentage,omitempty"`
	// PreDiscountBudget 仅为 TopView 广告返回本字段。
	// 预估折前预算。
	// 示例：14616。
	PreDiscountBudget model.Float64 `json:"pre_discount_budget,omitempty"`
	// ScheduleInfos 覆盖和频次广告组的广告投放信息
	ScheduleInfos []ScheduleInfo `json:"schedule_infos,omitempty"`
	// DeliveryMode 覆盖和频次广告组中广告投放的排序与排期策略。
	// 枚举值:
	// STANDARD：标准投放。每个广告均匀投放，预计会获得大致相同的访问量。
	// SCHEDULE：计划投放。为每个广告设置投放的特定时间段。
	// SEQUENCE：顺序投放。为广告设置特定的投放顺序。
	// OPTIMIZE：优选投放。
	DeliveryMode enum.DeliveryMode `json:"delivery_mode,omitempty"`
	// Dapparting 广告投放安排。格式为48x7的字符串。字符只能为0或者1。0代表不投放，1代表投放。每个字符对应一个30分钟的时间段。第一个字符对应的是周一的凌晨0:01-0:30，第二个字符对应0:31-1:00，依次类推。最后一个字符代表周日23:31-0:00。
	// 注意：
	// 不设置，全部设置为0，或者全部都设置为1都代表了要进行全时投放。
	Dayparting string `json:"dayparting,omitempty"`
	// OptimizationGoal 优化目标
	// 枚举值: 详见枚举值-优化目标。
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// SecondaryOptimizationEvent 当优化目标为INSTALL或VALUE时的次要优化目标。枚举值：详见转化事件-App深度转化事件类型。
	SecondaryOptimizationEvent enum.SecondaryOptimizationEvent `json:"secondary_optimization_event,omitempty"`
	// MessageEventSetID 即时通讯广告组中使用的消息事件集的 ID。
	// 若广告组设置中的即时通讯应用账号匹配到了唯一消息事件集，无论是通过 messaging_app_account_id 手动指定的 Messenger 账号还是基于指定的phone_region_code、phone_region_calling_code 和 phone_number 填充的 WhatsApp 账号，则本字段将自动填充为与所选即时通讯应用账号绑定的唯一消息事件集。
	MessageEventSetID string `json:"message_event_set_id,omitempty"`
	// Frequency 频次，与frequency_schedule共同构成用户看到广告的频次上限（仅适用于REACH广告）。
	// 需要满足以下两个条件：
	// 1 <= frequency <= 1000
	// 1 <= frequency_schedule <=30。
	// 例如：frequency = 2 且 frequency_schedule = 3 表示每 3 天至多 2 次展示。
	Frequency int `json:"frequency,omitempty"`
	// FrequencySchedule 频次天数，与frequency共同构成用户看到广告的频次上限（仅适用于REACH广告）。详见frequency字段说明。
	FrequencySchedule int `json:"frequency_schedule,omitempty"`
	// BidType 竞价策略。
	// 详见枚举值-竞价策略
	BidType enum.BidType `json:"bid_type,omitempty"`
	// BidPrice 出价。（成本上限策略下）想尽可能达到的平均单次成效成本
	BidPrice model.Float64 `json:"bid_price,omitempty"`
	// ConversionBidPrice oCPM转化出价
	ConversionBidPrice model.Float64 `json:"conversion_bid_price,omitempty"`
	// DeepBidType 深度事件出价类型
	// 枚举值: 详见枚举值-深度事件出价类型。
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
	// RoasBid 用于价值优化的 ROAS 目标值
	RoasBid model.Float64 `json:"roas_bid,omitempty"`
	// VboWindow 为广告收入价值优化或应用内购价值优化指定的竞价策略的时间窗口期。
	// 枚举值：
	// SEVEN_DAYS：前 7 天。
	// ZERO_DAY：第 0 天（当天）
	VboWindow enum.Window `json:"vbo_window,omitempty"`
	// BidDisplayMode 每次浏览成本的计算和测算模式。枚举值：CPMV(每千次浏览成本) / CPV(每次浏览成本) 。默认值：CPMV
	BidDisplayMode enum.BidDisplayMode `json:"bid_display_mode,omitempty"`
	// DeepCpaBid 当您为app内事件选择出价策略后，您需要在该字段中传入出价价格
	DeepCpaBid model.Float64 `json:"deep_cpa_bid,omitempty"`
	// CPVVideoDuration 优化目标之视频播放时长
	// 枚举值: SIX_SECONDS（视频播放6秒），TWO_SECONDS（视频播放2秒）
	CPVVideoDuration enum.CPVVideoDuration `json:"cpv_video_duration,omitempty"`
	// NextDayRetention 次日留存率。计算公式：next_day_retention = conversion_bid_price/deep_cpa_bid，范围 (0，1]。仅当placements为PLACEMENT_PANGLE且深度转化事件secondary_optimization_event为NEXT_DAY_OPEN时生效
	NextDayRetention model.Float64 `json:"next_day_retention,omitempty"`
	// ClickAttributionWindow 广告组的点击归因窗口期。该窗口期指用户从点击广告到采取行动之间的时间间隔。
	// 枚举值：
	// OFF：不启用。
	// ONE_DAY：1天（点击）。
	// SEVEN_DAYS：7天（点击）。
	// FOURTEEN_DAYS：14天（点击）。
	// TWENTY_EIGHT_DAYS：28天（点击）。
	ClickAttributionWindow enum.Window `json:"click_attribution_window,omitempty"`
	// EngagedViewAttributionWindow 广告组的深度互动观看归因窗口期。该窗口期指用户从观看视频广告 6 秒到采取行动之间的时间间隔。
	// 枚举值：
	// ONE_DAY：1天（深度互动观看）。
	// SEVEN_DAYS：7天（深度互动观看）。
	EngagedViewAttributionWindow enum.Window `json:"engaged_view_attribution_window,omitempty"`
	// ViewAttributionWindow 广告组的展示归因窗口期。该窗口期指用户从查看广告到采取行动之间的时间间隔。
	// 枚举值：
	// OFF：不启用。
	// ONE_DAY：1天（展示）。
	// SEVEN_DAYS：7天（展示）。
	ViewAttributionWindow enum.Window `json:"view_attribution_window,omitempty"`
	// AttributionEventCount 广告组的事件统计方式（统计类型）。
	// 统计用户在查看或点击了广告后采取操作的次数的方式。
	// 枚举值：
	// UNSET ：未设置。
	// EVERY：每一次。某位用户进行的多个购买事件将分别计入转化量。
	// ONCE：仅一次。某位用户进行的多个事件将只算作 1 次转化。
	AttributionEventCount enum.AttributionEventCount `json:"attribution_event_count,omitempty"`
	// BillingEvent 计费事件。
	// 枚举值: 详见枚举值-计费事件
	BillingEvent enum.BillingEvent `json:"billing_event,omitempty"`
	// Pacing 广告投放速度类型。您可以选择PACING_MODE_SMOOTH（在预定的时间内平均分配预算）和PACING_MODE_FAST（尽快消耗预算并产出结果）。当您开启推广系列预算优化（budget_optimize_on）时，该字段将自动设置为PACING_MODE_SMOOTH。否则您需要填写该字段。
	Pacing enum.PacingMode `json:"pacing,omitempty"`
	// OperationStatus 广告组的操作状态。
	// ENABLE：广告组处于启用（“开”）状态。
	// DISABLE：广告组处于未启用（“关”）状态。
	// FROZEN：广告组处于冻结状态，无法再次启用。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// SecondaryStatus 广告组状态（二级状态)。枚举值详见枚举值-二级状态-广告组状态。
	// 注意：沙箱环境下本字段不返回，因为广告组未实际投放。
	SecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// StatisticType 转化事件出价统计类型。枚举值 EVERYTIME(每一次计费广告)，NONE(仅一次计费广告)
	StatisticType string `json:"statistic_type,omitempty"`
	// IsHfss 广告推广对象是否是 HFSS 食品（高脂、高盐、高糖食品）请注意，欧洲国家不允许向未成年人推送 HFSS 食品广告
	IsHfss bool `json:"is_hfss,omitempty"`
	// CreativeMaterialMode 创意投放方式。
	// 枚举值: CUSTOM（自定义）、DYNAMIC（已废弃）（程序化创意）、SMART_CREATIVE（智能创意）
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// AdgroupAppProfilePageState 该广告组是否使用了App中间页。枚举值：INVALID， UNSET， ON， OFF。
	AdgroupAppProfilePageState enum.AppProfilePageState `json:"adgroup_app_profile_page_state,omitempty"`
	// FeedType 信息流类型。枚举值: STANDARD_FEED， TOP_FEED
	FeedType enum.FeedType `json:"feed_type,omitempty"`
	// RfPurchasedType 覆盖和频次广告购买方式，详见枚举值 - 覆盖和频次广告购买方式。
	// 注意：对于非覆盖和频次广告，本字段的值始终返回 null。
	RfPurchasedType enum.RfPurchasedType `json:"rf_purchased_type,omitempty"`
	// PurchasedImpression 覆盖和频次广告购买的展示量。
	// 注意：对于非覆盖和频次广告，本字段的值始终返回 null。
	PurchasedImpression model.Int64 `json:"purchased_impression,omitempty"`
	// PurchasedReach 覆盖和频次广告购买的触达人数。
	// 注意：对于非覆盖和频次广告，本字段的值始终返回 null。
	PurchasedReach model.Int64 `json:"purchased_reach,omitempty"`
	// RfEstimatedCPR 覆盖和频次广告预测的千人覆盖成本。
	// 注意：对于非覆盖和频次广告，本字段的值始终返回 null。
	RfEstimatedCPR model.Float64 `json:"rf_estimated_cpr,omitempty"`
	// RfEstimatedFrequency 覆盖和频次广告预测的人均展示频次 。
	// 注意：对于非覆盖和频次广告，本字段的值始终返回 null。
	RfEstimatedFrequency model.Float64 `json:"rf_estimated_frequency,omitempty"`
	// SplitTestGroupID 拆分对比测试组ID. 若广告组在拆分对比测试中时返回
	SplitTestGroupID string `json:"split_test_group_id,omitempty"`
	// SplitTestStatus 拆分对比测试状态。若广告组在拆分对比测试中时返回
	SplitTestStatus string `json:"split_test_status,omitempty"`
	// IsNewStructure 推广系列是否是新结
	IsNewStructure bool `json:"is_new_structure,omitempty"`
	// SkipLearningPhase 是否跳过学习阶段
	SkipLearningPhase bool `json:"skip_learning_phase,omitempty"`
}

// AudienceRule 受众规则
type AudienceRule struct {
	// EventSourceIDs 规则的事件源ID列表。当audience_type不为ENGAGEMENT或LEAD_GENERATION时，该字段为必填。
	// 对于广告互动受众，需要填入广告组ID。不填默认使用所有事件源ID。
	// 对于网站访客受众，需要填入Pixel ID。
	// 对于应用事件受众，需要填入App ID。
	// 对于线索生成受众，不能填入该字段，否则会出错。不填默认使用所有事件源ID。
	// 对于商店活动受众，填入店铺ID。（您可以使用 /store/list/接口获取店铺ID。）
	// 对于企业号受众，需要填入广告主的核心用户ID。（您可以使用 /user/info/ 接口获取核心用户ID。）
	// 对于自然流量互动受众，使用TikTok视频ID作为事件源ID。（您可以使用/identity/video/get/ 接口获取TikTok视频ID。）最多填入10个TikTok视频ID。
	// 对于直播互动受众，使用直播视频ID作为事件源ID。 (您可以使用/identity/video/get/接口获取直播视频ID。）最多填入10个直播视频ID。
	// 对于线下活动受众，使用线下事件组ID作为事件源ID。（您可以使用/offline/get/ 获取线下事件组ID。）
	EvetntSourceIDs []string `json:"event_source_ids,omitempty"`
	// RetentionDays 受众留存天数。 枚举值详情可查看枚举值 - 回溯期。
	// 注意：
	// 如果audience_type为BUSINESS_ACCOUNT，且filters对象中的value不为BUSINESS ACCOUNT PROFILE FOLLOW，则retention_days必须为7，14或30。
	// 如果audience_type为ENGAGEMENT_LIVE_VIDEO 或ENGAGEMENT_ORGANIC_VIDEO，retention_days必须为7，14或30。
	RetentionDays int `json:"retention_days,omitempty"`
	// FilterSet 筛选集合列表。格式见下方的“筛选集合格式”。
	FilterSet *model.FilterSet `json:"filter_set,omitempty"`
}

// CustomAction 自定义行为
type CustomAction struct {
	//   Code 用于筛选出受众再营销对象的自定义行为。
	// 枚举值：
	// VIEW_PRODUCT：受众浏览过商品。
	// ADD_TO_CART : 受众将商品加过购物车。
	// PURCHASE : 受众购买过商品。
	Code enum.CustomActionCode `json:"code,omitempty"`
	//   Days 用于筛选出完成了自定义行为的受众的时间区间。
	// 取值范围：[1,180]。
	Days int `json:"days,omitempty"`
}

// TaregetingAction 用于定向的行为分类对象数组
type TargetingAction struct {
	// ActionScene 用于定向的用户行为种类。
	// 枚举值：
	// VIDEO_RELATED：视频互动。
	// CREATOR_RELATED：创作者互动。
	// HASHTAG_RELATED：话题互动。
	ActionScene enum.ActionScene `json:"actiotn_scene,omitempty"`
	// ActionPeriod 选择n天内发生的行为
	ActionPeriod int `json:"action_period,omitempty"`
	// VideoUserActions 所选用户行为种类下想要定向的的具体用户行为。
	// 若action_scene为VIDEO_RELATED，枚举值为：WATCHED_TO_END（完播），LIKED（点赞），COMMENTED（评论），SHARED（分享）。
	// 若action_scene为CREATOR_RELATED，枚举值为：FOLLOWING（关注），VIEW_HOMEPAGE（浏览主页）。
	// 若action_scene为HASHTAG_RELATED，枚举值为： VIEW_HASHTAG（浏览话题标签）。
	VideoUserActions []enum.VideoUserAction `json:"video_user_actions,omitempty"`
	// ActionCategoryIDs 用于定向受众的视频互动分类 ID、创作者互动分类 ID、话题标签 ID 或话题包 ID 列表
	ActionCategoryIDs []string `json:"action_category_ids,omitempty"`
}

// ScheduleInfo 覆盖和频次广告组的广告投放信息
type ScheduleInfo struct {
	// ScheduleID 广告投放信息ID。
	// 您可通过/ad/create/请求中的schedule_id字段将广告投放信息 ID 与广告绑定。广告投放信息与广告绑定后，本字段的值将变为与广告投放信息绑定的广告的 ID，而is_draft字段的值将变为false。
	ScheduleID string `json:"schedule_id,omitempty"`
	// IsDraft 该广告投放信息是否为草稿。
	// 若本字段值为 true，代表您在schedule_infos中定义的广告投放信息还未与广告绑定。
	// 若想将广告投放信息与广告绑定，需将返回的schedule_id值传入/ad/create/请求中的schedule_id字段。广告投放信息与广告绑定后，本字段的值将自动变为false，而schedule_id的值将变为该广告投放信息绑定的广告的 ID。
	IsDraft bool `json:"is_draft,omitempty"`
	// Schedules 为广告设置的顺序投放的详情
	Schedules []Schedule `json:"schedules,omitempty"`
	// ExpectedOrders 该广告组中广告的投放顺序。
	// 例如，值为[1]代表会第一个投放该广告。
	ExpectedOrders []int `json:"expected_orders,omitempty"`
}

// Schedule 为广告设置的顺序投放的详情
type Schedule struct {
	// StartTime 广告投放开始时间
	StartTime model.DateTime `json:"start_time,omitzero"`
	// EndTime 广告投放结束时间
	EndTime model.DateTime `json:"end_time,omitzero"`
}
