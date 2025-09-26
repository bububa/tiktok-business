package webhook

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Request Webhook Request
type Request struct {
	// RequestID Webhook 请求的唯一 ID
	RequestID string `json:"request_id,omitempty"`
	// Object 订阅对象类型。
	// 有效值：
	// 1：线索
	// 2：广告组的审核状态。
	// 3 ：广告的审核状态。
	// 8：单个广告、广告组中所有广告或广告账户下所有广告的疲劳状态。
	// 9 ：TCM 订单中视频的 Spark Ads 授权状态。
	Object enum.WebhookObject `json:"object,omitempty"`
	// Time 发送事件通知时的 Unix 时间戳（非触发广告审核状态变化事件的时间）。
	Time model.DateTime `json:"time,omitzero"`
	Suspension
}
