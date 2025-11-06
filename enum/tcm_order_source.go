package enum

// TCMOrderSource 订单来源，用于区分订单是否为托管服务提供商（MSP）订单。
type TCMOrderSource string

const (
	// TCMOrderSource_MSP ：该订单是托管服务提供商订单。
	TCMOrderSource_MSP TCMOrderSource = "MSP"
	// TCMOrderSource_OTHER ：该订单非托管服务提供商订单。
	TCMOrderSource_OTHER TCMOrderSource = "OTHER"
)
