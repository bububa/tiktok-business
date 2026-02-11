package campaign

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CopyTaskCheckRequest 获取推广系列异步复制任务的结果 API Request
type CopyTaskCheckRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// TaskID 推广系列异步复制任务的 ID。
	// 若想获取任务 ID，需使用 /campaign/copy/task/create/。
	TaskID string `json:"task_id,omitempty"`
}

// Encode implements GetRequest
func (r *CopyTaskCheckRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("task_id", r.TaskID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// CopyTaskCheckResponse 获取推广系列异步复制任务的结果 API Response
type CopyTaskCheckResponse struct {
	model.BaseResponse
	Data *CopyTaskCheckResult `json:"data,omitempty"`
}

type CopyTaskCheckResult struct {
	// TaskStatus 推广系列异步复制任务的状态。
	// 枚举值：
	// RUNNING：任务正在处理中。
	// SUCCESS：任务已经处理完成。请检查 task_result中的结果，确认任务是否成功。
	// FAILURE：任务处理失败。
	TaskStatus string `json:"task_status,omitempty"`
	// TaskInfo 任务结果概述
	TaskInfo *CopyTaskInfo `json:"task_info,omitempty"`
	// TaskResult 任务结果详情
	TaskResult *CopyTaskResult `json:"task_result,omitempty"`
}

// CopyTaskInfo 任务结果概述
type CopyTaskInfo struct {
	// TotalAdCount 尝试复制的广告总数
	TotalAdCount int64 `json:"total_ad_count,omitempty"`
	// SuccessAdCount 成功复制的广告数量
	SuccessAdCount int64 `json:"success_ad_count,omitempty"`
	// TotalSmartCreativeAdgroupCount 尝试复制的智能创意广告组的总数
	TotalSmartCreativeAdgroupCount int64 `json:"total_smart_creative_adgroup_count,omitempty"`
	// SuccessSmartCreativeAdgroupCount 成功复制的智能创意广告组的数量
	SuccessSmartCreativeAdgroupCount int64 `json:"success_smart_creative_adgroup_count,omitempty"`
}

// CopyTaskResult
type CopyTaskResult struct {
	// CampaignID 新生成的推广系列的 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName 新生成的推广系列的名称。
	// 若推广系列复制失败，此字段仍将返回，可用于识别未成功复制的推广系列。
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignErrorInfos 推广系列复制过程中发生的报错。
	// 若无报错，此字段的值将为空列表（[]）。
	CampaignErrorInfos []string `json:"campaign_error_infos,omitempty"`
	// AdgroupResultList 广告组复制结果详情
	AdgroupResultList []CopyAdgroupResult `json:"adgroup_result_list,omitempty"`
}

type CopyAdgroupResult struct {
	// AdgroupID 新生成的广告组的 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName 新生成的广告组的名称
	AdgroupName string `json:"adgroup_name,omitempty"`
	// TotalAdCount 尝试复制到新广告组的广告总数
	TotalAdCount int64 `json:"total_ad_count,omitempty"`
	// SuccessAdCount 成功复制到新广告组的广告数量
	SuccessAdCount int64 `json:"success_ad_count,omitempty"`
	// AdgroupErrorList 广告组复制过程中发生的报错。
	// 若无报错，此字段的值将为空列表（[]）
	AdgroupErrorList []string `json:"adgroup_error_list,omitempty"`
	// AdStatus 将广告从原广告组复制到新生成的广告组的结果。
	// 枚举值：
	// ALL_SUCCESS：所有原广告组中的广告全部复制成功。
	// PARTIAL_SUCCESS：原广告组中的广告部分复制失败或全部复制失败。
	AdStatus string `json:"ad_status,omitempty"`
	// IsSmartCreative 新生成的广告组是否为智能创意广告组。
	// 枚举值：true，false。
	IsSmartCreative bool `json:"is_smart_creative,omitempty"`
	// SmartCreativeResult 仅当is_smart_creative 为 true 时返回。
	// 原智能创意广告组中广告素材的复制结果详情。
	SmartCreativeResult *CopySmartCreativeResult `json:"smart_creative_result,omitempty"`
	// AdListResult 广告复制结果详情
	AdListResult []CopyAdResult `json:"ad_list_result,omitempty"`
}

type CopySmartCreativeResult struct {
	// IsSuccess 原智能创意广告组中的广告素材是否复制成功。
	// 枚举值：true，false。
	IsSuccess bool `json:"is_success,omitempty"`
	// SmartCreativeErrorList 智能创意广告组中广告素材的复制过程中发生的 报错。
	// 若无报错，此字段的值将为空列表（[]）。
	SmartCreativeErrorList []string `json:"smart_creative_error_list,omitempty"`
}

type CopyAdResult struct {
	// IsSuccess 原广告是否复制成功。
	// 枚举值：true，false。
	IsSuccess bool `json:"is_success,omitempty"`
	// AdID 新生成的广告的 ID
	AdID string `json:"ad_id,omitempty"`
	// AdName 新生成的广告的名称
	AdName string `json:"ad_name,omitempty"`
	// AdErrorList 广告复制过程中发生的报错
	AdErrorList []string `json:"ad_error_list,omitempty"`
}
