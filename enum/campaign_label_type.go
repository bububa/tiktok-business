package enum

// CampaignLabelType 推广系列标签类型
type CampaignLabelType string

const (
	// CampaignLabelType_GENERAL：通用标签。
	CampaignLabelType_GENERAL CampaignLabelType = "GENERAL"
	// CampaignLabelType_MARKETING_EVENT：营销事件标签。
	CampaignLabelType_MARKETING_EVENT CampaignLabelType = "MARKETING_EVENT"
)
