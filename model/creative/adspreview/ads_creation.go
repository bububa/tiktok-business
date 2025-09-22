package adspreview

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/ad"
	"github.com/bububa/tiktok-business/util"
)

// AdsCreationPreviewRequest 预览计划创建的广告 API Request
type AdsCreationPreviewRequest struct {
	CreateRequest
	// ObjectiveType 推广目标。
	// 枚举值详见推广目标。
	// 注意：目前仅支持下列推广目标: REACH，TRAFFIC, VIDEO_VIEWS，ENGAGEMENT，APP_PROMOTION，LEAD_GENERATION，WEB_CONVERSIONS ，PRODUCT_SALES，和RF_REACH。
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// IsSmartPerformanceCampaign 推广系列是否为 Smart+ 推广系列。
	// 默认值：false。
	// 注意：在以下任意情况下，本字段均只支持设置为false ：
	//
	// ad_fomat 设置为 CAROUSEL_ADS 或 CATALOG_CAROUSEL 时。
	// placement 设置为 PLACEMENT_PANGLE或 ALL 时。
	// objective_type 设置为 TRAFFIC ， REACH ， VIDEO_VIEWS 或 PRODUCT_SALES 。
	IsSmartPerformanceCampaign bool `json:"is_smart_performance_campaign,omitempty"`
	// Placement 广告版位。
	// ad_format设置为CAROUSEL_ADS或CATALOG_CAROUSEL时，将本字段设置为PLACEMENT_TIKTOK，或不传入本字段。
	// 枚举值：
	// PLACEMENT_TIKTOK：仅在objective_type 为REACH，TRAFFIC, VIDEO_VIEWS，ENGAGEMENT，APP_PROMOTION，LEAD_GENERATION， WEB_CONVERSIONS 或 PRODUCT_SALES 时支持。
	// PLACEMENT_PANGLE：仅在objective_type 为TRAFFIC，APP_PROMOTION，LEAD_GENERATION（仅当 promotion_type为LEAD_GENERATION时）或 WEB_CONVERSIONS 时支持。
	// ALL：仅在objective_type 为TRAFFIC，APP_PROMOTION，LEAD_GENERATION（仅当 promotion_type为LEAD_GENERATION时）或 WEB_CONVERSIONS 时支持。
	Placement enum.Placement `json:"placement,omitempty"`
	// PreviewFormat 不同placement的预览格式有所不同：
	// 当placement为PLACEMENT_TIKTOK时，允许的预览格式为：
	// IN_FEED：信息流。广告将投放在“推荐”动态中，可能投放在“个人资料页”和“关注”推荐内容中。
	// SEARCH_RESULTS：搜索结果页面。
	// SEARCH_FEED：搜索推荐内容。
	// TIKTOK_LITE：TikTok Lite，占用内容更小、视频加载速度更快的精简版 TikTok 应用。TikTok Lite 子版位仅支持定向日本或韩国时使用。
	// PRODUCT_SEARCH: 搜索结果页面（Shopping Center）。在“商城”选项卡的搜索结果页面以及 TikTok 应用的常规搜索结果页面中展示您的商品。了解关于可投放店铺广告的“商城”选项卡广告位的详情。
	// PRODUCT_SHOP_CENTER: “为你推荐”板块（Shopping Center）。在“为你推荐”板块，将您的商品展示给可能会购买的用户。
	// 默认值：IN_FEED。
	//
	// 注意：
	//
	// SEARCH_FEED 和TIKTOK_LITE 仅在满足以下条件时生效：
	// objective_type 为 REACH、VIDEO_VIEWS 或 ENGAGEMENT。
	// placement 为 PLACEMENT_TIKTOK。
	// PRODUCT_SEARCH 和 PRODUCT_SHOP_CENTER 枚举值仅对视频购物广告和商品购物广告可用。
	// 当placement为PLACEMENT_PANGLE时，允许的预览格式为INTERSTITIAL, REWARDED, APP_OPEN， NORMAL_BANNER, VIDEO_THUMBNAIL_BANNER, SMALL_VIDEO_BANNER, ICON_ONLY_BANNER, NORMAL_NATIVE, VIDEO_THUMBNAIL_NATIVE。默认值：INTERSTITIAL。
	PreviewFormat enum.AdPreviewFormat `json:"preview_format,omitempty"`
	// ShoppingAdsType 当 objective_type 为 PRODUCT_SALES 且未传入 catalog_id 时必填。
	// 购物广告类型。
	// 枚举值：
	// VIDEO：视频购物广告。
	// LIVE：直播购物广告。
	// PRODUCT_SHOPPING_ADS: 商品购物广告。
	//
	// ad_format设置为CATALOG_CAROUSEL时，将本字段设置为VIDEO，或不传入本字段。
	ShoppingAdsType enum.ShoppingAdsType `json:"shopping_ads_type,omitempty"`
	// ProductSource 当为视频购物广告和商品购物广告时必填。
	// 您希望从中获取推广商品的产品来源。
	// 枚举值：STORE (TikTok Shop),SHOWCASE(TikTok Showcase)。
	// 对于视频购物广告，您可以将此字段设为 STORE 或 SHOWCASE。
	// 对于商品购物广告，将此字段设为 STORE。
	ProductSource enum.ProductSource `json:"product_source,omitempty"`
	// StoreID 在以下任一场景必填：
	// shopping_ads_type 为 VIDEO 且 product_source 为 STORE。
	// shopping_ads_type 为 PRODUCT_SHOPPING_ADS 且 product_source 为 STORE。
	// 当 shopping_ads_type 为 LIVE 时，此字段选填。
	//
	// TikTok Shop ID。
	//
	// 若想获取 TikTok Shop ID，可调用 /bc/asset/get/ 并将 asset_type 设置为 TIKTOK_SHOP。返回的 asset_id 将是 TikTok Shop ID。
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizeBcID 当传入 store_id 时必填。
	// 有权管理该 TikTok Shop (store_id) 的商务中心的 ID。
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// ShowcaseProducts 当 shopping_ads_type 为 VIDEO 且 product_source 为 SHOWCASE 时必填。
	//
	// 了解详情请查看创建商品来源为橱窗的视频购物广告。
	//
	// 您想要用在广告中的橱窗商品列表。
	//
	// 最大数量：20。
	//
	// 若想获取可用的橱窗商品，可调用 /showcase/product/get/。
	ShowcaseProducts []ad.ShowcaseProduct `json:"showcase_products,omitempty"`
	// PromotionType 推广对象类型。
	// 枚举值：
	// APP_ANDROID：Android应用。
	// APP_IOS：iOS应用。
	// WEBSITE：落地页。
	// LEAD_GENERATION: 即时表单（线索表单）或落地页。
	// LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE: TikTok 私信。
	// LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE: 即时通讯应用。
	// LEAD_GEN_CLICK_TO_CALL: 电话通话。
	//
	// 注意：以下枚举值仅支持在 objective_type 为 LEAD_GENERATION 时使用：
	//
	// LEAD_GENERATION
	// LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE
	// LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE
	// LEAD_GEN_CLICK_TO_CALL
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// IdentityID 认证身份 ID。
	// 如果您不使用 Spark Ads，传入 identity_id 和 identity_type（CUSTOMIZED_USER类型）可帮助您更好地管理广告信息。
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。
	// 枚举值：CUSTOMIZED_USER，AUTH_CODE ，TT_USER，BC_AUTH_TT。
	// 对于自定义认证身份，请使用CUSTOMIZED_USER，在 Spark Ads 中使用AUTH_CODE ，TT_USER，BC_AUTH_TT。
	// 参阅认证身份获取详细信息。
	//
	// ad_format设置为CATALOG_CAROUSEL时，仅可将本字段设置为CUSTOMIZED_USER。
	// objective_type 为 LEAD_GENERATION 时：
	// 若 promotion_type 为 LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE ，必须使用 Spark Ads 且 identity_type 需设置为 TT_USER 或 BC_AUTH_TT。
	// 若 promotion_type 非 LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE ，支持 Spark Ads 和非 Spark Ads，identity_type 可设置为任意枚举值。
	// objective_type 为 ENGAGEMENT 时，必须使用 Spark Ads 且identity_type 需设置为AUTH_CODE，TT_USER 或 BC_AUTH_TT。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizeBcID 当identity_type 为 BC_AUTH_TT时必填。
	// 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的 ID。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// AdFormat 广告样式。
	// 对于单视频广告，将此字段设为 SINGLE_VIDEO。
	// 对于单图片广告，将此字段设为 SINGLE_IMAGE。
	// 对于 TikTok 标准轮播广告，将此字段设为 CAROUSEL_ADS。
	// 对于 TikTok VSA 轮播广告，将此字段设为 CATALOG_CAROUSEL。
	// 点击这里查看创建轮播广告的步骤。
	// 对于直播购物广告，将此字段设为 LIVE_CONTENT。
	AdFormat enum.AdFormat `json:"ad_format,omitempty"`
	// VideoID 视频素材 ID。
	// ad_format为SINGLE_VIDEO时必填。
	// ad_format为SINGLE_IMAGE，CAROUSEL_ADS或CATALOG_CAROUSEL时不支持本字段。
	// 若想创建非 Spark Ads 视频广告或通过 Spark Ads 推送创建 Spark Ads 视频广告 ，需使用video_id。您可以使用/file/video/ad/search/接口查询视频素材 ID。
	// 若想通过 Spark Ads 拉取创建 Spark Ads 视频广告，需使用tiktok_item_id指定TikTok 帖子 ID。
	VideoID string `json:"video_id,omitempty"`
	// ImageIDs 仅在以下任意场景必填：
	// ad_format 为SINGLE_IMAGE 。
	// ad_format 为 CAROUSEL_ADS 且您想创建非 Spark Ads 标准轮播广告或通过 Spark Ads 推送创建 Spark Ads 标准轮播广告。
	//
	// 图片 ID 列表（图片素材或视频封面）。
	//
	// 当 ad_format 设置为 CAROUSEL_ADS 时，本字段的 ID 数量范围为[1, 35]，且本字段指定的图片ID的顺序即为轮播广告中对应图片的展示顺序。
	// 请查看创建轮播广告，了解可在 TikTok 轮播广告中使用的图片的要求。
	// 当 ad_format 设置为 SINGLE_IMAGE 时，仅可向本字段传入一个图片 ID。
	// 若想了解单图片广告中的图片要求，请查看创建单图片广告。
	ImageIDs []string `json:"image_ids,omitempty"`
	// MusicID 仅在以下任意场景必填：
	// ad_format 为 CATALOG_CAROUSEL。
	// ad_format 为 CAROUSEL_ADS 且您创建非 Spark Ads 标准轮播广告或通过 Spark Ads 推送创建 Spark Ads 标准轮播广告。
	//
	// 想要在 TikTok 轮播广告中使用的音乐的 ID。
	MusicID string `json:"music_id,omitempty"`
	// TiktokItemID identity_type 设置为AUTH_CODE，TT_USER 或 BC_AUTH_TT 且您想要通过 Spark Ads 提取创建 Spark Ads 时必填。
	// identity_type 设置为 CUSTOMIZED_USER 时不支持本字段。
	//
	// 用作 Spark Ads 的 TikTok 帖子 ID。
	//
	// 若想了解 Spark Ads 详情，请查看创建 Spark Ads。
	//
	// ID 可以从/tt_video/info/或/identity/video/get/接口的返回值中的item_id字段值获取。
	//
	// 注意：使用此功能，代表您声明拥有该视频中音乐的版权，可用于营销使用。
	TiktokItemID string `json:"tiktok_item_id,omitempty"`
	// CarouselImageIndex 仅当ad_format设置为CATALOG_CAROUSEL，且所指定的商品已设置对应的附加图片列表时生效。
	// 用于指定视频购物类型的轮播广告中要使用的附加图片的索引。若未传入本字段，将在视频购物类型的轮播广告中使用商品库商品的主图。
	// 取值范围：[0, 9]。
	// 您通过本字段传入的数值代表每个商品库商品的附加图片列表（additional_image_urls）中要使用的附加图片的位置。例如，若carousel_image_index 设置为 1，代表您将使用每个additional_image_urls的值中的第二个图片作为轮播广告中使用的图片。您可通过/catalog/product/get/返回的additional_image_urls字段获取商品对应的附加图片列表。
	// 若您指定的数值超过了商品对应的附加图片列表中附加图片URL的数量，将忽略本字段，使用商品的主图（image_url）。您可通过/catalog/product/get/返回的image_url 字段获取商品对应的主图。
	CarouselImageIndex int `json:"carousel_image_index,omitempty"`
	// AdText 广告文案，将作为广告创意的一部分展示给你的受众，向他们传达你想要推广的信息。不知道怎么写？试试智能文案。
	// 对于非Spark Ads广告，本字段必填。
	// 广告文案长度允许值为 1-100 字符，不能包含emoji。
	// 中文或日文每个字占用2个字符，每个英语字符占用1个字符。
	AdText string `json:"ad_text,omitempty"`
	// CallToAction 行动引导文案，引导用户在看到你的广告后完成你期望的行为。对于TikTok广告，本字段或者 call_to_action_id 必须传入一个，如果同时传入，此字段将被忽略。枚举值见枚举值 - 行动引导文案。直播购物事件下，此字段的值必须为 WATCH_LIVE。
	// 注意
	//
	// 如果objective_type为REACH或VIDEO_VIEW，且您已传入landing_page_url或page_id，则必须传入call_to_action或 call_to_action_id其中一个字段。
	// 如果objective_type为APP_PROMOTION， WEB_CONVERSIONS，或TRAFFIC，则必须传入call_to_action或 call_to_action_id其中一个字段。
	// 如果 objective_type 为 LEAD_GENERATION ：
	// promotion_type 为 LEAD_GENERATION 时，需传入 call_to_action 和 call_to_action_id 其中之一。
	// promotion_type 为LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE时，必须传入call_to_action，且必须设置为 SEND_MESSAGE。
	// promotion_type 为 LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE 时，必须传入call_to_action，且不允许设置为CALL_NOW。
	// promotion_type 为 LEAD_GEN_CLICK_TO_CALL时，必须传入call_to_action。
	// 如果objective_type为PRODUCT_SALES，则必须传入call_to_action。
	CallToAction enum.CallToAction `json:"call_to_action,omitempty"`
	// CallToActionID 行动引导文案素材包ID。行动引导文案素材包指的是一组自动优化的行动引导文案。如果本字段和 call_to_action 都传入，call_to_action 会被忽略。关于自动优化行动引导文案的更多信息，请参见智能推荐CTA > 动态优选CTA。
	// 注意
	//
	// 如果objective_type为REACH或VIDEO_VIEW，且您已传入landing_page_url或page_id，则必须传入call_to_action或 call_to_action_id其中一个字段。
	// 如果objective_type为APP_PROMOTION， WEB_CONVERSIONS，或TRAFFIC，则必须传入call_to_action或 call_to_action_id其中一个字段。
	// 如果ad_format为CAROUSEL_ADS或CATALOG_CAROUSEL，则不支持传入call_to_action_id。
	// 如果objective_type 为 LEAD_GENERATION：
	// promotion_type 为 LEAD_GENERATION 时，需传入 call_to_action 和call_to_action_id 其中之一。
	// promotion_type 为 LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE，LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE 或 LEAD_GEN_CLICK_TO_CALL 时，不支持 call_to_action_id 。
	CallToActionID string `json:"call_to_action_id,omitempty"`
	// CardID 创意素材包 ID。
	// 传入以下类型的创意素材包的 ID：
	// 展示卡片（又称图片卡片）。
	// 网站信息卡片。
	// 商品卡片。
	// 商品磁贴。
	// 倒计时贴纸。
	// 礼品码贴纸。
	// 交互手势。
	// 点赞彩蛋。
	//
	// 若想创建素材包，可使用 /creative/portfolio/create/。
	//
	// 若想了解如何获取展示卡片、网站信息卡片、商品卡片或商品磁贴素材包的 ID，请参考卡片。
	// 若想了解如何获取倒计时贴纸或礼品码贴纸素材包的 ID，请参考贴纸。
	// 若想了解如何获取交互手势或点赞彩蛋素材包的 ID，请参考高级创新互动样式。
	//
	// 注意：您可以在以下场景使用本字段。
	//
	// 当objective_type= REACH，promotion_type= WEBSITE，placements= PLACEMENT_TIKTOK，且ad_format = SINGLE_VIDEO。
	// 当objective_type= TRAFFIC，promotion_type= WEBSITE，placements= PLACEMENT_TIKTOK，且ad_format = SINGLE_VIDEO。
	// 当objective_type= VIDEO_VIEWS，placements= PLACEMENT_TIKTOK，且ad_format = SINGLE_VIDEO。
	// 当 objective_type = ENGAGEMENT ，仅支持通过本字段指定倒计时贴纸素材包。
	// 当objective_type= APP_PROMOTION，promotion_type= WEBSITE，placements= PLACEMENT_TIKTOK，且ad_format = SINGLE_VIDEO。
	// 当 objective_type = LEAD_GENERATION ，仅支持通过本字段指定展示卡片素材包。
	// 当objective_type=WEB_CONVERSIONS，placements= PLACEMENT_TIKTOK，且ad_format = SINGLE_VIDEO。
	// 当objective_type=PRODUCT_SALES，placements= PLACEMENT_TIKTOK，且 ad_format = SINGLE_VIDEO。
	CardID string `json:"card_id,omitempty"`
	// LandingPageURL 落地页URL。
	// 注意
	//
	// 当objective_type为APP_PROMOTION，WEB_CONVERSIONS或TRAFFIC，且promotion_type为WEBSITE时，page_id或landing_page_url必须传入其中一个。
	// 如果objective_type为APP_PROMOTION或 TRAFFIC，且promotion_type为APP_ANDROID/APP_IOS，page_id或landing_page_url都不需要传入。
	LandingPageURL string `json:"landing_page_url,omitempty"`
	// PageID 页面 ID。
	// 注意
	//
	// 当objective_type为APP_PROMOTION，WEB_CONVERSIONS或TRAFFIC，且promotion_type为WEBSITE时，page_id或landing_page_url必须传入其中一个。
	// 如果objective_type为APP_PROMOTION或 TRAFFIC，且promotion_type为APP_ANDROID/APP_IOS，page_id或landing_page_url都不需要传入。
	// 如果ad_format为CATALOG_CAROUSEL，不支持本字段。
	PageID model.Uint64 `json:"page_id,omitempty"`
	// CatalogID 商品库 ID。广告中的商品来自该商品库。
	// ad_format设置为CATALOG_CAROUSEL时，本字段必填。
	// 注意：您需对商品库有广告推广或商品库管理权限。若想查看您对某个商品库的权限，可使用/bc/asset/get/，返回的 catalog_role 需 为 ADMIN 或 AD_PROMOTE。
	CatalogID string `json:"catalog_id,omitempty"`
	// ProductSpecificType 选择商品的维度。
	// ad_format设置为CATALOG_CAROUSEL时，本字段必填。轮播图片将从所指定的商品库商品范围中动态选取，且所指定的商品库商品范围需包含至少两个商品。
	// 枚举值：
	// ALL：允许TikTok动态选择商品库中的所有商品。
	// PRODUCT_SET：请选择一个商品系列。TikTok将动态选择该商品系列中的商品。
	// CUSTOMIZED_PRODUCTS：最多可从您的商品库中选择20个商品。
	// 注意：
	//
	// 若将本字段设置为ALL，则无需传入sku_ids，item_group_ids 或 product_set_id。
	// 若本字段设置为PRODUCT_SET，则需要传入item_group_ids 或 product_set_id。
	// 若本字段设置为CUSTOMIZED_PRODUCTS，sku_ids必填。
	//
	// 注意:
	//
	// 当 shopping_ads_type 为 PRODUCT_SHOPPING_ADS 且 product_source 为 STORE 时，您需将此字段设为 ALL 或 CUSTOMIZED_PRODUCTS。
	// 若此字段设为 ALL，无需传入 sku_ids，item_group_ids 和 product_set_id。
	// 若此字段设为 CUSTOMIZED_PRODUCTS，item_group_ids 必填。
	// 当 shopping_ads_type 为 VIDEO 且 product_source 为 STORE 时，需通过 item_group_ids 指定商品。
	// 当 shopping_ads_type 为 VIDEO 且 product_source 为 SHOWCASE 时，需通过 showcase_products 指定商品。
	// 当 shopping_ads_type 为 LIVE 时，无需指定商品。
	ProductSpecificType enum.ProductSpecificType `json:"product_specific_type,omitempty"`
	// ItemGroupIDs 若指定了 catalog_id且product_specific_type设置为PRODUCT_SET，需传入item_group_ids或product_set_id。
	// 在以下任一场景中必填：
	// shopping_ads_type 为 VIDEO 且 product_source 为 STORE。
	// shopping_ads_type 为 PRODUCT_SHOPPING_ADS，product_source 为 STORE，且 product_specific_type 为 CUSTOMIZED_PRODUCTS。
	//
	// 商品的 SPU（标准化产品单元） ID列表。
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
	// ProductSetID 若指定了 catalog_id且product_specific_type设置为PRODUCT_SET，需传入item_group_ids或product_set_id。
	// 商品系列 ID。
	ProductSetID string `json:"prodduct_set_id,omitempty"`
	// SkuIDs 当指定了 catalog_id且product_specific_type设置为CUSTOMIZED_PRODUCTS时必填。
	// SKU（库存量单位）ID列表。
	SkuIDs []string `json:"sku_ids,omitempty"`
	// CatalogAuthorizedBcID 拥有商品库权限的商务中心ID。预览商品库商品时，请使用该字段。
	CatalogAuthorizedBcID string `json:"catalog_authorized_bc_id,omitempty"`
	// DynamicFormat ad_format 设置为 CAROUSEL_ADS或CATALOG_CAROUSEL 时不支持本字段。
	// 动态样式。
	// 枚举值：
	// UNSET：未设置。
	// DYNAMIC_CREATIVE：智能创意。根据用户的购买意图向每位用户展示不同的视频，商品卡片和目标页面，从而最大限度地提高广告转化量。
	DynamicFormat enum.DynamicFormat `json:"dynamic_format,omitempty"`
	// VerticalVideoStrategy ad_format 设置为 CAROUSEL_ADS或CATALOG_CAROUSEL 时不支持本字段。
	// 视频策略。
	// 枚举值：
	// SINGLE_VIDEO：利用视频创意推广您的商品。
	// CATALOG_VIDEOS：基于商品库中的信息生成动态视频。
	// LIVE_STREAM: 使用直播推广您的商品。本枚举值仅适用于广告样式为实时直播的直播购物广告。
	// UNSET：未设置。
	// 若dynamic_format为DYNAMIC_CREATIVE，该字段必须设置为UNSET。
	VerticalVideoStrategy enum.VerticalVideoStrategy `json:"vertical_video_strategy,omitempty"`
	// ShoppingAdsVideoPackageID 商品视频包 ID。
	// 使用/catalog/video/get/接口获取视频包 ID。
	ShoppingAdsVideoPackageID string `json:"shopping_ads_video_package_id,omitempty"`
	// DynamicDestination ad_format 设置为 CAROUSEL_ADS或CATALOG_CAROUSEL 时不支持本字段。
	// 动态目标页面策略。
	// 枚举值：
	// DLP：动态落地页。基于用户的主页、行为和意向，向每位用户展示不同的目标页面，包括自定义网站或即时体验商品页面等，从而最大限度地提高广告转化量。
	// UNSET：未设置。
	DynamicDestination enum.DynamicDestination `json:"dynamic_destination,omitempty"`
	// InstantProductPageUsed ad_format 设置为 CAROUSEL_ADS或CATALOG_CAROUSEL 时不支持本字段。
	// 是否使用TikTok即时体验商品页面（自动生成的TikTok即时体验页面，展示多种商品，提供原生落地页体验）。
	// 若dynamic_destination为DLP, 则不需要该参数。
	// 若dynamic_destination为UNSET, 您可以将本字段设置为true，且无需同时传入page_id。
	InstantProductPageUsed *bool `json:"instantt_product_page_used,omitempty"`
}

// Encode implements PostRequest interface
func (r *AdsCreationPreviewRequest) Encode() []byte {
	r.PreviewType = enum.AdPreviewType_ADS_CREATION
	return util.JSONMarshal(r)
}
