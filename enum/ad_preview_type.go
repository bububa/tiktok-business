package enum

// AdPreviewType 预览类型。
type AdPreviewType string

const (
	// AdPreviewType_ADS_CREATION 根据广告创建参数预览广告
	AdPreviewType_ADS_CREATION AdPreviewType = "ADS_CREATION"
	// AdPreviewType_AD 预览已有的常规广告
	AdPreviewType_AD AdPreviewType = "AD"
	// AdPreviewType_CARD 预览创新互动样式
	AdPreviewType_CARD AdPreviewType = "CARD"
	// AdPreviewType_PAGE 预览商品落地页
	AdPreviewType_PAGE AdPreviewType = "PAGE"
	// AdPreviewType_SINGLE_VIDEO 根据参数预览视频广告
	AdPreviewType_SINGLE_VIDEO AdPreviewType = "SINGLE_VIDEO"
	// AdPreviewType_SINGLE_IMAGE 根据参数预览图片广告
	AdPreviewType_SINGLE_IMAGE AdPreviewType = "SINGLE_IMAGE"
)
