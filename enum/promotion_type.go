package enum

// PromotionType 推广对象类型 （优化位置）
type PromotionType string

const (
	// PromotionType_APP_ANDROID 安卓应用
	PromotionType_APP_ANDROID PromotionType = "APP_ANDROID"
	// PromotionType_APP_IOS iOS 应用
	PromotionType_APP_IOS PromotionType = "APP_IOS"
	// PromotionType_GAME 游戏
	// 仅支持在应用帖子推广（即 app_promotion_type 为 APP_POSTS_PROMOTION）类型的推广系列中推广游戏。
	// 注意：目前不支持通过 API 创建应用帖子推广类型的推广系列，仅支持通过 API 更新或检索此类推广系列。
	PromotionType_GAME PromotionType = "GAME"
	// PromotionType_WEBSITE 落地页
	PromotionType_WEBSITE PromotionType = "WEBSITE"
	// PromotionType_LEAD_GENERATION 线索广告中的即时表单（线索表单）或落地页。
	// promotion_type 设置为 LEAD_GENERATION 时，您可将 promotion_target_type 设置为：
	// INSTANT_PAGE ：即时表单。
	// 或 EXTERNAL_WEBSITE ：落地页。
	PromotionType_LEAD_GENERATION PromotionType = "LEAD_GENERATION"

	// PromotionType_LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE 通过 TikTok 私信收集线索。
	// 若想了解如何创建推广对象类型为 TikTok 私信的线索广告，请查看这里。
	//
	// 注意：
	//
	// 创建优化位置为 TikTok 私信的线索广告对注册地为以下地区的广告主已全量开放：
	// 亚太地区：澳大利亚，孟加拉国，柬埔寨，印度尼西亚，日本，马来西亚，新西兰，巴基斯坦，菲律宾，新加坡，韩国，斯里兰卡，中国台湾地区，泰国和越南。
	// 拉美地区：阿根廷，玻利维亚，巴西，智利，哥伦比亚，哥斯达黎加，多米尼加共和国，厄瓜多尔，危地马拉，墨西哥，巴拿马，巴拉圭，秘鲁，波多黎各和乌拉圭。
	// 对于注册地为其他地区的广告主，此功能目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	PromotionType_LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE PromotionType = "LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE"
	// PromotionType_LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE 通过即时通讯应用收集线索。
	// 注意：
	//
	// 创建优化位置为即时通讯应用的线索广告对注册地为以下亚太地区的广告主已全量开放：澳大利亚，孟加拉国，柬埔寨，印度尼西亚，日本，马来西亚，新西兰，巴基斯坦，菲律宾，新加坡，韩国，斯里兰卡，中国台湾地区，泰国和越南。
	// 对于注册地为其他地区的广告主，此功能目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	PromotionType_LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE PromotionType = "LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE"
	// PromotionType_LEAD_GEN_CLICK_TO_CALL 通过电话通话收集线索。
	// 注意：创建推广对象类型为电话通话的线索广告目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	PromotionType_LEAD_GEN_CLICK_TO_CALL PromotionType = "LEAD_GEN_CLICK_TO_CALL"
	// PromotionType_WEBSITE_OR_DISPLAY 落地页或纯展示页。如果推广对象为网站，且推广对象不需传入landing_page_url时，使用WEBSITE_OR_DISPLAY (对应推广目标为Reach / Video Views / Engagement) 。对于其他推广目标(Traffic / Conversion)，如果推广对象为网站，那么请使用WEBSITE
	PromotionType_WEBSITE_OR_DISPLAY PromotionType = "WEBSITE_OR_DISPLAY"
	// PromotionType_TIKTOK_SHOP TikTok商店
	PromotionType_TIKTOK_SHOP PromotionType = "TIKTOK_SHOP"
	// PromotionType_VIDEO_SHOPPING 仅当shopping_ads_type 设置为VIDEO 且 product_source 设置为 STORE 或 SHOWCASE 时生效。
	// 若 shopping_ads_type 设置为VIDEO 且 product_source 设置为CATALOG，需将 promotion_type 设置为 WEBSITE，APP_ANDRIOD 或 APP_IOS。
	// 视频购物。在短视频上展示商品链接。观众可点击购买产品。
	PromotionType_VIDEO_SHOPPING PromotionType = "VIDEO_SHOPPING"
	// PromotionType_LIVE_SHOPPING 仅当 shopping_ads_type 设置为 LIVE 时生效。
	// 直播购物。在直播间介绍商品并展示商品链接。观众可点击购买。
	PromotionType_LIVE_SHOPPING PromotionType = "LIVE_SHOPPING"
	// PromotionType_PSA_PRODUCT 仅当shopping_ads_type 设置为PRODUCT_SHOPPING_ADS 时生效。
	// 通过购物中心和搜索流量推广商品。
	PromotionType_PSA_PRODUCT PromotionType = "PSA_PRODUCT"
	// PromotionType_EXTERNAL_OR_DISPLAY
	PromotionType_EXTERNAL_OR_DISPLAY PromotionType = "EXTERNAL_OR_DISPLAY"
)
