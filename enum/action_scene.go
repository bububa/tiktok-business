package enum

// ActionScene 用于定向的用户行为种类。
type ActionScene string

const (
	// ActionScene_VIDEO_RELATED：视频互动。
	ActionScene_VIDEO_RELATED ActionScene = "VIDEO_RELATED"
	// ActionScene_CREATOR_RELATED：创作者互动。
	ActionScene_CREATOR_RELATED ActionScene = "CREATOR_RELATED"
	// ActionScene_HASHTAG_RELATED：话题互动。
	ActionScene_HASHTAG_RELATED ActionScene = "HASHTAG_RELATED"
)
