package enum

// PrivacyLevelOption TikTok 账号可选的隐私级别设置
type PrivacyLevelOption string

const (
	// PrivacyLevelOption_PUBLIC_TO_EVERYONE：所有人。
	PrivacyLevelOption_PUBLIC_TO_EVERYONE PrivacyLevelOption = "PUBLIC_TO_EVERYONE"
	// PrivacyLevelOption_FOLLOWER_OF_CREATOR：粉丝。
	PrivacyLevelOption_FOLLOWER_OF_CREATOR PrivacyLevelOption = "FOLLOWER_OF_CREATOR"
	// PrivacyLevelOption_MUTUAL_FOLLOW_FRIENDS：好友（您回关的粉丝）。
	PrivacyLevelOption_MUTUAL_FOLLOW_FRIENDS PrivacyLevelOption = "MUTUAL_FOLLOW_FRIENDS"
	// PrivacyLevelOption_SELF_ONLY：仅自己。
	PrivacyLevelOption_SELF_ONLY PrivacyLevelOption = "SELF_ONLY"
)
