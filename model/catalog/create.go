package catalog

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建商品库 API Request
type CreateRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// Name 商品库名称。
	// 长度限制：128 字符。
	Name string `json:"name,omitempty"`
	// CatalogType Catalog type.
	// Enum values:
	// ECOM: E-commerce catalog.
	// HOTEL: hotel catalog.
	// FLIGHT: flight catalog.
	// DESTINATION: destination catalog.
	// ENTERTAINMENT: entertainment catalog.
	// AUTO_VEHICLE: Auto-Inventory catalog.
	// When catalog_type is AUTO_MODEL, see List of region codes and currencies for Auto-Inventory and Auto-Model catalogs to find out the supported region_code and currency values.
	// AUTO_MODEL: Auto-Model catalog.
	// When catalog_type is AUTO_MODEL, see List of region codes and currencies for Auto-Inventory and Auto-Model catalogs to find out the supported region_code and currency values.
	// MINI_SERIES: mini series catalog.
	// GENERIC: generic catalog.
	// ONLINE_TO_OFFLINE: online-to-offline catalog.
	// When catalog_type is ONLINE_TO_OFFLINE, see List of region codes and currencies for online-to-offline catalogs to find out the supported region_code and currency values.
	// Note:
	//
	// The entertainment catalog is currently an allowlist-only feature and is invitation-only because this catalog type is under Alpha Testing. If you would like to access it, please contact your TikTok representative. However, acceptance into the Alpha Test is not guaranteed.
	// The mini series catalog is currently an allowlist-only feature and is invitation-only because the catalog type is under testing. If you would like to access it, please contact your TikTok representative. However, acceptance into the test is not guaranteed.
	CatalogType enum.CatalogType `json:"catalog_type,omitempty"`
	// CatalogConf 商品库配置信息
	CatalogConf *CatalogConf `json:"catalog_conf,omitempty"`
}

// Encode implements PostRequest interface
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建商品库 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// CatalogID 商品库ID
		CatalogID string `json:"catalog_id,omitempty"`
	} `json:"data"`
}
