package enum

// VideoUserAction 所选用户行为种类下想要定向的的具体用户行为。
// 若action_scene为VIDEO_RELATED，枚举值为：WATCHED_TO_END（完播），LIKED（点赞），COMMENTED（评论），SHARED（分享）。
// 若action_scene为CREATOR_RELATED，枚举值为：FOLLOWING（关注），VIEW_HOMEPAGE（浏览主页）。
// 若action_scene为HASHTAG_RELATED，枚举值为： VIEW_HASHTAG（浏览话题标签）。
type VideoUserAction string

const (
	// VideoUserAction_WATCHED_TO_END（完播）
	VideoUserAction_WATCHED_TO_END VideoUserAction = "WATCHED_TO_END"
	// VideoUserAction_LIKED（点赞）
	VideoUserAction_LIKED VideoUserAction = "LIKED"
	// VideoUserAction_COMMENTED（评论）
	VideoUserAction_COMMENTED VideoUserAction = "COMMENTED"
	// VideoUserAction_SHARED（分享)
	VideoUserAction_SHARED VideoUserAction = "SHARED"
	// VideoUserAction_FOLLOWING（关注）
	VideoUserAction_FOLLOWING VideoUserAction = "FOLLOWING"
	// VideoUserAction_VIEW_HOMEPAGE（浏览主页）
	VideoUserAction_VIEW_HOMEPAGE VideoUserAction = "VIEW_HOMEPAGE"
	// VideoUserAction_VIEW_HASHTAG（浏览话题标签）
	VideoUserAction_VIEW_HASHTAG VideoUserAction = "VIEW_HASHTAG"
)
