package enum

// BcType 商务中心类型
type BcType string

const (
	// BcType_NORMAL：普通商务中心。
	BcType_NORMAL BcType = "NORMAL"
	// BcType_DIRECT：非自助直客商务中心。
	BcType_DIRECT BcType = "DIRECT"
	// BcType_AGENCY：非自助代理商商务中心。
	BcType_AGENCY BcType = "AGENCY"
	// BcType_SELF_SERVICE：自助直客商务中心。
	BcType_SELF_SERVICE BcType = "SELF_SERVICE"
	// BcType_SELF_SERVICE_AGENCY：自助代理商商务中心
	BcType_SELF_SERVICE_AGENCY BcType = "SELF_SERVICE_AGENCY"
)
