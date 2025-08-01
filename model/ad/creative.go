package ad

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Creative 广告创意
type Creative struct {
	// AdID 广告 ID
	AdID string `json:"ad_id,omitempty"`
	// AdName 广告名称。
	// 若想使广告名称自动生成，可传入 "" （空字符串）。自动生成广告名格式：广告 ID （ad_id）。
	// 长度限制：512字符且不能包含表情符号。
	// 注意：中文或日文每个字占用2个字符，每个英语字符占用1个字符。
	AdName *string `json:"ad_name,omitempty"`
	// IdentityType 认证身份类型。
	// 枚举值： CUSTOMIZED_USER (自定义用户），AUTH_CODE (授权用户)， TT_USER (TikTok 企业号用户）， BC_AUTH_TT（添加到商务中心的TikTok用户）。
	// 关于认证身份的更多信息，请访问认证身份。
	// 注意：类型为 AUTH_CODE 、 TT_USER 或 BC_AUTH_TT 的认证身份仅可用于创建最多10,000个广告。若达到此限制，您可以通过删除使用该身份的广告来释放与该身份关联的广告配额。若想查询此类身份的剩余广告配额，您可以调用/asset/bind/quota/并查看返回的 available_quota 字段值。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityID 认证身份ID。
	// 如果您不使用Spark Ads，传入 identity_id 和 identity_type (CUSTOMIZED_USER类型)可帮助您更好地管理广告信息。
	// 注意：类型为 AUTH_CODE 、 TT_USER 或 BC_AUTH_TT 的认证身份仅可用于创建最多10,000个广告。若达到此限制，您可以通过删除使用该身份的广告来释放与该身份关联的广告配额。若想查询此类身份的剩余广告配额，您可以调用/asset/bind/quota/并查看返回的 available_quota 字段值。
	IdentityID *string `json:"identity_id,omitempty"`
	// IdentityAuthorizedBcID 当identity_type 为 BC_AUTH_TT时必填。
	// 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。
	IdentityAuthorizedBcID *string `json:"identity_authorized_bc_id,omitempty"`
	// CatalogID 您为广告选择的商品库ID。商品来源为商品库的购物广告必填。广告中所展示的商品即来自该商品库。
	// 注意：自 2024 年 6 月 30 日起，广告组层级 product_source 为 STORE 时，您将无需再传入本字段，因为本字段将被忽略。
	CatalogID *string `json:"catalog_id,omitempty"`
	// ProductSpecificType 选择商品的维度
	// 仅以下任意场景下必填：
	// - shopping_ads_type 设置为 VIDEO 且 product_source 设置为 CATALOG。
	// - shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS 且 product_source 设置为 STORE。
	//
	// 枚举值：
	// - ALL：允许 TikTok 从所有商品中动态选择。
	// - PRODUCT_SET：请选择一个商品系列。TikTok 将动态选择该商品系列中的商品。
	// - CUSTOMIZED_PRODUCTS：选择自定义数量的特定商品。
	//
	// 若 shopping_ads_type 设置为 VIDEO 且 product_source 设置为 CATALOG，支持将本字段设置为 ALL ，PRODUCT_SET 或CUSTOMIZED_PRODUCTS 。
	// - 若本字段设置为 ALL，无需传入 sku_ids，item_group_ids 和 product_set_id。
	// - 若本字段设置为 PRODUCT_SET，需传入item_group_ids 和 product_set_id 其中之一。
	// - 若本字段设置为 CUSTOMIZED_PRODUCTS，需传入sku_ids 。
	//
	// 若 shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS 且 product_source 设置为 STORE，支持将本字段设置为 ALL 或 CUSTOMIZED_PRODUCTS。
	// - 若本字段设置为 ALL，无需传入 sku_ids，item_group_ids 和 product_set_id。
	// - 若本字段设置为 CUSTOMIZED_PRODUCTS，需传入item_group_ids 。
	//
	// 注意：
	// - 若 shopping_ads_type 设置为 VIDEO 且 product_source 设置为STORE，需通过 item_group_ids 字段直接指定商品。
	// - 若 shopping_ads_type 设置为 VIDEO 且 product_source 设置为SHOWCASE，需通过 showcase_products 字段直接指定商品。
	// - 若 shopping_ads_type 设置为 LIVE，则无需指定商品。
	ProductSpecificType enum.ProductSpecificType `json:"product_specific_type,omitempty"`
	// ItemGroupIDs 商品的 SPU （标准化产品单元）ID 列表
	// - 若shopping_ads_type设置为 VIDEO， product_source 设置为 CATALOG，且 product_specific_type 设置为 PRODUCT_SET，您需传入item_group_ids 和 product_set_id 其中之一。
	// - 若shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS , product_source 设置为 STORE ，且product_specific_type 设置为 CUSTOMIZED_PRODUCTS，本字段必填。
	//
	// 最大数量：
	// - shopping_ads_type 设置为 VIDEO， product_source 设置为 CATALOG，且 product_specific_type 设置为 PRODUCT_SET 时：20。
	// - shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS , product_source 设置为 STORE，且product_specific_type 设置为 CUSTOMIZED_PRODUCTS 时：50。
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
	// ProductSetID 商品系列ID
	// 若shopping_ads_type 设置为 VIDEO， product_source 设置为 CATALOG，且 product_specific_type 设置为 PRODUCT_SET，您需传入item_group_ids 和 product_set_id 其中之一。
	ProductSetID *string `json:"product_set_id,omitempty"`
	// SkuIDs 商品的 SKU（库存量单位）ID列表
	// shopping_ads_type 设置为 VIDEO ，product_source 设置为 CATALOG，且 product_specific_type 设置为 CUSTOMIZED_PRODUCTS 时必填
	// 最大数量：20
	SkuIDs []string `json:"sku_ids,omitempty"`
	// ShowcaseProducts 广告组层级product_source设置为SHOWCASE时必填。您可查看创建商品来源为橱窗的视频购物广告了解详情。
	// 您想要用在广告中的橱窗商品列表。最大数量：20。
	// 您可调用/showcase/product/get/，获取可用的橱窗商品。
	ShowcaseProducts []ShowcaseProduct `json:"showcase_products,omitempty"`
	// AdFormat 广告样式。
	// 枚举值：SINGLE_IMAGE，SINGLE_VIDEO，LIVE_CONTENT，CAROUSEL_ADS，CATALOG_CAROUSEL。
	// 本字段在以下场景必填：
	// 单视频广告需将该字段设为 SINGLE_VIDEO。
	// 单图片广告需将该字段设为SINGLE_IMAGE。
	// TikTok 标准轮播广告需将该字段设为 CAROUSEL_ADS。
	// 视频购物类型的TikTok轮播广告需将该字段设为CATALOG_CAROUSEL。
	// 点击这里查看创建轮播广告的步骤。
	// 直播购物广告需将该字段设为 LIVE_CONTENT。
	// 注意：
	//
	// 本字段设置为 CAROUSEL_ADS或CATALOG_CAROUSEL后不允许更新。
	AdFormat enum.AdFormat `json:"ad_format,omitempty"`
	// VerticalVideoStrategy 商品销量场景中使用的视频类型。
	// 枚举值： UNSET（未设置），SINGLE_VIDEO（单个视频），CATALOG_VIDEOS（使用视频模板的商品库视频），CATALOG_UPLOADED_VIDEOS（使用已上传视频的商品库视频）， LIVE_STREAM（直播视频）。
	//
	// 如果dynamic_format为DYNAMIC_CREATIVE，该字段必须设置为UNSET。
	//
	// 若想了解如何管理已上传的商品库视频，从而用于商品来源为商品库的视频购物广告中，可查看上传商品库视频并与商品绑定。
	//
	// 注意：
	//
	// 使用已上传的商品库视频创建广告样式为商品库视频、商品来源为商品库的视频购物广告（即vertical_video_strategy 设置为 CATALOG_UPLOADED_VIDEOS）目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 将 vertical_video_strategy 设置为 CATALOG_UPLOADED_VIDEOS 后，不允许将该字段更新为其他值。
	VerticalVideoStrategy enum.VerticalVideoStrategy `json:"vertical_video_strategy,omitempty"`
	// DynamicFormat 动态样式。
	// 枚举值：
	// DYNAMIC_CREATIVE：智能创意。根据用户的购买意图向每位用户展示不同的视频，商品卡片和目标页面，从而最大限度地提高广告转化量。
	// UNSET：未设置。
	DynamicFormat enum.DynamicFormat `json:"dynamic_format,omitempty"`
	// VideoID 视频 ID
	// - ad_format 设置为SINGLE_VIDEO且identity_type 为 CUSTOMIZED_USER时必填。
	// - ad_format 设置为SINGLE_VIDEO 且identity_type 为TT_USER或BC_AUTH_TT时，您需传入video_id 和 tiktok_item_id 其中之一。
	// - ad_format 设置为 SINGLE_IMAGE 或 CAROUSEL_ADS 时不支持本字段。
	// 若要创建 Spark Ads，您需使用 video_id 或 tiktok_item_id。若想了解 Spark Ads 详情，请查看创建Spark Ads。
	// 注意：我们建议您使用/file/video/ad/upload/(upload_type = UPLOAD_BY_VIDEO_ID)将视频 ID 与广告主 ID 进行绑定。
	VideoID *string `json:"video_id,omitempty"`
	// ImageIDs 图片 ID 列表（图片素材或视频封面）。
	// 可使用接口/file/image/ad/search/或/file/image/ad/upload/获取。
	// 当ad_format设置为SINGLE_VIDEO时：
	// 若想创建非 Spark Ads 的视频广告或通过 Spark Ads 推送创建 Spark Ads 视频广告，本字段必填且传入的值会作为视频封面（缩略图）（只允许传入一个值），此时您传入的图片宽高比需要与视频的宽高比保持一致。
	// 若想通过 Spark Ads 提取创建 Spark Ads 视频广告，不支持本字段。此时应传入tiktok_item_id 。
	// 当 ad_format设置为CAROUSEL_ADS时：
	// 若想创建非 Spark Ads 的标准轮播广告或通过 Spark Ads 推送创建 Spark Ads 标准轮播广告，本字段必填且本字段的 ID 数量范围为[1, 35]，且本字段指定的图片 ID 的顺序即为轮播广告中对应图片的展示顺序。
	// 请查看创建轮播广告，了解可在 TikTok 标准轮播广告中使用的图片的要求。
	// 若想通过 Spark Ads 提取创建 Spark Ads 标准轮播广告，不支持本字段。此时应传入tiktok_item_id。
	// 当 ad_format 设置为 SINGLE_IMAGE 时，本字段必填且您仅可向本字段传入一个图片 ID。
	// 若想了解单图片广告中的图片要求，请查看创建单图片广告。
	ImageIDs []string `json:"image_ids,omitempty"`
	// CarouselImageIndex 用于指定视频购物类型的轮播广告中要使用的附加图片的索引。若未传入本字段，将在视频购物类型的轮播广告中使用商品库商品的主图。
	// 仅当ad_format设置为CATALOG_CAROUSEL，且所指定的商品已设置对应的附加图片 URL 列表时生效。
	// 取值范围：[0, 9]。
	// 您通过本字段传入的数值代表每个商品库商品的附加图片列表（additional_image_urls）中要使用的附加图片的位置。例如，若carousel_image_index 设置为 1，代表您将使用每个additional_image_urls的值中的第二个图片作为轮播广告中使用的图片。您可通过/catalog/product/get/返回的additional_image_urls字段获取商品对应的附加图片列表。
	// 若您使用了本字段指定附加图片索引，系统将计算有可用附加图片的电商商品数量。若根据您提供的索引无法匹配到商品的对应附加图片，或从匹配到的附加图片 URL 拉取商品的附加图片失败，则该商品将不会计入。
	// 有可用附加图片的商品总数大于等于 2 时，仅在广告中展示这些有可用附加图片的商品，并使用商品的附加图片。
	// 有可用附加图片的商品总数小于 2 时，将展示所有商品，并均使用商品的主图。
	CarouselImageIndex *int `json:"carousel_image_index,omitempty"`
	// MusicID 想要在 TikTok 轮播广告中使用的音乐的 ID
	// 以下任意场景下必填：
	// - ad_format 设置为 CAROUSEL_ADS 且您想要创建非 Spark Ads 的标准轮播广告或通过 Spark Ads 推送创建 Spark Ads 标准轮播广告。
	// - ad_format 设置为 CATALOG_CAROUSEL
	// 想要在 TikTok 轮播广告中使用的音乐的 ID。
	// 请查看创建轮播广告-步骤-创建一个广告，了解如何获取可在TikTok轮播广告中使用的音乐ID。
	MusicID *string `json:"music_id,omitempty"`
	// TiktokItemID 用作 Spark Ads 的 TikTok 帖子 ID
	// - identity_type 设置为AUTH_CODE，TT_USER 或 BC_AUTH_TT 且您想要通过 Spark Ads 提取创建 Spark Ads 时必填。
	// - identity_type 设置为 CUSTOMIZED_USER 时不支持本字段。
	// 若想了解 Spark Ads 详情，请查看创建 Spark Ads。
	// ID 可以从/tt_video/info/或/identity/video/get/接口的返回值中的item_id字段值获取。
	// 注意：使用此功能，代表您声明拥有该视频中音乐的版权，可用于营销使用。
	TiktokItemID *string `json:"tiktok_item_id,omitempty"`
	// PromotionalMusicDisabled 是否关闭推广音乐开关。true： 关闭；false： 打开。默认true。如果您想为TikTok视频开启合拍或跟拍，请将该字段设置为false，即打开推广音乐。
	PromotionalMusicDisabled *bool `json:"promotional_music_disabled,omitempty"`
	// ItemDuetStatus 合拍开关状态。枚举值：ENABLE： 打开合拍；DISABLE： 禁止合拍。只有在promotional_music_disabled设置为false时，该字段才有效。
	ItemDuetStatus enum.EnableDisable `json:"item_duet_status,omitempty"`
	// ItemStitchStatus 跟拍开关状态。枚举值：ENABLE： 打开跟拍；DISABLE： 禁止跟拍。只有在promotional_music_disabled设置为false时，该字段才有效
	ItemStitchStatus enum.EnableDisable `json:"item_stitch_status,omitempty"`
	// DarkPostStatus 表示该视频是否为dark post。枚举值：ON，OFF
	DarkPostStatus enum.OnOff `json:"dark_post_status,omitempty"`
	// ShoppingAdsVideoPackageID 商品库视频模板ID。
	// 使用/catalog/video/get/接口获取视频包ID
	ShoppingAdsVideoPackageID *string `json:"shopping_ads_video_package_id,omitempty"`
	// AdText 广告文案，将作为广告创意的一部分展示给你的受众，向他们传达你想要推广的信息。不知道怎么写？试试智能文案。
	// - 对于单图片广告或视频广告（ad_format 为SINGLE_IMAGE或SINGLE_VIDEO），本字段必填。对于非Spark Ads广告，本字段必填。
	// - 广告文案长度允许值为 1-100 字符，不能包含emoji。
	// - 中文或日文每个字占用2个字符，每个英语字符占用1个字符。
	AdText *string `json:"ad_text,omitempty"`
	// CallToAction 行动引导文案，引导用户在看到您的广告后完成您期望的行为。
	// 对于TikTok广告，本字段或者call_to_action_id必须传入一个。
	// 枚举值见 枚举值-行动引导文案。
	// 直播购物广告中若creative_type设置为SHORT_VIDEO_LIVE，则本字段必填且必须设置为WATCH_LIVE。
	CallToAction enum.CallToAction `json:"call_to_action,omitempty"`
	// CallToActionID 行动引导文案素材包ID。行动引导文案素材包指的是一组自动优化的行动引导文案。
	// 如果本字段和call_to_action都传入，call_to_action会被忽略。
	// 关于自动优化行动引导文案的更多信息，请参见智能推荐CTA > 动态优选CTA。
	CallToActionID *string `json:"call_to_action_id,omitempty"`
	// CardID 创意素材包 ID。
	// 传入以下类型的创意素材包的 ID：
	// 行动引导文案。
	// 展示卡片（又称图片卡片）。
	// 网站信息卡片。
	// 下载卡。
	// 商品卡片。
	// 商品磁贴。
	// 倒计时贴纸。
	// 礼品码贴纸。
	// 弹出展示（又称福袋）。
	// 交互手势。
	// 点赞彩蛋。
	//
	// 若想创建素材包，可使用 /creative/portfolio/create/。
	// 若想了解如何获取展示卡片、网站信息卡片、下载卡、商品卡片或商品磁贴素材包的 ID，请参考卡片。
	// 若想了解如何获取倒计时贴纸或礼品码贴纸素材包的 ID，请参考贴纸。
	// 若想了解如何获取弹出展示、交互手势或点赞彩蛋素材包的 ID ，请参考高级创新互动样式。
	//
	// 注意：
	//
	// 商品磁贴素材包在广告中使用后，不允许再从广告中清空，但允许将其更新为其他的素材包。
	CardID *string `json:"card_id,omitempty"`
	// LandingPageURL 用户跳转至的落地页。
	// 落地页 URL 支持在末尾拼接 URL 参数。若想了解 URL 参数详情，请查看字段utm_params的对应描述。
	// 示例：https://www.example.com?utm_source=TikTok&utm_medium=video&utm_campaign=Campaign1&utm_content=Content1
	LandingPageURL *string `json:"landing_page_url,omitempty"`
	// UtmParams URL 参数列表。URL 参数是可以添加到 URL 末尾的代码片段，可以帮助您通过第三方分析平台监测来自不同渠道的点击，并了解访客与网站的互动情况。每个 URL 参数由key和value指定的键值对组成。
	// 本字段支持在以下场景中使用：
	// - 商品来源为商品库的视频购物广告场景：
	//   - 推广系列层级：
	//     - objective_type 设置为 PRODUCT_SALES
	//   - 广告组层级：
	//     - shopping_ads_type 设置为 VIDEO（视频购物广告）
	//     - product_source 设置为 CATALOG
	//     - promotion_type 设置为 WEBSITE，APP_ANDROID，或 APP_IOS
	// - 非视频购物广告场景：
	//   - 推广系列层级：
	//     - objective_type 设置为 REACH，TRAFFIC，VIDEO_VIEWS， LEAD_GENERATION，ENGAGEMENT，APP_PROMOTION （app_promotion_type 需设置为APP_PREREGISTRATION），WEB_CONVERSIONS 或 RF_REACH
	//
	// 最大数量：14。
	//
	// - 若您在广告中使用商品库链接（shopping_ads_fallback_type 为 SHOPPING_ADS），您可通过传入utm_params为商品库链接添加 URL 参数。URL 参数将在广告投放时自动拼接到商品库链接。
	// - 若您在广告中使用自定义链接（网站 URL）（shopping_ads_fallback_type 为 CUSTOM），您可手动为自定义链接添加 URL 参数，然后将 landing_page_url 直接设置为带有 URL 参数的 URL。
	//   - 然而，若您在广告中使用自定义链接且同时传入了 utm_params，在广告投放时 URL 参数不会自动拼接到自定义链接（landing_page_url）。此时，您需确保通过landing_page_url 指定的 URL 已经带有 URL 参数，且utm_params 与 URL 中的 URL 参数完全匹配。否则，将出现报错。
	// - 在非视频购物广告场景，若您将 landing_page_url 设置为一个已带有 URL 参数的 URL，您可同时传入 utm_params，以记录 URL 中带有的 URL 参数。此时，您需确保 utm_params 与 URL中的 URL 参数完全匹配。在广告投放时 URL 参数不会自动拼接到landing_page_url。
	//
	// 注意：在非视频购物广告场景使用 utm_params 目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// ----
	// KEY
	// 您可传入自定义参数或 UTM 参数。
	// 支持的 UTM 参数有：
	// utm_source：为您的网站带来访问量的应用程序、网站等。示例：TikTok。
	// utm_medium：广告或营销媒介。示例：cpm、cpc、横幅广告、视频。
	// utm_content：用于推广的创意内容。示例：广告名称、CTA文本、资产、颜色等。
	// utm_campaign：单个推广系列的名称、标语或促销代码。示例：BlackFridayProm。
	//
	// 请注意，UTM 参数区分大小写。
	//
	// 若传入自定义参数，则长度限制为 100 个字符。
	// ----
	// VALUE
	// URL 参数的值。
	// 您可以传入自定义值或宏的名称。
	//
	// 支持的宏包括：
	// __CAMPAIGN_NAME__: 该值将被替换为您的推广系列名称。
	// __CAMPAIGN_ID__: 该值将被替换为您的推广系列ID。
	// __AID_NAME__: 该值将被替换为您的广告组名称。
	// __AID__: 该值将被替换为您的广告组ID。
	// __CID_NAME__: 该值将被替换为您的广告名称。
	// __CID__: 该值将被替换为您的广告ID。
	// __PLACEMENT__: 该值将被替换为您的版位。
	//
	// 若传入自定义值，则长度限制为 600 个字符。
	UtmParams []model.KV `json:"utm_params,omitempty"`
	// PageID 页面 ID。
	// 若想获取您的广告账号中的页面，可使用/page/get/。
	// 若想创建即时体验页面，可使用即时体验页面编辑器。
	//
	// 即时表单仅当推广系列的推广目标为 LEAD_GENERATION 时有效，仅支持 TikTok 版位。（首次创建前需要通过协议管理-签订签订 LeadAds 协议）。
	// 注意: 我们将于2023年2月16起不再支持创建精品栏广告。为保证流畅的API体验，请勿将商品橱窗页面设为落地页。
	PageID uint64 `json:"page_id,omitempty"`
	// CppURL 自定产品页面 URL。
	// 自定产品页面是一种定制化的产品页面，支持自定义截图、推广文字和应用预览，可用于突出介绍所推广应用的特定内容或功能，或针对特定受众群体进行推广。
	// 自定产品页面 URL 需与所推广的应用（app_id）保持一致。
	// 格式：https://apps.apple.com/{region}/app/{app_name}/id{app_id}?ppid={ppid}。
	// 长度限制：512 字符。
	// 示例：https://apps.apple.com/us/app/tiktok/id835599320?ppid=12345a6b-c789-12d3-e4f5-g6h78i91jk2l。
	// 若想了解支持使用自定产品页面的场景以及相关步骤，请查看自定产品页面。
	// 注意：在广告中使用自定产品页面目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	CppURL *string `json:"cpp_url,omitempty"`
	// TiktokPageCategory 您希望推广的 TikTok 页面的类别
	// 当广告组层级的optimization_goal 为 PAGE_VISIT时，此字段为必填项。
	// 枚举值：
	// - PROFILE_PAGE：TikTok 账号主页。无需传入landing_page_url 或 page_id。
	// - OTHER_TIKTOK_PAGE：其他 TikTok 页面，包括播放列表页面、话题标签页面、音乐页面和商业化特效页面。您需同时通过 landing_page_url 指定具体的 TikTok 页面 URL。
	//   - 若想获取 TikTok 页面 URL，可在 TikTok 应用中点按你要推广的 TikTok 页面右上角的分享按钮，然后点击复制 URL。
	// - TIKTOK_INSTANT_PAGE：自定义 TikTok 即时体验页面。您需同时通过 page_id 指定具体的自定义 TikTok 即时体验页面。
	//   - 若想获取广告账户中的自定义 TikTok 即时体验页面，可使用 /page/get/ 并将 business_type 设置为 TIKTOK_INSTANT_PAGE。
	//   注意：
	//
	// 广告创建后，此字段不支持更新。
	// 若您将 tiktok_page_category 设置为 OTHER_TIKTOK_PAGE，则广告创建后不支持更新 landing_page_url
	TiktokPageCategory enum.TiktokPageCategory `json:"tiktok_page_category,omitempty"`
	// PhoneRegionCode 受众可点击广告拨打的电话号码的地区代码
	// 若广告组层级 promotion_type 设置为 LEAD_GEN_CLICK_TO_CALL ， 则 phone_number ， phone_region_calling_code 以及 phone_region_code 参数均为必填。
	// 若想了解如何创建推广对象类型为电话通话的线索广告，请查看这里。
	// 若想获取与某一电话号码所绑定的地区对应的地区代码（ phone_region_code ）和电话区号 （ phone_region_calling_code ），需使用/tool/phone_region_code/。
	// 注意：创建推广对象类型为电话通话的线索广告目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	PhoneRegionCode *string `json:"phone_region_code,omitempty"`
	// PhoneRegionCallingCode 受众可点击广告拨打的电话号码的区号
	// 若广告组层级 promotion_type 设置为 LEAD_GEN_CLICK_TO_CALL ， 则 phone_number ， phone_region_calling_code 以及 phone_region_code 参数均为必填。
	// 若想了解如何创建推广对象类型为电话通话的线索广告，请查看这里。
	// 若想获取与某一电话号码所绑定的地区对应的地区代码（ phone_region_code ）和电话区号 （ phone_region_calling_code ），需使用/tool/phone_region_code/。
	// 注意：创建推广对象类型为电话通话的线索广告目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	PhoneRegionCallingCode *string `json:"phone_region_calling_code,omitempty"`
	// PhoneNumber 受众可点击广告拨打的电话号码
	// 若广告组层级 promotion_type 设置为 LEAD_GEN_CLICK_TO_CALL ， 则 phone_number ， phone_region_calling_code 以及 phone_region_code 参数均为必填。
	// 若想了解如何创建推广对象类型为电话通话的线索广告，请查看这里。
	// 注意：创建推广对象类型为电话通话的线索广告目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Deeplink 深度链接，即已安装应用的用户点击链接后跳转至的特定页面。
	// 参见deeplink_type。
	//
	// 若想了解非延迟深度链接和延迟深度链接的支持场景，请查看“深度链接 - 支持范围 - 深度链接”小节。
	// 注意：
	//
	// 延迟深度链接目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。请注意，设置延迟深度链接需要相应的技术资源和开发能力。
	// 非延迟深度链接在某些场景下目前为白名单功能。若想了解非延迟深度链接为白名单功能的具体场景，请查看“深度链接 - 支持范围 - 深度链接”小节。如需在这些场景中使用此功能，请联系您的TikTok销售代表。
	Deeplink *string `json:"deeplink,omitempty"`
	// DeeplinkType 深度链接类型。
	// 枚举值：
	// NORMAL：非延迟深度链接。
	// DEFERRED_DEEPLINK：延迟深度链接。
	//
	// 默认值：NORMAL。
	//
	// 若想了解非延迟深度链接和延迟深度链接的支持场景，请查看“深度链接 - 支持范围 - 深度链接”小节。
	// 注意：
	//
	// 延迟深度链接目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。请注意，设置延迟深度链接需要相应的技术资源和开发能力。
	// 非延迟深度链接在某些场景下目前为白名单功能。若想了解非延迟深度链接为白名单功能的具体场景，请查看“深度链接 - 支持范围 - 深度链接”小节。如需在这些场景中使用此功能，请联系您的TikTok销售代表。
	// 专属推广系列中不支持延迟推广系列。因此，若广告所属的推广系列的 campaign_type 为 IOS14_CAMPAIGN，则不支持将本字段设置为 DEFERRED_DEEPLINK 。
	DeeplinkType enum.DeeplinkType `json:"deeplink_type,omitempty"`
	// DeeplinkFormatType 深度链接的格式类型
	// 推广系列层级 objective_type 设置为 TRAFFIC ，且广告组层级optimization_goal 设置为 DESTINATION_VISIT 时，deeplink_format_type 和 deeplink 均必填。
	// 枚举值：
	// UNIVERSAL_OR_APP：通用/应用链接，即以http://或 https:// 开头的 iOS 通用链接或 Android 应用链接。
	// SCHEME：URL 架构，即格式为scheme://resource 的自定义网址架构。例如，WhatsApp 的自定义网址架构格式应为whatsapp://resource。
	//
	// 若想了解如何在访问量广告中使用目标页面访问这一优化目标，请查看优化访问量广告中的目标页面访问。
	//
	// 注意：本字段设置后不允许更新。
	DeeplinkFormatType enum.DeeplinkFormatType `json:"deeplink_format_type,omitempty"`
	// ShoppingAdsDeeplinkType 在购物广告中要使用的深度链接的来源。
	//
	// 枚举值：
	// NONE：不设置深度链接。
	// CUSTOM：自定义深度链接。在广告中使用你提供的深度链接。您需同时传入deeplink。
	// SHOPPING_ADS：商品库深度链接。使用商品库中为各商品添加的深度链接。
	//
	// 默认值：NONE。
	//
	// SHOPPING_ADS 的值不支持在满足以下条件的场景中使用：
	// 广告组层级：
	// shopping_ads_type 为 VIDEO （视频购物广告）
	// 且 product_source 为 CATALOG
	// 广告层级：
	// ad_format 为 SINGLE_VIDEO 且 vertical_video_strategy 为 SINGLE_VIDEO
	// 或 ad_format 为 CATALOG_CAROUSEL
	//
	// 若想了解非延迟深度链接和延迟深度链接的支持场景，请查看“深度链接 - 支持范围 - 深度链接”小节。
	// 注意：
	//
	// 延迟深度链接目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。请注意，设置延迟深度链接需要相应的技术资源和开发能力。
	// 非延迟深度链接在某些场景下目前为白名单功能。若想了解非延迟深度链接为白名单功能的具体场景，请查看“深度链接 - 支持范围 - 深度链接”小节。如需在这些场景中使用此功能，请联系您的TikTok销售代表。
	ShoppingAdsDeeplinkType enum.ShoppingAdsDeeplinkType `json:"shopping_ads_deeplink_type,omitempty"`
	// DeeplinkUtmParams 深度链接 URL 参数列表。
	// 注意：在非视频购物广告场景使用 utm_params 目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	DeeplinkUtmParams []model.KV `json:"deeplink_utm_params,omitempty"`
	// ShoppingAdsFallbackType 在购物广告再营销场景下，深度链接唤起失败后的返回落地页（失败路径）类型。
	//
	// 枚举值：
	// DEFAULT：未设置。
	// CUSTOM: 自定义链接。您可以在广告组中提供一个网页地址, 并回退到这个地址。您需同时传入landing_page_url。
	// SHOPPING_ADS：商品库链接。 回退到目录中为各商品添加的网页地址。
	//
	// SHOPPING_ADS 的值不支持在满足以下条件的场景中使用：
	// 广告组层级：
	// shopping_ads_type 为 VIDEO （视频购物广告）
	// 且 product_source 为 CATALOG
	// 广告层级：
	// ad_format 为 SINGLE_VIDEO 且 vertical_video_strategy 为 SINGLE_VIDEO
	// 或 ad_format 为 CATALOG_CAROUSEL
	//
	// 若想了解失败路径的支持场景，请查看“深度链接 - 支持范围 - 失败路径”小节。
	ShoppingAdsFallbackType enum.ShoppingAdsFallbackType `json:"shopping_ads_fallback_type,omitempty"`
	// FallbackType 失败路径。如果用户没有安装过应用，您可以选择让用户回退去安装应用，或者去您想推广的网页。
	// 已设置延迟深度链接（deeplink_type 为 DEFERRED_DEEPLINK）时不支持
	// 枚举值：
	// APP_INSTALL：广告组层级通过 app_id 指定的推广应用对应的 App Store 或 Google Play 页面。
	// WEBSITE：通过 landing_page_url 指定的推广网页。
	// UNSET：未设置。
	//
	// 若本字段设置为 WEBSITE，需传入 landing_page_url 。
	//
	// 若想了解失败路径的支持场景，请查看“深度链接 - 支持范围 - 失败路径”小节。
	FallbackType enum.FallbackType `json:"fallback_type,omitempty"`
	// DynamicDestination 动态目标页面策略。
	//
	// 枚举值：
	// DLP：动态落地页。基于用户的主页、行为和意向，向每位用户展示不同的目标页面，包括自定义网站或即时体验商品页面等，从而最大限度地提高广告转化量。
	// UNSET：未设置。
	//
	// 若promotion_type 为APP_ANDROID或APP_IOS，则此字段不可设置为DLP。
	//
	// 注意：动态目标页面在商品来源为商品库的视频购物广告推广网站且使用优化目标点击或落地页浏览量时不支持使用。因此，满足以下条件时本字段不支持设置为DLP：
	//
	// 广告组层级：
	// shopping_ads_type 设置为 VIDEO（视频购物广告）。
	// product_source 设置为 CATALOG。
	// promotion_type 设置为 WEBSITE。
	// optimization_goal 设置为 CLICK 或 TRAFFIC_LANDING_PAGE_VIEW。
	DynamicDestination enum.DynamicDestination `json:"dynamic_destination,omitempty"`
	// InstantProductPageUsed 是否使用TikTok即时体验商品页面（自动生成的TikTok即时体验页面，展示多种商品，提供原生落地页体验）。
	// 若dynamic_destination为 DLP，则不需要传入该参数。
	// 若dynamic_destination 为UNSET，您可以将本字段设置为true，且无需同时传入page_id。
	InstantProductPageUsed *bool `json:"instantt_product_page_used,omitempty"`
	// PageImageIndex 仅当instant_product_page_used设置为true，且所指定的商品已设置对应的附加图片列表时生效。
	// 用于指定TikTok 即时体验商品页面中要使用的附加图片的索引。若未传入本字段，将在TikTok 即时体验商品页面中使用商品库商品的主图。
	// 取值范围：[0, 9]。
	// 您通过本字段传入的数值代表每个商品库商品的附加图片列表（additional_image_urls）中要使用的附加图片的位置。例如，若page_image_index 设置为 1，代表您将使用每个additional_image_urls的值中的第二个图片作为TikTok 即时体验商品页面中使用的图片。您可通过/catalog/product/get/返回的additional_image_urls字段获取商品对应的附加图片列表。
	// 若您指定的数值超过了商品对应的附加图片列表中附加图片URL的数量，将忽略本字段，使用商品的主图（image_url）。您可通过/catalog/product/get/返回的image_url 字段获取商品对应的主图。
	PageImageIndex *int `json:"page_image_index,omitempty"`
	// AutoMessageID TikTok 私信广告中使用的自动消息的 ID。
	// 目前唯一支持的自动消息类型为欢迎消息
	// 仅当广告组层级 promotion_type 设置为 LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE 时必填。
	// 若想获取广告账户中欢迎消息的列表，可使用/creative/auto_message/get/。
	// 若想在广告账户中创建欢迎消息，可使用/creative/auto_message/create/。
	// 若想了解如何创建使用了欢迎消息的 TikTok 私信广告，可查看创建优化位置为 TikTok 私信的线索广告。
	AutoMessageID *string `json:"auto_message_id,omitempty"`
	// AigcDisclosureType 是否启用 AI 生成内容自主声明开关，以表明广告中包含 AI 生成内容。启用该开关后，当用户查看完整广告时，您的广告将带有“广告主标记为 AI 生成”的标签。
	// 枚举值：
	// SELF_DISCLOSURE：启用开关，声明广告包含 AI 生成内容。
	// NOT_DECLARED：不声明广告包含 AI 生成内容。
	// 默认值： NOT_DECLARED
	// 若想了解此开关支持的推广目标和广告样式，以及使用此开关的步骤，请查看 AI 生成内容自主声明开关。
	//
	// 注意：
	//
	// 对于广告样式为单视频或标准轮播的广告，您仅可在更新广告素材（视频或图片）时更新本字段。例如，若您仅修改标准轮播广告中图片的顺序，则不支持更新本字段。
	// 对于广告样式为动态样式、商品库视频或商品库轮播的广告，不支持更新本字段。
	AigcDisclosureType enum.AigcDisclosureType `json:"aigc_disclosure_type,omitempty"`
	// DisclaimerType 您想要在广告中添加的免责声明类型。枚举值：TEXT_LINK（文字链免责声明），TEXT_ONLY（纯文本免责声明）。请查看使用免责声明，了解如何进行免责声明设置。
	// 注意：
	//
	// 免责声明目前对注册地非美国或加拿大且使用竞价推广目标的广告主，以及使用覆盖和频次推广目标的所有广告主为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 免责声明功能支持以下推广目标：APP_PROMOTION，WEB_CONVERSIONS， REACH，TRAFFIC， VIDEO_VIEWS，ENGAGEMENT，LEAD_GENERATION，RF_REACH。
	// 免责声明功能仅支持TikTok版位。
	// 免责声明功能不支持在程序化创意广告中使用。
	// 广告中加入免责声明后，不允许删除免责声明。
	DisclaimerType enum.DisclaimerType `json:"disclaimer_type,omitempty"`
	// DisclaimerText 您想要在广告中添加的纯文本免责声明。当disclaimer_type为TEXT_ONLY时必填。
	DisclaimerText *DisclaimerText `json:"disclaimer_text,omitempty"`
	// DisclaimerClickableTexts 您想要在广告中添加的文字链免责声明。当disclaimer_type为 TEXT_LINK时必填。 最大数量：3。
	// 若使用多个免责声明，则所有文案的总长度以及单个文案的长度均应不超过40个字符。
	DisclaimerClickableTexts []DisclaimerClickableText `json:"disclaimer_clickable_texts,omitempty"`
	// TrackingPixelID 要追踪的 Pixel 的 ID。您可以使用此字段监测发生在外部网站的事件。支持的推广目标包括竞价广告推广目标(REACH,VIDEO_VIEWS,TRAFFIC, WEB_CONVERSIONS, LEAD_GENERATION, APP_PROMOTION, PRODUCT_SALES, ENGAGEMENT) 和覆盖和频次广告推广目标(RF_REACH)。
	// 竞价广告推广目标
	// 如果在 广告组 层级中 pixel_id!=0 且pixel用于追踪外部网站事件， 那么 tracking_pixel_id 中传入的值须和广告组中 pixel_id 的值保持一致。否则您可以在 tracking_pixel_id 中填入任意有效 ID 值。
	// 覆盖和频次广告推广目标
	// 您可以在 tracking_pixel_id中填入任意有效ID值。
	TrackingPixelID uint64 `json:"tracking_pixel_id,omitempty"`
	// TrackingAppID 您想要监测的应用的 ID。您可以使用此字段监测发生在外部网站的应用事件。
	// 支持的推广目标包括竞价广告推广目标（REACH，VIDEO_VIEWS， TRAFFIC， WEB_CONVERSIONS，LEAD_GENERATION， APP_PROMOTION，PRODUCT_SALES，ENGAGEMENT）和覆盖和频次广告推广目标（RF_REACH）。
	// 若已在广告组层级已设置 app_id ，且您想要追踪站外应用事件，则您通过本字段传入的App ID必须为您在广告组层级设置的App ID。若未在广告组层级设置app_id，则您可以通过本字段传入任意想追踪的App的ID。
	// 您可以通过/app/list/ 接口获取App ID（app_id）。
	TrackingAppID *string `json:"tracking_app_id,omitempty"`
	// TrackingOfflineEventSetIDs 想要监测的线下事件组 ID 列表。您可以使用本字段监测和衡量浏览广告或与之互动的用户的线下活动。
	// 最大数量：50。
	// 点击这里了解如何创建和管理线下事件组。
	// 注意：
	//
	// 新建的广告中需对当前开启自动监测的线下事件组进行追踪：
	// 您可不传入本字段。本字段的值将自动默认为广告账户下当前所有开启自动监测的线下事件组，或
	// 您可手动传入本字段。您需至少传入广告账户下当前所有开启自动监测的线下事件组，未开启自动监测的线下事件组为可选。欲获取广告账户下当前所有开启自动监测的线下事件组的ID，需调用/offline/get/先获取广告账户下的所有线下事件组ID，再根据返回的auto_tracking值筛选出其中开启自动监测的线下事件组的ID。
	// 本字段支持以下推广目标： REACH，VIDEO_VIEWS，TRAFFIC, LEAD_GENERATION，APP_PROMOTION，WEB_CONVERSIONS，ENGAGEMENT，和PRODUCT_SALES。
	// 若推广目标（objective_type）设置为PRODUCT_SALES，您仅可在商品来源为商品库的视频购物广告中使用本字段。因此广告组层级需满足以下要求：
	// shopping_ads_type = VIDEO ，且
	// product_source = CATALOG。
	TrackingOfflineEventSetIDs []string `json:"tracking_offline_event_set_ids,omitempty"`
	// TrackingMessageEventSetID 即时通讯广告中监测的消息事件集的 ID
	// 仅当广告组层级同时满足以下条件时生效：
	// - promotion_type 为LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE
	// - optimization_goal 为 CLICK
	// - messaging_app_type 为 MESSENGER 或 WHATSAPP
	// - 未传入 page_id
	//
	// 若想获取可用的消息事件集列表，可使用/ctm/message_event_set/get/。
	// 若想了解如何创建 TikTok 即时通讯广告，可查看创建优化位置为即时通讯应用的线索广告。
	//
	// 注意：本字段设置后不支持更新。
	TrackingMessageEventSetID *string `json:"tracking_message_event_set_id,omitempty"`
	// ViewabilityPostbidPartner 出价后第三方可见性监测合作伙伴。枚举值：UNSET，MOAT，DOUBLE_VERIFY，IAS，ZEFR。
	// 注意：
	// 出价后第三方可见性监测目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 请查看品牌安全，了解出价后监测的支持推广目标，支持市场以及简要介绍。
	// 当您同时传入brand_safety_postbid_partner 和viewability_postbid_partner时：
	// 一旦将两字段其中之一设置为 IAS，则另一字段也需要设置为IAS，且您需要为brand_safety_vast_url和viewability_vast_url传入相同的值（即您从监测合作伙伴IAS处获取的封装的 VAST 网址）。
	// 若您设置了两个不同的出价后第三方监测合作伙伴，仅出价后第三方品牌安全监测设置生效。
	// 出价后第三方监测不支持在程序化创意广告中使用。
	// 出价后第三方监测不支持在promotion_type设置为LIVE_SHOPPING的广告组中使用。
	ViewabilityPostbidPartner enum.ViewabilityPostbidPartner `json:"viewability_postbid_partner,omitempty"`
	// ViewabilityVastURL 出价后第三方监测合作伙伴用于监测可见性的封装的 VAST 网址。当brand_safety_postbid_partner非IAS时本字段为空值。
	// 注意：
	//
	// 出价后第三方可见性监测目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 当您将brand_safety_postbid_partner和 viewability_postbid_partner均设置为IAS时，brand_safety_vast_url 和 viewability_vast_url的值需需保持一致。
	// 出价后第三方监测不支持在程序化创意广告中使用。
	// 出价后第三方监测不支持在promotion_type设置为LIVE_SHOPPING的广告组中使用。
	ViewabilityVastURL *string `json:"viewability_vast_url,omitempty"`
	// BrandSafetyPostbidPartner 出价后第三方品牌安全监测合作伙伴。枚举值： UNSET，DOUBLE_VERIFY，IAS，ZEFR。
	// 注意：
	//
	// 出价后第三方品牌安全监测目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 请查看品牌安全，了解出价后监测的支持推广目标，支持市场以及简要介绍。
	// 您在广告组层级配置的出价前品牌安全筛选设置决定了您能否同时使用出价后第三方品牌安全监测和出价后第三方可见性监测：
	// 若您在广告组层级将brand_safety_type设置为 NO_BRAND_SAFETY ，则在广告组层级可传入viewability_postbid_partner，不能传入brand_safety_postbid_partner 。
	// 若您在广告组层级将brand_safety_type设置的值非 NO_BRAND_SAFETY ，则在广告组层级可同时传入viewability_postbid_partner和brand_safety_postbid_partner。
	// 当您同时传入brand_safety_postbid_partner 和viewability_postbid_partner时：
	// 一旦将两字段其中之一设置为IAS，则另一字段也需要设置为IAS，且您需要为brand_safety_vast_url和viewability_vast_url传入相同的值（即您从监测合作伙伴IAS处获取的封装的 VAST 网址）。
	// 若您设置了两个不同的出价后第三方监测合作伙伴，仅出价后第三方品牌安全监测设置生效。
	// 出价后第三方监测不支持在程序化创意广告中使用。
	// 出价后第三方监测不支持在promotion_type设置为LIVE_SHOPPING的广告组中使用。
	BrandSafetyPostbidPartner enum.BrandSafetyPostbidPartner `json:"brand_safety_postbid_partner,omitempty"`
	// BrandSafetyVastURL 出价后第三方监测合作伙伴用于监测品牌安全的封装的 VAST 网址。当brand_safety_postbid_partner非IAS时本字段为空值。
	// 注意：
	//
	// 出价后第三方品牌安全监测目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 当您将brand_safety_postbid_partner和 viewability_postbid_partner均设置为 IAS时，brand_safety_vast_url 和 viewability_vast_url的值需需保持一致。
	// 出价后第三方监测不支持在程序化创意广告中使用。
	// 出价后第三方监测不支持在promotion_type设置为LIVE_SHOPPING的广告组中使用。
	BrandSafetyVastURL *string `json:"brand_safety_vast_url,omitempty"`
	// ImpressionTrackingURL 默认展示监测 URL
	// 由数据合作伙伴生成的 URL，用于跟踪广告中的展示事件。 您一般可以在合作伙伴的平台上找到此 URL，并可以复制。
	// 若您希望监测的应用（即广告组层级设置的app_id）的合作伙伴ID是44 （TikTok Business SDK）或 49（TikTok App API），则您不需要传入本字段。若传入本字段，则将忽略本字段。您可通过/app/list/ 或 /app/info/返回的partner_id字段获取应用的合作伙伴ID。
	// 若您希望监测的应用（即广告组层级设置的app_id ）已开启自归因，且其合作伙伴ID不是44（TikTok Business SDK）或 49（TikTok App API），则您同样不需要传入本字段，因为本字段将默认设置为您为应用配置的默认展示监测 URL，且该URL不支持更新。您可通过 /app/list/ 或 /app/info/返回的self_attribution_enabled字段查看应用是否已开启自归因。
	ImpressionTrackingURL *string `json:"impression_tracking_url,omitempty"`
	// ClickTrackingURL 点击监测 URL
	// 由数据合作伙伴生成的 URL，用于跟踪广告中的点击事件。 您通常可以在合作伙伴的平台上找到此 URL，并可以复制。
	// 若您希望监测的应用（即广告组层级设置的app_id ）的合作伙伴ID是44（TikTok Business SDK）或 49（TikTok App API），则您不需要传入本字段。若传入本字段，则将忽略本字段。您可通过/app/list/ 或/app/info/返回的partner_id字段获取应用的合作伙伴ID。
	// 若您希望监测的应用（即广告组层级设置的app_id ）已开启自归因，且其合作伙伴ID不是44（TikTok Business SDK）或 49（TikTok App API），则您同样不需要传入本字段，因为本字段将默认设置为您为应用配置的点击监测URL，且该URL不支持更新。您可通过/app/list/ 或/app/info/返回的self_attribution_enabled字段查看应用是否已开启自归因。
	// 目前 Pangle 不支持 DCM、Sizmek 或 Flashtalking。
	ClickTrackingURL *string `json:"click_tracking_url,omitempty"`
	// PlayableURL 试玩广告 URL
	// 在 Pangle 和 TikTok 版位生效。可通过/playable/get/接口获取。同一个广告组共用同一个试玩广告。
	// 注意：自 2024 年 9 月 30 日起，不再支持创建投放至 TikTok 版位的试玩广告。
	PlayableURL *string `json:"playable_url,omitempty"`
	// OperationStatus 广告的操作状态。
	// 注意
	// 对于覆盖和频次广告，需将operation_status设置为ENABLE，或不传入operation_status，请不要将operation_status设置为DISABLE。
	// 枚举值：
	// ENABLE：广告创建时处于启用状态。
	// DISABLE：广告创建时处于未启用状态。
	// 默认值：ENABLE。若想在广告创建后更新广告的启用/关闭状态，请使用 /ad/status/update/接口。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// CreativeType 直播购物广告、商品购物广告或应用预注册场景中的创意素材类型。枚举值：
	// SHOP_PDP：短视频带货广告，包含一个商品链接。
	// SHOP_PLP: 短视频带货广告，包含6-50个商品链接。
	// SHORT_VIDEO_LIVE: 短视频引流至直播间广告。
	// DIRECT_LIVE: 直播间带货。每个推广系列只能有一个直播间。
	// PSA：商品视觉元素将作为广告创意使用。
	// CUSTOM_INSTANT_PAGE：自定义 TikTok 即时体验页面。
	//
	// 若广告组层级 shopping_ads_type 设置为 LIVE ，本字段必填，且需设置为 SHORT_VIDEO_LIVE 或 DIRECT_LIVE。
	// 若广告组层级 shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS，可省略本字段或将本字段设置为PSA。
	// 当推广系列层级 app_promotion_type设置为APP_PREREGISTRATION，且广告组层级promotion_website_type设置为TIKTOK_NATIVE_PAGE时，您可省略本字段或将本字段设置为CUSTOM_INSTANT_PAGE。
	CreativeType enum.CreativeType `json:"creative_type,omitempty"`
	// VideoViewTrackingURL 视频播放监测URL。
	VideoViewTrackingURL *string `json:"video_view_tracking_url,omitempty"`
	// AppName 广告展示的App名称。默认为应用商店名称，长度限制：1-40个字符。如果应用商店中的应用名称超过了40个字符，您必须使用本字段传入一个新的应用名称，否则创建广告将失败。
	AppName *string `json:"app_name,omitempty"`
	// DisplayName 落地页或纯曝光广告的显示名称，长度限制：1-40个英文字符或1-20个中文/日文/韩文字符，当推广对象类型为落地页或纯曝光时必填
	DisplayName *string `json:"display_name,omitempty"`
	// AvatarIconWebURL 广告头像图片 ID，可以通过素材管理-图片上传接口上传图片（图片比例要求1:1）
	AvatarIconWebURL *string `json:"avatar_icon_web_url,omitempty"`
	// CreativeAuthorized 是否允许在我们的 TikTok For Business 创意中心展示您的广告。仅对非美国广告主有效。默认值：false。
	// 注意：creative_authorized仅可用于视频广告。
	CreativeAuthorized *bool `json:"creative_authorized,omitempty"`
	// ScheduleID 仅对覆盖和频次广告生效。
	// 广告投放信息ID。
	// 本字段用于将广告投放信息ID与您要在已设置顺序投放的覆盖和频次广告组中创建的广告绑定。
	// 若想获取广告投放信息ID，可调用/adgroup/rf/create/并将delivery_mode 设置为 SEQUENCE。广告投放信息与广告绑定后，is_draft字段的值将变为false，而schedule_id字段的值将变为与广告投放信息绑定的广告的 ID。若想确认广告投放信息ID与广告已成功绑定，调用/adgroup/get/并查看返回中is_draft 和schedule_id字段的值。
	ScheduleID *string `json:"schedule_id,omitempty"`
}
