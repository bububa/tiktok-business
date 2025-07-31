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
	Clicks model.Float64 `json:"clicks,omitempty"`
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
	//
	//互动指标
	// 互动指标衡量受众与您的广告的互动
	// VideoPlayActions 视频播放量
	// 视频开始播放的次数。每次视频展示的播放次数会单独统计，且不包括重播次数。
	VideoPlayActions model.Int64 `json:"video_play_actions,omitempty"`
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
	// 归因于广告的去重应用内加入心愿单操作次数。
	AddToWishlist model.Int64 `json:"add_to_wishlist,omitempty"`
	// CostPerAddToWishlist 去重加入心愿单平均成本（应用）
	// 归因于广告的去重应用内加入心愿单操作平均成本。
	CostPerAddToWishlist model.Float64 `json:"cost_per_add_to_wishlist,omitempty"`
	// AddToWishlistRate 去重加入心愿单率（应用）
	// 归因于广告的应用内加入心愿单操作占应用安装总量的百分比。
	AddToWishlistRate model.Float64 `json:"add_to_wishlist_rate,omitempty"`
	// TotalAddToWishlist 加入心愿单次数（应用）
	// 归因于广告的应用内加入心愿单操作次数。
	TotalAddToWishlist model.Int64 `json:"total_add_to_wishlist,omitempty"`
	// CostPerTotalAddToWishlist 加入心愿单平均成本（应用）
	// 归因于广告的应用内加入心愿单操作平均成本。
	CostPerTotalAddToWishlist model.Float64 `json:"cost_per_total_add_to_wishlist,omitempty"`
	// ValuePerTotalAddToWishlist 加入心愿单平均价值（应用）
	// 归因于广告的应用内加入心愿单操作平均价值。
	ValuePerTotalAddToWishlist model.Float64 `json:"value_per_total_add_to_wishlist,omitempty"`
	// TotalAddToWishlistValue 加入心愿单价值（应用）
	// 归因于广告的应用内加入心愿单操作总价值。
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
	//
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
	// Form 提交表单次数（网站）
	// 归因于广告的网站提交表单操作次数。
	Form model.Int64 `json:"form,omitempty"`
	// CostPerForm 提交表单平均成本（网站）
	// 归因于广告的网站提交表单操作的平均成本。
	CostPerForm model.Float64 `json:"cost_per_form,omitempty"`
	// FormRate 去重提交表单率（网站） (%)
	// 归因于广告的网站提交表单的平均价值。
	FormRate model.Float64 `json:"form_rate,omitempty"`
	// ValuePerForm 提交表单平均价值（网站）
	// 归因于广告的网站提交表单操作的总价值。
	ValuePerForm model.Float64 `json:"value_per_form,omitempty"`
	// TotalFormValue 提交表单价值（网站）
	// 归因于广告的网站提交表单操作次数占总点击数的百分比。
	TotalFormValue model.Float64 `json:"total_form_value,omitempty"`
}
