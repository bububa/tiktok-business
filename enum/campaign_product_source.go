package enum

// CampaignProductSource 按推广系列的商品来源筛选。
type CampaignProductSource string

const (
	// CampaignProductSource_CATALOG ：商品库。
	CampaignProductSource_CATALOG CampaignProductSource = "CATALOG"
	// CampaignProductSource_STORE：TikTok Shop 店铺，或 TikTok 橱窗。
	CampaignProductSource_STORE CampaignProductSource = "STORE"
)
