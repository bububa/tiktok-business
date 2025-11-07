package enum

// CreativeDeliveryStatus  创意状态
type CreativeDeliveryStatus string

const (
	// CreativeDeliveryStatus_IN_QUEUE：排队中。作品已准备好进行转化驱动潜力测试。
	CreativeDeliveryStatus_IN_QUEUE CreativeDeliveryStatus = "IN_QUEUE"
	// CreativeDeliveryStatus_LEARNING：学习中。作品正在接受转化驱动潜力测试。
	CreativeDeliveryStatus_LEARNING CreativeDeliveryStatus = "LEARNING"
	// CreativeDeliveryStatus_DELIVERING：投放中。作品已作为推广系列的一部分进行定期推广。
	CreativeDeliveryStatus_DELIVERING CreativeDeliveryStatus = "DELIVERING"
	// CreativeDeliveryStatus_NOT_DELIVERING：未投放。作品在测试过程中未表现出强大的转化潜力，不会作为广告计划的一部分进行定期推广。
	CreativeDeliveryStatus_NOT_DELIVERING CreativeDeliveryStatus = "NOT_DELIVERING"
	// CreativeDeliveryStatus_AUTHORIZATION_NEEDED：需要授权。此视频尚未被授权用于广告。您可以从 TikTok 账号授权自己的视频，或向其他达人申请授权。在授权后 20 分钟内，此处的状态将会更新。
	CreativeDeliveryStatus_AUTHORIZATION_NEEDED CreativeDeliveryStatus = "AUTHORIZATION_NEEDED"
	// CreativeDeliveryStatus_EXCLUDED：已排除。作品已被手动删除，无法作为广告计划的一部分进行定期推广。您可以随时将此作品重新添加到你的广告计划中。
	CreativeDeliveryStatus_EXCLUDED CreativeDeliveryStatus = "EXCLUDED"
	// CreativeDeliveryStatus_UNAVAILABLE：不可用。私密视频、被拒视频、已删除视频或来自已禁用 TikTok 账号的视频不能用于广告。在某些情况下，您可能会看到来自不可用视频的消耗，这是因为在推广系列投放期间被 TikTok 内容审核流程拒绝的视频仍可能会投放一段时间。
	CreativeDeliveryStatus_UNAVAILABLE CreativeDeliveryStatus = "UNAVAILABLE"
	// CreativeDeliveryStatus_REJECTED：拒绝。视频已被内容审查团队拒绝。如果您已提交申诉，那么在申诉审核完成前，视频状态将保持为“已拒绝”。
	CreativeDeliveryStatus_REJECTED CreativeDeliveryStatus = "REJECTED"
	// CreativeDeliveryStatus_NOT_ACTIVE：不活跃。此视频已被降级处理，因为它上传超过 30 天且近 30 天内未产生任何收入。
	CreativeDeliveryStatus_NOT_ACTIVE CreativeDeliveryStatus = "NOT_ACTIVE"
)
