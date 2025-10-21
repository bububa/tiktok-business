package product

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// MiniSeriesProduct 短剧商品
type MiniSeriesProduct struct {
	// SeriesID 短剧的自定义唯一 ID。
	// 长度限制：100 字符。
	SeriesID string `json:"series_id,omitempty" csv:"series_id"`
	// SeriesName 短剧名称。
	// 长度限制：100 字符。
	SeriesName string `json:"series_name,omitempty" csv:"series_name"`
	// ImageURL 短剧封面图的 URL。
	// 图片规格必须大于等于500 x 500像素，否则图片将被过滤，商品无法通过审核。详细信息请参考商品图片要求。
	// 所有图片应为 JPG 或 PNG 格式。
	// 示例： https://www.tiktok.com/short_drama_series_image_001.jpg。
	ImageURL string `json:"image_url,omitempty" csv:"image_link"`
	// Recharge 短剧收费详情
	Recharge []MiniSeriesRecharge `json:"recharge,omitempty" csv:"recharge"`
	// ProductDetail 短剧的其他信息
	ProductDetail *MiniSeriesProductDetail `json:"mini_series_product_detail,omitempty" csv:""`
}

// MiniSeriesRecharge 短剧收费详情
type MiniSeriesRecharge struct {
	// Type 短剧收费类型。
	// 枚举值：
	// BY_TIERS: 按级别收费。
	// SUBSCRIPTION: 按周期订阅收费。
	// BY_EPISODES: 按集数收费。
	Type enum.MiniSeriesType `json:"type,omitempty" csv:"type"`
	// PurchaseUnit 收费单位。
	// 当 recharge_type 为 BY_TIERS 时，枚举值为：
	// TIER 1: 1 级。
	// TIER_2: 2 级。
	// TIER 3: 3 级。
	// TIER 4: 4 级。
	// TIER 5: 5 级。
	// 当 recharge_type 为 SUBSCRIPTION，枚举值为：
	// WEEKLY: 按周订阅。
	// MONTHLY: 按月订阅。
	// QUARTERLY: 按季度订阅。
	// YEARLY: 按年订阅。
	// 当 recharge_type 为 BY_EPISODES，需指定一个或多个数字字符串，字符串代表的数字不可大于总集数。
	PurchaseUnit model.CSVStringList `json:"purchase_unit,omitempty" csv:"purchase_unit"`
	// Cost 级别、订阅周期或集数对应的价格。
	Cost model.CSVFloatList `json:"cost,omitempty" csv:"cost"`
}

// MiniSeriesProductDetail 短剧的其他信息
type MiniSeriesProductDetail struct {
	// CompanyType 短剧的公司类型。
	// 枚举值：
	// COPYRIGHT_HOLDER: 版权方。
	// DISTRIBUTOR: 发行方。
	CompanyType enum.MiniSeriesCompanyType `json:"company_type,omitempty" csv:"company_type"`
	// CopyrightHolderName 指定 company_type 时必填。
	// 版权方公司名称。
	// 长度限制：100 个字符。
	CopyrightHolderName string `json:"copyright_holder_name,omitempty" csv:"copyright_holder_name"`
	// AppID 需至少传入 app_id 和 minis_id 其中之一或同时传入 app_id 和 minis_id。
	// 广告中要推广的应用的应用 ID。
	// 若想获取您的广告账户中的应用 ID 列表，可使用 /app/list/。
	AppID string `json:"app_id,omitempty" csv:"app_id"`
	// MinisID 需至少传入 app_id 和 minis_id 其中之一或同时传入 app_id 和 minis_id。
	// 小程序 ID。
	// 需传入您的小程序在 TikTok 开放平台上的 ID。
	MinisID string `json:"minis_id,omitempty" csv:"minis_id"`
	// TotalEpisodes 短剧的总集数。
	// 取值范围：1-9,999。
	TotalEpisodes int `json:"total_episodes,omitempty" csv:"total_episodes"`
	// InitialPaidEpisodes 短剧付费起始集数。例如，若第 11 集起需付费，则该字段需设置为 11。
	// 取值范围：1-9,999。
	// 本字段的值需小于 total_episodes 的值。
	InitialPaidEpisodes int `json:"initial_paid_episodes,omitempty" csv:"initial_paid_espisodes"`
	// PerEpisodeDuration 每集的平均时长，单位为分钟。
	// 取值范围：0-15。
	PerEpisodeDuration int `json:"per_episode_duration" csv:"per_episode_duration"`
	// SpokenLanguage 短剧原声语言。
	// 枚举值：
	// ar: 阿拉伯语。
	// de: 德语。
	// cs: 捷克语。
	// en: 英语。
	// es: 西班牙语。
	// fr: 法语。
	// he: 希伯来语。
	// hi: 印地语。
	// id: 印尼语。
	// it: 意大利语。
	// ja: 日语。
	// ko: 韩语。
	// ms: 马来语。
	// pl: 波兰语。
	// pt: 葡萄牙语。
	// ro: 罗马尼亚语。
	// ru: 俄语。
	// ta: 泰米尔语。
	// th: 泰语。
	// tr: 土耳其语。
	// vi: 越南语。
	// zh: 简体中文。
	// zh-hant: 繁体中文。
	SpokenLanguage string `json:"spoken_language,omitempty" csv:"spoken_language"`
	// SubtitleLanguage 短剧字幕语言。
	// 枚举值：
	// ar: 阿拉伯语。
	// de: 德语。
	// cs: 捷克语。
	// en: 英语。
	// es: 西班牙语。
	// fr: 法语。
	// he: 希伯来语。
	// hi: 印地语。
	// id: 印尼语。
	// it: 意大利语。
	// ja: 日语。
	// ko: 韩语。
	// ms: 马来语。
	// pl: 波兰语。
	// pt: 葡萄牙语。
	// ro: 罗马尼亚语。
	// ru: 俄语。
	// ta: 泰米尔语。
	// th: 泰语。
	// tr: 土耳其语。
	// vi: 越南语。
	// zh: 简体中文。
	// zh-hant: 繁体中文。
	SubtitleLanguage string `json:"subtitle_language,omitempty" csv:"subtitle_language"`
	// ProductionType 制作类型，即短剧是本地短剧还是翻译短剧。
	// 枚举值：
	// LOCAL: 本地短剧。本地演员、本地语言、旁白和字幕。
	// TRANSLATION: 翻译短剧，为国际发行，附加字幕、旁白和配音。
	ProductionType enum.MiniSeriesProductionType `json:"production_type,omitempty" csv:"production_type"`
	// TargetAudience 短剧的目标受众。
	// 枚举值：
	// MALE: 男频。目标受众主体为男性，更多地围绕男主人公或从男性视角叙事强调权力财富声望的积累、强大能力的提升、运筹帷幄与设计营谋。
	// FEMALE: 女频。目标受众主体为女性，更多地围绕女主人公或从女性视角叙事，大多含有浪漫爱情元素。
	// NEUTRAL: 中立。没有明确目标受众性别。
	TargetAudience enum.MiniSeriesTargetAudience `json:"target_audience,omitempty" csv:"target_audience"`
	// Characters 短剧的角色列表。
	// 最大数量：3。
	// 枚举值参见 characters 的可选值。
	Characters model.CSVStringList `json:"characters,omitempty" csv:"characters"`
	// Genres 短剧的类别，通常基于主线剧情或情节。
	// 最大数量：3。
	// 枚举值参见 genres 的可选值。
	Genres model.CSVStringList `json:"genres,omitempty" csv:"genres"`
	// HistoricalContext 短剧的时代背景或设定。
	// 最大数量：3。
	// 枚举值参见 historical_context 的可选值。
	HistoricalContext model.CSVStringList `json:"historical_context,omitempty" csv:"historical_context"`
	// Actors 短剧中的演员的名字。
	// 最大数量：20。
	// 每个名字的长度限制：100 个字符。
	Actors model.CSVStringList `json:"actors,omitempty" csv:"actors"`
}

func (p MiniSeriesProduct) CatalogType() enum.CatalogType {
	return enum.CatalogType_MINI_SERIES
}
