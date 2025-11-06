package report

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest TTCM 订单表现报表 API Request
type GetRequest struct {
	// TCMAccountID TikTok Creator Marketplace 账户 ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// OrderIDs 订单ID列表。
	// 注意：若您传入不属于tcm_account_id的订单ID，则该订单会被忽略
	OrderIDs []string `json:"order_ids,omitempty"`
	// StartDate 查询开始日期。检索结果包含开始日期的数据。示例："2020-01-01"。当need_daily_stats 设置为 True时必填
	StartDate string `json:"start_date,omitempty"`
	// EndDate 查询结束日期。检索结果包含结束日期的数据。示例："2020-01-01"。当need_daily_stats 设置为 True时必填。
	// 注意：查询结束日期（end_date）应为开始日期（start_date）后60天以内的一个日期，即 0 <= end_date - start_date <= 60 天。
	EndDate string `json:"end_date,omitempty"`
	// NeedDailyStats 是否在报表中包含每日细分数据。枚举值： True，False。默认值： False
	NeedDailyStats bool `json:"need_daily_stats,omitempty"`
	// Metrics 要查询的数据指标。
	// 若您在请求中传入metrics字段，则所有请求字段均为必填。
	// 返回中metrics下包含的指标字段即为本字段的允许值，如未指定，将返回以下默认指标：
	// engagement_count
	// engagement_rate
	// video_views
	// likes
	// shares
	// comments
	// reach
	// total_play_time
	// average_view_time
	// two_seconds_views
	// six_seconds_views
	// video_completion_rate
	Metrics []string `json:"metrics,omitempty"`
	// Page 当前页码。默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小。默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("tcm_account_id", r.TCMAccountID)
	values.Set("order_ids", string(util.JSONMarshal(r.OrderIDs)))
	if r.StartDate != "" {
		values.Set("start_date", r.StartDate)
	}
	if r.EndDate != "" {
		values.Set("end_date", r.EndDate)
	}
	if r.NeedDailyStats {
		values.Set("need_daily_stats", "true")
	}
	if len(r.Metrics) > 0 {
		values.Set("metrics", string(util.JSONMarshal(r.Metrics)))
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

// GetResponse TTCM 订单表现报表 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Insights 数据列表
	Insights []Insight `json:"insights,omitempty"`
}
