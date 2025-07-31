package enum

// SecondaryOptimizationEvent 深度转化事件类型
type SecondaryOptimizationEvent string

const (
	// SecondaryOptimizationEvent_ACTIVE_PAY 购买。用户在您的应用中完成了付款过程。 当optimization_goal为INSTALL时可用。
	SecondaryOptimizationEvent_ACTIVE_PAY SecondaryOptimizationEvent = "ACTIVE_PAY"
	// SecondaryOptimizationEvent_ACTIVE_REGISTER 注册成功。用户在您的应用中注册。 当optimization_goal为INSTALL时可用。
	SecondaryOptimizationEvent_ACTIVE_REGISTER SecondaryOptimizationEvent = "ACTIVE_REGISTER"
	// SecondaryOptimizationEvent_ADD_PAYMENT_INFO 已废弃 添加付款信息
	SecondaryOptimizationEvent_ADD_PAYMENT_INFO SecondaryOptimizationEvent = "ADD_PAYMENT_INFO"
	// SecondaryOptimizationEvent_IN_APP_CART 加入购物车。用户在您的应用中将商品添加到他们的收藏夹或心愿单中。 当optimization_goal为INSTALL时可用。
	SecondaryOptimizationEvent_IN_APP_CART SecondaryOptimizationEvent = "IN_APP_CART"
	// SecondaryOptimizationEvent_IN_APP_ORDER 已废弃 结账
	SecondaryOptimizationEvent_IN_APP_ORDER SecondaryOptimizationEvent = "IN_APP_ORDER"
	// SecondaryOptimizationEvent_LOAN_APPLY 已废弃 贷款申请
	SecondaryOptimizationEvent_LOAN_APPLY SecondaryOptimizationEvent = "LOAN_APPLY"
	// SecondaryOptimizationEvent_LOAN_CREDIT 已废弃 贷款批准
	SecondaryOptimizationEvent_LOAN_CREDIT SecondaryOptimizationEvent = "LOAN_CREDIT"
	// SecondaryOptimizationEvent_NEXT_DAY_OPEN 次日留存。 当optimization_goal为INSTALL时可用。
	SecondaryOptimizationEvent_NEXT_DAY_OPEN SecondaryOptimizationEvent = "NEXT_DAY_OPEN"
	// SecondaryOptimizationEvent_PURCHASE_ROI 购买操作的ROI。 当optimization_goal为VALUE时可用。
	SecondaryOptimizationEvent_PURCHASE_ROI SecondaryOptimizationEvent = "PURCHASE_ROI"
	// SecondaryOptimizationEvent_START_TRIAL 已废弃 开始试用
	SecondaryOptimizationEvent_START_TRIAL SecondaryOptimizationEvent = "START_TRIAL"
	// SecondaryOptimizationEvent_SUBSCRIBE 订阅。 用户订阅了您应用中的频道或服务。 当optimization_goal为INSTALL时可用。
	SecondaryOptimizationEvent_SUBSCRIBE SecondaryOptimizationEvent = "SUBSCRIBE"
	// SecondaryOptimizationEvent_UPDATE_LEVEL 达到等级。用户达到您在游戏中定义的特定级别。当optimization_goal为INSTALL时可用。
	SecondaryOptimizationEvent_UPDATE_LEVEL SecondaryOptimizationEvent = "UPDATE_LEVEL"
	// SecondaryOptimizationEvent_DAY7_RETENTION 7日留存。
	// 仅当以下条件均满足时生效：
	// 推广系列层级：objective_type为 APP_PROMOTION 或APP_INSTALL
	// 广告组层级：
	// placements = PLACEMENT_PANGLE
	// optimization_goal = INSTALL
	// optimization_event = ACTIVE
	// 注意：7日留存目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	SecondaryOptimizationEvent_DAY7_RETENTION SecondaryOptimizationEvent = "DAY7_RETENTION"
	// SecondaryOptimizationEvent_PREFERRED_LEAD 偏好的潜在客户。向最有可能符合你偏好的潜在客户资质的人群展示广告。
	SecondaryOptimizationEvent_PREFERRED_LEAD SecondaryOptimizationEvent = "PREFERRED_LEAD"
)
