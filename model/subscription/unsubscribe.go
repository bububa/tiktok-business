package subscription

import "github.com/bububa/tiktok-business/util"

// UnsubscribeRequest 取消订阅 API Request
type UnsubscribeRequest struct {
	// AppID 开发者应用ID。您可以导航至应用管理 > App Detail > 基本信息，获取应用ID。
	AppID string `json:"app_id,omitempty"`
	// Secret 开发者应用密钥。您可以导航至应用管理 > App Detail > 基本信息，获取密钥。
	Secret string `json:"secret,omitempty"`
	// SubscriptionID 订阅ID。
	// 您可通过/subscription/get/返回的subscription_id获取单个App下所有订阅的ID。
	SubscriptionID string `json:"subscription_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *UnsubscribeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
