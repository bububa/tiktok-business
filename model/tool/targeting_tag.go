package tool

import "github.com/bububa/tiktok-business/enum"

// TargetingTag 定向标签
type TargetingTag struct {
	// Keyword 用于搜索定向标签的关键词
	Keyword string `json:"keyword,omitempty"`
	// TargetingType 定向标签所适用的定向类型。
	// 枚举值：GEO (地理位置定向，包括行政区域定向和邮政编码定向)
	TargetingType enum.TargetingType `json:"targeting_type,omitempty"`
	// Name 定向标签的名称
	Name string `json:"name,omitempty"`
	// StatusInfo 定向标签的状态信息
	StatusInfo *TargetingTagStatus `json:"status_info,omitempty"`
	// Geo 用于地域定向的定向标签的信息。
	// 当targeting_type为GEO时返回。
	Geo *Geo `json:"geo,omitempty"`
	// ISP 用于互联网服务提供商定向的定向标签的信息。
	// 当targeting_type为ISP时返回。
	ISP *ISP `json:"isp,omitempty"`
}

// TargetingTagStatus 定向标签的状态信息
type TargetingTagStatus struct {
	// Status 定向标签的状态。
	// 枚举值: ENABLED（可用于定向）, DISABLED（不可用于定向）
	Status enum.TargetingTagStatus `json:"status,omitempty"`
	// Reason 定向标签不可用于定向的原因。
	// 仅当status为DISABLED时返回。
	// 枚举值：
	// OFFLINE：该定向标签已下线。
	// NOT_SUPPORTED：该定向标签未下线，但在所选的设置下不支持。
	Reason string `json:"reason,omitempty"`
}

// Geo 用于地域定向的定向标签的信息。
type Geo struct {
	// GeoID 地域ID或邮政编码ID。
	// 地域ID可用于行政区域定向，邮政编码ID则可用于邮政编码定向。
	// 当对应的 geo_type为ZIP_CODE时，本字段的值为邮政编码ID，可在调用/adgroup/create/，/adgroup/update/，或/ad/audience_size/estimate/时传给zipcode_ids字段。
	// 当对应的 geo_type非 ZIP_CODE时，本字段的值为地域ID，可在调用/adgroup/create/，/adgroup/update/，或/ad/audience_size/estimate/时传给location_ids 字段。
	GeoID string `json:"geo_id,omitempty"`
	// GeoType geo_id对应的地域类型。
	// 枚举值：COUNTRY（国家或地区），PROVINCE（省），CITY（城市/郡/县），DISTRICT（区/城市），DMA（指定市场区域），ZIP_CODE（邮政编码对应地域）。
	GeoType enum.GeoType `json:"geo_type,omitempty"`
	// Description 所定向地域的描述
	Description string `json:"description,omitempty"`
	// RegionCode 该定向地域所属的国家或地区的代码
	RegionCode string `json:"region_code,omitempty"`
	// ParentID 父层级地域的ID。
	// 若无对应父层级地域，则返回0
	ParentID string `json:"parent_id,omitempty"`
}

// ISP 用于互联网服务提供商定向的定向标签的信息。
type ISP struct {
	// ISPID 可定向的互联网服务提供商ID
	ISPID string `json:"isp_id,omitempty"`
	// LocationID 可与该isp_id同时使用的地域ID。
	// 仅支持国家或地区层级的地域ID。
	LocationID string `json:"location_id,omitempty"`
	// RegionCode isp_id 所对应的国家或地区代码
	RegionCode string `json:"region_code,omitempty"`
}
