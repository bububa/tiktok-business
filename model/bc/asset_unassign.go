package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// AssetUnassignRequest 撤销用户对资产的权限 API Request
type AssetUnassignRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// UserID 资产所属的用户 ID。
	UserID string `json:"user_id,omitempty"`
	// AssetType 资产类型。
	// 枚举值：
	// ADVERTISER：广告账户（广告主账号）。
	// CATALOG：商品库。
	// TIKTOK_SHOP：TikTok Shop（商店）。
	// PIXEL：Pixel。
	// LEAD：线索。
	// TT_ACCOUNT：TikTok 账号。此资产类型对应 BC_AUTH_TT身份，详情见认证身份
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// AssetID 资产ID。如果资产类型为广告账户，则传入广告账户ID。如果资产类型为商品库，则传入商品库ID
	AssetID string `json:"asset_id,omitempty"`
}

// Encode implements PostRequest
func (r *AssetUnassignRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
