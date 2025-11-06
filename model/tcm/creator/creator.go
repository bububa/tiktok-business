package creator

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

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
	// IsOnVacation 创作者是否处于休假模式中。 枚举值：
	// True - 创作者处于休假模式，目前暂不接受订单邀请。
	// False - 创作者未处于休假模式，目前可以接受订单邀请。
	IsOnVacation bool `json:"is_on_vacation,omitempty"`
	// VocationStartDate 休假起始日期
	VocationStartDate model.DateTime `json:"vocation_start_date,omitzero"`
	// VocationEndDate 休假结束日期
	VocationEndDate model.DateTime `json:"vocation_end_date,omitzero"`
	// CreatorRate 创作者起始费率
	CreatorRate *CreatorRate `json:"creator_rate,omitempty"`
	// Location 创作者所在地的国家或地区代码，格式为ISO 3166-1 两位字母。
	Location string `json:"location,omitempty"`
	// Topics 创作者最近 30 个视频的主要主题类别。
	// 可传入的值参见下文 “topics可选值 ”小节。
	Topics []string `json:"topics,omitempty"`
	// CreatorBadges 该创作者目前拥有的徽章。
	// 枚举值：
	// RESPONSIVE ：响应迅速。创作者在48小时内回复品牌商。
	// EXPERIENCED ：经验丰富。创作者已完成第一个推广系列。
	// PROFILE_COMPLETE ：已完善个人主页。创作者已完整填写个人资料。
	CreatorBadges []enum.CreatorBadge `json:"creator_badges,omitempty"`
	// AudienceCountries 受众规模最大的三个国家或地区
	AudienceCountries []AudienceCountry `json:"audience_countries,omitempty"`
	// AudienceGenders 受众性别信息
	AudienceGenders []AudienceGender `json:"audience_genders,omitempty"`
	// AudienceAges 受众年龄分布
	AudienceAges []AudienceAge `json:"audience_ages,omitempty"`
	// AudienceLocales 对于美国创作者，指受众最多的 3 个州
	AudienceLocales []AudienceLocale `json:"audience_locales,omitempty"`
	// AudienceDevice 受众设备信息
	AudienceDevice []AudienceDevice `json:"audience_device,omitempty"`
	// AudienceUsage 受众的 TikTok 使用情况
	AudienceUsage []AudienceUsage `json:"audience_usage,omitempty"`
	// AverageViews 平均播放量等于视频总播放次数除以视频数（取最近 30 个视频的平均数）
	AverageViews float64 `json:"average_views,omitempty"`
	// AverageViewsBenchmark 该创作者视频平均播放量的排名（取最近 30 个视频的平均数），排名基于与其他在TTCM上拥有相似粉丝数量且涉及相同主题类别的创作者的对比。如果该创作者的视频涉及多个主题类别，则将显示所有主题类别中排名最高的一个。
	// 例如，1.0000表示该创作者的视频的平均播放量在TTCM上的相同主题类别的视频中排名处于前1%。
	AverageViewsBenchmark float64 `json:"average_views_benchmark,omitempty"`
	// SixSecondsViewsRate 6 秒视频观看率，等于观看了6秒钟的视频观看次数除以总视频观看次数（取最近 30 个视频的平均数）
	SixSecondsViewsRate float64 `json:"six_seconds_views_rate,omitempty"`
	// SixSecondsViewsRateBenchmark 该创作者视频的 6 秒视频观看率的排名（取最近 30 个视频的平均数），排名基于与其他在TTCM上拥有相似粉丝数量且涉及相同主题类别的创作者的对比。如果该创作者的视频涉及多个主题类别，则将显示所有主题类别中排名最高的一个。
	// 例如，1.0000表示该创作者的视频的 6 秒视频观看率在TTCM上的相同主题类别的视频中排名处于前1%
	SixSecondsViewsRateBenchmark float64 `json:"six_seconds_views_rate_benchmark,omitempty"`
	// EngagementRate 互动率等于帖子互动（点赞、分享和评论）次数除以视频总播放次数（取最近 30 个视频的平均数）
	EngagementRate float64 `json:"engagement_rate,omitempty"`
	// EngagementRateBenchmark 该创作者视频的互动率的排名（取最近 30 个视频的平均数），排名基于与其他在TTCM上拥有相似粉丝数量且涉及相同主题类别的创作者的对比。如果该创作者的视频涉及多个主题类别，则将显示所有主题类别中排名最高的一个。
	// 例如，1.0000表示该创作者的视频的互动率在TTCM上的相同主题类别的视频中排名处于前1%
	EngagementRateBenchmark float64 `json:"engagement_rate_benchmark,omitempty"`
	// EngagementsBenchmark 该创作者视频的平均互动次数的排名（取最近 30 个视频的平均数），排名基于与其他在TTCM上拥有相似粉丝数量且涉及相同主题类别的创作者的对比。如果该创作者的视频涉及多个主题类别，则将显示所有主题类别中排名最高的一个。
	// 例如，1.0000表示该创作者的视频的互动数在TTCM上的相同主题类别的视频中排名处于前1%
	EngagementsBenchmark float64 `json:"engagements_benchmark,omitempty"`
	// CompletionRate 完播率等于完整播放视频的次数除以视频总播放次数（取最近 30 个视频的平均数）
	CompletionRate float64 `json:"completion_rate,omitempty"`
	// CompletionRateBenchmark 该创作者视频的完播率的排名（取最近 30 个视频的平均数），排名基于与其他在TTCM上拥有相似粉丝数量且涉及相同主题类别的创作者的对比。如果该创作者的视频涉及多个主题类别，则将显示所有主题类别中排名最高的一个。
	// 例如，1.0000表示该创作者的视频的完播率在TTCM上的相同主题类别的视频中排名处于前1%
	CompletionRateBenchmark float64 `json:"completion_rate_benchmark,omitempty"`
	// FollowersGrowthRate 粉丝增长率指粉丝数与 30 天前相比的增长率。不适用于没有 30 天前数据的新创作者
	FollowersGrowthRate float64 `json:"followers_growth_rate,omitempty"`
	// FollowersGrowthRateBenchmark 该创作者的粉丝增长率的排名，排名基于与其他在TTCM上拥有相似粉丝数量且涉及相同主题类别的创作者的对比。如果该创作者涉及多个主题类别，则将显示所有主题类别中排名最高的一个。
	// 例如，1.0000表示该创作者的粉丝增长率在TTCM上的相同主题类别的创作者中排名处于前1%。
	FollowersGrowthRateBenchmark float64 `json:"followers_growth_rate_benchmark,omitempty"`
	// AverageLikes 最近 30 个视频的平均点赞数
	AverageLikes float64 `json:"average_likes,omitempty"`
	// AverageShares 最近 30 个视频的平均分享次数
	AverageShares float64 `json:"average_shares,omitempty"`
	// AverageComments 最近 30 个视频的平均评论数
	AverageComments float64 `json:"average_comments,omitempty"`
	// SponsoredVideo 该创作者发布的最近 30 个品牌内容视频的相关信息
}

// CreatorRate 创作者起始费率
type CreatorRate struct {
	// Rate 起始费率的金额
	Rate float64 `json:"rate,omitempty"`
	// Currency 起始费率对应的币种
	Currency string `json:"currency,omitempty"`
	// IsNegotiableType 创作者费率是否可协商。枚举值：
	// True: 无论创作者是否有起始费率，均可与创作者协商价格。
	// 注意: 若创作者无起始费率， is_negotiable_type 将会为True，且rate和 currency 的值均为空。
	// False: 创作者有起始费率，且价格不可协商。此时将返回rate和 currency 的具体值 。
	IsNegotiableType bool `json:"is_negotiable_type,omitempty"`
}

// AudienceCountry 受众规模最大的三个国家或地区
type AudienceCountry struct {
	// Country 国家或地区名称
	Country string `json:"country,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// AudienceGender 受众性别信息
type AudienceGender struct {
	// Gender 枚举值：FEMALE，MALE，NONE
	Gender enum.Gender `json:"gender,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// AudienceAge 受众年龄分布
type AudienceAge struct {
	// Age 年龄组。枚举值：18-24，25-34，35+
	Age string `json:"age,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// AudienceLocale 对于美国创作者，指受众最多的 3 个州
type AudienceLocale struct {
	// Locale 美国的州（的2 位字母代码）。相关详情请参见美国各州和领土缩写列表。
	Locale string `json:"locale,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// AudienceDevice 受众设备信息
type AudienceDevice struct {
	// Device 枚举值：ios，android
	Device string `json:"device,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// AudienceUsage 受众的 TikTok 使用情况
type AudienceUsage struct {
	// Usage 受众是否经常使用 TikTok。枚举值：Active（受众近三十天内登陆了TikTok App时显示此状态），Inactive。
	Usage string `json:"usage,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// SponsoredVideo 该创作者发布的最近 30 个品牌内容视频的相关信息
type SponsoredVideo struct {
	// AvgViewsPerVideo 该创作者的品牌内容视频平均播放量（取最近 30 个品牌内容视频的数据）
	AvgViewsPerVideo float64 `json:"avg_views_per_video,omitempty"`
	// AvgViewsPerVideoBenchmarking 该创作者的品牌内容视频平均播放量的排名 （取最近 30 个品牌内容视频的数据），排名基于与其他在TTCM上拥有相似粉丝数量且涉及相同主题类别的创作者的对比。如果该创作者的品牌内容视频涉及多个主题类别，则将显示所有主题类别中排名最高的一个。
	// 例如，1.0000表示该创作者的品牌内容视频平均播放量在TTCM上的相同主题类别的品牌内容视频中排名处于前1%。
	AvgViewsPerVideoBenchmarking float64 `json:"avg_views_per_video_benchmarking,omitempty"`
	// AvgSixSecondsViewsPerVideo 该创作者的品牌内容视频 6 秒观看率（取最近 30 个品牌内容视频的数据）
	AvgSixSecondsViewsPerVideo float64 `json:"avg_six_seconds_views_per_video,omitempty"`
	// AvgSixSecondsViewsPerVideoBenchmarking 该创作者的品牌内容视频 6 秒观看率的排名（取最近 30 个品牌内容视频的数据），排名基于与其他在TTCM上拥有相似粉丝数量且涉及相同主题类别的创作者的对比。如果该创作者的品牌内容视频涉及多个主题类别，则将显示所有主题类别中排名最高的一个。
	// 例如，1.0000表示该创作者的品牌内容视频 6 秒观看率在TTCM上的相同主题类别的品牌内容 视频中排名处于前1%。
	AvgSixSecondsViewsPerVideoBenchmarking float64 `json:"avg_six_seconds_views_per_video_benchmarking,omitempty"`
	// VideoCompletionRate 该创作者的品牌内容视频完播率（取最近 30 个品牌内容视频的数据）
	VideoCompletionRate float64 `json:"video_completion_rate,omitempty"`
	// VideoCompletionRateBenchmarking 该创作者的品牌内容视频完播率的排名（ 取最近 30 个品牌内容视频的数据），排名基于与其他在TTCM上拥有相似粉丝数量且涉及相同主题类别的创作者的对比。如果该创作者的品牌内容视频涉及多个主题类别，则将显示所有主题类别中排名最高的一个。
	// 例如，1.0000表示该创作者的品牌内容视频完播率在TTCM上的相同主题类别的品牌内容视频中排名处于前1%。
	VideoCompletionRateBenchmarking float64 `json:"video_completion_rate_benchmarking,omitempty"`
	// RecentVideosLikes 该创作者的品牌内容视频平均点赞数（取最近 30 个品牌内容视频的数据）
	RecentVideosLikes float64 `json:"recent_videos_likes,omitempty"`
	// RecentVideosShares 该创作者的品牌内容视频平均分享数（取最近 30 个品牌内容视频的数据）
	RecentVideosShares float64 `json:"recent_videos_shares,omitempty"`
	// RecentVideosComments 该创作者的品牌内容视频平均评论数（取最近 30 个品牌内容视频的数据）
	RecentVideosComments float64 `json:"recent_videos_comments,omitempty"`
	// RecentVideosEngagements 该创作者的品牌内容视频平均互动次数（取最近 30 个品牌内容视频的数据）
	RecentVideosEngagements float64 `json:"recent_videos_engagements,omitempty"`
	// RecentVideosEngagementsBenchmarking	该创作者的品牌内容视频平均互动次 数的排名（取最近 30 个品牌内容视频的数据），排名基于与其他在TTCM上拥有相似粉丝数量且涉及相同主题类别的创作者的对比。如果该创作者的品牌内容视频涉及多个主题类别，则将显示所有主题类别中排名最高的一个。
	// 例如，1.0000表示该创作者的品牌内容视频平均互动次数在TTCM上的相同主题类别的品牌内容视频中排名处于前1%
	RecentVideosEngagementsBenchmarking float64 `json:"recent_videos_engagements_benchmarking,omitempty"`
	// EngagementRate 该创作者的品牌内容视频互动率（取最近 30 个品牌内容视频的数据）
	EngagementRate float64 `json:"engagement_rate,omitempty"`
	// EngagementRateBenchmarking 该创作者的品牌内容视频互动率的排名（取最近 30 个品牌内容视频的数据），排名基于与其他在TTCM上拥有相似粉丝数量且涉及相同主题类别的 创作者的对比。如果该创作者的品牌内容视频涉及多个主题类别，则将显示所有主题类别中排名最高的一个。
	// 例如，1.0000表示该创作者的品牌内容视频互动率在TTCM上的相同主题类别的品牌内容视频中排名处于前1%
	EngagementRateBenchmarking float64 `json:"engagement_rate_benchmarking,omitempty"`
}
