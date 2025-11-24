package enum

// CommentAction 要对评论进行的特定操作 - 点赞或取消点赞。
type CommentAction string

const (
	// CommentAction_LIKE
	CommentAction_LIKE CommentAction = "LIKE"
	// CommentAction_UNLIKE
	CommentAction_UNLIKE CommentAction = "UNLIKE"
	// CommentAction_HIDE
	CommentAction_HIDE CommentAction = "HIDE"
	// CommentAction_UNHIDE
	CommentAction_UNHIDE CommentAction = "UNHIDE"
)
