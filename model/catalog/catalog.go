package catalog

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Catalog 商品库
type Catalog struct {
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// CatalogName 商品库名称
	CatalogName string `json:"catalog_name,omitempty"`
	// CatalogType 商品库类型。
	// 枚举值：
	// ECOM：电商商品库。
	// HOTEL：酒店商品库。
	// FLIGHT：航班商品库。
	// DESTINATION：目的地商品库。
	// ENTERTAINMENT：娱乐商品库。
	// MINI_SERIES：短剧商品库。
	// 注意：
	//
	// 旅游类商品库（酒店商品库、航班商品库和目的地商品库）目前在 Alpha 测试阶段，是仅向受邀开发者开放的白名单功能。如需使用此功能，请联系您的TikTok销售代表，但无法保证申请就能入选。
	// 娱乐商品库目前在 Alpha 测试阶段，是仅向受邀开发者开放的白名单功能。如需使用此功能，请联系您的 TikTok 销售代表，但无法保证申请就能入选。
	CatalogType enum.CatalogType `json:"catalog_type,omitempty"`
	// AdCreationEligible 商品库中含有广告样式所需的足量可用商品时，商品库是否可用于广告中。
	// 枚举值：
	// NOT_AVAILABLE: 商品库不可用于广告中。
	// AVAILABLE: 若商品库中含有广告样式所需的足量可用商品，商品库可用于广告中。
	AdCreationEligible enum.Availability `json:"ad_creation_eligible,omitempty"`
	// CreateTime 创建时间
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// UpdateTime 更新时间
	UpdateTime model.DateTime `json:"update_time,omitzero"`
	// BcInfo 商务中心信息
	BcInfo *BcInfo `json:"bc_info,omitempty"`
	// CatalogConfi 商品库配置信息
	CatalogConf *CatalogConf `json:"catalog_conf,omitempty"`
}

// BcInfo 商务中心信息
type BcInfo struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// BcName 商务中心名称
	BcName string `json:"bc_name,omitempty"`
	// PictureURL 商务中心头像 URL。若无头像，返回空字符串
	PictureURL string `json:"picture_url,omitempty"`
}

// CatalogConf 商品库配置信息
type CatalogConf struct {
	// RegionCode 定向地区的二字代码。
	// 请先调用/catalog/available_country/get/接口并传入bc_id，返回中的region_codes字段会列出支持的地区列表，这里region_code 的值需为支持地区列表中的地区代码。
	// 若想了解某个地区代码对应的地区，请查看附录-地区代码。
	RegionCode string `json:"region_code,omitempty"`
	// Country 定向国家
	Country string `json:"country,omitempty"`
	// Currency 货币单位，详见附录-币种。
	Currency string `json:"currency,omitempty"`
	// Channel 商品库创建渠道。枚举值: PARTNER: 第三方合作伙伴; CLIENT: 直客广告主
	Channel string `json:"channel,omitempty"`
}
