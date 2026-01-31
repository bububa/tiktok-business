package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Ad Ad info
type Ad struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName 该广告组所属的推广系列的名称。
	CampaignName string `json:"campaign_name,omitempty"`
	// AdgroupID 广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName 广告组名称
	AdgroupName string `json:"adgroup_name,omitempty"`
	// SmartPlusAdID 广告 ID
	SmartPlusAdID string `json:"smart_plus_ad_id,omitempty"`
	// AdName 广告名称
	AdName string `json:"ad_name,omitempty"`
	// OperationStatus 广告的操作状态。
	// ENABLE：广告处于启用（“开”）状态。
	// DISABLE：广告处于未启用（“关”）状态。
	// FROZEN：广告处于冻结状态，无法再次启用。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// SecondaryStatus 广告状态（二级状态）。枚举值详见枚举值-二级状态-广告状态。
	// 注意：沙箱环境下本字段不返回，因为广告未实际投放。
	SecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// CreateTime 广告创建时间。格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2023-01-01 00:00:01".
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ModifyTime 广告修改时间。格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2023-01-01 00:00:01".
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
	// CreativeList A list of creatives
	CreativeList []Creative `json:"creative_list,omitempty"`
	// AdTextList A list of ad texts.
	// Ad texts are shown to your audience as part of your ad creatives, to deliver the message you intend to communicate to them.
	AdTextList []AdText `json:"ad_text_list,omitempty"`
	// AutoMessageList Returned only when promotion_type is LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE at the ad group level.
	// Details of the automatic message to use in a TikTok Direct Messaging Ad.
	AutoMessageList []AutoMessage `json:"auto_message_list,omitempty"`
	// CallToActionList Call-to-action list
	CallToActionList []CallToAction `json:"call_to_action,omitempty"`
	// InteractiveAddOnList A list of interactive add-on (creative portfolio) IDs.
	InteractiveAddOnList []InteractiveAddOn `json:"interactive_add_on_list,omitempty"`
	// PageList A list of pages
	PageList []Page `json:"page_list,omitempty"`
	// LandingPageURLList A list of landing page URLs
	LandingPageURLList []LandingPageURL `json:"landing_page_url_list,omitempty"`
	// CustomProductPageList Details of the custom product page to use in the ad
	CustomProductPageList []CustomProductPage `json:"custom_product_page_url,omitempty"`
	// DeeplinkList A list of deeplinks
	DeeplinkList []Deeplink `json:"deeplink_list,omitempty"`
	// Disclaimer Disclaimer information
	Disclaimer *Disclaimer `json:"disclaimer,omitempty"`
	// AdConfiguration Additional configurations
	AdConfiguration *AdConfiguration `json:"ad_configuration,omitempty"`
}

// AdText Ad texts are shown to your audience as part of your ad creatives, to deliver the message you intend to communicate to them.
type AdText struct {
	// AdText Ad text.
	AdText string `json:"ad_text,omitempty"`
}

// AutoMessage Details of the automatic message to use in a TikTok Direct Messaging Ad.
type AutoMessage struct {
	// AutoMessageID The ID of the automatic message to use in a TikTok Direct Messaging Ad.
	// Currently, the only supported automatic message type is welcome message.
	// To obtain a list of welcome messages within your ad account, use /creative/auto_message/get/.
	// To create a welcome message within your ad account, use /creative/auto_message/create/.
	// To learn more about how to create TikTok Direct Messaging Ads with welcome messages, see Create an Upgraded Smart+ Lead Generation Campaign with optimization location as TikTok direct messages.
	AutoMessageID string `json:"auto_message_id,omitempty"`
}

type CallToAction struct {
	// CallToAction Call-to-action text.
	// For enum values, see Enumeration - Call-to-action.
	CallToAction enum.CallToAction `json:"call_to_action,omitempty"`
}

type InteractiveAddOn struct {
	// CardID The ID of an interactive add-on
	CardID string `json:"card_id,omitempty"`
}

type Page struct {
	// PageID Page ID
	PageID string `json:"page_id,omitempty"`
}

type LandingPageURL struct {
	// LandingPageURL Landing page URL
	LandingPageURL string `json:"landing_page_url,omitempty"`
}

// CustomProductPage Details of the custom product page to use in the ad
type CustomProductPage struct {
	// CustomProductPageURL The Custom Product Page (CPP) URL.
	// CPPs are product pages with customized screenshots, promotional text, and app previews that can be used to highlight specific content or features of the promoted app, or reach specific audiences.
	// Example: https://apps.apple.com/us/app/tiktok/id835599320?ppid=12345a6b-c789-12d3-e4f5-g6h78i91jk2l
	CustomProductPageURL string `json:"custom_product_page_url,omitempty"`
}

type Deeplink struct {
	// Deeplink Deeplink, the specific location where you want your audience to go if they have your app installed.
	Deeplink string `json:"deeplink,omitempty"`
	// DeeplinkType The deeplink type.
	// Enum values:
	// NORMAL: non-deferred deeplink.
	// DEFERRED_DEEPLINK: deferred deeplink.
	DeeplinkType enum.DeeplinkType `json:"deeplink_type,omitempty"`
}

// Disclaimer Disclaimer information
type Disclaimer struct {
	// DisclaimerType The type of disclaimer in the ad.
	// Enum values:
	// TEXT_LINK: clickable disclaimers.
	// TEXT_ONLY: text-only disclaimers.
	DisclaimerType enum.DisclaimerType `json:"disclaimer_type,omitempty"`
	// DisclaimerText The text-only disclaimer in the ad
	DisclaimerText *DisclaimerText `json:"disclaimer_text,omitempty"`
	// DisclaimerClickableTexts The clickable disclaimer or clickable disclaimers in the ad
	DisclaimerClickableTexts []DisclaimerClickableText `json:"disclaimer_clickable_texts,omitempty"`
}

// DisclaimerText The text-only disclaimer in the ad
type DisclaimerText struct {
	// Text The disclaimer text
	Text string `json:"text,omitempty"`
}

// DisclaimerClickableText The clickable disclaimer or clickable disclaimers in the ad
type DisclaimerClickableText struct {
	// Text The disclaimer text
	Text string `json:"text,omitempty"`
	// URL The URL for the clickable disclaimer.
	// When users tap each text, they will be redirected to the URL and see more disclaimer details.
	URL string `json:"url,omitempty"`
}

// AdConfiguration Additional configurations
type AdConfiguration struct {
	// IdentityType Identity type when you are creating non-Spark Ads.
	// Enum values: CUSTOMIZED_USER.
	// For details about identities, see Identities.
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityID Identity ID when you are creating non-Spark Ads.
	IdentityID string `json:"identity_id,omitempty"`
	// DarkPostStatus The status of the "Ads-only mode" for your creatives.
	// Enum values:
	// ON: Enable the ads-only mode to limit your posts to paid traffic.
	// OFF: Disable the ads-only mode. The post will appear on your TikTok profile and will be eligible to receive organic traffic.
	DarkPostStatus enum.OnOff `json:"dark_post_status,omitempty"`
	// ProductSpecificType Different dimensions to choose products.
	// Enum values:
	// ALL: Allow TikTok to dynamically choose from all products.
	// PRODUCT_SET: Specify a product set. TikTok will dynamically choose products from this set.
	// CUSTOMIZED_PRODUCTS: Specify a customized number of products.
	ProductSpecificType enum.ProductSpecificType `json:"product_specific_type,omitempty"`
	// ProductSetID The ID of a product set
	ProductSetID string `json:"product_set_id,omitempty"`
	// ProductIDs The product IDs of catalog products
	ProductIDs []string `json:"product_ids,omitempty"`
	// CatalogCreativeToggle Whether to enable auto-selection of creatives from your catalog.
	// Supported values: true, false.
	CatalogCreativeToggle *bool `json:"catalog_creative_toggle,omitempty"`
	// CreativeAutoAddToggle Whether to auto-add newly generated assets during delivery.
	// During your campaign, the system will generate new creative assets from your existing campaign assets and creative library. By using TikTok's latest trending elements (hooks, music, text, and more), you'll get high-quality, ready-to-use content that delivers alongside your other selected assets.
	// New assets will be added to your campaign automatically and won't appear on your TikTok profile. You can turn off this feature anytime.
	// Supported values: true, false.
	CreativeAutoAddToggle *bool `json:"creative_auto_add_toggle,omitempty"`
	// CreativeAutoEnhancementStrategyList The list of automatic enhancement strategies to apply to your ads.
	// Automatic enhancements are real-time edits applied to your ads during your campaign. They can improve performance by creating more engaging and impactful visuals, sound, and more.
	// Enum values:
	// TRANSLATE_AND_DUB: Translate and dub. Connect with global audiences by delivering your ad in 50+ languages.
	// MUSIC_REFRESH: Music refresh. Stay on-trend by swapping in music currently popular on TikTok.
	// VIDEO_QUALITY: Video quality. Improve overall visual quality with increased resolution and clarity.
	// IMAGE_QUALITY: Image quality. Improve overall visual quality with increased resolution and clarity.
	// IMAGE_RESIZE: Resize. Resize your image to take advantage of full-screen capabilities.
	// CALL_TO_ACTION_ENHANCEMENT: CTA enhancement. Tailor and optimize your call to action based on specific use cases to enhance performance.
	// AIGC_CARD: Generate ad card. Increase engagement by showcasing key messages and visuals designed to drive clicks and conversions.
	// Note: Currently, you cannot enable the strategies CALL_TO_ACTION_ENHANCEMENT and AIGC_CARD through creative_auto_enhancement_strategy_list via API.
	CreativeAutoEnhancementStrategyList enum.CreativeAutoEnhancementStrategy `json:"creative_auto_enhancement_strategy_list,omitempty"`
	// DeeplinkUTMParams A list of deeplink URL parameters. URL parameters are snippets of code that can be added to the end of the URLs to help you track clicks across different channels and understand how visitors interact with a website through third-party analytics platforms. They consist of key-value pairs that are specified through key and value.
	DeeplinkUTMParams []model.KV `json:"deeplink_utm_params,omitempty"`
	// EndCardCTA Call-to-action for the end card (image_ids) of an Automotive Carousel Ad for Inventory or Upgraded Smart+ Automotive Carousel Ad for Models, or the catalog solution for Upgraded Smart+ TikTok Direct Messaging Ads.
	// Enum values:
	// SEARCH_INVENTORY (Recommended): Search inventory.
	// LEARN_MORE: Learn more.
	// SHOP_NOW: Shop now.
	// SIGN_UP: Sign up.
	// CONTACT_US: Contact us.
	// BOOK_NOW: Book now.
	// READ_MORE: Read more.
	// VIEW_MORE: View now.
	// ORDER_NOW: Order now.
	// SEND_MESSAGE: Send message.
	// This value is only available for the catalog solution for Upgraded Smart+ TikTok Direct Messaging Ads.
	EndCardCTA string `json:"end_card_cta,omitempty"`
	// ProductDisplayFieldList A list of product details to display in your Automotive Carousel Ad for Inventory.
	// Enum values:
	// DEALER_NAME: the dealer_name field of vehicles.
	// MAKE: the make field of vehicles.
	// MODEL: the model field of vehicles.
	// YEAR: the year field of vehicles.
	// MILEAGE: the mileage field of vehicles.
	// PRICE: the price field of vehicles.
	// SALE_PRICE: the sale_price field of vehicles.
	// EXTERIOR_COLOR: the exterior_color field of vehicles.
	// TRIM: the trim field of vehicles.
	// ADDRESS_CITY: the city field of vehicles.
	// VEHICLE_STATE: the state_of_vehicle field of vehicles.
	// The title of each vehicle is displayed automatically. For instance, if you set this field to ["MILEAGE","YEAR"], the ad will display the mileage, price and title fields of each vehicle.
	// To retrieve the product detail fields of your vehicles in an Auto-Inventory catalog and check whether you need to update the details, use /catalog/product/get/.
	ProductDisplayFieldList []string `json:"product_display_field_list,omitempty"`
	// AutoDisclaimerTypes The type of disclaimer to show in the Automotive Carousel Ad for Models.
	// Enum values:
	// EMISSION: Emission disclaimer.
	// DISCOUNT: Discount or offer disclaimer.
	AutoDisclaimerTypes []enum.AutoDisclaimerType `json:"auto_disclaimer_types,omitempty"`
	// UTMParams A list of URL parameters.
	// URL parameters are snippets of code that can be added to the end of the URLs to help you track clicks across different channels and understand how visitors interact with a website through third-party analytics platforms. They consist of key-value pairs that are specified through key and value
	UTMParams []model.KV `json:"utm_params,omitempty"`
	// FallbackType Fallback type.
	// The destination when the user is unable to access the deeplink because the app has not been installed.
	// Enum value:
	// WEBSITE: The promoted web page as fallback URL. You need to simultaneously specify the fallback URL through landing_page_url_list.
	// To learn more about how to configure a fallback URL in TikTok Instant Messaging Ads, see Create an Upgraded Smart+ Lead Generation Campaign with optimization location as instant messaging apps.
	FallbackType enum.FallbackType `json:"fallback_type,omitempty"`
	// ProductInfo Product information.
	// This information will be used in different ad variations to create personalized ad delivery with the goal of improving ad performance. Based on past data, ads with complete information often have more clicks and conversions. Results may vary.
	ProductInfo *ProductInfo `json:"product_info,omitempty"`
	// CallToActionID The ID of the call-to-action portfolio to use in your ad
	CallToActionID string `json:"call_to_action_id,omitempty"`
	// PhoneInfo Returned only when promotion_type is set to LEAD_GEN_CLICK_TO_CALL at the ad group level.
	// Details of the phone number that the ad audience can directly reach out to you (the advertiser) through when they click the call-to-action button within the ad.
	// Learn about how to create an Upgraded Smart+ Lead Generation Campaign with optimization location as phone call.
	PhoneInfo *PhoneInfo `json:"phone_info,omitempty"`
	// TrackingInfo Tracking information
	TrackingInfo *TrackingInfo `json:"tracking_info,omitempty"`
}

// ProductInfo Product information.
type ProductInfo struct {
	// SellingPoints A list of selling points
	SellingPoints []string `json:"selling_points,omitempty"`
}

// PhoneInfo Details of WhatsApp or Zalo phone number
type PhoneInfo struct {
	// PhoneRegionCode The region code for WhatsApp or Zalo phone number.
	PhoneRegionCode string `json:"phone_region_code,omitempty"`
	// PhoneRegionCallingCode The region calling code for the WhatsApp or Zalo phone number.
	PhoneRegionCallingCode string `json:"phone_region_calling_code,omitempty"`
	// PhoneNumber The WhatsApp or Zalo phone number
	PhoneNumber string `json:"phone_number,omitempty"`
}

// TrackingInfo Tracking information
type TrackingInfo struct {
	// ImpressionTrackingURL Default Impression Tracking URL
	ImpressionTrackingURL *string `json:"impression_tracking_url,omitempty"`
	// ClickTrackingURL Click Tracking URL
	ClickTrackingURL *string `json:"click_tracking_url,omitempty"`
	// TrackingAppID The ID of the app that is measured
	TrackingAppID *string `json:"tracking_app_id,omitempty"`
	// TrackingMessageEventSetID The ID of the message event set that you want to measure in the Instant Messaging Ad.
	TrackingMessageEventSetID *string `json:"tracking_message_event_set_id,omitempty"`
	// AppTrackingInfoList Details of Third-party tracking settings
	AppTrackingInfoList []AppTrackingInfo `json:"app_tracking_info_list,omitempty"`
}

// AppTrackingInfo Details of Third-party tracking setting
type AppTrackingInfo struct {
	// AppType Returned when app_tracking_info_list is specified.
	// App type.
	// Enum values:
	// APP_ANDROID: Android App.
	// APP_IOS: iOS App.
	AppType enum.AppType `json:"app_type,omitempty"`
	// AppID Returned when app_tracking_info_list is specified.
	// The App ID of the App.
	// When app_type is APP_ANDROID, this field represents the App ID for an Android App.
	// When app_type is APP_IOS, this field represents the App ID for an iOS App.
	AppID string `json:"app_id,omitempty"`
	// ImpressionTrackingURL Impression tracking URL.
	// When app_type is APP_ANDROID, this field represents an impression tracking URL for Android.
	// When app_type is APP_IOS, this field represents an impression tracking URL for iOS.
	ImpressionTrackingURL string `json:"impression_tracking_url,omitempty"`
	// ClickTrackingURL Click tracking URL.
	// When app_type is APP_ANDROID, this field represents a click tracking URL for Android.
	// When app_type is APP_IOS, this field represents a click tracking URL for iOS.
	ClickTrackingURL string `json:"click_tracking_url,omitempty"`
}
