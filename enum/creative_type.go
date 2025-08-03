package enum

// CreativeType 直播购物广告、商品购物广告或应用预注册场景中的创意素材类型。枚举值：
type CreativeType string

const (
	// CreativeType_SHOP_PDP：短视频带货广告，包含一个商品链接。
	CreativeType_SHOP_PDP CreativeType = "SHOP_PDP"
	// CreativeType_SHOP_PLP: 短视频带货广告，包含6-50个商品链接。
	CreativeType_SHOP_PLP CreativeType = "SHOP_PLP"
	// CreativeType_SHORT_VIDEO_LIVE: 短视频引流至直播间广告。
	CreativeType_SHORT_VIDEO_LIVE CreativeType = "SHORT_VIDEO_LIVE"
	// CreativeType_DIRECT_LIVE: 直播间带货。每个推广系列只能有一个直播间。
	CreativeType_DIRECT_LIVE CreativeType = "DIRECT_LIVE"
	// CreativeType_PSA：商品视觉元素将作为广告创意使用。
	CreativeType_PSA CreativeType = "PSA"
	// CreativeType_CUSTOM_INSTANT_PAGE：自定义 TikTok 即时体验页面。
	CreativeType_CUSTOM_INSTANT_PAGE CreativeType = "CUSTOM_INSTANT_PAGE"
)
