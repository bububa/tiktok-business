package enum

type GeoType string

const (
	// GeoType_COUNTRY（国家或地区）
	GeoType_COUNTRY GeoType = "COUNTRY"
	// GeoType_PROVINCE（省）
	GeoType_PROVINCE GeoType = "PROVINCE"
	// GeoType_CITY（城市/郡/县）
	GeoType_CITY GeoType = "CITY"
	// GeoType_DISTRICT（区/城市）
	GeoType_DISTRICT GeoType = "DISTRICT"
	// GeoType_DMA（指定市场区域）
	GeoType_DMA GeoType = "DMA"
	// GeoType_ZIP_CODE（邮政编码对应地域）。
	GeoType_ZIP_CODE GeoType = "ZIP_CODE"
)
