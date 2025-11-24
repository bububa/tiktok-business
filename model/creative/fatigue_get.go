package creative

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// FatigueGetRequest 获取创意疲劳检测结果 API Request
type FatigueGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdID 广告ID，用于获取对应的创意疲劳检测结果
	AdID string `json:"ad_id,omitempty"`
	// Filtering 筛选条件
	Filtering *FatigueGetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值：1。
	// 取值范围：≥ 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值：10。
	// 取值范围：[1, 500]
	PageSize int `json:"page_size,omitempty"`
}

// FatigueGetFilter 筛选条件
type FatigueGetFilter struct {
	// StartDate 查询开始时间（闭区间）。格式：YYYY-MM-DD（广告主账户时区）。
	// 只允许指定过去60天内的日期
	StartDate string `json:"start_date,omitempty"`
	// EndDate 查询结束时间（开区间）。格式：YYYY-MM-DD（广告主账户时区）。
	// 只允许指定过去60天内的日期
	EndDate string `json:"end_date,omitempty"`
}

// Encode implements GetRequest
func (r *FatigueGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("ad_id", r.AdID)
	values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// FatigueGetResponse 获取创意疲劳检测结果 API Response
type FatigueGetResponse struct {
	model.BaseResponse
	Data *FatigueGetResult `json:"data,omitempty"`
}

type FatigueGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 创意疲劳检测结果信息。
	List []Fatigue `json:"list,omitempty"`
}

// Fatigue 创意疲劳检测结果信息。
type Fatigue struct {
	// AdgroupID 广告（ad_id）所属的广告组ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdID 广告ID
	AdID string `json:"ad_id,omitempty"`
	// Date 所获取的指标（metrics）数据对应的日期
	Date string `json:"date,omitempty"`
	// Metrics 日期（date）对应的指标数据
	Metrics *FatigueMetrics `json:"metrics,omitempty"`
}

// FatigueMetrics 日期（date）对应的指标数据
type FatigueMetrics struct {
	// HasFatigue 是否检测到创意疲劳
	HasFatigue bool `json:"has_fatigue,omitempty"`
	// FatigueIndex 疲劳指数是衡量CPA（每行动成本），消耗，新用户触达率及其他表现指标每日变化的综合指标。指标数值越高代表创意疲劳程度越高
	FatigueIndex int `json:"fatigue_index,omitempty"`
	// Dnu 日新增用户人数，即该广告当日所吸引的新增用户人数
	Dnu int64 `json:"dnu,omitempty"`
	// DnuRatio 日新增用户比率。
	// 本指标计算方式为：该广告当日所吸引的新增用户人数/该广告过去60天内的日新增用户人数最大值
	DnuRatio float64 `json:"dnu_ratio,omitempty"`
	// Spend 总消耗，即投放广告产生的花费金额。对应币种为广告账户默认币种
	Spend float64 `json:"spend,omitempty"`
	// CostPerConversion 当广告属于非iOS 14专属推广系列时，本指标返回实际值。 当广告属于iOS 14专属推广系列时，本指标返回0.0。 您可通过/campaign/get/返回的campaign_type值确定推广系列是否为iOS 14专属推广系列。
	// 转化成本，即广告花费平均到单次转化的成本，计数结果基于展示时间点统计。对应币种为广告账户默认币种。
	CostPerConversion float64 `json:"cost_per_conversion,omitempty"`
	// SkanCostPerConversion 当广告属于iOS 14专属推广系列时，本指标返回实际值。当广告属于非iOS 14专属推广系列时，本指标返回0.0。 您可通过/campaign/get/返回的campaign_type值确定推广系列是否为iOS 14专属推广系列。
	// 转化成本 (SKAN)，即广告花费平均到单次转化的成本。由于SKAdnetwork(SKAN)限制，可能只有部分数据。计数结果基于回传时间统计，与实际发生时间存在延迟。对应币种为广告账户默认币种。
	SkanCostPerConversion float64 `json:"skan_cost_per_conversion,omitempty"`
}
