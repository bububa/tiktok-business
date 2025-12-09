package spc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/ad/aco"
)

// Campaign 推广系列
type Campaign struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// OperationStatus 推广系列的操作状态。
	// ENABLE：推广系列处于启用（“开”）状态。
	// DISABLE：推广系列处于未启用（“关”）状态。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// CampaignSecondaryStatus 广告状态（二级状态）。
	// 枚举值见枚举值-二级状态。
	// 对应/campaign/get/接口返回的secondary_status。
	CampaignSecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// IsSmartPerformanceCampaign 是否为自动化类型的推广系列。
	// 支持的值：
	// true：推广系列为 Smart+ 推广系列或智能效果网站推广系列。
	// false：推广系列为普通类型的推广系列。
	// 注意：若is_smart_performance_campaign 为 true 且 objective_type 为 WEB_CONVERSIONS，您可根据/campaign/spc/get/返回的spc_type查看该网站推广系列为 Smart+ 网站推广系列还是智能效果网站推广系列。
	IsSmartPerformanceCampaign bool `json:"is_smart_performance_campaign,omitempty"`
	// ObjectiveType 推广目标，获取枚举值，参阅 枚举值-推广目标。
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// SpcType 网站推广系列类型。
	// 枚举值：
	// WEB_ALL_IN_ONE：Smart+ 网站推广系列，支持的优化目标包括转化、价值、点击和落地页浏览量。
	// UNSET：智能效果网站推广系列，支持的优化目标仅包括转化、价值和点击。
	SpcType enum.SpcType `json:"spc_type,omitempty"`
	// WebAllInOneCatalogStatus Smart+ 网站推广系列中的商品库状态。
	// 枚举值：
	// OPEN：启用。
	// UNSET：未启用。
	WebAllInOneCatalogStatus enum.WebAllInOneCatalogStatus `json:"web_all_in_one_catalog_status,omitempty"`
	// AppPromotionType 应用推广类型。枚举值：
	// APP_INSTALL：应用安装。让新用户安装您的应用。
	// APP_RETARGETING：应用再营销。重新吸引现有用户在您的应用中执行操作。
	// APP_PREREGISTRATION：应用预注册。让新用户在您的应用发布前完成预注册。
	// APP_POSTS_PROMOTION：应用帖子推广。推广 TikTok 帖子，提升应用认知度并度量转化量。
	// 注意：目前不支持通过 API 创建应用帖子推广类型的推广系列，仅支持通过 API 更新或检索此类推广系列。
	AppPromotionType enum.AppPromotionType `json:"aepp_promotion_type,omitempty"`
	// VirtualObjectiveType 新推广目标类型。

	// 枚举值：SALES（销量）。
	//
	// 使用本字段创建推广目标为销量的推广系列，该目标合并了网站转化量目标和商品销量目标。了解销量目标的详情，参见销量目标。
	VirtetualObjectiveType enum.VirtualObjectiveType `json:"virtual_objective_type,omitempty"`
	// salesdestination 销售目标页面，即想要推动销售的渠道。
	// 枚举值：
	// tiktok_shop：tiktok shop。推动 tiktok shop 上的销售。
	// website：网站。推动网站上的销售。
	// app：应用。推动应用上的销售（需要商品库）。
	SalesDestination enum.SalesDestination `json:"sales_destination,omitempty"`
	// CampaignType 推广系列类型。
	// 枚举值: REGULAR_CAMPAIGN（普通推广系列）、IOS14_CAMPAIGN（iOS 14专属推广系列）。
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// CampaignAppProfilePageState 下载中间页使用情况。
	// 枚举值：INVALID，UNSET，ON，OFF。
	CampaignAppProfilePageState enum.AppProfilePageState `json:"campaign_app_profile_page_state,omitempty"`
	// IsAdvancedDedicatedCampaign 该推广系列是否为高级专属推广系列。高级专属推广系列采用捕捉实时信号的高级投放模型。
	// 枚举值：true，false。
	IsAdvancedDedicatedCampaign bool `json:"is_advanced_dedicated_campaign,omitempty"`
	// DisableSkanCampaign 是否关闭 SKAN（SKAdNetwork）归因（Apple 针对 iOS 推广系列的转化归因解决方案）。
	// 枚举值：
	// true：关闭 SKAN 归因。推广系列将不受专属推广系列配额限制的约束，您可为此推广系列拉取自归因平台（SAN）指标数据。但是，不支持为此推广系列拉取 SKAN 指标数据。了解自归因平台集成的详情。
	// false：启用 SKAN 归因。推广系列将受专属推广系列配额限制的约束，您可为此推广系列拉取 SKAN 指标数据。
	DisableSkanCampaign bool `json:"disable_skan_campaign,omitempty"`
	// BidAlignType 专属推广系列的归因类型。此类型决定广告组层级设置的目标 CPA （conversion_bid_price） 或目标 ROAS（roas_bid） 基于的归因网络类型。
	// 枚举值：
	// SAN：自归因平台归因。
	// SKAN：SKAN 归因。
	BidAlignType enum.BidAlignType `json:"bid_atlign_type,omitempty"`
	// CampaignName 推广系列名称
	CampaignName string `json:"campaign_name,omitempty"`
	// SpecialIndustries 特殊广告分类。
	// 枚举值:
	// HOUSING：房地产，房屋保险，抵押贷款或相关的广告。
	// EMPLOYMENT：工作机会，实习机会，职业认证项目或相关的广告。
	// CREDIT：信用卡申请，汽车贷款，长期融资或相关的广告。
	// 注意：
	// 在推广系列商品来源为商品库和推广系列商品来源非商品库的购物广告中使用特殊广告分类目前为单独的白名单功能。如需使用这些功能，请联系您的TikTok销售代表。
	// 本字段对于注册地为美国或加拿大的广告主已全量。若您是注册地非美国或加拿大的广告主且想在定向美国或加拿大的广告中使用特殊广告分类，您需申请额外的白名单。
	SpecialIndustries []enum.SpecialIndustry `json:"special_industries,omitempty"`
	// AdgroupSecondaryStatus 广告组状态（二级状态）。
	// 枚举值参见枚举值-二级状态。
	// 对应/adgroup/get/接口返回的secondary_status。
	AdgroupSecondaryStatus enum.SecondaryStatus `json:"adgroup_secondary_status,omitempty"`
	// ProductSource 仅当推广系列对应的objective_type为PRODUCT_SALES，且设置了商品来源时返回。
	// 推广系列的商品来源。
	// 枚举值：
	// CATALOG ：商品库。
	// STORE：TikTok Shop 店铺，或 TikTok 橱窗。
	ProductSource enum.ProductSource `json:"campaign_product_source,omitempty"`
	// CatalogID Smart+ 商品库广告中要使用的商品库的 ID
	CatalogID string `json:"catalog_id,omitempty"`
	// CatalogAuthorizedBcID 商品库（catalog_id）所属的商务中心的 ID
	CatalogAuthorizedBcID string `json:"catalog_authorized_bc_id,omitempty"`
	// ProductSpecificType 选择商品的维度。
	//  枚举值：
	// ALL：允许 TikTok 从所有商品中动态选择。
	// PRODUCT_SET：请选择一个商品系列。TikTok 将动态选择该商品系列中的商品。
	// CUSTOMIZED_PRODUCTS：选择自定义数量的特定商品
	ProductSpecificType enum.ProductSpecificType `json:"product_specific_type,omitempty"`
	// ProductSetID 商品系列 ID
	ProductSetID string `json:"prodduct_set_id,omitempty"`
	// ProductIDs 商品库商品的商品 ID 列表
	ProductIDs []string `json:"product_ids,omitempty"`
	// PromotionType 推广对象类型（优化位置）。您可以使用该字段决定在哪里投放广告。
	// 枚举值情见枚举值-推广对象类型。
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// AppID 推广的应用的 ID。
	// 您可通过/app/list/获取app_id。
	AppID string `json:"app_id,omitempty"`
	// AppType 推广的应用的的类型。
	// 目前仅支持APP_ANDROID和APP_IOS。
	AppType enum.AppType `json:"app_type,omitempty"`
	// PromotionWebsiteType TikTok 即时体验页面类型。
	// 目前仅支持TIKTOK_NATIVE_PAGE（使用 TikTok 即时体验页面）
	PromotionWebsiteType enum.PromotionWebsiteType `json:"promotion_website_type,omitempty"`
	// PromotionTargetType 推广目标线索收集的专属优化位置。
	// 枚举值：
	// INSTANT_PAGE: 即时表单（线索表单）。创建可快速加载的应用内 TikTok 即时表单来收集更多线索。
	// EXTERNAL_WEBSITE: 第三方网站表单。使用包含第三方网站表单的落地页，或使用会重定向至包含第三方网站表单网站的 TikTok 即时体验页面来收集更多线索。
	PromotionTargetType enum.PromotionTargetType `json:"promotion_target_type,omitempty"`
	// OptimizationGoal 广告使用的优化目标。枚举值详见枚举值-优化目标。
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// PixelID Pixel ID。仅对落地页生效
	PixelID string `json:"pixel_id,omitempty"`
	// OptimizationEvent 广告组转化事件。
	// 详情见转化事件。
	OptimizationEvent enum.OptimizationEvent `json:"optimization_event,omitempty"`
	// DeepFunnelOptimizationStatus 深层漏斗优化的状态。
	// 深层漏斗优化同时优化浅层和深层漏斗事件。您可以在主要优化事件之外再选择一个次要事件，这有助于提升推广 系列的效果。
	//
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
	// LocationIDs 定向地域 ID。使用tool/region 获取定向地域ID
	LocationIDs []string `json:"location_ids,omitempty"`
	// Langauges 受众语言。
	// 枚举值：详见枚举值-受众语言
	Languages []string `json:"languages,omitempty"`
	// Gender 定向受众性别
	// 枚举值: GENDER_FEMALE,GENDER_MALE,GENDER_UNLIMITED。
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// ExcludeAgeUnderEighteen 是否排除十八岁以下受众
	ExcludeAgeUnderEighteen bool `json:"exclude_age_under_eighteen,omitempty"`
	// SpcAudienceAge 该推广系列定向的年龄区间。
	//
	// 举值：18+，25+。
	//
	// 对于推广目标为 LEAD_GENERATION 的推广系列（Smart+ 线索收集推广系列），本字段的值为18+ 或 25+。
	// 对于推广目标为 APP_PROMOTION的推广系列（Smart+ 应用推广系列）或WEB_CONVERSIONS 的推广系列（Smart+ 网站推广系列或智能效果网站推广系列），本字段的值固定为 18+。
	//
	// 注意：目前，不支持通过 API 创建 Smart+ 线索收集推广系列。若想筛选您的广告账号中存量的 Smart+ 线索收集推广系列，可调用/campaign/get/ 并将filtering设置为 {"objective_type":"LEAD_GENERATION", "is_smart_performance_campaign": true}。
	SpcAudienceAge string `json:"spc_audience_age,omitempty"`
	// ExcludeAudienceIDs 排除受众 ID 列表
	ExcludeAudienceIDs []string `json:"exclude_audience_ids,omitempty"`
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
	// CommentDisabled 是否允许用户在TikTok上评论您的广告
	CommentDisabled bool `json:"comment_disabled,omitempty"`
	// ShareDisabled 本广告组中的广告是否禁止分享到第三方平台
	ShareDisabled bool `json:"share_disabled,omitempty"`
	// VideoDownloadDisabled 用户是否可以在 TikTok 上下载您的广告视频
	VideoDownloadDisabled bool `json:"video_download_disabled,omitempty"`
	// BlockedPangleAppIDs Pangle 媒体屏蔽列表 ID
	BlockedPangleAppIDs []string `json:"blocked_pangle_app_ids,omitempty"`
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
	// CategoryExcludeIDs 不希望在该 Smart+ 推广系列前后显示的内容排除类别 ID 列表。
	// 若想根据类别 ID 获取内容排除类别的详细信息，可使用/tool/content_exclusion/info/。
	// 注意: 该设置将根据您的品牌安全中心配置自动应用。若要检索广告账户这些设置的详情，可使用/TikTok_inventory_filters/get/。
	CategoryExcludeIDs []string `json:"category_exclude_ids,omitempty"`
	// BudgetMode 广告预算类型。如果开启了推广系列预算优化(budget_optimize_on)， 将返回BUDGET_MODE_INFINITE。枚举值及设置详见预算
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 预算。
	// 当开启了推广系列预算优化(budget_optimize_on)时，返回0.0。
	// 对于 TopView 广告，本字段代表预估折后预算。
	Budget model.Float64 `json:"budget,omitempty"`
	// BudgetAutoAdjustStrategy 仅当 budget_mode 为 BUDGET_MODE_DYNAMIC_DAILY_BUDGET 时生效。
	// 日预算策略。
	// 枚举值：
	// AUTO_BUDGET_INCREASE：启用自动增加预算。当广告成效良好并满足目标 CPA、第 0 天目标 ROAS 和预算要求时，允许预算自动增加。
	// budget_auto_adjust_strategy 为 AUTO_BUDGET_INCREASE 时，指定的 budget 为初始日预算。当预算使用率达到 90% 或以上时，允许日预算自动增加 20%，每天最多可增加 10 次。日预算每天都将重置为初始日预算。
	// 注意：启用自动增加预算目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	BudgetAutoAdjustStrategy enum.BudgetAutoAdjustStrategy `json:"budget_auto_adjust_strategy,omitempty"`
	// CurrentBudget budget_mode 为BUDGET_MODE_DYNAMIC_DAILY_BUDGET 时返回本字段。
	// 启用了自动增加预算的推广系列的当前预算。
	CurrentBudget model.Float64 `json:"currenct_budget,omitempty"`
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
	// Dapparting 广告投放安排。格式为48x7的字符串。字符只能为0或者1。0代表不投放，1代表投放。每个字符对应一个30分钟的时间段。第一个字符对应的是周一的凌晨0:01-0:30，第二个字符对应0:31-1:00，依次类推。最后一个字符代表周日23:31-0:00。
	// 注意：
	// 不设置，全部设置为0，或者全部都设置为1都代表了要进行全时投放。
	Dayparting string `json:"dayparting,omitempty"`
	// SkipLearningPhase 是否跳过学习阶段
	SkipLearningPhase bool `json:"skip_learning_phase,omitempty"`
	// IdentityType 不使用 Spark Ads 时的认证身份类型。
	// 枚举值：CUSTOMIZED_USER。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityID 不使用 Spark Ads 时的认证身份 ID
	IdentityID string `json:"identity_id,omitempty"`
	// AppDownloadURL App下载链接
	AppDownloadURL string `json:"app_download_url,omitempty"`
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
	// MediaInfoList 媒体信息列表
	MediaInfoList []aco.MediaInfo `json:"media_info_list,omitempty"`
	// CatalogCreativeToggle 是否启用商品库素材的自动选择。
	//  支持的值：true，false。
	// 若本字段设置为 true，系统将使用可用的广告样式（商品库视频、商品库轮播和单视频）自动生成广告。
	CatalogCreativeToggle bool `json:"catalog_creative_toggle,omitempty"`
	// TitleList 广告文案列表。广告文案作为您的广告素材一部分向受众展示，传递您想传达的信息
	TitleList []aco.Title `json:"title_list,omitempty"`
	// CallToActionList 行动引导文案列表
	CallToActionList []aco.CallToAction `json:"call_to_action,omitempty"`
	// CallToActionID 行动引导文案素材包 ID
	CallToActionID string `json:"call_to_action_id,omitempty"`
	// CardList 卡片ID列表。长度范围：[0,1]
	CardList []aco.Card `json:"card_list,omitempty"`
	// AutomaticAddOnEnabled 是否启用创新互动样式的自动选择。
	// 支持的值：true，false。
	//  若本字段设置为 true，系统将在广告中自动应用最合适的创新互动样式。
	AutomaticAddOnEnabled bool `json:"automatic_add_on_enabled,omitempty"`
	// PageList 页面 ID 列表
	PageList []aco.Page `json:"page_list,omitempty"`
	// DeeplinkType 深度链接类型
	DeeplinkType enum.DeeplinkType `json:"deeplink_type,omitempty"`
	// Deeplink 深度链接，用户下载应用后点击链接跳转至的特定页面
	Deeplink string `json:"deeplink,omitempty"`
	// LandingPageURLs 落地页URL列表
	LandingPageURLs []aco.LandingPageURL `json:"landing_page_urls,omitempty"`
	// ImpressionTrackingURL 默认展示监测 URL
	ImpressionTrackingURL string `json:"impression_tracking_url,omitempty"`
	// ClickTrackingURL 点击监测 URL
	ClickTrackingURL string `json:"click_tracking_url,omitempty"`
	// CreateTime 推广系列的创建时间，格式为YYYY-MM-DD HH:MM:SS（UTC+0）
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ModifyTime 推广系列的修改时间，格式为YYYY-MM-DD HH:MM:SS（UTC+0）
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
}
