package enum

// RfCampaignType 仅当推广目标（objective_type）设置为RF_REACH时，可根据本字段判断具体的合约推广系列类型。
type RfCampaignType string

const (
	// RfCampaignType_STANDARD：覆盖和频次推广系列。
	RfCampaignType_STANDARD RfCampaignType = "STANDARD"
	// RfCampaignType_PULSE：TikTok Pulse 推广系列。
	RfCampaignType_PULSE RfCampaignType = "PULSE"
	// RfCampaignType_TOPVIEW： TopView 推广系列。
	RfCampaignType_TOPVIEW RfCampaignType = "TOPVIEW"
)
