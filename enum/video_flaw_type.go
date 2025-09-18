package enum

// VideoFlawType 视频问题种类
type VideoFlawType string

const (
	// VideoFlawType_LOW_RESOLUTION：视频分辨率低于540 x 960 像素, 不符合要求。
	VideoFlawType_LOW_RESOLUTION VideoFlawType = "LOW_RESOLUTION"
	// VideoFlawType_ILLEGAL_VIDEO_SIZE：视频尺寸的比例不符合竖版（9:16），方版（1:1），横版（16:9）的要求。
	VideoFlawType_ILLEGAL_VIDEO_SIZE VideoFlawType = "ILLEGAL_VIDEO_SIZE"
	// VideoFlawType_NO_BGM（已废弃）：广告或视频没有背景音乐，或背景音乐不连贯/不清晰。
	VideoFlawType_NO_BGM VideoFlawType = "NO_BGM"
	// VideoFlawType_BLACK_EDGE（已废弃）：视频画面中出现影响用户体验的黑边。
	VideoFlawType_BLACK_EDGE VideoFlawType = "BLACK_EDGE"
	// VideoFlawType_ILLEGAL_DURATION（已废弃）：视频长度长于 60 秒或短于 5 秒，不符合要求。
	VideoFlawType_ILLEGAL_DURATION VideoFlawType = "ILLEGAL_DURATION"
)
