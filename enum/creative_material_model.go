package enum

// CreativeMaterialMode 创意投放方式。
type CreativeMaterialMode string

const (
	// CreativeMaterialMode_CUSTOM（自定义）
	CreativeMaterialMode_CUSTOM CreativeMaterialMode = "CUSTOM"
	// CreativeMaterialMode_SMART_CREATIVE（智能创意）。
	CreativeMaterialMode_SMART_CREATIVE CreativeMaterialMode = "SMART_CREATIVE"
	// CreativeMaterialMode_DYNAMIC 程序化
	CreativeMaterialMode_DYNAMIC CreativeMaterialMode = "DYNAMIC"
)
