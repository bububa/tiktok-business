package enum

// RegionLevel 地域层级
type RegionLevel string

const (
	// RegionLevel_COUNTRY（国家或地区级）
	RegionLevel_COUNTRY RegionLevel = "COUNTRY"
	// RegionLevel_PROVINCE (省/州)
	RegionLevel_PROVINCE RegionLevel = "PROVINCE"
	// RegionLevel_CITY（市或郡/县级）
	RegionLevel_CITY RegionLevel = "CITY"
	// RegionLevel_DISTRICT（区或市级）
	RegionLevel_DISTRICT RegionLevel = "DISTRICT"
)
