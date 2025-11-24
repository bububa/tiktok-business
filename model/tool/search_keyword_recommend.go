package tool

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SearchKeywordRecommendRequest 获取推荐搜索关键词 API Request
type SearchKeywordRecommendRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// SearchQueries 搜索词列表。
	// 您可以指定商品、服务或品牌名称作为搜索词。
	// 最大数量：5。
	// 每个搜索词的长度限制：80 个字符。
	// 仅支持英文单词。
	// 不支持表情符号以及特殊字符（例如 #）。
	SearchQueries []string `json:"search_queries,omitempty"`
	// Regions 国家或地区代码列表。
	// 目前，仅支持值 US。
	Regions []string `json:"regions,omitempty"`
	// OrderField 排序逻辑。
	// 枚举值：
	// RELEVANCE：按照相关性排序。
	// MONTHLY_SEARCHES：按照返回参数 monthly_searches 的值排序。
	// 默认值：RELEVANCE。
	OrderField string `json:"order_field,omitempty"`
	// OrderType 仅当 order_field 为 MONTHLY_SEARCHES 时生效。
	// 排序顺序。
	// 枚举值：
	// ASC：升序。
	// DESC：降序。
	// 默认值：DESC。
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// TotalSize 要返回的推荐关键词的数量。
	// 取值范围：1-1,000。
	// 默认值：1,000。
	TotalSize int `json:"total_size,omitempty"`
	// Page 当前页数。
	// 取值范围：≥1。
	// 默认值： 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 取值范围：1-1,000。
	// 默认值：50。
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r *SearchKeywordRecommendRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("search_queries", string(util.JSONMarshal(r.SearchQueries)))
	values.Set("regions", string(util.JSONMarshal(r.Regions)))
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.TotalSize > 0 {
		values.Set("total_size", strconv.Itoa(r.TotalSize))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// SearchKeywordRecommendResponse 获取推荐搜索关键词 API Response
type SearchKeywordRecommendResponse struct {
	model.BaseResponse
	Data *SearchKeywordRecommendResult `json:"data,omitempty"`
}

type SearchKeywordRecommendResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// RecommendedKeywords 推荐的搜索关键词列表
	RecommendedKeywords []RecommendedSearchKeyword `json:"recommended_search_keywords,omitempty"`
}

// RecommendedSearchKeyword 推荐的搜索关键词
type RecommendedSearchKeyword struct {
	// RecommendedKeyword 推荐的搜索关键词
	RecommendedKeyword string `json:"recommended_keyword,omitempty"`
	// MonthlySearches 预测月搜索量。
	// 预测以历史数据为基础，仅供参考。
	MonthlySearches int64 `json:"monthly_searches,omitempty"`
}
