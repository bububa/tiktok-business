package enum

// OptimizationGoal 优化目标
type OptimizationGoal string

const (
	// OptimizationGoal_CLICK 点击/点击数。
	// CLICK 在以下任意场景中对应优化目标点击数：
	// shopping_ads_type 设置为 VIDEO 且 product_source 设置为 STORE 或 SHOWCASE。
	// shopping_ads_type 设置为 LIVE。
	// shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS 且 product_source 设置为 STORE。
	OptimizationGoal_CLICK OptimizationGoal = "CLICK"
	// OptimizationGoal_CONVERT 转化/付费数/开始结账数。
	// CONVERT（对应的 optimization_event 设置为SHOPPING）在以下任意场景中对应优化目标付费数：
	// shopping_ads_type 设置为 VIDEO 且 product_source 设置为 STORE 或 SHOWCASE。
	// shopping_ads_type 设置为 LIVE。
	// shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS 且 product_source 设置为 STORE。
	//
	// CONVERT（对应的 optimization_event 设置为 INITIATE_ORDER） 在以下任意场景中对应优化目标开始结账数：
	// shopping_ads_type 设置为 VIDEO 且 product_source 设置为 STORE 或 SHOWCASE。
	// shopping_ads_type 设置为 LIVE。
	OptimizationGoal_CONVERT OptimizationGoal = "CONVERT"
	// OptimizationGoal_INSTALL 安装。
	// 若要在 Pangle 版位或全球化应用组合版位的广告组中使用优化目标“安装与应用内事件数据”，需将 optimization_goal 设置为INSTALL，传入 secondary_optimization_event 合法值，且在placements的值中仅包含 PLACEMENT_PANGLE或 PLACEMENT_GLOBAL_APP_BUNDLE 或PLACEMENT_PANGLE和 PLACEMENT_GLOBAL_APP_BUNDLE。
	//
	// 注意：
	//
	// 自 2024 年 11 月 30 日起，您将无法创建或复制优化目标为“安装与应用内事件数据”且版位为仅 TikTok 或自动版位的广告组。此变动将影响使用以下配置的广告组：
	// 推广系列层级：objective_type 为 APP_PROMOTION
	// 广告组层级：
	// optimization_goal 为 INSTALL 且指定了 secondary_optimization_event 合法值
	// placement_type 为 PLACEMENT_TYPE_NORMAL 且 placements 为["PLACEMENT_TIKTOK"]，或 placement_type 为PLACEMENT_TYPE_AUTOMATIC
	// 现有的投放在 TikTok 版位的“安装与应用内事件数据”广告组将不受影响。此外，Pangle 版位和全球化应用组合版位（即 placement_type 为 PLACEMENT_TYPE_NORMAL 且 placements 的值仅包含PLACEMENT_PANGLE 和 PLACEMENT_GLOBAL_APP_BUNDLE）也不受影响。请注意此变更，并在必要时对您的集成进行适当的调整。
	OptimizationGoal_INSTALL OptimizationGoal = "INSTALL"
	// OptimizationGoal_IN_APP_EVENT 应用内事件
	OptimizationGoal_IN_APP_EVENT OptimizationGoal = "IN_APP_EVENT"
	// OptimizationGoal_SHOW 展示
	OptimizationGoal_SHOW OptimizationGoal = "SHOW"
	// OptimizationGoal_VIDEO_VIEW 视频播放
	OptimizationGoal_VIDEO_VIEW OptimizationGoal = "VIDEO_VIEW"
	// OptimizationGoal_REACH 覆盖人数
	OptimizationGoal_REACH OptimizationGoal = "REACH"
	// OptimizationGoal_LEAD_GENERATION 线索量
	OptimizationGoal_LEAD_GENERATION OptimizationGoal = "LEAD_GENERATION"
	// OptimizationGoal_CONVERSATION 会话数。向最有可能发起 TikTok 私信会话或即时通讯应用会话的用户展示你的广告。
	// 本优化目标与CLICK 相比，更注重深层次的优化，因此有更大机会实现更好的表现。
	//
	// 若想了解在 TikTok 私信广告中使用优化目标会话数的步骤，请查看创建优化位置为 TikTok 私信的线索广告。
	// 若想了解在 TikTok 即时通讯广告中使用优化目标会话数的步骤，请查看创建优化位置为即时通讯应用的线索广告。
	//
	// 注意：
	//
	// 优化目标会话数（CONVERSATION）目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 您将 optimization_goal 设置为 CONVERSATION 后， 可不传入 optimization_event，optimization_event 将默认设置为 MESSAGE。若您在optimization_event 字段手动传入其他值，则将产生报错。
	OptimizationGoal_CONVERSATION OptimizationGoal = "CONVERSATION"
	// OptimizationGoal_FOLLOWERS 账号关注
	OptimizationGoal_FOLLOWERS OptimizationGoal = "FOLLOWERS"
	// OptimizationGoal_PROFILE_VIEWS 主页访问
	// 注意: 我们计划在新版本API中废弃 PROFILE_VIEWS。为确保集成过程顺利，我们建议您将 PAGE_VISIT 用作优化目标。
	OptimizationGoal_PROFILE_VIEWS OptimizationGoal = "PROFILE_VIEWS"
	// OptimizationGoal_PAGE_VISITS TikTok 页面访问。为各种 TikTok 页面吸引访问量，包括账号主页、其他 TikTok 页面（播放列表页面、话题标签页面、音乐页面或商业化特效页面）、自定义 TikTok 即时体验页面。
	// 只有在推广系列层级的 objective_type 为 ENGAGEMENT（社区互动）时，才能使用此优化目标。
	// 若想了解如何在社区互动这一推广目标下使用优化目标 TikTok 页面访问，请查看创建社区互动广告。
	OptimizationGoal_PAGE_VISITS OptimizationGoal = "PAGE_VISITS"
	// OptimizationGoal_VALUE 价值/总收入。
	// VALUE 在以下任意场景中对应优化目标总收入：
	// shopping_ads_type 设置为 VIDEO 且 product_source 设置为 STORE 或 SHOWCASE。
	// shopping_ads_type 设置为 LIVE。
	// shopping_ads_type 设置为 PRODUCT_SHOPPING_ADS 且 product_source 设置为 STORE。
	// 若想了解如何使用价值这一优化目标，请查看开启价值优化。
	OptimizationGoal_VALUE OptimizationGoal = "VALUE"
	// OptimizationGoal_AUTOMATIC_VALUE_OPTIMIZATION 自动价值优化。通过机器学习进行动态优化，以实现最大价值或最大投放量。
	// 若想了解如何在网站转化量推广目标下使用优化目标自动价值优化，请查看开启网站价值优化。
	// 注意：
	// 自动价值优化目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 同时启用自动价值优化和推广系列预算优化目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	OptimizationGoal_AUTOMATIC_VALUE_OPTIMIZATION OptimizationGoal = "AUTOMATIC_VALUE_OPTIMIZATION"
	// OptimizationGoal_GMV 一个组合筛选器，筛选出推广目标（objective_type）为 PRODUCT_SALES、优化目标（optimization_goal）为 VALUE、推广系列的商品来源（campaign_product_source）为 STORE，以及购物广告类型（shopping_ads_type）为 VIDEO/LIVE/PRODUCT_SHOPPING_ADS 的推广系列、广告组和广告。
	// 注意: 该优化目标不能用于创建广告组，只能在获取推广系列/广告组/广告的GET 接口和报表管理的接口中筛选时使用。
	OptimizationGoal_GMV OptimizationGoal = "GMV"
	// OptimizationGoal_PURCHASES 一个组合筛选器，筛选出推广目标（objective_type）为 PRODUCT_SALES、优化目标（optimization_goal）为 CONVERT、推广系列的商品来源（campaign_product_source）为 STORE，以及匹配的事件（optimization_event）为 SHOPPING 的推广系列、广告组和广告。
	// 注意: 该优化目标不能用于创建广告组，只能在获取推广系列/广告组/广告的GET 接口和报表管理的接口中筛选时使用。
	OptimizationGoal_PURCHASES OptimizationGoal = "PURCHASES"
	// OptimizationGoal_INITIATE_CHECKOUTS 一个组合筛选器，筛选出推广目标（objective_type）为 PRODUCT_SALES、优化目标（optimization_goal）为 CONVERT、推广系列的商品来源（campaign_product_source）为 STORE，以及匹配的事件（optimization_event）为 INITIATE_ORDER 的推广系列、广告组和广告。
	// 注意:该优化目标不能用于创建广告组，只能在获取推广系列/广告组/广告的GET 接口和报表管理的接口中筛选时使用。
	OptimizationGoal_INITIATE_CHECKOUTS OptimizationGoal = "INITIATE_CHECKOUTS"
	// OptimizationGoal_MT_LIVE_ROOM 直播间观众留存。本优化目标只适用于直播带货场景 (promotion_type = LIVE_SHOPPING)。匹配的事件（optimization_event）为 LIVE_STAY_TIME
	OptimizationGoal_MT_LIVE_ROOM OptimizationGoal = "MT_LIVE_ROOM"
	// OptimizationGoal_PRODUCT_CLICK_IN_LIVE LIVE商品点击数。本优化目标只适用于直播带货场景(promotion_type = LIVE_SHOPPING)。匹配的事件(optimization_event)为 LIVE_CLICK_PRODUCT_ACTION
	OptimizationGoal_PRODUCT_CLICK_IN_LIVE OptimizationGoal = "PRODUCT_CLICK_IN_LIVE"
	// OptimizationGoal_ENGAGED_VIEW 观看6秒次数（专注观看）。
	// 用户观看视频时长达到6秒，或观看付费广告的一天内与视频互动至少一次。纳入统计的互动包括：点赞，分享，关注，访问主页，点击，话题标签点击，音乐点击，锚点点击，以及创新互动样式行为点击。
	// 注意：
	//
	// 广告组创建后，将无法把optimization_goal从ENGAGED_VIEW更新为其他值。
	// 将optimization_goal设置为ENGAGED_VIEW后，系统将自动将 ENGAGED_VIEW 传入optimization_event。若已经将 optimization_event设置为其他值，则传入值会被忽略。
	OptimizationGoal_ENGAGED_VIEW OptimizationGoal = "ENGAGED_VIEW"
	// OptimizationGoal_ENGAGED_VIEW_FIFTEEN 观看15秒次数（专注观看）。
	// 用户观看视频时长达到15秒，或观看付费广告的一天内与视频互动至少一次。纳入统计的互动包括：点赞，分享，关注，访问主页，点击，话题标签点击，音乐点击，锚点点击，以及创新互动样式行为点击。
	// 注意：
	//
	// ENGAGED_VIEW_FIFTEEN为白名单功能，若想使用该功能，请联系您的TikTok销售代表。
	// 广告组创建后，将无法把optimization_goal从ENGAGED_VIEW_FIFTEEN更新为其他值。
	// 将optimization_goal设置为ENGAGED_VIEW_FIFTEEN后，系统将自动将 ENGAGED_VIEW_FIFTEEN 传入optimization_event。若已经将 optimization_event设置为其他值，则传入值会被忽略。
	// 将optimization_goal设置为ENGAGED_VIEW_FIFTEEN后，schedule_type仅支持设置为SCHEDULE_START_END，且通过 schedule_start_time 和 schedule_end_time 定义的广告组排期不可超过 30 天。
	OptimizationGoal_ENGAGED_VIEW_FIFTEEN OptimizationGoal = "ENGAGED_VIEW_FIFTEEN"
	// OptimizationGoal_TRAFFIC_LANDING_PAGE_VIEW 落地页浏览量。用户点击并成功加载落地页。对应的optimization_event为LANDING_PAGE_VIEW。
	// 若要在访问量目标下使用 TRAFFIC_LANDING_PAGE_VIEW，需满足以下条件：
	// 推广系列层级：objective_type 设置为 TRAFFIC
	// 广告组层级：
	// promotion_type 设置为 WEBSITE
	// 版位：
	// placement_type设置为 PLACEMENT_TYPE_AUTOMATIC 或
	// placement_type设置为 PLACEMENT_TYPE_NORMAL，且 placements 未设置为 ["PLACEMENT_GLOBAL_APP_BUNDLE"]（仅包含全球化应用组合）
	// optimization_goal 设置为 TRAFFIC_LANDING_PAGE_VIEW
	// 广告层级：
	// 需通过page_id传入TikTok 即时体验页面的 ID。
	// 未设置深度链接（直达链接）。
	//
	// 若想了解如何商品销量目标下使用落地页浏览量优化目标，请查看优化视频购物广告中的落地页浏览量。
	//
	// 注意：
	//
	// 在购物广告中使用优化目标落地页浏览量目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 您将 optimization_goal 设置为 TRAFFIC_LANDING_PAGE_VIEW 后， 可不传入 optimization_event，optimization_event 将默认设置为 LANDING_PAGE_VIEW。若您在optimization_event 字段手动传入其他值，则将产生报错。
	OptimizationGoal_TRAFFIC_LANDING_PAGE_VIEW OptimizationGoal = "TRAFFIC_LANDING_PAGE_VIEW"
	// OptimizationGoal_CONVERSION_LEADS 转化线索量。
	// 向最有可能在提供联系信息后完成转化的用户展示广告
	OptimizationGoal_CONVERSION_LEADS OptimizationGoal = "CONVERSION_LEADS"
	// OptimizationGoal_DESTINATION_VISIT 目标页面访问。使用此优化目标的广告会将用户引导至应用内页面，或回退到网站。
	// 若想了解如何在访问量目标下使用目标页面访问，请查看优化访问量广告中的目标页面访问。
	// 若想了解如何在商品销量目标下使用目标页面访问，请查看优化视频购物广告中的目标页面访问。
	//
	// 注意：
	//
	// 优化目标目标页面访问（DESTINATION_VISIT）目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 您将 optimization_goal 设置为 DESTINATION_VISIT 后， 可不传入 optimization_event，optimization_event 将默认设置为 DESTINATION_VISIT。若您在optimization_event 字段手动传入其他值，则将忽略传入值。
	OptimizationGoal_DESTINATION_VISIT OptimizationGoal = "DESTINATION_VISIT"
	// OptimizationGoal_PREFERRED_LEAD 偏好的潜在客户。向最有可能符合你偏好的潜在客户资质的人群展示广告。
	OptimizationGoal_PREFERRED_LEAD OptimizationGoal = "PREFERRED_LEAD"
)
