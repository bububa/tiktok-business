package enum

// RegionLevelRange 想要返回的地域层级
type RegionLevelRange string

const (
	// RegionLevelRange_ALL: 返回所有层级的地域。
	RegionLevelRange_ALL RegionLevelRange = "ALL"
	// RegionLevelRange_TO_COUNTRY：只返回国家或地区级地域。
	RegionLevelRange_TO_COUNTRY RegionLevelRange = "TO_COUNTRY"
	// RegionLevelRange_TO_PROVINCE：返回国家和省级地域。指定市场区域和都会区属于省级地域。
	RegionLevelRange_TO_PROVINCE RegionLevelRange = "TO_PROVINCE"
	// RegionLevelRange_TO_CITY: 返回国家、省、市级地域。
	RegionLevelRange_TO_CITY RegionLevelRange = "TO_CITY"
	// RegionLevelRange_TO_DISTRICT: 返回国家、省、市、区级地域。
	RegionLevelRange_TO_DISTRICT RegionLevelRange = "TO_DISTRICT"
)
