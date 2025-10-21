package enum

// MiniSeriesTargetAudience 短剧的目标受众。
type MiniSeriesTargetAudience string

const (
	// MiniSeriesTargetAudience_MALE: 男频。目标受众主体为男性，更多地围绕男主人公或从男性视角叙事强调权力财富声望的积累、强大能力的提升、运筹帷幄与设计营谋。
	MiniSeriesTargetAudience_MALE MiniSeriesTargetAudience = "MALE"
	// MiniSeriesTargetAudience_FEMALE: 女频。目标受众主体为女性，更多地围绕女主人公或从女性视角叙事，大多含有浪漫爱情元素。
	MiniSeriesTargetAudience_FEMALE MiniSeriesTargetAudience = "FEMALE"
	// MiniSeriesTargetAudience_NEUTRAL: 中立。没有明确目标受众性别。
	MiniSeriesTargetAudience_NEUTRAL MiniSeriesTargetAudience = "NEUTRAL"
)
