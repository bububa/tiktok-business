package enum

// CommentStatus 评论（及评论回复）可见状态。
type CommentStatus string

const (
	// CommentStatus_PUBLIC：所有 TikTok 用户均可公开查看的评论（及评论回复）。
	CommentStatus_PUBLIC CommentStatus = "PUBLIC"
	// CommentStatus_HIDDEN：评论或评论回复已隐藏，无法公开浏览
	CommentStatus_HIDDEN CommentStatus = "HIDDEN"
	// CommentStatus_ALL：视频发布者的隐藏评论（及评论回复）以及所有 TikTok 用户均可公开查看的评论（及评论回复）。
	CommentStatus_ALL CommentStatus = "ALL"
)
