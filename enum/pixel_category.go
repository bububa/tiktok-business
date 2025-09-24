package enum

// PixelCategory Pixel 追踪事件类型。
type PixelCategory string

const (
	// PixelCategory_ONLINE_STORE ：电商事件。
	PixelCategory_ONLINE_STORE PixelCategory = "ONLINE_STORE"
	// PixelCategory_FILLING_FORM ：表单事件。
	PixelCategory_FILLING_FORM PixelCategory = "FILLING_FORM"
	// PixelCategory_CONTACTS ：联系事件。
	PixelCategory_CONTACTS PixelCategory = "CONTACTS"
	// PixelCategory_LANDING_PAGE ：APK 下载事件。
	PixelCategory_LANDING_PAGE PixelCategory = "LANDING_PAGE"
	// PixelCategory_CUSTOMIZE_EVENTS ：自定义事件。
	PixelCategory_CUSTOMIZE_EVENTS PixelCategory = "CUSTOMIZE_EVENTS"
)
