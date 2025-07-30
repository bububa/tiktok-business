package enum

// ServiceType 广告服务类型
type ServiceType string

const (
	// ServiceType_AUCTION：竞价广告，或竞价和合约广告。
	ServiceType_AUCTION ServiceType = "AUCTION"
	// ServiceType_RESERVATION （已废弃）：合约广告。
	ServiceType_RESERVATION ServiceType = "RESERVATION"
)
