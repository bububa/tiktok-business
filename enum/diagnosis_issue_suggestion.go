package enum

// DiagnosisIssueSuggestion 问题建议。
type DiagnosisIssueSuggestion string

const (
	// DiagnosisIssueSuggestion_NOBGM: 视频无背景音乐。您需为视频添加背景音乐。可使用 /creative/quick_optimization/create/为视频创建一键优化任务。
	DiagnosisIssueSuggestion_NOBGM DiagnosisIssueSuggestion = "NOBGM"
	// DiagnosisIssueSuggestion_VIDEO_LENGTH：视频时长过短。需编辑视频增加时长。可使用 /creative/quick_optimization/create/为视频创建一键优化任务。
	DiagnosisIssueSuggestion_VIDEO_LENGTH DiagnosisIssueSuggestion = "VIDEO_LENGTH"
	// DiagnosisIssueSuggestion_VIDEO_RESOLUTION：视频分辨率过低。需替换为高分辨率视频。
	DiagnosisIssueSuggestion_VIDEO_RESOLUTION DiagnosisIssueSuggestion = "VIDEO_RESOLUTION"

	// DiagnosisIssueSuggestion_SUGGEST_BID: 建议更改出价
	DiagnosisIssueSuggestion_SUGGEST_BID DiagnosisIssueSuggestion = "SUGGEST_BID"
	// DiagnosisIssueSuggestion_SUGGEST_BUDGET: 建议更改预算
	DiagnosisIssueSuggestion_SUGGEST_BUDGET DiagnosisIssueSuggestion = "SUGGEST_BUDGET"
	// DiagnosisIssueSuggestion_NOBID_SWITCH: 切换为最大投放量出价策略
	DiagnosisIssueSuggestion_NOBID_SWITCH DiagnosisIssueSuggestion = "NOBID_SWITCH"
	// DiagnosisIssueSuggestion_BUDGET_EDR：预算预估成效建议。若想了解预估成效详情，可查看预估成效。
	DiagnosisIssueSuggestion_BUDGET_EDR DiagnosisIssueSuggestion = "BUDGET_EDR"
	// DiagnosisIssueSuggestion_BID_EDR：出价预估成效建议。
	DiagnosisIssueSuggestion_BID_EDR DiagnosisIssueSuggestion = "BID_EDR"

	// DiagnosisIssueSuggestion_PIXEL：该 Pixel 过去 7 日内无任何活动。请检查您的 Pixel 设置。
	DiagnosisIssueSuggestion_PIXEL DiagnosisIssueSuggestion = "PIXEL"
)
