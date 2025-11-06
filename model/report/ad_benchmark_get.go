package report

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdBenchmarkGetRequest 获取广告表现对比数据 API Request
type AdBenchmarkGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CompareTimeWindow 对比的时间窗口。 枚举值：7 (默认), 14, 30, 60。
	CompareTimeWindow int `json:"compare_time_window,omitempty"`
	// Dimensions 对比维度。可以传入1-4个维度。枚举值： LOCATION（地域）, AD_CATEGORY（兴趣分类）, EXTERNAL_ACTION（转化事件）, PLACEMENT（版位）。
	Dimensions []enum.AdBenchmarkDimension `json:"dimensions,omitempty"`
	// MetricsFields 你想获取的指标。默认返回所有指标。支持的指标可见支持指标表格。
	MetricsFields []string `json:"metrics_fields,omitempty"`
	// Filtering 筛选条件。你只能传入三个条件中的一个条件，且必须传入一个条件。
	Filtering *AdBenchmarkGetFilter `json:"filtering,omitempty"`
	// SortField 排序依据的字段。默认根据创建时间(CREATE_TIME)排序
	SortField string `json:"sort_field,omitempty"`
	// SortType 排序类型。枚举值: ASC (升序，从最早创建的开始), DES (降序，从最晚创建的开始)
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Page 页号
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

type AdBenchmarkGetFilter struct {
	// AdIDs 你想获取对比数据的广告ID列表
	AdIDs []string `json:"ad_ids,omitempty"`
	// AdgroupIDs 你想获取对比数据的广告组ID列表
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// CampaignIDs 你想获取对比数据的广告推广系列ID列表
	CampaignIDs []string `json:"campaign_ids,omitempty"`
}

// Encode implements GetRequest
func (r *AdBenchmarkGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if r.CompareTimeWindow > 0 {
		values.Set("compare_time_window", strconv.Itoa(r.CompareTimeWindow))
	}
	values.Set("dimensions", string(util.JSONMarshal(r.Dimensions)))
	if len(r.MetricsFields) > 0 {
		values.Set("metrics_fields", string(util.JSONMarshal(r.MetricsFields)))
	}
	values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	if r.SortField != "" {
		values.Set("sort_field", r.SortField)
	}
	if r.SortType != "" {
		values.Set("sort_type", string(r.SortType))
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

// AdBenchmarkGetResponse 获取广告表现对比数据 API Response
type AdBenchmarkGetResponse struct {
	model.BaseResponse
	Data *AdBenchmarkGetResult `json:"data,omitempty"`
}

type AdBenchmarkGetResult struct {
	// CompareDate 对比的日期
	CompareDate string `json:"compare_date,omitempty"`
	// List 结果列表
	List []AdBenchmark `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type AdBenchmark struct {
	// Dimensions 所有请求的维度组合数据
	Dimensions *AdBenchmarkInfo `json:"dimensions,omitempty"`
	// Metrics 所有请求的指标数据
	Metrics *Metrics `json:"metrics,omitempty"`
}

type AdBenchmarkInfo struct {
	// AdID 广告ID。当dimensions包含 ad_id 时返回
	AdID string `json:"ad_id,omitempty"`
	// Location 2个字母的国家或地区代码
	Location string `json:"location,omitempty"`
	// Placement 投放版位，详见枚举值-广告管理-版位
	Placement enum.Placement `json:"placement,omitempty"`
	// AdCategory 二级兴趣分类ID。详细信息，请参考 获取兴趣分类列表。若需确定游戏相关兴趣分类ID对应的兴趣分类名称，请查看兴趣分类。
	AdCategory string `json:"ad_category,omitempty"`
	// ExternalAction 广告的外部事件/转化事件
	ExternalAction enum.OptimizationEvent `json:"external_action,omitempty"`
}
