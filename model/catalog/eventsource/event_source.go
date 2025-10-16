package eventsource

// EventSource 事件源数据
type EventSource struct {
	// EventSourceName 事件源名
	EventSourceName string `json:"event_source_name,omitempty"`
	// AppID 移动应用ID
	AppID string `json:"app_id,omitempty"`
	// PixelCode Pixel代码
	PixelCode string `json:"pixel_code,omitempty"`
}
