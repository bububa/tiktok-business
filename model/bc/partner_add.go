package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// PartnerAddRequest 添加合作伙伴 API Request
type PartnerAddRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// PartnerID 要添加为合作伙伴的商务中心的 ID
	PartnerID string `json:"partner_id,omitempty"`
	// AssetType 要分享给合作伙伴商务中心的资产类型。
	// 枚举值：ADVERTISER（广告账号）
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// AssetIDs 要分享给合作伙伴商务中心的资产 ID 列表。
	// 最大数量：50
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
func (r *PartnerAddRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
