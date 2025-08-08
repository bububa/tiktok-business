package enum

// FallbackType 失败路径。如果用户没有安装过应用，您可以选择让用户回退去安装应用，或者去您想推广的网页。
type FallbackType string

const (
	// FallbackType_APP_INSTALL：广告组层级通过 app_id 指定的推广应用对应的 App Store 或 Google Play 页面。
	FallbackType_APP_INSTALL FallbackType = "APP_INSTALL"
	// FallbackType_WEBSITE：通过 landing_page_url 指定的推广网页。
	FallbackType_WEBSITE FallbackType = "WEBSITE"
	// FallbackType_UNSET：未设置。
	FallbackType_UNSET FallbackType = "UNSET"
)
