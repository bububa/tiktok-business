package enum

// MiniSeriesProductionType 制作类型，即短剧是本地短剧还是翻译短剧。
type MiniSeriesProductionType string

const (
	// MiniSeriesProductionType_LOCAL: 本地短剧。本地演员、本地语言、旁白和字幕。
	MiniSeriesProductionType_LOCAL MiniSeriesProductionType = "LOCAL"
	// MiniSeriesProductionType_TRANSLATION: 翻译短剧，为国际发行，附加字幕、旁白和配音。
	MiniSeriesProductionType_TRANSLATION MiniSeriesProductionType = "TRANSLATION"
)
