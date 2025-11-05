package creator

// Creator TikTok 创作者
type Creator struct {
	// CreatorID TikTok 创作者 ID
	CreatorID string `json:"creator_id,omitempty"`
	// DisplayName 创作者主页中显示的名称
	DisplayName string `json:"display_name,omitempty"`
	// HandleName 相似创作者的唯一用户名
	HandleName string `json:"handle_name,omitempty"`
	// ProfileImage 相似创作者的头像 URI
	ProfileImage string `json:"profile_image,omitempty"`
	// FollowersCount 相似创作者的粉丝总数
	FollowersCount int64 `json:"followers_count,omitempty"`
	// FollowingsCount 相似创作者在 TikTok 上关注的账户总数
	FollowingsCount int64 `json:"followings_count,omitempty"`
	// VideosCount 相似创作者的视频总数
	VideosCount int `json:"videos_count,omitempty"`
	// LikesCount 相似创作者所有视频的点赞总数
	LikesCount int64 `json:"likes_count,omitempty"`
	// Bio 在 TikTok 创作者个人主页的个人简介部分显示的文本
	Bio string `json:"bio,omitempty"`
}
