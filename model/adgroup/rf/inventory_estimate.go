package rf

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InventoryEstimateRequest 获取覆盖和频次广告库存预估 API Request
type InventoryEstimateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AudienceInfo 受众定向规则
	AudienceInfo *Audience `json:"audience_info,omitempty"`
	// ScheduleStartTime 排期开始时间（UTC+0）。格式： "YYYY-MM-DD HH:MM:SS"，"HH:MM:SS"的值需由投放地时区某天的00:00:00时间转换成UTC+0的时间。例如，若想在UTC-5 时区地区的 2020-10-24 00:00:00时间开始投放，需在此字段传入2020-10-24 05:00:00（UTC+0时间）。
	// 注意：若您指定的定向地域 (location_ids)遵守夏令时，请忽略夏令时的影响，传入标准时间。
	// 例如，若您想在荷兰当地时间2022-04-02 00:00:00 （夏令时已开始，UTC+2时区）开始投放，需将时区视作标准时区（UTC+1），并将本字段设置为2022-04-02 23:00:00
	ScheduleStartTime string `json:"schedule_start_time,omitempty"`
	// ScheduleEndTime 排期结束时间（UTC+0）。格式： "YYYY-MM-DD HH:MM:SS"。"HH:MM:SS"的值需由投放地时区某天的23:59:59时间转换成UTC+0的时间。例如，若想在UTC-5 时区地区的 2020-11-10 23:59:59时间结束投放，需在此字段传入2020-11-11 04:59:59（UTC+0时间）。schedule_end_time 与 schedule_start_time 的时间间隔不能超过30天。
	// 注意：若您指定的定向地域 (location_ids)遵守夏令时，请忽略夏令时的影响，传入标准时间。
	// 例如，若您想在荷兰当地时间2022-04-02 23:59:59（夏令时已开始，为UTC+2时区）结束投放，需将时区视作标准时区（UTC+1），并将本字段设置为2022-04-02 22:59:59
	ScheduleEndTime string `json:"schedule_end_time,omitempty"`
	// Frequency 频次。本字段与 frequency_schedule 共同定义用户看到广告的频次上限（仅适用于覆盖和频次广告）。
	// 例如，frequency = 2 且frequency_schedule = 3， 表示每 3 天至多 2 次展示。
	// 需满足以下条件：
	// 1 ≤ frequency ≤ frequency_schedule
	// 1 ≤ frequency_schedule ≤ min(schedule_end_time - schedule_start_time, 30)
	// 由frequency 和 frequency_schedule 定义的频次上限不可超过每天 4 次展示。
	Frequency int `json:"frequency,omitempty"`
	// FrequencySchedule 频次天数。与 frequency 共同定义用户看到广告的频次上限（仅适用于覆盖和频次广告）。请查看frequency字段的描述
	FrequencySchedule int `json:"frequency_schedule,omitempty"`
	// ObjectiveType 推广目标。枚举值：RF_REACH
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// CPVVideoDuration 目标视频播放时长。
	// 枚举值：SIX_SECONDS(视频播放6秒)，TWO_SECONDS(视频播放2秒)
	CPVVideoDuration enum.CPVVideoDuration `json:"cpv_video_duration,omitempty"`
	// FeedType 信息流类型。枚举值: STANDARD_FEED, TOP_FEED。STANDARD_FEED指库存高，且CPM价格正常。TOP_FEED 指库存受限，且CPM价格较高
	FeedType enum.FeedType `json:"feed_type,omitempty"`
	// RfPurchasedType 购买类型。枚举值：
	// FIXED_SHOW: 按固定展示量购买。
	// FIXED_REACH:按固定覆盖人数购买。
	// FIXED_BUDGET: 按固定预算购买。
	RfPurchasedType enum.RfPurchasedType `json:"rf_purchased_type,omitempty"`
	// Budget 预算金额。当rf_purchased_type 为 FIXED_BUDGET时必填。基于您设置的预算金额，其他字段会返回预估值
	Budget float64 `json:"budget,omitempty"`
	// PurchasedImpression 想购买的展示量，以1000为单位。当 rf_purchased_type为 FIXED_SHOW时必填。基于您设置的展示量，其他字段会返回预估值
	PurchasedImpression int64 `json:"purchased_impression,omitempty"`
	// PurchasedReach 想购买的覆盖人数，以1000为单位。当 rf_purchased_type 为FIXED_REACH时必填。 基于您设置的覆盖人数，其他字段会返回预估值。
	PurchasedReach int64 `json:"purchased_reach,omitempty"`
	// RfCampaignType 仅当推广目标（objective_type）设置为RF_REACH时，可用此字段将推广系列设置为TikTok Pulse推广系列，从而获取PREMIUM标签。
	// 枚举值： STANDARD （普通覆盖和频次推广系列）， PULSE（TikTok Pulse推广系列）。
	// 若指定为PULSE，则仅返回一个国家或地区ID(location_id)。
	// 注意：
	// 推广目标非RF_REACH时，本字段无需传入。
	// 使用/adgroup/rf/create/ 创建TikTok Pulse广告组时，请求中需将本接口返回的location_id值传给location_ids。
	RfCampaignType enum.RfCampaignType `json:"rf_campaign_type,omitempty"`
}

// Encode implements GetRequest
func (r *InventoryEstimateRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("audience_info", string(util.JSONMarshal(r.AudienceInfo)))
	values.Set("schedule_start_time", r.ScheduleStartTime)
	values.Set("schedule_end_time", r.ScheduleEndTime)
	values.Set("frequency", strconv.Itoa(r.Frequency))
	values.Set("frequency_schedule", strconv.Itoa(r.FrequencySchedule))
	values.Set("objective_type", string(r.ObjectiveType))
	if r.CPVVideoDuration != "" {
		values.Set("cpv_video_duration", string(r.CPVVideoDuration))
	}
	if r.FeedType != "" {
		values.Set("feed_type", string(r.FeedType))
	}
	if r.RfPurchasedType != "" {
		values.Set("rf_purchased_type", string(r.RfPurchasedType))
	}
	if r.Budget > 1e-15 {
		values.Set("budget", strconv.FormatFloat(r.Budget, 'f', 0, 64))
	}
	if r.PurchasedImpression > 0 {
		values.Set("purchased_impression", strconv.FormatInt(r.PurchasedImpression, 10))
	}
	if r.PurchasedReach > 0 {
		values.Set("purchased_reach", strconv.FormatInt(r.PurchasedReach, 10))
	}
	if r.RfCampaignType != "" {
		values.Set("rf_campaign_type", string(r.RfCampaignType))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// InventoryEstimateResponse 获取覆盖和频次广告库存预估 API Response
type InventoryEstimateResponse struct {
	model.BaseResponse
	Data *InventoryEstimateResult `json:"data,omitempty"`
}

type InventoryEstimateResult struct {
	// DefaultResult 库存预估结果。
	// 若您在请求中为purchased_impression, budget或 purchased_reach设置了固定值，则会返回预估库存的对应预估值。例如，若您在请求中为budget设置了固定值，则base_info对象会返回能够满足您固定预算要求的impression, cpm, reach 以及 average_frequency对应值。
	DefaultResult *EstimatedInfo `json:"default_result,omitempty"`
	// Results 不同预算对应的预估库存结果。每条结果中会返回预算对应的预估展示量、千次展示成本、覆盖人数、频次、每日花费和人均展示频次分布。
	// 本对象的数据结构与default_result 相同。
	Results []EstimatedInfo `json:"results,omitempty"`
}
