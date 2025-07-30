package enum

// DeepBidType 深度事件出价类型
type DeepBidType string

const (
	// DeepBidType_DEFAULT 无深度事件出价
	DeepBidType_DEFAULT DeepBidType = "DEFAULT"
	// DeepBidType_MIN 双出价。运作方式 - 我们会控制平均成本稳定在您的目标成本上下，包括安装和深层目标成本, 最适合以下情况 - 对安装成本和深层目标成本均有预期, 请注意 - 安装量可能低于其他出价策略
	DeepBidType_MIN DeepBidType = "MIN"
	// DeepBidType_PACING 自动优化。运作方式 - 我们会控制平均安装成本稳定在您的目标成本上下，同时尽可能多获取深层目标事件.最适合以下情况 - 对于安装成本有预期，但对于深层目标成本希望系统自动优化, 请注意 - 深层目标成本可能高于双出价策略
	DeepBidType_PACING DeepBidType = "PACING"
	// DeepBidType_VO_MIN 已废弃,在有成本上限的条件下进行出价。 本出价策略只适用于价值优化目标 (optimization_goal = VALUE)。详情参见价值优化。
	DeepBidType_VO_MIN DeepBidType = "VO_MIN"
	// DeepBidType_VO_MIN_ROAS 在有ROAS目标的条件下出价。 本出价策略只适用于价值优化目标 (optimization_goal = VALUE)。详情参见价值优化。
	DeepBidType_VO_MIN_ROAS DeepBidType = "VO_MIN_ROAS"
	// DeepBidType_VO_HIGHEST_VALUE 无限制出价。 本出价策略只适用于价值优化目标 (optimization_goal = VALUE)。详情参见价值优化。
	DeepBidType_VO_HIGHEST_VALUE DeepBidType = "VO_HIGHEST_VALUE"
	// DeepBidType_AEO 以应用事件优化为目标进行出价。
	DeepBidType_AEO DeepBidType = "AEO"
)
