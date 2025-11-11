package enum

// TTAccountRole 商务中心用户对于 TikTok 账户的权限。
type TTAccountRole string

const (
	// TTAccountRole_POST：现有帖子权限，包括访问你自己的个人主页信息、访问你自己的公开视频和访问你自己的广告授权视频。
	TTAccountRole_POST TTAccountRole = "POST"
	// TTAccountRole_LIVE：直播视频权限（仅部分地区可以使用），包括访问你自己的个人主页信息和访问直播视频。
	TTAccountRole_LIVE TTAccountRole = "LIVE"
	// TTAccountRole_DIRECT_MESSAGE：私信权限，包括允许所有人向您发送私信，阅读和管理您收件箱中的私信，以及代表您向其他账号发送消息或进行互动。
	TTAccountRole_DIRECT_MESSAGE TTAccountRole = "DIRECT_MESSAGE"
)
