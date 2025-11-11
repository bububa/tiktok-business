package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// AssetAssignRequest 将资产分配给用户 API Request
type AssetAssignRequest struct {
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
	// AdvertiserRole 商务中心用户对于广告账户的权限。
	// 枚举值：
	// ADMIN：管理员。管理广告账号财务、设置和权限。 创建和编辑广告推广系列并查看投放效果数据。
	// OPERATOR：操作员。创建和编辑广告推广系列并查看投放效果数据。
	// ANALYST：分析师。查看广告和投放效果数据。
	// 仅当asset_type为ADVERTISER时，本字段返回非 null 值
	AdvertiserRole enum.AdvertiserRole `json:"advertiser_role,omitempty"`
	// CatalogRole 商务中心用户对于商品库的权限。
	// 枚举值：
	// ADMIN：商品库管理。可以通过自然内容和广告内容推广商品，也可以修改商品和商品库。
	// AD_PROMOTE：广告推广。使用商品库创建广告。
	// 仅当asset_type为CATALOG时，本字段返回非 null 值
	CatalogRole enum.CatalogRole `json:"catalog_role,omitempty"`
	// FormLibraryRole asset_type为LEAD时必填。
	// 用户对于表单库的权限/角色。目前只支持ADMIN
	FormLibraryRole string `json:"form_library_role,omitempty"`
	// StoreRole 商务中心用户对于 TikTok Shop 的权限。
	// 枚举值：
	// AD_PROMOTION: 广告推广。创建广告来推广此 TikTok Shop 的商品。
	// 仅当asset_type 为 TIKTOK_SHOP 时，本字段返回非 null 值。
	StoreRole enum.StoreRole `json:"store_role,omitempty"`
	// TTAcountRoles 商务中心用户对于 TikTok 账户的权限。
	// 枚举值:
	// POST：现有帖子权限，包括访问你自己的个人主页信息、访问你自己的公开视频和访问你自己的广告授权视频。
	// LIVE：直播视频权限（仅部分地区可以使用），包括访问你自己的个人主页信息和访问直播视频。
	// DIRECT_MESSAGE：私信权限，包括允许所有人向您发送私信，阅读和管理您收件箱中的私信，以及代表您向其他账号发送消息或进行互动。
	// 仅当asset_type为TT_ACCOUNT时，本字段返回非 null 值
	TTAcountRoles []enum.TTAccountRole `json:"tt_account_roles,omitempty"`
}

// Encode implements PostRequest
func (r *AssetAssignRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
