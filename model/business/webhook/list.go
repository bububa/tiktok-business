package webhook

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取 TikTok 账号 Webhook 配置 API Request
type ListRequest struct {
	// AppID 开发者应用ID。
	// 若想获取应用ID，您可以导航至应用管理 > App Detail > 基本信息。
	AppID string `json:"app_id,omitempty"`
	// Secret 开发者应用密钥。
	// 若想获取密钥，可以导航至应用管理 > App Detail > 基本信息。
	Secret string `json:"secret,omitempty"`
	// EventType 想要订阅的 Webhook 事件类型。
	// 枚举值：
	// VIDEO： 帖子发布事件，使您能获取通过某些 TikTok API 接口发布的内容的实时发布状态通知。
	// 若想了解此类事件类型中事件的结构和有效负载示例，请查看帖子发布事件。
	// COMMENT：评论更新事件，包括 comment.update 事件。
	// 若想了解此类事件类型中事件的结构和有效负载示例，请查看评论更新事件。
	EventType string `json:"event_type,omitempty"`
}

// Encode implements GetRequest
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("app_id", r.AppID)
	values.Set("secret", r.Secret)
	values.Set("event_type", r.EventType)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ListResponse 获取 TikTok 账号 Webhook 配置 API Response
type ListResponse struct {
	model.BaseResponse
	Data *Webhook `json:"data,omitempty"`
}
