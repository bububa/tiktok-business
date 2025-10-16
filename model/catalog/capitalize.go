package catalog

import "github.com/bububa/tiktok-business/util"

// CapitalizeRequest 将商品库迁移至商务中心 API Request
type CapitalizeRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// AdvertiserID 要迁移的商品库目前所属的广告账户
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CatalogID 要迁移的商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *CapitalizeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
