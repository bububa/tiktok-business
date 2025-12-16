package report

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Metrics 所有请求的指标数据
type Metrics struct {
	// 属性指标
	// 属性指标（如广告组名称和推广对象类型）是推广系列、广告组或广告的基本属性。属性指标仅在使用ID类型的维度时有效。
	//
	// AdvertiserName 账户名称
	// 广告主、推广系列、广告组、广告层级支持。
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// AdvertiserID 账户ID
	// 广告主、推广系列、广告组、广告层级支持。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignName 推广系列名称
	// 推广系列、广告组、广告层级支持。
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignID 推广系列ID
	// 广告组、广告层级支持。
	CampaignID string `json:"campaign_id,omitempty"`
	// ObjectiveType 推广目标
	// 推广系列、广告组、广告层级支持。
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// SplitTest 拆分对比实验状态
	// 推广系列、广告组、广告层级支持。
	SplitTest string `json:"split_test,omitempty"`
	// CampaignBudget 推广系列预算
	// 推广系列、广告组、广告层级支持。
	CampaignBudget model.Float64 `json:"campaign_budget,omitempty"`
	// CampaignDedicateType 推广系列类型
	// iOS14 专属推广系列或普通推广系列。推广系列、广告组、广告层级支持。
	CampaignDedicateType string `json:"campaign_dedicate_type,omitempty"`
	// AppPromotionType 应用推广类型
	// 推广系列、广告组、广告层级支持。枚举值：APP_INSTALL，APP_RETARGETING。
	// 仅当 objective_type为APP_PROMOTION时会返回APP_INSTALL或APP_RETARGETING的取值。其他情况下仅返回UNSET。
	AppPromotionType enum.AppPromotionType `json:"app_promotion_type,omitempty"`
	// AdgroupName 广告组名称
	// 广告组、广告层级支持。
	AdgroupName string `json:"adgroup_name,omitempty"`
	// AdgroupID 广告组ID
	// 仅广告层级支持。
	AdgroupID string `json:"adgroup_id,omitempty"`
	// VideoID 仅当material_type设置为SPC_VIDEO时返回本字段。
	// 视频 ID。
	VideoID string `json:"video_id,omitempty"`
	// FileName 仅当material_type设置为SPC_VIDEO时返回本字段。
	// 视频名称。
	FileName string `json:"file_name,omitempty"`
	// IsRemovedFromSpc 仅当 material_type 为 SPC_VIDEO， SPC_AD_TEXT 或 SPC_SPARK 时返回。
	// 视频是否已从 Smart+ 推广系列广告中移除。
	// 枚举值：
	// "1"：是。
	// "0"：否。
	IsRemovedFromSpc model.Int `json:"is_removed_from_spc,omitempty"`
	// IdentityID 仅当material_type设置为SPC_SPARK时返回本字段。
	// 认证身份 ID。
	IdentityID string `json:"identity_id,omitempty"`
	// TiktokItemID 仅当material_type设置为SPC_SPARK时返回本字段。
	// 用作 Spark Ads 的 TikTok 帖子 ID。
	TiktokItemID string `json:"tiktok_item_id,omitempty"`
	// TiktokName 仅当material_type设置为SPC_SPARK时返回本字段。
	// TikTok 账户昵称。
	// 注意：目前，Smart+ 推广系列支持 identity_type 为 AUTH_CODE 或 TT_USER的 TikTok 帖子。 关于认证身份的更多信息，请访问认证身份。
	TiktokName string `json:"tiktok_name,omitempty"`
	// CatalogName 仅当 material_type 设置为SPC_CATALOG_VIDEO 或 SPC_CATALOG_CAROUSEL 时返回本字段。
	// 商品库名称。
	CatalogName string `json:"catalog_name,omitempty"`
	// CatalogModifyTime 仅当 material_type 设置为SPC_CATALOG_VIDEO 或 SPC_CATALOG_CAROUSEL 时返回本字段。
	// 商品库最近一次修改的时间，格式为YYYY-MM-DD HH:MM:SS（UTC+0）。
	// 示例：2024-00-01 00:00:00。
	CatalogModifyTime model.DateTime `json:"catalog_modify_time,omitzero"`
	// PlacementType 版位类型
	// 广告组、广告层级支持。
	PlacementType enum.Placement `json:"placement_type,omitempty"`
	// PromotionType 推广对象类型
	// 推广对象类型。可以为App，网站或其他。广告组和广告层级，同步报表和异步报表都支持本指标。
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// OptStatus 程序化创意
	// 广告组、广告层级支持。
	OptStatus string `json:"opt_status,omitempty"`
	// AdgroupDownloadURL 下载URL/网站URL
	// 广告组、广告层级支持。
	AdgroupDownloadURL string `json:"adgroup_download_url,omitempty"`
	// ProfileImage 头像
	// 广告组、广告层级支持。
	ProfileImage string `json:"profile_image,omitempty"`
	// DPATargetAudienceType DPA广告的目标受众类型
	// DPA广告的目标受众类型。广告组和广告层级，同步报表和异步报表都支持本指标。
	DPATargetAudienceType string `json:"dpa_target_audience_type,omitempty"`
	// Budget 广告组预算
	// 广告组、广告层级支持。
	Budget model.Float64 `json:"budget,omitempty"`
	// SmartTarget 优化目标
	// 广告组、广告层级支持。
	SmartTarget enum.OptimizationGoal `json:"smart_target,omitempty"`
	// PricingCategory 即将废弃, 计费点
	// 广告组、广告层级支持。
	// 若您想获取广告的计费点（计费事件）信息，请使用指标billing_event。
	PricingCategory string `json:"pricing_category,omitempty"`
	// BillingEvent 计费点
	// 广告组、广告层级支持。
	// 示例："Clicks"，"Impression"。
	BillingEvent enum.BillingEvent `json:"billing_event,omitempty"`
	// BidStrategy 竞价策略
	// 广告组、广告层级支持。
	BidStrategy enum.BidStrategy `json:"bid_strategy,omitempty"`
	// Bid 出价
	// 广告组、广告层级支持。
	Bid model.Float64 `json:"bid,omitempty"`
	// BidSecondaryGoal 深层目标出价
	// 广告组、广告层级支持。
	BidSecondaryGoal model.Float64 `json:"bid_secondary_goal,omitempty"`
	// AeoType 应用内事件优化类型
	// 广告组、广告层级支持（广告组层级目前已支持，广告层级即将支持）。
	AoeType string `json:"aeo_type,omitempty"`
	// AdName 广告名称
	// 仅广告层级支持。
	AdName string `json:"ad_name,omitempty"`
	// AdID 广告ID
	// 广告层级支持。
	AdID string `json:"ad_id,omitempty"`
	// AdText 广告标题
	// 仅广告层级支持。
	AdText string `json:"ad_text,omitempty"`
	// CallToAction 行动引导文案
	// 仅广告层级支持。
	CallToAction string `json:"call_to_action,omitempty"`
	// AdProfileImage 头像（广告层级）
	// 广告层级支持。
	AdProfileImage string `json:"ad_profile_image,omitempty"`
	// AdURL URL（广告层级）
	// 广告层级支持。
	AdURL string `json:"ad_url,omitempty"`
	// TtAppID 推广应用ID
	// 您在创建广告组时使用的应用ID。广告组、广告层级支持，当一个广告组的推广对象类型为 App 时返回。
	TtAppID string `json:"tt_app_id,omitempty"`
	// TtAppName 推广应用名称
	// 广告组、广告层级支持，当一个广告组的推广对象类型为 App 时返回。
	TtAppName string `json:"tt_app_name,omitempty"`
	// MobileAppID 推广应用在Google Play或Apple App Store中的ID
	// 例如：App Store: https://apps.apple.com/us/app/angry-birds/id343200656; Google Play：https://play.google.com/store/apps/details?id=com.rovio.angrybirds。
	// 仅广告组、广告层级支持，当一个广告组的推广对象类型为 App 时返回。
	MobileAppID string `json:"mobile_app_id,omitempty"`
	// ImageMode 样式
	// 仅广告层级均支持。
	ImageMode enum.AdFormat `json:"ad_format,omitempty"`
	// Currency 货币
	// 货币代码，比如USD。请注意如果要使用currency作为指标，则请求中的dimensions字段必须包含adgroup_id/ad_id/campaign_id/advertiser_id。
	Currency string `json:"currency,omitempty"`
	// IsAco 广告是否为程序化创意广告或智能创意广告。若为程序化创意广告或智能创意广告，则为True。
	// 支持竞价广告组（AUCTION_ADGROUP）层级。
	IsAco bool `json:"is_aco,omitempty"`
	// IsSmartCreative 广告是否为智能创意广告。
	// 支持竞价广告（AUCTION_AD）层级。
	IsSmartCreative bool `json:"is_smart_creative,omitempty"`

	//   核心指标
	// 核心指标提供关于广告表现的基本洞察，涵盖消耗和展示量等基本数据。
	// Spend 消耗
	// 您的广告消耗总金额。
	Spend model.Float64 `json:"spend,omitempty"`
	// BilledCost 净消耗
	// 您的广告消耗总金额，不包括使用的广告赠款或优惠券。
	//
	// 注意：
	//
	// 此指标仅支持在同步基础报表中使用。
	// 此指标可能会延迟11小时，并且它在2023年9月1日后才开始有数据。
	BilledCost model.Float64 `json:"billed_cost,omitempty"`
	// CashSpend 现金消耗
	// 所选时间范围内投放广告产生的现金消耗。仅支持广告主层级和分天查询，不支持 lifetime，不支持分时查询。
	//
	// 注意：指标更新可能有 24-48 小时的延迟。
	CashSpend model.Float64 `json:"cash_spend,omitempty"`
	// VoucherSpend 赠款消耗
	// 所选时间范围内投放广告产生的赠款消耗。仅支持广告主层级和分天查询，不支持 lifetime，不支持分时查询。
	//
	// 注意：指标更新可能有 24-48 小时的延迟。
	VoucherSpend model.Float64 `json:"voucher_spend,omitempty"`
	// CPC 平均点击成本（目标页面）
	// 跳转到指定目标页面的每次点击平均成本。
	CPC model.Float64 `json:"cpc,omitempty"`
	// CPM 千次展示成本 (CPM)
	// 您为每 1,000 次展示平均花费的金额。
	CPM model.Float64 `json:"cpm,omitempty"`
	// Impressions 展示量
	// 您的广告展示的次数。
	Impressions model.Int64 `json:"impressions,omitempty"`
	// GrossImpressions 总展示数（包括无效展示）
	// 广告显示在屏幕上的次数，包括无效展示。
	GrossImpressions model.Int64 `json:"gross_impressions,omitempty"`
	// Clicks 点击量（目标页面）
	// 您的广告获得的跳转到指定目标页面的点击次数。
	Clicks model.Int64 `json:"clicks,omitempty"`
	// CTR 点击率（目标页面）
	// 促成目标页面点击的展示次数占总展示次数的百分比。
	CTR model.Float64 `json:"ctr,omitempty"`
	// Reach 覆盖人数
	// 至少看过您的广告一次的去重用户数。
	Reach model.Int64 `json:"reach,omitempty"`
	// CostPer1000Reached 覆盖千人成本
	// 触达 1,000 个去重用户的平均成本。
	CostPer1000Reached model.Float64 `json:"cost_per_1000_reached,omitempty"`
	// Frequency 频次
	// 在特定时间段内每位用户看到您的广告的平均次数。
	Frequency model.Float64 `json:"frequency,omitempty"`
	// Conversion 转化量
	// 广告达成您选择的优化事件的次数。
	Conversion model.Int64 `json:"conversion,omitempty"`
	// CostPerConversion 平均转化成本
	// 在转化上花费的平均金额。
	CostPerConversion model.Float64 `json:"cost_per_conversion,omitempty"`
	// ConversionRate 转化率（点击）
	// 您获得的成效数占广告获得的目标页面总点击量的百分比。
	// 注意：2023年10月末起，本字段将调整为基于展示次数计数（即使用与 conversion_rate_v2 相同的计算逻辑）。为保证流畅的 API 体验，避免计算逻辑变更带来的影响，建议您尽快切换到基于展示次数的 conversion_rate_v2 指标。
	ConversionRate model.Float64 `json:"conversion_rate,omitempty"`
	// ConversionRateV2 转化率（CVR）
	// 您获得的成效数占广告获得的总展示次数的百分比。
	ConversionRateV2 model.Float64 `json:"conversion_rate_v2,omitempty"`
	// RealTimeConversion
	// 按转化时间计算的转化量
	// 广告达成您选择的优化事件的次数。
	RealTimeConversion model.Int64 `json:"real_time_conversion,omitempty"`
	// RealTimeCostPerConversion 按转化时间计算的平均转化成本
	// 在转化上花费的平均金额。
	RealtimeCostPerConversion model.Float64 `json:"real_time_cost_per_conversion,omitempty"`
	// RealTimeConversionRate 实时转化率（点击）
	// 您获得的转化数占广告获得的目标页面总点击量的百分比。
	// 注意：2023年10月末起，本字段将调整为基于展示次数计数（即使用与 real_time_conversion_rate_v2 相同的计算逻辑）。为保证流畅的 API 体验，避免计算逻辑变更带来的影响，建议您尽快切换到基于展示次数的 real_time_conversion_rate_v2 指标。
	RealTimeConversionRate model.Float64 `json:"real_time_conversion_rate,omitempty"`
	// RealTimeConversionRateV2 按转化时间计算的转化率 （CVR）
	// 您获得的转化数占广告获得的总展示次数的百分比。
	RealTimeConversionRateV2 model.Float64 `json:"real_time_conversion_rate_v2,omitempty"`
	// Result 成效数
	// 广告达到您预期的推广目标和优化目标的次数。
	Result model.Int64 `json:"result,omitempty"`
	// CostPerResult 平均成效成本
	// 广告获得单次成效的平均成本。
	CostPerResult model.Float64 `json:"cost_per_result,omitempty"`
	// ResultRate 成效率
	// 产生的成效数占广告获得的总展示次数的百分比。
	ResultRate model.Float64 `json:"result_rate,omitempty"`
	// RealTimeResult 实时成效数
	// 广告达到您预期的推广目标和优化目标的次数。
	RealTimeResult model.Int64 `json:"real_time_result,omitempty"`
	// RealTimeCostPerResult 实时平均成效成本
	// 广告获得单次成效的平均成本。
	RealTimeCostPerResult model.Float64 `json:"real_time_cost_per_result,omitempty"`
	// RealTimeResultRate 实时成效率
	// 产生的成效数占广告获得的总展示次数的百分比。
	RealTimeResultRate model.Float64 `json:"real_time_result_rate,omitempty"`
	// SecondaryGoalResult 深层漏斗成效
	// 广告达成您选择的预期深层漏斗事件的次数。
	SecondaryGoalResult model.Int64 `json:"secondary_goal_result,omitempty"`
	// CostPerSecondaryGoalResult 深层漏斗成效平均成本
	// 广告获得单次深层漏斗成效的平均成本。
	CostPerSecondaryGoalResult model.Float64 `json:"cost_per_secondary_goal_result,omitempty"`
	// SecondaryGoalResultRate 深层漏斗成效率
	// 深层漏斗成效数占广告获得的总曝光数的百分比。
	SecondaryGoalResultRate model.Float64 `json:"secondary_goal_result_rate,omitempty"`
	// SkanResult 成效 (SKAN)
	SkanResult model.Int64 `json:"skan_result,omitempty"`
	// SkanCostPerResult 单次成效费用 (SKAN)
	SkanCostPerResult model.Float64 `json:"skan_cost_per_result,omitempty"`
	// SkanResultRate 成效率 (SKAN)
	SkanResultRate model.Float64 `json:"skan_result_rate,omitempty"`
	// SkanConversion 转化数 (SKAN)
	SkanConversion model.Int64 `json:"skan_conversion,omitempty"`
	// SkanCostPerConversion 转化成本 (SKAN)
	SkanCostPerConversion model.Float64 `json:"skan_cost_per_conversion,omitempty"`
	// SkanConversionRateV2 转化率（SKAN，展示）
	SkanConversionRateV2 model.Float64 `json:"skan_conversion_rate_v2,omitempty"`
	//
	//互动指标
	// 互动指标衡量受众与您的广告的互动
	// VideoPlayActions 视频播放量
	// 视频开始播放的次数。每次视频展示的播放次数会单独统计，且不包括重播次数。
	VideoPlayActions model.Int64 `json:"video_play_actions,omitempty"`
	// OrganicVideoViews 自然流量视频播放量
	OrganicVideoViews model.Int64 `json:"organic_video_views,omitempty"`
	// VideoWatched2s 视频播放 2 秒次数
	// 视频至少播放 2 秒的次数。每次视频展示的播放次数会单独统计，且不包括重播次数。
	VideoWatched2s model.Int64 `json:"video_watched_2s,omitempty"`
	// VideoWatched6s 视频播放 6 秒次数
	// 视频至少播放 6 秒的次数。每次视频展示的播放次数会单独统计，且不包括重播次数。
	VideoWatched6s model.Int64 `json:"video_watched_6s,omitempty"`
	// VideoViewsP25 播放 25% 次数
	// 视频至少播放总时长的 25% 的次数。每次视频展示的播放次数会单独统计，且不包括重播次数。
	VideoViewsP25 model.Int64 `json:"video_views_p25,omitempty"`
	// VideoViewsP50 播放 50% 次数
	// 视频至少播放总时长的 50% 的次数。每次视频展示的播放次数会单独统计，且不包括重播次数。
	VideoViewsP50 model.Int64 `json:"video_views_p50,omitempty"`
	// VideoViewsP75 播放 75% 次数
	// 视频至少播放总时长的 75% 的次数。每次视频展示的播放次数会单独统计，且不包括重播次数。
	VideoViewsP75 model.Int64 `json:"video_views_p75,omitempty"`
	// VideoViewsP100 播放完成次数
	// 您的视频完整播放的次数。每次视频展示的播放次数会单独统计，且不包括重播次数。
	VideoViewsP100 model.Int64 `json:"video_views_p100,omitempty"`
	// AverageVideoPlay 平均视频播放时长
	// 您的视频的平均播放时长，重播的播放时长也计算在内。
	AverageVideoPlay model.Float64 `json:"average_video_play,omitempty"`
	// AverageVideoPlayPerUser 用户平均播放时长
	// 每位用户播放您的视频的平均时长，重播的时长也计算在内。
	AverageVideoPlayPerUser model.Float64 `json:"average_video_play_per_user,omitempty"`
	// EngagedView 播放 6 秒次数（专注观看）
	// 您的视频至少播放 6 秒（如视频短于 6 秒，则完整播放）或在前 6 秒至少收到 1 次互动的次数。
	EngagedView model.Int64 `json:"engaged_view,omitempty"`
	// PaidEngagedView 专注观看 6 秒次数（付费观看次数）
	// 视频播放至少 6 秒（如果视频短于 6 秒，则完整播放）的次数。每次视频展示的播放次数会单独计算，且不包括重播次数。
	//
	// 注意：若想使用本指标，需在请求中的 dimensions 字段指定以下任一维度组合：
	//
	// 一个 ID 维度（advertiser_id、campaign_id、adgroup_id 或 ad_id）
	// 一个 ID 维度 + stat_time_day
	PaidEngagedView model.Int64 `json:"paid_engaged_view,omitempty"`
	// PaidEngagementEngagedView 专注观看 6 秒次数（付费互动次数）
	// 您的视频在开始观看后 6 秒至少获得 1 次正面互动（点赞、分享、关注或点击）的次数。 此指标仅提供给使用专注观看商品的广告主。对于每次视频曝光，付费互动次数会单独计算，不含重播次数。
	//
	// 注意：若想使用本指标，需在请求中的 dimensions 字段指定以下任一维度组合：
	//
	// 一个 ID 维度（advertiser_id、campaign_id、adgroup_id 或 ad_id）
	// 一个 ID 维度 + stat_time_day
	PaidEngagementEngagedView model.Int64 `json:"paid_engagement_engaged_view,omitempty"`
	// EngagedView15s 播放 15 秒次数（专注观看）
	// 您的视频至少播放 15 秒（如视频短于 15 秒，则完整播放）或在前 15 秒至少收到 1 次互动的次数。
	EngagedView15s model.Int64 `json:"engaged_view_15s,omitempty"`
	// PaidEngagedView15s 专注观看 15 秒次数（付费观看次数）
	// 观众观看您的视频至少 15 秒（如果视频短于 15 秒，则完整观看）的次数。对于每次视频曝光，观看次数会单独计算，不含重播次数。 此指标仅提供给使用专注观看商品的广告主。
	//
	// 注意：若想使用本指标，需在请求中的 dimensions 字段指定以下任一维度组合：
	//
	// 一个 ID 维度（advertiser_id、campaign_id、adgroup_id 或 ad_id）
	// 一个 ID 维度 + stat_time_day
	PaidEngagedView15s model.Int64 `json:"paid_engaged_view_15s,omitempty"`
	// PaidEngagementEngagedView15s 专注观看 15 秒次数（付费互动次数）
	// 您的视频在开始观看后 15 秒至少获得 1 次正面互动（点赞、分享、关注或点击）的次数。 此指标仅提供给使用专注观看商品的广告主。对于每次视频曝光，付费互动次数会单独计算，不含重播次数。
	//
	// 注意：若想使用本指标，需在请求中的 dimensions 字段指定以下任一维度组合：
	//
	// 一个 ID 维度（advertiser_id、campaign_id、adgroup_id 或 ad_id）
	// 一个 ID 维度 + stat_time_day
	PaidEngagementEngagedView15s model.Int64 `json:"paid_engagement_engaged_view_15s,omitempty"`
	//
	// 点击指标
	// 点击指标衡量广告产生的点击次数，包括与按钮和锚点等不同组件的互动。
	//
	// Engagements 点击数（全部）
	// 您的广告获得的点击次数。此项包括了跳转到目标页面的点击以及出于社交和互动目的进行的点击。
	Engagements model.Int64 `json:"engagements,omitempty"`
	// EngagementRate 点击率（全部）
	// 促成点击的展示次数占总展示次数的百分比。
	EngagementRate model.Float64 `json:"engagement_rate,omitempty"`
	// ClicksOnMusicDisc 音乐点击数
	// 广告展示期间，与广告关联的音乐获得的点击次数。
	ClicksOnMusicDisc model.Int64 `json:"clicks_on_music_disc,omitempty"`
	// DuetClicks "合拍"按钮点击数
	// "合拍"按钮的点击次数。
	DuetClicks model.Int64 `json:"duet_clicks,omitempty"`
	// StitchClicks "拼接"按钮点击数
	// "拼接"按钮的点击次数。
	StitchClicks model.Int64 `json:"stitch_clicks,omitempty"`
	// SoundUsageClicks "使用音乐"按钮点击数
	// "使用音乐"按钮的点击次数。
	SoundUsageClicks model.Int64 `json:"sound_usage_clicks,omitempty"`
	// AnchorClicks 锚点点击数
	// 广告锚点的点击次数。
	AnchorClicks model.Int64 `json:"anchor_clicks,omitempty"`
	// AnchorClickRate 锚点点击率
	// 锚点点击次数占广告获得的总展示次数的百分比。
	AnchorClickRate model.Float64 `json:"anchor_click_rate,omitempty"`
	// ClicksOnHashtagChallenge 话题点击数
	// 广告展示期间，与广告关联的挑战赛获得的点击次数。
	ClicksOnHashtagChallenge model.Int64 `json:"clicks_on_hashtag_challenge,omitempty"`
	//
	// 社交指标
	// 社交指标衡量与广告的社交互动，例如关注、点赞、评论和分享。
	//
	// Follows 付费关注数
	// 广告展示期间，关联主页获得的关注次数。
	Follows model.Int64 `json:"follows,omitempty"`
	// Likes 付费点赞数
	// 广告展示期间，您的广告获得的点赞次数。
	Likes model.Int64 `json:"likes,omitempty"`
	// Comments 付费评论数
	// 广告展示期间，您的广告获得的评论次数。
	Comments model.Int64 `json:"comments,omitempty"`
	// Shares 付费分享数
	// 广告展示期间，您的广告获得的分享次数。
	Shares model.Int64 `json:"shares,omitempty"`
	// ProfileVisits 付费主页访问量
	// 广告展示期间，关联主页获得的访问次数。
	ProfileVisits model.Int64 `json:"profile_visits,omitempty"`
	// ProfileVisitsRate 付费主页访问率
	// 广告带来的主页访问次数占广告获得的总展示次数的百分比。
	ProfileVisitsRate model.Float64 `json:"profile_visits_rate,omitempty"`
	// TtPlaylistVisit 播放列表页面访问量
	// 付费广告在推广期间带来的播放列表页面访问次数。
	TtPlaylistVisit model.Int64 `json:"tt_playlist_visit,omitempty"`
	// TtPlaylistVisitRate 播放列表页面访问率
	// 付费广告在推广期间带来的播放列表页面访问数占广告展示数的比率。
	TtPlaylistVisitRate model.Float64 `json:"tt_playlist_visit_rate,omitempty"`
	// EarnedPageVideoViews 获得的页面视频播放量
	// 用户在单个 TikTok 应用会话中观看您的广告后，您所设置的 TikTok 页面中的视频开始播放的次数。每次视频展示的播放次数会单独统计，且不包括重播次数。
	EarnedPageVideoViews model.Int64 `json:"earned_page_video_views,omitempty"`
	// EarnedPageVideoViewsP100 获得的页面播放完成次数
	// 用户在单个 TikTok 应用会话 中观看您的广告后，您所设置的 TikTok 页面中的视频完整播放的次数。每次视频展示的播放次数会单独统计，且不包括重播次数。
	EarnedPageVideoViewsP100 model.Int64 `json:"earned_page_video_views_p100,omitempty"`
	// EarnedPageClicks 获得的页面点击量（全部）
	// 用户在单个 TikTok 应用会话中观看您的广告后，您所设置的 TikTok 页面获得的点击次数。此项包括了出于社交和互动目的进行的点击。
	EarendPageClicks model.Int64 `json:"earned_page_clicks,omitempty"`
	//
	// 直播指标
	// 直播指标衡量直播期间观众的参与度，包括播放量和点击数等指标。
	//
	// LiveViews 直播播放量
	// 用户通过广告进入您的直播间的次数。包含重复访问。
	LiveViews model.Int64 `json:"live_views,omitempty"`
	// LiveUniqueViews 直播去重播放量
	// 至少有一次通过广告进入您的直播间的去重用户数。
	LiveUniqueViews model.Int64 `json:"live_unique_views,omitempty"`
	// LiveEffectiveViews 直播播放 10 秒次数
	// 用户通过广告进入您的直播间并至少观看 10 秒的次数。包含重复访问。
	LiveEffectiveViews model.Int64 `json:"live_effective_views,omitempty"`
	// LiveProductClicks 直播商品点击数
	// 观众在直播过程中点击商品并查看商品详情页的次数。
	LiveProductClicks model.Int64 `json:"live_product_clicks,omitempty"`
	//
	// 图片指标
	// 图片指标衡量轮播广告中图片的表现。
	//
	// SingleImageImpressions 单图展示数
	// 仅支持 image_id 维度。
	// 轮播中的图片在屏幕上显示的次数。
	SingleImageImpressions model.Int64 `json:"single_image_impressions,omitempty"`
	// SingleImageImpressionRate 单图展示率
	// 仅支持 image_id 维度。
	// 图片在屏幕上显示的次数占广告总展示量的百分比。
	SingleImageImpressionRate model.Float64 `json:"single_image_impression_rate,omitempty"`
	// SingleImageCTR 单图广告点击率（CTR）
	// 仅支持 image_id 维度。
	// 用户点击图片的次数占图片总展示量的百分比。
	SingleImageCTR model.Float64 `json:"single_image_ctr,omitempty"`
	//
	// 创新互动样式指标
	// 创新互动样式指标衡量广告中使用的创新互动样式的效果。
	//
	// InteractiveAddOnImpressions 创新互动样式展示次数
	// 您的创新互动样式展示的次数。
	InteractiveAddOnImpressions model.Int64 `json:"interactive_add_on_impressions,omitempty"`
	// InteractiveAddOnDestinationClicks 创新互动样式目标页面点击数
	// 用户点击创新互动样式跳转到指定目标页面的次数。
	InteractiveAddOnDestinationClicks model.Int64 `json:"interactive_add_on_destination_clicks,omitempty"`
	// InteractiveAddOnActivityClicks 创新互动样式操作点击数
	// 创新互动样式中的交互式操作点击次数。
	InteractiveAddOnActivityClicks model.Int64 `json:"interactive_add_on_activity_clicks,omitempty"`
	// InteractiveAddOnOptionAClicks 创新互动样式选项 A 点击数
	// 创新互动样式中选项 A 获得的点击次数。
	InteractiveAddOnOptionAClicks model.Int64 `json:"interactive_add_on_option_a_clicks,omitempty"`
	// InteractiveAddOnOptionBClicks 创新互动样式选项 B 点击数
	// 创新互动样式中选项 B 获得的点击次数。
	InteractiveAddOnOptionBClicks model.Int64 `json:"interactive_add_on_option_b_clicks,omitempty"`
	// CountdownStickerRecallClicks 倒计时贴纸提醒消息点击数
	// 用户点击倒计时贴纸提醒消息跳转到指定目标页面的次数。
	CountdownStickerRecallClicks model.Int64 `json:"countdown_sticker_recall_clicks,omitempty"`
	//
	// 即时体验指标
	// 即时体验指标衡量与查看或播放即时体验相关的表现。
	//
	// IxPageDurationAvg 即时体验页面平均浏览时长
	// 您的即时体验页面展示的平均时长。
	IxPageDurationAvg model.Float64 `json:"ix_page_duration_avg,omitempty"`
	// IxPageViewrateAvg 即时体验页面平均浏览率
	// 您的即时体验页面上的内容被浏览的平均百分比。
	IxPageViewrateAvg model.Float64 `json:"ix_page_viewrate_avg,omitempty"`
	// IxVideoViews 即时体验视频播放量
	// 您的即时体验视频被播放的次数。不包括重播。
	IxVideoViews model.Int64 `json:"ix_video_views,omitempty"`
	// IxVideoViewsP25 即时体验视频播放 25% 次数
	// 您的即时体验视频至少播放总时长 25% 的次数。不包括重播。
	IxVideoViewsP25 model.Int64 `json:"ix_video_views_p25,omitempty"`
	// IxVideoViewsP50 即时体验视频播放 50% 次数
	// 您的即时体验视频至少播放 50% 长度的次数。不包括重播。
	IxVideoViewsP50 model.Int64 `json:"ix_video_views_p50,omitempty"`
	// IxVideoViewsP75 即时体验视频播放 75% 次数
	// 您的即时体验视频至少播放总时长 75% 的次数。不包括重播。
	IxVideoViewsP75 model.Int64 `json:"ix_video_views_p75,omitempty"`
	// IxVideoViewsP100 即时体验视频播放完成次数
	// 您的即时体验视频完整播放的次数。不包括重播。
	IxVideoViewsP100 model.Int64 `json:"ix_video_views_p100,omitempty"`
	// IxAverageVideoPlay 即时体验视频平均播放时长
	// 您的即时体验视频平均每次视频浏览的播放时长，重播的时长也计算在内。
	IxAverageVideoPlay model.Int64 `json:"ix_average_video_play,omitempty"`
	//
	// 转化指标
	// 转化指标衡量不同渠道（如应用、网站和线下）中的转化事件。
	//
	// 应用指标
	// 应用指标衡量与移动应用相关的操作，例如安装和注册等。
	//
	// RealTimeAppInstall 按转化时间计算的应用安装量
	// 归因于您的广告的"应用安装"事件数。
	RealTimeAppInstall model.Int64 `json:"real_time_app_install,omitempty"`
	// RealTimeAppInstallCost 按转化时间计算的应用安装平均成本
	// 单次"应用安装"事件的平均成本。
	RealTimeAppInstallCost model.Float64 `json:"real_time_app_install_cost,omitempty"`
	// AppInstall 应用安装
	// 归因于您的广告的"应用安装"事件数。
	AppInstall model.Int64 `json:"app_install,omitempty"`
	// CostPerAppInstall 应用安装平均成本
	// 单次"应用安装"事件的平均成本。
	CostPerAppInstall model.Float64 `json:"cost_per_app_install,omitempty"`
	// Registration 去重注册次数（应用）
	// 归因于广告的去重应用注册操作次数。
	Registration model.Int64 `json:"registration,omitempty"`
	// CostPerRegistration 去重平均注册成本（应用）
	// 归因于广告的去重应用注册操作平均成本。
	CostPerRegistration model.Float64 `json:"cost_per_registration,omitempty"`
	// RegistrationRate 去重注册率（应用）
	// 归因于广告的应用注册操作占应用安装总量的百分比。
	RegistrationRate model.Float64 `json:"registration_rate,omitempty"`
	// TotalRegistration 注册次数（应用）
	// 归因于广告的应用注册操作次数。
	TotalRegistration model.Int64 `json:"total_registration,omitempty"`
	// CostPerTotalRegistration 平均注册成本（应用）(%)
	// 归因于广告的应用注册操作平均成本。
	CostPerTotalRegistration model.Float64 `json:"cost_per_total_registration,omitempty"`
	// Purchase 去重付费次数（应用）
	// 归因于广告的去重应用内付费操作次数。
	Purchase model.Int64 `json:"purchase,omitempty"`
	// CostPerPurchase 去重平均付费成本（应用）
	// 归因于广告的去重应用内付费操作平均成本。
	CostPerPurchase model.Float64 `json:"cost_per_purchase,omitempty"`
	// PurchaseRate 去重付费率（应用）
	// 归因于广告的应用内付费操作占应用安装总量的百分比。
	PurchaseRate model.Float64 `json:"purchase_rate,omitempty"`
	// TotalPurchase 付费次数（应用）
	// 归因于广告的应用内付费操作次数。
	TotalPurchase model.Int64 `json:"total_purchase,omitempty"`
	// CostPerTotalPurchase 平均付费成本（应用）
	// 归因于广告的应用内付费操作平均成本。
	CostPerTotalPurchase model.Float64 `json:"cost_per_total_purchase,omitempty"`
	// ValuePerTotalPurchase 平均付费价值（应用）
	// 归因于广告的应用内付费操作平均价值。
	ValuePerTotalPurchase model.Float64 `json:"value_per_total_purchase,omitempty"`
	// TotalPurchaseValue 付费价值（应用）
	// 归因于广告的应用内付费操作总价值。
	TotalPurchaseValue model.Float64 `json:"total_purchase_value,omitempty"`
	// TotalActivePayRoas 付费 ROAS（应用）
	// 应用付费操作带来的广告支出回报率 (ROAS)。
	TotalActivePayRoas model.Float64 `json:"total_active_pay_roas,omitempty"`
	// TotalPurchaseRoasDay0 第 0 天付费 ROAS
	// 应用安装后 24 小时内的付费操作带来的总广告支出回报率 (ROAS)。
	TotalPurchaseRoasDay0 model.Float64 `json:"total_purchase_roas_day0,omitempty"`
	// TotalPurchaseRoasDay2 第 2 天付费 ROAS
	// 应用安装后 72 小时内的付费操作带来的总广告支出回报率 (ROAS)。
	TotalPurchaseRoasDay2 model.Float64 `json:"total_purchase_roas_day2,omitempty"`
	// TotalPurchaseRoasDay6 第 6 天付费 ROAS
	// 应用安装后 168 小时内的付费操作带来的总广告支出回报率 (ROAS)。
	TotalPurchaseRoasDay6 model.Float64 `json:"total_purchase_roas_day6,omitempty"`
	// TotalPurchaseValueDay0 第 0 天付费价值
	// 应用安装后 24 小时内的所有付费操作的总价值。
	TotalPurchaseValueDay0 model.Float64 `json:"total_purchase_value_day0,omitempty"`
	// TotalPurchaseValueDay2 第 2 天付费价值
	// 应用安装后 72 小时内的所有付费操作的总价值。
	TotalPurchaseValueDay2 model.Float64 `json:"total_purchase_value_day2,omitempty"`
	// TotalPurchaseValueDay6 第 6 天付费价值
	// 应用安装后 168 小时内的所有付费操作的总价值。
	TotalPurchaseValueDay6 model.Float64 `json:"total_purchase_value_day6,omitempty"`
	// AppEventAddToCart 去重加入购物车次数（应用）
	// 归因于广告的去重应用内加入购物车操作次数。
	AppEventAddToCart model.Int64 `json:"app_event_add_to_cart,omitempty"`
	// CostPerAppEventAddToCart 去重加入购物车平均成本（应用）
	// 归因于广告的去重应用内加入购物车操作平均成本。
	CostPerAppEventAddToCart model.Float64 `json:"cost_per_app_event_add_to_cart,omitempty"`
	// AppEventAddToCartRate 去重加入购物车率（应用）
	// 归因于广告的应用内加入购物车操作占应用安装总量的百分比。
	AppEventAddToCartRate model.Float64 `json:"app_event_add_to_cart_rate,omitempty"`
	// TotalAppEventAddToCart 加入购物车次数（应用）
	// 归因于广告的应用内加入购物车操作次数。
	TotalAppEventAddToCart model.Int64 `json:"total_app_event_add_to_cart,omitempty"`
	// CostPerTotalAppEventAddToCart 加入购物车平均成本（应用）
	// 归因于广告的应用内加入购物车操作平均成本。
	CostPerTotalAppEventAddToCart model.Float64 `json:"cost_per_total_app_event_add_to_cart,omitempty"`
	// ValuePerTotalAppEventAddToCart 加入购物车平均价值（应用）
	// 归因于广告的应用内加入购物车操作平均价值。
	ValuePerTotalAppEventAddToCart model.Float64 `json:"value_per_total_app_event_add_to_cart,omitempty"`
	// TotalAppEventAddToCartValue 加入购物车价值（应用）
	// 归因于广告的应用内加入购物车操作总价值。
	TotalAppEventAddToCartValue model.Float64 `json:"total_app_event_add_to_cart_value,omitempty"`
	// Checkout 去重结账次数（应用）
	// 归因于广告的去重应用内结账操作次数。
	Checkout model.Int64 `json:"checkout,omitempty"`
	// CostPerCheckout 去重平均结账成本（应用）
	// 归因于广告的去重应用内结账操作平均成本。
	CostPerCheckout model.Float64 `json:"cost_per_checkout,omitempty"`
	// CheckoutRate 去重结账率（应用）
	// 归因于广告的应用内结账操作占应用安装总量的百分比。
	CheckoutRate model.Float64 `json:"checkout_rate,omitempty"`
	// TotalCheckout 结账次数（应用）
	// 归因于广告的应用内结账操作次数。
	TotalCheckout model.Int64 `json:"total_checkout,omitempty"`
	// CostPerTotalCheckout 平均结账成本（应用）
	// 归因于广告的应用内结账操作平均成本。
	CostPerTotalCheckout model.Float64 `json:"cost_per_total_checkout,omitempty"`
	// ValuePerCheckout 平均结账价值（应用）
	// 归因于广告的应用内结账操作平均价值。
	ValuePerCheckout model.Float64 `json:"value_per_checkout,omitempty"`
	// TotalCheckoutValue 结账价值（应用）
	// 归因于广告的应用内结账操作总价值。
	TotalCheckoutValue model.Float64 `json:"total_checkout_value,omitempty"`
	// ViewContent 去重查看内容次数（应用）
	// 归因于广告的去重应用内查看内容操作次数。
	ViewContent model.Int64 `json:"view_content,omitempty"`
	// CostPerViewContent 去重查看内容平均成本（应用）
	// 归因于广告的去重应用内查看内容操作平均成本。
	CostPerViewContent model.Float64 `json:"cost_per_view_content,omitempty"`
	// ViewContentRate 去重查看内容率（应用）
	// 归因于广告的应用内查看内容操作占应用安装总量的百分比。
	ViewContentRate model.Float64 `json:"view_content_rate,omitempty"`
	// TotalViewContent 查看内容次数（应用）
	// 归因于广告的应用内查看内容操作次数。
	TotalViewContent model.Int64 `json:"total_view_content,omitempty"`
	// CostPerTotalViewContent 查看内容平均成本（应用）
	// 归因于广告的应用内查看内容操作平均成本。
	CostPerTotalViewContent model.Float64 `json:"cost_per_total_view_content,omitempty"`
	// ValuePerTotalViewContent 查看内容平均价值（应用）
	// 归因于广告的应用内查看内容操作平均价值。
	ValuePerTotalViewContent model.Float64 `json:"value_per_total_view_content,omitempty"`
	// TotalViewContentValue 查看内容价值（应用）
	// 归因于广告的应用内查看内容操作总价值。
	TotalViewContentValue model.Float64 `json:"total_view_content_value,omitempty"`
	// NextDayOpen 去重次日留存数
	// 归因于广告的去重次日留存数。
	NextDayOpen model.Int64 `json:"next_day_open,omitempty"`
	// CostPerNextDayOpen 去重次日留存平均成本
	// 去重次日留存平均成本
	CostPerNextDayOpen model.Float64 `json:"cost_per_next_day_open,omitempty"`
	// NextDayOpenRate 次日留存率 (%)
	// 去重次日留存数占所有应用安装事件数的百分比。
	NextDayOpenRate model.Float64 `json:"next_day_open_rate,omitempty"`
	// TotalNextDayOpen 次日留存数
	// 归因于广告的次日留存数。
	TotalNextDayOpen model.Int64 `json:"total_next_day_open,omitempty"`
	// CostPerTotalNextDayOpen 次日留存平均成本
	// 次日留存的平均成本。
	CostPerTotalNextDayOpen model.Int64 `json:"cost_per_total_next_day_open,omitempty"`
	// Day7Retention 去重 7 日留存数
	// 来自移动应用、作为应用事件统计并归因于广告的去重 7 日留存数。
	Day7Retention model.Int64 `json:"day7_retention,omitempty"`
	// CostPerDay7Retention 去重 7 日留存平均成本
	// 归因于广告的去重 7 日留存的平均成本。
	CostPerDay7Retention model.Float64 `json:"cost_per_day7_retention,omitempty"`
	// Day7RetentionRate 7 日留存率 (%)
	// 7 日留存数占所有应用安装事件数的百分比。
	Day7RetentionRate model.Float64 `json:"day7_retention_rate,omitempty"`
	// TotalDay7Retention 7 日留存数
	// 来自移动应用、作为应用事件统计并归因于广告的 7 日留存总数。
	TotalDay7Retention model.Int64 `json:"total_day7_retention,omitempty"`
	// CostPerTotalDay7Retention 7 日留存平均成本
	// 归因于广告的 7 日留存的平均成本。
	CostPerTotalDay7Retention model.Float64 `json:"cost_per_total_day7_retention,omitempty"`
	// ValuePerTotalDay7Retention 7 日留存平均价值
	// 归因于广告的 7 日留存的平均价值。
	ValuePerTotalDay7Retention model.Float64 `json:"value_per_total_day7_retention,omitempty"`
	// TotalDay7RetentionValue 7 日留存价值
	// 所有 7 日留存返回的总价值。
	TotalDay7RetentionValue model.Float64 `json:"total_day7_retention_value,omitempty"`
	// AddPaymentInfo 去重添加支付信息次数（应用）
	// 归因于广告的去重应用内添加支付信息操作次数。
	AddPaymentInfo model.Int64 `json:"add_payment_info,omitempty"`
	// CostPerAddPaymentInfo 去重添加支付信息平均成本（应用）
	// 归因于广告的去重应用内添加支付信息操作平均成本。
	CostPerAddPaymentInfo model.Float64 `json:"cost_per_add_payment_info,omitempty"`
	// AddPaymentInfoRate 去重添加支付信息率（应用）
	// 归因于广告的应用内添加支付信息操作占应用安装总量的百分比。
	AddPaymentInfoRate model.Float64 `json:"add_payment_info_rate,omitempty"`
	// TotalAddPaymentInfo 添加支付信息次数（应用）
	// 归因于广告的应用内添加支付信息操作次数。
	TotalAddPaymentInfo model.Int64 `json:"total_add_payment_info,omitempty"`
	// CostTotalAddPaymentInfo 添加支付信息平均成本（应用）
	// 归因于广告的应用内添加支付信息操作平均成本。
	CostTotalAddPaymentInfo model.Float64 `json:"cost_total_add_payment_info,omitempty"`
	// AddToWishlist 去重加入心愿单次数（应用）
	AddToWishlist model.Int64 `json:"add_to_wishlist,omitempty"`
	// CostPerAddToWishlist 去重加入心愿单平均成本（应用）
	CostPerAddToWishlist model.Float64 `json:"cost_per_add_to_wishlist,omitempty"`
	// AddToWishlistRate 去重加入心愿单率（应用）
	AddToWishlistRate model.Float64 `json:"add_to_wishlist_rate,omitempty"`
	// TotalAddToWishlist 加入心愿单次数（应用）
	TotalAddToWishlist model.Int64 `json:"total_add_to_wishlist,omitempty"`
	// CostPerTotalAddToWishlist 加入心愿单平均成本（应用）
	CostPerTotalAddToWishlist model.Float64 `json:"cost_per_total_add_to_wishlist,omitempty"`
	// ValuePerTotalAddToWishlist 加入心愿单平均价值（应用）
	ValuePerTotalAddToWishlist model.Float64 `json:"value_per_total_add_to_wishlist,omitempty"`
	// TotalAddToWishlistValue 加入心愿单价值（应用）
	TotalAddToWishlistValue model.Float64 `json:"total_add_to_wishlist_value,omitempty"`
	// LaunchApp 去重打开应用次数（应用）
	// 归因于广告的去重打开应用操作次数。
	LaunchApp model.Int64 `json:"launch_app,omitempty"`
	// CostPerLaunchApp 去重打开应用平均成本（应用）
	// 归因于广告的去重打开应用操作平均成本。
	CostPerLaunchApp model.Float64 `json:"cost_per_launch_app,omitempty"`
	// LaunchAppRate 去重打开应用率（应用）
	// 归因于广告的打开应用操作占应用安装总量的百分比。
	LaunchAppRate model.Float64 `json:"launch_app_rate,omitempty"`
	// TotalLaunchApp 打开应用次数（应用）
	// 归因于广告的打开应用操作次数。
	TotalLaunchApp model.Int64 `json:"total_launch_app,omitempty"`
	// CostPerTotalLaunchApp 打开应用平均成本（应用）
	// 归因于广告的打开应用操作平均成本。
	CostPerTotalLaunchApp model.Float64 `json:"cost_per_total_launch_app,omitempty"`
	// CompleteTutorial 去重完成教程次数（应用）
	CompleteTutorial model.Int64 `json:"complete_tutorial,omitempty"`
	// CostPerCompleteTutorial 去重完成教程平均成本（应用）
	CostPerCompleteTutorial model.Float64 `json:"cost_per_complete_tutorial,omitempty"`
	// CompleteTutorialRate 去重完成教程率（应用）
	CompleteTutorialRate model.Float64 `json:"complete_tutorial_rate,omitempty"`
	// TotalCompleteTutorial 完成教程次数（应用）
	TotalCompleteTutorial model.Int64 `json:"total_complete_tutorial,omitempty"`
	// CostPerTotalCompleteTutorial 完成教程平均成本（应用）
	CostPerTotalCompleteTutorial model.Float64 `json:"cost_per_total_complete_tutorial,omitempty"`
	// ValuePerTotalCompleteTutorial 完成教程平均价值（应用）
	ValuePerTotalCompleteTutorial model.Float64 `json:"value_per_total_complete_tutorial,omitempty"`
	// TotalCompleteTutorialValue 完成教程价值（应用）
	TotalCompleteTutorialValue model.Float64 `json:"total_complete_tutorial_value,omitempty"`
	// CreateGroup 去重创建战队次数（应用）
	CreateGroup model.Int64 `json:"create_group,omitempty"`
	// CostPerCreateGroup 去重创建战队平均成本（应用）
	CostPerCreateGroup model.Float64 `json:"cost_per_create_group,omitempty"`
	// CreateGroupRate 去重创建战队率（应用）
	CreateGroupRate model.Float64 `json:"create_group_rate,omitempty"`
	// TotalCreateGroup 创建战队次数（应用）
	TotalCreateGroup model.Int64 `json:"total_create_group,omitempty"`
	// CostPerTotalCreateGroup 创建战队平均成本（应用）
	CostPerTotalCreateGroup model.Float64 `json:"cost_per_total_create_group,omitempty"`
	// ValuePerTotalCreateGroup 创建战队平均价值（应用）
	ValuePerTotalCreateGroup model.Float64 `json:"value_per_total_create_group,omitempty"`
	// TotalCreateGroupValue 创建战队价值（应用）
	TotalCreateGroupValue model.Float64 `json:"total_create_group_value,omitempty"`
	// JoinGroup 去重加入战队次数（应用）
	JoinGroup model.Int64 `json:"join_group,omitempty"`
	// CostPerJoinGroup 去重加入战队平均成本（应用）
	CostPerJoinGroup model.Float64 `json:"cost_per_join_group,omitempty"`
	// JoinGroupRate 去重加入战队率（应用）
	JoinGroupRate model.Float64 `json:"join_group_rate,omitempty"`
	// TotalJoinGroup 加入战队次数（应用）
	TotalJoinGroup model.Int64 `json:"total_join_group,omitempty"`
	// CostPerTotalJoinGroup 加入战队平均成本（应用）
	CostPerTotalJoinGroup model.Float64 `json:"cost_per_total_join_group,omitempty"`
	// ValuePerTotalJoinGroup 加入战队平均价值（应用）
	ValuePerTotalJoinGroup model.Float64 `json:"value_per_total_join_group,omitempty"`
	// TotalJoinGroupValue 加入战队价值（应用）
	TotalJoinGroupValue model.Float64 `json:"total_join_group_value,omitempty"`
	// CreateGamerole 去重创建角色次数（应用）
	CreateGamerole model.Int64 `json:"create_gamerole,omitempty"`
	// CostPerCreateGamerole 去重创建角色平均成本（应用）
	CostPerCreateGamerole model.Float64 `json:"cost_per_create_gamerole,omitempty"`
	// CreateGameroleRate 去重创建角色率（应用）
	CreateGameroleRate model.Float64 `json:"create_gamerole_rate,omitempty"`
	// TotalCreateGamerole 创建角色次数（应用）
	TotalCreateGamerole model.Int64 `json:"total_create_gamerole,omitempty"`
	// CostPerTotalCreateGamerole 创建角色平均成本（应用）
	CostPerTotalCreateGamerole model.Float64 `json:"cost_per_total_create_gamerole,omitempty"`
	// ValuePerTotalCreateGamerole 创建角色平均价值（应用）
	ValuePerTotalCreateGamerole model.Float64 `json:"value_per_total_create_gamerole,omitempty"`
	// TotalCreateGameroleValue 创建角色价值（应用）
	TotalCreateGameroleValue model.Float64 `json:"total_create_gamerole_value,omitempty"`
	// SpendCredits 去重花费点数次数（应用）
	SpendCredits model.Int64 `json:"spend_credits,omitempty"`
	// CostPerSpendCredits 去重花费点数平均成本（应用）
	CostPerSpendCredits model.Float64 `json:"cost_per_spend_credits,omitempty"`
	// SpendCreditsRate 去重花费点数率（应用）
	SpendCreditsRate model.Float64 `json:"spend_credits_rate,omitempty"`
	// TotalSpendCredits 花费点数次数（应用）
	TotalSpendCredits model.Int64 `json:"total_spend_credits,omitempty"`
	// CostPerTotalSpendCredits 花费点数平均成本（应用）
	CostPerTotalSpendCredits model.Float64 `json:"cost_per_total_spend_credits,omitempty"`
	// ValuePerTotalSpendCredits 花费点数平均价值（应用）
	ValuePerTotalSpendCredits model.Float64 `json:"value_per_total_spend_credits,omitempty"`
	// TotalSpendCreditsValue 花费点数价值（应用）
	TotalSpendCreditsValue model.Float64 `json:"total_spend_credits_value,omitempty"`
	// AchieveLevel 去重达到等级次数（应用）
	AchieveLevel model.Int64 `json:"achieve_level,omitempty"`
	// CostPerAchieveLevel 去重达到等级平均成本（应用）
	CostPerAchieveLevel model.Float64 `json:"cost_per_achieve_level,omitempty"`
	// AchieveLevelRate 去重达到等级率（应用）
	AchieveLevelRate model.Float64 `json:"achieve_level_rate,omitempty"`
	// TotalAchieveLevel 达到等级次数（应用）
	TotalAchieveLevel model.Int64 `json:"total_achieve_level,omitempty"`
	// CostPerTotalAchieveLevel 达到等级平均成本（应用）
	CostPerTotalAchieveLevel model.Float64 `json:"cost_per_total_achieve_level,omitempty"`
	// ValuePerTotalAchieveLevel 达到等级平均价值（应用）
	ValuePerTotalAchieveLevel model.Float64 `json:"value_per_total_achieve_level,omitempty"`
	// TotalAchieveLevelValue 达到等级价值（应用）
	TotalAchieveLevelValue model.Float64 `json:"total_achieve_level_value,omitempty"`
	// UnlockAchievement 去重解锁关卡次数（应用）
	UnlockAchievement model.Int64 `json:"unlock_achievement,omitempty"`
	// CostPerUnlockAchievement 去重解锁关卡平均成本（应用）
	CostPerUnlockAchievement model.Float64 `json:"cost_per_unlock_achievement,omitempty"`
	// UnlockAchievementRate 去重解锁关卡率（应用）
	UnlockAchievementRate model.Float64 `json:"unlock_achievement_rate,omitempty"`
	// TotalUnlockAchievement 解锁关卡次数（应用）
	TotalUnlockAchievement model.Int64 `json:"total_unlock_achievement,omitempty"`
	// CostPerTotalUnlockAchievement 解锁关卡平均成本（应用）
	CostPerTotalUnlockAchievement model.Float64 `json:"cost_per_total_unlock_achievement,omitempty"`
	// ValuePerTotalUnlockAchievement 解锁关卡平均价值（应用）
	ValuePerTotalUnlockAchievement model.Float64 `json:"value_per_total_unlock_achievement,omitempty"`
	// TotalUnlockAchievementValue 解锁关卡价值（应用）
	TotalUnlockAchievementValue model.Float64 `json:"total_unlock_achievement_value,omitempty"`
	// SalesLead 去重收集线索次数（应用）
	SalesLead model.Int64 `json:"sales_lead,omitempty"`
	// CostPerSalesLead 去重收集线索平均成本（应用）
	CostPerSalesLead model.Float64 `json:"cost_per_sales_lead,omitempty"`
	// SalesLeadRate 去重收集线索率（应用）
	SalesLeadRate model.Float64 `json:"sales_lead_rate,omitempty"`
	// TotalSalesLead 收集线索次数（应用）
	TotalSalesLead model.Int64 `json:"total_sales_lead,omitempty"`
	// CostPerTotalSalesLead 收集线索平均成本（应用）
	CostPerTotalSalesLead model.Float64 `json:"cost_per_total_sales_lead,omitempty"`
	// ValuePerTotalSalesLead 收集线索平均价值（应用）
	ValuePerTotalSalesLead model.Float64 `json:"value_per_total_sales_lead,omitempty"`
	// TotalSalesLeadValue 收集线索价值（应用）
	TotalSalesLeadValue model.Float64 `json:"total_sales_lead_value,omitempty"`
	// InAppAdClick 去重应用内广告点击次数（应用）
	InAppAdClick model.Int64 `json:"in_app_ad_click,omitempty"`
	// CostPerInAppAdClick 去重应用内广告点击平均成本（应用）
	CostPerInAppAdClick model.Float64 `json:"cost_per_in_app_ad_click,omitempty"`
	// InAppAdClickRate 去重应用内广告点击率（应用）
	InAppAdClickRate model.Float64 `json:"in_app_ad_click_rate,omitempty"`
	// TotalInAppAdClick 应用内广告点击次数（应用）
	TotalInAppAdClick model.Int64 `json:"total_in_app_ad_click,omitempty"`
	// CostPerTotalInAppAdClick 应用内广告点击平均成本（应用）
	CostPerTotalInAppAdClick model.Float64 `json:"cost_per_total_in_app_ad_click,omitempty"`
	// ValuePerTotalInAppAdClick 应用内广告点击平均价值（应用）
	ValuePerTotalInAppAdClick model.Float64 `json:"value_per_total_in_app_ad_click,omitempty"`
	// TotalInAppAdClickValue 应用内广告点击价值（应用）
	TotalInAppAdClickValue model.Float64 `json:"total_in_app_ad_click_value,omitempty"`
	// InAppAdImpr 去重应用内广告展示次数（应用）
	InAppAdImpr model.Int64 `json:"in_app_ad_impr,omitempty"`
	// CostPerInAppAdImpr 去重应用内广告展示平均成本（应用）
	CostPerInAppAdImpr model.Float64 `json:"cost_per_in_app_ad_impr,omitempty"`
	// InAppAdImprRate 去重应用内广告展示率（应用）
	InAppAdImprRate model.Float64 `json:"in_app_ad_impr_rate,omitempty"`
	// TotalInAppAdImpr 应用内广告展示次数（应用）
	TotalInAppAdImpr model.Int64 `json:"total_in_app_ad_impr,omitempty"`
	// CostPerTotalInAppAdImpr 应用内广告展示平均成本（应用）
	CostPerTotalInAppAdImpr model.Float64 `json:"cost_per_total_in_app_ad_impr,omitempty"`
	// ValuePerTotalInAppAdImpr 应用内广告展示平均价值（应用）
	ValuePerTotalInAppAdImpr model.Float64 `json:"value_per_total_in_app_ad_impr,omitempty"`
	// TotalInAppAdImprValue 应用内广告展示价值（应用）
	TotalInAppAdImprValue model.Float64 `json:"total_in_app_ad_impr_value,omitempty"`
	// LoanApply 去重申请贷款次数（应用）
	LoanApply model.Int64 `json:"loan_apply,omitempty"`
	// CostPerLoanApply 去重申请贷款平均成本（应用）
	CostPerLoanApply model.Float64 `json:"cost_per_loan_apply,omitempty"`
	// LoanApplyRate 去重申请贷款率（应用）
	LoanApplyRate model.Float64 `json:"loan_apply_rate,omitempty"`
	// TotalLoanApply 申请贷款次数（应用）
	TotalLoanApply model.Int64 `json:"total_loan_apply,omitempty"`
	// CostPerTotalLoanApply 申请贷款平均成本（应用）
	CostPerTotalLoanApply model.Float64 `json:"cost_per_total_loan_apply,omitempty"`
	// LoanCredit 去重批准贷款次数（应用）
	LoanCredit model.Int64 `json:"loan_credit,omitempty"`
	// CostPerLoanCredit 去重批准贷款平均成本（应用）
	CostPerLoanCredit model.Float64 `json:"cost_per_loan_credit,omitempty"`
	// LoanCreditRate 去重批准贷款率（应用）
	LoanCreditRate model.Float64 `json:"loan_credit_rate,omitempty"`
	// TotalLoanCredit 批准贷款次数（应用）
	TotalLoanCredit model.Int64 `json:"total_loan_credit,omitempty"`
	// CostPerTotalLoanCredit 批准贷款平均成本（应用）
	CostPerTotalLoanCredit model.Float64 `json:"cost_per_total_loan_credit,omitempty"`
	// LoanDisbursement 去重发放贷款次数（应用）
	LoanDisbursement model.Int64 `json:"loan_disbursement,omitempty"`
	// CostPerLoanDisbursement 去重发放贷款平均成本（应用）
	CostPerLoanDisbursement model.Float64 `json:"cost_per_loan_disbursement,omitempty"`
	// LoanDisbursementRate 去重发放贷款率（应用）
	LoanDisbursementRate model.Float64 `json:"loan_disbursement_rate,omitempty"`
	// TotalLoanDisbursement 发放贷款次数（应用）
	TotalLoanDisbursement model.Int64 `json:"total_loan_disbursement,omitempty"`
	// CostPerTotalLoanDisbursement 发放贷款平均成本（应用）
	CostPerTotalLoanDisbursement model.Float64 `json:"cost_per_total_loan_disbursement,omitempty"`
	// Login 去重登录次数（应用）
	Login model.Int64 `json:"login,omitempty"`
	// CostPerLogin 去重登录平均成本（应用）
	CostPerLogin model.Float64 `json:"cost_per_login,omitempty"`
	// LoginRate 去重登录率（应用）
	LoginRate model.Float64 `json:"login_rate,omitempty"`
	// TotalLogin 登录次数（应用）
	TotalLogin model.Int64 `json:"total_login,omitempty"`
	// CostPerTotalLogin 登录平均成本（应用）
	CostPerTotalLogin model.Float64 `json:"cost_per_total_login,omitempty"`
	// Ratings 去重评分次数（应用）
	Ratings model.Int64 `json:"ratings,omitempty"`
	// CostPerRatings 去重评分平均成本（应用）
	CostPerRatings model.Float64 `json:"cost_per_ratings,omitempty"`
	// RatingsRate 去重评分率（应用）
	RatingsRate model.Float64 `json:"ratings_rate,omitempty"`
	// TotalRatings 评分次数（应用）
	TotalRatings model.Int64 `json:"total_ratings,omitempty"`
	// CostPerTotalRatings 评分平均成本（应用）
	CostPerTotalRatings model.Float64 `json:"cost_per_total_ratings,omitempty"`
	// ValuePerTotalRatings 评分平均价值（应用）
	ValuePerTotalRatings model.Float64 `json:"value_per_total_ratings,omitempty"`
	// TotalRatingsValue 评分价值（应用）
	TotalRatingsValue model.Float64 `json:"total_ratings_value,omitempty"`
	// Search 去重搜索次数（应用）
	Search model.Int64 `json:"search,omitempty"`
	// CostPerSearch 去重搜索平均成本（应用）
	CostPerSearch model.Float64 `json:"cost_per_search,omitempty"`
	// SearchRate 去重搜索率（应用）
	SearchRate model.Float64 `json:"search_rate,omitempty"`
	// TotalSearch 搜索次数（应用）
	TotalSearch model.Int64 `json:"total_search,omitempty"`
	// CostPerTotalSearch 搜索平均成本（应用）
	CostPerTotalSearch model.Float64 `json:"cost_per_total_search,omitempty"`
	// StartTrial 去重开始试用次数（应用）
	StartTrial model.Int64 `json:"start_trial,omitempty"`
	// CostPerStartTrial 去重开始试用平均成本（应用）
	CostPerStartTrial model.Float64 `json:"cost_per_start_trial,omitempty"`
	// StartTrialRate 去重开始试用率（应用）
	StartTrialRate model.Float64 `json:"start_trial_rate,omitempty"`
	// TotalStartTrial 开始试用次数（应用）
	TotalStartTrial model.Int64 `json:"total_start_trial,omitempty"`
	// CostPerTotalStartTrial 开始试用平均成本（应用）
	CostPerTotalStartTrial model.Float64 `json:"cost_per_total_start_trial,omitempty"`
	// Subscribe 去重订阅次数（应用）
	Subscribe model.Int64 `json:"subscribe,omitempty"`
	// CostPerSubscribe 去重订阅平均成本（应用）
	CostPerSubscribe model.Float64 `json:"cost_per_subscribe,omitempty"`
	// SubscribeRate 去重订阅率（应用）
	SubscribeRate model.Float64 `json:"subscribe_rate,omitempty"`
	// TotalSubscribe 订阅次数（应用）
	TotalSubscribe model.Int64 `json:"total_subscribe,omitempty"`
	// CostPerTotalSubscribe 订阅平均成本（应用）
	CostPerTotalSubscribe model.Float64 `json:"cost_per_total_subscribe,omitempty"`
	// ValuePerTotalSubscribe 订阅平均价值（应用）
	ValuePerTotalSubscribe model.Float64 `json:"value_per_total_subscribe,omitempty"`
	// TotalSubscribeValue 订阅价值（应用）
	TotalSubscribeValue model.Float64 `json:"total_subscribe_value,omitempty"`
	// UniqueCustomAppEvents 去重自定义事件次数（应用）
	UniqueCustomAppEvents model.Int64 `json:"unique_custom_app_events,omitempty"`
	// CostPerUniqueCustomAppEvent 去重自定义事件平均成本（应用）
	CostPerUniqueCustomAppEvent model.Float64 `json:"cost_per_unique_custom_app_event,omitempty"`
	// CustomAppEventRate 自定义事件率（应用）
	CustomAppEventRate model.Float64 `json:"custom_app_event_rate,omitempty"`
	// CustomAppEvents 自定义事件次数（应用）
	CustomAppEvents model.Int64 `json:"custom_app_events,omitempty"`
	// CostPerCustomAppEvent 自定义事件平均成本（应用）
	CostPerCustomAppEvent model.Float64 `json:"cost_per_custom_app_event,omitempty"`
	// ValuePerCustomAppEvent 自定义事件平均价值（应用）
	ValuePerCustomAppEvent model.Float64 `json:"value_per_custom_app_event,omitempty"`
	// CustomAppEventsValue 自定义事件价值（应用）
	CustomAppEventsValue model.Float64 `json:"custom_app_events_value,omitempty"`
	// UniqueAdImpressionEvents 去重广告展示数
	UniqueAdImpressionEvents model.Int64 `json:"unique_ad_impression_events,omitempty"`
	// CostPerUniqueAdImpressionEvent 去重广告展示平均成本
	CostPerUniqueAdImpressionEvent model.Float64 `json:"cost_per_unique_ad_impression_event,omitempty"`
	// CustomizedAdImpressionEventRate 去重广告展示率（%）
	CustomizedAdImpressionEventRate model.Float64 `json:"customized_ad_impression_event_rate,omitempty"`
	// AdsImpressionEvents 广告展示数
	AdsImpressionEvents model.Int64 `json:"ads_impression_events,omitempty"`
	// CostPerAdImpressionEvent 广告展示平均成本
	CostPerAdImpressionEvent model.Float64 `json:"cost_pre_ad_impression_event,omitempty"`
	// ValuePerAdImpressionEvent 广告展示平均价值
	ValuePerAdImpressionEvent model.Float64 `json:"value_per_ad_impression_event,omitempty"`
	// TotalAdImpressionEventsValue 广告展示总价值
	TotalAdImpressionEventsValue model.Float64 `json:"total_ad_impression_events_value,omitempty"`
	// TotalAdImpressionROAS 广告展示 ROAS
	TotalAdImpressionROAS model.Float64 `json:"total_ad_impression_roas,omitempty"`
	// AdImpressionAdRevenueROASDay0 第 0 天广告收入 ROAS
	AdImpressionAdRevenueROASDay0 model.Float64 `json:"ad_impression_ad_revenue_roas_day0,omitempty"`
	// AdImpressionAdRevenueROASDay2 第 2 天广告收入 ROAS
	AdImpressionAdRevenueROASDay2 model.Float64 `json:"ad_impression_ad_revenue_roas_day2,omitempty"`
	// AdImpressionAdRevenueROASDay6 第 6 天广告收入 ROAS
	AdImpressionAdRevenueROASDay6 model.Float64 `json:"ad_impression_ad_revenue_roas_day6,omitempty"`
	// AdImpressionAdRevenueDay0 第 0 天广告收入
	AdImpressionAdRevenueDay0 model.Float64 `json:"ad_impression_ad_revenue_day0,omitempty"`
	// AdImpressionAdRevenueDay2 第 2 天广告收入
	AdImpressionAdRevenueDay2 model.Float64 `json:"ad_impression_ad_revenue_day2,omitempty"`
	// AdImpressionAdRevenueDay6 第 6 天广告收入
	AdImpressionAdRevenueDay6 model.Float64 `json:"ad_impression_ad_revenue_day6,omitempty"`
	//网站指标
	// 网站指标衡量广告带来的网站操作，例如付费。
	//
	// CompletePaymentRoas 付费 ROAS（网站）
	// 网站付费操作带来的广告支出回报率 (ROAS)。
	CompletePaymentRoas model.Float64 `json:"complete_payment_roas,omitempty"`
	// CompletePayment 付费数（网站）
	// 归因于广告的网站付费操作次数。
	CompletePayment model.Int64 `json:"complete_payment,omitempty"`
	// CostPerCompletePayment 付费平均成本（网站）
	// 归因于广告的网站付费操作的平均成本。
	CostPerCompletePayment model.Float64 `json:"cost_per_complete_payment,omitempty"`
	// CompletePaymentRate 去重付费率（网站） (%)
	// 归因于广告的网站付费操作次数占总点击数的百分比。
	CompletePaymentRate model.Float64 `json:"complete_payment_rate,omitempty"`
	// ValuePerCompletePayment 付费平均价值（网站）
	// 归因于广告的网站付费操作的平均价值。
	ValuePerCompletePayment model.Float64 `json:"value_per_complete_payment,omitempty"`
	// TotalCompletePaymentRate 付费价值（网站）
	// 归因于广告的网站付费操作的总价值。
	TotalCompletePaymentRate model.Float64 `json:"total_complete_payment_rate,omitempty"`
	// TotalLandingPageView 落地页总浏览量（网站）
	// 来自广告的落地页浏览事件数。
	TotalLandingPageView model.Int64 `json:"total_landing_page_view,omitempty"`
	// CostPerLandingPageView 落地页浏览总平均成本（网站）
	// 广告主网站上发生落地页浏览事件的平均成本。
	CostPerLandingPageView model.Float64 `json:"cost_per_landing_page_view,omitempty"`
	// LandingPageViewRate 落地页总浏览率（网站） (%)
	// 落地页浏览事件数占点击事件数的百分比。
	LandingPageViewRate model.Float64 `json:"landing_page_view_rate,omitempty"`
	// TotalPageview 网页浏览量（网站）
	// 网页浏览事件发生的次数。
	TotalPageview model.Int64 `json:"total_pageview,omitempty"`
	// CostPerPageview 网页浏览平均成本（网站）
	// 网页浏览事件的平均成本。
	CostPerPageview model.Float64 `json:"cost_per_pageview,omitempty"`
	// PageviewRate 网页浏览率（网站） (%)
	// 网页浏览事件数占所有点击事件数的百分比。
	PageviewRate model.Float64 `json:"tageview_rate,omitempty"`
	// AvgValuePerPageview 网页浏览平均价值（网站）
	// 单次网页浏览事件返回的平均价值。
	AvgValuePerPageview model.Float64 `json:"avg_value_per_pageview,omitempty"`
	// TotalValuePerPageview 网页浏览总价值（网站）
	// 所有网页浏览事件返回的总价值。
	TotalValuePerPageview model.Float64 `json:"total_value_per_pageview,omitempty"`
	// ButtonClick 点击按钮次数（网站）
	// 归因于广告的在广告主网站上点击按钮操作次数。
	ButtonClick model.Int64 `json:"button_click,omitempty"`
	// CostPerButtonClick 点击按钮平均成本（网站）
	// 归因于广告的在广告主网站上点击按钮操作的平均成本。
	CostPerButtonClick model.Float64 `json:"cost_per_button_click,omitempty"`
	// ButtonClickRate 去重点击按钮率（网站） (%)
	// 归因于广告的在广告主网站上点击按钮操作次数占总点击量的百分比。
	ButtonClickRate model.Float64 `json:"button_click_rate,omitempty"`
	// ValuePerButtonClick 点击按钮平均价值（网站）
	// 归因于广告的在广告主网站上点击按钮操作的平均价值。
	ValuePerButtonClick model.Float64 `json:"value_per_button_click,omitempty"`
	// TotalButtonClickValue 点击按钮价值（网站）
	// 归因于广告的在广告主网站上点击按钮操作的总价值。
	TotalButtonClickValue model.Float64 `json:"total_button_click_value,omitempty"`
	// OnlineConsult 联系次数（网站）
	// 归因于广告的在广告主网站上进行联系操作的次数。
	OnlineConsult model.Int64 `json:"online_consult,omitempty"`
	// CostPerOnlineConsult 平均联系成本（网站）
	// 归因于广告的在广告主网站上进行联系操作的平均成本。
	CostPerOnlineConsult model.Float64 `json:"cost_per_online_consult,omitempty"`
	// OnlineConsultRate 去重联系率（网站） (%)
	// 归因于广告的在广告主网站上进行联系操作次数占总点击量的百分比。
	OnlineConsultRate model.Float64 `json:"online_consult_rate,omitempty"`
	// ValuePerOnlineConsult 平均联系价值（网站）
	// 归因于广告的在广告主网站上进行联系操作的平均价值。
	ValuePerOnlineConsult model.Float64 `json:"value_per_online_consult,omitempty"`
	// TotalOnlineConsultValue 联系价值（网站）
	// 归因于广告的在广告主网站上进行联系操作的总价值。
	TotalOnlineConsultValue model.Float64 `json:"total_online_consult_value,omitempty"`
	// UserRegistration 注册次数（网站）
	// 归因于广告的在广告主网站上完成注册操作次数。
	UserRegistration model.Int64 `json:"user_registration,omitempty"`
	// CostPerUserRegistration 注册平均成本（网站）
	// 归因于广告的在广告主网站上完成注册操作的平均成本。
	CostPerUserRegistration model.Float64 `json:"cost_per_user_registration,omitempty"`
	// UserRegistrationRate 去重注册率（网站） (%)
	// 归因于广告的在广告主网站上完成注册操作次数占总点击量的百分比。
	UserRegistrationRate model.Float64 `json:"user_registration_rate,omitempty"`
	// ValuePerUserRegistration 注册平均价值（网站）
	// 归因于广告的在广告主网站上完成注册操作的平均价值。
	ValuePerUserRegistration model.Float64 `json:"value_per_user_registration,omitempty"`
	// TotalUserRegistrationValue 注册价值（网站）
	// 归因于广告的在广告主网站上完成注册操作的总价值。
	TotalUserRegistrationValue model.Float64 `json:"total_user_registration_value,omitempty"`
	// PageContentViewEvents 查看内容次数（网站）
	// 归因于广告的在广告主网站上查看内容操作的次数。
	PageContentViewEvents model.Int64 `json:"page_content_view_events,omitempty"`
	// CostPerPageContentViewEvent 查看内容平均成本（网站）
	// 归因于广告的在广告主网站上查看内容操作的平均成本。
	CostPerPageContentViewEvent model.Float64 `json:"cost_per_page_content_view_event,omitempty"`
	// PageContentViewEventRate 去重查看内容率（网站） (%)
	// 归因于广告的在广告主网站上查看内容操作次数占总点击量的百分比。
	PageContentViewEventRate model.Float64 `json:"page_content_view_event_rate,omitempty"`
	// ValuePerPageContentViewEvent 查看内容平均价值（网站）
	// 归因于广告的在广告主网站上查看内容操作的平均价值。
	ValuePerPageContentViewEvent model.Float64 `json:"value_per_page_content_view_event,omitempty"`
	// TotalPageViewContentEventsValue 查看内容价值（网站）
	// 归因于广告的在广告主网站上查看内容操作的总价值。
	TotalPageViewContentEventsValue model.Float64 `json:"total_page_view_content_events_value,omitempty"`
	// WebEventAddToCart 加入购物车次数（网站）
	// 归因于广告的在广告主网站上将商品加入购物车操作次数。
	WebEventAddToCart model.Int64 `json:"web_event_add_to_cart,omitempty"`
	// CostPerWebEventAddToCart 加入购物车平均成本（网站）
	// 归因于广告的在广告主网站上将商品加入购物车操作的平均成本。
	CostPerWebEventAddToCart model.Float64 `json:"cost_per_web_event_add_to_cart,omitempty"`
	// WebEventAddToCartRate 去重加入购物车率（网站） (%)
	// 归因于广告的在广告主网站上将商品加入购物车操作次数占总点击量的百分比。
	WebEventAddToCartRate model.Float64 `json:"web_event_add_to_cart_rate,omitempty"`
	// ValuePerWebEventAddToCart 加入购物车平均价值（网站）
	// 归因于广告的在广告主网站上将商品加入购物车操作的平均价值。
	ValuePerWebEventAddToCart model.Float64 `json:"value_per_web_event_add_to_cart,omitempty"`
	// TotalWebEventAddToCartValue 加入购物车价值（网站）
	// 归因于广告的在广告主网站上将商品加入购物车操作的总价值。
	TotalWebEventAddToCartValue model.Float64 `json:"total_web_event_add_to_cart_value,omitempty"`
	// OnWebOrder 下单次数（网站）
	// 归因于广告的在广告主网站上下单操作的次数。
	OnWebOrder model.Int64 `json:"on_web_order,omitempty"`
	// CostPerOnWebOrder 平均下单成本（网站）
	// 归因于广告的在广告主网站上下单操作的平均成本。
	CostPerOnWebOrder model.Float64 `json:"cost_per_on_web_order,omitempty"`
	// OnWebOrderRate 去重下单率（网站） (%)
	// 归因于广告的在广告主网站上下单操作次数占总点击量的百分比。
	OnWebOrderRate model.Float64 `json:"on_web_order_rate,omitempty"`
	// ValuePerOnWebOrder 平均下单价值（网站）
	// 归因于广告的在广告主网站上下单操作的平均价值。
	ValuePerOnWebOrder model.Float64 `json:"value_per_on_web_order,omitempty"`
	// TotalOnWebOrderValue 下单价值（网站）
	// 归因于广告的在广告主网站上下单操作的总价值。
	TotalOnWebOrderValue model.Float64 `json:"total_on_web_order_value,omitempty"`
	// InitiateCheckout 开始结账次数（网站）
	// 归因于广告的在广告主网站上开始结账操作的次数。
	InitiateCheckout model.Int64 `json:"initiate_checkout,omitempty"`
	// CostPerInitiateCheckout 开始结账平均成本（网站）
	// 归因于广告的在广告主网站上开始结账操作的平均成本。
	CostPerInitiateCheckout model.Float64 `json:"cost_per_initiate_checkout,omitempty"`
	// InitiateCheckoutRate 去重开始结账率（网站） (%)
	// 归因于广告的在广告主网站上开始结账操作次数占总点击量的百分比。
	InitiateCheckoutRate model.Float64 `json:"initiate_checkout_rate,omitempty"`
	// ValuePerInitiateCheckout 开始结账平均价值（网站）
	// 归因于广告的在广告主网站上开始结账操作的平均价值。
	ValuePerInitiateCheckout model.Float64 `json:"value_per_initiate_checkout,omitempty"`
	// TotalInitiateCheckoutValue 开始结账价值（网站）
	// 归因于广告的在广告主网站上开始结账操作的总价值。
	TotalInitiateCheckoutValue model.Float64 `json:"total_initiate_checkout_value,omitempty"`
	// AddBilling 添加支付信息（网站）
	// 网站上的添加支付信息操作次数。
	AddBilling model.Int64 `json:"add_billing,omitempty"`
	// CostPerAddBilling 添加支付信息平均成本（网站）
	// 归因于广告的在广告主网站上添加支付信息操作的平均成本。
	CostPerAddBilling model.Float64 `json:"cost_per_add_billing,omitempty"`
	// AddBillingRate 去重添加支付信息率（网站） (%)
	// 归因于广告的在广告主网站上添加支付信息操作次数占总点击量的百分比。
	AddBillingRate model.Float64 `json:"add_billing_rate,omitempty"`
	// ValuePerAddBilling 添加支付信息平均价值（网站）
	// 归因于广告的在广告主网站上添加支付信息操作的平均价值。
	ValuePerAddBilling model.Float64 `json:"value_per_add_billing,omitempty"`
	// TotalAddBillingValue 添加支付信息价值（网站）
	// 归因于广告的在广告主网站上添加支付信息操作的总价值。
	TotalAddBillingValue model.Float64 `json:"total_add_billing_value,omitempty"`
	// PageEventSearch 搜索次数（网站）
	PageEventSearch model.Int64 `json:"page_event_search,omitempty"`
	// CostPerPageEventSearch 平均搜索成本（网站）
	CostPerPageEventSearch model.Float64 `json:"cost_per_page_event_search,omitempty"`
	// PageEventSearchRate 平均搜索成本（网站）
	PageEventSearchRate model.Float64 `json:"page_event_search_rate,omitempty"`
	// ValuePerPageEventSearch 平均搜索价值（网站）
	ValuePerPageEventSearch model.Float64 `json:"value_per_page_event_search,omitempty"`
	// TotalPageEventSearchValue 搜索价值（网站）
	TotalPageEventSearchValue model.Float64 `json:"total_page_event_search_value,omitempty"`
	// Form 提交表单次数（网站）
	Form model.Int64 `json:"form,omitempty"`
	// CostPerForm 提交表单平均成本（网站）
	CostPerForm model.Float64 `json:"cost_per_form,omitempty"`
	// FormRate 去重提交表单率（网站） (%)
	FormRate model.Float64 `json:"form_rate,omitempty"`
	// ValuePerForm 提交表单平均价值（网站）
	ValuePerForm model.Float64 `json:"value_per_form,omitempty"`
	// TotalFormValue 提交表单价值（网站）
	TotalFormValue model.Float64 `json:"total_form_value,omitempty"`
	// DownloadStart 下载次数（网站）
	DownloadStart model.Int64 `json:"download_start,omitempty"`
	// CostPerDownloadStart 下载平均成本（网站）
	CostPerDownloadStart model.Float64 `json:"cost_per_download_start,omitempty"`
	// DownloadStartRate 去重下载率（网站） (%)
	DownloadStartRate model.Float64 `json:"download_start_rate,omitempty"`
	// ValuePerDownloadStart 下载平均价值（网站）
	ValuePerDownloadStart model.Float64 `json:"value_per_download_start,omitempty"`
	// TotalDownloadStartValue 下载价值（网站）
	TotalDownloadStartValue model.Float64 `json:"total_download_start_value,omitempty"`
	// OnWebAddToWishlist 加入心愿单次数（网站）
	OnWebAddToWishlist model.Int64 `json:"on_web_add_to_wishlist,omitempty"`
	// CostPerOnWebAddToWishlist 加入心愿单平均成本（网站）
	CostPerOnWebAddToWishlist model.Float64 `json:"cost_per_on_web_add_to_wishlist,omitempty"`
	// OnWebAddToWishlistRate 去重加入心愿单率（网站） (%)
	OnWebAddToWishlistRate model.Float64 `json:"on_web_add_to_wishlist_rate,omitempty"`
	// ValuePerOnWebAddToWishlist 加入心愿单平均价值（网站）
	ValuePerOnWebAddToWishlist model.Float64 `json:"value_per_on_web_add_to_wishlist,omitempty"`
	// TotalOnWebAddToWishlistValue 加入心愿单价值（网站）
	TotalOnWebAddToWishlistValue model.Float64 `json:"total_on_web_add_to_wishlist_value,omitempty"`
	// OnWebSubscribe 订阅次数（网站）
	OnWebSubscribe model.Int64 `json:"on_web_subscribe,omitempty"`
	// CostPerOnWebSubscribe 订阅平均成本（网站）
	CostPerOnWebSubscribe model.Float64 `json:"cost_per_on_web_subscribe,omitempty"`
	// OnWebSubscribeRate 去重订阅率（网站） (%)
	OnWebSubscribeRate model.Float64 `json:"on_web_subscribe_rate,omitempty"`
	// ValuePerOnWebSubscribe 订阅平均价值（网站）
	ValuePerOnWebSubscribe model.Float64 `json:"value_per_on_web_subscribe,omitempty"`
	// TotalOnWebSubscribeValue 订阅价值（网站）
	TotalOnWebSubscribeValue model.Float64 `json:"total_on_web_subscribe_value,omitempty"`
	// WebsiteTotalFindLocation 寻找位置次数（网站）
	WebsiteTotalFindLocation model.Int64 `json:"website_total_find_location,omitempty"`
	// WebsiteCostPerFindLocation 寻找位置平均成本（网站）
	WebsiteCostPerFindLocation model.Float64 `json:"website_cost_per_find_location,omitempty"`
	// WebsiteFindLocationRate 寻找位置率（网站） (%)
	WebsiteFindLocationRate model.Float64 `json:"website_find_location_rate,omitempty"`
	// WebsiteValuePerFindLocation 寻找位置平均价值（网站）
	WebsiteValuePerFindLocation model.Float64 `json:"website_value_per_find_location,omitempty"`
	// WebsiteTotalFindLocationValue 寻找位置价值（网站）
	WebsiteTotalFindLocationValue model.Float64 `json:"website_total_find_location_value,omitempty"`
	// WebsiteTotalSchedule 预约次数（网站）
	WebsiteTotalSchedule model.Int64 `json:"website_total_schedule,omitempty"`
	// WebsiteCostPerSchedule 平均预约成本（网站）
	WebsiteCostPerSchedule model.Float64 `json:"website_cost_per_schedule,omitempty"`
	// WebsiteScheduleRate 预约率（网站） (%)
	WebsiteScheduleRate model.Float64 `json:"website_schedule_rate,omitempty"`
	// WebsiteValuePerSchedule 平均预约价值（网站）
	WebsiteValuePerSchedule model.Float64 `json:"website_value_per_schedule,omitempty"`
	// WebsiteTotalScheduleValue 预约价值（网站）
	WebsiteTotalScheduleValue model.Float64 `json:"website_total_schedule_value,omitempty"`
	// WebsiteTotalCustomizeProduct 定制商品次数（网站）
	WebsiteTotalCustomizeProduct model.Int64 `json:"website_total_customize_product,omitempty"`
	// WebsiteCostPerCustomizeProduct 定制商品平均成本（网站）
	WebsiteCostPerCustomizeProduct model.Float64 `json:"website_cost_per_customize_product,omitempty"`
	// WebsiteCustomizeProductRate 定制商品率（网站） (%)
	WebsiteCustomizeProductRate model.Float64 `json:"website_customize_product_rate,omitempty"`
	// WebsiteValuePerCustomizeProduct 定制商品平均价值（网站）
	WebsiteValuePerCustomizeProduct model.Float64 `json:"website_value_per_customize_product,omitempty"`
	// WebsiteTotalCustomizeProductValue 定制商品总价值（网站）
	WebsiteTotalCustomizeProductValue model.Float64 `json:"website_total_customize_product_value,omitempty"`
	// CustomPageEvents 自定义事件数（网站）
	CustomPageEvents model.Int64 `json:"custom_page_events,omitempty"`
	// CostPerCustomPageEvent 自定义事件平均成本（网站）
	CostPerCustomPageEvent model.Float64 `json:"cost_per_custom_page_event,omitempty"`
	// CustomPageEventRate 自定义事件率（网站） (%)
	CustomPageEventRate model.Float64 `json:"custom_page_event_rate,omitempty"`
	// ValuePerCustomPageEvent 自定义事件平均价值（网站）
	ValuePerCustomPageEvent model.Float64 `json:"value_per_custom_page_event,omitempty"`
	// CustomPageEventsValue 自定义事件价值（网站）
	CustomPageEventsValue model.Float64 `json:"custom_page_events_value,omitempty"`
	//
	// TikTok 指标
	// OnsitePurchasesROAS 付费 ROAS (TikTok)
	OnsitePurchasesROAS model.Float64 `json:"onsite_purchases_roas,omitempty"`
	// OnsiteTotalPurchase 付费量 (TikTok)
	OnsiteTotalPurchase model.Int64 `json:"onsite_total_purchase,omitempty"`
	// OnsiteCostPerPurchase 付费平均成本 (TikTok)
	OnsiteCostPerPurchase model.Float64 `json:"onsite_cost_per_purchase,omitempty"`
	// OnsitePurchaseRate 付费率 (TikTok) (%)
	OnsitePurchaseRate model.Float64 `json:"onsite_purchase_rate,omitempty"`
	// OnsiteValuePerPurchase 单次付费价值 (TikTok)
	OnsiteValuePerPurchase model.Float64 `json:"onsite_value_per_purchase,omitempty"`
	// OnsiteTotalPurchaseValue 付费价值 (TikTok)
	OnsiteTotalPurchaseValue model.Float64 `json:"onsite_total_purchase_value,omitempty"`
	// OnsiteUniquePurchase 去重付费量
	OnsiteUniquePurchase model.Int64 `json:"onsite_unique_purchase,omitempty"`
	// OnsiteCostPerUniquePurchase 去重付费平均成本
	OnsiteCostPerUniquePurchase model.Float64 `json:"onsite_cost_per_unique_purchase,omitempty"`
	// OnsiteTotalPurchaseValueDay0 第 0 天付费价值
	OnsiteTotalPurchaseValueDay0 model.Int64 `json:"onsite_total_purchase_value_day0,omitempty"`
	// OnsiteTotalPurchaseValueDay6 第 6 天付费价值
	OnsiteTotalPurchaseValueDay6 model.Float64 `json:"onsite_total_purchase_value_day6,omitempty"`
	// OnsiteTotalPurchaseValueDay13 第 13 天付费价值
	OnsiteTotalPurchaseValueDay13 model.Float64 `json:"onsite_total_purchase_value_day13,omitempty"`
	// OnsiteTotalPurchaseValueCalendarDay0 第 0 日历日付费价值
	OnsiteTotalPurchaseValueCalendarDay0 model.Float64 `json:"onsite_total_purchase_value_calendar_day0,omitempty"`
	// OnsiteTotalPurchaseValueCalendarDay6 第 6 日历日付费价值
	OnsiteTotalPurchaseValueCalendarDay6 model.Float64 `json:"onsite_total_purchase_value_calendar_day6,omitempty"`
	// OnsiteTotalPurchaseValueCalendarDay13 第 13 日历日付费价值
	OnsiteTotalPurchaseValueCalendarDay13 model.Float64 `json:"onsite_total_purchase_value_calendar_day13,omitempty"`
	// OnsiteTotalPurchaseROASDay0 第 0 天付费 (ROAS)
	OnsiteTotalPurchaseROASDay0 model.Float64 `json:"onsite_total_purchase_roas_day0,omitempty"`
	// OnsiteTotalPurchaseROASDay6 第 6 天付费 (ROAS)
	OnsiteTotalPurchaseROASDay6 model.Float64 `json:"onsite_total_purchase_roas_day6,omitempty"`
	// OnsiteTotalPurchaseROASDay13 第 13 天付费 (ROAS)
	OnsiteTotalPurchaseROASDay13 model.Float64 `json:"onsite_total_purchase_roas_day13,omitempty"`
	// OnsiteTotalPurchaseROASCalendarDay0 第 0 日历日付费 (ROAS)
	OnsiteTotalPurchaseROASCalendarDay0 model.Float64 `json:"onsite_total_purchase_roas_calendar_day0,omitempty"`
	// OnsiteTotalPurchaseROASCalendarDay6 第 6 日历日付费 (ROAS)
	OnsiteTotalPurchaseROASCalendarDay6 model.Float64 `json:"onsite_total_purchase_roas_calendar_day6,omitempty"`
	// OnsiteTotalPurchaseROASCalendarDay13 第 13 日历日付费 (ROAS)
	OnsiteTotalPurchaseROASCalendarDay13 model.Float64 `json:"onsite_total_purchase_roas_calendar_day13,omitempty"`
	// OnsiteTotalCheckoutInitiation 开始结账数量 (TikTok)
	OnsiteTotalCheckoutInitiation model.Int64 `json:"onsite_total_checkout_initiation,omitempty"`
	// OnsiteCostPerCheckoutInitiation 开始结账平均成本 (TikTok)
	OnsiteCostPerCheckoutInitiation model.Float64 `json:"onsite_cost_per_checkout_initiation,omitempty"`
	// OnsiteCheckoutInitiationRate 开始结账率 (TikTok) (%)
	OnsiteCheckoutInitiationRate model.Float64 `json:"onsite_checkout_initiation_rate,omitempty"`
	// OnsiteValuePerCheckoutInitiation 开始结账平均价值 (TikTok)
	OnsiteValuePerCheckoutInitiation model.Float64 `json:"onsite_value_per_checkout_initiation,omitempty"`
	// OnsiteTotalCheckoutInitiationValue 开始结账价值 (TikTok)
	OnsiteTotalCheckoutInitiationValue model.Float64 `json:"onsite_total_checkout_initiation_value,omitempty"`
	// OnsiteUniqueCheckoutInitiation 去重结账数
	OnsiteUniqueCheckoutInitiation model.Int64 `json:"onsite_unique_checkout_initiation,omitempty"`
	/// OnsiteCostPerUniqueCheckoutInitiation 去重结账平均成本
	OnsiteCostPerUniqueCheckoutInitiation model.Float64 `json:"onsite_cost_per_unique_checkout_initiation,omitempty"`
	// OnsiteTotalAddToCart 加入购物车次数 (TikTok)
	OnsiteTotalAddToCart model.Int64 `json:"onsite_total_add_to_cart,omitempty"`
	// OnsiteCostPerAddToCart 加入购物车平均成本 (TikTok)
	OnsiteCostPerAddToCart model.Float64 `json:"onsite_cost_per_add_to_cart,omitempty"`
	// OnsiteAddToCartRate 加入购物车率 (TikTok) (%)
	OnsiteAddToCartRate model.Float64 `json:"onsite_add_to_cart_rate,omitempty"`
	// OnsiteValuePerAddToCart 加入购物车平均价值 (TikTok)
	OnsiteValuePerAddToCart model.Float64 `json:"onsite_value_per_add_to_cart,omitempty"`
	// OnsiteTotalAddToCartValue 加入购物车价值 (TikTok)
	OnsiteTotalAddToCartValue model.Float64 `json:"onsite_total_add_to_cart_value,omitempty"`
	// OnsiteUniqueAddToCart 去重加入购物车事件数
	OnsiteUniqueAddToCart model.Int64 `json:"onsite_unique_add_to_cart,omitempty"`
	// OnsiteCostPerUniqueAddToCart 去重加入购物车操作平均成本
	OnsiteCostPerUniqueAddToCart model.Float64 `json:"onsite_cost_per_unique_add_to_cart,omitempty"`
	// OnsiteTotalProductDetailsPageView 商品详情页浏览数 (TikTok)
	OnsiteTotalProductDetailsPageView model.Int64 `json:"onsite_total_product_details_page_view,omitempty"`
	// OnsiteCostPerProductDetailsPageView 单次商品详情页浏览成本 (TikTok)
	OnsiteCostPerProductDetailsPageView model.Float64 `json:"onsite_cost_per_product_details_page_view,omitempty"`
	// OnsiteProductDetailsPageViewRate 商品详情页浏览率 (TikTok) (%)
	OnsiteProductDetailsPageViewRate model.Float64 `json:"onsite_product_details_page_view_rate,omitempty"`
	// OnsiteValuePerProductDetailsPageView 单次商品详情页浏览价值 (TikTok)
	OnsiteValuePerProductDetailsPageView model.Float64 `json:"onsite_value_per_product_details_page_view,omitempty"`
	// OnsiteTotalProductDetailsPageViewValue 商品详情页浏览价值 (TikTok)
	OnsiteTotalProductDetailsPageViewValue model.Float64 `json:"onsite_total_product_details_page_view_value,omitempty"`
	// OnsiteUniqueFirstLaunchConversionTime 按转化时间计算的活跃度
	OnsiteUniqueFirstLaunchConversionTime model.Int64 `json:"onsite_unique_first_launch_conversion_time,omitempty"`
	// OnsiteCostPerUniqueFirstLaunchConversionTime 按转化时间计算的活跃度平均成本
	OnsiteCostPerUniqueFirstLaunchConversionTime model.Float64 `json:"onsite_cost_per_unique_first_launch_conversion_time,omitempty"`
	// OnsiteUniqueFirstLaunch 活跃度
	OnsiteUniqueFirstLaunch model.Int64 `json:"onsite_unique_first_launch,omitempty"`
	// OnsiteCostPerUniqueFirstLaunch 活跃度平均成本
	OnsiteCostPerUniqueFirstLaunch model.Float64 `json:"onsite_cost_per_unique_first_launch,omitempty"`
	// OnsiteUniqueNonFirstLaunch 去重打开次数
	OnsiteUniqueNonFirstLaunch model.Int64 `json:"onsite_unique_non_first_launch,omitempty"`
	// OnsiteCostPerUniqueNonFirstLaunch 去重打开平均成本
	OnsiteCostPerUniqueNonFirstLaunch model.Float64 `json:"onsite_cost_per_unique_non_first_launch,omitempty"`
	// OnsiteLaunchRate 打开率 (%)
	OnsiteLaunchRate model.Float64 `json:"onsite_launch_rate,omitempty"`
	// OnsiteTotalNonFirstLaunch 总打开次数
	OnsiteTotalNonFirstLaunch model.Int64 `json:"onsite_total_non_first_launch,omitempty"`
	// OnsiteCostPerNonFirstLaunch 打开平均成本
	OnsiteCostPerNonFirstLaunch model.Float64 `json:"onsite_cost_per_non_first_launch,omitempty"`
	// OnsiteUniqueRegistration 去重注册次数
	OnsiteUniqueRegistration model.Int64 `json:"onsite_unique_registration,omitempty"`
	// OnsiteCostPerUniqueRegistration 去重注册平均成本
	OnsiteCostPerUniqueRegistration model.Float64 `json:"onsite_cost_per_unique_registration,omitempty"`
	// OnsiteRegistrationRate 注册率 (%)
	OnsiteRegistrationRate model.Float64 `json:"onsite_registration_rate,omitempty"`
	// OnsiteTotalRegistration 总注册次数
	OnsiteTotalRegistration model.Int64 `json:"onsite_total_registration,omitempty"`
	// OnsiteCostPerRegistration 注册平均成本
	OnsiteCostPerRegistration model.Float64 `json:"onsite_cost_per_registration,omitempty"`
	// OnsiteUniqueAdImpression 去重广告曝光事件数
	OnsiteUniqueAdImpression model.Int64 `json:"onsite_unique_ad_impression,omitempty"`
	// OnsiteCostPerUniqueAdImpression 去重广告曝光事件平均成本
	OnsiteCostPerUniqueAdImpression model.Float64 `json:"onsite_cost_per_unique_ad_impression,omitempty"`
	// OnsiteAdImpressionRate 广告曝光事件率
	OnsiteAdImpressionRate model.Float64 `json:"onsite_ad_impression_rate,omitempty"`
	// OnsiteTotalAdImpression 广告曝光事件总数
	OnsiteTotalAdImpression model.Int64 `json:"onsite_total_ad_impression,omitempty"`
	// OnsiteCostPerAdImpression 广告曝光事件平均成本
	OnsiteCostPerAdImpression model.Float64 `json:"onsite_cost_per_ad_impression,omitempty"`
	// OnsiteTotalAdImpressionValue 广告曝光总价值
	OnsiteTotalAdImpressionValue model.Float64 `json:"onsite_total_ad_impression_value,omitempty"`
	// OnsiteValuePerAdImpression 广告曝光事件平均价值
	OnsiteValuePerAdImpression model.Float64 `json:"onsite_value_per_ad_impression,omitempty"`
	// OnsiteAdImpressionAdRevenueDay0 第 0 天广告收入
	OnsiteAdImpressionAdRevenueDay0 model.Float64 `json:"onsite_ad_impression_ad_revenue_day0,omitempty"`
	// OnsiteAdImpressionAdRevenueDay6 第 6 天广告收入
	OnsiteAdImpressionAdRevenueDay6 model.Float64 `json:"onsite_ad_impression_ad_revenue_day6,omitempty"`
	// OnsiteAdImpressionAdRevenueDay13 第 13 天广告收入
	OnsiteAdImpressionAdRevenueDay13 model.Float64 `json:"onsite_ad_impression_ad_revenue_day13,omitempty"`
	// OnsiteAdImpressionAdRevenueCalendarDay0 第 0 日历日广告收入
	OnsiteAdImpressionAdRevenueCalendarDay0 model.Float64 `json:"onsite_ad_impression_ad_revenue_calendar_day0,omitempty"`
	// OnsiteAdImpressionAdRevenueCalendarDay6 第 6 日历日广告收入
	OnsiteAdImpressionAdRevenueCalendarDay6 model.Float64 `json:"onsite_ad_impression_ad_revenue_calendar_day6,omitempty"`
	// OnsiteAdImpressionAdRevenueCalendarDay13 第 13 日历日广告收入
	OnsiteAdImpressionAdRevenueCalendarDay13 model.Float64 `json:"onsite_ad_impression_ad_revenue_calendar_day13,omitempty"`
	// OnsiteAdImpressionAdRevenueROASDay0 第 0 天广告收入 (ROAS)
	OnsiteAdImpressionAdRevenueROASDay0 model.Float64 `json:"onsite_ad_impression_ad_revenue_roas_day0,omitempty"`
	// OnsiteAdImpressionAdRevenueROASDay6 第 6 天广告收入 (ROAS)
	OnsiteAdImpressionAdRevenueROASDay6 model.Float64 `json:"onsite_ad_impression_ad_revenue_roas_day6,omitempty"`
	// OnsiteAdImpressionAdRevenueROASDay13 第 13 天广告收入 (ROAS)
	OnsiteAdImpressionAdRevenueROASDay13 model.Float64 `json:"onsite_ad_impression_ad_revenue_roas_day13,omitempty"`
	// OnsiteAdImpressionAdRevenueROASCalendarDay0 第 0 日历日广告收入 (ROAS)
	OnsiteAdImpressionAdRevenueROASCalendarDay0 model.Float64 `json:"onsite_ad_impression_ad_revenue_roas_calendar_day0,omitempty"`
	// OnsiteAdImpressionAdRevenueROASCalendarDay6 第 6 日历日广告收入 (ROAS)
	OnsiteAdImpressionAdRevenueROASCalendarDay6 model.Float64 `json:"onsite_ad_impression_ad_revenue_roas_calendar_day6,omitempty"`
	// OnsiteAdImpressionAdRevenueROASCalendarDay13 第 13 日历日广告收入 (ROAS)
	OnsiteAdImpressionAdRevenueROASCalendarDay13 model.Float64 `json:"onsite_ad_impression_ad_revenue_roas_calendar_day13,omitempty"`
	// OnsiteAddToWishlist 加入心愿单次数 (TikTok)
	OnsiteAddToWishlist model.Int64 `json:"onsite_add_to_wishlist,omitempty"`
	// CostPerOnsiteAddToWishlist 加入心愿单平均成本 (TikTok)
	CostPerOnsiteAddToWishlist model.Float64 `json:"cost_per_onsite_add_to_wishlist,omitempty"`
	// OnsiteAddToWishlistRate 去重加入心愿单率 (TikTok)
	OnsiteAddToWishlistRate model.Float64 `json:"onsite_add_to_wishlist_rate,omitempty"`
	// ValuePerOnsiteAddToWishlist 加入心愿单平均价值 (TikTok)
	ValuePerOnsiteAddToWishlist model.Float64 `json:"value_per_onsite_add_to_wishlist,omitempty"`
	// TotalOnsiteAddToWishlistValue 加入心愿单价值 (TikTok)
	TotalOnsiteAddToWishlistValue model.Float64 `json:"total_onsite_add_to_wishlist_value,omitempty"`
	// OnsiteAddBilling 添加结算信息次数 (TikTok)
	OnsiteAddBilling model.Int64 `json:"onsite_add_billing,omitempty"`
	// CostPerOnsiteAddBilling 添加结算信息平均成本 (TikTok)
	CostPerOnsiteAddBilling model.Float64 `json:"cost_per_onsite_add_billing,omitempty"`
	// OnsiteAddBillingRate 添加结算信息率 (TikTok) (%)
	OnsiteAddBillingRate model.Float64 `json:"onsite_add_billing_rate,omitempty"`
	// ValuePerOnsiteAddBilling 添加结算信息平均价值 (TikTok)
	ValuePerOnsiteAddBilling model.Float64 `json:"value_per_onsite_add_billing,omitempty"`
	// TotalOnsiteAddBillingValue 添加结算信息价值 (TikTok)
	TotalOnsiteAddBillingValue model.Float64 `json:"total_onsite_add_billing_value,omitempty"`
	// OnsiteForm 提交表单次数 (TikTok)
	OnsiteForm model.Int64 `json:"onsite_form,omitempty"`
	// CostPerOnsiteForm 提交表单平均成本 (TikTok)
	CostPerOnsiteForm model.Float64 `json:"cost_per_onsite_form,omitempty"`
	// OnsiteFormRate 去重提交表单率 (TikTok)
	OnsiteFormRate model.Float64 `json:"onsite_form_rate,omitempty"`
	// ValuePerOnsiteForm 提交表单平均价值 (TikTok)
	ValuePerOnsiteForm model.Float64 `json:"value_per_onsite_form,omitempty"`
	// TotalOnsiteFormValue 提交表单价值 (TikTok)
	TotalOnsiteFormValue model.Float64 `json:"total_onsite_form_value,omitempty"`
	// OnsiteTotalFilteredOutFormSubmission 过滤后的表单提交数 (TikTok)
	OnsiteTotalFilteredOutFormSubmission model.Int64 `json:"onsite_total_filtered_out_form_submission,omitempty"`
	// OnsiteDownloadStart 应用商店点击数 (TikTok)
	OnsiteDownloadStart model.Int64 `json:"onsite_download_start,omitempty"`
	// CostPerOnsiteDownloadStart 应用商店点击平均成本 (TikTok)
	CostPerOnsiteDownloadStart model.Float64 `json:"cost_per_onsite_download_start,omitempty"`
	// OnsiteDownloadStartRate 应用商店点击率 (TikTok) (%)
	OnsiteDownloadStartRate model.Float64 `json:"onsite_download_start_rate,omitempty"`
	// OnsiteDestinationVisits 目标页面访问数 (TikTok)
	OnsiteDestinationVisits model.Int64 `json:"onsite_destination_visits,omitempty"`
	// CostPerOnsiteDestinationVisit 目标页面访问平均成本 (TikTok)
	CostPerOnsiteDestinationVisit model.Float64 `json:"cost_per_onsite_destionation_visit,omitempty"`
	// OnsiteDestinationVisitRate 目标页面访问率 (TikTok) (%)
	OnsiteDestinationVisitRate model.Float64 `json:"onsite_destination_visit_rate,omitempty"`
	// IxPageViewCount 网页浏览量 (TikTok)
	IxPageViewCount model.Int64 `json:"ix_page_view_count,omitempty"`
	// CostPerIxPageViewCount 网页浏览平均成本 (TikTok)
	CostPerIxPageViewCount model.Float64 `json:"cost_per_ix_page_view_count,omitempty"`
	// IxPageViewCountRate 网页浏览率 (TikTok) (%)
	IxPageViewCountRate model.Float64 `json:"ix_page_view_count_rate,omitempty"`
	// IxButtonClickCount 出站点击数 (TikTok)
	IxButtonClickCount model.Int64 `json:"ix_button_click_count,omitempty"`
	// CostPerIxButtonClickCount 出站点击平均成本 (TikTok)
	CostPerIxButtonClickCount model.Float64 `json:"cost_per_ix_button_click_count,omitempty"`
	// IxButtonClickCountRate 出站点击率 (TikTok) (%)
	IxButtonClickCountRate model.Float64 `json:"ix_button_click_count_rate,omitempty"`
	// IxProductClickCount 商品点击量 (TikTok)
	IxProductClickCount model.Int64 `json:"ix_product_click_count,omitempty"`
	// CostPerIxProductClickCount 商品点击平均成本 (TikTok)
	CostPerIxProductClickCount model.Float64 `json:"cost_per_ix_product_click_count,omitempty"`
	// IxProductClickCountRate 商品点击率 (TikTok) (%)
	IxProductClickCountRate model.Float64 `json:"ix_product_click_count_rate,omitempty"`
	//
	// 线下指标
	// OfflineTotalCRMEvents CRM 事件数
	OfflineTotalCRMEvents model.Int64 `json:"offline_total_crm_events,omitempty"`
	// OfflineCostPerCRMEvent CRM 事件平均成本
	OfflineCostPerCRMEvent model.Float64 `json:"offline_cost_per_crm_event,omitempty"`
	// OfflineCRMEventRate CRM 事件率 (%)
	OfflineCRMEventRate model.Float64 `json:"offline_crm_event_rate,omitempty"`
	// OfflineValuePerCRMEvent CRM 事件平均价值
	OfflineValuePerCRMEvent model.Float64 `json:"offline_value_per_crm_event,omitempty"`
	// OfflineCRMEventValue CRM 事件价值
	OfflineCRMEventValue model.Float64 `json:"offline_crm_event_value,omitempty"`
	// OfflineShoppingEvents 付费次数（线下）
	OfflineShoppingEvents model.Int64 `json:"offline_shopping_events,omitempty"`
	// CostPerOfflineShoppingEvent 付费平均成本（线下）
	CostPerOfflineShoppingEvent model.Float64 `json:"cost_per_offline_shopping_event,omitempty"`
	// OfflineShoppingEventRate 去重付费率（线下） (%)
	OfflineShoppingEventRate model.Float64 `json:"offline_shopping_event_rate,omitempty"`
	// ValuePerOfflineShoppingEvent 付费平均价值（线下）
	ValuePerOfflineShoppingEvent model.Float64 `json:"value_per_offline_shopping_event,omitempty"`
	// OfflineShoppingEventsValue 付费价值（线下）
	OfflineShoppingEventsValue model.Float64 `json:"offline_shopping_events_value,omitempty"`
	// OfflineShoppingEventsROAS 付费 ROAS（线下）
	OfflineShoppingEventsROAS model.Float64 `json:"offline_shopping_events_roas,omitempty"`
	// OfflineContactEvents 联系次数（线下）
	OfflineContactEvents model.Int64 `json:"offline_contact_events,omitempty"`
	// CostPerOfflineContactEvent 平均联系成本（线下）
	CostPerOfflineContactEvent model.Float64 `json:"cost_per_offline_contact_event,omitempty"`
	// OfflineContactEventRate 去重联系率（线下）
	OfflineContactEventRate model.Float64 `json:"offline_contact_event_rate,omitempty"`
	// ValuePerOfflineContactEvent 平均联系价值（线下）
	ValuePerOfflineContactEvent model.Float64 `json:"value_per_offline_contact_event,omitempty"`
	// OfflineContectEventsValue 联系价值（线下）
	OfflineContactEventsValue model.Float64 `json:"offline_contact_events_value,omitempty"`
	// OfflineSubscribeEvents 订阅次数（线下）
	OfflineSubscribeEvents model.Int64 `json:"offline_subscribe_events,omitempty"`
	// CostPerOfflineSubscribeEvent 平均订阅成本（线下）
	CostPerOfflineSubscribeEvent model.Float64 `json:"cost_per_offline_subscribe_event,omitempty"`
	// OfflineSubscribeEventRate 去重订阅率（线下）
	OfflineSubscribeEventRate model.Float64 `json:"offline_subscribe_event_rate,omitempty"`
	// ValuePerOfflineSubscribeEvent 平均订阅价值（线下）
	ValuePerOfflineSubscribeEvent model.Float64 `json:"value_per_offline_subscribe_event,omitempty"`
	// OfflineSubscribeEventsValue 订阅价值（线下）
	OfflineSubscribeEventsValue model.Float64 `json:"offline_subscribe_events_value,omitempty"`
	// OfflineFormEvents 提交表单次数（线下）
	OfflineFormEvents model.Int64 `json:"offline_form_events,omitempty"`
	// CostPerOfflineFormEvent 提交表单平均成本（线下）
	CostPerOfflineFormEvent model.Float64 `json:"cost_per_offline_form_event,omitempty"`
	// OfflineFormEventRate 去重提交表单率（线下）
	OfflineFormEventRate model.Float64 `json:"offline_form_event_rate,omitempty"`
	// ValuePerOfflineFormEvent 提交表单平均价值（线下）
	ValuePerOfflineFormEvent model.Float64 `json:"value_per_offline_form_event,omitempty"`
	// OfflineFormEventsValue 提交表单价值（线下）
	OfflineFormEventsValue model.Float64 `json:"offline_form_events_value,omitempty"`
	// OfflineAddPaymentInfoEvents 添加支付信息次数（线下）
	OfflineAddPaymentInfoEvents model.Int64 `json:"offline_add_payment_info_events,omitempty"`
	// CostPerOfflineAddPaymentInfoEvent 添加支付信息平均成本（线下）
	CostPerOfflineAddPaymentInfoEvent model.Float64 `json:"cost_per_offline_add_payment_info_event,omitempty"`
	// OfflineAddPaymentInfoEventRate 去重添加支付信息率（线下）
	OfflineAddPaymentInfoEventRate model.Float64 `json:"offline_add_payment_info_event_rate,omitempty"`
	// ValuePerOfflineAddPaymentInfoEvent 添加支付信息平均价值（线下）
	ValuePerOfflineAddPaymentInfoEvent model.Float64 `json:"value_per_offline_add_payment_info_event,omitempty"`
	// OfflineAddPaymentInfoEventsValue 添加支付信息价值（线下）
	OfflineAddPaymentInfoEventsValue model.Float64 `json:"offline_add_payment_info_events_value,omitempty"`
	// OfflineAddToCartEvents 加入购物车次数（线下）
	OfflineAddToCartEvents model.Int64 `json:"offline_add_to_cart_events,omitempty"`
	// CostPerOfflineAddToCartEvent 加入购物车平均成本（线下）
	CostPerOfflineAddToCartEvent model.Float64 `json:"cost_per_offline_add_to_cart_event,omitempty"`
	// OfflineAddToCartEventRate 去重加入购物车率（线下）
	OfflineAddToCartEventRate model.Float64 `json:"offline_add_to_cart_event_rate,omitempty"`
	// ValuePerOfflineAddToCartEvent 加入购物车平均价值（线下）
	ValuePerOfflineAddToCartEvent model.Float64 `json:"value_per_offline_add_to_cart_event,omitempty"`
	// OfflineAddToCartEventsValue 加入购物车价值（线下）
	OfflineAddToCartEventsValue model.Float64 `json:"offline_add_to_cart_events_value,omitempty"`
	// OfflineAddToWishlistEvents 加入心愿单次数（线下）
	OfflineAddToWishlistEvents model.Int64 `json:"offline_add_to_wishlist_events,omitempty"`
	// CostPerOfflineAddToWishlistEvent 加入心愿单平均成本（线下）
	CostPerOfflineAddToWishlistEvent model.Float64 `json:"cost_per_offline_add_to_wishlist_event,omitempty"`
	// OfflineAddToWishlistEventRate 去重加入心愿单率（线下）
	OfflineAddToWishlistEventRate model.Float64 `json:"offline_add_to_wishlist_event_rate,omitempty"`
	// ValuePerOfflineAddToWishlistEvent 加入心愿单平均价值（线下）
	ValuePerOfflineAddToWishlistEvent model.Float64 `json:"value_per_offline_add_to_wishlist_event,omitempty"`
	// OfflineAddToWishlistEventsValue 加入心愿单价值（线下）
	OfflineAddToWishlistEventsValue model.Float64 `json:"offline_add_to_wishlist_events_value,omitempty"`
	// OfflineClickButtonEvents 按钮点击次数（线下）
	OfflineClickButtonEvents model.Int64 `json:"offline_click_button_events,omitempty"`
	// CostPerOfflineClickButtonEvent 按钮点击平均成本（线下）
	CostPerOfflineClickButtonEvent model.Float64 `json:"cost_per_offline_click_button_event,omitempty"`
	// OfflineClickButtonEventRate 去重按钮点击率（线下）
	OfflineClickButtonEventRate model.Float64 `json:"offline_click_button_event_rate,omitempty"`
	// ValuePerOfflineClickButtonEvent 按钮点击平均价值（线下）
	ValuePerOfflineClickButtonEvent model.Float64 `json:"value_per_offline_click_button_event,omitempty"`
	// OfflineClickButtonEventsValue 按钮点击价值（线下）
	OfflineClickButtonEventsValue model.Float64 `json:"offline_click_button_events_value,omitempty"`
	// OfflineRegistrationEvents 完成注册次数（线下）
	OfflineRegistrationEvents model.Int64 `json:"offline_registration_events,omitempty"`
	// CostPerOfflineRegistrationEvent 完成注册平均成本（线下）
	CostPerOfflineRegistrationEvent model.Float64 `json:"cost_per_offline_registration_event,omitempty"`
	// OfflineRegistrationEventRate 去重完成注册率（线下）
	OfflineRegistrationEventRate model.Float64 `json:"offline_registration_event_rate,omitempty"`
	// ValuePerOfflineRegistrationEvent 完成注册平均价值（线下）
	ValuePerOfflineRegistrationEvent model.Float64 `json:"value_per_offline_registration_event,omitempty"`
	// OfflineRegistrationEventsValue 完成注册价值（线下）
	OfflineRegistrationEventsValue model.Float64 `json:"offline_registration_events_value,omitempty"`
	// OfflineDownloadEvents 下载次数（线下）
	OfflineDownloadEvents model.Int64 `json:"offline_download_events,omitempty"`
	// CostPerOfflineDownloadEvent 下载平均成本（线下）
	CostPerOfflineDownloadEvent model.Float64 `json:"cost_per_offline_download_event,omitempty"`
	// OfflineDownloadEventRate 去重下载率（线下）
	OfflineDownloadEventRate model.Float64 `json:"offline_download_event_rate,omitempty"`
	// ValuePerOfflineDownloadEvent 下载平均价值（线下）
	ValuePerOfflineDownloadEvent model.Float64 `json:"value_per_offline_download_event,omitempty"`
	// OfflineDownloadEventsValue 下载价值（线下）
	OfflineDownloadEventsValue model.Float64 `json:"offline_download_events_value,omitempty"`
	// OfflineInitiateCheckoutEvents 开始结账次数（线下）
	OfflineInitiateCheckoutEvents model.Int64 `json:"offline_initiate_checkout_events,omitempty"`
	// CostPerOfflineInitiateCheckoutEvent 开始结账平均成本（线下）
	CostPerOfflineInitiateCheckoutEvent model.Float64 `json:"cost_per_offline_initiate_checkout_event,omitempty"`
	// OfflineInitiateCheckoutEventRate 去重开始结账率（线下）
	OfflineInitiateCheckoutEventRate model.Float64 `json:"offline_initiate_checkout_event_rate,omitempty"`
	// ValuePerOfflineInitiateCheckoutEvent 开始结账平均价值（线下）
	ValuePerOfflineInitiateCheckoutEvent model.Float64 `json:"value_per_offline_initiate_checkout_event,omitempty"`
	// OfflineInitiateCheckoutEventsValue 开始结账价值（线下）
	OfflineInitiateCheckoutEventsValue model.Float64 `json:"offline_initiate_checkout_events_value,omitempty"`
	// OfflinePlaceOrderEvents 下单次数（线下）
	OfflinePlaceOrderEvents model.Int64 `json:"offline_place_order_events,omitempty"`
	// CostPerOfflinePlaceOrderEvent 下单平均成本（线下）
	CostPerOfflinePlaceOrderEvent model.Float64 `json:"cost_per_offline_place_order_event,omitempty"`
	// OfflinePlaceOrderEventRate 去重下单率（线下）
	OfflinePlaceOrderEventRate model.Float64 `json:"offline_place_order_event_rate,omitempty"`
	// ValuePerOfflinePlaceOrderEvent 下单平均价值（线下）
	ValuePerOfflinePlaceOrderEvent model.Float64 `json:"value_per_offline_place_order_event,omitempty"`
	// OfflinePlaceOrderEventsValue 下单价值（线下）
	OfflinePlaceOrderEventsValue model.Float64 `json:"offline_place_order_events_value,omitempty"`
	// OfflineSearchEvents 搜索次数（线下）
	OfflineSearchEvents model.Int64 `json:"offline_search_events,omitempty"`
	// CostPerOfflineSearchEvent 搜索平均成本（线下）
	CostPerOfflineSearchEvent model.Float64 `json:"cost_per_offline_search_event,omitempty"`
	// OfflineSearchEventRate 去重搜索率（线下）
	OfflineSearchEventRate model.Float64 `json:"offline_search_event_rate,omitempty"`
	// ValuePerOfflineSearchEvent 搜索平均价值（线下）
	ValuePerOfflineSearchEvent model.Float64 `json:"value_per_offline_search_event,omitempty"`
	// OfflineSearchEventsValue 搜索价值（线下）
	OfflineSearchEventsValue model.Float64 `json:"offline_search_events_value,omitempty"`
	// OfflineViewContentEvents 查看内容次数（线下）
	OfflineViewContentEvents model.Int64 `json:"offline_view_content_events,omitempty"`
	// CostPerOfflineViewContentEvent 查看内容平均成本（线下）
	CostPerOfflineViewContentEvent model.Float64 `json:"cost_per_offline_view_content_event,omitempty"`
	// OfflineViewContentEventRate 去重查看内容率（线下）
	OfflineViewContentEventRate model.Float64 `json:"offline_view_content_event_rate,omitempty"`
	// ValuePerOfflineViewContentEvent 查看内容平均价值（线下）
	ValuePerOfflineViewContentEvent model.Float64 `json:"value_per_offline_view_content_event,omitempty"`
	// OfflineViewContentEventsValue 查看内容价值（线下）
	OfflineViewContentEventsValue model.Float64 `json:"offline_view_content_events_value,omitempty"`
	// OfflinePerferredLeads 偏好线索数（线下）
	OfflinePerferredLeads model.Int64 `json:"offline_perferred_leads,omitempty"`
	// OfflineCostPerPreferredLead 偏好线索平均成本（线下）
	OfflineCostPerPreferredLead model.Float64 `json:"offline_cost_per_preferred_lead,omitempty"`
	// OfflinePreferredLeadRate 偏好线索率（线下） (%)
	OfflinePreferredLeadRate model.Float64 `json:"offline_preferred_lead_rate,omitempty"`
	// OfflineTotalFindLocation 寻找位置总数（线下）
	OfflineTotalFindLocation model.Int64 `json:"offline_total_find_location,omitempty"`
	// OfflineCostPerFindLocation 寻找位置平均成本（线下）
	OfflineCostPerFindLocation model.Float64 `json:"offline_cost_per_find_location,omitempty"`
	// OfflineFindLocationRate 寻找位置率（线下）
	OfflineFindLocationRate model.Float64 `json:"offline_find_location_rate,omitempty"`
	// OfflineValuePerFindLocation 寻找位置平均价值（线下）
	OfflineValuePerFindLocation model.Float64 `json:"offline_value_per_find_location,omitempty"`
	// OfflineTotalFindLocationValue 寻找位置价值（线下）
	OfflineTotalFindLocationValue model.Float64 `json:"offline_total_find_location_value,omitempty"`
	// OfflineTotalSchedule 预约总数（线下）
	OfflineTotalSchedule model.Int64 `json:"offline_total_schedule,omitempty"`
	// OfflineCostPerSchedule 预约平均成本（线下）
	OfflineCostPerSchedule model.Float64 `json:"offline_cost_per_schedule,omitempty"`
	// OfflineScheduleRate 预约率（线下）
	OfflineScheduleRate model.Float64 `json:"offline_schedule_rate,omitempty"`
	// OfflineValuePerSchedule 预约平均价值（线下）
	OfflineValuePerSchedule model.Float64 `json:"offline_value_per_schedule,omitempty"`
	// OfflineTotalScheduleValue 预约价值（线下）
	OfflineTotalScheduleValue model.Float64 `json:"offline_total_schedule_value,omitempty"`
	// OfflineTotalCustomizeProduct 定制商品总数（线下）
	OfflineTotalCustomizeProduct model.Int64 `json:"offline_total_customize_product,omitempty"`
	// OfflineCostPerCustomizeProduct 定制商品平均成本（线下）
	OfflineCostPerCustomizeProduct model.Float64 `json:"offline_cost_per_customize_product,omitempty"`
	// OfflineCustomizeProductRate 定制商品率（线下）
	OfflineCustomizeProductRate model.Float64 `json:"offline_customize_product_rate,omitempty"`
	// OfflineValuePerCustomizeProduct 定制商品平均价值（线下）
	OfflineValuePerCustomizeProduct model.Float64 `json:"offline_value_per_customize_product,omitempty"`
	// OfflineTotalCustomizeProductValue 定制商品价值（线下）
	OfflineTotalCustomizeProductValue model.Float64 `json:"offline_total_customize_product_value,omitempty"`
	// OfflineTotalCustomEvents 自定义事件总数（线下）
	OfflineTotalCustomEvents model.Int64 `json:"offline_total_custom_events,omitempty"`
	// OfflineCostPerCustomEvent 自定义事件平均成本（线下）
	OfflineCostPerCustomEvent model.Float64 `json:"offline_cost_per_custom_event,omitempty"`
	// OfflineCustomEventRate 自定义事件率（线下）
	OfflineCustomEventRate model.Float64 `json:"offline_custom_event_rate,omitempty"`
	// OfflineValuePerCustomEvent 自定义事件平均价值（线下）
	OfflineValuePerCustomEvent model.Float64 `json:"offline_value_per_custom_event,omitempty"`
	// OfflineTotalCustomEventValue 自定义事件价值（线下）
	OfflineTotalCustomEventValue model.Float64 `json:"offline_total_custom_event_value,omitempty"`
	//
	// 消息指标
	// MessagingTotalConversationTikTokDirectMessage 对话数（TikTok 私信）
	MessagingTotalConversationTikTokDirectMessage model.Int64 `json:"messaging_total_conversation_tiktok_direct_message,omitempty"`
	// MessagingCostPerConversationTikTokDirectMessage 对话平均成本（TikTok 私信）
	MessagingCostPerConversationTikTokDirectMessage model.Float64 `json:"messaging_cost_per_conversation_tiktok_direct_message,omitempty"`
	// MessagingConversationRateTikTokDirectMessage 对话率（TikTok 私信） (%)
	MessagingConversationRateTikTokDirectMessage model.Float64 `json:"messaging_conversation_rate_tiktok_direct_message,omitempty"`
	// MessagingTotalConversationInstantMessageApp 对话数（即时通讯应用）
	MessagingTotalConversationInstantMessageApp model.Int64 `json:"messaging_total_conversation_instant_messaging_app,omitempty"`
	// MessagingCostPerConversationInstantMessageApp 对话平均成本（即时通讯应用）
	MessagingCostPerConversationInstantMessageApp model.Float64 `json:"messaging_cost_per_conversation_instant_messaging_app,omitempty"`
	// MessagingConversationRateInstantMessageApp 对话率（即时通讯应用） (%)
	MessagingConversationRateInstantMessageApp model.Float64 `json:"messaging_conversation_rate_instant_messaging_app,omitempty"`
	//
	// 店铺指标
	// OnsiteShoppingROAS ROAS（店铺）
	OnsiteShoppingROAS model.Float64 `json:"onsite_shopping_roas,omitempty"`
	// OnsiteShopping 付费数（店铺）
	OnsiteShopping model.Int64 `json:"onsite_shopping,omitempty"`
	// CostPerOnsiteShopping 平均付费成本（店铺）
	CostPerOnsiteShopping model.Float64 `json:"cost_per_onsite_shopping,omitempty"`
	// OnsiteShoppingRate 付费率（店铺） (%)
	OnsiteShoppingRate model.Float64 `json:"onsite_shopping_rate,omitempty"`
	// ValuePerOnsiteShopping 平均订单价值（店铺）
	ValuePerOnsiteShopping model.Float64 `json:"value_per_onsite_shopping,omitempty"`
	// TotalOnsiteShoppingValue 总收入（店铺）
	TotalOnsiteShoppingValue model.Float64 `json:"total_onsite_shopping_value,omitempty"`
	// ShopTotalPurchaseByOrderSubmission 按订单提交时间计算的付费数（店铺）
	ShopTotalPurchaseByOrderSubmission model.Int64 `json:"shop_total_purchase_by_order_submission,omitempty"`
	// ShopGrossRevenueByOrderSubmission 按订单提交时间计算的总收入（店铺）
	ShopGrossRevenueByOrderSubmission model.Float64 `json:"shop_gross_revenue_by_order_submission,omitempty"`
	// ShopTotalItemsPurchased 付费商品数（店铺）
	ShopTotalItemsPurchased model.Int64 `json:"shop_total_items_purchased,omitempty"`
	// OnsiteInitiateCheckoutCount 开始结账数（店铺)
	OnsiteInitiateCheckoutCount model.Int64 `json:"onsite_initiate_checkout_count,omitempty"`
	// CostPerOnsiteInitiateCheckoutCount 开始结账平均成本（店铺）
	CostPerOnsiteInitiateCheckoutCount model.Float64 `json:"cost_per_onsite_initiate_checkout_count,omitempty"`
	// OnsiteInitiateCheckoutCountRate 开始结账率（店铺）
	OnsiteInitiateCheckoutCountRate model.Float64 `json:"onsite_initiate_checkout_count_rate,omitempty"`
	// ValuePerOnsiteInitiateCheckoutCount 开始结账平均价值（店铺）
	ValuePerOnsiteInitiateCheckoutCount model.Float64 `json:"value_per_onsite_initiate_checkout_count,omitempty"`
	// TotalOnsiteInitiateCheckoutCountValue 开始结账价值（店铺）
	TotalOnsiteInitiateCheckoutCountValue model.Float64 `json:"total_onsite_initiate_checkout_count_value,omitempty"`
	// OnsiteOnWebDetail 商品页浏览数（店铺）
	OnsiteOnWebDetail model.Int64 `json:"onsite_on_web_detail,omitempty"`
	// CostPerOnsiteOnWebDetail 商品页浏览平均成本（店铺）
	CostPerOnsiteOnWebDetail model.Float64 `json:"cost_per_onsite_on_web_detail,omitempty"`
	// OnsiteOnWebDetailRate 商品页浏览率（店铺）
	OnsiteOnWebDetailRate model.Float64 `json:"onsite_on_web_detail_rate,omitempty"`
	// ValuePerOnsiteOnWebDetail 商品页浏览平均价值（店铺）
	ValuePerOnsiteOnWebDetail model.Float64 `json:"value_per_onsite_on_web_detail,omitempty"`
	// TotalOnsiteOnWebDetailValue 商品页浏览价值（店铺）
	TotalOnsiteOnWebDetailValue model.Float64 `json:"total_onsite_on_web_detail_value,omitempty"`
	// OnsiteOnWebCart 加入购物车数（店铺）
	OnsiteOnWebCart model.Int64 `json:"onsite_on_web_cart,omitempty"`
	// CostPerOnsiteOnWebCart 加入购物车平均成本（店铺）
	CostPerOnsiteOnWebCart model.Float64 `json:"cost_per_onsite_on_web_cart,omitempty"`
	// OnsiteOnWebCartRate 加入购物车率（店铺）
	OnsiteOnWebCartRate model.Float64 `json:"onsite_on_web_cart_rate,omitempty"`
	// ValuePerOnsiteOnWebCart 加入购物车平均价值（店铺）
	ValuePerOnsiteOnWebCart model.Float64 `json:"value_per_onsite_on_web_cart,omitempty"`
	// TotalOnsiteOnWebCartValue 加入购物车价值（店铺）
	TotalOnsiteOnWebCartValue model.Float64 `json:"total_onsite_on_web_cart_value,omitempty"`
	//
	// 归因指标
	// VTAConversion 展示归因转化数
	VTAConversion model.Int64 `json:"vta_conversion,omitempty"`
	// CostPerVTAConversion 展示归因转化平均成本
	CostPerVTAConversion model.Float64 `json:"cost_per_vta_conversion,omitempty"`
	// CTAConversion 点击归因转化数
	CTAConversion model.Int64 `json:"cta_conversion,omitempty"`
	// CostPerCTAConversion 点击归因转化平均成本
	CostPerCTAConversion model.Float64 `json:"cost_per_cta_conversion,omitempty"`
	// EngagedViewThroughConversions 深度互动观看归因转化数
	EngagedViewThroughConversions model.Int64 `json:"engaged_view_through_conversions,omitempty"`
	// CostPerEngagedViewThroughConversion 深度互动观看归因转化平均成本
	CostPerEngagedViewThroughConversion model.Float64 `json:"cost_per_engaged_view_through_conversion,omitempty"`
	// VTAAppInstall 展示归因应用安装数
	VTAAppInstall model.Int64 `json:"vta_app_install,omitempty"`
	// CostPerVTAAppInstall 展示归因应用安装平均成本
	CostPerVTAAppInstall model.Float64 `json:"cost_per_vta_app_install,omitempty"`
	// CTAAppInstall 点击归因应用安装数
	CTAAppInstall model.Int64 `json:"cta_app_install,omitempty"`
	// CostPerCTAAppInstall 点击归因应用安装平均成本
	CostPerCTAAppInstall model.Float64 `json:"cost_per_cta_app_install,omitempty"`
	// EVTAAppInstall 深度互动观看归因应用安装数
	EVTAAppInstall model.Int64 `json:"evta_app_install,omitempty"`
	// CostPerEVTAAppInstall 深度互动观看归因应用安装平均成本
	CostPerEVTAAppInstall model.Float64 `json:"cost_per_evta_app_install,omitempty"`
	// VTARegistration 展示归因注册数
	VTARegistration model.Int64 `json:"vta_registration,omitempty"`
	// CostPerVTARegistration 展示归因注册平均成本
	CostPerVTARegistration model.Float64 `json:"cost_per_vta_registration,omitempty"`
	// CTARegistration 点击归因注册数
	CTARegistration model.Int64 `json:"cta_registration,omitempty"`
	// CostPerCTARegistration 点击归因注册平均成本
	CostPerCTARegistration model.Float64 `json:"cost_per_cta_registration,omitempty"`
	// EVTARegistration 深度互动观看归因注册数
	EVTARegistration model.Int64 `json:"evta_registration,omitempty"`
	// CostPerEVTARegistration 深度互动观看归因注册平均成本
	CostPerEVTARegistration model.Float64 `json:"cost_per_evta_registration,omitempty"`
	// VTAPurchase 展示归因付费数
	VTAPurchase model.Int64 `json:"vta_purchase,omitempty"`
	// CostPerVTAPurchase 展示归因付费平均成本
	CostPerVTAPurchase model.Float64 `json:"cost_per_vta_purchase,omitempty"`
	// CTAPurchase 点击归因付费数
	CTAPurchase model.Int64 `json:"cta_purchase,omitempty"`
	// CostPerCTAPurchase 点击归因付费平均成本
	CostPerCTAPurchase model.Float64 `json:"cost_per_cta_purchase,omitempty"`
	// EVTAPurchase 深度互动观看归因付费数
	EVTAPurchase model.Int64 `json:"evta_purchase,omitempty"`
	// CostPerEVTAPurchase 深度互动观看归因付费平均成本
	CostPerEVTAPurchase model.Float64 `json:"cost_per_evta_purchase,omitempty"`
	// VTACompletePayment 展示归因付费数（网站）
	VTACompletePayment model.Int64 `json:"vta_complete_payment,omitempty"`
	// CostPerVTAPaymentsCompleted 展示归因付费平均成本（网站）
	CostPerVTAPaymentsCompleted model.Float64 `json:"cost_per_vta_payments_completed,omitempty"`
	// EVTAPaymentsCompleted 深度互动观看归因付费数（网站）
	EVTAPaymentsCompleted model.Int64 `json:"evta_payments_completed,omitempty"`
	// CostPerEVTAPaymentsCompleted 深度互动观看归因付费平均成本（网站）
	CostPerEVTAPaymentsCompleted model.Float64 `json:"cost_per_evta_payments_completed,omitempty"`
	// VTACompletePaymentValue 展示归因付费价值
	VTACompletePaymentValue model.Float64 `json:"vta_complete_payment_value,omitempty"`
	// VTACompletePaymentROAS 展示归因付费ROAS
	VTACompletePaymentROAS model.Float64 `json:"vta_complete_payment_roas,omitempty"`
}
