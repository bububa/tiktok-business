package enum

// BusinessMessageDirectReplyType The type of direct reply.
type BusinessMessageDirectReplyType string

const (
	// BusinessMessageDirectReplyType_COMMENT_TO_MESSAGE: Comment-to-Message. Once you enable this feature, the TikTok system will identify high-intent comments, enabling you to send customized replies to assist these users more effectively.
	BusinessMessageDirectReplyType_COMMENT_TO_MESSAGE BusinessMessageDirectReplyType = "COMMENT_TO_MESSAGE"
)
