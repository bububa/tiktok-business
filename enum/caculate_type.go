package enum

// CalculateType 加密类型枚举值详见枚举值-加密类型
type CalculateType string

const (
	// CalculateType_FIRST_SHA256 以SHA256加密的IDFA或GAID。
	CalculateType_FIRST_SHA256 CalculateType = "FIRST_SHA256"
	// CalculateType_FIRST_MD5 以MD5加密的IDFA或GAID。
	CalculateType_FIRST_MD5 CalculateType = "FIRST_MD5"
	// CalculateType_EMAIL_SHA256 以SHA256加密的邮件地址。
	CalculateType_EMAIL_SHA256 CalculateType = "EMAIL_SHA256"
	// CalculateType_PHONE_SHA256 以SHA256加密电话号码。
	CalculateType_PHONE_SHA256 CalculateType = "PHONE_SHA256"
	// CalculateType_IDFA_SHA256 以SHA256加密的IDFA（即Apple Identifiers for Advertisers，Apple面向广告主的标识符）。
	CalculateType_IDFA_SHA256 CalculateType = "IDFA_SHA256"
	// CalculateType_IDFA_MD5 以MD5加密的IDFA。
	CalculateType_IDFA_MD5 CalculateType = "IDFA_MD5"
	// CalculateType_GAID_SHA256 以SHA256加密的GAID（即Google Advertising ID，谷歌广告ID）
	CalculateType_GAID_SHA256 CalculateType = "GAID_SHA256"
	// CalculateType_GAID_MD5 以MD5加密的GAID，
	CalculateType_GAID_MD5 CalculateType = "GAID_MD5"
	// CalculateType_MULTIPLE_TYPES 多种标识符。该加密类型仅适用于使用客户文件API创建的客户文件受众。
	CalculateType_MULTIPLE_TYPES CalculateType = "MULTIPLE_TYPES"
	// CalculateType_FIRST_NORMAL（已废弃）IDFA或GAID原始数据
	CalculateType_FIRST_NORMAL CalculateType = "FIRST_NORMAL"
)
