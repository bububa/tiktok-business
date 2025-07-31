package enum

// Window 窗口期
type Window string

const (
	// Window_OFF：不启用。
	Window_OFF Window = "OFF"
	// Window_ZERO_DAY 第0天
	Window_ZERO_DAY Window = "ZERO_DAY"
	// Window_ONE_DAY：1天
	Window_ONE_DAY Window = "ONE_DAY"
	// Window_SEVEN_DAYS：7天
	Window_SEVEN_DAYS Window = "SEVEN_DAYS"
	// Window_FOURTEEN_DAYS：14天
	Window_FOURTEEN_DAYS Window = "FOURTEEN_DAYS"
	// Window_TWENTY_EIGHT_DAYS：28天
	Window_TWENTY_EIGHT_DAYS Window = "TWENTY_EIGHT_DAYS"
)
