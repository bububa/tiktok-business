package rf

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// EstimatedInfoRequest 获取覆盖和频次广告组预估信息 API Request
type EstimatedInfoRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdgroupID 广告组ID
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
}

// Encode implements GetRequest
func (r *EstimatedInfoRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("adgroup_ids", string(util.JSONMarshal(r.AdgroupIDs)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// EstimatedInfoResponse 获取覆盖和频次广告组预估信息 API Response
type EstimatedInfoResponse struct {
	model.BaseResponse
	Data struct {
		// EstimatedInfo 预估信息
		EstimatedInfo []EstimatedInfo `json:"estimated_info,omitempty"`
	} `json:"data"`
}

// EstimatedInfo 预估信息
type EstimatedInfo struct {
	// AdgroupID 广告组ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// BaseInfo 基本信息
	BaseInfo *EstimatedBaseInfo `json:"base_info,omitempty"`
	// DailyCost 每日花费
	DailyCost *EstimatedDailyCost `json:"daily_cost,omitempty"`
	// FrequencyPerPerson 人均展示频次分布。
	// 注意：由 frequency 和 frequency_schedule 定义的频次上限超过每天 4 次展示时，frequency_per_person 的值将不再变动。
	FrequencyPerPerson *FrequencyPerPerson `json:"frequency_per_person,omitempty"`
}

// EstimatedBaseInfo 基本信息
type EstimatedBaseInfo struct {
	// Budget 总预算
	Budget float64 `json:"budget,omitempty"`
	// CPM 千次展示成本。公式：cpm = budget / impression
	CPM float64 `json:"cpm,omitempty"`
	// Impression 总展示量，以 1,000 为单位。
	// 例如，本字段值为 110 代表总展示量为 110,000。
	Impression int64 `json:"impression,omitempty"`
	// Reach 总覆盖人数，以 1,000 为单位。
	// 例如，本字段值为 110 代表总覆盖人数为 110,000
	Reach int64 `json:"reach,omitempty"`
	// MaxReach 可预订的全部覆盖人数，以 1,000 为单位。此数据基于所指定的信息流类型、定向设置和推广系列排期。
	// 例如，本字段值为 200 代表可用库存中人数为 200,000。
	MaxReach int64 `json:"max_reach,omitempty"`
	// AverageFrequency 公式：人均展示频次 （average_frequency）= 总展示量（impression）/总覆盖人数（reach）
	AverageFrequency float64 `json:"average_frequency,omitempty"`
}

// EstimatedDailyCost 每日花费
type EstimatedDailyCost struct {
	// Cost 当日花费
	Cost float64 `json:"cost,omitempty"`
	// Date 日期。
	// 格式："YYYY-MM-DD"。
	// 示例："2024-01-01"
	Date string `json:"date,omitempty"`
}

// FrequencyPerPerson 人均展示频次分布。
type FrequencyPerPerson struct {
	// Frequency 人均展示频次
	Frequency int `json:"frequency,omitempty"`
	// Percentage frequency 对应的用户人数占比。
	// 例如，若frequency 为 1， percentage 为0.42304230423042305，则触达次数为 1 的用户占比为约 42.3%。
	Percentage float64 `json:"percentage,omitempty"`
}
