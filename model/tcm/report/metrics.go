package report

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Metrics 指标数据
type Metrics struct {
	// VideoID 与订单绑定的视频 ID
	VideoID string `json:"video_id,omitempty"`
	// EngagementCount 点赞数 + 分享次数 + 评论数
	EngagementCount int64 `json:"engagement_count,omitempty"`
	// EngagementRate 互动率。计算公式：engagement_count / video_views
	EngagementRate float64 `json:"engagement_rate,omitempty"`
	// EngagementRateBenchmarkP25 相似创作者的原生视频互动率的第25个百分位数。
	// 若互动率高于本字段值，代表对应视频的互动率超过25%的相似创作者的原生视频
	EngagementRateBenchmarkP25 float64 `json:"engagement_rate_benchmark_p25,omitempty"`
	// EngagementRateBenchmarkP50 相似创作者的原生视频互动率的第50个百分位数。
	// 若互动率高于本字段值，代表对应视频的互动率超过50%的相似创作者的原生视频。
	EngagementRateBenchmarkP50 float64 `json:"engagement_rate_benchmark_p50,omitempty"`
	// VideoViews 视频播放次数
	VideoViews int64 `json:"video_views,omitempty"`
	// Likes 点赞数
	Likes int64 `json:"likes,omitempty"`
	// Shares 分享次数
	Shares int64 `json:"shares,omitempty"`
	// Comments 评论数
	Comments int64 `json:"comments,omitempty"`
	// LifetimeGenderDistribution 视频播放量按性别的分布情况
	LifetimeGenderDistribution *GenderDistribution `json:"lifetime_gender_distribution,omitempty"`
	// LifetimeTopAgeDistribution 视频播放量按年龄组的分布情况
	LifetimeTopAgeDistribution *AgeDistribution `json:"lifetime_top_age_distribution,omitempty"`
	// LifetimeTopCountryDistribution 视频播放量按受众地域的分布情况
	LifetimeTopCountryDistribution *CountryDistribution `json:"lifetime_top_country_distribution,omitempty"`
	// LifetimeDeviceDistribution 视频播放量按受众设备的分布情况
	LifetimeDeviceDistribution *DeviceDistribution `json:"lifetime_device_distribution,omitempty"`
	// MostPopularCountryStateDistribution 该视频最受欢迎的国家或地区中的州/省级别地域的分布信息
	MostPopularCountryStateDistribution *CountryStateDistribution `json:"most_popular_country_state_distribution,omitempty"`
	// AudienceInterestDistribution 占比前10的受众兴趣分布
	AudienceInterestDistribution *AudienceInterestDistribution `json:"audience_interest_distribution,omitempty"`
	// Reach 覆盖人数。观看过该视频至少一次的独立用户数，此数值为估算值
	Reach int64 `json:"reach,omitempty"`
	// VideoCompletionRate 视频完播率 = 视频完播次数/视频总播放次数
	VideoCompletionRate float64 `json:"video_completion_rate,omitempty"`
	// VideoCompletionRateBenchmarkP25 相似创作者的原生视频完播率的第25个百分位数。
	// 若完播率高于本字段值，代表对应视频的完播率超过25%的相似创作者的原生视频。
	VideoCompletionRateBenchmarkP25 float64 `json:"video_completion_rate_benchmark_p25,omitempty"`
	// VideoCompletionRateBenchmarkP50 相似创作者的原生视频完播率的第50个百分位数。
	// 若完播率高于本字段值，代表对应视频的完播率超过50%的相似创作者的原生视频。
	VideoCompletionRateBenchmarkP50 float64 `json:"video_completion_rate_benchmark_p50,omitempty"`
	// VideoViewRetention 多组数字构成的列表。每组数字的格式为[x,y]，其中x为当前为第几秒，y代表在x秒时的视频留存率。
	// 例如，[[22.0,0.1], [8.0,0.2], [5.0,0.3],[3.0,0.5], [2.0,0.6],[1.0,0.9]]代表：
	// 第8-22秒内，视频留存率从20%下降为 10%。
	// 第5-8秒内，视频留存率从30%下降为 20%。
	// 第3-5秒内，视频留存率从50%下降为 30%。
	// 第3秒内，视频留存率从60%下降为 50%。
	// 第2秒内，视频留存率从90%下降为 60%。
	// 第1秒内，视频留存率从100%下降为 90%。
	VideoViewRetention [][2]float64 `json:"video_view_retention,omitempty"`
	// TotalPlayTime 视频总播放时间，单位为秒
	TotalPlayTime int64 `json:"total_play_time,omitempty"`
	// AverageViewTime 视频平均播放时间，单位为秒
	AverageViewTime float64 `json:"average_view_time,omitempty"`
	// TwoSecondsViews 视频播放达2秒次数 / 视频总播放次数
	TwoSecondsViews float64 `json:"two_seconds_views,omitempty"`
	// SixSecondsViews 视频播放达6秒次数 / 视频总播放次数
	SixSecondsViews float64 `json:"six_seconds_views,omitempty"`
	// VideoViewsBySource 视频播放信息
	VideoViewsBySource *VideoViewsBySource `json:"video_views_by_source,omitempty"`
	// VideoViewsOrganic 自然流量产生的视频播放总量。
	// 即通过自然流量，而非Spark Ads产生的视频播放总量。
	VideoViewsOrganic int64 `json:"video_views_organic,omitempty"`
	// VideoViewsPaid 付费流量产生的视频播放总量。
	// 即通过Spark Ads产生的视频播放总量
	VideoViewsPaid int64 `json:"video_views_paid,omitempty"`
	// VideoBoostedDate 视频推广日期。即该视频首次推广的日期。格式为："YYYY-MM-DD HH:MM:SS"（UTC+0时间）。例如，"2023-01-01 00:00:01"。
	// 注意：若视频未经过推广，则本字段会显示默认值 "1970-01-01"
	VideoBoostedDate model.DateTime `json:"video_boosted_date,omitzero"`
	// DailyStats 每日细分数据。仅当请求中传入start_date和 end_date时返回。
	// 注意: 缺失的指标数据不会进行默认值（如0）填充，因此若当日数据中无某一指标，代表该指标无有效数据。
	DailyStats []DailyStat `json:"daily_stats,omitempty"`
}

// GenderDistribution 视频播放量按性别的分布情况
type GenderDistribution struct {
	// Gender 性别。枚举值：FEMALE，MALE
	Gender enum.Gender `json:"gender,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// AgeDistribution 视频播放量按年龄组的分布情况
type AgeDistribution struct {
	// Age 年龄组。枚举值：18-24，25-34，35+
	Age string `json:"age,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// CountryDistribution 视频播放量按受众地域的分布情况
type CountryDistribution struct {
	// Country 两个字母组成的国家或地区缩写
	Country string `json:"country,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// DeviceDistribution 视频播放量按受众设备的分布情况
type DeviceDistribution struct {
	// Device 设备平台。示例：ios 或 android
	Device string `json:"device,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// CountryStateDistribution 该视频最受欢迎的国家或地区中的州/省级别地域的分布信息
type CountryStateDistribution struct {
	// Country 国家或地区代码。枚举值：
	// US：美国
	// GB：英国
	// AU：澳大利亚
	// BR： 巴西
	// ID： 印度尼西亚
	// MY： 马来西亚
	// PH： 菲律宾
	// SA：沙地阿拉伯
	// SG： 新加坡
	Country string `json:"country,omitempty"`
	// LocaleDistribution 该国家或地区中州/省级别地域的分布信息
	LocaleDistribution []LocaleDistribution `json:"locale_distribution,omitempty"`
}

// LocaleDistribution 该国家或地区中州/省级别地域的分布信息
type LocaleDistribution struct {
	// Locale 地域名称
	Locale string `json:"locale,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// AudienceInterestDistribution 占比前10的受众兴趣分布
type AudienceInterestDistribution struct {
	// InterestTag 受众兴趣类别名称
	InterestTag string `json:"interest_tag,omitempty"`
	// Percentage 百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// VideoViewsBySource 视频播放信息
type VideoViewsBySource struct {
	// ForYou 从推荐页产生的视频播放次数
	ForYou float64 `json:"for_you,omitempty"`
	// Hashtag 从话题标签产生的视频播放次数
	Hashtag float64 `json:"hashtag,omitempty"`
	// Sound 从视频声音产生的视频播放次数
	Sound float64 `json:"sound,omitempty"`
	// PersonalProfile 从个人主页产生的视频播放次数
	PersonalProfile float64 `json:"personal_profile,omitempty"`
	// Search 从搜索产生的视频播放次数
	Search float64 `json:"search,omitempty"`
	// Following 从关注页面产生的视频播放次数
	Following float64 `json:"following,omitempty"`
}

// DailyStat 每日细分数据
type DailyStat struct {
	// Date 日期
	Date string `json:"date,omitempty"`
	// Reach 覆盖人数。当日观看了该推广系列下至少一个视频的人数
	Reach float64 `json:"reach,omitempty"`
	// EngagementRate 互动次数/ 视频总播放次数
	EngagementRate float64 `json:"engagement_rate,omitempty"`
	// VideoCompletionRate 视频完播率 = 视频完播次数/视频总播放次数
	VideoCompletionRate float64 `json:"video_completion_rate,omitempty"`
	// Likes 点赞数
	Likes int64 `json:"likes,omitempty"`
	// Shares 分享次数
	Shares int64 `json:"shares,omitempty"`
	// Comments 评论数
	Comments int64 `json:"comments,omitempty"`
	// Views 播放次数
	Views int64 `json:"views,omitempty"`
	// VideoViewsOrganic 当日自然流量产生的视频播放量。
	// 即通过自然流量，而非Spark Ads产生的当日视频播放量。
	VideoViewsOrganic int64 `json:"video_views_organic,omitempty"`
	// VideoViesPaid 当日付费流量产生的视频播放量。
	// 即通过Spark Ads产生的当日视频播放量
	VideoViewsPaid int64 `json:"video_views_paid,omitempty"`
}
