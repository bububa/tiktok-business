package enum

// Destination 目标页类型。
type Destination string

const (
	// Destination_APP：所推广 App 在Google Play 或 Apple Store 的下载页面。广告对应的广告组中promotion_type 设置为APP_ANDROID 或 APP_IOS。
	Destination_APP Destination = "APP"
	// Destination_TIKTOK_INSTANT_PAGE: TikTok 即时体验页面，包括自定义页面、线索表单（即时体验表单）、应用介绍页和即时体验商品页面。广告中指定了TikTok即时体验页面的页面 ID（page_id）。
	Destination_TIKTOK_INSTANT_PAGE Destination = "TIKTOK_INSTANT_PAGE"
	// Destination_WEBSITE：网站页面。广告中指定了落地页（landing_page_url）。
	Destination_WEBSITE Destination = "WEBSITE"
	// Destination_SOCIAL_MEDIA_APP：社交媒体 URL。若想了解支持将目标页面设置为社交媒体 URL 的广告类型，请查看创建推广对象类型为即时通讯应用的线索广告。
	Destination_SOCIAL_MEDIA_APP Destination = "SOCIAL_MEDIA_APP"
	// Destination_PHONE_CALL：电话号码。若想了解支持将目标页面设置为电话号码的广告类型，请查看创建推广对象类型为电话通话的线索广告。
	Destination_PHONE_CALL Destination = "PHONE_CALL"
)
