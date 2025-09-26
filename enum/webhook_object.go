package enum

// WebhookObject 订阅对象类型。
type WebhookObject int

const (
	// LeadWebhook 线索
	LeadWebhook WebhookObject = 1
	// AdgroupWebhook 广告组的审核状态。
	AdgroupWebhook WebhookObject = 2
	// AdWebhook 广告的审核状态。
	AdWebhook WebhookObject = 3
	// 8：单个广告、广告组中所有广告或广告账户下所有广告的疲劳状态。
	CreativeFatigueWebhook WebhookObject = 8
	// 9 ：TCM 订单中视频的 Spark Ads 授权状态。
	TCMSparkAdsWebhook WebhookObject = 9
)
