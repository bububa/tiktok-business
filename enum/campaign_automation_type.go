package enum

// CampaignAutomationType Campaign automation type.
type CampaignAutomationType string

const (
	// CampaignAutomationType_MANUAL: Manual Campaigns.
	CampaignAutomationType_MANUAL CampaignAutomationType = "MANUAL"
	// CampaignAutomationType_SMART_PLUS: Smart+ Campaigns.
	CampaignAutomationType_SMART_PLUS CampaignAutomationType = "SMART_PLUS"
	// CampaignAutomationType_UPGRADED_SMART_PLUS: Upgraded Smart+ Campaigns. To learn more about Upgraded Smart+ Campaigns, see Upgraded Smart+ Campaign.
	CampaignAutomationType_UPGRADED_SMART_PLUS CampaignAutomationType = "UPGRADED_SMART_PLUS"
	// When campaign_automation_type is UPGRADED_SMART_PLUS, the returned ad_id_v2 is the ID of the Upgraded Smart+ Ad.
	// CampaignAutomationType_UPGRADED_SMART_PLUS_CREATIVE: Upgraded Smart+ Campaigns. To learn more about Upgraded Smart+ Campaigns, see Upgraded Smart+ Campaign.
	// When campaign_automation_type is UPGRADED_SMART_PLUS_CREATIVE, the returned ad_id is the ID of a creative in the Upgraded Smart+ Ad.
	CampaignAutomationType_UPGRADED_SMART_PLUS_CREATIVE CampaignAutomationType = "UPGRADED_SMART_PLUS_CREATIVE"
)
