package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SearchKeywordIdeaKeywordRequest 发现新关键词 API Request
type SearchKeywordIdeaKeywordRequest struct {
	// AdvertiserID advertiser id
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Keywords 关键词列表。
	// 你可以传入与你的业务密切相关的商品或服务。
	// 最大数量：10。
	// 每个关键词的长度限制：80 个字符。
	// 仅支持英文单词。
	// 不支持表情符号和特殊符号（例如 #）。
	Keywords []string `json:"keywords,omitempty"`
	// OrderField 排序逻辑。
	// 枚举值：
	// AVG_MONTHLY_SEARCHES：按（由返回参数 avg_monthly_searches_lower 和 avg_monthly_searches_upper 定义的）月均搜索量范围对结果进行排序。
	// THREE_MONTH_CHANGE：按返回参数 three_month_change 对结果进行排序。
	// YEAR_OVER_YEAR_CHANGE：按返回参数 year_over_year_change 对结果进行排序。
	// COMPETITION：按返回参数 competition 对结果进行排序。
	// 默认值：AVG_MONTHLY_SEARCHES。
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序的顺序。
	// 枚举值：
	// ASC：升序。
	// DESC：降序。
	// 默认值：DESC。
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// BrandType 品牌类型。
	// 枚举值：
	// BRAND：品牌。
	// NON_BRAND：非品牌。
	// ALL：全部。
	// 默认值：ALL。
	BrandType string `json:"brand_type,omitempty"`
	// CountryCodes 定向国家/地区的地域代码。
	// 最大数量：1。
	// 默认值：US。
	// 目前仅支持 US 值。
	CountryCodes []string `json:"country_codes,omitempty"`
}

// Encode implements GetRequest
func (r *SearchKeywordIdeaKeywordRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("keywords", string(util.JSONMarshal(r.Keywords)))
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.BrandType != "" {
		values.Set("brand_type", r.BrandType)
	}
	if len(r.CountryCodes) > 0 {
		values.Set("country_codes", string(util.JSONMarshal(r.CountryCodes)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// SearchKeywordIdeaKeywordResponse 发现新关键词 API Response
type SearchKeywordIdeaKeywordResponse struct {
	model.BaseResponse
	Data *SearchKeywordIdeaKeywordResult `json:"data,omitempty"`
}

type SearchKeywordIdeaKeywordResult struct {
	// TotalSearchVolumeLower 所有建议关键词 (keywords) 的最低总搜索量。
	TotalSearchVolumeLower int64 `json:"total_search_volume_lower,omitempty"`
	// TotalSearchVolumeUpper 所有建议关键词 (keywords) 的最高总搜索量
	TotalSearchVolumeUpper int64 `json:"total_search_volume_upper,omitempty"`
	// Keywords 基于传入的定向国家或地区 (country_codes)、与传入的一个或多个关键词相关的建议关键词列表。
	Keywords []IdeaKeyword `json:"keywords,omitempty"`
}

type IdeaKeyword struct {
	// RecommendedKeyword 关键词。
	// 与传入的一个或多个关键词相关的建议关键词。
	RecommendedKeyword string `json:"recommended_keyword,omitempty"`
	// AvgMonthlySearchesLower 月均搜索量（最低）。
	// 根据传入的定向国家或地区，近 1 个月内某个关键词及其类似词的最低月均搜索量。
	AvgMonthlySearchesLower int64 `json:"avg_monthly_searches_lower,omitempty"`
	// AvgMonthlySearchesUpper 月均搜索量（最高）。
	// 根据传入的定向国家或地区，近 1 个月内某个关键词及其类似词的最高月均搜索量
	AvgMonthlySearchesUpper int64 `json:"avg_monthly_searches_upper,omitempty"`
	// ThreeMonthChange 近 3 个月变化。
	// 将近 1 个月的月搜索量与 2 个月前的月搜索量进行对比，得出此数据。
	ThreeMonthChange float64 `json:"three_month_change,omitempty"`
	// YearOverYearChange 同比变化。
	// 将上月的月搜索量与去年同月的月搜索量进行对比，得出此数据。
	YearOverYearChange float64 `json:"year_over_year_change,omitempty"`
	// Competition 竞争程度。
	// 根据上月对某个关键词出价的去重广告主数量，表示该关键词的广告位竞争程度。
	// 枚举值：
	// HIGH：高。
	// MEDIUM：中。
	// LOW：低。
	Competition string `json:"competition,omitempty"`
	// EstimatedCPCLower 预估 CPC（最低）。
	// 上月广告主将该关键词投放在传入的定向国家/地区的搜索页面首个广告位所支付的最低 CPC。具体的关键词 CPC 可能因实际情况而异。
	EstimatedCPCLower float64 `json:"estimated_cpc_lower,omitempty"`
	// EstimatedCPCUpper 预估 CPC（最高）。
	// 上月广告主将该关键词投放在传入的定向国家/地区的搜索页面首个广告位所支付的最高 CPC。具体的关键词 CPC 可能因实际情况而异。
	EstimatedCPCUpper float64 `json:"estimated_cpc_upper,omitempty"`
}
