package ad

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Ad 广告
type Ad struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName 该广告组所属的推广系列的名称。
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignSystemOrigin 仅为内容加热推广系列返回本字段。
	// 非内容加热推广系列不返回本字段。
	//
	// 推广系列来源。
	//
	// 枚举值：
	// PROMOTE：该推广系列为通过 TikTok 移动应用创建的内容加热推广系列。
	//
	// 仅为内容加热推广系列（campaign_system_origin 为 PROMOTE）返回以下信息：
	// advertiser_id
	// campaign_id
	// create_time
	// campaign_name
	// operation_status
	// secondary_status
	CampaignSystemOrigin enum.CampaignSystemOrigin `json:"campaign_system_origin,omitempty"`
	// AdgroupID 广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName 广告组名称
	AdgroupName string `json:"adgroup_name,omitempty"`
	// AdID 广告 ID
	AdID string `json:"ad_id,omitempty"`
	// AdName 广告名称
	AdName string `json:"ad_name,omitempty"`
	// CreateTime 广告创建时间。格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2023-01-01 00:00:01".
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ModifyTime 广告修改时间。格式为"YYYY-MM-DD HH:MM:SS"。
	// 示例："2023-01-01 00:00:01".
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
	// IdentityType 认证身份类型。枚举值： CUSTOMIZED_USER (自定义用户），AUTH_CODE (授权用户)， TT_USER (TikTok用户）， BC_AUTH_TT（添加到商务中心的TikTok用户）。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityID 认证身份ID。
	// 如果您不使用Spark Ads，传入 identity_id 和 identity_type (CUSTOMIZED_USER类型)可帮助您更好地管理广告信息。
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityAuthorizedBcID 当identity_type 为 BC_AUTH_TT时返回。
	// 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// CatalogID 您为广告选择的商品库 ID。
	// 注意：自 2024 年 6 月 30 日起，将不再为新创建的广告组层级 product_source 为 STORE 的购物广告返回 catalog_id，传入的 catalog_id 字段将被忽略。存量的 product_source 为 STORE 的购物广告若未被更新，则不受影响。若您更新存量的这类广告，且广告已绑定 catalog_id ，则绑定的 catalog_id 将清空且该 catalog_id 字段将不再返回。
	CatalogID string `json:"catalog_id,omitempty"`
	// ProductSpecificType 选择商品的维度。
	// 枚举值：
	// ALL：允许 TikTok 从所有商品中动态选择。
	// PRODUCT_SET：请选择一个商品系列。TikTok 将动态选择该商品系列中的商品。
	// CUSTOMIZED_PRODUCTS：选择自定义数量的特定商品。
	ProductSpecificType enum.ProductSpecificType `json:"product_specific_type,omitempty"`
	// ItemGroupIDs 商品的 SPU ID 列表
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
	// ProductSetID 商品系列 ID
	ProductSetID string `json:"product_set_id,omitempty"`
	// SkuIDs 商品的 SKU ID 列表
	SkuIDs []string `json:"sku_ids,omitempty"`
	// ShowcaseProducts 您想要用在广告中的橱窗商品列表。
	// 您可查看创建商品来源为橱窗的视频购物广告了解详情。
	ShowcaseProducts []ShowcaseProduct `json:"showcase_products,omitempty"`
	// AdFormat 广告样式。枚举值：SINGLE_IMAGE， SINGLE_VIDEO， LIVE_CONTENT, CAROUSEL_ADS（非视频购物类型的轮播广告），CATALOG_CAROUSEL（视频购物类型的轮播广告） 。
	AdFormat enum.AdFormat `json:"ad_format,omitempty"`
	// VerticalVideoStrategy 商品销量场景中使用的视频类型。
	// 枚举值： UNSET（未设置），SINGLE_VIDEO（单个视频），CATALOG_VIDEOS（使用视频模板的商品库视频），CATALOG_UPLOADED_VIDEOS（使用已上传视频的商品库视频）， LIVE_STREAM（直播视频）。
	VerticalVideoStrategy enum.VerticalVideoStrategy `json:"vertical_video_strategy,omitempty"`
	// DynamicFormat 动态样式。
	// 枚举值：
	// DYNAMIC_CREATIVE：智能创意。根据用户的购买意图向每位用户展示不同的视频，商品卡片和目标页面，从而最大限度地提高广告转化量。
	// UNSET：未设置。
	DynamicFormat enum.DynamicFormat `json:"dynamic_format,omitempty"`
	// VideoID 视频 ID。
	// 若广告为通过从关联的企业号、创作者授权的帖子或商务中心中创作者授权的账户提取原生视频而创建的 Spark Ads，则本字段值为 null。此时您可以通过tiktok_item_id字段获取用作 Spark Ads 的 TikTok 帖子 ID。
	// 若想了解 Spark Ads 详情，请查看创建 Spark Ads。
	VideoID string `json:"video_id,omitempty"`
	// ImageIDs 图片 ID 列表
	ImageIDs []string `json:"image_ids,omitempty"`
	// CarouselImageIndex 用于指定视频购物类型的轮播广告中要使用的附加图片的索引。
	// 您通过本字段传入的数值代表每个商品库商品的附加图片列表（additional_image_urls）中要使用的附加图片的位置。
	CarouselImageIndex int `json:"carousel_image_index,omitempty"`
	// MusicID TikTok轮播广告中使用的音乐的ID
	MusicID string `json:"music_id,omitempty"`
	// TiktokItemID 用作 Spark Ads 的 TikTok 帖子 ID，
	// 若想了解 Spark Ads 详情，请查看创建 Spark Ads。
	// 注意：本字段为以下类型的 Spark Ads 返回：
	// 通过 Spark Ads 提取创建的 Spark Ads。
	// 通过 Spark Ads 推送创建且通过广告审核的 Spark Ads。
	TiktokItemID string `json:"tiktok_item_id,omitempty"`
	// PromotionalMusicDisabled 是否关闭推广音乐开关。true： 关闭；false： 打开。默认true。如果您想为TikTok视频开启合拍或跟拍，请将该字段设置为false，即打开推广音乐。
	PromotionalMusicDisabled bool `json:"promotional_music_disabled,omitempty"`
	// ItemDuetStatus 合拍开关状态。枚举值：ENABLE： 打开合拍；DISABLE： 禁止合拍。只有在promotional_music_disabled设置为false时，该字段才有效。
	ItemDuetStatus enum.EnableDisable `json:"item_duet_status,omitempty"`
	// ItemStitchStatus 跟拍开关状态。枚举值：ENABLE： 打开跟拍；DISABLE： 禁止跟拍。只有在promotional_music_disabled设置为false时，该字段才有效
	ItemStitchStatus enum.EnableDisable `json:"item_stitch_status,omitempty"`
	// DarkPostStatus 表示该视频是否为dark post。枚举值：ON，OFF
	DarkPostStatus enum.OnOff `json:"dark_post_status,omitempty"`
	// BrandedContentDisabled 默认值为false。当branded_content_disabled为true时，dark_post_status将无法被更改
	BrandedContentDisabled bool `json:"branded_content_disabled,omitempty"`
	// ShoppingAdsVideoPackageID 商品库视频模板ID。
	// 当对应的推广系列objective_type为PRODUCT_SALES时返回该字段。
	ShoppingAdsVideoPackageID string `json:"shopping_ads_video_package_id,omitempty"`
	// AdText 广告文案，将作为广告创意的一部分展示给您的受众，向他们传达您想要推广的信息。关键词匹配类型为精确匹配
	AdText string `json:"ad_text,omitempty"`
	// AdTexts 广告文案列表
	AdTexts string `json:"ad_texts,omitempty"`
	// CallToAction 行动引导文案，见 枚举值-行动引导文案
	CallToAction enum.CallToAction `json:"call_to_action,omitempty"`
	// CallToActionID 行动引导文案素材包ID。行动引导文案素材包指的是一组自动优化的行动引导文案。如果本字段和call_to_action都传入，call_to_action会被忽略。关于动态优选CTA的更多信息，请参见智能推荐CTA > 动态优选CTA。
	CallToActionID string `json:"call_to_action_id,omitempty"`
	// CardID 创意素材包 ID
	CardID string `json:"card_id,omitempty"`
	// LandingPageURL 落地页 URL
	LandingPageURL string `json:"landing_page_url,omitempty"`
	// UtmParams URL 参数列表
	UtmParams []model.KV `json:"utm_params,omitempty"`
	// PageID 页面 ID
	PageID model.Uint64 `json:"page_id,omitempty"`
	// CppURL 自定产品页面 URL。
	// 自定产品页面是一种定制化的产品页面，支持自定义截图、推广文字和应用预览，可用于突出介绍所推广应用的特定内容或功能，或针对特定受众群体进行推广。
	CppURL string `json:"cpp_url,omitempty"`
	// TiktokPageCategory 您希望推广的 TikTok 页面的类型。
	// 枚举值：
	// PROFILE_PAGE：TikTok 账号主页。
	// OTHER_TIKTOK_PAGE：其他 TikTok 页面，包括播放列表页面、话题标签页面、音乐页面和商业化特效页面。
	// TIKTOK_INSTANT_PAGE：自定义 TikTok 即时体验页面。
	TiktokPageCategory enum.TiktokPageCategory `json:"tiktok_page_category,omitempty"`
	// PhoneRegionCode 仅当广告组层级 promotion_type 设置为 LEAD_GEN_CLICK_TO_CALL 时返回。
	// 受众可点击广告拨打的电话号码的地区代码。
	// 若想了解如何创建推广对象类型为电话通话的线索广告，请查看这里。
	// 注意：创建推广对象类型为电话通话的线索广告目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	PhoneRegionCode string `json:"phone_region_code,omitempty"`
	// PhoneRegionCallingCode 仅当广告组层级 promotion_type 设置为 LEAD_GEN_CLICK_TO_CALL 时返回。
	// 受众可点击广告拨打的电话号码的区号。
	// 注意：创建推广对象类型为电话通话的线索广告目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表
	PhoneRegionCallingCode string `json:"phone_region_calling_code,omitempty"`
	// PhoneNumber 仅当广告组层级 promotion_type 设置为 LEAD_GEN_CLICK_TO_CALL 时返回。
	// 受众可点击广告拨打的电话号码。
	// 注意：创建推广对象类型为电话通话的线索广告目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	PhoneNumber string `json:"phone_number,omitempty"`
	// Deeplink 深度链接，即已安装应用的用户点击链接后跳转至的特定页面
	Deeplink string `json:"deeplink,omitempty"`
	// DeeplinkType 深度链接类型。
	// 枚举值：
	// NORMAL：非延迟深度链接。
	// DEFERRED_DEEPLINK：延迟深度链接
	DeeplinkType enum.DeeplinkType `json:"deeplink_type,omitempty"`
	// DeeplinkFormatType 深度链接的格式类型。
	// 枚举值：
	// UNIVERSAL_OR_APP_LINK：通用/应用链接，即以http://或 https:// 开头的 iOS 通用链接或 Android 应用链接。
	// SCHEME_LINK：URL 架构，即格式为scheme://resource 的自定义网址架构。例如，WhatsApp 的自定义网址架构格式应为whatsapp://resource。
	// NONE：非深度链接。
	DeeplinkFormatType enum.DeeplinkFormatType `json:"deeplink_format_type,omitempty"`
	// ShoppingAdsDeeplinkType 在购物广告中要使用的深度链接的来源。
	// 枚举值：
	// NONE：不设置深度链接。
	// CUSTOM：自定义深度链接。在广告中使用你提供的深度链接。
	// SHOPPING_ADS：商品库深度链接。使用商品库中为各商品添加的深度链接。
	ShoppingAdsDeeplinkType enum.ShoppingAdsDeeplinkType `json:"shopping_ads_deeplink_type,omitempty"`
	// DeeplinkUtmParams 深度链接 URL 参数列表。
	// 注意：在非视频购物广告场景使用 utm_params 目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	DeeplinkUtmParams []model.KV `json:"deeplink_utm_params,omitempty"`
	// ShoppingAdsFallbackType 在购物广告再营销场景下，深度链接唤起失败后的返回落地页类型。
	// 枚举值：
	// DEFAULT：未设置。
	// CUSTOM: 自定义链接。您可以在广告组中提供一个网页地址, 并回退到这个地址。
	// SHOPPING_ADS：商品库链接。 回退到目录中为各商品添加的网页地址。
	ShoppingAdsFallbackType enum.ShoppingAdsFallbackType `json:"shopping_ads_fallback_type,omitempty"`
	// FallbackType 失败路径。如果用户没有安装过应用，您可以选择让用户回退去安装应用，或者去您想推广的网页。
	// 枚举值：
	// APP_INSTALL：广告组层级通过 app_id 指定的推广应用对应的 App Store 或 Google Play 页面。
	// WEBSITE：通过 landing_page_url 指定的推广网页。
	// UNSET：未设置。
	FallbackType enum.FallbackType `json:"fallback_type,omitempty"`
	// DynamicDestination 动态目标页面策略。
	// 枚举值：
	// DLP：动态落地页。基于用户的主页、行为和意向，向每位用户展示不同的目标页面，包括自定义网站或即时体验商品页面等，从而最大限度地提高广告转化量。
	// UNSET：未设置。
	DynamicDestination enum.DynamicDestination `json:"dynamic_destination,omitempty"`
	// AutoMessageID TikTok 私信广告中使用的自动消息的 ID。
	// 目前唯一支持的自动消息类型为欢迎消息
	AutoMessageID string `json:"auto_message_id,omitempty"`
	// AigcDisclosureType 是否启用 AI 生成内容自主声明开关，以表明广告中包含 AI 生成内容。启用该开关后，当用户查看完整广告时，您的广告将带有“广告主标记为 AI 生成”的标签。
	// 枚举值：
	// SELF_DISCLOSURE：启用开关，声明广告包含 AI 生成内容。
	// NOT_DECLARED：不声明广告包含 AI 生成内容。
	AigcDisclosureType enum.AigcDisclosureType `json:"aigc_disclosure_type,omitempty"`
	// DisclaimerType 广告中的免责声明类型。枚举值：TEXT_LINK（文字链免责声明），TEXT_ONLY（纯文本免责声明）。
	// 注意：免责声明目前对注册地非美国或加拿大且使用竞价推广目标的广告主，以及使用覆盖和频次推广目标的所有广告主为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	DisclaimerType enum.DisclaimerType `json:"disclaimer_type,omitempty"`
	// DisclaimerText 广告中的纯文本免责声明
	DisclaimerText *DisclaimerText `json:"disclaimer_text,omitempty"`
	// DisclaimerClickableText 广告中的文字链免责声明
	DisclaimerClickableTexts []DisclaimerClickableText `json:"disclaimer_clickable_texts,omitempty"`
	// TrackingPixelID 正在监测的 Pixel ID。你可以使用此字段追踪发生在外部网站的事件
	TrackingPixelID model.Uint64 `json:"tracking_pixel_id,omitempty"`
	// TrackingAppID 正在监测的应用的 ID
	TrackingAppID string `json:"tracking_app_id,omitempty"`
	// TrackingOfflineEventSetIDs 正在监测的线下事件组 ID 列表。
	// 点击这里了解如何创建和管理线下事件组。
	TrackingOfflineEventSetIDs []string `json:"tracking_offline_event_set_ids,omitempty"`
	// TrackingMessageEventSetID 即时通讯广告中监测的消息事件集的 ID
	TrackingMessageEventSetID string `json:"tracking_message_event_set_id,omitempty"`
	// VastMoatEnabled 是否开启Moat（第三方视频可见性监测）。TikTok与Moat合作，为您在TikTok for Business上购买的品牌竞价、覆盖和频次信息流广告提供可见性监测。
	// 注意：若您想开启Moat第三方可见性监测，相较于本字段，更推荐使用viewability_postbid_partner，将值设为 MOAT 。
	VastMoatEnabled bool `json:"vast_moat_enabled,omitempty"`
	// ViewabilityPostbidPartner 出价后第三方可见性监测合作伙伴。枚举值：UNSET，MOAT，DOUBLE_VERIFY，IAS，ZEFR。
	// 注意：出价后第三方品牌安全监测目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	ViewabilityPostbidPartner enum.ViewabilityPostbidPartner `json:"viewability_postbid_partner,omitempty"`
	// ViewabilityVastURL 出价后第三方监测合作伙伴用于监测可见性的封装的 VAST 网址。当brand_safety_postbid_partner非IAS时本字段为空值。
	// 注意：出价后第三方品牌安全监测目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	ViewabilityVastURL string `json:"viewability_vast_url,omitempty"`
	// BrandSafetyPostbidPartner 出价后第三方品牌安全监测合作伙伴。枚举值： UNSET，DOUBLE_VERIFY，IAS，ZEFR。
	// 注意：出价后第三方品牌安全监测目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	BrandSafetyPostbidPartner enum.BrandSafetyPostbidPartner `json:"brand_safety_postbid_partner,omitempty"`
	// BrandSafetyVastURL 出价后第三方监测合作伙伴用于监测品牌安全的封装的 VAST 网址。当brand_safety_postbid_partner非IAS时本字段为空值。
	// 注意：出价后第三方品牌安全监测目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	BrandSafetyVastURL string `json:"brand_safety_vast_url,omitempty"`
	// ImpressionTrackingURL 默认展示监测 URL
	ImpressionTrackingURL string `json:"impression_tracking_url,omitempty"`
	// ClickTrackingURL 点击监测 URL
	ClickTrackingURL string `json:"click_tracking_url,omitempty"`
	// PlayableURL 试玩广告 URL
	PlayableURL string `json:"playable_url,omitempty"`
	// VideoViewTrackingURL 视频播放监测URL。
	VideoViewTrackingURL string `json:"video_view_tracking_url,omitempty"`
	// OperationStatus 广告的操作状态。
	// ENABLE：广告处于启用（“开”）状态。
	// DISABLE：广告处于未启用（“关”）状态。
	// FROZEN：广告处于冻结状态，无法再次启用。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// SecondaryStatus 广告状态（二级状态）。枚举值详见枚举值-二级状态-广告状态。
	// 注意：沙箱环境下本字段不返回，因为广告未实际投放。
	SecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// CreativeType 直播购物广告、商品购物广告或应用预注册场景中的创意素材类型。枚举值：
	// SHOP_PDP：短视频带货广告，包含一个商品链接。
	// SHOP_PLP: 短视频带货广告，包含6-50个商品链接。
	// SHORT_VIDEO_LIVE: 短视频引流至直播间广告。
	// DIRECT_LIVE: 直播间带货。每个推广系列只能有一个直播间。
	// PSA：商品视觉元素将作为广告创意使用。
	// CUSTOM_INSTANT_PAGE：自定义 TikTok 即时体验页面。
	// 仅当推广系列层级 app_promotion_type设置为APP_PREREGISTRATION，且广告组层级promotion_website_type设置为TIKTOK_NATIVE_PAGE时返回此值。
	CreativeType enum.CreativeType `json:"creative_type,omitempty"`
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
	// DisplayName 落地页或纯曝光广告的显示名称
	DisplayName string `json:"display_name,omitempty"`
	// AvatarIconWebURL 广告头像图片 ID
	AvatarIconWebURL string `json:"avatar_icon_web_url,omitempty"`
	// ProfileImageURL 头像 URL
	ProfileImageURL string `json:"profile_image_url,omitempty"`
	// CreativeAuthorized 是否允许在我们的 TikTok For Business 创意中心展示您的广告。仅对非美国广告主有效。默认值：false。
	// 注意：creative_authorized仅可用于视频广告。
	CreativeAuthorized bool `json:"creative_authorized,omitempty"`
	// IsAco 是否为程序化广告或智能创意广告。若为程序化广告或智能创意广告，则为true
	IsAco bool `json:"is_aco,omitempty"`
	// IsNewStructure 该广告是否为新结构（广告与对应推广系列及广告组结构相同）
	IsNewStructure bool `json:"is_new_structure,omitempty"`
	// OptimizationEvent 广告组转化事件。参阅转化事件，获取详细信息
	OptimizationEvent enum.OptimizationEvent `json:"optimization_event,omitempty"`
}

// ShowcaseProduct 您想要用在广告中的橱窗商品列表。
type ShowcaseProduct struct {
	// ItemGroupID 商品的 SPU ID
	ItemGroupID string `json:"item_group_id,omitempty"`
	// StoreID 商品（item_group_id）所属的店铺 ID。
	// 请注意本字段仅支持 TikTok Shop 店铺。
	StoreID string `json:"store_id,omitempty"`
	// CatalogID 商品（item_group_id）所属的商品库 ID。
	// 注意：自 2024 年 6 月 30 日起，将不再为新创建的广告组层级 product_source 为 STORE 的购物广告返回本字段，传入的 showcase_products 对象数组中的 catalog_id 字段将被忽略。存量的 product_source 为 STORE 的购物广告若未被更新，则不受影响。若您更新存量的这类广告，且广告已绑定 showcase_products 对象数组中的 catalog_id，则绑定的catalog_id 将清空且该 catalog_id 字段将不再返回。
	CatalogID string `json:"catalog_id,omitempty"`
}

// DisclaimerText 广告中的纯文本免责声明
type DisclaimerText struct {
	// Text 免责声明文案。
	Text string `json:"text,omitempty"`
}

// DisclaimerClickableText 广告中的文字链免责声明
type DisclaimerClickableText struct {
	// Text 免责声明文案
	Text string `json:"text,omitempty"`
	// URL 文字链免责声明的URL 。用户点击对应文字时，将打开您提供的对应URL，查看免责声明的详细内容
	URL string `json:"url,omitempty"`
}
