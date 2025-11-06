package webhook

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DeleteRequest 删除 TikTok 账号 Webhook 配置 API Request
type DeleteRequest struct {
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
	// ItemList 仅当 event_type 设置为 COMMENT 时生效。
	// 不想再接收对应的评论更新事件通知的 TikTok 帖子 ID 列表。
	// 最大数量：100。
	// 若想获取已有的评论更新事件的 Webhook 配置中指定的 TikTok 帖子 ID 列表， 可使用/business/webhook/list/。
	// 若 event_type 设置为 COMMENT 且传入了本字段，则您将不再接收授权过该开发者应用（app_id）的 TikTok 账号下所指定的评论的更新事件通知。
	// 若 event_type 设置为 COMMENT 且未传入本字段，则您将不再接收授权过该开发者应用（app_id）的 TikTok 账号下所有评论的更新事件通知。
	ItemList []string `json:"item_list,omitempty"`
}

// Encode implements PostRequest
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeleteResponse 删除 TikTok 账号 Webhook 配置 API Response
type DeleteResponse struct {
	model.BaseResponse
	Data *Webhook `json:"data,omitempty"`
}
