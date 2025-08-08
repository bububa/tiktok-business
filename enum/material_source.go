package enum

// MaterialSource 素材来源
type MaterialSource string

const (
	// MaterialSource_UPLOADED 用户上传
	MaterialSource_UPLOADED MaterialSource = "UPLOADED"
	// MaterialSource_CREATIVE_TOOL_OTHERS 创意工具-其他
	MaterialSource_CREATIVE_TOOL_OTHERS MaterialSource = "CREATIVE_TOOL_OTHERS"
	// MaterialSource_CREATIVE_TOOL_SMART_VIDEO 创意工具-微电影
	MaterialSource_CREATIVE_TOOL_SMART_VIDEO MaterialSource = "CREATIVE_TOOL_SMART_VIDEO"
	// MaterialSource_OTHERS 其它
	MaterialSource_OTHERS MaterialSource = "OTHERS"
)
