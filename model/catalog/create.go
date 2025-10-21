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
	// 短剧商品库目前在测试阶段，是仅向受邀开发者开放的白名单功能。如需使用此功能，请联系您的 TikTok 销售代表，但无法保证申请就能入选。
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
