package enum

// CommentType 评论类型。
type CommentType string

const (
	// CommentType_ALL 全部
	CommentType_ALL CommentType = "ALL"
	// CommentType_COMMENT 原始评论
	CommentType_COMMENT CommentType = "COMMENT"
	// CommentType_REPLY 回复评论
	CommentType_REPLY CommentType = "REPLY"
)
