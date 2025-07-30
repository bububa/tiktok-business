package enum

// ObjectiveType 推广目标
type ObjectiveType string

const (
	// ObjectiveType_APP_PROMOTION 应用推广。吸引用户安装您的移动应用，并再次吸引应用的已有用户。
	ObjectiveType_APP_PROMOTION ObjectiveType = "APP_PROMOTION"
	// ObjectiveType_WEB_CONVERSIONS 网站转化量。引导用户前往您的网站或即时体验页面采取特定行动，例如购买商品或将商品加入购物车。
	ObjectiveType_WEB_CONVERSIONS ObjectiveType = "WEB_CONVERSIONS"
	// ObjectiveType_APP_INSTALL 应用安装量，吸引更多用户安装您的应用。
	// 已废弃
	// 注意:
	//
	// 新目标应用推广下后续将支持更多新功能，旧目标应用安装量下则不会支持新功能。我们建议您使用新的推广目标应用推广。
	// 现有创建于2023年10月27日以前的应用安装量广告将不受影响，预算用尽前将始终保持有效。您可以随时终止或删除这些推广系列、广告组、广告及程序化创意广告。
	ObjectiveType_APP_INSTALL ObjectiveType = "APP_INSTALL"
	// ObjectiveType_CONVERSIONS 转化量，吸引用户在您的网站/应用中执行更多有价值的操作。
	// 已废弃
	// 注意:
	//
	// 新目标网站转化量下后续将支持更多新功能，旧目标转化量下则不会支持新功能。我们建议您使用新的推广目标网站转化量。
	// 现有创建于2023年10月27日以前的转化量广告将不受影响，预算用尽前将始终保持有效。您可以随时终止或删除这些推广系列、广告组、广告及程序化创意广告。
	ObjectiveType_CONVERSIONS ObjectiveType = "CONVERSIONS"
	// ObjectiveType_REACH 竞价覆盖。
	// 向尽可能多的用户展示竞价广告。
	ObjectiveType_REACH ObjectiveType = "REACH"
	// ObjectiveType_TRAFFIC 访问量，吸引更多用户访问您的目标网站或应用。
	ObjectiveType_TRAFFIC ObjectiveType = "TRAFFIC"
	// VIDEO_VIEWS 视频播放量，吸引更多用户观看视频。
	ObjectiveType_VIDEO_VIEWS ObjectiveType = "VIDEO_VIEWS"
	// ObjectiveType_CATALOG_SALES（白名单功能） 目录促销，吸引用户在商品库中购买您的商品
	// 已废弃
	// 注意：自2023年2月16日起，您将不能创建目录促销目标下的Dynamic Showcase Ads。为保证流畅的API体验，我们建议您创建商品销量目标下的视频购物广告。
	ObjectiveType_CATALOG_SALES ObjectiveType = "CATALOG_SALES"
	// ObjectiveType_PRODUCT_SALES 商品销量，为商业类广告主提供更加简单便捷的投放广告方式，销售来自 TikTok Shop、网站和应用的商品。
	// 注意：商品来源为商品库的视频购物广告目前对于商业类的移动应用广告主为白名单功能。若您为商业类的移动应用广告主且需使用此功能，请联系您的 TikTok 销售代表。若您是商业类的网络广告主，您已自动拥有该功能的权限。
	ObjectiveType_PRODUCT_SALES ObjectiveType = "PRODUCT_SALES"
	// ObjectiveType_ENGAGEMENT 社区互动，获得更多账号关注或主页访问等
	ObjectiveType_ENGAGEMENT ObjectiveType = "ENGAGEMENT"
	// ObjectiveType_LEAD_GENERATION 线索收集，为业务或品牌开发潜在客户
	ObjectiveType_LEAD_GENERATION ObjectiveType = "LEAD_GENERATION"
	// ObjectiveType_SHOP_PURCHASES 店铺推广，用于短视频带货或者直播带货
	// 已废弃
	ObjectiveType_SHOP_PURCHASES ObjectiveType = "SHOP_PURCHASES"
	// ObjectiveType_RF_APP_INSTALL （白名单功能） 应用安装量，吸引更多用户安装您的应用, 用于广告
	// 已废弃
	ObjectiveType_RF_APP_INSTALL ObjectiveType = "RF_APP_INSTALL"
	// ObjectiveType_RF_ENGAGEMENT （白名单功能） 社区互动，获得更多账号关注或主页访问等，用于覆盖和频次广告。 该推广目标目前不支持品牌安全(Brand Safety)
	// 已废弃
	ObjectiveType_RF_ENGAGEMENT ObjectiveType = "RF_ENGAGEMENT"
	// ObjectiveType_RF_REACH （白名单功能）
	// 合约覆盖和频次或 TikTok Pulse。
	// 向尽可能多的用户展示提前预订的广告保证覆盖人数和展示频次，或以固定的千次展示成本（CPM）在推荐页上排名前 4% 的热门容后面展示的广告。
	//
	// 注意：
	//
	// 若想使用覆盖和频次目标，可使用/campaign/create/将 objective_type 设置为 RF_REACH 且将 rf_campaign_type 设置为 STANDARD。
	// 若想使用 TikTok Pulse 目标，可使用/campaign/create/将 objective_type 设置为 RF_REACH 且将 rf_campaign_type 设置为 PULSE。
	ObjectiveType_RF_REACH ObjectiveType = "RF_REACH"
	// ObjectiveType_TOPVIEW_REACH  合约 TopView 。
	// 在尽可能多的用户打开应用时展示广告。
	//
	// 注意：
	//
	// 不支持在 /campaign/get/，/adgroup/get/ 和 /ad/get/ 接口中的筛选字段 objective_type 指定本枚举值从而筛选 TopView 数据。若想筛选 TopView 推广系列、广告组或广告的数据，需将筛选字段 buying_types 设置为 ["RESERVATION_TOP_VIEW"]。
	// 不支持使用 API 创建、更新或复制 TopView 广告，您仅可通过 API 检索已有的 TopView 广告。若想了解 API 中支持的 TopView 广告功能，请查看 TopView。
	ObjectiveType_TOPVIEW_REACH ObjectiveType = "TOPVIEW_REACH"
	// ObjectiveType_RF_TRAFFIC（白名单功能）访问量，吸引更多用户访问您的目标网站或应用，用于覆盖和频次广告。该推广目标目前不支持品牌安全(Brand Safety)。
	// 已废弃
	// 注意：覆盖和频次广告的推广目标视频播放量和访问量于2023年6月1日废弃。2023年6月1日前保存的存量订单可继续投放，同时支持编辑与暂停，但对应的存量广告的投放结束时间需不晚于2024年1月12日。
	ObjectiveType_RF_TRAFFIC ObjectiveType = "RF_TRAFFIC"
	// ObjectiveType_RF_VIDEO_VIEW （白名单功能）视频播放量，吸引更多用户观看视频 , 用于覆盖和频次广告。
	// 已废弃
	// 注意：覆盖和频次广告的推广目标视频播放量和访问量于2023年6月1日废弃。2023年6月1日前保存的存量订单可继续投放，同时支持编辑与暂停，但对应的存量广告的投放结束时间需不晚于2024年1月12日。
	ObjectiveType_RF_VIDEO_VIEW ObjectiveType = "RF_VIDEO_VIEW"
)
