package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// PartnerAssetDeleteRequest 取消资产分享 API Request
type PartnerAssetDeleteRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// PartnerID 要添加为合作伙伴的商务中心的 ID
	PartnerID string `json:"partner_id,omitempty"`
	// AssetType 资产类型。枚举值：ADVERTISER，CATALOG，TIKTOK_SHOP，STOREFRANT，PIXEL
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// AssetIDs 您想要取消分享的资产ID列表。目前仅支持传入一个表单库ID
	AssetIDs []string `json:"asset_ids,omitempty"`
	// AdvertiserRole 只有在asset_type 为ADVERTISER时有效。
	// 授予合作伙伴的广告账号角色。
	// 枚举值：
	// ADMIN：管理员。管理广告账号财务、设置和权限。 创建和编辑广告推广系列并查看投放效果数据。
	// OPERATOR：操作员。创建和编辑广告推广系列并查看投放效果数据。
	// ANALYST：分析师。查看广告和投放效果数据。
	// asset_type 为 ADVERTISER 时的默认值：ANALYST。
	AdvertiserRole enum.AdvertiserRole `json:"advertiser_role,omitempty"`
}

// Encode implements PostRequest
func (r *PartnerAssetDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
