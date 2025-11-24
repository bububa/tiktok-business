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
	// AD_ACCOUNT_SUSPENSION：广告账户的暂时停用状态。
	// LEAD：线索。
	// AD_GROUP：广告组的审核状态。
	// AD：广告的审核状态。
	// TCM_SPARK_ADS：上传至某个 TCM 工作流程 2.0 订单的视频的 Spark Ads 授权状态。
	// CREATIVE_FATIGUE：单个广告、广告组中所有广告或广告账户下所有广告的疲劳状态。
	// API_SERVICE_STATUS: API 服务状态
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
