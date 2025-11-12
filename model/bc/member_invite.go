package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// MemberInviteRequest 邀请成员 API Request
type MemberInviteRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// Emails 邮箱地址列表。
	// 之前已邀请过的邮箱地址会被忽略。
	// 最大数量：50
	Emails string `json:"emails,omitempty"`
	// UserRole 用户在商务中心的基础角色。
	// 枚举值：
	// ADMIN：管理员可以访问商务中心内的所有功能。
	// STANDARD：标准成员只能对分配给他们的账户进行操作。
	UserRole enum.UserRole `json:"user_role,omitempty"`
	// AssetIDs 分配给邀请用户的广告账号 ID 列表。
	// 最大数量：50。
	// 若想获取商务中心中的广告账户，可使用/bc/asset/get/并将 asset_type 设置为 ADVERTISER。
	AssetIDs []string `json:"asset_ids,omitempty"`
	// AdvertiserRole 分配给邀请用户的广告账户权限。
	// 枚举值：
	// ADMIN：管理员。管理广告账号财务、设置和权限。 创建和编辑广告推广系列并查看投放效果数据。
	// OPERATOR：操作员。创建和编辑广告推广系列并查看投放效果数据。
	// ANALYST：分析师。查看广告和投放效果数据
	AdvertiserRole enum.AdvertiserRole `json:"advertiser_role,omitempty"`
	// ExtUserRole 用户在商务中心除基本角色外的扩展角色
	ExtUserRole *ExtUserRole `json:"ext_user_role,omitempty"`
}

// Encode implements PostRequest
func (r *MemberInviteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
