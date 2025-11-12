package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// PartnerAssetGetRequest 获取合作伙伴资产 API Request
type PartnerAssetGetRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// PartnerID 合作伙伴商务中心ID
	PartnerID string `json:"partner_id,omitempty"`
	// AssetType 资产类型。
	// 枚举值：
	// ADVERTISER：广告账户（广告主账号）。
	// CATALOG：商品库。
	// TIKTOK_SHOP：TikTok Shop（商店）。
	// PIXEL：Pixel。
	// LEAD：线索。
	// TT_ACCOUNT：TikTok 账号。此资产类型对应 BC_AUTH_TT身份，详情见认证身份
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// ShareType 分享类型。枚举值：
	// SHARED: 合作伙伴分享给您的资产。
	// SHARING: 您分享给合作伙伴的资产。
	ShareType enum.ShareType `json:"share_type,omitempty"`
	// Filtering 使用关键词筛选
	Filtering *PartnerAssetGetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// PartnerAssetGetFilter 使用关键词筛选
type PartnerAssetGetFilter struct {
	// Keyword 您希望用作筛选条件的关键词
	Keyword string `json:"keyword,omitempty"`
}

// Encode implements GetRequest
func (r *PartnerAssetGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("partner_id", r.PartnerID)
	values.Set("asset_type", string(r.AssetType))
	values.Set("share_type", string(r.ShareType))
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
