package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AssetGetRequest 获取资产列表 API Request
type AssetGetRequest struct {
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
	// Filtering 筛选条件
	Filtering *AssetGetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 取值范围：1-50。
	// 默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// AssetGetFilter 筛选条件
type AssetGetFilter struct {
	// UserID 资产所属的用户 ID。
	// 注意：user_id 和user_mail 不能同时传入。
	// 对于已经加入（BOUND ）的用户，传入user_id获取他们的资产。
	// 对于邀请中（PENDING ）的用户，传入user_email 获取资产。
	UserID string `json:"user_id,omitempty"`
	// UserEmail 资产所属用户的电子邮件。
	// 注意：
	// user_id 和user_mail 不能同时传入。
	// 对于已经加入（BOUND ）的用户，传入user_id获取他们的资产。
	// 对于邀请中（PENDING ）的用户，传入user_email 获取资产。
	// 仅在如下情形下用user_email作为筛选项：已经将asset_type 设置为ADVERTISER ，且邀请加入商务中心的请求仍在待确认状态
	UserEmail string `json:"user_email,omitempty"`
	// Keyword 您想要查找的资产的关键词
	Keyword string `json:"keyword,omitempty"`
}

// Encode implements GetRequest
func (r *AssetGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("asset_type", string(r.AssetType))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AssetGetResponse 获取资产列表 API Response
type AssetGetResponse struct {
	model.BaseResponse
	Data *AssetGetResult `json:"data,omitempty"`
}

type AssetGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 资产列表
	List []Asset `json:"list,omitempty"`
}

// Asset 资产
type Asset struct {
	// AssetType 资产类型。
	// 枚举值：
	// ADVERTISER：广告账户（广告主账号）。
	// CATALOG：商品库。
	// TIKTOK_SHOP：TikTok Shop（商店）。
	// PIXEL：Pixel。
	// LEAD：线索。
	// TT_ACCOUNT：TikTok 账号。此资产类型对应 BC_AUTH_TT身份，详情见认证身份
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// AssetID 资产 ID。
	// 若 asset_type 为 ADVERTISER，本字段代表广告账户 ID。
	// 若 asset_type 为 CATALOG，本字段代表商品库 ID。
	// 若 asset_type 为 TIKTOK_SHOP，本字段代表 TikTok Shop 的 ID。
	// 若 asset_type 为 PIXEL，本字段代表 Pixel ID。
	// 若 asset_type 为 LEAD，本字段代表表单库 ID。若想获取您有权访问的表单库的详情，可使用/page/library/get/。
	// 若 asset_type 为 TT_ACCOUNT，本字段代表商务中心中的 TikTok 账号的 ID。
	AssetID string `json:"asset_id,omitempty"`
	// AssetName 资产名称
	AssetName string `json:"asset_name,omitempty"`
	// RelationType 商务中心与资产的关系类型。
	// 枚举值:
	// OWNER_BC: 该资产由当前商务中心创建并拥有。
	// OWNER_PARTNER: 该资产由合作伙伴商务中心拥有。
	// OWNER_INDIVIDUAL: 该资产由一个个人账户拥有。
	RelationType enum.BcAssetRelationType `json:"relation_type,omitempty"`
	// RelationStatus 仅当 asset_type 为 ADVERTISER、PIXEL 或 TIKTOK_SHOP 时生效。
	// 商务中心与资产的关系状态。
	// 枚举值:
	// BOUND: 管理资产的申请已获拥有者批准。
	// UNBOUND: 管理资产的申请获拥有者批准后已被取消。
	// PENDING: 管理资产的申请正在等待拥有者批准。
	// REJECTED: 管理资产的申请已被拥有者拒绝。
	RelationStatus enum.BcAssetRelationStatus `json:"relation_status,omitempty"`
	// AdvertiserStatus 广告账户状态。
	// 枚举值可参考枚举值 - 广告主状态
	AdvertiserStatus enum.AdvertiserStatus `json:"advertiser_status,omitempty"`
	// AdvertiserAccountType 广告账户类型。
	// 枚举值：RESERVATION （合约广告账户），AUCTION（竞价广告账户）
	AdvertiserAccountType enum.ServiceType `json:"advertiser_account_type,omitempty"`
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
	// AdCreationEligible 仅当 asset_type 为 CATALOG 时返回本字段。
	// 商品库中含有广告样式所需的足量可用商品时，商品库是否可用于广告中。
	// 枚举值：
	// NOT_AVAILABLE: 商品库不可用于广告中。
	// AVAILABLE: 若商品库中含有广告样式所需的足量可用商品，商品库可用于广告中
	AdCreationEligible enum.Availability `json:"ad_creation_eligible,omitempty"`
	// StoreRole 商务中心用户对于 TikTok Shop 的权限。
	// 枚举值：
	// AD_PROMOTION: 广告推广。创建广告来推广此 TikTok Shop 的商品。
	// 仅当asset_type 为 TIKTOK_SHOP 时，本字段返回非 null 值。
	StoreRole enum.StoreRole `json:"store_role,omitempty"`
	// PixelCode Pixel 代码。
	// 仅当 asset_type 为 PIXEL 时，本字段返回非 null 的值
	PixelCode string `json:"pixel_code,omitempty"`
	// TTAcountRoles 商务中心用户对于 TikTok 账户的权限。
	// 枚举值:
	// POST：现有帖子权限，包括访问你自己的个人主页信息、访问你自己的公开视频和访问你自己的广告授权视频。
	// LIVE：直播视频权限（仅部分地区可以使用），包括访问你自己的个人主页信息和访问直播视频。
	// DIRECT_MESSAGE：私信权限，包括允许所有人向您发送私信，阅读和管理您收件箱中的私信，以及代表您向其他账号发送消息或进行互动。
	// 仅当asset_type为TT_ACCOUNT时，本字段返回非 null 值
	TTAcountRoles []enum.TTAccountRole `json:"tt_account_roles,omitempty"`
	// OwernBcName 拥有该资产的商务中心名称。
	// asset_type 为 TT_ACCOUNT 时，本字段的值将为 null
	OwnerBcName string `json:"owner_bc_name,omitempty"`
}
