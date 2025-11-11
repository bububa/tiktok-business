package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// AssetAdminDeleteRequest 删除资产 API Request
type AssetAdminDeleteRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// AssetType 资产类型。
	// 枚举值：
	// ADVERTISER：广告账户（广告主账号）。
	// CATALOG：商品库。
	// TIKTOK_SHOP：TikTok Shop（商店）。
	// PIXEL：Pixel。
	// LEAD：线索。
	// TT_ACCOUNT：TikTok 账号。此资产类型对应 BC_AUTH_TT身份，详情见认证身份
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// AssetIDs 您要删除的资产ID列表。目前只支持传入一个ID。
	// 若 asset_type为 LEAD，则资产ID为表单库ID。
	// 若 asset_type为 TT_ACCOUNT，则资产ID为商务中心中TikTok账户的ID。
	AssetIDs []string `json:"asset_ids,omitempty"`
}

// Encode implements PostRequest
func (r *AssetAdminDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
