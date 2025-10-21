package enum

	// EntertainmentProductTimeline 媒体标题项的上线状态。
type EntertainmentProductTimeline string

const (
// EntertainmentProductTimeline_COMING_SOON ：该媒体标题项即将上线，暂不支持购买或浏览。
	EntertainmentProductTimeline_COMING_SOON EntertainmentProductTimeline = "COMING_SOON"
// EntertainmentProductTimeline_ONLINE：该媒体标题项已上线，支持购买或浏览。
	EntertainmentProductTimeline_ONLINE EntertainmentProductTimeline = "ONLINE"
// EntertainmentProductTimeline_EXPIRE_SOON：该媒体标题项即将下线，将不支持支持购买或浏览。
EntertainmentProductTimeline_EXPIRE_SOON EntertainmentProductTimeline = "EXPIRE_SOON"
)
