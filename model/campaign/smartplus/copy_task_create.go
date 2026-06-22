package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	ad "github.com/bububa/tiktok-business/model/ad/smartplus"
	adgroup "github.com/bububa/tiktok-business/model/adgroup/smartplus"
	"github.com/bububa/tiktok-business/util"
)

// CopyTaskCreateRequest Create an asynchronous copy task for an Upgraded Smart+ Campaign API Request
type CopyTaskCreateRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// RequestID Request ID that supports idempotency to prevent you from sending the same request twice. If you retry requests with the same request ID multiple times within the 10-second cache time, then only one request will succeed. If a duplicate request with the expired request ID is received after the cache time, the server will treat it as a new request and process it accordingly.
	// It is different from the request_id returned in the response parameters, which is used to uniquely identify an HTTP request.
	// The value should be a string representation of a 64-bit integer.
	RequestID string `json:"request_id,omitempty"`
	// CampaignID The ID of an Upgraded Smart+ Campaign that you want to copy.
	// To retrieve Upgraded Smart+ Campaigns within your ad account, use /smart_plus/campaign/get/.
	// Note:
	// The source campaign must use one of these advertising objectives: APP_PROMOTION, LEAD_GENERATION, WEB_CONVERSIONS
	// The source campaign must not be deleted.
	// The source campaign must contain at least 1 undeleted ad group, which must contain at least 1 undeleted ad.
	// The source campaign must satisfy all per-level limits, and the total copied assets must also not exceed the global new campaign limits:
	// On the source campaign: you may have a maximum of 10 undeleted ad groups, each ad group may contain a maximum of 30 undeleted ads, and each individual ad may use a maximum of 50 creatives.
	// Separately, there are hard global caps that apply to the resulting new campaign: no more than 200 total ads may be copied over, and no more than 1,000 total creatives may be copied into the new campaign.
	CampaignID string `json:"campaign_id,omitempty"`
	// OperationStatus The status of the new campaign when created.
	// Enum values:
	// ENABLE: The campaign is enabled when created.
	// DISABLE: The campaign is disabled when created.
	// Default value: DISABLE.
	// If you want to update the status of the campaign after creation, use /smart_plus/campaign/status/update/.
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// CampaignName The name for the new campaign.
	// Length limit: 512 characters. Emojis are not supported. Each word in Chinese or Japanese counts as two characters, while each letter in English counts as one character.
	// If not specified, this field will default to "COPIED_{{name_of_the_source_campaign}}". For instance, if the source campaign is named "FIRST_CAMPAIGN" and this field is not specified, the name of the new campaign will be "COPIED_FIRST_CAMPAIGN".
	CampaignName string `json:"campaign_name,omitempty"`
	// Budget The budget for the new campaign or the budget limit for all ad groups under the new campaign.
	// When budget_optimize_on of the source campaign is true, this field represents the fixed budget or initial budget for the new campaign.
	// When budget_auto_adjust_strategy is UNSET, this field represents the fixed budget for the new campaign.
	// When budget_auto_adjust_strategy is AUTO_BUDGET_INCREASE, this field represents the initial budget for the new campaign. To retrieve the current campaign budget, use /smart_plus/campaign/get/ after the copy task is completed and check the returned current_budget.
	// When budget_optimize_on of the source campaign is false, this field represents the budget limit for all ad groups under the new campaign.
	// When budget_mode is BUDGET_MODE_DAY, this field represents the daily limit for all ad groups under the new campaign.
	// When budget_mode is BUDGET_MODE_TOTAL, this field represents the total limit for all ad groups under the new campaign.
	// If not specified, this field will default to the budget of the source campaign or the budget limit for all ad groups under the source campaign.
	Budget float64 `json:"budget,omitempty"`
	// ScheduleType Specify this field only when you want to set the same schedule for all ad groups in the new campaign.
	// If you want to use schedules of the source ad groups in the new campaign, leave this field unspecified.
	// Schedule type for all new ad groups.
	// Enum values:
	// SCHEDULE_START_END: Set start time and end time to run ad groups. You need to pass both schedule_start_time and schedule_end_time at the same time.
	// SCHEDULE_FROM_NOW: Set start time to run ad groups continuously. You need to pass schedule_start_time and the end time will be automatically set to 10 years later than the start time.
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime Required when schedule_type is passed.
	// Schedule start time (UTC+0) for all new ad groups, in the format of YYYY-MM-DD HH:MM:SS.
	// The start time can be up to 12 hours earlier than the current time, but cannot be later than 2028-01-01 00:00:00.
	ScheduleStartTime string `json:"schedule_start_time,omitempty"`
	// ScheduleEndTime Required when schedule_type is SCHEDULE_START_END.
	// Not supported when schedule_type is SCHEDULE_FROM_NOW.
	// Schedule end time (UTC+0) for all new ad groups, in the format of YYYY-MM-DD HH:MM:SS.
	// The end time cannot be later than 2038-01-01 00:00:00.
	ScheduleEndTime string `json:"schedule_end_time,omitempty"`
	// Dayparting Specify this field only when you want to set the same ad delivery arrangement for all ad groups in the new campaign.
	// If you want to use the same ad delivery arrangements of the source ad groups in the new campaign, leave this field unspecified.
	// Ad delivery arrangement, in the format of a string that consists of 48 x 7 characters. Each character is mapped to a 30-minute timeframe from Monday to Sunday. Each character can be set to either 0 or 1. 1 represents delivery in the 30-minute timeframe, and 0 stands for non-delivery in the 30-minute timeframe. The first character is mapped to 0:01-0:30 of Monday; The second character is mapped to 0:31-1:00 of Monday, and the last character represents 23:31-0:00 Sunday.
	// Note: An all-1 value and when this field is not specified, are considered full-time delivery.
	Dayparting string `json:"dayparting,omitempty"`
	// DeepCopyMode The copying mode determining how you create ad groups and ads (asset groups) in the new campaign.
	// Enum values:
	// DEFAULT: The default copy mode. You copy all undeleted ad groups and ads from the source campaign to the new one. You can omit the adgroup_list field, as the API ignores it in this mode.
	// CUSTOM: The custom copy mode. You copy only the specified undeleted ad groups and ads from the source campaign to the new one. You also have options to customize specific ad group or ad settings to override the source configurations.
	// When using this mode, you need to provide the adgroup_list field simultaneously. You can copy a maximum of 20 ads per ad group, across a maximum of 30 ad groups.
	// Default value: DEFAULT.
	DeepCopyMode string `json:"deep_copy_mode,omitempty"`
	// AdgroupList When deep_copy_mode is set to DEFAULT or not passed, this field is ignored.
	// When deep_copy_mode is set to CUSTOM, this field is required.
	// The customized settings for ad groups in the new campaign.
	// Max size: 10.
	AdgroupList []Adgroup `json:"adgroup_list,omitempty"`
}

// Adgroup Upgraded Smart+ ad group
type Adgroup struct {
	adgroup.Adgroup
	AdList []ad.Ad `json:"ad_list,omitempty"`
}

// Encode implements PostRequest
func (r *CopyTaskCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CopyTaskCreateResponse Create an asynchronous copy task for an Upgraded Smart+ Campaign API Response
type CopyTaskCreateResponse struct {
	model.BaseResponse
	Data *CopyTaskCreateResult `json:"data,omitempty"`
}

type CopyTaskCreateResult struct {
	// TaskID ID of the asynchronous campaign copy task.
	// To check the results of the task, use /smart_plus/campaign/copy/task/check/.
	TaskID string `json:"task_id,omitempty"`
	// AdgroupErrorList The errors encountered during the process of copying specific ad groups.
	AdgroupErrorList []AdgroupError `json:"adgroup_error_list,omitempty"`
}

type AdgroupError struct {
	// AdgroupID The ID of the ad group that fails to be copied.
	AdgroupID string `json:"adgroup_id,omitempty"`
	// ErrorMessage The error encountered during the process of copying the ad group (adgroup_id).
	ErrorMessage string `json:"error_message,omitempty"`
}

func (e AdgroupError) Error() string {
	return util.StringsJoin(e.ErrorMessage, "[", e.AdgroupID, "]")
}
