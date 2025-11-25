package enum

// OperatingSystem 受众操作系统
type OperatingSystem string

const (
	// ALL
	OSAll OperatingSystem = "ALL"
	// ANDROID Android
	ANDROID OperatingSystem = "ANDROID"
	// IOS iOS
	IOS OperatingSystem = "IOS"
	// PC 已废弃 PC
	PC OperatingSystem = "PC"
)
