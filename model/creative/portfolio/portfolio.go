package portfolio

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/showcase"
)

// Portfolio 素材包
type Portfolio struct {
	// CreativePortfolioID 创意素材包 ID
	CreativePortfolioID string `json:"creative_portfolio_id,omitempty"`
	// CreativePortfolioType 素材包类型。
	// 枚举值：
	// CTA：行动引导文案。
	// CARD：卡片（展示卡片）。
	// WEB_INFO_CARD：网站信息卡片。
	// DOWNLOAD_CARD：下载卡。
	// PRODUCT_CARD：商品卡片。
	// PRODUCT_TILE： 商品磁贴。
	// STICKER：倒计时贴纸或礼品码贴纸。
	// PREMIUM_BADGE：弹出展示（又称福袋）。
	// GESTURE：交互手势。
	// SUPER_LIKE：点赞彩蛋。点赞彩蛋能够在受众为视频广告点赞后，为受众提供令人惊喜的视觉元素。有了点赞彩蛋，当受众点赞或点按两次广告时，屏幕上会飘过心形图标和自定义图标。
	//
	// 默认值：CTA。
	//
	// 注意：
	//
	// 所有类型的素材包单次仅可创建一个。
	// 网站信息卡片可用于objective_type为WEB_CONVERSIONS 或 TRAFFIC（promotion_type 需设置为 WEBSITE ）的竞价广告。
	// 商品卡片目前仅支持在视频购物广告中使用。
	// 商品磁贴目前仅支持在商品来源为商品库的视频购物广告中使用。
	// 在创建商品磁贴素材包前，需确保您的商品库中至少包含 4 个过审的商品。若想确认您的商品已有足够的过审商品用于商品磁贴素材包，可使用/catalog/product/get/并确认至少 4 个商品的 audit_status为approved 。
	// 使用商品磁贴素材包创建广告目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 倒计时贴纸可用于以覆盖人数（ REACH ）、访问量 （ TRAFFIC ）、视频播放量（ VIDEO_VIEW ）、应用推广 （ APP_PROMOTION）、线索收集（ LEAD_GENERATION ）、网站转化量 （ WEB_CONVERSIONS ）或商品销量（ PRODUCT_SALES ） 为推广目标（ objective_type ）的竞价广告，以及以覆盖人数（ RF_REACH ）为推广目标（ objective_type ）的覆盖和频次广告。
	// 若您想在线索广告中使用倒计时贴纸，您在广告组层级仅可将优化位置设置为落地页。若想了解如何创建此类广告，请查看创建优化位置为落地页的线索广告。
	// 在推广目标为访问量 （ TRAFFIC ）、应用推广（ APP_PROMOTION ）和网站转化量 （ WEB_CONVERSIONS ）的广告中使用含非直播活动或直播活动提醒的倒计时贴纸（即 sticker_type 为 REMINDER_COUNTDOWN 或 LIVE_REMINDER_COUNTDOWN ）目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 礼品码贴纸仅可用于以覆盖人数（REACH），访问量 （TRAFFIC）， 视频播放量（VIDEO_VIEW），应用推广 （APP_PROMOTION），线索收集（LEAD_GENERATION），网站转化量 （WEB_CONVERSIONS）或商品销量（PRODUCT_SALES） 为推广目标（objective_type）的竞价广告，以及以覆盖人数（RF_REACH）为推广目标（objective_type）的覆盖和频次广告。
	// 交互手势仅可用于以覆盖人数（REACH）、访问量（TRAFFIC）、视频播放量（VIDEO_VIEW）（optimization_goal 需设置为 ENGAGED_VIEW）、应用推广（APP_PROMOTION）或网站转化量（WEB_CONVERSIONS）为推广目标（objective_type）的竞价广告，以及以覆盖人数（RF_REACH）为推广目标（objective_type）的覆盖和频次广告。
	// 在推广目标为应用推广（APP_PROMOTION）、访问量（TRAFFIC）或网站转化量（WEB_CONVERSIONS）的推广系列广告中使用交互手势目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 点赞彩蛋仅可用于以覆盖人数（REACH）、访问量（TRAFFIC）、视频播放量（VIDEO_VIEW）（optimization_goal 需设置为 ENGAGED_VIEW）、社区互动（ENGAGEMENT）（optimization_goal 需设置为 PAGE_VISIT）、应用推广（APP_PROMOTION）或网站转化量（WEB_CONVERSIONS）为推广目标（objective_type）的竞价广告，以及以覆盖人数（RF_REACH）为推广目标（objective_type）的覆盖和频次广告。
	// 在推广目标为访问量（TRAFFIC）或应用推广（APP_PROMOTION）的推广系列广告中使用点赞彩蛋目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	CreativePortfolioType enum.CreativePortfolioType `json:"creative_portfolio_type,omitempty"`
	// PortfolioContent 素材包内容
	PortfolioContent *PortfolioContent `json:"portfolio_content,omitempty"`
}

// PortfolioContent 素材包内容
type PortfolioContent struct {
	// AssetContent 当creative_portfolio_type为CTA时，本字段必填。
	// 行动引导文案。例如："Learn More"
	AssetContent string `json:"asset_content,omitempty"`
	// AssetIDs 当creative_portfolio_type为CTA时，本字段必填。
	// 行动引导文案ID列表。例如，[201781, 201535]。你需要使用/creative/cta/recommend/接口获取一组自动优化的行动引导文案。然后将返回的数据传入本字段。
	// 注意：只有创建自动优化行动引导文案的广告主账号，才能使用该文案。
	AssetIDs []string `json:"asset_ids,omitempty"`
	// CardType 当素材包内容（creative_portfolio_type）为卡片（CARD）时，本字段必填。
	// 卡片类型。
	// 枚举值：IMAGE（展示卡片）
	CardType enum.CreativePortfolioCardType `json:"card_type,omitempty"`
	// GestureType creative_portfolio_type 为 GESTURE 时，本字段必填。
	// 交互手势类型。
	// 枚举值：
	// CLICK：点击型。
	// STRAIGHT_SLIDE：直线滑动型。
	// CURVED_SLIDE：曲线滑动型。
	GestureType enum.CreativePortfolioGestureType `json:"gesture_type,omitempty"`
	// ImageID creative_portfolio_type 为 CARD 、WEB_INFO_CARD 或 SUPER_LIKE 时，本字段必填。
	// 图片 ID。
	// 对于展示卡片（CARD），图片尺寸必须为 750（宽）x 421（高）像素。
	// 对于网站信息卡片（WEB_INFO_CARD），推荐图片尺寸为 800（宽）x 800（高）像素。
	// 对于点赞彩蛋（SUPER_LIKE），本字段代表受众点赞或点按两次广告时，屏幕上会与心形图标一起飘过的自定义图标的图片。该图片需满足以下规格要求：
	// 推荐图片尺寸：200（宽）x 200（高）像素。
	// 文件类型：JPG、JPEG 或 PNG。
	// 文件大小：不超过 50KB。
	// 您可通过/file/image/ad/upload/或/file/image/ad/search/获取图片 ID。
	ImageID string `json:"image_id,omitempty"`
	// PopUpWindowImageID creative_portfolio_type 为 GESTURE 时，本字段必填。
	// creative_portfolio_type 为 SUPER_LIKE 时，本字段为选填字段。
	// 交互手势中跳转图的图片 ID 或点赞彩蛋中弹出窗口的图片 ID。
	// creative_portfolio_type 为 GESTURE 时，本字段代表受众跟随手势后将出现的跳转图。
	// creative_portfolio_type 为 SUPER_LIKE 时，本字段代表在点赞彩蛋播放结束后出现的弹出窗口，当受众点击弹窗，就会跳转到您的广告落地页。
	// 支持的规格：
	// 推荐图片尺寸：620（宽）x 788（高）像素。
	// 文件类型：JPG、JPEG、PNG 或 GIF。
	// 文件大小：不超过 3MB。
	// 若想上传图片获取图片 ID，可使用 /file/image/ad/upload/。
	PopUpWindowImageID string `json:"pop_up_window_image_id,omitempty"`
	// Title creative_portfolio_type 为 WEB_INFO_CARD 时，本字段必填。
	// 卡片标题。
	// 长度限制：25 个字符。
	Title string `json:"title,omitempty"`
	// SellingPoints creative_portfolio_type 为 WEB_INFO_CARD 时，本字段必填。
	// 卖点，即用户应该购买您的商品的理由。
	// 最大数量：5。
	// 本字符串数组中的每个字符串代表一个卖点，单个卖点的长度限制：25 个字符。
	SellingPoints []string `json:"selling_points,omitempty"`
	// CallToActionText creative_portfolio_type 为 GESTURE 或 PREMIUM_BADGE 时，本字段必填。
	// 交互手势或弹出展示中的交互引导语。
	// creative_portfolio_type 为 GESTURE 时，本字段代表交互手势中的交互引导语。
	// 不同交互手势类型的示例文案：
	// 点击型交互手势：Click here for more details
	// 直线滑动型交互手势：Slide in a straight line
	// 曲线滑动型交互手势：Slide along the curve
	// creative_portfolio_type 为 PREMIUM_BADGE 时，本字段代表弹出展示中的交互引导语。
	// 示例文案： Click here for details
	// 长度限制：27 个字符。每个英文字母占用 1 个字符，中文、日文或韩文每个字占用 2 个字符。
	// 不支持表情符号和特殊字符 {、}、@ 和 # 。
	// 请勿包含条形码、商品的库存代码或任何社交媒体信息。
	CallToActionText string `json:"call_to_action_text,omitempty"`
	// BadgeShowTime creative_portfolio_type 为 GESTURE 或 PREMIUM_BADGE 时，本字段必填。
	// 交互手势中手势图标或弹出展示中福袋的出现时机，单位为秒。
	// 视频前 3 秒和后 5 秒不可设置为起始点。
	// 示例：5。
	BadgeShowTime int `json:"badge_show_time,omitempty"`
	// BadgePosition creative_portfolio_type 为 GESTURE 或 PREMIUM_BADGE 时，本字段必填。
	// 交互手势中手势图标或弹出展示中福袋的位置信息。
	BadgePosition *BadgePosition `json:"badge_position,omitempty"`
	// BadgeImageInfo creative_portfolio_type 为 PREMIUM_BADGE 时，本字段必填。
	// 弹出展示中福袋图片信息。
	BadgeImageInfo *BadgeImageInfo `json:"badge_image_info,omitempty"`
	// SlideLength creative_portfolio_type 为 GESTURE 且 gesture_type 为 STRAIGHT_SLIDE 时必填。
	// 直线滑动型手势的长度。
	SlideLength int `json:"slide_length,omitempty"`
	// SlideDimension creative_portfolio_type 为 GESTURE 且gesture_type为 CURVED_SLIDE 时必填。
	// 曲线滑动型手势的尺寸。
	SlideDimension *SlideDimension `json:"slide_dimension,omitempty"`
	// InteractiveMusicID creative_portfolio_type 为 GESTURE 时必填。
	// 交互手势中手势音效的 URL。
	// 规格要求：
	// 文件大小：不超过 1 MB。
	// 文件格式：MP3。
	// 推荐时长：6 秒。
	// 示例：https://sf16-sg-default.akamaized.net/obj/tiktok-obj/0123456789.mp3。
	// 为确保流畅的广告体验，可将视频广告的背景音乐设置为手势音效。
	// 若想将音乐上传至素材库并获取音乐 URL（url），可使用/file/music/upload/。
	InteractiveMusicID string `json:"interactive_music_id,omitempty"`
	// Layouts 当creative_portfolio_type为DOWNLOAD_CARD时必填。
	// 下载卡的布局类型。枚举值：
	// TYPE_1：类型1，在下载卡中展示应用描述。
	// TYPE_2：类型2，在下载卡中不展示应用描述。
	// 您可传入一种或两种类型，当您同时传入 TYPE_1 和 TYPE_2时，将同时创建一个类型1和一个类型2的下载卡。
	Layouts []enum.CreativePortfolioLayout `json:"layouts,omitempty"`
	// Description 当creative_portfolio_type为DOWNLOAD_CARD且layouts 的传入值包含 TYPE_1时必填。
	// 下载卡中展示的您想推广的应用的描述。
	Description string `json:"description,omitempty"`
	// Tags 当creative_portfolio_type为DOWNLOAD_CARD时必填。
	// 下载卡中的标签。枚举值：
	// CATEGORIES：在下载卡中展示您想推广的应用的分类。
	// FILESIZE：在下载卡中展示您想推广的应用的文件大小。
	// RATING：在下载卡中展示您想推广的应用的评分。
	// RANKING：在下载卡中展示您想推广的应用的排行。
	// COMMENT_VOLUME：在下载卡中展示您想推广的应用的评论数。
	// 注意:
	// 您可同时传入多个标签，当传入标签大于两个时，仅按以下优先级显示两个标签：CATEGORIES > FILESIZE > RATING > RANKING>COMMENT_VOLUME。
	// 对于类型1 的下载卡，您需至少传入一个标签。
	// 对于类型2 的下载卡，您需至少传入两个标签。
	Tags []enum.CreativePortfolioTag `json:"tags,omitempty"`
	// CategoryLabel 当creative_portfolio_type为DOWNLOAD_CARD且tags的传入值包含CATEGORIES时必填。
	// 下载卡中展示的您想推广的应用的分类标签。枚举值请查看枚举值-创意管理-分类标签，您需将取值以字串传入
	CategoryLabel string `json:"category_label,omitempty"`
	// AppID 当creative_portfolio_type为DOWNLOAD_CARD时必填。
	// 下载卡中您想推广的应用的ID。该ID为事件管理平台中显示的应用ID
	AppID string `json:"app_id,omitempty"`
	// ProfileImage 当creative_portfolio_type为DOWNLOAD_CARD时必填。
	// 头像URL。
	ProfileImage string `json:"profile_image,omitempty"`
	// CallToAction 当creative_portfolio_type为DOWNLOAD_CARD时必填。
	// 下载卡中的行动引导文案
	CallToAction enum.CallToAction `json:"call_to_action,omitempty"`
	// MobileAppID 当creative_portfolio_type为DOWNLOAD_CARD时必填。
	// 移动应用ID，用于从 Google Play 或App Store中提取标签信息。
	// 对于iOS应用，请传入应用对应的App Store URL中id后的数字串作为ID。例如https://apps.apple.com/us/app/hyperpure/id1203646221中移动应用ID为1203646221。
	// 对于Android应用，请传入应用对应的Google Play 商店 URL中id后的应用程序包名作为ID。例如https://play.google.com/store/apps/details?id=com.innersloth.spacemafia 中移动应用ID为com.innersloth.spacemafia。
	MobileAppID string `json:"mobile_app_id,omitempty"`
	// CountryCode 当creative_portfolio_type为DOWNLOAD_CARD时必填。
	// 想要定位的地域的国家或地区代码。枚举值可参见 附录-地区代码。
	CountryCode string `json:"country_code,omitempty"`
	// StickerParam 当creative_portfolio_type为STICKER时，该字段为必填。
	// 想要创建的倒计时贴纸和礼品码贴纸详情。倒计时贴纸和礼品码贴纸示例可查看倒计时贴纸和礼品码贴纸。
	// 注意：单次仅可创建一个倒计时贴纸或礼品码贴纸。
	StickerParam *StickerParam `json:"sticker_param,omitempty"`
	// ProductSource creative_portfolio_type为PRODUCT_CARD或 PRODUCT_TILE时必填。
	// creative_portfolio_type为 PRODUCT_TILE 时，本字段仅支持设置为CATALOG 。
	// 商品来源，即您要推广的商品的来源。
	// 枚举值：UNSET， CATALOG （商品库），STORE（TikTok Shop 店铺），SHOWCASE （TikTok 橱窗）。
	// 若您将本字段设置为STORE，您需同时传入item_group_ids。
	// 注意：若您创建完商品卡片后希望在创建广告时使用商品卡片，需确保在广告组层级本字段传入创建商品卡片时所用的值。否则，广告实际投放时，商品卡片中的商品缩略图可能与所推广的商品不符。
	ProductSource enum.ProductSource `json:"product_source,omitempty"`
	// IdentityID creative_portfolio_type为PRODUCT_CARD或 PRODUCT_TILE时必填。
	// 认证身份ID。
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType creative_portfolio_type为PRODUCT_CARD或 PRODUCT_TILE时必填。
	// creative_portfolio_type为PRODUCT_CARD且product_source为STORE时，本字段支持的值为AUTH_CODE，TT_USER和BC_AUTH_TT。
	// 认证身份类型。
	// 枚举值: CUSTOMIZED_USER（自定义用户）, AUTH_CODE（授权用户）, TT_USER（TikTok用户），BC_AUTH_TT（已添加到商务中心的TikTok用户）。
	// 关于认证身份的更多信息，请查看认证身份。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID creative_portfolio_type为PRODUCT_CARD或 PRODUCT_TILE且identity_type为BC_AUTH_TT时必填。
	// 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// CatalogID creative_portfolio_type为PRODUCT_CARD且product_source为CATALOG，或creative_portfolio_type为PRODUCT_TILE 时必填。
	// 商品库ID。
	// 注意：
	// 您可使用/store/list/获取广告账户下可用商店对应的catalog_id，store_id和store_authorized_bc_id。
	// 若您创建完商品卡片后希望在创建广告时使用商品卡片，需确保在广告组层级本字段传入创建商品卡片时所用的值。否则，广告实际投放时，商品卡片中的商品缩略图可能与所推广的商品不符。
	CatalogID string `json:"catalog_id,omitempty"`
	// CatalogAuthorizedBcID creative_portfolio_type为PRODUCT_CARD且product_source为CATALOG，或creative_portfolio_type为PRODUCT_TILE时必填。
	// 拥有商品库（catalog_id）权限的商务中心的ID。
	// 注意：若您创建完商品卡片后希望在创建广告时使用商品卡片，需确保在广告组层级本字段传入创建商品卡片时所用的值。否则，广告实际投放时，商品卡片中的商品缩略图可能与所推广的商品不符。
	CatalogAuthorizedBcID string `json:"catalog_authorized_bc_id,omitempty"`
	// StoreID creative_portfolio_type为PRODUCT_CARD且product_source为STORE时必填。
	// creative_portfolio_type为PRODUCT_TILE 时不支持本字段。
	// TikTok Shop 店铺 ID。
	// 注意：
	// 欲获取 TikTok Shop 店铺 ID，您可以调用/bc/asset/get/：
	// 当返回中的asset_type为TIKTOK_SHOP时，对应的asset_id为TikTok Shop店铺ID。
	// 若您创建完商品卡片后希望在创建广告时使用商品卡片，需确保在广告组层级本字段传入创建商品卡片时所用的值。否则，广告实际投放时，商品卡片中的商品缩略图可能与所推广的商品不符。
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID creative_portfolio_type为PRODUCT_CARD且product_source为STORE时必填。
	// creative_portfolio_type为PRODUCT_TILE 时不支持本字段。
	// 有该店铺 (store_id）权限的商务中心的ID。
	// 注意：若您创建完商品卡片后希望在创建广告时使用商品卡片，需确保在广告组层级本字段传入创建商品卡片时所用的值。否则，广告实际投放时，商品卡片中的商品缩略图可能与所推广的商品不符。
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// ProductSpecificType creative_portfolio_type为PRODUCT_CARD且product_source为CATALOG，或creative_portfolio_type为PRODUCT_TILE时必填。
	// 指定商品的方式。 枚举值：
	// ALL：所有商品。 允许TikTok动态选择商品库中的所有商品。您此时不需同时传入sku_ids，item_group_ids 或product_set_id 。
	// PRODUCT_SET：商品系列。请选择一个商品系列，TikTok将动态选择该商品系列中的商品。您此时需同时传入item_group_ids或product_set_id。
	// CUSTOMIZED_PRODUCTS：特定商品。最多可从您的商品库中选择20个商品。您需同时传入sku_ids 。
	// 注意：若您创建完商品卡片后希望在创建广告时使用商品卡片，需确保在广告层级本字段传入创建商品卡片时所用的值。否则，广告实际投放时，商品卡片中的商品缩略图可能与所推广的商品不符。
	ProductSpecificType enum.ProductSpecificType `json:"product_specific_type,omitempty"`
	// ProductSetID creative_portfolio_type为PRODUCT_CARD或 PRODUCT_TILE且product_specific_type为PRODUCT_SET时，您需传入product_set_id或item_group_ids。
	// 商品系列ID。您可以使用/catalog/set/get/获取商务中心下商品库的商品系列列表。
	// 注意：若您创建完商品卡片后希望在创建广告时使用商品卡片，需确保在广告层级本字段传入创建商品卡片时所用的值。否则，广告实际投放时，商品卡片中的商品缩略图可能与所推广的商品不符。
	ProductSetID string `json:"product_set_id,omitempty"`
	// ItemGroupIDs creative_portfolio_type为PRODUCT_CARD或 PRODUCT_TILE且product_source为STORE时必填。
	// creative_portfolio_type为PRODUCT_CARD且product_specific_type为PRODUCT_SET时，您需传入product_set_id或item_group_ids。
	// 商品 SPU ID列表。最大数量： 20。
	// 注意：若您创建完商品卡片后希望在创建广告时使用商品卡片，需确保在广告层级本字段传入创建商品卡片时所用的值。否则，广告实际投放时，商品卡片中的商品缩略图可能与所推广的商品不符。
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
	// SkuIDs creative_portfolio_type为PRODUCT_CARD 或 PRODUCT_TILE且product_specific_type为CUSTOMIZED_PRODUCTS时必填。
	// 商品 SKU ID列表。最大数量： 20。
	// 注意：若您创建完商品卡片后希望在创建广告时使用商品卡片，需确保在广告层级本字段传入创建商品卡片时所用的值。否则，广告实际投放时，商品卡片中的商品缩略图可能与所推广的商品不符。
	SkuIDs []string `json:"sku_ids,omitempty"`
	// DynamicFormat creative_portfolio_type为PRODUCT_CARD且product_source为CATALOG时必填。
	// creative_portfolio_type为PRODUCT_TILE 时不支持本字段。
	// creative_portfolio_type为PRODUCT_CARD且product_source为STORE时，则此字段不可设置为DYNAMIC_CREATIVE，您可不传入此字段，或将本字段设置为UNSET。
	// 是否使用动态样式。动态样式功能将视频素材、商品卡片以及落地页根据顾客的购买意图进行不同的组合，以实现转化最大化。
	// 枚举值： UNSET，DYNAMIC_CREATIVE（使用动态样式生成创意）。
	// 注意：
	// 受众再营销广告（广告组的shopping_ads_retargeting_type不是OFF）无法使用dynamic_format。当 dynamic_format 设置为DYNAMIC_CREATIVE时，会自动创建卡片。这意味着您无需传入与卡片相关的字段（card_show_price， card_tags和 card_image_index ）。
	// 若您创建完商品卡片后希望在创建广告时使用商品卡片，需确保在广告层级本字段传入创建商品卡片时所用的值。
	DynamicFormat enum.DynamicFormat `json:"dynamic_format,omitempty"`
	// VerticalVideoStrategy creative_portfolio_type为PRODUCT_CARD且product_source为CATALOG时必填。
	// creative_portfolio_type为PRODUCT_TILE 时不支持本字段。
	// creative_portfolio_type为PRODUCT_CARD且product_source为STORE时，您可不传入此字段，或将本字段设置为SINGLE_VIDEO。
	// 商品销量（即购物广告）场景中使用的视频类型。 枚举值： UNSET (未设置), SINGLE_VIDEO (单个视频），CATALOG_VIDEOS（商品库视频）。
	// 如果dynamic_format为DYNAMIC_CREATIVE，该字段必须设置为UNSET。
	// 注意：若您创建完商品卡片后希望在创建广告时使用商品卡片，需确保在广告层级本字段传入创建商品卡片时所用的值。
	VerticalVideoStrategy enum.VerticalVideoStrategy `json:"vertical_video_strategy,omitempty"`
	// AdText creative_portfolio_type为PRODUCT_CARD时生效。
	// 广告文案，将作为广告创意的一部分展示给你的受众，向他们传达你想要推广的信息。若您不清楚如何创建有效文案，可以使用智能文案功能，根据行业和语言获取文案推荐。
	// 注意：
	// 广告文案长度允许值为1-100 字符，不能包含emoji。
	// 中文或日文每个字占用2个字符，每个英语字符占用1个字符。
	AdText string `json:"ad_text,omitempty"`
	// CardShowPrice creative_portfolio_type为PRODUCT_CARD时必填。
	// 是否在商品卡片中显示商品价格
	CardShowPrice *bool `json:"card_show_price,omitempty"`
	// CardTags reative_portfolio_type为PRODUCT_CARD时必填。
	// 想要在商品卡片中显示的商品标签。 枚举值：BRAND （品牌名称）， DESC（说明）
	CardTags []string `json:"card_tags,omitempty"`
	// CardImageIndex creative_portfolio_type为PRODUCT_CARD或 PRODUCT_TILE时生效。
	// 您可使用本字段指定在商品卡片或商品磁贴中使用的图片。您通过本字段设置的数字将作为索引，查询您调用/catalog/product/upload/时通过additional_image_urls字段传入的图片URL。
	// 例如，card_image_index为1代表您将使用additional_image_urls的值中的第二图片URL对应的图片。
	CardImageIndex int `json:"card_image_index,omitempty"`
	// ShowcaseProducts creative_portfolio_type为PRODUCT_CARD且product_source为SHOWCASE时必填。您可查看创建商品来源为橱窗的视频购物广告了解详情。
	// 您想要用在广告中的橱窗商品列表。最大数量：20。
	// 您可调用/showcase/product/get/，获取可用的橱窗商品。
	ShowcaseProducts []showcase.Product `json:"showcase_products,omitempty"`
}

// BadgePosition 交互手势中手势图标或弹出展示中福袋的位置信息。
type BadgePosition struct {
	// PositionX 传入 badge_position 时必填。
	// 交互手势中手势图标或弹出展示中福袋左上角的相对横坐标，以左上角横坐标除以 720 表示。
	// 取值范围：0-1。
	PositionX float64 `json:"position_x,omitempty"`
	// PositionY 传入 badge_position 时必填。
	// 交互手势中手势图标或弹出展示中福袋左上角的相对纵坐标，以左上角横坐标除以 1280 表示。
	// 取值范围：0-1。
	PositionY float64 `json:"position_y,omitempty"`
	// Angle 传入 badge_position 且满足以下任一条件时，本字段必填。
	// creative_portfolio_type 为 GESTURE 且 gesture_type 为 STRAIGHT_SLIDE 或 CURVED_SLIDE。
	// creative_portfolio_type 为 PREMIUM_BADGE。
	// 交互手势中直线滑动型或曲线滑动型手势图标的手势角度或弹出展示中福袋的弹出角度。
	// 取值范围：
	// creative_portfolio_type 为 GESTURE 且 gesture_type 为 STRAIGHT_SLIDE 或 CURVED_SLIDE 时，推荐角度为 30-150 或 210-330 区间内的值。
	// creative_portfolio_type 为 PREMIUM_BADGE 时：0-360。
	Angle int `json:"angle,omitempty"`
}

// BadgeImageInfo 弹出展示中福袋图片信息。
type BadgeImageInfo struct {
	// ImageID 传入 badge_image_info 时必填。
	// 弹出展示中福袋的图片 ID。
	// 规格要求：
	// 尺寸：448（宽）x 448（高）像素。
	// 图片格式：JPG、JPEG 或 PNG。
	// 文件大小：不超过 1 MB。
	ImageID string `json:"image_id,omitempty"`
}

// SlideDimension 曲线滑动型手势的尺寸。
type SlideDimension struct {
	// DimensionWidth 传入 slide_dimension 时必填。
	// 曲线滑动型手势的宽度
	DimensionWidth int `json:"dimension_width,omitempty"`
	// DimensionHeight 传入 slide_dimension 时必填。
	// 曲线滑动型手势的高度。
	DimensionHeight int `json:"dimension_height,omitempty"`
}

// StickerParam 想要创建的倒计时贴纸和礼品码贴纸详情。倒计时贴纸和礼品码贴纸示例可查看倒计时贴纸和礼品码贴纸。
type StickerParam struct {
	// StickerType 贴纸类型。默认值： COUNTDOWN。 枚举值：
	// COUNTDOWN：无提醒的倒计时贴纸。
	// REMINDER_COUNTDOWN： 含非直播活动提醒的倒计时贴纸。
	// LIVE_REMINDER_COUNTDOWN：含直播活动提醒的倒计时贴纸。
	// GIFTCODE：礼品码贴纸，用于通过提供虚拟礼品或优惠券兑换码来吸引新客户和促进现有客户。
	StickerType enum.CreativePortfolioStickerType `json:"sticker_type,omitempty"`
	// Title 当creative_portfolio_type为STICKER时，该字段为必填。
	// 贴纸标题。
	// 对于倒计时贴纸，最大长度为54个UTF-8字节，且不支持话题标签符号 (#)。
	// 对于礼品码贴纸，最大长度为20个UTF-8字节。
	Title string `json:"title,omitempty"`
	// GiftCode 当sticker_type为GIFTCODE时，该字段为必填。
	// 当sticker_type不为GIFTCODE时，不支持传入该字段。
	// 礼品码。
	// 本字段仅适用于礼品码贴纸。最大长度为10个UTF-8字节（10个英文字母和数字）。
	GiftCode string `json:"giftcode,omitempty"`
	// CutoffTime 当sticker_type为COUNTDOWN，REMINDER_COUNTDOWN，或 LIVE_REMINDER_COUNTDOWN时，该字段为必填。
	// 当sticker_type为GIFTCODE时，不支持传入该字段。
	// 贴纸对应的倒计时截止日期(sticker_type为COUNTDOWN或REMINDER_COUNTDOWN时)或直播开始时间 (sticker_type为LIVE_REMINDER_COUNTDOWN时)，格式为"2022-10-30 00:00:00" (UTC+0 时间)。
	// 注意：您在此字段传入的时间将默认为UTC+0时间，且该默认时区无法更改。
	CutoffTime model.DateTime `json:"cutoff_time,omitzero"`
	// Color 当creative_portfolio_type为STICKER时，该字段为必填。
	// 贴纸的背景色。
	// 对于倒计时贴纸，支持的枚举值：ORANGE(橘色)、BLACK(黑色)、RED(红色)、BLUE(蓝色)。
	// 对于礼品码贴纸，支持的枚举值： BLACK(黑色)和BLUE(蓝色)
	Color string `json:"color,omitempty"`
	// DisplayAngle 当sticker_type为COUNTDOWN，REMINDER_COUNTDOWN，或LIVE_REMINDER_COUNTDOWN时，该字段为必填。
	// 当sticker_type为GIFTCODE时不支持传入。贴纸的显示角度。
	// 取值范围： [-180, +180]。 +90 表示将贴纸顺时针旋转90度。
	DisplayAngle int `json:"display_angle,omitempty"`
	// PredefinedPlacement 当sticker_type为COUNTDOWN，REMINDER_COUNTDOWN或LIVE_REMINDER_COUNTDOWN时，需要传入predefined_placement字段，或者传入position_x，position_y和size字段。
	// 当sticker_type为GIFTCODE时，需要传入predefined_placement，或者传入position_x，position_y字段。
	// 该字段用于预定义：
	// 礼品码贴纸的横坐标和纵坐标。
	// 或倒计时贴纸的横坐标和纵坐标以及大小。
	// 例如，若将predefined_placement设置为55x143x0.9，对于倒计时贴纸，x轴坐标为55，y轴坐标为143，大小为90%。
	// 同样，若将predefined_placement设置为86x179，那么对于礼品码贴纸，x轴坐标为86，y轴坐标为179。
	// 本字段的枚举值，请参见“预定义位置”章节。
	// 注意：
	// 对于倒计时贴纸，若指定了predefined_placement，同时也指定了position_x，position_y和size，position_x，position_y和size的值将被忽略。
	// 对于礼品码贴纸，若指定了predefined_placement，同时也指定了position_x，position_y，那么position_x和position_y的值将被忽略。
	PredefinedPlacement string `json:"predefined_placement,omitempty"`
	// PositionX 当sticker_type为COUNTDOWN，REMINDER_COUNTDOWN或LIVE_REMINDER_COUNTDOWN时，需要传入predefined_placement字段，或者传入position_x，position_y和size字段。
	// 当sticker_type为GIFTCODE时，需要传入predefined_placement，或者传入position_x，position_y字段。
	// 贴纸左上角相对于屏幕左上角的横坐标。
	// 对于倒计时贴纸：
	// 本字段的取值范围随贴纸大小变动：
	// size设置为0.6时，本字段取值范围为[50, 301]。
	// size设置为0.7时，本字段取值范围为[50, 253]。
	// size设置为0.8时，本字段取值范围为[50, 205]。
	// size设置为0.9时，本字段取值范围为[50, 157]。
	// size设置为1时，本字段取值范围为[50, 109]。
	// size设置为1.1时，本字段取值范围为[50, 61]。
	// 对于礼品码贴纸：本字段取值范围[79, 145]。
	// 注意：对于倒计时贴纸和礼品码贴纸，若同时传入predefined_placement和position_x，那么position_x将被忽略
	PositionX int `json:"position_x,omitempty"`
	// PositionY 当sticker_type为COUNTDOWN，REMINDER_COUNTDOWN或LIVE_REMINDER_COUNTDOWN时，需要传入predefined_placement字段，或者传入position_x，position_y和size字段。
	// 当sticker_type为GIFTCODE时，需要传入predefined_placement，或者传入position_x，position_y字段。
	// 贴纸左上角相对于屏幕左上角的纵坐标。
	// 对于倒计时贴纸：
	// 本字段的取值范围随贴纸大小变动：
	// size设置为0.6时，本字段取值范围为[141, 639]。
	// size设置为0.7时，本字段取值范围为[141, 606]。
	// size设置为0.8时，本字段取值范围为[141, 573]。
	// size设置为0.9时，本字段取值范围为[141, 540]。
	// size设置为1时，本字段取值范围为[141, 506]。
	// size设置为1.1时，本字段取值范围为[141, 473]。
	// 对于礼品码贴纸：本字段取值范围[169, 683]。
	// 注意：对于倒计时贴纸和礼品码贴纸，若同时传入predefined_placement和position_x，那么position_x将被忽略。
	PositionY int `json:"position_y,omitempty"`
	// Size 当sticker_type为COUNTDOWN，REMINDER_COUNTDOWN或LIVE_REMINDER_COUNTDOWN时，需要传入predefined_placement字段，或者传入position_x，position_y和size字段。
	// 当sticker_type为GIFTCODE时，不支持传入size字段。
	// 以百分比表示的贴纸大小。枚举值：0.6(60%), 0.7(70%), 0.8(80%), 0.9(90%), 1(100%), 1.1(110%)。
	// 注意：对于倒计时贴纸，如果同时传入predefined_placement和size，那么size将被忽略。
	Size string `json:"size,omitempty"`
	// ReminderTime 设置的提醒时间。 当sticker_type为REMINDER_COUNTDOWN或LIVE_REMINDER_COUNTDOWN时必填。
	// 对于非直播倒计时贴纸 (即sticker_type为REMINDER_COUNTDOWN)，枚举值为：
	// ONE_MINUTE_EARLIER: 在非直播活动前一分钟发送提醒。
	// ONE_HOUR_EARLIER : 在非直播活动前一小时发送提醒。
	// ONE_DAY_EARLIER: 在非直播活动前一天发送提醒。
	// 对于非直播倒计时贴纸 (即sticker_type为LIVE_REMINDER_COUNTDOWN), 枚举值为：
	// ONE_MINUTE_AFTER: 在直播活动开始一分钟后发送提醒。
	// FIVE_MINUTES_AFTER: 在直播活动开始五分钟后发送提醒。
	// TEN_MINUTES_AFTER: 在直播活动开始十分钟后发送提醒。
	ReminderTime string `json:"reminder_time,omitempty"`
	// LandingPageURL 落地页URL，在非直播活动的倒计时结束后，点击提醒时会跳转至此URL。 仅当sticker_type为REMINDER_COUNTDOWN时必填
	LandingPageURL string `json:"landing_page_url,omitempty"`
	// LiveTikTokUserID 直播活动主播的TikTok用户ID 。仅当 sticker_type为LIVE_REMINDER_COUNTDOWN时必填。 传入本字段后，当您点击直播活动的提醒时，会自动跳转至主播的直播间。
	LiveTikTokUserID string `json:"live_tiktok_user_id,omitempty"`
}
