package report

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Dimensions 维度是数据的一项属性，用于进行数据组合。基础报表支持以下维度
type Dimensions struct {
	// AdvertiserID 广告主ID。当dimensions包含 advertiser_id 时返回
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列ID。当dimensions包含 campaign_id 时返回
	CampaignID string `json:"campaign_id,omitempty"`
	// AdgroupID 广告组ID。当dimensions包含 adgroup_id 时返回
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdID 广告ID。当dimensions包含 ad_id 时返回
	AdID string `json:"ad_id,omitempty"`
	// MaterialID 素材 ID
	MaterialID string `json:"material_id,omitempty"`
	// PageID 按页面 ID 分组。
	// 注意：
	// 您可以查看下文的维度组合小节了解允许的组合。
	// 本维度支持的指标详见基础报表中维度所支持指标 - page_id支持的指标。
	PageID string `json:"page_id,omitempty"`
	// ComponentName 按 TikTok 即时表单组件名称分组。
	// 注意：
	// 您可以查看下文的维度组合小节了解允许的组合。
	// 本维度支持的指标详见基础报表中维度所支持指标 - component_name支持的指标。
	ComponentName string `json:"component_name,omitempty"`
	// RoomID 按直播间 ID 分组。
	// 注意：
	// 您可以查看下文的维度组合小节了解允许的组合。
	// 本维度支持的指标详见基础报表中维度所支持指标 - room_id支持的指标。
	RoomID string `json:"room_id,omitempty"`
	// PostID 按帖子 ID 分组。
	// 注意：
	// 您可以查看下文的维度组合小节了解允许的组合。
	// 本维度支持的指标详见基础报表中维度所支持指标 - post_id支持的指标。
	PostID string `json:"post_id,omitempty"`
	// ImageID 图片 ID
	ImageID string `json:"image_id,omitempty"`
	// VideoID 视频 ID
	VideoID string `json:"video_id,omitempty"`
	// VideoMaterialID 按视频素材 ID 分组。
	// 注意：
	// 您可以查看下文的维度组合小节了解允许的组合。
	// 本维度支持所有数据层级（data_level= AUCTION_CAMPAIGN ， AUCTION_ADGROUP，AUCTION_AD 或AUCTION_ADVERTISER）。
	// 本维度不支持下列指标：
	// 独立访客（UV）类指标：reach，cost_per_1000_reached，frequency以及当 optimization_goal（优化目标）为REACH时的成效指标：
	// result，cost_per_result，result_rate
	// real_time_result， real_time_cost_per_result，real_time_result_rate
	// secondary_goal_result，cost_per_secondary_goal_result，secondary_goal_result_rate
	// skan_result，skan_cost_per_result，skan_result_rate
	// 直播指标：live_views，live_unique_views，live_effective_views，live_product_clicks
	// video_views_p100
	VideoMaterialID string `json:"video_material_id,omitempty"`
	// Source 素材来源。枚举值见：枚举值 - 素材来源
	Source enum.MaterialSource `json:"source,omitempty"`
	// CatalogID 仅当 material_type 为 SPC_CATALOG_VIDEO 或SPC_CATALOG_CAROUSEL 时返回。
	// 商品库 ID。
	CatalogID string `json:"catalog_id,omitempty"`
	// StatTimeDay 消耗发生的时间（天）。格式：2020-01-01 00:00:00
	StatTimeDay model.DateTime `json:"stat_time_day,omitzero"`
	// StatTimeHour 消耗发生的时间（小时）。格式：2020-01-01 10:00:00
	StatTimeHour model.DateTime `json:"stat_time_hour,omitzero"`
	// AC 受众网络类型。详见枚举值-广告管理-网络类型
	AC enum.NetworkType `json:"ac,omitempty"`
	// Age 受众年龄区间。详见枚举值-广告管理-受众年龄区间
	Age enum.Age `json:"age,omitempty"`
	// CountryCode 受众国家或地区代码。详见附录-地区代码
	CountryCode string `json:"country_code,omitempty"`
	// InterestCategory 一级兴趣分类
	InterestCategory string `json:"interest_category,omitempty"`
	// InterestCategoryV2 新版一级兴趣分类
	// 使用/tool/interest_category/接口获取兴趣分类完整列表
	InterestCategoryV2 string `json:"interest_category_v2,omitempty"`
	// Location 2个字母的国家或地区代码
	Location string `json:"location,omitempty"`
	// Gender 受众性别。枚举值: FEMALE，MALE，NONE
	Gender enum.Gender `json:"gender,omitempty"`
	// Language 受众语言。详见枚举值-广告管理-受众语言
	Language string `json:"language,omitempty"`
	// Placement 投放版位，详见枚举值-广告管理-版位
	Placement enum.Placement `json:"placement,omitempty"`
	// Platform 受众操作系统，详见枚举值-广告管理-受众操作系统
	Platform enum.OperatingSystem `json:"platform,omitempty"`
	// ContextualTag 内容相关定向标签
	ContextualTag string `json:"contextual_tag,omitempty"`
	// AdType 按广告类型分组。
	// 您可以使用本维度区分搜索广告和非搜索广告。若为搜索广告，值显示为Search，本维度的值非Search则代表广告不是搜索广告。
	// 注意：
	// 您可以查看下文的维度组合小节了解允许的组合。
	// 本维度支持所有数据层级（data_level= AUCTION_CAMPAIGN ， AUCTION_ADGROUP，AUCTION_AD 或AUCTION_ADVERTISER）。
	// 请查看基础报表中维度所支持指标 - ad_type支持的指标，了解本维度所支持的指标。
	// 在同步基础报表中，我们支持拉取的数据最早至2022年2月22日；在异步基础报表中，我们支持拉取的数据最早至2022年3月1日。
	AdType enum.AdType `json:"ad_type,omitempty"`
	// CustomEventType 按自定义事件类型分组。
	// 本维度支持以下指标：
	// 基础数据指标clicks
	// 全部自定义app事件指标（具体指标详见支持的指标 - 应用内事件数据指标）
	// 全部自定义页面事件指标（具体指标详见支持的指标 - 网页事件数据指标）
	CustomEventType enum.OptimizationEvent `json:"custom_event_type,omitempty"`
	// EventSourceID 按事件源 ID 分组。
	// 注意：您可以查看下文的维度组合小节了解允许的组合。
	// 本维度支持的指标：
	// 全部网页事件数据指标（具体指标详见支持的指标 - 网页事件数据指标）
	// 全部应用内事件数据指标 （具体指标详见支持的指标 - 全部应用内事件数据指标）
	EventSourceID string `json:"event_source_id,omitempty"`
}
