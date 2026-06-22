package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CopyTaskCheckRequest Get the results of an asynchronous copy task for an Upgraded Smart+ Campaign API Request
type CopyTaskCheckRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// TaskID ID of the asynchronous campaign copy task.
	// To get the task ID, use /smart_plus/campaign/copy/task/create/.
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

// CopyTaskCheckResponse Get the results of an asynchronous copy task for an Upgraded Smart+ Campaign API Response
type CopyTaskCheckResponse struct {
	model.BaseResponse
	Data *CopyTask `json:"data,omitempty"`
}

type CopyTask struct {
	// TaskStatus The status of the asynchronous campaign copy task.
	// Enum values:
	// RUNNING: The task is being processed.
	// SUCCESS: The task has been processed. Check the task_result to see if the task has succeeded.
	// FAILURE: The task fails to be processed.
	TaskStatus enum.CampaignCopyTaskStatus `json:"task_status,omitempty"`
	// TaskInfo Overview of the task result.
	TaskInfo *CopyTaskInfo `json:"task_info,omitempty"`
	// TaskResult The details of the task result.
	TaskResult *CopyTaskResult `json:"task_result,omitempty"`
}

// CopyTaskInfo Overview of the task result.
type CopyTaskInfo struct {
	// TotalAdCount The total number of ads that you tried to copy.
	TotalAdCount int `json:"total_ad_count,omitempty"`
	// SuccessAdCount The number of ads that have been successfully copied
	SuccessAdCount int `json:"success_ad_count,omitempty"`
}

// CopyTaskResult The details of the task result.
type CopyTaskResult struct {
	// CampaignID The ID of the newly created campaign.
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName The name of the newly created campaign.
	// If the campaign copy process fails, this field is still returned to indicate the campaign that was not successfully copied.
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignErrorInfos The errors encountered during the campaign copy process.
	// If no errors occurred, the value of this field is an empty list ([]).
	CampaignErrorInfos []string `json:"campaign_error_infos,omitempty"`
	// AdgroupResultList The details of the ad group copy results
	AdgroupResultList []CopyAdgroupResult `json:"adgroup_result_list,omitempty"`
}

type CopyAdgroupResult struct {
	// AdgroupID The ID of the newly created ad group.
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName The name of the newly created ad group.
	// If the ad group copy process fails, this field is still returned to indicate the ad group that was not successfully copied.
	AdgroupName string `json:"adgroup_name,omitempty"`
	// TotalAdCount The number of ads that you tried to copy into the new ad group.
	TotalAdCount int `json:"total_ad_count,omitempty"`
	// SuccessAdCount The number of ads that have been successfully copied into the new ad group.
	SuccessAdCount int `json:"success_ad_count,omitempty"`
	// AdgroupErrorList The errors encountered during the ad group copy process.
	// If no errors occurred, the value of this field is an empty list ([]).
	AdgroupErrorList []string `json:"adgroup_error_list,omitempty"`
	// AdStatus The result of copying the ads from the source ad group to the newly created ad group.
	// Enum values:
	// ALL_SUCCESS: All ads from the source ad group were successfully copied.
	// PARTIAL_SUCCESS: Some or all ads from the source ad group failed to be copied.
	AdStatus string `json:"ad_status,omitempty"`
	// AdResultList The details of the ad copy results.
	AdResultList []CopyAdResult `json:"ad_result_list,omitempty"`
}

type CopyAdResult struct {
	// IsSuccess Whether the copy of the source ad was successful.
	// Supported values: true, false.
	IsSuccess bool `json:"is_success,omitempty"`
	// SmartPlusAdID The ID of the newly created ad.
	SmartPlusAdID string `json:"smart_plus_ad_id,omitempty"`
	// AdName The name of the newly created ad.
	// If the ad copy process fails, this field is still returned to indicate the ad that was not successfully copied.
	AdName string `json:"ad_name,omitempty"`
	// AdErrorList The errors encountered during the ad copy process.
	// If no errors occurred, the value of this field is an empty list ([]).
	AdErrorList []string `json:"ad_error_list,omitempty"`
}
