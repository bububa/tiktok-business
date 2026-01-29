package enum

// SmartPlusAdgroupMode The ad group mode for the campaign.
type SmartPlusAdgroupMode string

const (
	// SmartPlusAdgroupMode_SINGLE: You can only create one single ad group within the campaign.
	SmartPlusAdgroupMode_SINGLE SmartPlusAdgroupMode = "SINGLE"
	// SmartPlusAdgroupMode_MULTIPLE: You can only create multiple ad groups within the campaign.
	SmartPlusAdgroupMode_MULTIPLE SmartPlusAdgroupMode = "MULTIPLE"
)
