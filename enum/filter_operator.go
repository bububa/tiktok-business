package enum

// FilterSetOperator 筛选条件间的操作符，AND 代表同时满足筛选关系， OR 代表任意一项满足，目前仅支持 OR
type FilterSetOperator string

const (
	// AND 代表同时满足筛选关系
	AND FilterSetOperator = "AND"
	// OR 代表任意一项满足
	OR FilterSetOperator = "OR"
)

// FilterOperator 连接筛选字段和筛选字段值的筛选操作符
type FilterOperator string

const (
	// EQ（等于）
	EQ FilterOperator = "EQ"
	// GT（大于）
	GT FilterOperator = "GT"
	// LT（小于）
	LT FilterOperator = "LT"
	// LTE: 小于或等于。
	LTE FilterOperator = "LTE"
	// GTE: 大于或等于。
	GTE FilterOperator = "GTE"
	// CONTAIN: 包含。
	CONTAIN FilterOperator = "CONTAIN"
	// NOT_CONTAIN: 不包含。
	NOT_CONTAIN FilterOperator = "NOT_CONTAIN"
	// START_WITH: 开头是。
	START_WITH FilterOperator = "START_WITH"
	// EQUAL: 等于。
	EQUAL FilterOperator = "EQUAL"
	// NOT_EQUAL: 不等于。
	NOT_EQUAL FilterOperator = "NOT EQUAL"
)

// FilterValue 筛选字段值
type FilterValue string

const (
	// FilterValue_CLICK 视频广告点击事件
	FilterValue_CLICK FilterValue = "CLICK"
	// FilterValue_IMPRESSION 视频广告展示事件
	FilterValue_IMPRESSION FilterValue = "IMPRESSION"
	// FilterValue_PLAY_2S 视频广告播放2秒
	FilterValue_PLAY_2S FilterValue = "PLAY 2S"
	// FilterValue_PLAY_6S 视频广告播放6秒
	FilterValue_PLAY_6S FilterValue = "PLAY 6S"
	// FilterValue_PLAY_25 视频广告播放进度25%
	FilterValue_PLAY_25 FilterValue = "PLAY 25"
	// FilterValue_PLAY_50 视频广告播放进度50%
	FilterValue_PLAY_50 FilterValue = "PLAY 50"
	// FilterValue_PLAY_75 视频广告播放进度75%
	FilterValue_PLAY_75 FilterValue = "PLAY 75"
	// FilterValue_PLAY_OVER 视频广告播放完成
	FilterValue_PLAY_OVER FilterValue = "PLAY OVER"
	// FilterValue_ENGAGEMENT_APP_PROFILE_CLICK 应用介绍页广告点击事件
	FilterValue_ENGAGEMENT_APP_PROFILE_CLICK FilterValue = "ENGAGEMENT APP PROFILE CLICK"
	// FilterValue_ENGAGEMENT_APP_PROFILE_IMPRESSION 应用介绍页广告展示事件
	FilterValue_ENGAGEMENT_APP_PROFILE_IMPRESSION FilterValue = "ENGAGEMENT APP PROFILE IMPRESSION"
	// FilterValue_ENGAGEMENT_TIKTOK_INSTANT_CLICK 即时页面广告点击事件
	FilterValue_ENGAGEMENT_TIKTOK_INSTANT_CLICK FilterValue = "ENGAGEMENT TIKTOK INSTANT CLICK"
	// FilterValue_ENGAGEMENT_TIKTOK_INSTANT_IMPRESSION 即时页面广告展示事件
	FilterValue_ENGAGEMENT_TIKTOK_INSTANT_IMPRESSION FilterValue = "ENGAGEMENT TIKTOK INSTANT IMPRESSION"
	// FilterValue_ENGAGEMENT_COLLECTION_ADS_CLICK 精品栏广告点击事件
	FilterValue_ENGAGEMENT_COLLECTION_ADS_CLICK FilterValue = "ENGAGEMENT COLLECTION ADS CLICK"
	// FilterValue_ENGAGEMENT_COLLECTION_ADS_IMPRESSION 精品栏广告展示事件
	FilterValue_ENGAGEMENT_COLLECTION_ADS_IMPRESSION FilterValue = "ENGAGEMENT COLLECTION ADS IMPRESSION"
	// FilterValue_LIVE_VIDEO_VIEW 观看
	FilterValue_LIVE_VIDEO_VIEW FilterValue = "LIVE VIDEO VIEW"
	// FilterValue_LIVE_VIDEO_ENGAGEMENT 互动
	FilterValue_LIVE_VIDEO_ENGAGEMENT FilterValue = "LIVE VIDEO ENGAGEMENT"
	// FilterValue_INSTALL 安装事件
	FilterValue_INSTALL FilterValue = "INSTALL"
	// FilterValue_DAY_2_RETENTION 次日留存
	FilterValue_DAY_2_RETENTION FilterValue = "DAY 2 RETENTION"
	// FilterValue_APP_ADD_PAYMENT_INFO 添加支付信息事件
	FilterValue_APP_ADD_PAYMENT_INFO FilterValue = "APP ADD PAYMENT INFO"
	// FilterValue_APP_ADD_TO_CART 应用中添加到购物车事件，请注意本事件区别于Pixel中添加到购物车事件
	FilterValue_APP_ADD_TO_CART FilterValue = "APP ADD TO CART"
	// FilterValue_ADD_TO_WISHLIST 添加到购物清单事件
	FilterValue_ADD_TO_WISHLIST FilterValue = "ADD TO WISHLIST"
	// FilterValue_LAUNCH_APP 启动应用事件
	FilterValue_LAUNCH_APP FilterValue = "LAUNCH APP"
	// FilterValue_CHECKOUT 下单事件
	FilterValue_CHECKOUT FilterValue = "CHECKOUT"
	// FilterValue_COMPLETE_TUTORIAL 完成导引事件
	FilterValue_COMPLETE_TUTORIAL FilterValue = "COMPLETE TUTORIAL"
	// FilterValue_VIEW_CONTENT 查看内容事件
	FilterValue_VIEW_CONTENT FilterValue = "VIEW CONTENT"
	// FilterValue_CREATE_GROUP 创建群组事件
	FilterValue_CREATE_GROUP FilterValue = "CREATE GROUP"
	// FilterValue_CREATE_ROLE 创建角色事件
	FilterValue_CREATE_ROLE FilterValue = "CREATE ROLE"
	// FilterValue_GENERATE_LEAD 生成线索
	FilterValue_GENERATE_LEAD FilterValue = "GENERATE LEAD"
	// FilterValue_IN_APP_AD_CLICK 应用内点击广告事件
	FilterValue_IN_APP_AD_CLICK FilterValue = "IN-APP AD CLICK"
	// FilterValue_IN_APP_AD_IMPRESSION 应用内广告展示事件
	FilterValue_IN_APP_AD_IMPRESSION FilterValue = "IN-APP AD IMPRESSION"
	// FilterValue_JOIN_GROUP 加入群组事件
	FilterValue_JOIN_GROUP FilterValue = "JOIN GROUP"
	// FilterValue_ACHIEVE_LEVEL 达到等级
	FilterValue_ACHIEVE_LEVEL FilterValue = "ACHIEVE LEVEL"
	// FilterValue_LOAN_APPLY 申请贷款
	FilterValue_LOAN_APPLY FilterValue = "LOAN APPLY"
	// FilterValue_LOAN APPROVAL 贷款获批
	FilterValue_LOAN_APPROVAL FilterValue = "LOAN APPROVAL"
	// FilterValue_LOAN_DISBURSAL 放款
	FilterValue_LOAN_DISBURSAL FilterValue = "LOAN DISBURSAL"
	// FilterValue_LOGIN 登录
	FilterValue_LOGIN FilterValue = "LOGIN"
	// FilterValue_PURCHASE 购买
	FilterValue_PURCHASE FilterValue = "PURCHASE"
	// FilterValue_RATE 评价
	FilterValue_RATE FilterValue = "RATE"
	// FilterValue_REGISTRATION 注册
	FilterValue_REGISTRATION FilterValue = "REGISTRATION"
	// FilterValue_APP_SEARCH 搜索
	FilterValue_APP_SEARCH FilterValue = "APP SEARCH"
	// FilterValue_SPEND_CREDITS 使用积分
	FilterValue_SPEND_CREDITS FilterValue = "SPEND CREDITS"
	// FilterValue_START_TRIAL 开始试用
	FilterValue_START_TRIAL FilterValue = "START TRIAL"
	// FilterValue_SUBSCRIBE 注册
	FilterValue_SUBSCRIBE FilterValue = "SUBSCRIBE"
	// FilterValue_UNLOCK_ACHIEVEMENT 解锁成就
	FilterValue_UNLOCK_ACHIEVEMENT FilterValue = "UNLOCK ACHIEVEMENT"
	// FilterValue_PAGE_BROWSE 对应“网页浏览事件”。访客何时登陆你的网站。
	FilterValue_PAGE_BROWSE FilterValue = "PAGE BROWSE"
	// FilterValue_CLICK_BUTTON 点击按钮
	FilterValue_CLICK_BUTTON FilterValue = "CLICK BUTTON"
	// FilterValue_PIXEL_SUBMIT_FORM 表单提交
	FilterValue_PIXEL_SUBMIT_FORM FilterValue = "PIXEL SUBMIT FORM"
	// FilterValue_CONTACT 问询
	FilterValue_CONTACT FilterValue = "CONTACT"
	// FilterValue_DOWNLOAD 点击下载按钮
	FilterValue_DOWNLOAD FilterValue = "DOWNLOAD"
	// FilterValue_PIXEL_ADD_PAYMENT_INFO 添加付款信息
	FilterValue_PIXEL_ADD_PAYMENT_INFO FilterValue = "PIXEL ADD PAYMENT INFO"
	// FilterValue_COMPLETE_PAYMENT 完成付款
	FilterValue_COMPLETE_PAYMENT FilterValue = "COMPLETE PAYMENT"
	// FilterValue_INITIATE_CHECKOUT 开始付款
	FilterValue_INITIATE_CHECKOUT FilterValue = "INITIATE CHECKOUT"
	// FilterValue_COMPLETE_REGISTRATION 完成注册
	FilterValue_COMPLETE_REGISTRATION FilterValue = "COMPLETE REGISTRATION"
	// FilterValue_PRODUCT_DETAIL_PAGE_BROWSE 对应“查看内容”事件。用户何时查看了详细的产品页面。
	FilterValue_PRODUCT_DETAIL_PAGE_BROWSE FilterValue = "PRODUCT DETAIL PAGE BROWSE"
	// FilterValue_PIXEL_SEARCH 搜索
	FilterValue_PIXEL_SEARCH FilterValue = "PIXEL SEARCH"
	// FilterValue_PIXEL_ADD_TO_CART 添加到购物车（Pixel）
	FilterValue_PIXEL_ADD_TO_CART FilterValue = "PIXEL ADD TO CART"
	// FilterValue_PLACE_AN_ORDER 下单
	FilterValue_PLACE_AN_ORDER FilterValue = "PLACE AN ORDER"
	// FilterValue_PIXEL_ADD_TO_WISHLIST 添加到购物清单（Pixel）
	FilterValue_PIXEL_ADD_TO_WISHLIST FilterValue = "PIXEL ADD TO WISHLIST"
	// FilterValue_PIXEL_SUBSCRIBE 订阅（Pixel）
	FilterValue_PIXEL_SUBSCRIBE FilterValue = "PIXEL SUBSCRIBE"
	// FilterValue_LEAD_GENERATION_FORM_SUBMISSION 提交表单
	FilterValue_LEAD_GENERATION_FORM_SUBMISSION FilterValue = "LEAD GENERATION FORM SUBMISSION"
	// FilterValue_FORM_VIEW 浏览表单
	FilterValue_FORM_VIEW FilterValue = "FORM VIEW"
	// FilterValue_TIKTOK_SHOP_PAGE_VIEW 产品详情页浏览事件
	FilterValue_TIKTOK_SHOP_PAGE_VIEW FilterValue = "TIKTOK SHOP PAGE VIEW"
	// FilterValue_TIKTOK_SHOP_ADD_TO_CART 添加到购物车事件
	FilterValue_TIKTOK_SHOP_ADD_TO_CART FilterValue = "TIKTOK SHOP ADD TO CART"
	// FilterValue_TIKTOK_SHOP_INITIATE_CHECKOUT 开始下单事件
	FilterValue_TIKTOK_SHOP_INITIATE_CHECKOUT FilterValue = "TIKTOK SHOP INITIATE CHECKOUT"
	// FilterValue_TIKTOK_SHOP_COMPLETE_PAYMENT 支付完成事件
	FilterValue_TIKTOK_SHOP_COMPLETE_PAYMENT FilterValue = "TIKTOK SHOP COMPLETE PAYMENT"
	// FilterValue_BUSINESS_ACCOUNT_PROFILE_FOLLOW 关注事件
	FilterValue_BUSINESS_ACCOUNT_PROFILE_FOLLOW FilterValue = "BUSINESS ACCOUNT PROFILE FOLLOW"
	// FilterValue_BUSINESS_ACCOUNT_PROFILE_VISIT 主页访问事件
	FilterValue_BUSINESS_ACCOUNT_PROFILE_VISIT FilterValue = "BUSINESS ACCOUNT PROFILE VISIT"
	// FilterValue_BUSINESS_ACCOUNT_ENGAGEMENT 视频互动事件（点赞/分享/评论）
	FilterValue_BUSINESS_ACCOUNT_ENGAGEMENT FilterValue = "BUSINESS ACCOUNT ENGAGEMENT"
	// FilterValue_BUSINESS_ACCOUNT_PLAY_2S 视频播放2秒事件
	FilterValue_BUSINESS_ACCOUNT_PLAY_2S FilterValue = "BUSINESS ACCOUNT PLAY 2S"
	// FilterValue_BUSINESS_ACCOUNT_PLAY_6S 视频播放6秒事件
	FilterValue_BUSINESS_ACCOUNT_PLAY_6S FilterValue = "BUSINESS ACCOUNT PLAY 6S"
	// FilterValue_BUSINESS_ACCOUNT_PLAY_OVER 视频播放完成事件
	FilterValue_BUSINESS_ACCOUNT_PLAY_OVER FilterValue = "BUSINESS ACCOUNT PLAY OVER"
	// FilterValue_OFFLINE_COMPLETE_PAYMENT 完成付款
	FilterValue_OFFLINE_COMPLETE_PAYMENT FilterValue = "OFFLINE COMPLETE PAYMENT"
	// FilterValue_OFFLINE_CONTACT 联系
	FilterValue_OFFLINE_CONTACT FilterValue = "OFFLINE CONTACT"
	// FilterValue_OFFLINE_SUBMIT_FORM 提交表单
	FilterValue_OFFLINE_SUBMIT_FORM FilterValue = "OFFLINE SUBMIT FORM"
	// FilterValue_OFFLINE_SUBSCRIBE 订阅
	FilterValue_OFFLINE_SUBSCRIBE FilterValue = "OFFLINE SUBSCRIBE"
	// FilterValue_OFFLINE_ADD_PAYMENT_INFO 添加支付信息
	FilterValue_OFFLINE_ADD_PAYMENT_INFO FilterValue = "OFFLINE ADD PAYMENT INFO"
	// FilterValue_OFFLINE_ADD_TO_CART 加入购物车
	FilterValue_OFFLINE_ADD_TO_CART FilterValue = "OFFLINE ADD TO CART"
	// FilterValue_OFFLINE_ADD_TO_WISHLIST 加入心愿单
	FilterValue_OFFLINE_ADD_TO_WISHLIST FilterValue = "OFFLINE ADD TO WISHLIST"
	// FilterValue_OFFLINE_CLICK_BUTTON 点击按钮
	FilterValue_OFFLINE_CLICK_BUTTON FilterValue = "OFFLINE CLICK BUTTON"
	// FilterValue_OFFLINE_COMPLETE_REGISTRATION 完成注册
	FilterValue_OFFLINE_COMPLETE_REGISTRATION FilterValue = "OFFLINE COMPLETE REGISTRATION"
	// FilterValue_OFFLINE_DOWNLOAD 下载
	FilterValue_OFFLINE_DOWNLOAD FilterValue = "OFFLINE DOWNLOAD"
	// FilterValue_OFFLINE_INITIATE_CHECKOUT 开始下单
	FilterValue_OFFLINE_INITIATE_CHECKOUT FilterValue = "OFFLINE INITIATE CHECKOUT"
	// FilterValue_OFFLINE_PLACE_AN_ORDER 下单事件
	FilterValue_OFFLINE_PLACE_AN_ORDER FilterValue = "OFFLINE PLACE AN ORDER"
)
