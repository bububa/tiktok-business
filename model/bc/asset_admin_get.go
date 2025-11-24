package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// AssetAdminGetRequest 以管理员身份获取资产列表 API Request
type AssetAdminGetRequest struct {
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
	Filtering *AssetAdminGetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 取值范围：1-50。
	// 默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// AssetAdminGetFilter 筛选条件
type AssetAdminGetFilter struct {
	// AdvertiserShowStatus 仅当 asset_type 为 ADVERTISER 时生效。
	// 广告账户显示状态。
	// 枚举值可参考 枚举值 - 广告主显示状态。若不传入，将会返回所有显示状态的广告账户。
	AdvertiserShowStatus enum.AdvertiserShowStatus `json:"advertiser_show_status,omitempty"`
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
	RelationStatus enum.BcRelationStatus `json:"relation_status,omitempty"`
}

// Encode implements GetRequest
func (r *AssetAdminGetRequest) Encode() string {
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
