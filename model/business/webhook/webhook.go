package webhook

type Webhook struct {
	// AppID 开发者应用 ID
	AppID string `json:"app_id,omitempty"`
	// EventType 订阅的 Webhook 事件类型。
	// 枚举值：
	// VIDEO： 帖子发布事件，使您能获取通过某些 TikTok API 接口发布的内容的实时发布状态通知。
	// 若想了解此类事件类型中事件的结构和有效负载示例，请查看帖子发布事件。
	// COMMENT：评论更新事件，包括 comment.update 事件。
	// 若想了解此类事件类型中事件的结构和有效负载示例，请查看评论更新事件。
	EventType string `json:"event_type,omitempty"`
	// CallbackURL 回调 URL
	CallbackURL string `json:"callback_url,omitempty"`
	// ItemList 仅当 event_type 为 COMMENT 且开发者应用（app_id）配置了 item_list 时返回本字段。
	// TikTok 帖子 ID 列表，用于获取这些帖子下评论的更新事件通知。
	ItemList []string `json:"item_list,omitempty"`
}
