package report

import (
	"encoding/json"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Metrics 指标数据
type Metrics struct {
	// VideoInfo 视频信息
	VideoInfo *VideoInfo `json:"video_info,omitempty"`
	// VideoID 与订单绑定的视频 ID
	VideoID string `json:"video_id,omitempty"`
	// CreatorInfo TikTok 创作者信息
	CreatorInfo *CreatorInfo `json:"creator_info,omitempty"`
	// AdClicks 广告点击数
	AdClicks model.Int64 `json:"ad_clicks,omitempty"`
	// AdConversions 广告转化数
	AdConversions model.Int64 `json:"ad_conversions,omitempty"`
	// AdCost 广告消耗
	AdCost model.Float64 `json:"ad_cost,omitempty"`
	// AdCPA 单次转化成本
	AdCPA model.Float64 `json:"ad_cpa,omitempty"`
	// AdCPC 单次点击成本
	AdCPC model.Float64 `json:"ad_cpc,omitempty"`
	// AdCPM 千次展示成本
	AdCPM model.Float64 `json:"ad_cpm,omitempty"`
	// AdCTR 广告点击率
	AdCTR model.Float64 `json:"ad_ctr,omitempty"`
	// AdCVR 广告转化率
	AdCVR model.Float64 `json:"ad_cvr,omitempty"`
	// AdImpressions 广告展示次数
	AdImpressions model.Int64 `json:"ad_impressions,omitempty"`
	// AdROAS 广告支出回报率
	AdROAS model.Float64 `json:"ad_roas,omitempty"`
	// EngagementCount 点赞数 + 分享次数 + 评论数
	EngagementCount model.Int64 `json:"engagement_count,omitempty"`
	// EngagementRate 互动率。计算公式：engagement_count / video_views
	EngagementRate model.Float64 `json:"engagement_rate,omitempty"`
	// EngagementRateOrganic 自然流量互动率
	EngagementRateOrganic model.Float64 `json:"engagement_rate_organic,omitempty"`
	// EngagementRatePaid 付费流量互动率
	EngagementRatePaid model.Float64 `json:"engagement_rate_paid,omitempty"`
	// EngagementRateBenchmarkP25 相似创作者的原生视频互动率的第25个百分位数。
	// 若互动率高于本字段值，代表对应视频的互动率超过25%的相似创作者的原生视频
	EngagementRateBenchmarkP25 model.Float64 `json:"engagement_rate_benchmark_p25,omitempty"`
	// EngagementRateBenchmarkP50 相似创作者的原生视频互动率的第50个百分位数。
	// 若互动率高于本字段值，代表对应视频的互动率超过50%的相似创作者的原生视频。
	EngagementRateBenchmarkP50 model.Float64 `json:"engagement_rate_benchmark_p50,omitempty"`
	// VideoViews 视频播放次数
	VideoViews model.Int64 `json:"video_views,omitempty"`
	// Likes 点赞数
	Likes model.Int64 `json:"likes,omitempty"`
	// LikesOrganic 自然流量点赞数
	LikesOrganic model.Int64 `json:"likes_organic,omitempty"`
	// LikesPaid 付费流量点赞数
	LikesPaid model.Int64 `json:"likes_paid,omitempty"`
	// Shares 分享次数
	Shares model.Int64 `json:"shares,omitempty"`
	// SharesOrganic 自然流量分享次数
	SharesOrganic model.Int64 `json:"shares_organic,omitempty"`
	// SharesPaid 付费流量分享次数
	SharesPaid model.Int64 `json:"shares_paid,omitempty"`
	// Comments 评论数
	Comments model.Int64 `json:"comments,omitempty"`
	// CommentsOrganic 自然流量评论数
	CommentsOrganic model.Int64 `json:"comments_organic,omitempty"`
	// CommentsPaid 付费流量评论数
	CommentsPaid model.Int64 `json:"comments_paid,omitempty"`
	// LifetimeGenderDistribution 视频播放量按性别的分布情况
	LifetimeGenderDistribution []GenderDistribution `json:"lifetime_gender_distribution,omitempty"`
	// LifetimeTopAgeDistribution 视频播放量按年龄组的分布情况
	LifetimeTopAgeDistribution []AgeDistribution `json:"lifetime_top_age_distribution,omitempty"`
	// LifetimeTopCountryDistribution 视频播放量按受众地域的分布情况
	LifetimeTopCountryDistribution []CountryDistribution `json:"lifetime_top_country_distribution,omitempty"`
	// LifetimeDeviceDistribution 视频播放量按受众设备的分布情况
	LifetimeDeviceDistribution []DeviceDistribution `json:"lifetime_device_distribution,omitempty"`
	// MostPopularCountryStateDistribution 该视频最受欢迎的国家或地区中的州/省级别地域的分布信息
	MostPopularCountryStateDistribution *CountryStateDistribution `json:"most_popular_country_state_distribution,omitempty"`
	// AudienceGendersDistribution 视频播放量按性别的分布情况
	AudienceGendersDistribution []GenderDistribution `json:"audience_genders_distribution,omitempty"`
	// AudienceCountriesDistribution 视频播放量按受众国家或地区的分布情况
	AudienceCountriesDistribution []CountryDistribution `json:"audience_countries_distribution,omitempty"`
	// AudienceAgeDistribution 视频播放量按年龄组的分布情况
	AudienceAgeDistribution []AgeDistribution `json:"audience_age_distribution,omitempty"`
	// AudienceDeviceDistribution 视频播放量按受众设备的分布情况
	AudienceDeviceDistribution []DeviceDistribution `json:"audience_device_distribution,omitempty"`
	// AudienceLocaleDistribution 视频播放量按州/省级地域的分布情况
	AudienceLocaleDistribution []LocaleDistribution `json:"audience_locale_distribution,omitempty"`
	// AudienceLanguageDistribution 视频播放量按受众语言的分布情况
	AudienceLanguageDistribution []LanguageDistribution `json:"audience_language_distribution,omitempty"`
	// AudienceInterestDistribution 占比前10的受众兴趣分布
	AudienceInterestDistribution []AudienceInterestDistribution `json:"audience_interest_distribution,omitempty"`
	// Reach 覆盖人数。观看过该视频至少一次的独立用户数，此数值为估算值
	Reach model.Int64 `json:"reach,omitempty"`
	// ReachOrganic 自然流量覆盖人数
	ReachOrganic model.Int64 `json:"reach_organic,omitempty"`
	// ReachPaid 付费流量覆盖人数
	ReachPaid model.Int64 `json:"reach_paid,omitempty"`
	// VideoCompletionRate 视频完播率 = 视频完播次数/视频总播放次数
	VideoCompletionRate model.Float64 `json:"video_completion_rate,omitempty"`
	// VideoCompletionRateOrganic 自然流量视频完播率
	VideoCompletionRateOrganic model.Float64 `json:"video_completion_rate_organic,omitempty"`
	// VideoCompletionRatePaid 付费流量视频完播率
	VideoCompletionRatePaid model.Float64 `json:"video_completion_rate_paid,omitempty"`
	// VideoCompletionRateBenchmarkP25 相似创作者的原生视频完播率的第25个百分位数。
	// 若完播率高于本字段值，代表对应视频的完播率超过25%的相似创作者的原生视频。
	VideoCompletionRateBenchmarkP25 model.Float64 `json:"video_completion_rate_benchmark_p25,omitempty"`
	// VideoCompletionRateBenchmarkP50 相似创作者的原生视频完播率的第50个百分位数。
	// 若完播率高于本字段值，代表对应视频的完播率超过50%的相似创作者的原生视频。
	VideoCompletionRateBenchmarkP50 model.Float64 `json:"video_completion_rate_benchmark_p50,omitempty"`
	// VideoViewRetention 每秒视频留存率列表
	VideoViewRetention []model.Float64 `json:"video_view_retention,omitempty"`
	// TotalPlayTime 视频总播放时间，单位为秒
	TotalPlayTime model.Int64 `json:"total_play_time,omitempty"`
	// AverageViewTime 视频平均播放时间，单位为秒
	AverageViewTime model.Float64 `json:"average_view_time,omitempty"`
	// AverageViewTimeOrganic 自然流量平均播放时间，单位为秒
	AverageViewTimeOrganic model.Float64 `json:"average_view_time_organic,omitempty"`
	// AverageViewTimePaid 付费流量平均播放时间，单位为秒
	AverageViewTimePaid model.Float64 `json:"average_view_time_paid,omitempty"`
	// TwoSecondsViews 视频播放达2秒次数 / 视频总播放次数
	TwoSecondsViews model.Float64 `json:"two_seconds_views,omitempty"`
	// OrganicTwoSecondsViews 自然流量 2 秒播放率
	OrganicTwoSecondsViews model.Float64 `json:"organic_two_seconds_views,omitempty"`
	// PaidTwoSecondsViews 付费流量 2 秒播放率
	PaidTwoSecondsViews model.Float64 `json:"paid_two_seconds_views,omitempty"`
	// SixSecondsViews 视频播放达6秒次数 / 视频总播放次数
	SixSecondsViews model.Float64 `json:"six_seconds_views,omitempty"`
	// OrganicSixSecondsViews 自然流量 6 秒播放率
	OrganicSixSecondsViews model.Float64 `json:"organic_six_seconds_views,omitempty"`
	// PaidSixSecondsViews 付费流量 6 秒播放率
	PaidSixSecondsViews model.Float64 `json:"paid_six_seconds_views,omitempty"`
	// VideoViewsBySource 视频播放信息
	VideoViewsBySource *VideoViewsBySource `json:"video_views_by_source,omitempty"`
	// VideoViewsOrganic 自然流量产生的视频播放总量。
	// 即通过自然流量，而非Spark Ads产生的视频播放总量。
	VideoViewsOrganic model.Int64 `json:"video_views_organic,omitempty"`
	// VideoViewsPaid 付费流量产生的视频播放总量。
	// 即通过Spark Ads产生的视频播放总量
	VideoViewsPaid model.Int64 `json:"video_views_paid,omitempty"`
	// VideoBoostedDate 视频推广日期。即该视频首次推广的日期。格式为："YYYY-MM-DD HH:MM:SS"（UTC+0时间）。例如，"2023-01-01 00:00:01"。
	// 注意：若视频未经过推广，则本字段会显示默认值 "1970-01-01"
	VideoBoostedDate model.DateTime `json:"video_boosted_date,omitzero"`
	// AnchorID 与推广系列关联的 anchor ID
	AnchorID string `json:"anchor_id,omitempty"`
	// AnchorMetrics Anchor 指标
	AnchorMetrics *AnchorMetrics `json:"anchor_metrics,omitempty"`
	// DailyStats 每日细分数据。仅当请求中传入start_date和 end_date时返回。
	// 注意: 缺失的指标数据不会进行默认值（如0）填充，因此若当日数据中无某一指标，代表该指标无有效数据。
	DailyStats []DailyStat `json:"daily_stats,omitempty"`
}

// GenderDistribution 视频播放量按性别的分布情况
type GenderDistribution struct {
	// Gender 性别。枚举值：FEMALE，MALE
	Gender enum.Gender `json:"gender,omitempty"`
	// Percentage 百分比
	Percentage model.Float64 `json:"percentage,omitempty"`
}

// AgeDistribution 视频播放量按年龄组的分布情况
type AgeDistribution struct {
	// Age 年龄组。枚举值：18-24，25-34，35+
	Age string `json:"age,omitempty"`
	// Percentage 百分比
	Percentage model.Float64 `json:"percentage,omitempty"`
}

// CountryDistribution 视频播放量按受众地域的分布情况
type CountryDistribution struct {
	// Country 两个字母组成的国家或地区缩写
	Country string `json:"country,omitempty"`
	// Percentage 百分比
	Percentage model.Float64 `json:"percentage,omitempty"`
}

// DeviceDistribution 视频播放量按受众设备的分布情况
type DeviceDistribution struct {
	// Device 设备平台。示例：ios 或 android
	Device string `json:"device,omitempty"`
	// Percentage 百分比
	Percentage model.Float64 `json:"percentage,omitempty"`
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
	Percentage model.Float64 `json:"percentage,omitempty"`
}

// AudienceInterestDistribution 占比前10的受众兴趣分布
type AudienceInterestDistribution struct {
	// InterestTag 受众兴趣类别名称
	InterestTag string `json:"interest_tag,omitempty"`
	// Percentage 百分比
	Percentage model.Float64 `json:"percentage,omitempty"`
}

// LanguageDistribution 视频播放量按受众语言的分布情况
type LanguageDistribution struct {
	// Language 受众浏览语言代码
	Language string `json:"language,omitempty"`
	// Percentage 百分比
	Percentage model.Float64 `json:"percentage,omitempty"`
}

// VideoViewsBySource 视频播放信息
type VideoViewsBySource struct {
	// ForYou 从推荐页产生的视频播放次数
	ForYou model.Int64 `json:"for_you,omitempty"`
	// Hashtag 从话题标签产生的视频播放次数
	Hashtag model.Int64 `json:"hashtag,omitempty"`
	// Sound 从视频声音产生的视频播放次数
	Sound model.Int64 `json:"sound,omitempty"`
	// PersonalProfile 从个人主页产生的视频播放次数
	PersonalProfile model.Int64 `json:"personal_profile,omitempty"`
	// Search 从搜索产生的视频播放次数
	Search model.Int64 `json:"search,omitempty"`
	// Following 从关注页面产生的视频播放次数
	Following model.Int64 `json:"following,omitempty"`
	// Other 其他来源产生的视频播放次数
	Other model.Int64 `json:"other,omitempty"`
}

// UnmarshalJSON supports both the documented object shape and the example array shape.
func (v *VideoViewsBySource) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	if data[0] != '[' {
		type alias VideoViewsBySource
		return json.Unmarshal(data, (*alias)(v))
	}

	var items []struct {
		Source     string      `json:"source,omitempty"`
		VideoViews model.Int64 `json:"video_views,omitempty"`
	}
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}
	for _, item := range items {
		switch item.Source {
		case "for_you":
			v.ForYou = item.VideoViews
		case "hashtag":
			v.Hashtag = item.VideoViews
		case "sound":
			v.Sound = item.VideoViews
		case "personal_profile":
			v.PersonalProfile = item.VideoViews
		case "search":
			v.Search = item.VideoViews
		case "following":
			v.Following = item.VideoViews
		case "other":
			v.Other = item.VideoViews
		}
	}
	return nil
}

// AnchorMetrics Anchor 指标
type AnchorMetrics struct {
	// AnchorViews Anchor 展示次数
	AnchorViews model.Int64 `json:"anchor_views,omitempty"`
	// AnchorClicks Anchor 点击次数
	AnchorClicks model.Int64 `json:"anchor_clicks,omitempty"`
	// AnchorUniqueViews Anchor 独立展示人数
	AnchorUniqueViews model.Int64 `json:"anchor_unique_views,omitempty"`
	// AnchorUniqueClicks Anchor 独立点击人数
	AnchorUniqueClicks model.Int64 `json:"anchor_unique_clicks,omitempty"`
}

// DailyStat 每日细分数据
type DailyStat struct {
	// Date 日期
	Date string `json:"date,omitempty"`
	// Reach 覆盖人数。当日观看了该推广系列下至少一个视频的人数
	Reach model.Int64 `json:"reach,omitempty"`
	// EngagementRate 互动次数/ 视频总播放次数
	EngagementRate model.Float64 `json:"engagement_rate,omitempty"`
	// EngagementRateOrganic 自然流量互动率
	EngagementRateOrganic model.Float64 `json:"engagement_rate_organic,omitempty"`
	// EngagementRatePaid 付费流量互动率
	EngagementRatePaid model.Float64 `json:"engagement_rate_paid,omitempty"`
	// VideoCompletionRate 视频完播率 = 视频完播次数/视频总播放次数
	VideoCompletionRate model.Float64 `json:"video_completion_rate,omitempty"`
	// Likes 点赞数
	Likes model.Int64 `json:"likes,omitempty"`
	// LikesOrganic 自然流量点赞数
	LikesOrganic model.Int64 `json:"likes_organic,omitempty"`
	// LikesPaid 付费流量点赞数
	LikesPaid model.Int64 `json:"likes_paid,omitempty"`
	// Shares 分享次数
	Shares model.Int64 `json:"shares,omitempty"`
	// SharesOrganic 自然流量分享次数
	SharesOrganic model.Int64 `json:"shares_organic,omitempty"`
	// SharesPaid 付费流量分享次数
	SharesPaid model.Int64 `json:"shares_paid,omitempty"`
	// Comments 评论数
	Comments model.Int64 `json:"comments,omitempty"`
	// CommentsOrganic 自然流量评论数
	CommentsOrganic model.Int64 `json:"comments_organic,omitempty"`
	// CommentsPaid 付费流量评论数
	CommentsPaid model.Int64 `json:"comments_paid,omitempty"`
	// Views 播放次数
	Views model.Int64 `json:"views,omitempty"`
	// VideoViewsOrganic 当日自然流量产生的视频播放量。
	// 即通过自然流量，而非Spark Ads产生的当日视频播放量。
	VideoViewsOrganic model.Int64 `json:"video_views_organic,omitempty"`
	// VideoViewsPaid 当日付费流量产生的视频播放量。
	// 即通过Spark Ads产生的当日视频播放量
	VideoViewsPaid model.Int64 `json:"video_views_paid,omitempty"`
}
