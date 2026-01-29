package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/adgroup"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest Create an Upgraded Smart+ Ad Group API Request
type CreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// RequestID 请求ID，用于创建同名广告组。该ID支持接口幂等，避免重复请求。如果您传入相同的request ID重试多次请求，只有一次请求会成功。
	// 该ID和返回参数中的 request_id 不同。返回的request_id用于唯一标识一次 HTTP 请求。
	// 本字段的值需为字符串格式的64位整数值。
	RequestID string `json:"request_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// OperationStatus 广告组的操作状态。
	// ENABLE：广告组处于启用（“开”）状态。
	// DISABLE：广告组处于未启用（“关”）状态。
	// 若想在广告组创建后更新广告组的启用/关闭状态，请使用 /adgroup/update/status/ 接口。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// AdgroupName 广告组名，长度小于 512 字符，不能包含emoji。
	AdgroupName string `json:"adgroup_name,omitempty"`
	// CatalogID 商品库ID
	// - 若 product_source 设置为 CATALOG 或 STORE，本字段必填。
	// - 若 shopping_ads_type 设置为 LIVE，则字段 catalog_id，catalog_authorized_bc_id，store_id 和 store_authorized_bc_id 为选填字段。若您通过这些字段为您的直播购物广告指定了TikTok Shop，将会报告此商店的商品销量。若您不指定 TikTok Shop，将会报告 TikTok 账号（identity_id）显示的商品销量。
	// - 若 product_source 设置为 CATALOG，需将本字段设置为您想推广的商品所在的商品库。
	// - 若 product_source 设置为 STORE，需将本字段设置为与您想推广的商品所在的TikTok Shop 绑定的商品库。
	// 您可使用/store/list/获取广告账户下可用商店对应的 catalog_id，store_id 和 store_authorized_bc_id。
	// 注意：自 2024 年 6 月 30 日起， product_source 为 STORE 时，您将无需再传入 catalog_id，因为该字段将被忽略
	CatalogID string `json:"catalog_id,omitempty"`
	// CatalogAuthorizedBcID 对该商品库有权限的商务中心的 ID
	// - 若 shopping_ads_type 设置为 VIDEO 且 product_source 设置为CATALOG 或 STORE，本字段必填。
	// - 若 shopping_ads_type 设置为 LIVE，本字段选填。
	// - 若 shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS 且 product_source 设置为 STORE，不支持本字段。
	// 如果要操作的是已经资产化的商品库，则需要传入相应的商品库 ID。
	// 注意：自 2024 年 6 月 30 日起， product_source 为 STORE 时，您将无需再传入 catalog_id，因为该字段将被忽略。
	CatalogAuthorizedBcID string `json:"catalog_authorized_bc_id,omitempty"`
	// PromotionType 推广对象类型（优化位置）。若对应的推广系列推广目标不是REACH，VIDEO_VIEWS或ENGAGEMENT，则此字段必填。
	// 枚举值详见 枚举值-推广对象类型。
	// 注意: 如果在推广系列层级的objective_type是 ENGAGEMENT，则该字段只能使用 EXTERNAL_OR_DISPLAY。
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// PromotionTargetType 推广目标线索收集的专属推广对象类型
	// 仅当推广系列层级的objective_type 设置为 LEAD_GENERATION时有效。
	// 枚举值：
	// - INSTANT_PAGE：即时表单（线索表单）。创建可快速加载的应用内 TikTok 即时表单来收集更多线索。
	// - EXTERNAL_WEBSITE：第三方网站表单。使用包含第三方网站表单的落地页，或使用会重定向至包含第三方网站表单网站的TikTok即时体验页面来收集更多线索。
	// 若本字段设置为EXTERNAL_WEBSITE且 optimization_goal 为 CLICK，无需同时传入pixel_id 和 optimization_event
	PromotionTargetType enum.PromotionTargetType `json:"promotion_target_type,omitempty"`
	// OptimizationGoal 优化目标
	// 若想获取完整的枚举值，请参阅枚举值 - 优化目标。
	// 若想获取 CBO 推广系列支持的枚举值，请参阅推广系列预算优化 - 支持的优化目标。
	//
	// 对于每个优化目标，您需同时手动指定对应的billing_event。若想了解与某一优化目标关联的计费事件，可查看优化目标关联的计费事件。
	//
	// 若要在 Pangle 版位或全球化应用组合版位的广告组中使用优化目标“安装与应用内事件数据”，需将 optimization_goal 设置为INSTALL，传入 secondary_optimization_event 合法值，且在placements的值中仅包含 PLACEMENT_PANGLE或 PLACEMENT_GLOBAL_APP_BUNDLE 或PLACEMENT_PANGLE和 PLACEMENT_GLOBAL_APP_BUNDLE。
	//
	// 注意：
	//
	// promotion_type 为LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE、LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE 或 LEAD_GEN_CLICK_TO_CALL 时，本字段默认值为 CLICK。
	// 自 2024 年 11 月 30 日起，您将无法创建或复制优化目标为“安装与应用内事件数据”且版位为仅 TikTok 或自动版位的广告组。此变动将影响使用以下配置的广告组：
	// 推广系列层级：objective_type 为 APP_PROMOTION
	// 广告组层级：
	// optimization_goal 为 INSTALL 且指定了 secondary_optimization_event 合法值
	// placement_type 为 PLACEMENT_TYPE_NORMAL 且 placements 为["PLACEMENT_TIKTOK"]，或 placement_type 为PLACEMENT_TYPE_AUTOMATIC
	// 现有的投放在 TikTok 版位的“安装与应用内事件数据”广告组将不受影响。此外，Pangle 版位和全球化应用组合版位（即 placement_type 为 PLACEMENT_TYPE_NORMAL 且 placements 的值仅包含PLACEMENT_PANGLE 和 PLACEMENT_GLOBAL_APP_BUNDLE）也不受影响。请注意此变更，并在必要时对您的集成进行适当的调整。
	// VIDEO_VIEW枚举值已废弃。当 objective_type 在推广计划层级设为 VIDEO_VIEWS 时，您可使用ENGAGED_VIEW 或 ENGAGED_VIEW_FIFTEEN（白名单功能）作为优化目标。
	// 若您要在一个已开启CBO的推广系列下创建新的广告组，需确保本字段的值与第一个广告组中的设置保持一致。
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// AppID 推广的App的ID。您可通过/app/list/获取app_id
	// 在以下使用场景中必填：
	// - objective_type 设置为 APP_PROMOTION，且 app_promotion_type设置为APP_RETARGETING（应用再营销）。
	// - objective_type 设置为 APP_PROMOTION，且 app_promotion_type设置为APP_INSTALL（应用安装）或APP_PREREGISTRATION（应用预注册），同时对应的推广系列非iOS 14专属推广系列。
	//
	// 注意: 您不能使用未在 MMP（移动监测合作伙伴）侧启用 SAN（自归因）模块的 App 创建广告组。若想确保您的 App 启用 TikTok 自归因集成，可查看如何将新应用集成到 SAN 和如何将现有应用迁移到 SAN。
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
	// PixelID Pixel ID，仅对落地页生效
	// 且需同时传入optimization_event字段。本字段支持的推广目标(objective_type)包括：WEB_CONVERSIONS, PRODUCT_SALES。如果您想追踪除LEAD_GENERATION外其他目标的事件，请在创意层级上使用 tracking_pixel_id 。
	// 您可以使用/pixel/list/ 接口获取所有Pixel ID。
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
	// OptimizationEvent 广告组转化事件。参阅转化事件 ，获取详细信息
	// 当推广对象是应用且拥有追踪 URL 或传入pixel_id时必填
	// 注意：
	// 当使用 Reach，Video View，或 Traffic 推广目标，并使用pixel_id时，不支持此字段。
	// 若您要在一个已开启CBO的推广系列下创建新的广告组，需确保本字段的值与第一个广告组中的设置保持一致。
	// 当optimization_goal为 PAGE_VISIT 时，optimization_event 将自动设置为 PAGE_VISIT。
	OptimizationEvent enum.OptimizationEvent `json:"optimization_event,omitempty"`
	// CustomConversionID The ID of the Custom Conversion used in the ad group.
	CustomConversionID string `json:"custom_conversion_id,omitempty"`
	// DeepFunnelOptimizationStatus 深层漏斗优化的状态。
	// 仅当 promotion_type 为 LEAD_GENERATION 时本字段生效 。
	// 深层漏斗优化同时优化浅层和深层漏斗事件。您可以在主要优化事件之外再选择一个次要事件，这有助于提升推广系列的效果。
	//
	// 枚举值：
	// ON ：启用深层漏斗优化。
	// OFF：不启用深层漏斗优化。
	// 默认值：OFF。
	//
	// 若想了解如何在线索广告中配置深层漏斗优化，可查看创建优化位置为即时表单的线索广告和创建优化位置为落地页的线索广告。
	//
	// 注意：
	//
	// 使用 CRM 事件组的深层漏斗优化目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 使用 Pixel 或线下事件组的深层漏斗优化已全量发布。
	DeepFunnelOptimizationStatus enum.OnOff `json:"deep_funnel_optimization_status,omitempty"`
	// DeepFunnelEventSource 事件源类型。
	// deep_funnel_optimization_status 为 ON 时本字段必填
	// 枚举值：
	// PIXEL：Pixel。
	// OFFLINE：线下事件组。
	// CRM：CRM 事件组。
	DeepFunnelEventSource enum.DeepFunnelEventSource `json:"deep_funnel_event_source,omitempty"`
	// DeepFunnelEventSourceID 事件源 ID。
	// deep_funnel_optimization_status 为 ON 时本字段必填。
	// deep_funnel_event_source 为 PIXEL 时，本字段代表 Pixel ID。
	// deep_funnel_event_source 为 OFFLINE 时，本字段代表线下事件组 ID。
	// deep_funnel_event_source为CRM时，本字段代表 CRM 事件组 ID。
	DeepFunnelEventSourceID string `json:"deep_funnel_event_source_id,omitempty"`
	// DeepFunnelOptimizationEvent 深层漏斗优化事件。
	// deep_funnel_optimization_status 为 ON 时本字段必填。
	// 若想获取 Pixel、线下事件组和 CRM 事件组支持的标准事件值，可查看 Pixel 事件类型中的“用于广告创编的事件名称”列。
	// 若想获取某个 Pixel 的优化事件，可使用 /pixel/list/ 并查看返回的 optimization_event
	DeepFunnelOptimizationEvent enum.OptimizationEvent `json:"deep_funnel_optimization_event,omitempty"`
	// IdentityID 仅当 shopping_ads_type 为 VIDEO 且 product_source 为 SHOWCASE 或 shopping_ads_type 为 LIVE 时生效且必填
	// 认证身份 ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 仅当 shopping_ads_type 为 VIDEO 且 product_source 为 SHOWCASE 或 shopping_ads_type 为 LIVE 时生效且必填
	// 认证身份类型。枚举值: AUTH_CODE (授权码认证)， TT_USER (TikTok商务用户)， BC_AUTH_TT（已添加到商务中心的TikTok用户）。参阅认证身份 获取详细信息。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。当identity_type 为 BC_AUTH_TT时必填。
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
	BidPrice float64 `json:"bid_price,omitempty"`
	// ConversionBidPrice oCPM转化出价
	ConversionBidPrice float64 `json:"conversion_bid_price,omitempty"`
	// DeepBidType 深度事件出价类型
	// 枚举值: 详见枚举值-深度事件出价类型。
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
	// RoasBid 用于价值优化的 ROAS 目标值
	RoasBid float64 `json:"roas_bid,omitempty"`
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
	// BudgetMode 广告预算类型
	// dule_type必须是SCHEDULE_START_END，您需要传入一个结束时间 (schedule_end_time)。当开启推广系列预算优化（CBO）时，广告组层级的budget_mode将被忽略。
	//
	// 当未开启CBO时， 广告组层级的budget_mode支持以下枚举值：
	// BUDGET_MODE_TOTAL：总预算。
	// BUDGET_MODE_DYNAMIC_DAILY_BUDGET：动态日预算，即一周七日的平均预算。单日实际消耗不超过设定的七日平均预算的125%。周消耗不会超过七日平均预算*7。
	// objective_type设置为TRAFFIC，APP_PROMOTION，WEB_CONVERSIONS，LEAD_GENERATIONPRODUCT_SALES（仅视频购物广告），REACH，VIDEO_VIEWS 或 ENGAGEMENT 且您不想设置总预算时，将本字段设置为BUDGET_MODE_DYNAMIC_DAILY_BUDGET。
	// 注意：为推广目标为REACH，VIDEO_VIEWS 或 ENGAGEMENT 的推广系列启用 BUDGET_MODE_DYNAMIC_DAILY_BUDGET 目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	//
	// BUDGET_MODE_DAY：日预算。
	// objective_type 设置为REACH，VIDEO_VIEWS 或 ENGAGEMENT时，推荐您使用动态日预算（BUDGET_MODE_DYNAMIC_DAILY_BUDGET）而非日预算（BUDGET_MODE_DAY）。
	//
	// 了解如何设置预算类型，请参阅预算。
	//
	// 注意：
	//
	// 如果本字段设置为BUDGET_MODE_DAY，那么schedule_type可以是SCHEDULE_START_END或SCHEDULE_FROM_NOW。
	// 如果本字段设置为BUDGET_MODE_TOTAL，那么sche
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
	// 当开启了推广系列预算优化（budget_optimize_on = TRUE）时，该字段将被忽略。
	// 详情请参阅预算。欲直接获取某一币种对应的每日预算取值范围，请参阅币种-每日预算取值范围。
	Budget float64 `json:"budget,omitempty"`
	// MinBudget Ad group minimum budget.
	// The system will aim to spend at least this amount, but it is not guaranteed
	MinBudget float64 `json:"min_budget,omitempty"`
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
	// IsHfss 广告推广对象是否是 HFSS 食品（高脂、高盐、高糖食品）请注意，欧洲国家不允许向未成年人推送 HFSS 食品广告
	// 请注意，欧洲国家不允许向未成年人推送 HFSS 食品广告
	IsHfss bool `json:"is_hfss,omitempty"`
	// IsLhfCompliance 推广内容是否符合 LHF（较不健康食品）合规要求。
	// 将 is_lhf_compliance 设置为 true，即表示您确认在英国 TikTok 上推广的任何食品或饮料均符合 2024 年较不健康食品法规以及其他所有适用法律。
	// 支持的值：true，false。
	// 默认值：false。
	// 注意：自 2026 年 1 月 1 日起，当广告定向到英国地域且 is_hfss 为 true 时，is_lhf_compliance 必填，且必须设置为 true。
	IsLhfCompliance bool `json:"is_lhf_compliance,omitempty"`
	// PlacementType 版位类型。
	// 枚举值：
	// PLACEMENT_TYPE_AUTOMATIC：自动版位。TikTok 的广告系统将使用智能计算，为您选择所有可用应用的版位最佳组合。
	// PLACEMENT_TYPE_NORMAL：编辑版位。通过手动编辑版位，你将能够自主选择广告投放的应用。
	//
	// 默认值: PLACEMENT_TYPE_NORMAL。
	//
	// 注意：
	//
	// 广告组创建后无法更新placement_type。
	// 若本字段设置为 PLACEMENT_TYPE_AUTOMATIC ，则返回的 placements 将显示实际支持的版位。例如，若实际仅支持 TikTok 版位，返回的 placements 值将为 PLACEMENT_TIKTOK。
	PlacementType enum.PlacementType `json:"placement_type,omitempty"`
	// Placements 投放版位。获取枚举值，详见枚举值-版位。
	// 当 placement_type 为 PLACEMENT_TYPE_NORMAL 时，本字段必填。
	// 当 placement_type 为 PLACEMENT_TYPE_AUTOMATIC 时，忽略本字段，系统将自动选择版位写入placements。
	// 注意：
	// 当placement_type为PLACEMENT_TYPE_AUTOMATIC（自动版位）时，本字段返回指定的推广目标下实际支持的版位。例如，若实际仅支持 TikTok版位，本字段返回的值将为PLACEMENT_TIKTOK 。
	// 若指定的推广目标支持 PLACEMENT_GLOBAL_APP_BUNDLE 版位，但您未申请全球化应用组合版位的白名单，则本字段的值将不会包含PLACEMENT_GLOBAL_APP_BUNDLE。
	// 当 placement_type 为 PLACEMENT_TYPE_NORMAL（手动版位）时，本字段通常将返回您在请求中通过placements字段指定的版位。
	// 唯一的例外是您在请求中通过 placements 字段指定了 PLACEMENT_GLOBAL_APP_BUNDLE 版位 ，但是您未申请 PLACEMENT_GLOBAL_APP_BUNDLE 的白名单，则本字段的返回值中将筛除PLACEMENT_GLOBAL_APP_BUNDLE 。例如，若您在请求中将 placements 设置为["PLACEMENT_TIKTOK", "PLACEMENT_GLOBAL_APP_BUNDLE"]，但是未申请全球化应用组合版位的白名单，本字段的返回值将为["PLACEMENT_TIKTOK"]。
	Placements []enum.Placement `json:"placements,omitempty"`
	// CommentDisabled 是否允许用户在TikTok上评论您的广告
	CommentDisabled *bool `json:"comment_disabled,omitempty"`
	// ShareDisabled 本广告组中的广告是否禁止分享到第三方平台
	// 默认值： False。仅当以下条件均满足时 True的值会生效：
	// 推广系列层级的objective_type 设置为以下推广目标之一： APP_PROMOTION，WEB_CONVERSIONS， REACH， TRAFFIC，VIDEO_VIEWS，ENGAGEMENT， LEAD_GENERATION。
	// 版位：
	// placement_type设置为 PLACEMENT_TYPE_AUTOMATIC 或
	// placement_type设置为PLACEMENT_TYPE_NORMAL且 placements = PLACEMENT_TIKTOK
	ShareDisabled bool `json:"share_disabled,omitempty"`
	// VideoDownloadDisabled 用户是否可以在 TikTok 上下载您的广告视频
	VideoDownloadDisabled *bool `json:"video_download_disabled,omitempty"`
}

// Encode implements PostRequest interface
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse Create an Upgraded Smart+ Ad Group API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Adgroup `json:"data,omitempty"`
}
