package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/adgroup"
)

// Adgroup Upgraded Smart+ ad group
type Adgroup struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName 该广告组所属的推广系列的名称。
	CampaignName string `json:"campaign_name,omitempty"`
	// AdgroupID 广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName 广告组名称
	AdgroupName string `json:"adgroup_name,omitempty"`
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
	// PromotionType 推广对象类型（优化位置）。若对应的推广系列推广目标不是REACH，VIDEO_VIEWS或ENGAGEMENT，则此字段必填。
	// 枚举值详见 枚举值-推广对象类型。
	// 注意: 如果在推广系列层级的objective_type是 ENGAGEMENT，则该字段只能使用 EXTERNAL_OR_DISPLAY。
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// AppID 推广的App的ID。您可通过/app/list/获取app_id
	AppID string `json:"app_id,omitempty"`
	// GamingAdComplianceAgreement Returned only in any of the following scenarios:
	// Scenario 1: When the following conditions are both met at the campaign level:
	// objective_type is APP_PROMOTION.
	// app_promotion_type is APP_RETARGETING.
	// Scenario 2: When the following conditions are all met at the campaign level:
	// objective_type is APP_PROMOTION.
	// app_promotion_type is APP_INSTALL.
	// campaign_type is REGULAR_CAMPAIGN.
	// Whether to agree to the Compliance Assurance Policy for Gaming Advertisers on TikTok.
	// The policy is as follows: Yo confirms and attest that any gaming application, product or service (game) you desire to advertise on TikTok, including any associated URL(s), (a) complies with all applicable laws and regulations of the jurisdictions where the game can be accessed or played, and upon request, can provide supporting documentation as evidence of why the game is not considered illegal gambling or lottery; and (b) has not been and is not part of any investigation or lawsuit regarding the game's legality or regulatory compliance.
	// Enum values:
	// ON: To agree to the policy.
	// OFF: To leave the policy not accepted.
	GamingAdComplianceAgreement enum.OnOff `json:"gaming_ad_compliance_agreement,omitempty"`
	// PromotionWebsiteType 广告组中的 TikTok 即时体验页面类型。
	// 枚举值：
	// UNSET：不使用TikTok 即时体验页面。
	// TIKTOK_NATIVE_PAGE：使用 TikTok 即时体验页面。
	PromotionWebsiteType enum.PromotionWebsiteType `json:"promotion_website_type,omitempty"`
	// OptimizationGoal The measurable results that you'd like to drive your ads with.
	// For enum values, see Enumeration - Optimization Goal.
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// PixelID pixel ID
	PixelID string `json:"pixel_id,omitempty"`
	// AppConfig 推广系列层级 sales_destination 为 WEB_AND_APP 时必填。
	// 想要推广的应用详情。
	// 最大数量：2。
	// 您可以通过本字段指定以下任意类型应用：
	// 一个安卓应用
	// 一个 iOS 应用
	// 一个安卓应用和一个 iOS 应用
	AppConfig []adgroup.AppConfig `json:"app_config,omitempty"`
	// MinisID Returned when promotion_type is MINI_APP or MINI_GAME.
	// The ID of the TikTok Minis.
	MinisID string `json:"minis_id,omitempty"`
	// Otiktok_subplacementsptimizationEvent 广告组转化事件。参阅转化事件 ，获取详细信息
	OptimizationEvent enum.OptimizationEvent `json:"optimization_event,omitempty"`
	// CustomConversionID The ID of the Custom Conversion used in the ad group.
	CustomConversionID string `json:"custom_conversion_id,omitempty"`
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
	// IdentityID 当 shopping_ads_type 为 VIDEO 且 product_source 为 SHOWCASE 或 shopping_ads_type 为 LIVE 时返回。
	// 认证身份 ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 当 shopping_ads_type 为 VIDEO 且 product_source 为 SHOWCASE 或 shopping_ads_type 为 LIVE 时返回。
	// 认证身份类型。枚举值: AUTH_CODE (授权码认证)， TT_USER (TikTok商务用户)， BC_AUTH_TT（已添加到商务中心的TikTok用户）。参阅认证身份 获取详细信息。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 当identity_type 为 BC_AUTH_TT时返回。
	// 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// MessagingAppType The type of instant messaging app or customized URL to use in the Instant Messaging Ad Group.
	// Enum values:
	// MESSENGER: Messenger.
	// WHATSAPP: WhatsApp.
	// ZALO: Zalo.
	// LINE: Line.
	// IM_URL: Instant Messaging URL.
	// To learn more about how to create Upgraded Smart+ TikTok Instant Messaging Ads, see Create an Upgraded Smart+ Lead Generation Campaign with optimization location as instant messaging apps.
	MessagingAppType enum.MessagingAppType `json:"messaging_app_type,omitempty"`
	// ZaloIDType The type of Zalo contact format.
	// Enum values:
	// ZALO_OFFICIAL_ACCOUNT: Zalo Official Account ID.
	// ZALO_PHONE_ACCOUNT: Zalo phone number.
	ZaloIDType enum.ZaloIDType `json:"zalo_id_type,omitempty"`
	// MessagingAppAccountID The ID of the instant messaging app account.
	// When messaging_app_type is MESSENGER, this field represents the Facebook Page ID.
	// When messaging_app_type is LINE, this field represents the LINE Business ID.
	// When zalo_id_type is ZALO_OFFICIAL_ACCOUNT, this field represents the Zalo Official Account ID.
	// When messaging_app_type is WHATSAPP or when zalo_id_type is ZALO_PHONE_ACCOUNT, this field represents the WhatsApp or Zalo phone number automatically populated based on the specified phone_info.
	MessagingAppAccountID string `json:"messaging_app_account_id,omitempty"`
	// MessageEventSetID The ID of the message event set to be used in the Instant Messaging Ad Group.
	// If the instant messaging app account, either the Messenger account or Zalo Official Account specified via messaging_app_account_id or the WhatsApp or Zalo phone account matched from the specified phone_info, in your ad group settings matches an existing event set, this field will be automatically populated with the unique message event set associated with the instant messaging app account you choose.
	MessageEventSetID string `json:"message_event_set_id,omitempty"`
	// PhoneInfo Details of WhatsApp or Zalo phone number
	PhoneInfo *PhoneInfo `json:"phone_info,omitempty"`
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
	// IncentiveOfferType The type of incentive offer applicable to the Upgraded Smart+ Ad Group.
	// If your ad group meets the eligibility criteria and exceeds certain CPA or Minimum ROAS/Target ROAS thresholds, you will be incentivized with ad credits to use within the same ad account.
	// To learn more about the incentive offer eligibility criteria and the calculation of incentive amount, see Smart+ Platform Incentive Offer (Cost Cap/Minimum ROAS/Target ROAS).
	// Enum values:
	// INELIGIBLE: The ad group is ineligible for any incentive offer.
	// COST_CAP_AND_MIN_ROAS: The ad group uses the Cost Cap or Minimum ROAS/Target ROAS bidding strategy and is eligible for the incentive offer.
	IncentiveOfferType enum.IncentiveOfferType `json:"incentive_offer_type,omitempty"`
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
	// THIRTY_DAYS: 30-day click.
	// THIRTY_TWO_DAYS: 32-day click.
	ClickAttributionWindow enum.Window `json:"click_attribution_window,omitempty"`
	// EngagedViewAttributionWindow 广告组的深度互动观看归因窗口期。该窗口期指用户从观看视频广告 6 秒到采取行动之间的时间间隔。
	// 枚举值：
	// ONE_DAY：1天（深度互动观看）。
	// SEVEN_DAYS：7天（深度互动观看）。
	// FOURTEEN_DAYS: 14-day engaged view.
	// TWENTY_EIGHT_DAYS: 28-day engaged view.
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
	// BudgetMode 广告预算类型。如果开启了推广系列预算优化(budget_optimize_on)， 将返回BUDGET_MODE_INFINITE。枚举值及设置详见预算
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// BudgetAutoAdjustStrategy Returned only when the following conditions are both met:
	// At the campaign level: budget_optimize_on is false.
	// At the ad group level: budget_mode is BUDGET_MODE_DYNAMIC_DAILY_BUDGET.
	// The ad group budget strategy for custom ad group budget.
	// Enum values:
	// AUTO_BUDGET_INCREASE: To enable automatic budget increase. Allow your budget to automatically increase when your ads are performing well and target CPA, Day 0 target ROAS, and budget requirements are met.
	// When budget_auto_adjust_strategy is AUTO_BUDGET_INCREASE, the specified budget will be the initial daily budget. Your daily budget will be allowed to automatically increase by 20%, up to 10 times per day, when your budget utilization reaches 90% or more. Your daily budget will reset to your original daily budget each day.
	// UNSET: To disable automatic budget increase.
	BudgetAutoAdjustStrategy enum.BudgetAutoAdjustStrategy `json:"budget_auto_adjust_strategy,omitempty"`
	// Budget 广告组预算。
	// 当开启了推广系列预算优化(budget_optimize_on)时，返回0.0。
	// 对于 TopView 广告，本字段代表预估折后预算。
	Budget model.Float64 `json:"budget,omitempty"`
	// CurrentBudget Returned only when budget_auto_adjust_strategy  is AUTO_BUDGET_INCREASE.
	// Current ad group budget for an ad group with automatic budget increase enabled.
	CurrentBudget model.Float64 `json:"current_budget,omitempty"`
	// MinBudget Ad group minimum budget.
	// The system will aim to spend at least this amount, but it is not guaranteed
	MinBudget model.Float64 `json:"min_budget,omitempty"`
	// ScheduleType 广告投放时间类型。
	// 枚举值: SCHEDULE_FROM_NOW，SCHEDULE_START_END。如果您选择了SCHEDULE_START_END，您需要明确投放开始和结束时间。 如果您选择了SCHEDULE_FROM_NOW，您只需要明确投放开始时间。
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime 广告投放起始时间(UTC+0)。
	// 格式：YYYY-MM-DD HH:MM:SS。
	ScheduleStartTime model.DateTime `json:"schedule_start_time,omitzero"`
	// ScheduleEndTime 广告投放结束时间(UTC+0)。
	// 格式： YYYY-MM-DD HH:MM:SS
	ScheduleEndTime model.DateTime `json:"schedule_end_time,omitzero"`
	// MoviePremiereDate The theatrical release date, in the format of YYYY-MM-DD (UTC+0).
	// Providing the movie release date allows the system to incorporate this timing into performance enhancements.
	MoviePremiereDate model.DateTime `json:"movie_premiere_date,omitzero"`
	// Dapparting 广告投放安排。格式为48x7的字符串。字符只能为0或者1。0代表不投放，1代表投放。每个字符对应一个30分钟的时间段。第一个字符对应的是周一的凌晨0:01-0:30，第二个字符对应0:31-1:00，依次类推。最后一个字符代表周日23:31-0:00。
	// 注意：
	// 不设置，全部设置为0，或者全部都设置为1都代表了要进行全时投放。
	Dayparting string `json:"dayparting,omitempty"`
	// TargetingOptimizationMode Audience targeting optimization mode.
	// Enum values:
	// MANUAL: Custom targeting. You can use custom targeting settings to precisely control who sees your ads. This may limit delivery and impact campaign performance.
	// AUTOMATIC: Automatic targeting. You can use automatic targeting to leverage real-time data and machine learning to target audiences most likely to engage with your ads.
	TargetingOptimziationMode enum.TargetingOptimizationMode `json:"targeting_optimization_mode,omitempty"`
	// SuggestionAudienceEnabled Whether to enable audience suggestions.
	// Audience suggestions guide automatic targeting by choosing additional audience settings. These serve as suggestions only, and delivery to those audiences is not guaranteed.
	// Supported values: true, false.
	SuggestionAudienceEnabled bool `json:"suggestion_audience_enabled,omitempty"`
	// TargetingSpec Targeting settings
	TargetingSpec *Targeting `json:"targeting_spec,omitempty"`
	// IsHfss 广告推广对象是否是 HFSS 食品（高脂、高盐、高糖食品）请注意，欧洲国家不允许向未成年人推送 HFSS 食品广告
	IsHfss bool `json:"is_hfss,omitempty"`
	// IsLhfCompliance 推广内容是否符合 LHF（较不健康食品）合规要求。
	// 将 is_lhf_compliance 设置为 true，即表示您确认在英国 TikTok 上推广的任何食品或饮料均符合 2024 年较不健康食品法规以及其他所有适用法律。
	// 支持的值：true，false。
	// 默认值：false。
	// 注意：自 2026 年 1 月 1 日起，当广告定向到英国地域且 is_hfss 为 true 时，is_lhf_compliance 必填，且必须设置为 true。
	IsLhfCompliance bool `json:"is_lhf_compliance,omitempty"`
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
	// TiktokSubplacements he subplacements within TikTok for your ads, allowing you to choose where your ads will appear.
	// Enum values:
	// LEMON8: Lemon8, a community app for lifestyle content, focusing on real-life experiences, tips, guides, and product reviews. By including Lemon8 as a subplacement, your ads will appear in its For You feed, Search feed, and Immersive Video feed. Learn more about Lemon8 for Performance Auction.
	// PINE_DRAMA: PineDrama, an app for watching short-form dramas. Ads will be placed on recommendation pages, within episode streams, and more. Learn more about PineDrama for Performance Auction.
	TiktokSubplacements []enum.TiktokSubplacement `json:"tiktok_subplacements,omitempty"`
	// SearchResultEnabled 是否在搜索广告中展示您的广告，即当用户在 TikTok上搜索您的业务相关信息时，系统是否向他们展示您的广告。
	SearchResultEnabled bool `json:"search_result_enabled,omitempty"`
	// CommentDisabled 是否允许用户在TikTok上评论您的广告
	CommentDisabled bool `json:"comment_disabled,omitempty"`
	// ShareDisabled 本广告组中的广告是否禁止分享到第三方平台
	ShareDisabled bool `json:"share_disabled,omitempty"`
	// VideoDownloadDisabled 用户是否可以在 TikTok 上下载您的广告视频
	VideoDownloadDisabled bool `json:"video_download_disabled,omitempty"`
	// SkipLearningPhase 是否跳过学习阶段
	SkipLearningPhase bool `json:"skip_learning_phase,omitempty"`
	// CreateTime 广告组创建时间。格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2023-01-01 00:00:01".
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ModifyTime 广告组修改时间。格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2023-01-01 00:00:01".
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
}

// PhoneInfo Details of WhatsApp or Zalo phone number
type PhoneInfo struct {
	// PhoneRegionCode The region code for WhatsApp or Zalo phone number.
	PhoneRegionCode string `json:"phone_region_code,omitempty"`
	// PhoneRegionCallingCode The region calling code for the WhatsApp or Zalo phone number.
	PhoneRegionCallingCode string `json:"phone_region_calling_code,omitempty"`
	// PhoneNumber The WhatsApp or Zalo phone number
	PhoneNumber string `json:"phone_number,omitempty"`
}

// Targeting Targeting settings
type Targeting struct {
	// AppTargetingType Returned only when the following conditions are all met:
	// At the campaign level:
	// objective_type is WEB_CONVERSIONS
	// sales_destination is APP.
	// optimization_goal is CLICK, IN_APP_EVENT, or VALUE.
	// App targeting type.
	// Enum values:
	// PROSPECT: Prospecting. Find prospective customers, including those who have not interacted with your products.
	// RETARGETING: Retargeting. Show ads to people who have already interacted with your business on and off TikTok.
	AppTargetingType enum.AppTargetingType `json:"app_targeting_type,omitempty"`
	// LocationIDs IDs of the locations that you want to target.
	LocationIDs []string `json:"location_ids,omitempty"`
	// ZipcodeIDs Zip code IDs or postal code IDs of the targeted locations
	ZipcodeIDs []string `json:"zipcode_ids,omitempty"`
	// SpcAudienceAge The age group that the ad group targets.
	// Enum values:
	// ALL: all age groups.
	// OVER_EIGHTEEN: 18+.
	// OVER_TWENTY_FIVE: 25+.
	SpcAudienceAge string `json:"spc_audience_age,omitempty"`
	// Languages Codes of the languages that you want to target.
	// For the list of language codes supported, see Enumerations - Language Code.
	Languages []string `json:"languages,omitempty"`
	// OperationSystem Device operating systems that you want to target.
	// Enum values: ANDROID, IOS
	OperationSystem []enum.OperatingSystem `json:"operation_system,omitempty"`
	// ExcludedAudienceIDs List of audience IDs to be excluded.
	ExcludedAudienceIDs []string `json:"excluded_audience_ids,omitempty"`
	// AgeGroups Age groups you want to target.
	// For enum values, see Enumeration - Age Group
	AgeGroups []enum.Age `json:"age_groups,omitempty"`
	// Gender Gender that you want to target.
	// Enum values: GENDER_FEMALE, GENDER_MALE, GENDER_UNLIMITED.
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// AudienceIDs A list of audience IDs
	AudienceIDs []string `json:"audience_ids,omitempty"`
	// ShoppingAdsRetargetingType Valid only when the following conditions are all met:
	// At the campaign level:
	// 	objective_type is WEB_CONVERSION
	// 	sales_destination is WEB_AND_APP
	// 	catalog_type is ECOMMERCE
	// At the ad group level: targeting_optimization_mode is MANUAL.
	// The retargeting type.
	// Enum values:
	// LAB1: Retargeting audiences who viewed products or added products to cart but didn't purchase products.
	// LAB2: Retargeting audiences who added products to cart but didn't purchase products.
	// LAB3: Retargeting audiences using custom combination.
	// OFF: No retargeting.
	// Default value: OFF.
	ShoppingAdsRetargetingType enum.ShoppingAdsRetargetingType `json:"shopping_ads_retargeting_type,omitempty"`
	// ShoppingAdsRetargetingActionsDays Required when shopping_ads_retargeting_type is LAB1 or LAB2.
	// The valid time range for the specified audience action. Audiences who have completed the specified action within the time range will be retargeted.
	// Value range: 1, 2, 3, 7, 14, 30, 60, 90, 180.
	ShoppingAdsRetargetingActionsDays int `json:"shopping_ads_retargeting_actions_days,omitempty"`
	// IncludedCustomActions When shopping_ads_retargeting_type is LAB3, you need to specify either included_custom_actions or excluded_custom_actions.
	// Details of the catalog audience to include.
	// Catalog audience is based on people's interactions with specific products and often drives better performance than custom audience.
	IncludedCustomActions []TargetingCustomAction `json:"included_custom_actions,omitempty"`
	// ExcludedCustomActions When shopping_ads_retargeting_type is LAB3, you need to specify either included_custom_actions or excluded_custom_actions.
	// Details of the catalog audience to exclude.
	// Improve ad performance by excluding products that people have already interacted with, ensuring they only see relevant ads from your brand.
	ExcludedCustomActions []TargetingCustomAction `json:"excluded_custom_actions,omitempty"`
	// ShoppingAdsRetargetingCustomAudienceRelation valid only when the following conditions are both met:
	// shopping_ads_retargeting_type is set to LAB1, LAB2, or LAB3.
	// audience_ids is specified.
	//
	// The logical relation between the retargeting audience specified by shopping_ads_retargeting_type and the custom audience specified by audience_ids.
	//
	// Enum values:
	// OR: To combine the retargeting audience and the custom audience to create the targeted audience. The ad group will target anyone in catalog or custom audiences.
	// AND: To intersect between the retargeting audience and the custom audience to create the targeted audience. The ad group will target individuals who belong to both the retargeting audience and the custom audience.
	//
	// If this field is not set, the targeted audience will consist of individuals who belong to both the retargeting audience and the custom audience.
	//
	// Note: Once set, this field cannot be updated to a null value.
	ShoppingAdsRetargetingCustomAudienceRelation enum.FilterSetOperator `json:"shopping_ads_retargeting_custom_audience_relation,omitempty"`
	// IncludedPangleAudiencePackageIDs IDs of the Pangle audiences that you want to include
	IncludedPangleAudiencePackageIDs []string `json:"included_pangle_audience_package_ids,omitempty"`
	// ExcludedPangleAudiencePackageIDs IDs of the Pangle audiences that you want to exclude.
	ExcludedPangleAudiencePackageIDs []string `json:"excluded_pangle_audience_package_ids,omitempty"`
	// InterestCategoryIDs IDs of general interest keywords that you want to use to target audiences
	InterestCategoryIDs []string `json:"interest_category_ids,omitempty"`
	// InterestKeywordIDs IDs of additional interest categories that you want to use to target audience
	InterestKeywordIDs []string `json:"interest_keyword_ids,omitempty"`
	// PurchaseIntentionKeywordIDs IDs of purchase intention categories that you want to use to target audiences with an interest in purchases related to a content category
	PurchaseIntentionKeywordIDs []string `json:"purchase_intention_keyword_ids,omitempty"`
	// Actions A list of targeting behavioral category objects
	Actions []adgroup.TargetingAction `json:"actions,omitempty"`
	// SmartInterestBehaviorEnabled 是否已启用智能型兴趣和行为。
	// 枚举值： true， false。
	// 若想了解智能型兴趣和行为详情及如何启用智能型兴趣和行为，请查看智能定向。
	SmartInterestBehaviorEnabled bool `json:"smart_interest_behavior_enabled,omitempty"`
	// SpendingPower 消费能力。枚举值：ALL, HIGH。如果设为HIGH，您就可以定向和普通用户相比 TikTok 广告投入支出通常更多的高消费能力用户。
	SpendingPower enum.SpendingPower `json:"spending_power,omitempty"`
	// HouseholdIncome 您想定向的收入群体。枚举值：TOP5(家庭收入群体的前5%)， TOP10(家庭收入群体的前10%)，TOP10_25(家庭收入群体的前10-25%)， TOP25_50(家庭收入群体的前25-50%)。
	HouseholdIncome []enum.HouseholdIncome `json:"household_income,omitempty"`
	// MinAndroidVersion 受众最低Android版本。枚举值：详见枚举值-最低Android版本。
	MinAndroidVersion string `json:"min_android_version,omitempty"`
	// MinIosVersion 受众最低 iOS 版本。枚举值：详见枚举值-最低 iOS 版本
	MinIosVersion string `json:"min_ios_version,omitempty"`
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
	// CategoryExcludeIDs 内容排除类别ID
	CategoryExcludeIDs []string `json:"category_exclude_ids,omitempty"`
}

// TargetingCustomAction When shopping_ads_retargeting_type is LAB3, you need to specify either included_custom_actions or excluded_custom_actions.
type TargetingCustomAction struct {
	// Code The custom action used to filter out the audiences.
	// Enum values:
	// VIEW_PRODUCT: The audience viewed the product.
	// ADD_TO_CART: The audience added the product to the cart.
	// PURCHASE: The audience purchased the product.
	Code enum.TargetingCustomActionCode `json:"code,omitempty"`
	// Days The time range used to filter out the audiences that completed the specified action.
	// Value range: 1-180.
	Days int `json:"days,omitempty"`
}
