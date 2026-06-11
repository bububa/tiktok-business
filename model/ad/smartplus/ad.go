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
	// CatalogCreativeInfo Valid only when the following conditions are met:
	// The ad is an Upgraded Smart+ E-commerce Catalog Ad.
	// catalog_creative_toggle is true.
	// Additional settings for catalog creatives to use in your ads.
	CatalogCreativeInfo *CatalogCreativeInfo `json:"catalog_creatiev_info,omitempty"`
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
	// ProductInfoEnabled The product information mode.
	// Enum values:
	// UNSET: To disable product information mode.
	// NON_CATALOG: To enable the non-catalog version of product information.
	// CATALOG: To enable the catalog version of product information.
	// Default value: UNSET.
	// To learn about the scenarios where you need to set this parameter, see Available product information settings in Upgraded Smart+ Campaigns.
	// Note:
	// This field will automatically be set to CATALOG when the following conditions are both met at the campaign level:
	// app_promotion_type is MINIS.
	// catalog_enabled is true.
	// This field can only be set to UNSET when an Inventory Card portfolio is specified through card_id in Upgraded Smart+ Automotive Ads.
	ProductInfoEnabled string `json:"product_info_enabled,omitempty"`
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
	// ProductTitles A list of product names.
	// Size range: 0-1.
	// Length limit for each product name: 40 characters and cannot contain emoji. Each word in Chinese or Japanese counts as two characters, while each letter in English counts as one character.
	// If you don't provide a product name, a default name will be used so that the name won't appear blank.
	ProductTitles []string `json:"product_titles,omitempty"`
	// ProductImageList A list of product images.
	// Size range: 0-10.
	// If you don't provide a product image, a generic image will be used so that the image won't appear blank.
	ProductImageList []ProductImage `json:"product_image_list,omitempty"`
	// SellingPoints A list of selling points
	SellingPoints []string `json:"selling_points,omitempty"`
	// CatalogTagList Valid in any of the following scenarios:
	// Scenario 1: At the campaign level objective_type is WEB_CONVERSIONS and catalog_type is ECOMMERCE.
	// Scenario 2:
	// At the campaign level objective_type is LEAD_GENERATION and catalog_enabled is true.
	// At the ad group level an Auto-Inventory catalog is specified through catalog_id.
	// Details of the product information for catalog ads.
	// For scenario 1, the specified product information will be automatically displayed from your catalog to drive results. The enum values are:
	// PRICE: Price.
	// STRIKETHROUGH_PRICE: Strikethrough price.
	// DISCOUNT: Discount.
	// FREE_SHIPPING: Free shipping.
	// Default value: ["PRICE", "STRIKETHROUGH_PRICE", "DISCOUNT", "FREE_SHIPPING"].
	// Max size: 4.
	// For scenario 2, the corresponding text for the specified product information will be displayed on the inventory card. The enum values are:
	// DEALER_NAME: Dealer name.
	// CURRENT_MILEAGE: Current mileage.
	// LEAD_PRICE: Price.
	// LEAD_SALE_PRICE: Sale price.
	// EXTERIOR_COLOR: Exterior color.
	// TRIM: Trim.
	// ADDRESS_CITY: City.
	// VEHICLE_STATE: Vehicle state.
	// Max size: 2 for primary catalog tags (DEALER_NAME,CURRENT_MILEAGE,EXTERIOR_COLOR,TRIM,ADDRESS_CITY,VEHICLE_STATE) and 2 for secondary catalog tags (LEAD_PRICE,LEAD_SALE_PRICE).
	CatalogTagList []string `json:"catalog_tag_list,omitempty"`
	// PromoInfoList Valid only for regular Upgraded Smart+ Web Campaigns.
	// Details of promo codes and offers.
	// Your offer details for promo codes, offers, or events will be highlighted to boost engagement and ad performance. Promo codes require shoppers to enter a code at checkout. Offers apply automatically and no code is needed.
	// Max size: 10.
	PromoInfoList []PromoInfo `json:"promo_info_list,omitempty"`
}

type ProductImage struct {
	// WebURI The image ID of a product image.
	// Specifications requirements:
	// Image format: .jpg, .png, .jpeg, or .webp.
	// File size: no more than 5 MB
	// Image size: no limit.
	// Recommend image resolution: 800 x 800 pixels.
	// To obtain the image ID, use the /file/image/ad/search/ or /file/image/ad/upload/ endpoint.
	WebURI string `json:"web_uri,omitempty"`
}

type PromoInfo struct {
	// DiscountType Required when promo_info_list is specified.
	// Discount type.
	// Enum values:
	// PERCENTAGE: Percentage off discount.
	// CASH: Cash off discount.
	DiscountType string `json:"discount_type,omitempty"`
	// DiscountValue Required when promo_info_list is specified.
	// Discount value.
	// When discount_type is PERCENTAGE, specify an integer between 1-100.
	// When discount_type is CASH, specify a float greater than 0.
	DiscountValue float64 `json:"discount_value,omitempty"`
	// DiscountCurrency Required when discount_type is CASH.
	// Discount currency.
	// For enum values, see List of values for discount_currency or minimum_purchase_currency.
	DiscountCurrency string `json:"discount_currency,omitempty"`
	// PromoCode The promo code.
	// For promo codes that require shoppers to enter a code at checkout, specify the code through this field.
	// For offers that apply automatically, omit this field.
	// Length limit: 30 characters long and cannot contain emoji. Each word in Chinese or Japanese counts as two characters, while each letter in English counts as one character.
	// Note: Promo codes exclusive to TikTok often perform best.
	PromoCode string `json:"promo_code,omitempty"`
	// MinimumPurchaseType Minimum purchase type.
	// Enum values:
	// QUANTITY: Minimum quantity.
	// SUBTOTAL: Minimum subtotal.
	MinimumPurchaseType string `json:"minimum_purchase_type,omitempty"`
	// MinimumPurchaseValue Required when minimum_purchase_type is specified.
	// Minimum purchase value.
	// When minimum_purchase_type is QUANTITY, specify an integer that is 0 or greater.
	// When minimum_purchase_type is SUBTOTAL, specify a float greater than 0.
	MinimumPurchaseValue *float64 `json:"minimum_purchase_value,omitempty"`
	// MinimumPurchaseCurrency Required when minimum_purchase_type is SUBTOTAL.
	// Minimum purchase currency.
	// For enum values, see List of values for discount_currency or minimum_purchase_currency.
	MinimumPurchaseCurrency string `json:"minimum_purchase_currency,omitempty"`
	// ValidStartTime Valid start time (UTC+0) for the promo code or offer, in the format of YYYY-MM-DD HH:MM:SS
	ValidStartTime model.DateTime `json:"valid_start_time,omitzero"`
	// ValidEndTime Valid end time (UTC+0) for the promo code or offer, in the format of YYYY-MM-DD HH:MM:SS.
	// The valid_end_time should be later than the current time.
	// If you specify valid_end_time, provide valid_start_time at the same time.
	ValidEndTime model.DateTime `json:"valid_end_time,omitzero"`
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
	// ViewabilityPostbidPartner Valid only in regular Upgraded Smart+ Web Campaigns and Upgraded Smart+ E-commerce Catalog Ads.
	// Post-bid third-party viewability measurement partner.
	// Enum values: DOUBLE_VERIFY, IAS, ZEFR.
	// Note:
	// Post-bid third-party viewability measurement is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative.
	// See Brand safety to learn about the supported advertising objectives, supported markets, and the general introduction of post-bid measurement.
	// When you pass in both brand_safety_postbid_partner and viewability_postbid_partner:
	// If you set either of these two fields as IAS, you need to specify both as IAS, and assign the same value (the wrapped VAST URL that you obtain from the measurement partner IAS) to brand_safety_vast_url and viewability_vast_url.
	// If you set either of these two fields as ZEFR, you need to specify both as ZEFR.
	// If you set two different post-bid measurement partners, only the post-bid brand safety measurement setting will become valid.
	ViewabilityPostbidPartner *enum.ViewabilityPostbidPartner `json:"viewability_postbid_partner,omitempty"`
	// ViewabilityVastURL Required when viewability_postbid_partner is IAS.
	// The wrapped VAST URL used by the post-bid third-party partner to measure viewability. You need to get the URL from the measurement partner IAS.
	// Note:
	// Post-bid third-party viewability measurement is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative.
	// When you set both brand_safety_postbid_partner and viewability_postbid_partner as IAS, brand_safety_vast_url should be the same as viewability_vast_url.
	ViewabilityVastURL *string `json:"viewability_vast_url,omitempty"`
	// BrandSafetyPostbidPartner Valid only in regular Upgraded Smart+ Web Campaigns and Upgraded Smart+ E-commerce Catalog Ads.
	// Post-bid third-party brand safety measurement partner.
	// Enum values: DOUBLE_VERIFY, IAS, ZEFR.
	// Note:
	// Post-bid third-party brand safety measurement is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative.
	// See Brand safety to learn about the supported advertising objectives, supported markets, and the general introduction of post-bid measurement.
	// When you pass in both brand_safety_postbid_partner and viewability_postbid_partner:
	// If you set either of these two fields as IAS, you need to specify both as IAS, and assign the same value (the wrapped VAST URL that you obtain from the measurement partner IAS) to brand_safety_vast_url and viewability_vast_url.
	// If you set two different post-bid measurement partners, only the post-bid brand safety measurement setting will become valid.
	BrandSafetyPostbidPartner *enum.ViewabilityPostbidPartner `json:"brand_safety_postbid_partner,omitempty"`
	// BrandSafetyVastURL Required when brand_safety_postbid_partner is IAS.
	// The wrapped VAST URL used by the post-bid third-party partner to measure brand safety. You need to get the URL from the measurement partner IAS.
	// Note:
	// Post-bid third-party brand safety measurement is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative.
	// When you set both brand_safety_postbid_partner and viewability_postbid_partner as IAS, brand_safety_vast_url should be the same as viewability_vast_url.
	BrandSafetyVastURL *string `json:"brand_safety_vast_url,omitempty"`
	// ImpressionTrackingURL Default Impression Tracking URL
	ImpressionTrackingURL *string `json:"impression_tracking_url,omitempty"`
	// ClickTrackingURL Click Tracking URL
	ClickTrackingURL *string `json:"click_tracking_url,omitempty"`
	// TrackingAppID The ID of the app that is measured
	TrackingAppID *string `json:"tracking_app_id,omitempty"`
	// TrackingPixelID Valid only when the following conditions are both met at the campaign level:
	// objective_type is APP_PROMOTION.
	// app_promotion_type is APP_INSTALL or APP_RETARGETING.
	// The ID of the pixel for tracking the actions people take on your website, such as page views, add to cart, purchases, and more.
	// If pixel_id is already set at the ad group level, this tracking_pixel_id must match it for tracking website events.
	// If pixel_id is not set at the ad group level, you may specify any pixel ID for tracking website events.
	TrackingPixelID *string `json:"tracking_pixel_id,omitempty"`
	// TrackingOfflineEventSetIDs Valid only when the following conditions are both met at the campaign level:
	// objective_type is APP_PROMOTION.
	// app_promotion_type is APP_INSTALL or APP_RETARGETING.
	// A list of Offline Event set IDs for tracking customer actions that happen in-person, such as in-store purchases or live event attendance.
	// Max size: 50.
	// See here to learn more about how to create and manage Offline Event sets.
	// Note:
	// For a new ad, existing auto-tracking Offline Event sets have to be tracked. You can choose to omit or provide this field.
	// If you omit this field, it automatically includes all existing auto-tracking Offline Event set IDs within your ad account.
	// If you provide this field, you need to specify all existing auto-tracking Offline Event set IDs. Adding non-auto-tracking Offline Event set IDs is optional. To find existing auto-tracking Offline Event sets within an ad account, call /offline/get/ and select sets where auto_tracking is true.
	TrackingOfflineEventSetIDs []string `json:"tracking_offline_event_set_ids,omitempty"`
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

// CatalogCreativeInfo The ad is an Upgraded Smart+ E-commerce Catalog Ad.
type CatalogCreativeInfo struct {
	// CatalogMediaSettings The types of creatives from your E-commerce catalog to use in the ad.
	// Enum values:
	// VIDEO: Video.
	// IMAGE: Image.
	// TEMPLATE_VIDEO: Video templates. If you include this value in catalog_media_settings, you can optionally specify catalog_template_video_id at the same time.
	// If catalog_template_video_id is not specified, all video templates and video URLs from your catalog will be used to generate the ad.
	// If catalog_template_video_id is specified, the selected video template will be used to generate the ad.
	CatalogMediaSettings []enum.CatalogMediaSetting `json:"catalog_media_settings,omitempty"`
	// CatalogTemplateVideoID Valid only when TEMPLATE_VIDEO is included in catalog_media_settings.
	// The ID of a video template in your catalog to use in the ad.
	// To obtain the IDs of video templates (video packages) in your catalog, use /catalog/video_package/get/.
	// To learn about how to create video packages on TikTok Ads Manager, see How to create video packages in a Catalog.
	CatalogTemplateVideoID string `json:"catalog_template_video_id,omitempty"`
}
