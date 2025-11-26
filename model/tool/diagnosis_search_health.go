package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DiagnosisSearchHealthRequest 获取搜索广告推广系列的健康状况诊断 API Request
type DiagnosisSearchHealthRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdgroupID 广告组 ID。
	// 若要获取现有广告组的诊断信息，需指定此字段。
	// 若要获取搜索广告推广系列中的广告组 ID，可使用 /adgroup/get/ 接口，并将筛选条件字段 campaign_ids 设置为搜索广告推广系列的 ID。
	// 注意：
	// 若同时指定 adgroup id 和 ad_ids，adgroup id将被忽略。
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdIDs 广告 ID 列表。
	// 若要获取现有广告的诊断信息，需指定此字段。
	// 广告 ID 均需属于同一广告组。
	// 最大数量：50。
	// 若要获取搜索广告组中的广告 ID，可使用 /ad/get/ 接口，并将筛选条件字段 adgroup_ids 设置为搜索广告推广系列中搜索广告组的 ID。
	// 注意：
	// 若同时指定 adgroup id 和 ad_ids，adgroup id将被忽略。
	AdIDs []string `json:"ad_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r *DiagnosisSearchHealthRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if r.AdgroupID != "" {
		values.Set("adgroup_id", r.AdgroupID)
	}
	if len(r.AdIDs) > 0 {
		values.Set("ad_ids", string(util.JSONMarshal(r.AdIDs)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// DiagnosisSearchHealthResponse 获取搜索广告推广系列的健康状况诊断 API Response
type DiagnosisSearchHealthResponse struct {
	model.BaseResponse
	Data *DiagnosisSearchHealthResult `json:"data,omitempty"`
}

type DiagnosisSearchHealthResult struct {
	// SearchHealthStatus 整体搜索健康状况。此字段对广告组或广告的设置情况的概述。
	// 枚举值：
	// GOOD：良好。您的设置有助于提升广告成效。
	// NEED_IMPROVEMENT：需要改进。
	// NO_DATA：暂无数据。
	SearchHealthStatus enum.SearchHealthStatus `json:"search_health_status,omitempty"`
	// SearchVolumn 关键词搜索量的详细指标
	SearchVolumn *DiagnosisSearchVolumn `json:"search_volumn,omitempty"`
	// KeywordRelevance 关键词与广告组或广告相关性的详情
	KeywordRelevance *DiagnosisKeywordRelevance `json:"keyword_relevance,omitempty"`
	// BidBudget 出价和预算诊断的详情
	BidBudget *DiagnosisBidBudget `json:"bid_budget,omitempty"`
}

// DiagnosisSearchVolumn 关键词搜索量的详细指标
type DiagnosisSearchVolumn struct {
	// DiagnosisResult 搜索量分析的结果。
	// 枚举值：
	// HIGH：高。
	// MEDIUM：中。
	// LOW：低。
	// INVALID：无效。
	DiagnosisResult string `json:"diagnosis_result,omitempty"`
	// TotalMonthlySearches 关键词每月获得的预估搜索次数。
	// 该数值会随匹配类型的设置而有所变化。广泛匹配有助于最大化搜索量。
	TotalMonthlySearches int64 `json:"total_monthly_searches,omitempty"`
	// TotalKeywordCount 搜索关键词总数
	TotalKeywordCount int64 `json:"total_keyword_count,omitempty"`
	// TotalReleventKeywordCount 与广告组或广告相关的搜索关键词总数
	TotalReleventKeywordCount int64 `json:"total_relevent_keyword_count,omitempty"`
}

// DiagnosisKeywordRelevance 关键词与广告组或广告相关性的详情
type DiagnosisKeywordRelevance struct {
	// AdgroupID 广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdID 广告 ID
	AdID string `json:"ad_id,omitempty"`
	// KeywordRelevanceStatus 关键词相关性分析的结果。
	// 枚举值：
	// TO_BE_CALCULATED：待计算。
	// PARTIALLY_RELEVANT：部分相关。
	// RELEVANT：相关。
	// IRRELEVANT：不相关。
	KeywordRelevanceStatus enum.KeywordRelevanceStatus `json:"keyword_relevance_status,omitempty"`
	// RelevantKeywordCount 相关的搜索关键词的数量
	RelevantKeywordCount int `json:"relevant_keyword_count,omitempty"`
	// RelevantKeywords 相关的搜索关键词列表
	RelevantKeywords []string `json:"relevant_keywords,omitempty"`
	// IrrelevantKeywordCount 不相关的搜索关键词的数量
	IrrelevantKeywordCount int `json:"irrelevant_keyword_count,omitempty"`
	// IrrelevantKeywords 不相关的搜索关键词列表
	IrrelevantKeywords []string `json:"irrelevant_keywords,omitempty"`
}

// DiagnosisBidBudget 出价和预算诊断的详情
type DiagnosisBidBudget struct {
	// BidBudgetStatus 出价和预算分析的结果。
	// 枚举值：
	// GOOD：良好。
	// LOW_BID_AND_BUDGET：出价和预算过低。
	// LOW_BUDGET：预算较低。
	// LOW_BID：出价较低。
	// NO_DATA：暂无数据。
	BidBudgetStatus enum.BidBudgetStatus `json:"bid_budget_status,omitempty"`
	// BidSuggestStatus 针对出价的诊断结果。
	// 枚举值：
	// GOOD：良好。
	// LOW：出价较低。
	// NO_DATA：暂无数据。
	BidSuggestStatus enum.BidSuggestStatus `json:"bid_suggest_status,omitempty"`
	// SuggestedValue 建议出价。
	// 根据历史数据，将出价至少增加到此建议值，广告组或广告可能会获得更多转化
	SuggestedValue model.Float64 `json:"suggested_value,omitempty"`
}
