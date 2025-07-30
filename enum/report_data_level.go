package enum

// ReportDataLevel 您想要查询的报表数据层级。
type ReportDataLevel string

const (
	// ReportDataLevel_AUCTION_AD：竞价广告或竞价和合约广告，广告层级。
	ReportDataLevel_AUCTION_AD ReportDataLevel = "AUCTION_AD"
	// ReportDataLevel_AUCTION_ADGROUP：竞价广告或竞价和合约广告，广告组层级。
	ReportDataLevel_AUCTION_ADGROUP ReportDataLevel = "AUCTION_ADGROUP"
	// ReportDataLevel_AUCTION_CAMPAIGN：竞价广告或竞价和合约广告，推广系列层级。
	ReportDataLevel_AUCTION_CAMPAIGN ReportDataLevel = "AUCTION_CAMPAIGN"
	// ReportDataLevel_AUCTION_ADVERTISER：竞价广告或竞价和合约广告，广告主层级。
	ReportDataLevel_AUCTION_ADVERTISER ReportDataLevel = "AUCTION_ADVERTISER"
	// ReportDataLevel_RESERVATION_AD（已废弃）：合约广告，广告层级。
	ReportDataLevel_RESERVATION_AD ReportDataLevel = "RESERVATION_AD"
	// ReportDataLevel_RESERVATION_ADGROUP（已废弃）：合约广告，广告组层级。
	ReportDataLevel_RESERVATION_ADGROUP ReportDataLevel = "RESERVATION_ADGROUP"
	// ReportDataLevel_RESERVATION_CAMPAIGN（已废弃）：合约广告，推广系列层级。
	ReportDataLevel_RESERVATION_CAMPAIGN ReportDataLevel = "RESERVATION_CAMPAIGN"
	// ReportDataLevel_RESERVATION_ADVERTISER（已废弃）：合约广告，广告主层级。
	ReportDataLevel_RESERVATION_ADVERTISER ReportDataLevel = "RESERVATION_ADVERTISER"
)
