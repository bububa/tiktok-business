package enum

// VerticalVideoStrategy 商品销量场景中使用的视频类型。
type VerticalVideoStrategy string

const (
	// VerticalVideoStrategy_UNSET（未设置）
	VerticalVideoStrategy_UNSET VerticalVideoStrategy = "UNSET"
	// VerticalVideoStrategy_SINGLE_VIDEO（单个视频）
	VerticalVideoStrategy_SINGLE_VIDEO VerticalVideoStrategy = "SINGLE_VIDEO"
	// VerticalVideoStrategy_CATALOG_VIDEOS（使用视频模板的商品库视频）
	VerticalVideoStrategy_CATALOG_VIDEOS VerticalVideoStrategy = "CATALOG_VIDEOS"
	// VerticalVideoStrategy_CATALOG_UPLOADED_VIDEOS（使用已上传视频的商品库视频）
	VerticalVideoStrategy_CATALOG_UPLOADED_VIDEOS VerticalVideoStrategy = "CATALOG_UPLOADED_VIDEOS"
	// VerticalVideoStrategy_LIVE_STREAM（直播视频）。
	VerticalVideoStrategy_LIVE_STREAM VerticalVideoStrategy = "LIVE_STREAM"
)
