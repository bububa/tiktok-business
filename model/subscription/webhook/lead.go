package webhook

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// LeadEntity  线索
type LeadEntity struct {
	// ID 线索 ID
	ID string `json:"id,omitempty"`
	// LeadSource 线索来源。
	// 枚举值：
	// INSTANT_FORM：通过即时表单生成的线索。
	// DIRECT_MESSAGE：通过绑定的企业号的私信生成的线索
	LeadSource enum.LeadSource `json:"lead_source,omitempty"`
	// PageID 即时表单的页面 ID
	PageID string `json:"page_id,omitempty"`
	// PageName 即时表单的页面名称
	PageName string `json:"page_name,omitempty"`
	// AdvertiserID 广告账户 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserName 广告账户名称
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// LibraryID 表单库 ID
	LibraryID string `json:"library_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName 推广系列名称
	CampaignName string `json:"campaign_name,omitempty"`
	// AdgroupID 广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName 广告组名称
	AdgroupName string `json:"adgroup_name,omitempty"`
	// AdID 广告 ID
	AdID string `json:"ad_id,omitempty"`
	// AdName 广告名称
	AdName string `json:"ad_name,omitempty"`
	// CreateTime 线索创建时间，格式为以秒为单位的 Epoch/Unix 时间戳。
	// 示例：1743639757。
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// Changes 即时表单中字段的变化或私信线索的变化
	Changes []struct {
		// Field 变化的字段
		Field string `json:"field,omitempty"`
		// Value 字段的值
		Value string `json:"value,omitempty"`
	} `json:"changes,omitempty"`
}
