package enum

// PostbackWindowMode 选定的模式决定您想要确保能接收到的 SKAN (SKAdNetwork) 4.0 回传。
type PostbackWindowMode string

const (
	// POSTBACK_WINDOW_MODE1：此模式确保能接收到第一次回传，对应的归因窗口期为 0-2 天。回传数据最长需要 4 天时间接收，推广系列配额释放的时间最长也为 4 天。
	POSTBACK_WINDOW_MODE1 PostbackWindowMode = "POSTBACK_WINDOW_MODE1"
	// POSTBACK_WINDOW_MODE2：此模式确保能接收到第一次及第二次回传，对应的归因窗口期为 3-7 天。回传数据最长需要 13 天时间接收，推广系列配额释放的时间最长也为 13 天。
	POSTBACK_WINDOW_MODE2 PostbackWindowMode = "POSTBACK_WINDOW_MODE2"
	// POSTBACK_WINDOW_MODE3：此模式确保能接收到三次回传，对应的归因窗口期为 8-35 天。回传数据最长需要 41 天时间接收，推广系列配额释放的时间最长也为 41 天。
	POSTBACK_WINDOW_MODE3 PostbackWindowMode = "POSTBACK_WINDOW_MODE3"
)
