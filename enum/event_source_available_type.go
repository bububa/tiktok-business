package enum

// EventSourceAvailableType 事件类别。
type EventSourceAvailableType string

const (
	// EVENT_RECEIVED：所有事件。
	EVENT_RECEIVED EventSourceAvailableType = "EVENT_RECEIVED"
	// EVENT_WITH_CONTENT_ID：带有内容 ID（content ID）的事件。
	EVENT_WITH_CONTENT_ID EventSourceAvailableType = "EVENT_WITH_CONTENT_ID"
	// EVENT_WITH_CONTENT_ID_MATCHING_INVENTORY：内容 ID 能匹配库存的事件。
	EVENT_WITH_CONTENT_ID_MATCHING_INVENTORY EventSourceAvailableType = "EVENT_WITH_CONTENT_ID_MATCHING_INVENTORY"
)
