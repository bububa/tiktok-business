package enum

// EventSourceType 事件源类型。
type EventSourceType string

const (
	// EventSourceType_PIXEL: Pixel。
	EventSourceType_PIXEL EventSourceType = "PIXEL"
	// EventSourceType_APP: 应用。
	EventSourceType_APP EventSourceType = "APP"
)
