package enum

// EventType 标准事件
type EventType string

const (
	//   网站标准事件 && 线下标准事件 && CRM 标准事件
	// 下表列出了所有的网站标准事件 （包括 event_source 设置为web时的线下事件），以及每种网站事件在properties对象中的推荐参数。需注意事件名区分大小写，因此请务必在/event/track/中的event参数使用事件列中显示的确切名称。
	// 线下标准事件包含的事件与网站标准事件完全相同。
	// 支持的 CRM 标准事件与支持的网站标准事件相同。

	// EventType_AddPaymentInfo 添加支付信息
	EventType_AddPaymentInfo EventType = "AddPaymentInfo"
	// EventType_AddToCart 加入购物车
	EventType_AddToCart EventType = "AddToCart"
	// EventType_AddToWishlist 加入心愿单
	EventType_AddToWishlist EventType = "AddToWishlist"
	// EventType_ApplicationApproval 批准申请
	EventType_ApplicationApproval EventType = "ApplicationApproval"
	// EventType_CompleteRegistration 完成注册
	EventType_CompleteRegistration EventType = "CompleteRegistration"
	// EventType_Contact 联系
	EventType_Contact EventType = "Contact"
	// EventType_CustomizeProduct 定制商品
	EventType_CustomizeProduct EventType = "CustomizeProduct"
	// EventType_Download 下载
	EventType_Download EventType = "Download"
	// EventType_FindLocation 寻找位置
	EventType_FindLocation EventType = "FindLocation"
	// EventType_InitiateCheckout 开始结账
	EventType_InitiateCheckout EventType = "InitiateCheckout"
	// EventType_Lead 提交表单
	EventType_Lead EventType = "Lead"
	// EventType_Purchase 支付完成
	EventType_Purchase EventType = "Purchase"
	// EventType_Schedule 预约
	EventType_Schedule EventType = "Schedule"
	// EventType_Search 搜索
	EventType_Search EventType = "Search"
	// EventType_StartTrial 开始试用
	EventType_StartTrial EventType = "StartTrial"
	// EventType_SubmitApplication 提交申请
	EventType_SubmitApplication EventType = "SubmitApplication"
	// EventType_Subscribe 订阅
	EventType_Subscribe EventType = "Subscribe"
	// EventType_ViewContent 查看内容
	EventType_ViewContent EventType = "ViewContent"

	//	应用标准事件
	//
	// EventType_AchieveLevel 达到级别
	EventType_AchieveLevel EventType = "AchieveLevel"
	// EventType_Checkout 结账
	EventType_Checkout EventType = "Checkout"
	// EventType_CompleteTutorial 完成教程
	EventType_CompleteTutorial EventType = "CompleteTutorial"
	// EventType_CreateGroup 创建组
	EventType_CreateGroup EventType = "CreateGroup"
	// EventType_CreateRole 创建角色
	EventType_CreateRole EventType = "CreateRole"
	// EventType_GenerateLead 收集线索
	EventType_GenerateLead EventType = "GenerateLead"
	// EventType_InAppADClick 应用内广告点击
	EventType_InAppADClick EventType = "InAppADClick"
	// EventType_InAppADImpr 应用内广告展示
	EventType_InAppADImpr EventType = "InAppADImpr"
	// EventType_InstallApp 应用安装
	EventType_InstallApp EventType = "InstallApp"
	// EventType_JoinGroup 加入组
	EventType_JoinGroup EventType = "JoinGroup"
	// EventType_LaunchAPP 启动应用
	EventType_LaunchAPP EventType = "LaunchAPP"
	// EventType_LoanApplication 贷款申请
	EventType_LoanApplication EventType = "LoanApplication"
	// EventType_LoanApproval 贷款批准
	EventType_LoanApproval EventType = "LoanApproval"
	// EventType_LoanDisbursal 贷款发放
	EventType_LoanDisbursal EventType = "LoanDisbursal"
	// EventType_Login 登陆
	EventType_Login EventType = "Login"
	// EventType_Rate 评分
	EventType_Rate EventType = "Rate"
	// EventType_Registration 注册
	EventType_Registration EventType = "Registration"
	// EventType_SpendCredits 花费点数
	EventType_SpendCredits EventType = "SpendCredits"
	// EventType_UnlockAchievement 解锁关卡
	EventType_UnlockAchievement EventType = "UnlockAchievement"
)
