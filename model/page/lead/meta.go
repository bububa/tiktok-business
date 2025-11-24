package lead

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Meta 线索元信息
type Meta struct {
	// LeadSource 线索来源。
	// 枚举值：
	// INSTANT_FORM：通过即时表单生成的线索。
	// DIRECT_MESSAGE：通过绑定的企业号的私信生成的线索。
	LeadSource enum.LeadSource `json:"lead_source,omitempty"`
	// LeadID 线索 ID
	LeadID string `json:"lead_id,omitempty"`
	// PageID 仅当 lead_source 为 INSTANT_FORM 时返回本字段。
	// 即时表单的页面 ID
	PageID string `json:"page_id,omitempty"`
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
	// CreateTime 测试线索的创建时间（UNIX时间戳
	CreateTime model.DateTime `json:"create_time,omitzero"`
}
