package enum

// DeepFunnelEventSource  事件源类型
type DeepFunnelEventSource string

const (
	// DeepFunnelEventSource_PIXEL：Pixel。
	DeepFunnelEventSource_PIXEL DeepFunnelEventSource = "PIXEL"
	// DeepFunnelEventSource_OFFLINE：线下事件组。
	DeepFunnelEventSource_OFFLINE DeepFunnelEventSource = "OFFLINE"
	// DeepFunnelEventSource_CRM：CRM 事件组。
	DeepFunnelEventSource_CRM DeepFunnelEventSource = "CRM"
)
