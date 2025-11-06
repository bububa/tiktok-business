package webhook

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 创建或更新 TikTok 账号 Webhook 配置 API Request
type UpdateRequest struct {
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
	// CallbackURL 回调 URL
	CallbackURL string `json:"callback_url,omitempty"`
	// ItemList 仅当 event_type 设置为 COMMENT 时生效。
	// 当 event_type 设置为 DEFAULT 或 VIDEO 时忽略本字段。
	// TikTok 帖子 ID 列表，用于获取这些帖子下评论的更新事件通知。
	// 最大数量：100。
	// 若想获取某个 TikTok 账号下的 TikTok 帖子的 ID 列表，可使用 /business/video/list/。
	// 若请求中 event_type 设置为 COMMENT 且指定了 item_list：
	// 开发者应用（app_id）从未配置过 item_list 时，您将仅接收到请求中通过 item_list 指定的帖子的 comment.update 事件通知。
	// 开发者应用（app_id）已配置过 item_list 时，您接收到之前已配置过的帖子以及请求中通过 item_list 指定的帖子的 comment.update 事件通知。
	// 若请求中 event_type 设置为 COMMENT 且未指定 item_list：
	// 开发者应用（app_id）从未配置过 item_list 时，您将接收到所有授权过该开发者应用（app_id）的 TikTok 账号下所有帖子的 comment.update 事件通知。
	// 开发者应用（app_id）已配置过 item_list 时，您接收到之前已配置过的帖子的 comment.update 事件通知。若想了解之前已配置过的具体帖子信息，请查看返回中的 item_list 字段值。
	ItemList []string `json:"item_list,omitempty"`
}

// Encode implements PostRequest
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 创建或更新 TikTok 账号 Webhook 配置 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *Webhook `json:"data,omitempty"`
}
