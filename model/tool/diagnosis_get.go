package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DiagnosisGetRequest 获取广告组诊断 API Request
type DiagnosisGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Filtering 筛选条件
	Filtering *DiagnosisGetFilter `json:"filtering,omitempty"`
}

// DiagnosisGetFilter 筛选条件
type DiagnosisGetFilter struct {
	// AdgroupIDs 您想要获取诊断的广告组 ID 列表。
	// 最大数量：20。
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// IssueCategory 问题类别。
	// 允许值：
	// CREATIVE：创意相关问题。
	// BID_AND_BUDGET：出价和预算相关问题。
	// EVENT_TRACK：事件追踪相关问题。
	// 若传入多个值，将返回属于指定的问题类别的诊断。
	// 若未传入本字段，将返回存在的所有问题类别的诊断。
	IssueCategory []enum.DiagnosisIssueCategory `json:"issue_category,omitempty"`
}

// Encode implements GetRequest interface
func (r *DiagnosisGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// DiagnosisGetResponse 获取广告组诊断 API Response
type DiagnosisGetResponse struct {
	model.BaseResponse
	Data struct {
		// Results 广告组列表及诊断信息
		Results []DiagnosisGetResult `json:"results,omitempty"`
	} `json:"data"`
}

// DiagnosisGetResult 诊断信息
type DiagnosisGetResult struct {
	// AdgroupID 广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName 广告组名称
	AdgroupName string `json:"adgroup_name,omitempty"`
	// Diagnosis 该广告组的诊断信息
	Diagnosis *Diagnosis `json:"diagnosis,omitempty"`
}

// Diagnosis 广告组的诊断信息
type Diagnosis struct {
	// DiagnosisTime 请求诊断数据的日期及时间（UTC+0 时间），格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2024-01-01 00:00:01"。
	DiagnosisTime model.DateTime `json:"diagnosis_time,omitzero"`
	// Suggestions 对于本广告组的建议
	Suggestions []Suggestion `json:"suggestions,omitempty"`
}

// Suggestion 广告组的建议
type Suggestion struct {
	// Creative 针对创意的建议
	Creative []CreativeSuggestion `json:"creative,omitempty"`
	// BidAndBudget 针对出价及预算的建议
	BidAndBudget []BidAndBudgetSuggestion `json:"bid_and_budget,omitempty"`
	// EventTrack 针对事件追踪的建议
	EventTrack []EventTrackSuggestion `json:"event_track,omitempty"`
}

// CreativeSuggestion 针对创意的建议
type CreativeSuggestion struct {
	// SuggstionTime 诊断出问题并生成建议的日期及时间（UTC+0 时间），格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2024-01-01 00:00:01"。
	SuggestionTime model.DateTime `json:"suggestion_time,omitzero"`
	// SuggestionID 建议 ID
	SuggestionID string `json:"suggestion_id,omitempty"`
	// Vid 视频 ID。
	// 示例：v07033g50000c1irpg9l3c4i8hcdp7hg。
	Vid string `json:"vid,omitempty"`
	// Name 广告名称
	Name string `json:"name,omitempty"`
	// AdID 广告 ID
	AdID string `json:"ad_id,omitempty"`
	// IssueSuggestion 问题建议。
	// 枚举值:
	// NOBGM: 视频无背景音乐。您需为视频添加背景音乐。可使用 /creative/quick_optimization/create/为视频创建一键优化任务。
	// VIDEO_LENGTH：视频时长过短。需编辑视频增加时长。可使用 /creative/quick_optimization/create/为视频创建一键优化任务。
	// VIDEO_RESOLUTION：视频分辨率过低。需替换为高分辨率视频。
	IssueSuggestion enum.DiagnosisIssueSuggestion `json:"issue_suggestion,omitempty"`
}

// BidAndBudgetSuggestion 针对出价及预算的建议
type BidAndBudgetSuggestion struct {
	// SuggstionTime 诊断出问题并生成建议的日期及时间（UTC+0 时间），格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2024-01-01 00:00:01"。
	SuggestionTime model.DateTime `json:"suggestion_time,omitzero"`
	// SuggestionID 建议 ID
	SuggestionID string `json:"suggestion_id,omitempty"`
	// IssueSuggestion 问题建议。枚举值:
	// SUGGEST_BID: 建议更改出价
	// SUGGEST_BUDGET: 建议更改预算
	// NOBID_SWITCH: 切换为最大投放量出价策略
	// BUDGET_EDR：预算预估成效建议。若想了解预估成效详情，可查看预估成效。
	// BID_EDR：出价预估成效建议。
	IssueSuggestion enum.DiagnosisIssueSuggestion `json:"issue_suggestion,omitempty"`
	// Bid 在issue_suggestion为 SUGGEST_BID时返回。
	// 当前出价。
	Bid float64 `json:"bid,omitempty"`
	// Budget 在issue_suggestion为SUGGEST_BID, SUGGEST_BUDGET, 或NOBID_SWITCH时返回。
	// 当前预算。
	Budget float64 `json:"budget,omitempty"`
	// SuggestBid 在issue_suggestion为SUGGEST_BID时返回。
	// 建议出价。
	// SuggestBudget 在issue_suggestion为SUGGEST_BUDGET 或NOBID_SWITCH时返回。
	// 建议预算。
	SuggestBudget float64 `json:"suggest_budget,omitempty"`
	// CostFloor 在issue_suggestion为SUGGEST_BID时返回。
	// 预估成本下限。
	CostFloor float64 `json:"cost_floor,omitempty"`
	// BidEdrInfo 若issue_suggestion非BID_EDR ，本字段的值将为空列表（[]）。
	// 针对该广告组的出价预估成效建议。
	// 9 组推荐出价及预估成效数据，推荐出价涨幅从-10%逐步增加5%，最高推荐增加30%出价。
	BidEdrInfo *BidEdrInfo `json:"bid_edr_info,omitempty"`
	// BudgetEdrInfo 若issue_suggestion非BUDGET_EDR ，本字段的值将为空列表（[]）。
	// 针对该广告组的预算预估成效建议。
	// 15 组推荐预算及预估成效数据，推荐预算涨幅从 -20% 逐步增加5%，最高推荐增加 50% 预算 。
	BudgetEdrInfo *BudgetEdrInfo `json:"budget_edr_info,omitempty"`
}

// BidEdrInfo  针对该广告组的出价预估成效建议
type BidEdrInfo struct {
	// RecommendedBid 推荐出价。
	// 若需更新合适的出价，请使用/adgroup/update/接口，将出价传给conversion_bid_price字段 ，在当天（广告主账户时区）进行更新。
	RecommendedBid float64 `json:"recommended_bid,omitempty"`
	// BidIncreaseRatio 出价增加的预估比例
	BidIncreaseRatio float64 `json:"bid_increase_ratio,omitempty"`
	// EstimatedCost 预估花费
	EstimatedCost float64 `json:"estimated_cost,omitempty"`
	// CostUplift 预估花费提升量
	CostUplift float64 `json:"cost_uplift,omitempty"`
	// CostUpliftRatio 预估花费提升比例
	CostUpliftRatio float64 `json:"cost_uplift_ratio,omitempty"`
}

// BudgetEdrInfo 针对该广告组的预算预估成效建议
type BudgetEdrInfo struct {
	// RecommendBudget 推荐预算。若需更新合适的预算，请使用 /adgroup/budget/update/接口在当天（广告主账户时区）更新。
	RecommendBudget float64 `json:"recommend_budget,omitempty"`
	// BudgetIncreaseRatio 预算增加的预估比例
	BudgetIncreaseRatio float64 `json:"budget_increase_ratio,omitempty"`
	// EstimatedConversion 预估转化量
	EstimatedConversion int64 `json:"estimated_conversion,omitempty"`
	// ConversionUplift 预估转化提升量
	ConversionUplift int64 `json:"conversion_uplift,omitempty"`
	// ConversionUpliftRatio 预估转化提升比例
	ConversionUpliftRatio float64 `json:"conversion_uplift_ratio,omitempty"`
}

// EventTrackSuggestion 针对事件追踪的建议
type EventTrackSuggestion struct {
	// SuggstionTime 诊断出问题并生成建议的日期及时间（UTC+0 时间），格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2024-01-01 00:00:01"。
	SuggestionTime model.DateTime `json:"suggestion_time,omitzero"`
	// SuggestionID 建议 ID
	SuggestionID string `json:"suggestion_id,omitempty"`
	// IssueSuggestion 问题建议。
	// 枚举值：
	// PIXEL：该 Pixel 过去 7 日内无任何活动。请检查您的 Pixel 设置。
	IssueSuggestion enum.DiagnosisIssueSuggestion `json:"issue_suggestion,omitempty"`
	// PixelID Pixel ID
	PixelID string `json:"pixel_id,omitempty"`
	// PixelCode Pixel 代码
	PixelCode string `json:"pixel_code,omitempty"`
}
