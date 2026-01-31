package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// StatusUpdateRequest Update the operation statuses of Upgraded Smart+ Campaigns API Request
type StatusUpdateRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignIDs The IDs of the campaigns to operate on.
	// To obtain campaign IDs, use /smart_plus/campaign/get/.
	// Size range: 1-20.
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// OperationStatus The operation to perform on the campaigns.
	// Enum values: DELETE (delete), DISABLE (pause), ENABLE (enable).
	// Note: The status of deleted campaigns cannot be modified.
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// PostbackWindowMode Valid only when the following conditions are all met:
	// The campaigns (campaign_ids) are Dedicated Campaigns (campaign_type is IOS14_CAMPAIGN).
	// You specify operation_status as DISABLE at the same time in the request.
	// postback_window_mode has not been configured for the campaign.
	// You have enabled SKAN 4.0 for your App.
	//
	// The mode that defines which SKAN (SKAdNetwork) 4.0 postback you want to secure. Options with longer windows require more time to receive and, as a result, more time to release the campaign back to the available number.
	//
	// Enum values:
	// POSTBACK_WINDOW_MODE1: This option secures the first postback, which corresponds to the 0-2 day attribution window. The data can take up to 4 days to return, and the campaign will wait for 4 days to release the campaign quota.
	// POSTBACK_WINDOW_MODE2: This mode secures the first two postbacks, which corresponds to the 3-7 day attribution window. The data can take up to 13 days to return, and the campaign will wait for 13 days to release the campaign quota.
	// POSTBACK_WINDOW_MODE3: This mode secures all three postbacks, which corresponds to the 8-35 day attribution window. The data can take up to 41 days to return, and the campaign will wait for 41 days to release the campaign quota.
	//
	// To learn more about SKAN 4.0 attribution window configuration and campaign quota, see Dedicated Campaigns.
	// Note:
	//
	// If you have set up Mobile Measurement Partner (MMP) Tracking with Adjust, Airbridge, Appsflyer, Branch, Kochava, or Singular for your App and your MMP SDK version is updated to a SKAN 4.0 supported SDK, you can transition your App to SKAN 4.0 on Events Manager. To learn about how to set up MMP tracking for your App, see How to Set Up App Attribution in TikTok Ads Manager. To learn more about how to transition your App to SKAN 4.0, see About SKAN 4.0 and TikTok and How to transition to SKAN 4.0.
	// If operation_status is set to ENABLE or not passed, and you pass in this field at the same time, an error will occur.
	// If you pass in this field when you have not enabled SKAN 4.0 for your App (app_id), an error will occur.
	// If this field is not passed when campaign_type is set to IOS14_CAMPAIGN, operation_status is set to DISABLE, and you have enabled SKAN 4.0 for your App, this field will default to POSTBACK_WINDOW_MODE1.
	// The postback window mode configuration can be partially successful. If you provide campaign_ids for non-Dedicated Campaigns, or Dedicated Campaigns for which postback_window_mode has been configured, those campaigns will be ignored and their postback_window_mode will not be modified. Only the Dedicated Campaigns that meet the conditions mentioned above will have their postback_window_mode configured.
	// If you have enabled SKAN 4.0 for your App, ensure that you target devices running iOS 16.1 and later so that you can receive SKAN 4.0 postbacks. To only target iOS 16.1+ devices, set min_ios_version to 16.1 at the ad group level.
	PostbackWindowMode enum.PostbackWindowMode `json:"postback_window_mode,omitempty"`
}

// Encode implements PostRequest interface
func (r *StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse Update the operation statuses of Upgraded Smart+ Campaigns API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data *StatusUpdateResult `json:"data,omitempty"`
}

type StatusUpdateResult struct {
	// CampaignList A list of campaigns for which the operation was successful.
	CampaignList []UpdatedCampaign `json:"campaign_list,omitempty"`
	// ErrorList A list of errors and the corresponding failed campaigns
	ErrorList []ErrUpdate `json:"error_list,omitempty"`
}

// UpdatedCampaign 所指定的推广系列的相关信息
type UpdatedCampaign struct {
	// CampaignID 推广系列ID
	CampaignID string `json:"campaign_id,omitempty"`
	// Status 该推广系列（ campaign_id ）当前的状态。
	// 枚举值：DELETE （已删除）， DISABLE （已暂停）， ENABLE （投放中）。
	Status enum.OperationStatus `json:"status,omitempty"`
	// PostbackWindowMode 选定的模式决定您想要确保能接收到的 SKAN (SKAdNetwork) 4.0 回传。
	// 枚举值：
	// POSTBACK_WINDOW_MODE1 ：此模式确保能接收到第一次回传，对应的归因窗口期为 0-2 天。 回传数据最长需要 4 天时间接收，推广系列配额释放的时间最长也为 4 天。
	// POSTBACK_WINDOW_MODE2 ：此模式确保能接收到第一次及第二次回传，对应的归因窗口期为 3-7 天。回传数据最长需要 13 天时间接收，推广系列配额释放的时间最长也为 13 天。
	// POSTBACK_WINDOW_MODE3 ：此模式确保能接收到三次回传，对应的归因窗口期为 8-35 天。回传数据最长需要 41 天时间接收，推广系列配额释放的时间最长也为 41 天。
	PostbackWindowMode enum.PostbackWindowMode `json:"postback_window_mode,omitempty"`
}

type ErrUpdate struct {
	// CampaignID The ID of the campaign for which the operation failed
	CampaignID string `json:"campaign_id,omitempty"`
	// ErrorMessage The error message associated with the failed campaign.
	ErrorMessage string `json:"error_message,omitempty"`
}

func (e ErrUpdate) Error() string {
	return util.StringsJoin(e.ErrorMessage, ", campaign_id:", e.CampaignID)
}
