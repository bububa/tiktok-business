package enum

// AppPromotionType 应用推广类型
type AppPromotionType string

const (
	// AppPromotionType_APP_INSTALL 应用安装。让新用户安装您的应用
	AppPromotionType_APP_INSTALL AppPromotionType = "APP_INSTALL"
	// AppPromotionType_APP_RETARGETING 应用再营销。重新吸引现有用户在您的应用中执行操作。
	AppPromotionType_APP_RETARGETING AppPromotionType = "APP_RETARGETING"
	// AppPromotionType_APP_PREREGISTRATION 应用预注册。让新用户在您的应用发布前完成预注册。
	AppPromotionType_APP_PREREGISTRATION AppPromotionType = "APP_PREREGISTRATION"
	// AppPromotionType_APP_POSTS_PROMOTION：应用帖子推广。推广 TikTok 帖子，提升应用认知度并度量转化量。
	AppPromotionType_APP_POSTS_PROMOTION AppPromotionType = "APP_POSTS_PROMOTION"
	// AppPromotionType_UNSET
	AppPromotionType_UNSET AppPromotionType = "UNSET"
)
