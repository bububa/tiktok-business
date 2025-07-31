package enum

// OptimizationEvent 转化事件
type OptimizationEvent string

const (
	//	应用事件类型
	//
	// 下表列出了推广系列管理中应用转化事件（optimization_event）的枚举值。
	// OptimizationEvent_ACTIVE 安装
	OptimizationEvent_ACTIVE OptimizationEvent = "ACTIVE"
	// OptimizationEvent_ACTIVE_PAY 购买
	OptimizationEvent_ACTIVE_PAY OptimizationEvent = "ACTIVE_PAY"
	// OptimizationEvent_ACTIVE_REGISTER 注册
	OptimizationEvent_ACTIVE_REGISTER OptimizationEvent = "ACTIVE_REGISTER"
	// OptimizationEvent_ADD_TO_WISHLIST 加入心愿单
	OptimizationEvent_ADD_TO_WISHLIST OptimizationEvent = "ADD_TO_WISHLIST"
	// OptimizationEvent_COMPLETE_TUTORIAL 完成教程
	OptimizationEvent_COMPLETE_TUTORIAL OptimizationEvent = "COMPLETE_TUTORIAL"
	// OptimizationEvent_CREATE_GAMEROLE 创建角色
	OptimizationEvent_CREATE_GAMEROLE OptimizationEvent = "CREATE_GAMEROLE"
	// OptimizationEvent_CREATE_GROUP 创建组
	OptimizationEvent_CREATE_GROUP OptimizationEvent = "CREATE_GROUP"
	// OptimizationEvent_IN_APP_AD_CLICK 应用内广告点击
	OptimizationEvent_IN_APP_AD_CLICK OptimizationEvent = "IN_APP_AD_CLICK"
	// OptimizationEvent_IN_APP_AD_IMPR 应用内广告展示
	OptimizationEvent_IN_APP_AD_IMPR OptimizationEvent = "IN_APP_AD_IMPR"
	// OptimizationEvent_IN_APP_CART 加入购物车
	OptimizationEvent_IN_APP_CART OptimizationEvent = "IN_APP_CART"
	// OptimizationEvent_IN_APP_DETAIL_UV 查看内容
	OptimizationEvent_IN_APP_DETAIL_UV OptimizationEvent = "IN_APP_DETAIL_UV"
	// OptimizationEvent_JOIN_GROUP 加入组
	OptimizationEvent_JOIN_GROUP OptimizationEvent = "JOIN_GROUP"
	// OptimizationEvent_LAUNCH_APP 启动应用
	OptimizationEvent_LAUNCH_APP OptimizationEvent = "LAUNCH_APP"
	// OptimizationEvent_LIVE_CLICK_PRODUCT_ACTION 直播带货期间商品链接的点击数。只在 promotion_type（推广对象）为LIVE_SHOPPING（直播带货）且 optimization_goal（优化目标）为PRODUCT_CLICK_IN_LIVE时有效。
	// 注意：LIVE_CLICK_PRODUCT_ACTION目前为白名单功能，如需使用此功能，请联系您的TikTok销售代表。
	OptimizationEvent_LIVE_CLICK_PRODUCT_ACTION OptimizationEvent = "LIVE_CLICK_PRODUCT_ACTION"
	// OptimizationEvent_LIVE_STAY_TIME 直播间的观众停留时间。只在 promotion_type（推广对象）为直播带货（LIVE_SHOPPING）且 optimization_goal（优化目标）为MT_LIVE_ROOM时有效。
	// 注意：LIVE_STAY_TIME目前为白名单功能，如需使用此功能，请联系您的TikTok销售代表。
	OptimizationEvent_LIVE_STAY_TIME OptimizationEvent = "LIVE_STAY_TIME"
	// OptimizationEvent_LOAN_COMPLETION 贷款发放
	OptimizationEvent_LOAN_COMPLETION OptimizationEvent = "LOAN_COMPLETION"
	// OptimizationEvent_LOGIN 登陆
	OptimizationEvent_LOGIN OptimizationEvent = "LOGIN"
	// OptimizationEvent_NEXT_DAY_OPEN 次日留存
	OptimizationEvent_NEXT_DAY_OPEN OptimizationEvent = "NEXT_DAY_OPEN"
	// OptimizationEvent_RATE 评分
	OptimizationEvent_RATE OptimizationEvent = "RATE"
	// OptimizationEvent_SALES_LEAD 收集线索
	OptimizationEvent_SALES_LEAD OptimizationEvent = "SALES_LEAD"
	// OptimizationEvent_SEARCH 搜索
	OptimizationEvent_SEARCH OptimizationEvent = "SEARCH"
	// OptimizationEvent_SPEND_CREDITS 花费点数
	OptimizationEvent_SPEND_CREDITS OptimizationEvent = "SPEND_CREDITS"
	// OptimizationEvent_UNLOCK_ACHIEVEMENT 解锁关卡
	OptimizationEvent_UNLOCK_ACHIEVEMENT OptimizationEvent = "UNLOCK_ACHIEVEMENT"
	// OptimizationEvent_UPDATE_LEVEL 达到级别
	OptimizationEvent_UPDATE_LEVEL OptimizationEvent = "UPDATE_LEVEL"
	// OptimizationEvent_ENGAGED_VIEW 仅当 optimization_goal（优化目标）为ENGAGED_VIEW时生效。
	// 观看6秒次数（专注观看）。
	// 用户观看视频时长达到6秒，或观看付费广告的一天内与视频互动至少一次。纳入统计的互动包括：点赞，分享，关注，访问主页，点击，话题标签点击，音乐点击，锚点点击，以及创新互动样式行为点击。
	OptimizationEvent_ENGAGED_VIEW OptimizationEvent = "ENGAGED_VIEW"
	// OptimizationEvent_ENGAGED_VIEW_FIFTEEN 仅当 optimization_goal（优化目标）为ENGAGED_VIEW_FIFTEEN时生效。
	// 观看15秒次数（专注观看）。
	// 用户观看视频时长达到15秒，或观看付费广告的一天内与视频互动至少一次。纳入统计的互动包括：点赞，分享，关注，访问主页，点击，话题标签点击，音乐点击，锚点点击，以及创新互动样式行为点击。
	// 注意：ENGAGED_VIEW_FIFTEEN目前为白名单功能，如需使用此功能，请联系您的TikTok销售代表。
	OptimizationEvent_ENGAGED_VIEW_FIFTEEN OptimizationEvent = "ENGAGED_VIEW_FIFTEEN"
	// OptimizationEvent_LANDING_PAGE_VIEW 落地页浏览量。用户点击并成功加载落地页。
	// 转化事件 LANDING_PAGE_VIEW 仅在访问量和商品销量目标下支持。
	// 若要在访问量目标下使用 LANDING_PAGE_VIEW，请查看枚举值 - 优化目标。
	// 若要在商品销量目标下使用 LANDING_PAGE_VIEW ，请查看优化视频购物广告中的落地页浏览量。
	// 注意：将推广系列的objective_type 设置为 TRAFFIC，且将广告组的optimization_event设置为LANDING_PAGE_VIEW后，不允许将广告组与拆分对比测试绑定。
	OptimizationEvent_LANDING_PAGE_VIEW OptimizationEvent = "LANDING_PAGE_VIEW"
	// OptimizationEvent_PAGE_VISIT TikTok 简介页面访问量和其他TikTok页面访问量（包括播放列表页面、标签页、音乐页面等）。
	// 注意：当optimization_goal为 PAGE_VISIT 时，optimization_event 将自动设置为 PAGE_VISIT。
	OptimizationEvent_PAGE_VISIT OptimizationEvent = "PAGE_VISIT"
	// OptimizationEvent_MESSAGE 仅当 promotion_type 设置为 LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE 或 LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE 且 optimization_goal 为 CONVERSATION 时生效。
	// 通过 TikTok 私信或即时通讯应用消息收集线索。
	// 注意： 若 promotion_type 设置为 LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE 或 LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE 且 optimization_goal 为 CONVERSATION， optimization_event可不传入，该字段将自动设置为 MESSAGE
	OptimizationEvent_MESSAGE OptimizationEvent = "MESSAGE"
	// OptimizationEvent_DESTINATION_VISIT 仅当optimization_goal 设置为 DESTINATION_VISIT 时生效。
	// 目标页面访问。使用此优化目标的广告会将用户引导至应用内页面，或回退到网站。
	// 注意：
	// 优化目标目标页面访问（DESTINATION_VISIT）目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 您将 optimization_goal 设置为 DESTINATION_VISIT 后， 可不传入 optimization_event，optimization_event 将默认设置为 DESTINATION_VISIT。若您在optimization_event 字段手动传入其他值，则将忽略传入值。
	OptimizationEvent_DESTINATION_VISIT OptimizationEvent = "DESTINATION_VISIT"
	// OptimizationEvent_AD_REVENUE_VALUE 广告收入
	// 仅支持广告收入价值优化场景。
	// 广告收入价值。尽可能增加你通过应用内广告获得的收入。
	// 若想了解广告收入价值优化详情，请查看价值优化 - 价值类型。
	OptimizationEvent_AD_REVENUE_VALUE OptimizationEvent = "AD_REVENUE_VALUE"
	//
	// Pixel事件类型
	// 下表列出了推广系列管理中网页转化事件（optimization_event）的枚举值。
	//
	// OptimizationEvent_ADD_BILLING 添加支付信息
	OptimizationEvent_ADD_BILLING OptimizationEvent = "ADD_BILLING"
	// OptimizationEvent_BUTTON 已废弃 点击按钮
	OptimizationEvent_BUTTON OptimizationEvent = "BUTTON"
	// OptimizationEvent_CONSULT 咨询
	OptimizationEvent_CONSULT OptimizationEvent = "CONSULT"
	// OptimizationEvent_DOWNLOAD_START 下载
	OptimizationEvent_DOWNLOAD_START OptimizationEvent = "DOWNLOAD_START"
	// OptimizationEvent_FORM 提交表单
	OptimizationEvent_FORM OptimizationEvent = "FORM"
	// OptimizationEvent_INITIATE_ORDER 开始结账。如果要在直播带货场景（promotion_type=LIVE_SHOPPING）下使用，本事件类型需要先开通白名单才能使用。在直播带货场景下，只有当优化目标（optimization_goal ）为转化（CONVERT）时本事件类型才有效。
	OptimizationEvent_INITIATE_ORDER OptimizationEvent = "INITIATE_ORDER"
	// OptimizationEvent_ON_WEB_ADD_TO_WISHLIST 加入心愿单
	OptimizationEvent_ON_WEB_ADD_TO_WISHLIST OptimizationEvent = "ON_WEB_ADD_TO_WISHLIST"
	// OptimizationEvent_ON_WEB_CART 加入购物车
	OptimizationEvent_ON_WEB_CART OptimizationEvent = "ON_WEB_CART"
	// OptimizationEvent_ON_WEB_DETAIL 查看内容
	OptimizationEvent_ON_WEB_DETAIL OptimizationEvent = "ON_WEB_DETAIL"
	// OptimizationEvent_ON_WEB_ORDER 已废弃 下单
	OptimizationEvent_ON_WEB_ORDER OptimizationEvent = "ON_WEB_ORDER"
	// OptimizationEvent_ON_WEB_REGISTER 完成注册
	OptimizationEvent_ON_WEB_REGISTER OptimizationEvent = "ON_WEB_REGISTER"
	// OptimizationEvent_ON_WEB_SEARCH 搜索
	OptimizationEvent_ON_WEB_SEARCH OptimizationEvent = "ON_WEB_SEARCH"
	// OptimizationEvent_ON_WEB_SUBSCRIBE 订阅
	OptimizationEvent_ON_WEB_SUBSCRIBE OptimizationEvent = "ON_WEB_SUBSCRIBE"
	// OptimizationEvent_SHOPPING 付费
	OptimizationEvent_SHOPPING OptimizationEvent = "SHOPPING"
)
