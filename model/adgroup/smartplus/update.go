package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest Update an Upgraded Smart+ Ad Group API Request
type UpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdgroupID 广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName 广告组名称
	AdgroupName string `json:"adgroup_name,omitempty"`
	// BidPrice Valid when the following conditions are met:
	// optimization_goal is CLICK.
	// bid_type is BID_TYPE_CUSTOM.
	// The target cost per click. The system will aim to get the most results while keeping the average cost per result around or lower than the specified amount.
	// bid_price needs to be lower than budget. See Bidding-Bidding limits to learn more about the bid verification mechanism.
	// Important: When multiple ad groups are present within your campaign, updating the bid_price for one ad group will automatically update the bid_price for all other ad groups under the same campaign with the new value.
	// Note: When Campaign Budget Optimization (CBO) is enabled (budget_optimize_on is true) at the campaign level, this setting, if specified, must be the same across all ad groups within the same campaign.
	BidPrice float64 `json:"bid_price,omitempty"`
	// ConversionBidPrice Valid when the following conditions are met:
	// optimization_goal is CONVERT, TRAFFIC_LANDING_PAGE_VIEW, INSTALL, or IN_APP_EVENT.
	// bid_type is BID_TYPE_CUSTOM.
	// The target cost per conversion or cost per landing page view. The system will aim to get the most results while keeping the average cost per result around or lower than the specified amount.
	// conversion_bid_price needs to be lower than budget. See Bidding-Bidding limits to learn more about the bid verification mechanism.
	// Important: When multiple ad groups are present within your campaign, updating the conversion_bid_price for one ad group will automatically update the bid_price for all other ad groups under the same campaign with the new value.
	// Note:When CBO is enabled (budget_optimize_on is true) at the campaign level, this setting, if specified, must be the same across all ad groups within the same campaign.
	ConversionBidPrice float64 `json:"conversion_bid_price,omitempty"`
	// RoasBid Valid when deep_bid_type is VO_MIN_ROAS.
	// Target ROAS for Value-Based Optimization.
	// Value range: 0.01-1,000.
	// Important: When multiple ad groups are present within your campaign, updating the roas_bid for one ad group will automatically update the bid_price for all other ad groups under the same campaign with the new value.
	// Note: When CBO is enabled (budget_optimize_on is true) at the campaign level, this setting, if specified, must be the same across all ad groups within the same campaign.
	RoasBid float64 `json:"roas_bid,omitempty"`
	// CommentDisabled 是否允许用户在TikTok上评论您的广告
	CommentDisabled *bool `json:"comment_disabled,omitempty"`
	// ShareDisabled 本广告组中的广告是否禁止分享到第三方平台
	ShareDisabled *bool `json:"share_disabled,omitempty"`
	// ScheduleType 广告投放时间类型。
	// 枚举值: SCHEDULE_FROM_NOW，SCHEDULE_START_END。如果您选择了SCHEDULE_START_END，您需要明确投放开始和结束时间。 如果您选择了SCHEDULE_FROM_NOW，您只需要明确投放开始时间。
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime 广告投放起始时间(UTC+0)。
	// 格式：YYYY-MM-DD HH:MM:SS。
	ScheduleStartTime *model.DateTime `json:"schedule_start_time,omitempty"`
	// ScheduleEndTime 广告投放结束时间(UTC+0)。
	// 格式： YYYY-MM-DD HH:MM:SS
	ScheduleEndTime *model.DateTime `json:"schedule_end_time,omitempty"`
	// Dapparting 广告投放安排。格式为48x7的字符串。字符只能为0或者1。0代表不投放，1代表投放。每个字符对应一个30分钟的时间段。第一个字符对应的是周一的凌晨0:01-0:30，第二个字符对应0:31-1:00，依次类推。最后一个字符代表周日23:31-0:00。
	// 注意：
	// 不设置，全部设置为0，或者全部都设置为1都代表了要进行全时投放。
	Dayparting string `json:"dayparting,omitempty"`
	// IsHfss 广告推广对象是否是 HFSS 食品（高脂、高盐、高糖食品）请注意，欧洲国家不允许向未成年人推送 HFSS 食品广告
	IsHfss *bool `json:"is_hfss,omitempty"`
	// IsLhfCompliance 推广内容是否符合 LHF（较不健康食品）合规要求。
	// 将 is_lhf_compliance 设置为 true，即表示您确认在英国 TikTok 上推广的任何食品或饮料均符合 2024 年较不健康食品法规以及其他所有适用法律。
	// 支持的值：true，false。
	// 默认值：false。
	// 注意：自 2026 年 1 月 1 日起，当广告定向到英国地域且 is_hfss 为 true 时，is_lhf_compliance 必填，且必须设置为 true。
	IsLhfCompliance *bool `json:"is_lhf_compliance,omitempty"`
	// TargetingOptimizationMode Audience targeting optimization mode.
	// Enum values:
	// MANUAL: Custom targeting. You can use custom targeting settings to precisely control who sees your ads. This may limit delivery and impact campaign performance.
	// AUTOMATIC: Automatic targeting. You can use automatic targeting to leverage real-time data and machine learning to target audiences most likely to engage with your ads.
	TargetingOptimziationMode enum.TargetingOptimizationMode `json:"targeting_optimization_mode,omitempty"`
	// SuggestionAudienceEnabled Whether to enable audience suggestions.
	// Audience suggestions guide automatic targeting by choosing additional audience settings. These serve as suggestions only, and delivery to those audiences is not guaranteed.
	// Supported values: true, false.
	SuggestionAudienceEnabled *bool `json:"suggestion_audience_enabled,omitempty"`
	// TargetingSpec Targeting settings
	TargetingSpec *Targeting `json:"targeting_spec,omitempty"`
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
	Budget *float64 `json:"budget,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse Update an Upgraded Smart+ Ad Group API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *Adgroup `json:"data,omitempty"`
}
