package enum

// AdFormat 素材类型
type AdFormat string

const (
	// AdFormat_SINGLE_IMAGE 单个图片。
	AdFormat_SINGLE_IMAGE AdFormat = "SINGLE_IMAGE"
	// AdFormat_SINGLE_VIDEO 单个视频。
	AdFormat_SINGLE_VIDEO AdFormat = "SINGLE_VIDEO"
	// AdFormat_CAROUSEL 已废弃 多图轮播，支持TopBuzz版位。
	AdFormat_CAROUSEL AdFormat = "CAROUSEL"
	// AdFormat_LIVE_CONTENT 直播内容。仅适用于直播带货（promotion_type = LIVE_SHOPPING）场景。
	AdFormat_LIVE_CONTENT AdFormat = "LIVE_CONTENT"
	// AdFormat_CAROUSEL_ADS 多图轮播，支持TikTok版位。
	AdFormat_CAROUSEL_ADS AdFormat = "CAROUSEL_ADS"
)
