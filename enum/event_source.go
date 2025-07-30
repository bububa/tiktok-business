package enum

// EventSource 本字段用于指定通过事件 API 上传的事件类型。
type EventSource string

const (
	// EventSourceWeb 在您的网站上发生的事件，通过 Pixel 代码追踪。
	EventSourceWeb EventSource = "web"
	// EventSourceApp  在您的应用程序上发生的事件，通过 TikTok 应用 ID 追踪。
	EventSourceApp EventSource = "app"
	// EventSourceOffline 在实体店铺中发生的转化事件，通过线下事件组 ID 追踪。
	EventSourceOffline EventSource = "offline"
	// EventSourceCRM 在 CRM 系统中发生的线索事件，通过 CRM 事件组 ID 追踪。
	EventSourceCRM EventSource = "crm"
)
