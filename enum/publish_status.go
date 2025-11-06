package enum

// PublishStatus TikTok 帖子的发布状态。
type PublishStatus string

const (
	// PublishStatus_PROCESSING_DOWNLOAD：正在从 URL 中提取内容。
	PublishStatus_PROCESSING_DOWNLOAD PublishStatus = "PROCESSING_DOWNLOAD"
	// PublishStatus_PUBLISH_COMPLETE：帖子已通过审核且发布成功。
	PublishStatus_PUBLISH_COMPLETE PublishStatus = "PUBLISH_COMPLETE"
	// PublishStatus_FAILED：帖子发布出现错误，帖子发布失败。
	PublishStatus_FAILED PublishStatus = "FAILED"
	// PublishStatus_SEND_TO_USER_INBOX：帖子草稿上传成功，系统已将相应通知发至创作者的收件箱中。
	PublishStatus_SEND_TO_USER_INBOX PublishStatus = "SEND_TO_USER_INBOX"
)
