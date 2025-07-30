package enum

// NetworkType 网络类型
type NetworkType string

const (
	// WIFI wifi
	WIFI NetworkType = "WIFI"
	// NetworkType_2G 2G
	NetworkType_2G NetworkType = "2G"
	// NetworkType_3G 3G
	NetworkType_3G NetworkType = "3G"
	// NetworkType_4G 4G
	NetworkType_4G NetworkType = "4G"
	// NetworkType_5G 5G
	NetworkType_5G NetworkType = "5G"
	// NetworkType_Unlimited 不限
	NetworkType_Unlimited NetworkType = "unlimited"
)
