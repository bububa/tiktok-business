package enum

// EntertainmentProductCategory 媒体标题项的种类。
type EntertainmentProductCategory string

const (
	// EntertainmentProductCategory_MOVIE：电影。
	EntertainmentProductCategory_MOVIE EntertainmentProductCategory = "MOVIE"
	// EntertainmentProductCategory_MUSIC：音乐。
	EntertainmentProductCategory_MUSIC EntertainmentProductCategory = "MUSIC"
	// EntertainmentProductCategory_TV_SHOW：电视节目。
	EntertainmentProductCategory_TV_SHOW EntertainmentProductCategory = "TV_SHOW"
	// EntertainmentProductCategory_TV_SERIES：电视连续剧。
	EntertainmentProductCategory_TV_SERIES EntertainmentProductCategory = "TV_SERIES"
	// EntertainmentProductCategory_SPORTS_GAME：体育比赛。
	EntertainmentProductCategory_SPORTS_GAME EntertainmentProductCategory = "SPORTS_GAME"
	// EntertainmentProductCategory_LIVE_EVENT：直播活动。
	EntertainmentProductCategory_LIVE_EVENT EntertainmentProductCategory = "LIVE_EVENT"
)
