package business

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Insight TikTok 账号主页数据
type Insight struct {
	// IsBusinessAccount 该 TikTok 账户 (business_id) 是否为 TikTok 企业号。
	// 枚举值：
	// true：该 TikTok 账户是 TikTok 企业号。
	// false：该 TikTok 账户是 TikTok 个人账号。
	IsBusinessAccount bool `json:"is_business_account,omitempty"`
	// ProfileImage TikTok 账号主页头像的临时URL。过期日期和时间以 Epoch/Unix 时间格式包含在x-expires查询参数中，以秒为单位。
	// 示例：https://p16-sign-va.tiktokcdn.com/tos-maliva-avt-0068/75dec21d63500917fb6ec8bc59415156~c5_300x300.jpeg?x-expires=1614099600&x-signature=PmK%2BWs3LzSzRL2tYs%2FZx7EjG3Gk%3D
	// 更新时间：实时
	ProfileImage string `json:"profile_image,omitempty"`
	// Username TikTok 账号的用户名。
	// 示例：la_flama_blanca
	// 更新时间：实时
	Username string `json:"username,omitempty"`
	// ProfileDeepLink TikTok 账号主页链接。
	// 示例：https://vm.tiktok.com/ABcdEfghi/
	// 更新时间：实时
	ProfileDeepLink string `json:"profile_deep_link,omitempty"`
	// DisplayName TikTok 账号的显示名称（昵称）。
	// 示例：La Flama Blanca
	// 更新时间：实时
	DisplayName string `json:"display_name,omitempty"`
	// BioDescription TikTok 账号的个人简介。
	// 更新时间：实时
	// 注意：若您从未设置过 TikTok 账号的个人简介，本字段将返回空字符串（""）。
	BioDescription string `json:"bio_description,omitempty"`
	// IsVerified TikTok 账号是否带有认证徽章（蓝 V 认证），从而表明账户所有者的真实性已经过核验。
	// 支持的值：true，false。
	// 示例：true
	// 更新时间：实时
	IsVerified bool `json:"is_verified,omitempty"`
	// FollowingCount TikTok 账号关注的账号总数。
	// 示例：2
	// 更新时间：实时
	FollowingCount int64 `json:"following_count,omitempty"`
	// FollowersCount TikTok 账号的总粉丝数。
	// 示例：123
	// 更新时间：实时
	FollowersCount int64 `json:"followers_count,omitempty"`
	// TotalLikes 总点赞数。
	// 用户对已发布视频的总点赞次数。
	// 示例：6
	// 更新时间：实时
	TotalLikes int64 `json:"total_likes,omitempty"`
	// VideosCount  TikTok 账号发布的公开视频总数。
	// 示例：6
	// 更新时间：实时
	VideosCount int64 `json:"videos_count,omitempty"`
	// Metrics 所有请求的每日数据指标。
	// 注意：仅返回 TikTok 账号创建后的数据指标。数据指标可能延迟最长 24-48 小时
	Metrics []Metrics `json:"metrics,omitempty"`
	// AudienceAges 粉丝年龄。
	// 按年龄列出粉丝的分布情况。该人口统计数据基于许多因素，包括用户在其个人主页中提供的信息。
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为粉丝数至少为 100 的 TikTok 账号返回本指标数据。
	AudienceAges []AudienceAge `json:"audience_ages,omitempty"`
	// AudienceGenders 粉丝性别。
	// 按性别列出粉丝的分布情况。该人口统计数据基于许多因素，包括用户在其个人主页中提供的信息。
	// 示例： [{gender: "Female", percentage: 0.7554470336505679}, {gender: "Male", percentage: 0.24455296634943213}, ...]
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为粉丝数至少为 100 的 TikTok 账号返回本指标数据。
	AudienceGenders []AudienceGender `json:"audience_genders,omitempty"`
	// AudienceCountries 粉丝热门国家/地区。
	// 按粉丝的位置（国家/地区）列出粉丝的分布情况。该人口统计数据基于许多因素，包括用户在其个人主页中提供的信息。
	// 示例： [{country: "US", percentage: 0.7554470336505679}, {country: "UK", percentage: 0.24455296634943213}, ...]
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为粉丝数至少为 100 的 TikTok 账号返回本指标数据。
	AudienceCountries []AudienceCountry `json:"audience_countries,omitempty"`
	// AudienceCities 粉丝热门城市。
	// 粉丝所在城市分布。该人口统计数据基于许多因素，包括用户在其个人主页中提供的信息。
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为粉丝数至少为 100 的 TikTok 账号返回本指标数据。
	AudienceCities []AudienceCity `json:"audience_cities,omitempty"`
}

// Metrics 所有请求的每日数据指标。
type Metrics struct {
	// Date 这些每日数据指标的对应日期，格式为YYYY-MM-DD，时区为UTC。
	// 示例：2021-08-12
	// 更新时间：T+ 24-48 小时（UTC 时间）
	Date string `json:"date,omitempty"`
	// VideoViews 每日观看次数。
	// 您的视频的每日观看次数。
	// 注意：
	// 此指标当视频回放时长大于0，且回放为一次曝光中的首次回放时也计数。
	// 若用户刷过广告后又回刷广告，将视作两次曝光。第二次曝光也会计入播放次数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	VideoViews int64 `json:"video_views,omitempty"`
	// UniqueVideoViews 每日触达的观众。
	// 每日至少观看过一次已发布内容的用户。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	UniqueVideoViews int64 `json:"unique_video_views,omitempty"`
	// ProfileViews 每日个人主页访问量。
	// 每日用户查看您的个人主页的次数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	ProfileViews int64 `json:"profile_views,omitempty"`
	// Likes 每日点赞数。
	// 每日用户对已发布视频的点赞次数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：此字段显示的点赞数不会减去取消点赞数（即后来被移除或撤回的点赞）。这个数值可能与企业套件 > 数据分析页面上企业号显示的点赞数有所不同。
	// 例如，若一条帖子收到 30 个点赞和 10 个取消点赞：
	// 此字段将显示点赞数为 30。
	// 企业号的数据分析页面将显示点赞数为 20（30个点赞 - 10个取消点赞）。
	Likes int64 `json:"likes,omitempty"`
	// Comments 每日评论数。
	// 每日用户对已发布视频的评论次数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	Comments int64 `json:"comments,omitempty"`
	// Shares 每日分享次数。
	// 每日用户对已发布视频的分享次数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	Shares int64 `json:"shares,omitempty"`
	// PhoneNumberClicks 每日手机号码点击数。
	// 在选定日期范围内收集的每日手机号码链接的总点击数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	PhoneNumberClicks int64 `json:"phone_number_clicks,omitempty"`
	// LeadSubmissions 每日潜在客户提交次数。
	// 在选定日期范围内从您的消费者收集的每日潜在客户总数（如报价、新闻资讯订阅等）。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	LeadSubmissions int64 `json:"lead_submissions,omitempty"`
	// AppDownloadClicks 每日应用下载链接点击数。
	// 在选定日期范围内收集的每日应用下载链接的总点击数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	AppDownloadClicks int64 `json:"app_download_clicks,omitempty"`
	// BioLinkClicks 每日个人简介链接点击数。
	// 在选定日期范围内收集的每日个人简介链接的总点击数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	BioLinkClicks int64 `json:"bio_link_clicks,omitempty"`
	// EmailClicks 每日邮箱点击数。
	// 在选定日期范围内收集的每日主页邮箱按钮的总点击数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	EmailClicks int64 `json:"email_clicks,omitempty"`
	// AddressClicks 每日地址点击数。
	// 在选定日期范围内收集的每日主页地址按钮的总点击数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	AddressClicks int64 `json:"address_clicks,omitempty"`
	// DailyTotalViewers 每日净增长。
	// 表示在一定时段内每日关注您账号的用户数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	DailyTotalViewers int64 `json:"daily_total_viewers,omitempty"`
	// DailyNewFollowers 每日新粉丝数。
	// 表示在一定时段内每日获得的新粉丝总数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	DailyNewFollowers int64 `json:"daily_new_followers,omitempty"`
	// DailyLostFollowers 每日掉粉数。
	// 表示在一定时段内每日失去的粉丝总数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	DailyLostFollowers int64 `json:"daily_lost_followers,omitempty"`
	// FollowersCount TikTok 账号在某一日期的粉丝总数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	FollowersCount int64 `json:"followers_count,omitempty"`
	// AudienceActivity 每小时粉丝活动。
	// 表示您的粉丝一天中在 TikTok 上活跃的时段。
	// 示例：[{"hour": "0", "count": 12}, {"hour": "1", "count": 18}, {"hour": "2", "count": 23}, {"hour": "3", "count": 28}, ...]
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为粉丝数至少为 100 的 TikTok 账号返回本指标数据。
	AudienceActivity []AudienceActivity `json:"audience_activity,omitempty"`
	// EngagedAudience 每日互动的观众 。
	// 表示每日与您发布的至少一项内容有过互动（点赞、评论、分享）的用户。
	EngagedAudience int64 `json:"engaged_audience,omitempty"`
}

// AudienceActivity 每小时粉丝活动。
type AudienceActivity struct {
	// Hour 一天中的某一小时
	Hour model.Int `json:"hour,omitempty"`
	// Count 某一小时的活跃粉丝数
	Count int64 `json:"count,omitempty"`
}

// AudienceAge 受众年龄分布
type AudienceAge struct {
	// Age 年龄组。枚举值：18-24，25-34，35-44，45-54，55+
	Age string `json:"age,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// AudienceGender 受众性别信息
type AudienceGender struct {
	// Gender 枚举值：Female，Male，Other
	Gender enum.Gender `json:"gender,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// AudienceCountry 受众规模最大的三个国家或地区
type AudienceCountry struct {
	// Country 国家或地区名称
	Country string `json:"country,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// AudienceCity 粉丝热门城市。
type AudienceCity struct {
	// CityName 城市名称
	CityName string `json:"city_name,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}
