package enum

// ContentLineupType 内容相关定向标签类型。
type ContentLineupType string

const (
	// ContentLineupType_MAX_PULSE（Max Pulse：您的广告将在 TikTok 上任意主题的热门内容前后展示。）
	ContentLineupType_MAX_PULSE ContentLineupType = "MAX_PULSE"
	// ContentLineupType_CUSTOM（类别分组：您的广告将在所选类别或季节性活动的热门内容前后展示。）
	ContentLineupType_CUSTOM ContentLineupType = "CUSTOM"
)
