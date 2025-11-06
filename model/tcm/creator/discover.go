package creator

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DiscoverRequest 发现创作者 API Request
type DiscoverRequest struct {
	// TCMAccountID TikTok Creator Marketplace 账户 ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// CountryCode 创作者所在的区域。枚举值：US，GB，IN，TW，DE，FR，IT，JP，KR，VN，TH，MY，PH，SG，ID，CA，AU，RU，HK，CN，SA，EG，AE，IL，TR，BR，ES。
	CountryCode string `json:"country_code,omitempty"`
	// MajorAudienceCountryCode 主体粉丝所在的区域。枚举值参见 country_code
	MajorAudienceCountryCode string `json:"major_audience_country_code,omitempty"`
	// MinFollowers 创作者的最低粉丝数。
	// 取值范围：100 - 10,000,000,000。
	MinFollowers int64 `json:"min_followers,omitempty"`
	// MaxFollowers 创作者的最高粉丝数。
	// 取值范围：100 - 10,000,000,000。
	// 若同时传入min_followers 和 max_followers，max_followers 的值需大于 min_followers。若为 min_followers 和 max_followers 指定相同值，将返回空结果。
	MaxFollowers int64 `json:"max_followers,omitempty"`
	// MinAvgViews 创作者视频获得的最低平均播放量，应大于 0。请注意，此平均数根据最近的 30 个视频计算得出
	MinAvgViews int64 `json:"min_avg_views,omitempty"`
	// MaxAvgViews 创作者视频获得的最高平均播放量，此数字不能超过 1,000 亿次。请注意，此平均数根据最近的 30 个视频计算得出
	MaxAvgViews int64 `json:"max_avg_views,omitempty"`
	// MajorAudienceGender 主体粉丝的性别。枚举值：male，female
	MajorAudienceGender string `json:"major_audience_gender,omitempty"`
	// MajorAudienceAge 主体粉丝的年龄组。
	// 枚举值：18-24，25-34，35-44，45-54，55+，35+。
	// 若想以大部分粉丝所属年龄组为 35 岁及以上中特定细分年龄组精细搜索创作者，请使用枚举值35-44，45-54 或55+。
	// 若想以大部分粉丝所属年龄组为 35 岁及以上笼统搜索创作者，请使用枚举值35+。
	MajorAudienceAge string `json:"major_audience_age,omitempty"`
	// EngagementRate 互动率 (%) 。广告在投放期间的总互动（比如点赞，分享，和评论）次数除以视频播放次数所得的比值，取最近30个视频的平均值。 枚举值：0-5, 5-10, 10-15, 15-20, >20。若未设置，则返回所有互动率的结果
	EngagementRate string `json:"engagement_rate,omitempty"`
	// MobileOS 希望定向的设备操作系统。 枚举值：android，ios。若未设置，则返回结果包含android和ios对应的结果
	MobileOS string `json:"mobile_os,omitempty"`
	// Languages 用于筛选结果的创作者语言。
	// 可传入的值参见下文 “languages可选值” 小节。
	// 若未设置，则默认返回所有支持语言的结果
	Languages []string `json:"languages,omitempty"`
	// CreatorCategories 用于筛选结果的创作者类别。
	// 可传入的值参见下文 “creator_categories可选值”小节。
	// 若设置了多个类别，则返回将所设置的任一类别的匹配结果。
	// 若未设置，则返回所有创作者类别的匹配结果。
	CreatorCategories []string `json:"creator_categories,omitempty"`
	// MaxCampaignCount 用于筛选结果的创作者已完成推广系列数量的最大值。
	// 取值范围：[0,100]。
	// 本字段无默认值
	MaxCampaignCount int `json:"max_campaign_count,omitempty"`
	// MinCampaignCount 用于筛选结果的创作者已完成推广系列数量的最小值。
	// 取值范围：[0,100]。
	// 本字段无默认值。
	// min_campaign_count 的值不可大于max_campaign_count的值
	MinCampaignCount int `json:"min_campaign_count,omitempty"`
	// HasResponsiveBadge 是否筛选拥有“响应迅速”的徽章的创作者。“响应迅速”徽章代表创作者通常在48小时内回复品牌方。
	// 枚举值：
	// true：筛选拥有“响应迅速”的徽章的创作者。
	// false：不筛选拥有“响应迅速”的徽章的创作者。
	// 若未设置，则返回结果默认不筛选创作者是否拥有“响应迅速”的徽章。
	HasResponsiveBadge bool `json:"has_responsive_badge,omitempty"`
	// Query 用于检索创作者或视频的关键词
	Query string `json:"query,omitempty"`
	// SortField 您可以使用此字段对返回的数据进行排序。
	// 枚举值：
	// follower_count：按照粉丝数对创作者进行排序。创作者的粉丝数越多，在返回列表中越靠前。
	// avg_views：按照视频平均播放量对创作者进行排序。创作者的视频平均播放量越高，在返回列表中越靠前。
	// engagement_rate：按照视频互动率对创作者进行排序。创作者的互动率越高，在返回列表中越靠前。
	// relevance：按照相关性对创作者进行排序。创作者的个人主页点击率越高，在返回列表中越靠前。
	// 默认值：follower_count 。
	SortField string `json:"sort_field,omitempty"`
	// Page 当前页码。
	// 默认值： 1
	Page int `json:"page,omitempty"`
	// PageSize 当前页面大小。
	// 默认值：10。
	// 取值范围：1-200
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements PostRequest
func (r *DiscoverRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DiscoverResponse 发现创作者 API Response
type DiscoverResponse struct {
	model.BaseResponse
	Data *DiscoverResult `json:"data,omitempty"`
}

type DiscoverResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Creators 返回的相似创作者。
	// 注意：目前仅会返回最多 250 个创作者的数据。
	Creators []Creator `json:"creators,omitempty"`
}
