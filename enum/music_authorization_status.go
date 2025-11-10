package enum

// MusicAuthorizationStatus TikTok帖子中所有音乐的聚合授权状态
type MusicAuthorizationStatus string

const (
	// MusicAuthorizationStatus_WITHOUT_SONG_ID
	MusicAuthorizationStatus_WITHOUT_SONG_ID MusicAuthorizationStatus = "WITHOUT_SONG_ID"
	// MusicAuthorizationStatus_AUTHORIZATION_MISSING
	MusicAuthorizationStatus_AUTHORIZATION_MISSING MusicAuthorizationStatus = "AUTHORIZATION_MISSING"
	// MusicAuthorizationStatus_WITH_FULL_AUTHORIZATION
	MusicAuthorizationStatus_WITH_FULL_AUTHORIZATION MusicAuthorizationStatus = "WITH_FULL_AUTHORIZATION"
)
