package subscription

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SubscribeRequest 新建订阅 API Request
type SubscribeRequest struct {
	// AppID 开发者应用ID。您可以导航至应用管理 > App Detail > 基本信息，获取应用ID。
	AppID string `json:"app_id,omitempty"`
	// Secret 开发者应用密钥。您可以导航至应用管理 > App Detail > 基本信息，获取密钥。
	Secret string `json:"secret,omitempty"`
	//	SubscribeEntity 订阅对象。
	//
	// 枚举值:
	// REPORT_DATA_CHANGE: the data change in any of the reporting metrics spend, impressions, clicks, and video_play_actions for one or more ad accounts. You need to specify advertiser_ids simultaneously. To find the detailed description for these reporting metrics, see Basic reports > Supported metrics.
	// Note: To use the REPORT_DATA_CHANGE value, ensure that you have enabled the Reporting > Consolidated Report permission for your developer app.
	// AD_ACCOUNT_SUSPENSION: the suspension statuses of ad accounts.
	// LEAD: leads.
	// TikTok is committed to providing "at least once delivery" for webhooks. Consequently, it’s possible that webhook endpoints may receive duplicate events. To mitigate potential issues, incorporate safeguards against processing duplicate events by designing your event handling to be idempotent.
	// AD_GROUP: the review status of an ad group.
	// AD: the review status of an ad.
	// TCM_VIDEOS: the linking of a TikTok video to a TTO campaign.
	// CREATIVE_FATIGUE: the fatigue status of an ad, ads within an ad group, or ads within an advertiser account.
	// API_SERVICE_STATUS: the status of API service.
	SubscribeEntity enum.SubscribeEntity `json:"subscribe_entity,omitempty"`
	// CallbackURL 回调链接
	CallbackURL string `json:"callback_url,omitempty"`
	// SubscriptionDetail 订阅内容
	SubscriptionDetail *SubscriptionDetail `json:"subscription_detail,omitempty"`
}

// Encode implements PostRequest interface
func (r *SubscribeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// SubscribeResponse 新建订阅 API Response
type SubscribeResponse struct {
	model.BaseResponse
	Data struct {
		// SubscriptionID 订阅 ID
		SubscriptionID string `json:"subscription_id,omitempty"`
	} `json:"subscription_id"`
}
