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
}
