package adgroup

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建广告组 API Request
type CreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// AdgroupName 广告组名，长度小于 512 字符，不能包含emoji。
	AdgroupName string `json:"adgroup_name,omitempty"`
	// RequestID 请求ID，用于创建同名广告组。该ID支持接口幂等，避免重复请求。如果您传入相同的request ID重试多次请求，只有一次请求会成功。
	// 该ID和返回参数中的 request_id 不同。返回的request_id用于唯一标识一次 HTTP 请求。
	// 本字段的值需为字符串格式的64位整数值。
	RequestID string `json:"request_id,omitempty"`
	// ShoppingAdsType 当对应的推广系列 objective_type 为 PRODUCT_SALES时必填。
	// 购物广告类型。
	//
	// 枚举值：
	// VIDEO：视频购物广告。
	// LIVE：直播购物广告。
	// PRODUCT_SHOPPING_ADS：商品购物广告。
	ShoppingAdsType enum.ShoppingAdsType `json:"shopping_ads_type,omitempty"`
	// IdentityID 仅当 shopping_ads_type 为 VIDEO 且 product_source 为 SHOWCASE 或 shopping_ads_type 为 LIVE 时生效且必填
	// 认证身份 ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 仅当 shopping_ads_type 为 VIDEO 且 product_source 为 SHOWCASE 或 shopping_ads_type 为 LIVE 时生效且必填
	// 认证身份类型。枚举值: AUTH_CODE (授权码认证)， TT_USER (TikTok商务用户)， BC_AUTH_TT（已添加到商务中心的TikTok用户）。参阅认证身份 获取详细信息。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。当identity_type 为 BC_AUTH_TT时必填。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// ProductSource 仅视频购物广告和商品购物广告必填。
	//
	// 商品来源，即您要推广的商品的来源。
	//
	// 枚举值：UNSET ,CATALOG(商品库)，STORE (TikTok Shop)，SHOWCASE（TikTok 橱窗）。
	// 对于视频购物广告：
	// 通过campaign_product_source设置的推广系列层级商品来源决定了广告组层级可选的商品来源：
	// 若在推广系列层级campaign_product_source设置为CATALOG，本字段仅可设置为CATALOG。
	// 若在推广系列层级campaign_product_source设置为 STORE，本字段仅可设置为STORE或SHOWCASE。
	// 若未在推广系列层级设置campaign_product_source，且同一推广系列下存在至少一个product_source设置为CATALOG 的历史广告组，和一个product_source设置为STORE 或SHOWCASE的历史广告组，则该推广系列下不允许新建广告组。
	// 同一推广系列下的广告组的商品来源需均设置为CATALOG，或为STORE和SHOWCASE的组合（包括均为STORE，或均为SHOWCASE，或为STORE和SHOWCASE的混合） 。
	// 对于商品购物广告，需将本字段设置为STORE。
	ProductSource enum.ProductSource `json:"product_source,omitempty"`
	// CatalogID 商品库ID
	// - 若 product_source 设置为 CATALOG 或 STORE，本字段必填。
	// - 若 shopping_ads_type 设置为 LIVE，则字段 catalog_id，catalog_authorized_bc_id，store_id 和 store_authorized_bc_id 为选填字段。若您通过这些字段为您的直播购物广告指定了TikTok Shop，将会报告此商店的商品销量。若您不指定 TikTok Shop，将会报告 TikTok 账号（identity_id）显示的商品销量。
	// - 若 product_source 设置为 CATALOG，需将本字段设置为您想推广的商品所在的商品库。
	// - 若 product_source 设置为 STORE，需将本字段设置为与您想推广的商品所在的TikTok Shop 绑定的商品库。
	// 您可使用/store/list/获取广告账户下可用商店对应的 catalog_id，store_id 和 store_authorized_bc_id。
	// 注意：自 2024 年 6 月 30 日起， product_source 为 STORE 时，您将无需再传入 catalog_id，因为该字段将被忽略
	CatalogID string `json:"catalog_id,omitempty"`
	// CatalogAuthorizedBcID 对该商品库有权限的商务中心的 ID
	// - 若 shopping_ads_type 设置为 VIDEO 且 product_source 设置为CATALOG 或 STORE，本字段必填。
	// - 若 shopping_ads_type 设置为 LIVE，本字段选填。
	// - 若 shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS 且 product_source 设置为 STORE，不支持本字段。
	// 如果要操作的是已经资产化的商品库，则需要传入相应的商品库 ID。
	// 注意：自 2024 年 6 月 30 日起， product_source 为 STORE 时，您将无需再传入 catalog_id，因为该字段将被忽略。
	CatalogAuthorizedBcID string `json:"catalog_authorized_bc_id,omitempty"`
	// StoreID TikTok Shop 店铺 ID
	// - 在以下任意场景必填：
	//   - shopping_ads_type 设置为 VIDEO 且 product_source 设置为 STORE。
	//   - shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS 且 product_source 设置为 STORE。
	//   - shopping_ads_type 设置为 LIVE 时，本字段选填。
	// 欲获取 TikTok Shop 店铺 ID，您可以调用/bc/asset/get/：当返回中的asset_type为 TIKTOK_SHOP时，对应的asset_id为 TikTok Shop 店铺 ID。
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID 传入 store_id 时必填
	// 有权管理该商店 (store_id）的商务中心的 ID。
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// PromotionType 推广对象类型（优化位置）。若对应的推广系列推广目标不是REACH，VIDEO_VIEWS或ENGAGEMENT，则此字段必填。
	// 枚举值详见 枚举值-推广对象类型。
	// 注意: 如果在推广系列层级的objective_type是 ENGAGEMENT，则该字段只能使用 EXTERNAL_OR_DISPLAY。
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// PromotionTargetType 推广目标线索收集的专属推广对象类型
	// 仅当推广系列层级的objective_type 设置为 LEAD_GENERATION时有效。
	// 枚举值：
	// - INSTANT_PAGE：即时表单（线索表单）。创建可快速加载的应用内 TikTok 即时表单来收集更多线索。
	// - EXTERNAL_WEBSITE：第三方网站表单。使用包含第三方网站表单的落地页，或使用会重定向至包含第三方网站表单网站的TikTok即时体验页面来收集更多线索。
	// 若本字段设置为EXTERNAL_WEBSITE且 optimization_goal 为 CLICK，无需同时传入pixel_id 和 optimization_event
	PromotionTargetType enum.PromotionTargetType `json:"promotion_target_type,omitempty"`
	// MessageAppType 即时通讯广告组中使用的即时通讯应用或自定义 URL。
	// 枚举值：
	// MESSENGER：Messenger。您需同时通过 messaging_app_account_id 指定 Facebook 公共主页。
	// WHATSAPP：WhatsApp。您需同时指定 phone_region_code、phone_region_calling_code 和 phone_number。系统将自动使用这些字段填充对应的 WhatsApp 电话号码。
	// ZALO: Zalo。您需同时指定 phone_region_code、phone_region_calling_code 和 phone_number 来自动填充对应的 Zalo 电话号码。
	// LINE: Line。您需同时通过 messaging_app_account_id 指定 LINE Business ID。
	// IM_URL：即时通讯 URL。
	//
	// 默认值：IM_URL。
	// 仅当 promotion_type 为 LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE 时生效。
	// optimization_goal 为 CONVERSATION 时，本字段必填。
	// optimization_goal 为 CLICK 时，本字段默认值为IM_URL。
	// 注意：
	// optimization_goal 为 CONVERSATION 时，不支持将本字段设置为 ZALO，LINE 或 IM_URL。
	// 本字段设置后不支持更新
	MesssageAppType enum.MessageAppType `json:"message_app_type,omitempty"`
	// MessageAppAccountID 即时通讯应用账号 ID。
	// - messaging_app_type 为 MESSENGER或 LINE 时，本字段必填。
	// - messaging_app_type 为 WHATSAPP, ZALO或 IM_URL 时，将忽略本字段。
	// - messaging_app_type 为 MESSENGER 时，需通过本字段指定 Facebook 公共主页编号。
	//   - 若要查找您的 Facebook 公共主页编号，请前往您的 Facebook 公共主页，点击封面照片下方的简介（About），然后点击公共主页信息公示（Page transparency）。
	// - 当 messaging_app_type 为 LINE 时，通过此字段指定 LINE Business ID。
	//   - 若要查找您的 LINE Business ID，请前往您的 LINE 官方账号页面。
	// - messaging_app_type 为 WHATSAPP 时，将忽略本字段并自动根据指定的phone_region_code、phone_region_calling_code 和 phone_number填充为对应的 WhatsApp 电话号码。
	// 注意：本字段设置后不支持更新
	MessageAppAccountID string `json:"message_app_account_id,omitempty"`
	// PhoneRegionCode WhatsApp 或 Zalo 电话号码的地区代码。
	// messaging_app_type 为 WHATSAPP 或 ZALO 时，本字段必填
	// 若想获取与某一电话号码所绑定的地区对应的地区代码（ phone_region_code ）和电话区号 （ phone_region_calling_code ），需使用/tool/phone_region_code/。
	// 注意：本字段设置后不支持更新。
	PhoneRegionCode string `json:"phone_region_code,omitempty"`
	// PhoneRegionCallingCode WhatsApp 或 Zalo 电话号码的区号。
	// messaging_app_type 为 WHATSAPP 或 ZALO 时，本字段必填
	// 若想获取与某一电话号码所绑定的地区对应的地区代码（ phone_region_code ）和电话区号 （ phone_region_calling_code ），需使用/tool/phone_region_code/。
	// 注意：本字段设置后不支持更新。
	PhoneRegionCallingCode string `json:"phone_region_calling_code,omitempty"`
	// PhoneNumber WhatsApp 或 Zalo 电话号码。
	// messaging_app_type 为 WHATSAPP 或 ZALO 时，本字段必填。
	// 对于 WhatsApp，需使用在消息管理工具（MMT）中绑定的 WhatsApp 电话号码。
	// 注意：本字段设置后不支持更新。
	PhoneNumber string `json:"phone_number,omitempty"`
	// PromotionWebsiteType 广告组中的 TikTok 即时体验页面类型。
	// 枚举值：
	// UNSET：不使用 TikTok 即时体验页面。
	// TIKTOK_NATIVE_PAGE：使用 TikTok 即时体验页面。
	// 若您在广告组层级将本字段设置为 TIKTOK_NATIVE_PAGE，则不支持传入 pixel_id，且您需要在广告层级通过 page_id 指定 TikTok 即时体验页面。
	// 若想了解如何创建应用预注册广告，请查看这里。
	// 注意：若推广系列层级app_promotion_type设置为APP_PREREGISTRATION，本字段必填。
	PromotionWebsiteType enum.PromotionWebsiteType `json:"promotion_website_type,omitempty"`
	// AppID 推广的App的ID。您可通过/app/list/获取app_id
	// 在以下使用场景中必填：
	// - objective_type 设置为 APP_PROMOTION，且 app_promotion_type设置为APP_RETARGETING（应用再营销）。
	// - objective_type 设置为 APP_PROMOTION，且 app_promotion_type设置为APP_INSTALL（应用安装）或APP_PREREGISTRATION（应用预注册），同时对应的推广系列非iOS 14专属推广系列。
	//
	// 注意: 您不能使用未在 MMP（移动监测合作伙伴）侧启用 SAN（自归因）模块的 App 创建广告组。若想确保您的 App 启用 TikTok 自归因集成，可查看如何将新应用集成到 SAN 和如何将现有应用迁移到 SAN。
	AppID string `json:"app_id,omitempty"`
	// PixelID Pixel ID，仅对落地页生效
	// 且需同时传入optimization_event字段。本字段支持的推广目标(objective_type)包括：WEB_CONVERSIONS, PRODUCT_SALES。如果您想追踪除LEAD_GENERATION外其他目标的事件，请在创意层级上使用 tracking_pixel_id 。
	// 您可以使用/pixel/list/ 接口获取所有Pixel ID。
	PixelID string `json:"pixel_id,omitempty"`
	// OptimizationEvent 广告组转化事件。参阅转化事件 ，获取详细信息
	// 当推广对象是应用且拥有追踪 URL 或传入pixel_id时必填
	// 注意：
	// 当使用 Reach，Video View，或 Traffic 推广目标，并使用pixel_id时，不支持此字段。
	// 若您要在一个已开启CBO的推广系列下创建新的广告组，需确保本字段的值与第一个广告组中的设置保持一致。
	// 当optimization_goal为 PAGE_VISIT 时，optimization_event 将自动设置为 PAGE_VISIT。
	OptimizationEvent enum.OptimizationEvent `json:"optimization_event,omitempty"`
	// DeepFunnelOptimizationStatus 深层漏斗优化的状态。
	// 仅当 promotion_type 为 LEAD_GENERATION 时本字段生效 。
	// 深层漏斗优化同时优化浅层和深层漏斗事件。您可以在主要优化事件之外再选择一个次要事件，这有助于提升推广系列的效果。
	//
	// 枚举值：
	// ON ：启用深层漏斗优化。
	// OFF：不启用深层漏斗优化。
	// 默认值：OFF。
	//
	// 若想了解如何在线索广告中配置深层漏斗优化，可查看创建优化位置为即时表单的线索广告和创建优化位置为落地页的线索广告。
	//
	// 注意：
	//
	// 使用 CRM 事件组的深层漏斗优化目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 使用 Pixel 或线下事件组的深层漏斗优化已全量发布。
	DeepFunnelOptimizationStatus enum.OnOff `json:"deep_funnel_optimization_status,omitempty"`
	// DeepFunnelEventSource 事件源类型。
	// deep_funnel_optimization_status 为 ON 时本字段必填
	// 枚举值：
	// PIXEL：Pixel。
	// OFFLINE：线下事件组。
	// CRM：CRM 事件组。
	DeepFunnelEventSource enum.DeepFunnelEventSource `json:"deep_funnel_event_source,omitempty"`
	// DeepFunnelEventSourceID 事件源 ID。
	// deep_funnel_optimization_status 为 ON 时本字段必填。
	// deep_funnel_event_source 为 PIXEL 时，本字段代表 Pixel ID。
	// deep_funnel_event_source 为 OFFLINE 时，本字段代表线下事件组 ID。
	// deep_funnel_event_source为CRM时，本字段代表 CRM 事件组 ID。
	DeepFunnelEventSourceID string `json:"deep_funnel_event_source_id,omitempty"`
	// DeepFunnelOptimizationEvent 深层漏斗优化事件。
	// deep_funnel_optimization_status 为 ON 时本字段必填。
	// 若想获取 Pixel、线下事件组和 CRM 事件组支持的标准事件值，可查看 Pixel 事件类型中的“用于广告创编的事件名称”列。
	// 若想获取某个 Pixel 的优化事件，可使用 /pixel/list/ 并查看返回的 optimization_event
	DeepFunnelOptimizationEvent enum.OptimizationEvent `json:"deep_funnel_optimization_event,omitempty"`
	// PlacementType 版位类型。
	// 枚举值：
	// PLACEMENT_TYPE_AUTOMATIC：自动版位。TikTok 的广告系统将使用智能计算，为您选择所有可用应用的版位最佳组合。
	// PLACEMENT_TYPE_NORMAL：编辑版位。通过手动编辑版位，你将能够自主选择广告投放的应用。
	//
	// 默认值: PLACEMENT_TYPE_NORMAL。
	//
	// 注意：
	//
	// 广告组创建后无法更新placement_type。
	// 若本字段设置为 PLACEMENT_TYPE_AUTOMATIC ，则返回的 placements 将显示实际支持的版位。例如，若实际仅支持 TikTok 版位，返回的 placements 值将为 PLACEMENT_TIKTOK。
	PlacementType enum.PlacementType `json:"placement_type,omitempty"`
	// Placements 投放版位。获取枚举值，详见枚举值-版位。
	// 当 placement_type 为 PLACEMENT_TYPE_NORMAL 时，本字段必填。
	// 当 placement_type 为 PLACEMENT_TYPE_AUTOMATIC 时，忽略本字段，系统将自动选择版位写入placements。
	// 注意：
	// 当placement_type为PLACEMENT_TYPE_AUTOMATIC（自动版位）时，本字段返回指定的推广目标下实际支持的版位。例如，若实际仅支持 TikTok版位，本字段返回的值将为PLACEMENT_TIKTOK 。
	// 若指定的推广目标支持 PLACEMENT_GLOBAL_APP_BUNDLE 版位，但您未申请全球化应用组合版位的白名单，则本字段的值将不会包含PLACEMENT_GLOBAL_APP_BUNDLE。
	// 当 placement_type 为 PLACEMENT_TYPE_NORMAL（手动版位）时，本字段通常将返回您在请求中通过placements字段指定的版位。
	// 唯一的例外是您在请求中通过 placements 字段指定了 PLACEMENT_GLOBAL_APP_BUNDLE 版位 ，但是您未申请 PLACEMENT_GLOBAL_APP_BUNDLE 的白名单，则本字段的返回值中将筛除PLACEMENT_GLOBAL_APP_BUNDLE 。例如，若您在请求中将 placements 设置为["PLACEMENT_TIKTOK", "PLACEMENT_GLOBAL_APP_BUNDLE"]，但是未申请全球化应用组合版位的白名单，本字段的返回值将为["PLACEMENT_TIKTOK"]。
	Placements []enum.Placement `json:"placements,omitempty"`
	// TiktokSubplacements TikTok 子版位，用于选择广告出现的具体位置。
	// 仅当满足以下条件时生效：
	// - 推广系列层级：objective_type 为 REACH、VIDEO_VIEWS 或 ENGAGEMENT。
	// - 广告组层级：placements 为 ["PLACEMENT_TIKTOK"] 且 placement_type 为 PLACEMENT_TYPE_NORMAL。
	//
	// 枚举值：
	// IN_FEED：信息流。广告将投放在“推荐”动态中，可能投放在“个人资料页”和“关注”推荐内容中。
	// SEARCH_FEED：搜索推荐内容。
	// TIKTOK_LITE即将废弃：TikTok Lite，占用内容更小、视频加载速度更快的精简版 TikTok 应用。TikTok Lite 子版位仅支持定向日本或韩国时使用。
	//
	// - 若您传入 tiktok_subplacements， tiktok_subplacements 的值中需至少包含 IN_FEED。
	// - 若想在 tiktok_subplacements 的值中包含 TIKTOK_LITE，您需要在定向地域（location_ids）中包含日本或韩国的地域。否则，将出现报错。
	// - 若想在 tiktok_subplacements 的值中包含 SEARCH_FEED 或 TIKTOK_LITE，不可同时将 brand_safety_type 设置为 THRID_PARTY
	//
	// 注意：
	//
	// 自 2025 年 7 月 19 日起，您将无法创建或复制使用子版位 TikTok Lite （TIKTOK_LITE）的广告组。为确保流畅、不受影响的广告创建体验，推荐您从广告组创建流程中移除此子版位或避免复制包含此类广告组的推广系列。
	// 在广告组中使用子版位 TikTok Lite（TIKTOK_LITE）目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 在广告组中使用子版位搜索推荐内容（SEARCH_FEED）目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 本字段设置后不允许更新。
	// 若您创建的广告组的 tiktok_subplacements 值中包含 TIKTOK_LITE，则不允许更新广告组时从定向地域（location_ids）中完全去除日本和韩国地域。
	// 子版位搜索推荐内容（SEARCH_FEED）支持的创新互动样式仅包括展示卡片。不过，该创新互动样式可能不支持您在推广系列层级设置的推广目标。若想了解展示卡片和推广目标的兼容情况，请查看卡片。
	// 子版位 TikTok Lite（TIKTOK_LITE）支持的创新互动样式仅包括展示卡片、倒计时贴纸和礼品码贴纸。不过，这些创新互动样式可能不支持您在推广系列层级设置的推广目标。若想了解这些新互动样式和推广目标的兼容情况，请查看卡片和贴纸。
	TiktokSubplacements []enum.TiktokSubplacement `json:"tiktok_subplacements,omitempty"`
	// SearchResultEnabled 是否在搜索广告中展示您的广告，即当用户在 TikTok上搜索您的业务相关信息时，系统是否向他们展示您的广告。
	// 若您使用以下推广目标 ( objective_type ) : APP_PROMOTION, WEB_CONVERSIONS , TRAFFIC , LEAD_GENERATION ，且将 placement_type 设置为 PLACEMENT_TYPE_AUTOMATIC，或将 placement_type设置为 PLACEMENT_TYPE_NORMAL的同时将 PLACEMENT_TIKTOK 包含在 placements 字段中, 则本字段会自动设置为true。此时您若想打开自动搜索广告位，可以不传入本字段，或为本字段传入 true。 此时您若想关闭自动搜索广告位，需将本字段设置为 false。
	// 若您未使用以下推广目标 ( objective_type ) : APP_PROMOTION, WEB_CONVERSIONS , TRAFFIC , LEAD_GENERATION ，或将 placement_type 设置为 PLACEMENT_TYPE_NORMAL 的同时未将 PLACEMENT_TIKTOK 包含在 placements 字段中，本字段将自动设置为false。此时您无法再将本字段设置为true。
	//
	// 若想了解搜索广告详情，请查看创建搜索广告。
	//
	// 注意：
	//
	// 本字段设置后不允许修改。
	SearchResultEnabled *bool `json:"search_result_enabled,omitempty"`
	// CommentDisabled 是否允许用户在TikTok上评论您的广告
	CommentDisabled *bool `json:"comment_disabled,omitempty"`
	// VideoDownloadDisabled 用户是否可以在 TikTok 上下载您的广告视频
	VideoDownloadDisabled *bool `json:"video_download_disabled,omitempty"`
	// ShareDisabled 本广告组中的广告是否禁止分享到第三方平台
	// 默认值： False。仅当以下条件均满足时 True的值会生效：
	// 推广系列层级的objective_type 设置为以下推广目标之一： APP_PROMOTION，WEB_CONVERSIONS， REACH， TRAFFIC，VIDEO_VIEWS，ENGAGEMENT， LEAD_GENERATION。
	// 版位：
	// placement_type设置为 PLACEMENT_TYPE_AUTOMATIC 或
	// placement_type设置为PLACEMENT_TYPE_NORMAL且 placements = PLACEMENT_TIKTOK
	ShareDisabled bool `json:"share_disabled,omitempty"`
	// BlockedPangleAppIDs Pangle 媒体屏蔽列表 ID
	// 可通过查询 Pangle 媒体屏蔽列表接口返回的app_package_id字段获取。仅当版位中含有 Pangle 时生效
	BlockedPangleAppIDs []string `json:"blocked_pangle_app_ids,omitempty"`
	// SavedAudienceID 已保存受众ID。
	// 以下情形支持该字段：
	// - 您的推广系列未选择房地产，就业，信贷行业类型（specical_industries）。
	// - 且您的推广系列的推广目标(objective_type) 不为商品销量(PRODUCT_SALES)。
	// - 且您的广告组选择TikTok广告位（即placement_type设置为PLACEMENT_TYPE_AUTOMATIC；或placement_type设置为PLACEMENT_TYPE_NORMAL且placements包含PLACEMENT_TIKTOK）。
	// - 且您的广告组未使用自动定向(即 auto_targeting_enabled 未设置为true)。
	//
	// 在使用此字段之前，请调用/dmp/saved_audience/create/创建已保存受众，并在返回中获取已保存受众ID。与您的已保存受众关联的广告主ID应与广告组中的广告主ID相同。否则，将会出现错误。
	// 如果您使用saved_audience_id创建广告组，将返回已保存受众ID和包含在已保存受众中的定向选项。
	// 请参阅创建已保存受众以了解详细的流程和代码示例。
	// 注意：
	// - 通过/dmp/saved_audience/create/创建已保存受众时，您可以指定gender等定向选项。但是，如果您使用已保存受众创建广告组，则必须避免同时设置saved_audience_id和已保存受众中已有的定向选项（如gender）。否则，将会出现错误。
	// - 如果您的已保存受众的设备设置与广告组的设备设置冲突，则会出现错误。例如：
	// - 当您的已保存受众的操作系统为Android时，但在创建广告组时，您将ios14_targeting设置为IOS14_PLUS。
	// - 当您在广告组中使用 Android 应用程序的 App ID，且将objective_type字段指定为APP_PROMOTION，但您的已保存受众操作系统为iOS。
	// - 如果您的API请求引发与已保存受众配置和广告组配置之间冲突相关的错误，则当前错误消息将指示问题涉及您的受众定向选项。为解决此问题，建议使用/dmp/saved_audience/list/检查相应定向选项的详细信息，以确定问题的来源，必要时创建新的已保存受众。
	// - 若saved_audience_id 创建时设置了age_groups， TikTok 广告的新增年龄定向限制中针对不同推广目标的年龄限制规定将也适用于该保存受众。在广告组中使用保存受众前，请确保其年龄定向合法。
	SavedAudienceID string `json:"saved_audience_id,omitempty"`
	// AutoTargetingEnabled 是否开启自动定向。
	// 注意：自 2024 年 6 月起，您将无法为广告组开启自动定向或定向扩量。为保证流畅的 API 体验，我们建议您使用智能定向。
	AutoTargetingEnabled *bool `json:"auto_targeting_enabled,omitempty"`
	// ShoppingAdsRetargetingType 购物广告受众再营销类型。
	// shopping_ads_type 设置为 VIDEO 且 product_source 设置为 CATALOG 时必填。
	//
	// 枚举值：
	// LAB1 : 对浏览过商品或者将商品加过购物车，但未购买过商品的受众进行再营销。
	// LAB2：对将商品加过购物车,但未购买过商品的受众进行再营销。
	// LAB3 : 使用自定义条件组合对受众进行再营销。
	// OFF : 不进行受众再营销。
	ShoppingAdsRetargetingType enum.ShoppingAdsRetargetingType `json:"shopping_ads_retargeting_type,omitempty"`
	// ShoppingAdsRetargetingActionsDays 所指定的受众行为的有效时间范围。在该时间范围内完成所指定行为的受众将成为购物广告受众再营销的对象。
	// 取值范围：1，2，3，7, 14，30，60，90，180。
	ShoppingAdsRetargetingActionsDays int `json:"shopping_ads_retargeting_actions_days,omitempty"`
	// IncludeCustomActions 作为包含条件筛选购物广告受众再营销对象的自定义行为
	IncludeCustomActions []CustomAction `json:"include_custom_actions,omitempty"`
	// ExcludeCustomActions 作为排除条件筛选购物广告受众再营销对象的自定义行为。
	ExcludeCustomActions []CustomAction `json:"exclude_custom_actions,omitempty"`
	// ShoppingAdsRetargetingCustomAudienceRelation 由 shopping_ads_retargeting_type 字段指定的视频购物广告重定向受众与由 audience_ids 指定的自定义受众间的逻辑关系。
	// 仅当以下条件均满足时生效：
	// - shopping_ads_type 设置为 VIDEO。
	// - product_source设置为 CATALOG。
	// - 传入shopping_ads_retargeting_type 且将其设置为LAB1，LAB2 或 LAB3。
	// - 传入audience_ids 。
	// 枚举值：
	// OR：取视频购物广告重定向受众和自定义受众的并集构建定向受众。定向受众中包含属于视频购物广告重定向受众或自定义受众的人群。
	// AND：取视频购物广告重定向受众和自定义受众的交集构建定向受众。定向受众中仅包含同时属于视频购物广告重定向受众和自定义受众的人群。
	// 若未设置本字段，将取视频购物广告重定向受众和自定义受众的交集构建定向受众。
	// 注意：设置本字段后，不允许手动将本字段更新为 null 值。
	ShoppingAdsRetargetingCustomAudienceRelation enum.FilterSetOperator `json:"shopping_ads_retargeting_custom_audience_relation,omitempty"`
	// LocationIDs 定向地域 ID。使用tool/region 获取定向地域ID
	// 您需至少传入location_ids或zipcode_ids其中一，允许同时传入location_ids和zipcode_ids。
	// 最大数量：3,000。若同时传入 location_ids 和 zipcode_ids，则单个广告组中地域 ID 和邮政编码 ID 的数量之和不可超过 3,000。
	// 注意：
	// 定向地域不允许重叠。例如，不支持同时定向美国和加利福尼亚州。
	// 若您在创建广告组时通过 location_ids 或 zipcode_ids 定向了美国地域，您可随后将定向地域更新为美国的其他地域，但不允许从定向地域中完全去除美国地域，从而只定向非美国地域。
	LocationIDs []string `json:"location_ids,omitempty"`
	// ZipcodeIDs 定向地域的邮政编码 ID。
	// 您需至少传入location_ids或zipcode_ids其中一，允许同时传入location_ids和zipcode_ids。
	// 最大数量：3,000。若同时传入 location_ids 和 zipcode_ids，则单个广告组中地域 ID 和邮政编码 ID 的数量之和不可超过 3,000。
	// 注意：
	// 邮政编码定向目前仅支持定向美国、加拿大、巴西、印度尼西亚、泰国、越南的地域。
	// 巴西、印度尼西亚、泰国、越南的邮政编码定向目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 不支持在启用了特殊广告分类（special_industries）的推广系列中使用邮政编码定向。
	// 不支持在 objective_type 为 RF_REACH 的推广系列中使用邮政编码定向。
	// 邮政编码定向仅支持 TikTok 版位。因此，placements 的值中不包含 PLACEMENT_TIKTOK。
	// 定向地域不允许重叠。例如，不支持同时定向美国和加利福尼亚州。
	// 若您在创建广告组时通过 location_ids 或 zipcode_ids 定向了美国地域，您可随后将定向地域更新为美国的其他地域，但不允许从定向地域中完全去除美国地域，从而只定向非美国地域。
	// 欲获取邮政编码 ID 的信息，您仅可使用/tool/targeting/info/。
	ZipcodeIDs []string `json:"zipcode_ids,omitempty"`
	// Langauges 受众语言。
	// 枚举值：详见枚举值-受众语言
	Languages []string `json:"languages,omitempty"`
	// Gender 定向受众性别
	// 枚举值: GENDER_FEMALE,GENDER_MALE,GENDER_UNLIMITED。
	Gender enum.GENDER `json:"gender,omitempty"`
	// AgeGroups 定向受众年龄
	// 在某些场景下，不允许创建定向美国、拉丁美洲、欧洲经济区、英国、瑞士或加拿大 13-17 年龄区间（AGE_13_17）受众的广告组。若想了解详情，请查看 TikTok 广告的新增年龄定向限制。
	// 若您使用了与美国、拉丁美洲、欧洲经济区、英国、瑞士或加拿大 13-17 年龄区间定向不兼容的定向设置，且将 age_groups 设置为 [] 或不传入字段 age_groups ，则该字段将默认为["AGE_18_24", "AGE_25_34", "AGE_35_44", "AGE_45_54", "AGE_55_100"]。
	// 注意：若您想要定向美国、拉丁美洲、欧洲经济区、英国、瑞士或加拿大 13-17 年龄区间且通常将 age_groups 设置为 [] 或不传入字段 age_groups，推荐您在该字段中明确指定您想要定向的具体年龄区间，例如AGE_18_24。
	AgeGroups []enum.Age `json:"age_groups,omitempty"`
	// SpendingPower 消费能力。
	// 枚举值：ALL, HIGH。如果设为HIGH，您就可以定向和普通用户相比 TikTok 广告投入支出通常更多的高消费能力用户。
	// 注意：
	// 本字段不支持推广目标PRODUCT_SALES和RF_REACH。
	// 本字段支持自动版位和包含 TikTok 的手动版位。因此，您需将 placement_type设置为 PLACEMENT_TYPE_AUTOMATIC，或在placements的值中包含PLACEMENT_TIKTOK。需注意，即使您指定了 Pangle 版位，消费能力定向也不会在 Pangle 版位生效。
	// 如果在广告组层级，auto_targeting_enabled 设为 true，那么该字段会自动设为ALL。
	SpendingPower enum.SpendingPower `json:"spending_power,omitempty"`
	// HouseholdIncome 您想定向的收入群体。
	// 枚举值：TOP5(家庭收入群体的前5%)， TOP10(家庭收入群体的前10%)，TOP10_25(家庭收入群体的前10-25%)， TOP25_50(家庭收入群体的前25-50%)。
	// 注意:
	// 该字段仅支持竞价广告。参考推广目标了解详情。
	// 该字段仅支持美国TikTok版位。
	HouseholdIncome []enum.HouseholdIncome `json:"household_income,omitempty"`
	// AudienceIDs 受众 ID 列表。
	// 通过/dmp/custom_audience/list/接口进行管理
	AudienceIDs []string `json:"audience_ids,omitempty"`
	// SmartAudienceEnabled 是否已启用智能受众。
	// 枚举值： true， false。
	// 若想了解智能受众详情及如何启用智能受众，请查看智能定向。
	SmartAudienceEnabled bool `json:"smart_audience_enabled,omitempty"`
	// ExcludeAudienceIDs 排除受众 ID 列表
	// 通过/dmp/custom_audience/list/ 接口进行管理
	ExcludeAudienceIDs []string `json:"exclude_audience_ids,omitempty"`
	// InterestCategoryIDs 用于定向受众的一般兴趣分类 ID 列表
	// 若想搜索或列举一般兴趣分类 ID，可使用/targeting/search/（推荐）或 /tool/interest_category/。
	// 若想获取基于您所在行业的推荐兴趣分类 ID 列表，可使用/tool/targeting_category/recommend/。
	InterestCategoryIDs []string `json:"interest_category_ids,omitempty"`
	// InterestKeywordIDs 用于定向受众的其他兴趣分类 ID 列表
	// 若想搜索其他兴趣分类 ID，可使用/targeting/search/（推荐）或/tool/interest_keyword/recommend/
	InterestKeywordIDs []string `json:"interest_keyword_ids,omitempty"`
	// PurchaseIntentionKeywordIDs 购买意向分类 ID 列表，用于定向近期行为表明其可能会购买某内容类别相关产品的受众。
	// 若想搜索或列举购买意向分类 ID，可使用/targeting/search/。
	// 注意：
	// 不支持同时传入 purchase_intention_keyword_ids 和interest_keyword_ids ，否则会出现关键词冲突。
	// purchase_intention_keyword_ids 只支持竞价广告，且对应的 objective_type（推广目标）应为 APP_PROMOTION，WEB_CONVERSIONS ，LEAD_GENERATION，TRAFFIC 或 PRODUCT_SALES（对应的 product_source 需设置为CATALOG），同时版位需包含 TikTok（PLACEMENT_TIKTOK）或 Pangle （PLACEMENT_PANGLE）。
	PurchaseIntentionKeywordIDs []string `json:"purchase_intention_keyword_ids,omitempty"`
	// Actions 用于定向的行为分类对象数组
	Actions []TargetingAction `json:"actions,omitempty"`
	// SmartInterestBehaviorEnabled 是否已启用智能型兴趣和行为。
	// 枚举值： true， false。
	// 若想了解智能型兴趣和行为详情及如何启用智能型兴趣和行为，请查看智能定向。
	SmartInterestBehaviorEnabled bool `json:"smart_interest_behavior_enabled,omitempty"`
	// IncludedPangleAudiencePackageIDs 包含 Pangle 流量包 ID。您可以使用接口/pangle_audience_package/get/获取受众ID (package_id) 。bind_type设置为INCLUDE。请不要同时传入该字段和excluded_pangle_audience_package_ids
	IncludedPangleAudiencePackageIDs []string `json:"include_pangle_audience_package_ids,omitempty"`
	// ExcludePangleAudiencePackageIDs 排除 Pangle 流量包 ID。您可以使用接口/pangle_audience_package/get/获取受众ID (package_id) 。bind_type设置为EXCLUDE。请不要同时传入该字段和included_pangle_audience_package_ids
	ExcludePangleAudiencePackageIDs []string `json:"exclude_pangle_audience_package_ids,omitempty"`
	// OperatingSystems 受众操作系统，枚举值：ANDROID, IOS
	// 仅支持传入一个值
	// 本字段在以下场景必填：
	// - objective_type 为 APP_PROMOTION时
	// - objective_type 为TRAFFIC 且 promotion_type 为 APP_IOS 或 APP_ANDROID 时
	OperatingSystems []enum.OperatingSystem `json:"operating_systems,omitempty"`
	// MinAndroidVersion 受众最低Android版本。枚举值：详见枚举值-最低Android版本。
	MinAndroidVersion string `json:"min_android_version,omitempty"`
	// Ios14Targeting 您想定向的 iOS 设备版本。
	// 推广系列的campaign_type为 IOS14_CAMPAIGN 时， 本字段必填，且必须设置为 IOS14_PLUS
	// 枚举值：
	// UNSET：运行 iOS 14.4 及更早版本的设备。
	// IOS14_MINUS：不受 iOS 14 新的 ATT 框架影响的 iOS 设备（运行 iOS 14.0 及更早版本的设备）。
	// IOS14_PLUS：运行iOS 14.5 及之后版本、已执行 ATT 框架的 iOS 设备。此值仅支持在专属推广系列中使用。
	// ALL：运行 iOS 14.5 及之后版本、已执行 ATT 框架的 iOS 设备。此值仅支持在 iOS 应用再营销广告或 iOS 重定向类型的视频购物广告（商品来源为商品库且优化位置为应用）中使用。
	//
	// 对于以下两种场景，将忽略您的传参，并将本字段设置为 ALL。
	// iOS 应用再营销广告：
	// app_promotion_type 为 APP_RETARGETING
	// 且 promotion_type 为 APP_IOS
	// iOS 重定向类型的视频购物广告（商品来源为商品库且优化位置为应用）：
	// shopping_ads_type 为 VIDEO
	// product_source 为 CATALOG
	// promotion_type 为 APP_IOS
	// 且 shopping_ads_retargeting_type 非 OFF
	// 对于其他场景，默认值为 UNSET。
	//
	// 本字段设置为 IOS14_PLUS 后不允许更新，且需满足以下条件：
	// 广告组层级：
	// app_id 字段设置为 iOS App 的 ID。
	// 若想获取iOS App 的 ID，可使用/app/list/。 iOS App 对应的platform 为 IOS 。
	// operating_systems 字段设置为["IOS"]。
	// 传入 min_ios_version 字段，且字段值与ios14_targeting 字段的值相匹配。
	// optimization_goal 字段设置为 CLICK，INSTALL，IN_APP_EVENT 或 VALUE。
	// 不传入以下字段：
	// min_android_version
	// shopping_ads_retargeting_type
	// shopping_ads_retargeting_actions_days
	// 广告层级：
	// deeplink_type 字段不可设置为 DEFERRED_DEEPLINK 。
	Ios14Targeting enum.Ios14Targeting `json:"ios14_targeting,omitempty"`
	// MinIosVersion 受众最低 iOS 版本。枚举值：详见枚举值-最低 iOS 版本
	// 传入ios14_targeting时本字段必填
	MinIosVersion string `json:"min_ios_version,omitempty"`
	// DeviceModelIDs 设备机型ID列表。关于设备机型更多信息，可参考获取设备型号列表
	// 可参考获取设备机型列表获取设备机型ID和状态的更多信息。创建广告时需确保对应设备处于活跃状态（ /tool/device_model/ 接口返回中is_active 为 True）。
	// 注意: 不可同时设置设备机型 (device_model_ids) 和设备价格(device_price_ranges) 。
	DeviceModelIDs []string `json:"device_model_ids,omitempty"`
	// NetworkTypes 网络类型。
	// 详见枚举值-网络类型。
	NetworkTypes []enum.NetworkType `json:"network_types,omitempty"`
	// CarrierIDs 运营商ID。 只有当运营商in_use字段为true时生效。获取详细信息，参阅获取运营商列表
	// 您可以使用接口/v1.3/tool/carrier/ 获取枚举值。
	CarrierIDs []string `json:"carrier_ids,omitempty"`
	// IspIDs 定向的互联网服务提供商的ID
	// 仅当同时通过location_ids传入有效的国家或地区层级的地域ID时生效。
	// 您可通过/tool/targeting/list/获取地域ID对应的可定向的互联网服务提供商ID。
	// 注意：若传入 isp_ids，不支持将版位设置为仅全球化应用组合（placements =PLACEMENT_GLOBAL_APP_BUNDLE）。
	IspIDs []string `json:"isp_ids,omitempty"`
	// DevicePriceRanges 设备价格区间（10000代表1000+）。该数字必须是50的倍数。
	// 重要提示: 最终的实际上限将在您设定的上限基础上增加50，您可以在TikTok广告管理平台的广告组设置中看到实际上限。例如，如果您设置的价格区间是[0, 250]，实际区间则为[0, 300]。
	DevicePriceRanges []int `json:"device_price_ranges,omitempty"`
	// AudienceType 应用重定向受众类型。枚举值可见枚举值 - 应用重定向受众类型
	// 仅当推广系列层级 app_promotion_type 设置为 APP_RETARGETING 时，本字段生效且必填。
	// 推广系列层级 app_promotion_type 未设置为 APP_RETARGETING 时，不支持本字段。
	// 若您想为应用再营销广告指定空自定义受众，则无需传入 audience_ids 和 excluded_audience_ids。
	// 若您想为应用再营销广告指定非空自定义受众，您可以通过 audience_ids 仅指定包含受众，或通过excluded_audience_ids 仅指定排除受众，或通过 audience_ids 和 excluded_audience_ids 同时指定包含受众和排除受众。
	// 如果audience_ids和excluded_audience_ids同时传入，这两个字段不能包含同样的受众 ID。
	AudienceType enum.AudienceType `json:"audience_type,omitempty"`
	// AudienceRule 受众规则
	// 仅支持 objective_type 为 TRAFFIC 或 APP_PROMOTION。详情见受众规则
	AudienceRule *AudienceRule `json:"audience_rule,omitempty"`
	// ContextualTagIDs 内容相关定向标签ID列表。请查看内容相关定向了解如何使用内容相关定向。
	// 您可使用 /tool/contextual_tag/get/ 获取可用的内容相关定向标签。请查看内容相关定向了解如何使用内容相关定向。
	// 注意：
	// 本字段目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 仅支持推广系列的推广目标为REACH 和 VIDEO_VIEWS 。
	// 不支持同时将brand_safety_type 设置为THIRD_PARTY 。
	ContextualTagIDs []string `json:"contextual_tag_ids,omitempty"`
	// BrandSafetyType 品牌安全类型。枚举值：
	// NO_BRAND_SAFETY：不使用任何品牌安全解决方案，即使用全部库存。您的广告可能会在某些包含成人主题的内容周围展示。
	// EXPANDED_INVENTORY：使用TikTok的品牌安全解决方案，即使用扩展库存。您的广告将不会出现在不当内容或包含成人主题的内容后面。在下个API版本中EXPANDED_INVENTORY将替代NO_BRAND_SAFETY，并成为默认的品牌安全选项。
	// STANDARD_INVENTORY：使用TikTok的品牌安全解决方案中的标准库存。您的广告将展示在适合大多数品牌的内容周围。
	// LIMITED_INVENTORY：使用TikTok的品牌安全解决方案中的有限库存。您的广告将展示在不包含成人主题的内容周围。
	// THIRD_PARTY：使用第三方品牌安全解决方案。
	// 您可以使用/tool/region/接口查询品牌安全设置对应的广告可投放国家或地区。
	// 注意：
	//
	// 出价前第一方品牌安全筛选对于使用应用推广（APP_PROMOTION），网站转化量（WEB_CONVERSIONS），访问量（TRAFFIC），或线索收集（LEAD_GENERATION） 或商品销量（PRODUCT_SALES）推广目标的竞价广告目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 出价前第三方品牌安全解决方案目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	BrandSafetyType enum.BrandSafetyType `json:"brand_safety_type,omitempty"`
	// BrandSafetyPartner 品牌安全合作伙伴。只有在品牌安全为第三方（brand_safety_type = THIRD_PARTY）时有效。枚举值：IAS， OPEN_SLATE（该合作伙伴在TikTok广告管理平台的对应名称为DoubleVerify，因为该合作伙伴已被DoubleVerify收购）。您可以使用/tool/region/接口，并根据自己的品牌安全设置获取可投放的国家和地区。您需要传入品牌安全类型和品牌安全合作伙伴。
	// 注意：
	// 出价前第三方品牌安全解决方案目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	BrandSafetyPartner string `json:"brand_safety_partner,omitempty"`
	// CategoryExcludeIDs 内容排除类别ID
	// 本字段仅支持满足以下条件时使用：
	// 推广系列层级：
	// objective_type 为 REACH，VIDEO_VIEWS，ENGAGEMENT，RF_REACH，APP_PROMOTION，WEB_CONVERSIONS，TRAFFIC 或 LEAD_GENERATION。
	// 广告组层级：
	// placements 设置为 ["PLACEMENT_TIKTOK"]。
	// brand_safety_type 设置为 STANDARD_INVENTORY 或LIMITED_INVENTORY。
	// 您可以使用/tool/content_exclusion/get/来获取内容类别 ID 列表（excluded_category_list），ID可以用于指定某内容类别，避免您的广告出现在该内容类型周围。
	CategoryExcludeIDs []string `json:"category_exclude_ids,omitempty"`
	// VerticalSensitiveID 行业敏感内容控制类型ID
	// 本字段仅支持满足以下条件时使用：
	// 推广系列层级：
	// objective_type 为 REACH，VIDEO_VIEWS 或 ENGAGEMENT。
	// 广告组层级：
	// placements 设置为 ["PLACEMENT_TIKTOK"]。
	// brand_safety_type 设置为 STANDARD_INVENTORY 或LIMITED_INVENTORY。
	//   您可以使用/tool/content_exclusion/get/ 获取包含敏感内容的行业类别（vertical_sensitivity_list），防止某类别出现在您的广告周围。
	VerticalSensitiveID string `json:"vertical_sensitive_id,omitempty"`
	// BudgetMode 广告预算类型
	// dule_type必须是SCHEDULE_START_END，您需要传入一个结束时间 (schedule_end_time)。当开启推广系列预算优化（CBO）时，广告组层级的budget_mode将被忽略。
	//
	// 当未开启CBO时， 广告组层级的budget_mode支持以下枚举值：
	// BUDGET_MODE_TOTAL：总预算。
	// BUDGET_MODE_DYNAMIC_DAILY_BUDGET：动态日预算，即一周七日的平均预算。单日实际消耗不超过设定的七日平均预算的125%。周消耗不会超过七日平均预算*7。
	// objective_type设置为TRAFFIC，APP_PROMOTION，WEB_CONVERSIONS，LEAD_GENERATIONPRODUCT_SALES（仅视频购物广告），REACH，VIDEO_VIEWS 或 ENGAGEMENT 且您不想设置总预算时，将本字段设置为BUDGET_MODE_DYNAMIC_DAILY_BUDGET。
	// 注意：为推广目标为REACH，VIDEO_VIEWS 或 ENGAGEMENT 的推广系列启用 BUDGET_MODE_DYNAMIC_DAILY_BUDGET 目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	//
	// BUDGET_MODE_DAY：日预算。
	// objective_type 设置为REACH，VIDEO_VIEWS 或 ENGAGEMENT时，推荐您使用动态日预算（BUDGET_MODE_DYNAMIC_DAILY_BUDGET）而非日预算（BUDGET_MODE_DAY）。
	//
	// 了解如何设置预算类型，请参阅预算。
	//
	// 注意：
	//
	// 如果本字段设置为BUDGET_MODE_DAY，那么schedule_type可以是SCHEDULE_START_END或SCHEDULE_FROM_NOW。
	// 如果本字段设置为BUDGET_MODE_TOTAL，那么sche
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 广告组预算。
	// 当开启了推广系列预算优化（budget_optimize_on = TRUE）时，该字段将被忽略。
	// 详情请参阅预算。欲直接获取某一币种对应的每日预算取值范围，请参阅币种-每日预算取值范围。
	Budget float64 `json:"budget,omitempty"`
	// ScheduleType 广告投放时间类型。
	// 枚举值: SCHEDULE_FROM_NOW，SCHEDULE_START_END。如果您选择了SCHEDULE_START_END，您需要明确投放开始和结束时间。 如果您选择了SCHEDULE_FROM_NOW，您只需要明确投放开始时间。
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime 广告投放起始时间(UTC+0)。
	// 格式：YYYY-MM-DD HH:MM:SS。
	ScheduleStartTime model.DateTime `json:"schedule_start_time,omitzero"`
	// ScheduleEndTime 广告投放结束时间(UTC+0)。
	// 格式： YYYY-MM-DD HH:MM:SS
	ScheduleEndTime model.DateTime `json:"schedule_end_time,omitzero"`
	// Dapparting 广告投放安排。格式为48x7的字符串。字符只能为0或者1。0代表不投放，1代表投放。每个字符对应一个30分钟的时间段。第一个字符对应的是周一的凌晨0:01-0:30，第二个字符对应0:31-1:00，依次类推。最后一个字符代表周日23:31-0:00。
	// 注意：
	// 不设置，全部设置为0，或者全部都设置为1都代表了要进行全时投放。
	Dayparting string `json:"dayparting,omitempty"`
	// OptimizationGoal 优化目标
	// 若想获取完整的枚举值，请参阅枚举值 - 优化目标。
	// 若想获取 CBO 推广系列支持的枚举值，请参阅推广系列预算优化 - 支持的优化目标。
	//
	// 对于每个优化目标，您需同时手动指定对应的billing_event。若想了解与某一优化目标关联的计费事件，可查看优化目标关联的计费事件。
	//
	// 若要在 Pangle 版位或全球化应用组合版位的广告组中使用优化目标“安装与应用内事件数据”，需将 optimization_goal 设置为INSTALL，传入 secondary_optimization_event 合法值，且在placements的值中仅包含 PLACEMENT_PANGLE或 PLACEMENT_GLOBAL_APP_BUNDLE 或PLACEMENT_PANGLE和 PLACEMENT_GLOBAL_APP_BUNDLE。
	//
	// 注意：
	//
	// promotion_type 为LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE、LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE 或 LEAD_GEN_CLICK_TO_CALL 时，本字段默认值为 CLICK。
	// 自 2024 年 11 月 30 日起，您将无法创建或复制优化目标为“安装与应用内事件数据”且版位为仅 TikTok 或自动版位的广告组。此变动将影响使用以下配置的广告组：
	// 推广系列层级：objective_type 为 APP_PROMOTION
	// 广告组层级：
	// optimization_goal 为 INSTALL 且指定了 secondary_optimization_event 合法值
	// placement_type 为 PLACEMENT_TYPE_NORMAL 且 placements 为["PLACEMENT_TIKTOK"]，或 placement_type 为PLACEMENT_TYPE_AUTOMATIC
	// 现有的投放在 TikTok 版位的“安装与应用内事件数据”广告组将不受影响。此外，Pangle 版位和全球化应用组合版位（即 placement_type 为 PLACEMENT_TYPE_NORMAL 且 placements 的值仅包含PLACEMENT_PANGLE 和 PLACEMENT_GLOBAL_APP_BUNDLE）也不受影响。请注意此变更，并在必要时对您的集成进行适当的调整。
	// VIDEO_VIEW枚举值已废弃。当 objective_type 在推广计划层级设为 VIDEO_VIEWS 时，您可使用ENGAGED_VIEW 或 ENGAGED_VIEW_FIFTEEN（白名单功能）作为优化目标。
	// 若您要在一个已开启CBO的推广系列下创建新的广告组，需确保本字段的值与第一个广告组中的设置保持一致。
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// SecondaryOptimizationEvent 当优化目标为INSTALL或VALUE时的次要优化目标。枚举值：详见转化事件-App深度转化事件类型。
	// 注意：
	//
	// 自 2024 年 11 月 30 日起，您将无法创建或复制优化目标为“安装与应用内事件数据”且版位为仅 TikTok 或自动版位的广告组。此变动将影响使用以下配置的广告组：
	// 推广系列层级：objective_type 为 APP_PROMOTION
	// 广告组层级：
	// optimization_goal 为 INSTALL 且指定了 secondary_optimization_event 合法值
	// placement_type 为 PLACEMENT_TYPE_NORMAL 且 placements 为["PLACEMENT_TIKTOK"]，或 placement_type为PLACEMENT_TYPE_AUTOMATIC
	// 现有的投放在 TikTok 版位的“安装与应用内事件数据”广告组将不受影响。此外，Pangle 版位和全球化应用组合版位（即 placement_type 为 PLACEMENT_TYPE_NORMAL 且 placements 的值仅包含PLACEMENT_PANGLE 和 PLACEMENT_GLOBAL_APP_BUNDLE）也不受影响。请注意此变更，并在必要时对您的集成进行适当的调整。
	// 若设置了 secondary_optimization_event，则需同时传入 deep_bid_type 。
	SecondaryOptimizationEvent enum.SecondaryOptimizationEvent `json:"secondary_optimization_event,omitempty"`
	// MessageEventSetID 即时通讯广告组中使用的消息事件集的 ID。
	// 若广告组设置中的即时通讯应用账号匹配到了唯一消息事件集，无论是通过 messaging_app_account_id 手动指定的 Messenger 账号还是基于指定的phone_region_code、phone_region_calling_code 和 phone_number 填充的 WhatsApp 账号，则本字段将自动填充为与所选即时通讯应用账号绑定的唯一消息事件集。
	MessageEventSetID string `json:"message_event_set_id,omitempty"`
	// Frequency 频次，与frequency_schedule共同构成用户看到广告的频次上限（仅适用于REACH广告）。
	// 需要满足以下两个条件：
	// 1 <= frequency <= 1000
	// 1 <= frequency_schedule <=30。
	// 例如：frequency = 2 且 frequency_schedule = 3 表示每 3 天至多 2 次展示。
	Frequency int `json:"frequency,omitempty"`
	// FrequencySchedule 频次天数，与frequency共同构成用户看到广告的频次上限（仅适用于REACH广告）。详见frequency字段说明。
	FrequencySchedule int `json:"frequency_schedule,omitempty"`
	// BidType 竞价策略。
	// 竞价策略将决定系统的单次成效费用管理方式、预算支出方式以及广告投放方式。
	// 当开启推广系列预算优化（CBO）时，本字段必填。
	// 若想查看枚举值，请参阅枚举值-竞价策略。
	// 注意：若您要在一个已开启 CBO 的推广系列下创建新的广告组，需确保本字段的值与第一个广告组中的设置保持一致。
	BidType enum.BidType `json:"bid_type,omitempty"`
	// BidPrice 出价。（成本上限策略下）想尽可能达到的平均单次成效成本
	// bid_type 为 BID_TYPE_CUSTOM 且 billing_event 为 CPC，CPM 或 CPV 时，本字段必填
	// 当开启了推广系列预算优化（budget_optimize_on = true）时，建议同一个推广系列下所有广告组出价保持一致。
	// bid_price需小于您在推广系列层级和广告组层级设置的budget金额。请查看出价-出价限制了解出价校验机制。
	BidPrice float64 `json:"bid_price,omitempty"`
	// ConversionBidPrice oCPM转化出价
	// bid_type 为 BID_TYPE_CUSTOM 且 billing_event 为 OCPM 时，本字段必填
	// conversion_bid_price需小于您在推广系列层级和广告组层级设置的budget金额。请查看出价-出价限制了解出价校验机制
	ConversionBidPrice float64 `json:"conversion_bid_price,omitempty"`
	// DeepBidType 深度事件出价类型
	// 当开启推广系列预算优化（CBO）且optimization_goal为VALUE时，本字段必填。
	// 枚举值：VO_MIN_ROAS（白名单功能），VO_HIGHEST_VALUE（白名单功能）。
	// 枚举值详见枚举值-深度事件出价类型。
	// 注意：若设置了 secondary_optimization_event，则需同时传入 deep_bid_type 。
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
	// RoasBid 用于价值优化的 ROAS 目标值
	// 当deep_bid_type为VO_MIN_ROAS时必填
	// 取值范围:
	// 对于应用内广告收入场景（objective_type 为 APP_PROMOTION， optimization_goal 为 VALUE，且 optimization_event 为 AD_REVENUE_VALUE）：0.01-10。
	// 对于商品来源为 TikTok Shop 的商品购物广告场景（shopping_ads_type 为 PRODUCT_SHOPPING_ADS 且 product_source 为 STORE）：0.01-20。
	// 对于其他场景：0.01-1,000。
	RoasBid float64 `json:"roas_bid,omitempty"`
	// VboWindow 为广告收入价值优化或应用内购价值优化指定的竞价策略的时间窗口期。
	// 枚举值：
	// SEVEN_DAYS：前 7 天。
	// ZERO_DAY：第 0 天（当天）。
	//
	// 使用竞价策略最低 ROAS（目标 ROAS）时，本字段代表目标 ROAS 值的窗口期类型（第 7 天 ROAS 或第 0 天 ROAS）。系统将努力让第 7 天或当天的 ROAS 保持在或超过目标 ROAS 值。
	// 使用竞价策略最高价值时，本字段代表最高价值的窗口期类型（前 7 天最高价值或第 0 天最高价值 ）。系统将在前 7 天或当天优先充分消耗预算，同时尽可能获得最大的成效价值。
	//
	// 若想了解如何为广告启用广告收入价值优化或应用内购价值优化，请查看开启应用价值优化。
	//
	// 注意：
	//
	// 不同场景下的广告收入价值优化或应用内购价值优化目前为白名单功能。如需使用这些功能，请联系您的 TikTok 销售代表。
	// 在高级专属推广系列或 Android 推广系列中使用应用内购价值优化的第 0 天出价或第 7 天出价（vbo_window 为 ZERO_DAY 或 SEVEN_DAYS）目前为白名单功能。如需使用这些功能，请联系您的 TikTok 销售代表。
	// 在高级专属推广系列或 Android 推广系列中使用广告收入价值优化的第 0 天出价或第 7 天出价（vbo_window 为 ZERO_DAY 或 SEVEN_DAYS）目前为白名单功能。如需使用这些功能，请联系您的 TikTok 销售代表。
	// 本字段设置后不允许更新。
	VboWindow enum.Window `json:"vbo_window,omitempty"`
	// BidDisplayMode 每次浏览成本的计算和测算模式。枚举值：CPMV(每千次浏览成本) / CPV(每次浏览成本) 。默认值：CPMV
	// 当在推广系列层级的objective_type设置为VIDEO_VIEWS，bid_display_mode必填且必须设为 CPV
	// 注意：
	//
	// bid_display_mode一旦设置无法更新。
	// 如果推广系列层级的objective_type不为VIDEO_VIEWS, 您不需要设置此字段。
	// 枚举值CPMV已经在v1.3版本废弃。
	BidDisplayMode enum.BidDisplayMode `json:"bid_display_mode,omitempty"`
	// DeepCpaBid 当您为app内事件选择出价策略后，您需要在该字段中传入出价价格
	DeepCpaBid float64 `json:"deep_cpa_bid,omitempty"`
	// NextDayRetention 次日留存率。计算公式：next_day_retention = conversion_bid_price/deep_cpa_bid，范围 (0，1]。仅当placements为PLACEMENT_PANGLE且深度转化事件secondary_optimization_event为NEXT_DAY_OPEN时生效
	NextDayRetention float64 `json:"next_day_retention,omitempty"`
	// ClickAttributionWindow 广告组的点击归因窗口期。该窗口期指用户从点击广告到采取行动之间的时间间隔。
	// 枚举值：
	// OFF：不启用。
	// ONE_DAY：1天（点击）。
	// SEVEN_DAYS：7天（点击）。
	// FOURTEEN_DAYS：14天（点击）。
	// TWENTY_EIGHT_DAYS：28天（点击）。
	// 注意：
	// 若想手动指定归因窗口期，您需至少同时传入click_attribution_window 和 view_attribution_window，可同时传入engaged_view_attribution_window。需注意，一旦指定 engaged_view_attribution_window，需同时指定 click_attribution_window 和 view_attribution_window。
	// 若您未传入click_attribution_window，engaged_view_attribution_window，view_attribution_window 和attribution_event_count字段，则如果有默认设置，会使用默认设置。若想确认您未手动传入本字段时，是否使用了默认设置，可使用 /adgroup/get/ 。
	// 本字段设置后不支持更新。
	ClickAttributionWindow enum.Window `json:"click_attribution_window,omitempty"`
	// EngagedViewAttributionWindow 广告组的深度互动观看归因窗口期。该窗口期指用户从观看视频广告 6 秒到采取行动之间的时间间隔。
	// 枚举值：
	// ONE_DAY：1天（深度互动观看）。
	// SEVEN_DAYS：7天（深度互动观看）。
	// 注意：
	// 若想手动指定归因窗口期，您需至少同时传入click_attribution_window 和 view_attribution_window，可同时传入engaged_view_attribution_window。需注意，一旦指定 engaged_view_attribution_window，需同时指定 click_attribution_window 和 view_attribution_window。
	// 若您未传入click_attribution_window，engaged_view_attribution_window，view_attribution_window 和attribution_event_count字段，则如果有默认设置，会使用默认设置。若想确认您未手动传入本字段时，是否使用了默认设置，可使用 /adgroup/get/ 。
	// 本字段设置后不支持更新。
	EngagedViewAttributionWindow enum.Window `json:"engaged_view_attribution_window,omitempty"`
	// ViewAttributionWindow 广告组的展示归因窗口期。该窗口期指用户从查看广告到采取行动之间的时间间隔。
	// 枚举值：
	// OFF：不启用。
	// ONE_DAY：1天（展示）。
	// SEVEN_DAYS：7天（展示）。
	// 注意：
	// 若想手动指定归因窗口期，您需至少同时传入click_attribution_window 和 view_attribution_window，可同时传入engaged_view_attribution_window。需注意，一旦指定 engaged_view_attribution_window，需同时指定 click_attribution_window 和 view_attribution_window。
	// 若您未传入click_attribution_window，engaged_view_attribution_window，view_attribution_window 和attribution_event_count字段，则如果有默认设置，会使用默认设置。若想确认您未手动传入本字段时，是否使用了默认设置，可使用 /adgroup/get/ 。
	// 本字段设置后不支持更新。
	ViewAttributionWindow enum.Window `json:"view_attribution_window,omitempty"`
	// AttributionEventCount 广告组的事件统计方式（统计类型）。
	// 统计用户在查看或点击了广告后采取操作的次数的方式。
	// 枚举值：
	// UNSET ：未设置。
	// EVERY：每一次。某位用户进行的多个购买事件将分别计入转化量。
	// ONCE：仅一次。某位用户进行的多个事件将只算作 1 次转化。
	// 注意：
	// 若您未传入click_attribution_window，engaged_view_attribution_window，view_attribution_window 和attribution_event_count字段，则如果有默认设置，会使用默认设置。若想确认您未手动传入本字段时，是否使用了默认设置，可使用 /adgroup/get/ 。
	// 本字段设置后不支持更新。
	AttributionEventCount enum.AttributionEventCount `json:"attribution_event_count,omitempty"`
	// BillingEvent 计费事件。
	// 若想了解与某一优化目标关联的计费事件，可查看优化目标关联的计费事件。
	BillingEvent enum.BillingEvent `json:"billing_event,omitempty"`
	// Pacing 广告投放速度类型。您可以选择PACING_MODE_SMOOTH（在预定的时间内平均分配预算）和PACING_MODE_FAST（尽快消耗预算并产出结果）。当您开启推广系列预算优化（budget_optimize_on）时，该字段将自动设置为PACING_MODE_SMOOTH。否则您需要填写该字段。
	Pacing enum.PacingMode `json:"pacing,omitempty"`
	// StatisticType 转化事件出价统计类型。枚举值 EVERYTIME(每一次计费广告)，NONE(仅一次计费广告)
	// 您可以通过/app/optimization_event/接口获取应用事件对应的statistic_type字段值。
	StatisticType string `json:"statistic_type,omitempty"`
	// IsHfss 广告推广对象是否是 HFSS 食品（高脂、高盐、高糖食品）请注意，欧洲国家不允许向未成年人推送 HFSS 食品广告
	// 请注意，欧洲国家不允许向未成年人推送 HFSS 食品广告
	IsHfss bool `json:"is_hfss,omitempty"`
	// OperationStatus 广告组的操作状态。
	// ENABLE：广告组处于启用（“开”）状态。
	// DISABLE：广告组处于未启用（“关”）状态。
	// 若想在广告组创建后更新广告组的启用/关闭状态，请使用 /adgroup/update/status/ 接口。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// CreativeMaterialMode 创意投放方式。
	// 枚举值: CUSTOM（自定义）、DYNAMIC（已废弃）（程序化创意）、SMART_CREATIVE（智能创意）
	// 注意: 程序化创意广告（creative_material_mode = DYNAMIC）已下线。我们建议您创建智能创意广告（creative_material_mode = CUSTOM），以便您更顺畅地迁移。欲知详细信息，请参考智能创意广告。
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// AdgroupAppProfilePageState 该广告组是否使用了App中间页。
	// 枚举值：ON，OFF。
	//
	// 注意: 应用介绍页不支持在优化目标为应用内事件且竞价策略为成本上限的专属推广系列中使用。因此本字段在下列条件均满足时不支持设置为ON：
	//
	// ios14_targeting 设置为 IOS14_PLUS
	// optimization_goal 设置为 IN_APP_EVENT
	// bid_type 设置为 BID_TYPE_CUSTOM
	AdgroupAppProfilePageState enum.AppProfilePageState `json:"adgroup_app_profile_page_state,omitempty"`
}

// Encode implements PostRequest interface
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建广告组 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Adgroup `json:"data,omitempty"`
}
